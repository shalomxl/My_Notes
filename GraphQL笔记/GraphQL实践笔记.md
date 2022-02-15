# GraphQL实践笔记

## Host Service

Host Service 是 the graph 提供的免费部署自定义子图的服务。在未来的两三年后，将会不再提供，原有的子图将会移植到将来的 The Graph 网络中。

### 使用基本流程

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