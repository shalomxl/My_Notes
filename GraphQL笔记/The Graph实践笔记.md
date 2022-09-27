# GraphQL实践笔记

## 1. 安装Graph CLI

```sh
yarn global add @graphprotocol/graph-cli

或

npm install -g @graphprotocol/graph-cli
```
以上都是全局安装，也可以按项目安装。

## 2. 初始化Graph项目

### 基于已部署的合约创建新的Subgraph
执行命令：
```
graph init
```
虽然这个命令有很多可选项，但是填上后，依然会有一个交互式的界面来确认各个选项。例如我执行下面命令：
```
graph init --product host-service --from-contract 0x6A5ebE5fd34B8078F6C8cE1C1E6d5Ea378502B6d --network chapel --index-events Created --contract-name GameDaoFixedNFT
```
依然会出现如下交互界面：
```
✔ Product for which to initialize · hosted-service
✔ Subgraph name · connectorgamefi/market
✔ Directory to create the subgraph in · market
✔ Ethereum network · chapel
✔ Contract address · 0x6A5ebE5fd34B8078F6C8cE1C1E6d5Ea378502B6d
✖ Failed to fetch ABI from Etherscan: request to https://api-chapel.etherscan.io/api?module=contract&action=getabi&address=0x6A5ebE5fd34B8078F6C8cE1C1E6d5Ea378502B6d failed, reason: getaddrinfo ENOTFOUND api-chapel.etherscan.io
✔ ABI file (path) · abis/GameDaoFixedNFT.json
✔ Contract Name · GameDaoFixedNFT
```
可以看到，获取合约ABI的时候，默认使用etherscan.io的api来获取，如果我们的合约是bsc上的，或者非以太系列的网络上的，那么就要指定本地ABI路径。

### 从现有的Subgraph中创建Subgraph
目前貌似只有 subgraph-studio 支持这个功能。

### 添加数据源到自己的Subgraph
同样只有 subgraph-studio 支持这个功能。

## 3. 项目目录解析

初始化好的项目目录如下：
```
market
├── abis
│   └── GameDaoFixedNFT.json
├── generated
│   ├── GameDaoFixedNFT
│   └── schema.ts
├── package.json
├── schema.graphql
├── src
│   └── mapping.ts
├── subgraph.yaml
└── yarn.lock
```
其中最重要的几个文件：
* `subgraph.yaml`: 项目的配置文件，指定项目的数据源等
* `schema.graphql`: 定义存储的数据结构
* `mapping.ts`: 是 AssemblyScript Mappings，数据如何从合约的结构转换成`schema.graphql`指定的数据结构，是在这里指定的，是我们自定义开发的地方。

## 4. Subgraph.yaml解析

Subgraph.yaml中所有内容的清单文档：[graph-node/subgraph-manifest.md at master · graphprotocol/graph-node](https://github.com/graphprotocol/graph-node/blob/master/docs/subgraph-manifest.md)
### 简单案例分析
这是从官方文档中摘取的Subgraph.yaml案例：
```yaml
specVersion: 0.0.4
description: Gravatar for Ethereum
repository: https://github.com/graphprotocol/example-subgraphs
schema:
  file: ./schema.graphql
dataSources:
  - kind: ethereum/contract
    name: Gravity # 在generated文件夹中使用
    network: mainnet
    source:
      address: '0x2E645469f354BB4F5c8a05B3b30A929361cf77eC'
      abi: Gravity
      startBlock: 6175244
    mapping:
      kind: ethereum/events
      apiVersion: 0.0.6
      language: wasm/assemblyscript
      entities:
        - Gravatar
      abis:
        - name: Gravity
          file: ./abis/Gravity.json
      eventHandlers:
        - event: NewGravatar(uint256,address,string,string)
          handler: handleNewGravatar
        - event: UpdatedGravatar(uint256,address,string,string)
          handler: handleUpdatedGravatar
      callHandlers: 
        - function: createGravatar(string,string)
          handler: handleCreateGravatar
      blockHandlers:
        - handler: handleBlock
        - handler: handleBlockWithCall
          filter:
            kind: call
      file: ./src/mapping.ts
```
其中的各个字段需要与项目其他文件关联上。

#### startBlock
指定需要从哪个区块开始提取合约数据。一般情况，会指定合约创建时的区块号，这样就可以避免去扫描与当前合约无关的区块。
#### callHandlers
通过 Parity 节点客户端的 tracing API 来实现追踪合约的调用。因此，只有支持 Parity 节点客户端的网络才可以使用 callHandlers。
#### blockHandlers
每产生一个新区块，就会调用blockHandlers下指定的函数。指定了过滤器，可以在特定的区块下调用。
```
blockHandlers:
  - handler: handleBlock
  - handler: handleBlockWithCall
    filter:
      kind: call
```
上面内容中，指定了两个blockHandler。第一个`handleBlock`是每个新区块都会调用，第二个`handleBlockWithCall`是只有包含当前合约调用的区块，才会触发，因为加了filter。

##### blockHandler如何编写
在`mapping.ts`中，handlerBlock将只会接受一个参数`ethereum.Block`，`ethereum.Block`相关信息在Mapping记录中🔗[Writing Mappings](brain://8LK7iE02KEKvfPNAXvcc4Q/WritingMappings)
```
import { ethereum } from '@graphprotocol/graph-ts'

export function handleBlock(block: ethereum.Block): void {
  let id = block.hash
  let entity = new Block(id)
  entity.save()
}
```
#### eventHandlers
eventHandlers中的两个小功能：指定topic0 和 接收交易receipt
##### 指定topic0
```
eventHandlers:
  - event: LogNote(bytes4,address,bytes32,bytes32,uint256,bytes)
    topic0: '0x644843f351d3fba4abcd60109eaff9f54bac8fb8ccf0bab941009c21df21cf31'
    handler: handleGive
```
指定了topic0后，需要签名和topic0同时满足，才会触发handler。
##### 接收交易receipt
```
eventHandlers:
  - event: NewGravatar(uint256,address,string,string)
    handler: handleNewGravatar
    receipt: true
```
打开开关后，在eventHandler中的event参数中，将可以通过`event.receipt`获得交易的receipt。
### 以上dataSource顺序规则
1. 在同一个block中，所有的event首先按照区块内的交易顺序进行排序
2. 同一个交易中的多个event，按照`Subgraph.yaml`中定义的先后顺序进行排序

## 5. 定义数据存储实体

在`schema.graphql`中定义实体结构的时候，最好避免直接将事件作为实体。应该结合查询的业务逻辑，将实体作为一个符合业务的对象来设计。一个实体可能是由多个事件数据组成。

有关定义实体的细节，看这里[创建子图 - The Graph Docs](https://thegraph.com/docs/zh/developing/creating-a-subgraph/#heading-10)

## 6. 编写mappint.ts

文档链接：[创建子图 - The Graph Docs](https://thegraph.com/docs/zh/developing/creating-a-subgraph/#heading-25)｜[assemblyscript-api](https://thegraph.com/docs/zh/developing/assemblyscript-api/)

## 7. 构建和部署

生成code
```sh
graph codegen
```

构建
```sh
graph build
```

部署
```sh
graph deploy --product hosted-service connectorGamefi/Gameloot
```

