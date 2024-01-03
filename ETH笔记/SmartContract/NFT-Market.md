# Blur

Blur的合约 0xb2ecfe4e4d61f8790bbb9de2d1259b9e2410cea5 一共有4个 event：

```
event Execution(
    Transfer transfer,
    bytes32 orderHash,
    uint256 listingIndex,
    uint256 price,
    FeeRate makerFee,
    Fees fees,
    OrderType orderType
);

event Execution721Packed(
    bytes32 orderHash,
    uint256 tokenIdListingIndexTrader,
    uint256 collectionPriceSide
);

event Execution721TakerFeePacked(
    bytes32 orderHash,
    uint256 tokenIdListingIndexTrader,
    uint256 collectionPriceSide,
    uint256 takerFeeRecipientRate
);

event Execution721MakerFeePacked(
    bytes32 orderHash,
    uint256 tokenIdListingIndexTrader,
    uint256 collectionPriceSide,
    uint256 makerFeeRecipientRate
);
```

## Execution

### 案例交易

暂未找到

### 如何解析Event

要用 Go 语言解析 `Execution` 事件的日志数据，首先需要理解这个事件的数据结构，以及如何在 Go 中处理相应的二进制数据。根据您提供的结构体定义，我们可以构建对应的 Go 结构体，并使用以太坊的 Go 库来解析日志数据。

以下是完整的示例代码，展示了如何解析 `Execution` 事件日志：

```go
package main

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// 定义 Execution 事件中使用的结构体
type Transfer struct {
	Trader    common.Address
	Id        *big.Int
	Amount    *big.Int
	Collection common.Address
	AssetType uint8 // 假设 AssetType 是 uint8 类型
}

type FeeRate struct {
	Recipient common.Address
	Rate      uint16
}

type Fees struct {
	ProtocolFee FeeRate
	TakerFee    FeeRate
}

type Execution struct {
	TransferData Transfer
	OrderHash    [32]byte
	ListingIndex *big.Int
	Price        *big.Int
	MakerFee     FeeRate
	FeesData     Fees
	OrderType    uint8
}

func main() {
	// 假设您有一个日志数据 logData
	var logData []byte // 这里应该是从区块链上获取的实际日志数据

	// 构建 ABI
	abiData := "" // 这里应填入 Execution 事件的 ABI 字符串
	parsedAbi, err := abi.JSON(strings.NewReader(abiData))
	if err != nil {
		panic(err)
	}

	// 解析日志
	event := new(Execution)
	err = parsedAbi.UnpackIntoInterface(event, "Execution", logData)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Execution Event: %+v\n", event)
}

// 注意：
// - 这段代码假设了 AssetType 是 uint8 类型。您需要根据实际情况调整。
// - 您需要提供 Execution 事件的 ABI 字符串和实际的日志数据。
// - go-ethereum 库提供了 UnpackIntoInterface 函数来解析日志数据。
```

#### 注意事项：

1. **ABI 字符串**：您需要提供 `Execution` 事件的具体 ABI 字符串。这通常可以从合约的 ABI 定义中获取。

2. **日志数据**：您需要提供实际的日志数据。这通常是从以太坊区块链上获取的原始日志数据。

3. **AssetType 类型**：这里我假设 `AssetType` 是 `uint8` 类型。如果它是不同的类型，请相应地调整代码。

4. **go-ethereum 库**：这个代码示例使用了 go-ethereum 库中的 `abi` 包来解析日志。您需要确保安装了 go-ethereum 库。

通过这段代码，您可以将 `Execution` 事件的日志数据解析成 Go 中的结构体实例。这样，您就能以程序化的方式访问这些数据了。

## Execution721Packed

### 案例交易

https://etherscan.io/tx/0xd6d09ba6a5cf3baeb3f3ee8ae7f6fef1ce6d7306089e06d1eceea412ba25d69a#eventlog

### 如何解析Event

要将您通过 Go 语言获取的十进制数 `452312848714801535733105423629729706736873637691851605814500335336667401207` 解析为 `orderType`、`price` 和 `collection`，您需要按照这个数值在 Solidity 函数中被打包的方式来进行逆向操作。这个数值是通过位操作组合 `orderType`（占据最高的8位）、`price`（占据随后的160位）和 `collection`（占据最低的160位）得到的。

在 Go 语言中，您可以使用位移和位掩码操作来提取这些值。首先，定义一个 `uint256` 类型的变量来存储这个大数值。然后，按照下面的步骤进行：

1. **提取 `collection`**：
   - `collection` 位于最低的 160 位。您可以通过对这个大数值应用一个只有最低 160 位为 1 的位掩码来提取 `collection`。

2. **提取 `price`**：
   - `price` 位于紧接着 `collection` 之上的 160 位。您可以先将整个数值向右移动 160 位，然后应用一个 160 位的位掩码来提取 `price`。

3. **提取 `orderType`**：
   - `orderType` 位于最高的 8 位。您可以将整个数值向右移动 248 位来获取 `orderType`。

以下是用 Go 语言实现这个逻辑的示例代码：

```go
package main

import (
	"fmt"
	"math/big"
)

func main() {
	// 将提供的数值转换为big.Int
	value := new(big.Int)
	value.SetString("452312848714801535733105423629729706736873637691851605814500335336667401207", 10)

	// 定义位掩码
	mask160 := new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 160), big.NewInt(1))
	mask8 := new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 8), big.NewInt(1))

	// 提取collection (最低160位)
	collection := new(big.Int).And(value, mask160)

	// 提取price (中间160位)
	price := new(big.Int).Rsh(value, 160)
	price.And(price, mask160)

	// 提取orderType (最高8位)
	orderType := new(big.Int).Rsh(value, 248)
	orderType.And(orderType, mask8)

	fmt.Printf("orderType: %d\n", orderType)
	fmt.Printf("price: %d\n", price)
	fmt.Printf("collection: 0x%x\n", collection)
}
```

在这个代码中，`value` 是您提供的大数值。我们使用 `big.Int` 来处理这个数值，因为它超出了 Go 语言内建整型能表示的范围。我们定义了两个位掩码 `mask160` 和 `mask8` 来帮助提取相应位数的值。然后，我们分别提取 `collection`、`price` 和 `orderType`。注意，`collection` 作为地址，在最后输出时采用了十六进制格式。

## Execution721TakerFeePacked

同上

## Execution721MakerFeePacked

同上

# Opensea

```
OrderFulfilled (bytes32 orderHash, index_topic_1 address offerer, index_topic_2 address zone, address recipient, tuple[] offer, tuple[] consideration)
```

价格匹配机制：

https://etherscan.io/tx/0x465f34dcee32751d2a34253efccab8486b127b208e7f288ff1321d5487ce9e12#eventlog

在这笔交易中，OrderFulfilled 被 emit 了两次，但是应该表示的是同一笔交易。两个不同的交易方向的 order 被匹配，然后被成交。这种交易往往会触发一个事件：

```
OrdersMatched (bytes32[] orderHashes)
```

##