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

opensea目前就这一个事件 OrderFulfilled

```
OrderFulfilled (bytes32 orderHash, index_topic_1 address offerer, index_topic_2 address zone, address recipient, tuple[] offer, tuple[] consideration)
```

其中，offer是下面这个结构体数组：

```
struct OfferItem {
    ItemType itemType;
    address token;
    uint256 identifierOrCriteria;
    uint256 startAmount;
    uint256 endAmount;
}
```

consideration是这个结构体数组：

```
struct ConsiderationItem {
    ItemType itemType;
    address token;
    uint256 identifierOrCriteria;
    uint256 startAmount;
    uint256 endAmount;
    address payable recipient;
}
```


其中有一笔交易触发两个这个事件的，是因为价格匹配机制，案例交易如下：

https://etherscan.io/tx/0x465f34dcee32751d2a34253efccab8486b127b208e7f288ff1321d5487ce9e12#eventlog

在这笔交易中，OrderFulfilled 被 emit 了两次，但是应该表示的是同一笔交易。两个不同的交易方向的 order 被匹配，然后被成交。这种交易往往会触发一个OrdersMatched 事件，其中的 orderHashed 会和其中一个 OrderFulfilled 的 orderHash 参数一样：

```
OrdersMatched (bytes32[] orderHashes)
```

## 如何解析Event

要在 Go 中精确地解析 OpenSea 的 `OrderFulfilled` 事件的 `OfferItem` 和 `ConsiderationItem` 结构体，我们首先需要了解这些结构体在 Solidity 中的定义，并将其转换为相应的 Go 结构体。基于您提供的 Solidity 结构体定义，我们可以创建对应的 Go 结构体，并使用 `go-ethereum` 库来解析这些结构体。

在开始之前，请确认 `ItemType` 在 Solidity 中的具体定义，例如它可能是一个枚举类型。在 Go 中，我们可以将其简单地表示为 `uint8` 或 `uint256`，这取决于 `ItemType` 的具体实现。

以下是更新后的代码示例：

```go
package main

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// 假设 ItemType 是一个 uint8 类型
type ItemType uint8

// OfferItem 结构体
type OfferItem struct {
	ItemType            ItemType
	Token               common.Address
	IdentifierOrCriteria *big.Int
	StartAmount         *big.Int
	EndAmount           *big.Int
}

// ConsiderationItem 结构体
type ConsiderationItem struct {
	ItemType            ItemType
	Token               common.Address
	IdentifierOrCriteria *big.Int
	StartAmount         *big.Int
	EndAmount           *big.Int
	Recipient           common.Address
}

type OrderFulfilledEvent struct {
	OrderHash     [32]byte
	Offerer       common.Address
	Zone          common.Address
	Recipient     common.Address
	Offer         []OfferItem
	Consideration []ConsiderationItem
}

func main() {
	// ABI 字符串
	abiData := "" // ABI 字符串，您需要从合约的 ABI 中获取
	parsedAbi, err := abi.JSON(strings.NewReader(abiData))
	if err != nil {
		panic(err)
	}

	// 日志数据，假设您已经有了
	var logData []byte

	// 解析日志
	event := new(OrderFulfilledEvent)
	err = parsedAbi.UnpackIntoInterface(event, "OrderFulfilled", logData)
	if err != nil {
		panic(err)
	}

	fmt.Printf("OrderFulfilled Event: %+v\n", event)
}

// 注意：
// - ItemType 的具体类型取决于其在 Solidity 中的定义。
// - ABI 字符串和日志数据需要您从合约的 ABI 和具体的事件日志中获取。
// - 确保已安装 go-ethereum 库。
```

### 注意事项

- **ItemType 类型**：在 Solidity 中，如果 `ItemType` 是一个枚举，您需要知道其确切的整数表示。在这个例子中，我假设它是 `uint8` 类型。请根据实际情况调整。
- **ABI 和日志数据**：您需要提供 `OrderFulfilled` 事件的具体 ABI 字符串，并从以太坊区块链中获取实际的日志数据。
- **数组解析**：请注意，`Offer` 和 `Consideration` 是数组类型。`go-ethereum` 的 ABI 解析库应该能够处理这种复杂的数据结构。

这段代码提供了一个框架，用于解析 OpenSea 的 `OrderFulfilled` 事件日志数据。根据您合约的具体实现，可能需要做进一步的调整。


# LooksRare

## TakerBid

```
event TakerBid(
    NonceInvalidationParameters nonceInvalidationParameters,
    address bidUser, // taker (initiates the transaction)
    address bidRecipient, // taker (receives the NFT)
    uint256 strategyId,
    address currency,
    address collection,
    uint256[] itemIds,
    uint256[] amounts,
    address[2] feeRecipients,
    uint256[3] feeAmounts
);

struct NonceInvalidationParameters {
    bytes32 orderHash;
    uint256 orderNonce;
    bool isNonceInvalidated;
}
```

### 案例交易

https://etherscan.io/tx/0x9a09678e0b81dc128db8b2075781e2207b644c29fc0adcd661fa552b9b418de8#eventlog

### 如何解析Event

要解析 LooksRare 的 `TakerBid` 事件日志数据，我们首先需要定义与 Solidity 结构体相对应的 Go 结构体，并使用 `go-ethereum` 库来解析日志数据。以下是一个 Go 语言示例，展示了如何实现这一解析过程：

#### 定义 Go 结构体

```go
package main

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// NonceInvalidationParameters 结构体
type NonceInvalidationParameters struct {
	OrderHash          [32]byte
	OrderNonce         *big.Int
	IsNonceInvalidated bool
}

// TakerBidEvent 结构体
type TakerBidEvent struct {
	NonceInvalidationParameters NonceInvalidationParameters
	BidUser                     common.Address
	BidRecipient                common.Address
	StrategyId                  *big.Int
	Currency                    common.Address
	Collection                  common.Address
	ItemIds                     []*big.Int
	Amounts                     []*big.Int
	FeeRecipients               [2]common.Address
	FeeAmounts                  [3]*big.Int
}

func main() {
	// ABI 字符串
	abiData := "" // ABI 字符串，您需要从合约的 ABI 中获取
	parsedAbi, err := abi.JSON(strings.NewReader(abiData))
	if err != nil {
		panic(err)
	}

	// 日志数据，假设您已经有了
	var logData []byte

	// 解析日志
	event := new(TakerBidEvent)
	err = parsedAbi.UnpackIntoInterface(event, "TakerBid", logData)
	if err != nil {
		panic(err)
	}

	fmt.Printf("TakerBid Event: %+v\n", event)
}

// 注意：
// - ABI 字符串和日志数据需要您从合约的 ABI 和具体的事件日志中获取。
// - 确保已安装 go-ethereum 库。
```

#### 注意事项

- **ABI 字符串**：您需要从 LooksRare 合约的 ABI 中提取相关的 ABI 字符串。
- **日志数据**：您需要从以太坊区块链上获取 `TakerBid` 事件的实际日志数据。
- **数组和固定大小数组**：`itemIds` 和 `amounts` 是动态大小的数组，而 `feeRecipients` 和 `feeAmounts` 是固定大小的数组。这在 Go 结构体中分别对应于切片 (`[]*big.Int`) 和固定大小的数组 (`[2]common.Address` 和 `[3]*big.Int`)。
- **安装 `go-ethereum` 库**：此代码依赖于 `go-ethereum` 库，您需要确保已经正确安装了这个库。

通过上述步骤和代码，您可以解析 LooksRare 的 `TakerBid` 事件日志数据，以便进一步处理和分析。

## TakerAsk

```
event TakerAsk(
    NonceInvalidationParameters nonceInvalidationParameters,
    address askUser, // taker (initiates the transaction)
    address bidUser, // maker (receives the NFT)
    uint256 strategyId,
    address currency,
    address collection,
    uint256[] itemIds,
    uint256[] amounts,
    address[2] feeRecipients,
    uint256[3] feeAmounts
);

struct NonceInvalidationParameters {
    bytes32 orderHash;
    uint256 orderNonce;
    bool isNonceInvalidated;
}
```

### 案例交易

暂未找到

### 解析Event

要解析 LooksRare 的 `TakerAsk` 事件日志数据，我们将遵循与解析 `TakerBid` 事件类似的步骤。首先，定义与 Solidity 结构体相对应的 Go 结构体，然后使用 `go-ethereum` 库来解析日志数据。以下是解析 `TakerAsk` 事件的 Go 代码示例：

#### 定义 Go 结构体

```go
package main

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// NonceInvalidationParameters 结构体
type NonceInvalidationParameters struct {
	OrderHash          [32]byte
	OrderNonce         *big.Int
	IsNonceInvalidated bool
}

// TakerAskEvent 结构体
type TakerAskEvent struct {
	NonceInvalidationParameters NonceInvalidationParameters
	AskUser                     common.Address
	BidUser                     common.Address
	StrategyId                  *big.Int
	Currency                    common.Address
	Collection                  common.Address
	ItemIds                     []*big.Int
	Amounts                     []*big.Int
	FeeRecipients               [2]common.Address
	FeeAmounts                  [3]*big.Int
}

func main() {
	// ABI 字符串
	abiData := "" // ABI 字符串，您需要从合约的 ABI 中获取
	parsedAbi, err := abi.JSON(strings.NewReader(abiData))
	if err != nil {
		panic(err)
	}

	// 日志数据，假设您已经有了
	var logData []byte

	// 解析日志
	event := new(TakerAskEvent)
	err = parsedAbi.UnpackIntoInterface(event, "TakerAsk", logData)
	if err != nil {
		panic(err)
	}

	fmt.Printf("TakerAsk Event: %+v\n", event)
}

// 注意：
// - ABI 字符串和日志数据需要您从合约的 ABI 和具体的事件日志中获取。
// - 确保已安装 go-ethereum 库。
```

#### 注意事项

- **ABI 字符串**：您需要从 LooksRare 合约的 ABI 中提取相关的 ABI 字符串。
- **日志数据**：您需要从以太坊区块链上获取 `TakerAsk` 事件的实际日志数据。
- **数组和固定大小数组**：`itemIds` 和 `amounts` 是动态大小的数组，而 `feeRecipients` 和 `feeAmounts` 是固定大小的数组。这在 Go 结构体中分别对应于切片 (`[]*big.Int`) 和固定大小的数组 (`[2]common.Address` 和 `[3]*big.Int`)。
- **安装 `go-ethereum` 库**：此代码依赖于 `go-ethereum` 库，您需要确保已经正确安装了这个库。

通过上述步骤和代码，您可以解析 LooksRare 的 `TakerAsk` 事件日志数据，以便进行进一步的处理和分析。


# Element

## ERC721SellOrderFilled

```
ERC721SellOrderFilled (bytes32 orderHash, address maker, address taker, uint256 nonce, address erc20Token, uint256 erc20TokenAmount, tuple[] fees, address erc721Token, uint256 erc721TokenId)

struct FeeItem {
    address recipient; // 接收者地址
    uint256 amount;    // 接收的费用金额
}
```

### 解析 Event data

为了解析 Element 的 `ERC721SellOrderFilled` 事件日志数据，我们将遵循前面讨论的步骤。首先，我们定义与 Solidity 结构体相对应的 Go 结构体，然后使用 `go-ethereum` 库来解析日志数据。以下是更新后的 Go 代码示例：

#### 定义 Go 结构体

```go
package main

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// FeeItem 结构体
type FeeItem struct {
	Recipient common.Address
	Amount    *big.Int
}

// ERC721SellOrderFilledEvent 结构体
type ERC721SellOrderFilledEvent struct {
	OrderHash       [32]byte
	Maker           common.Address
	Taker           common.Address
	Nonce           *big.Int
	Erc20Token      common.Address
	Erc20TokenAmount *big.Int
	Fees            []FeeItem
	Erc721Token     common.Address
	Erc721TokenId   *big.Int
}

func main() {
	// ABI 字符串
	abiData := "" // ABI 字符串，您需要从合约的 ABI 中获取
	parsedAbi, err := abi.JSON(strings.NewReader(abiData))
	if err != nil {
		panic(err)
	}

	// 日志数据，假设您已经有了
	var logData []byte

	// 解析日志
	event := new(ERC721SellOrderFilledEvent)
	err = parsedAbi.UnpackIntoInterface(event, "ERC721SellOrderFilled", logData)
	if err != nil {
		panic(err)
	}

	fmt.Printf("ERC721SellOrderFilled Event: %+v\n", event)
}

// 注意：
// - ABI 字符串和日志数据需要您从合约的 ABI 和具体的事件日志中获取。
// - 确保已安装 go-ethereum 库。
```

#### 注意事项

- **ABI 字符串**：您需要从 Element 合约的 ABI 中提取相关的 ABI 字符串。
- **日志数据**：您需要从以太坊区块链上获取 `ERC721SellOrderFilled` 事件的实际日志数据。
- **安装 `go-ethereum` 库**：此代码依赖于 `go-ethereum` 库，您需要确保已经正确安装了这个库。

这段代码提供了一个框架，用于解析 Element 的 `ERC721SellOrderFilled` 事件日志数据。根据您合约的具体实现，可能需要做进一步的调整。

## ERC721BuyOrderFilled

```
ERC721BuyOrderFilled (bytes32 orderHash, address maker, address taker, uint256 nonce, address erc20Token, uint256 erc20TokenAmount, tuple[] fees, address erc721Token, uint256 erc721TokenId)

struct FeeItem {
    address recipient; // 接收者地址
    uint256 amount;    // 接收的费用金额
}
```

### 解析 Event data

要解析 Element 的 `ERC721BuyOrderFilled` 事件日志数据，我们将采用与之前相同的方法。首先定义与 Solidity 结构体相对应的 Go 结构体，然后使用 `go-ethereum` 库来解析日志数据。假设 `fees` 数组的结构与之前相同（包含接收者地址和金额），我们可以使用相同的 `FeeItem` 结构体定义。以下是解析 `ERC721BuyOrderFilled` 事件的 Go 代码示例：

#### 定义 Go 结构体

```go
package main

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// FeeItem 结构体
type FeeItem struct {
	Recipient common.Address
	Amount    *big.Int
}

// ERC721BuyOrderFilledEvent 结构体
type ERC721BuyOrderFilledEvent struct {
	OrderHash       [32]byte
	Maker           common.Address
	Taker           common.Address
	Nonce           *big.Int
	Erc20Token      common.Address
	Erc20TokenAmount *big.Int
	Fees            []FeeItem
	Erc721Token     common.Address
	Erc721TokenId   *big.Int
}

func main() {
	// ABI 字符串
	abiData := "" // ABI 字符串，您需要从合约的 ABI 中获取
	parsedAbi, err := abi.JSON(strings.NewReader(abiData))
	if err != nil {
		panic(err)
	}

	// 日志数据，假设您已经有了
	var logData []byte

	// 解析日志
	event := new(ERC721BuyOrderFilledEvent)
	err = parsedAbi.UnpackIntoInterface(event, "ERC721BuyOrderFilled", logData)
	if err != nil {
		panic(err)
	}

	fmt.Printf("ERC721BuyOrderFilled Event: %+v\n", event)
}

// 注意：
// - ABI 字符串和日志数据需要您从合约的 ABI 和具体的事件日志中获取。
// - 确保已安装 go-ethereum 库。
```

#### 注意事项

- **ABI 字符串**：您需要从 Element 合约的 ABI 中提取相关的 ABI 字符串。
- **日志数据**：您需要从以太坊区块链上获取 `ERC721BuyOrderFilled` 事件的实际日志数据。
- **安装 `go-ethereum` 库**：此代码依赖于 `go-ethereum` 库，您需要确保已经正确安装了这个库。

这段代码提供了一个框架，用于解析 Element 的 `ERC721BuyOrderFilled` 事件日志数据。根据您合约的具体实现，可能需要做进一步的调整。
