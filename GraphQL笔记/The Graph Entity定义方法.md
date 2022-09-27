# Defining Entities

## Entity基本定义方法
每个类型都需要注释`@entity`，像如下例子一样：
```
type Gravatar @entity(immutable: true) {
    id: Bytes!
    owner: Bytes
    displayName: String
    imageUrl: String
    accepted: Boolean
}
```
若想要类型在今后的升级中不可变，可以使用`@entity(immutable: true)`进行注释。不可变的entity读写速度要快得多，所以要尽可能多的使用。

与定制SQL表一样，需要考虑字段是否是必填字段，必填字段需要在类型后边加上`!`，就像上面的id字段。

每个Entity必须要有一个唯一的必填字段：id。id的类型可以是 `Bytes!` or `String!`。但是`Bytes!`要快很多，所以尽量使用`Bytes!`。可以发现的唯一主键与SQL不同，Entity的主键是Bytes类型的，它的生成规则是在`mapping.ts`中DIY的，文档中有推荐一种规则：

id的生成：`left.id.concatI32(count)`，使用左侧相同的部分加上数量。

## 字段类型选择
支持的字段类型：

|Type|Description|
|`Bytes`|Byte array, represented as a hexadecimal string. Commonly used for Ethereum hashes and addresses.|
|`String`|Scalar for `string` values. Null characters are not supported and are automatically removed.|
|`Boolean`|Scalar for `boolean` values.|
|`Int`|The GraphQL spec defines `Int` to have a size of 32 bytes.|
|`BigInt`|Large integers. Used for Ethereum's `uint32`, `int64`, `uint64`, ..., `uint256` types. Note: Everything below `uint32`, such as `int32`, `uint24` or `int8` is represented as `i32`.|
|`BigDecimal`|`BigDecimal` High precision decimals represented as a significand and an exponent. The exponent range is from −6143 to +6144. Rounded to 34 significant digits.|
## Entity之间的关系
Entity之间的关系是单向的，可以定义两个互逆的单向关系来模拟双向关系。

### 一对一
互相添加可选字段作为对方的类型。
```
type Transaction @entity(immutable: true) {
  id: Bytes!
  transactionReceipt: TransactionReceipt
}
type TransactionReceipt @entity(immutable: true) {
  id: Bytes!
  transaction: Transaction
}
```
相互有个可选字段对应。

### 一对多
```
type Token @entity(immutable: true) {
  id: Bytes!
}

type TokenBalance @entity {
  id: Bytes!
  amount: Int!
  token: Token!
}
```
子集类型中，定一个字段存储父级。
#### 定义反向查找
这是一个虚拟字段，不增加额外存储，从指定字段派生出来的。使用 `@derivedFrom(field: "token")`
```
type Token @entity(immutable: true) {
  id: Bytes!
  tokenBalances: [TokenBalance!]! @derivedFrom(field: "token")
}

type TokenBalance @entity {
  id: Bytes!
  amount: Int!
  token: Token!
}
```
### 多对多
是建立中间Entity来保存Entity之间的关系。
```
type Organization @entity {
  id: Bytes!
  name: String!
  members: [UserOrganization!]! @derivedFrom(field: "organization")
}

type User @entity {
  id: Bytes!
  name: String!
  organizations: [UserOrganization!] @derivedFrom(field: "user")
}

type UserOrganization @entity {
  id: Bytes! # Set to `user.id.concat(organization.id)`
  user: User!
  organization: Organization!
}
```
查询方式：
```
query usersWithOrganizations {
  users {
    organizations {
      # this is a UserOrganization entity
      organization {
        name
      }
    }
  }
}
```
