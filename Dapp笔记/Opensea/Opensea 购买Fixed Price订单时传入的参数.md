Opensea主网Seaport合约地址：0x00000000006c3852cbEf3e08E8dF289169EdE581

接下Fixed Price要调用的函数是`fulfillBasicOrder`，函数签名为`fb0f3ee1`。要绕过前端，就需要理解合约中`fulfillBasicOrder`入参的意义。

`fulfillBasicOrder`合约代码整理如下：
```solidity
struct AdditionalRecipient {
    uint256 amount;
    address payable recipient;
}
struct BasicOrderParameters {
    // calldata offset
    address considerationToken; // 0x24
    uint256 considerationIdentifier; // 0x44
    uint256 considerationAmount; // 0x64
    address payable offerer; // 0x84
    address zone; // 0xa4
    address offerToken; // 0xc4
    uint256 offerIdentifier; // 0xe4
    uint256 offerAmount; // 0x104
    BasicOrderType basicOrderType; // 0x124 enum
    uint256 startTime; // 0x144
    uint256 endTime; // 0x164
    bytes32 zoneHash; // 0x184
    uint256 salt; // 0x1a4
    bytes32 offererConduitKey; // 0x1c4
    bytes32 fulfillerConduitKey; // 0x1e4
    uint256 totalOriginalAdditionalRecipients; // 0x204
    AdditionalRecipient[] additionalRecipients; // 0x224
    bytes signature; // 0x244
    // Total length, excluding dynamic array data: 0x264 (580)
}

function fulfillBasicOrder(BasicOrderParameters calldata parameters)
        external
        payable
        returns (bool fulfilled);
```

计算Selector时，最终取Hash的部分：
```solidity
fulfillBasicOrder((address,uint256,uint256,address,address,address,uint256,uint256,uint8,uint256,uint256,bytes32,uint256,bytes32,bytes32,uint256,(uint256,address)[],bytes))
```

实际交易的入参样例与注释：
```json
{
    "BasicOrderParameters":{
        // offerer想要的token，0地址代表ETH
        "considerationToken":"0x0000000000000000000000000000000000000000",
        // 对应order中的 consideration第零项的 identifierOrCriteria 字段
        "considerationIdentifier":"0",
        // offerer想要的资产数量
        "considerationAmount":"950000000000000000",
        // offerer地址，可能是NFT卖家，也可能是offer提供者(这个offer是报价，与代码中的不一样)
        "offerer":"0xBCcC2073ADfC46421308f62cfD9868dF00D339a8",
        // 对应order中的 zone 字段
        "zone":"0x00000000E88FE2628EbC5DA81d2b3CeaD633E89e",
        // offer的合约地址
        "offerToken":"0xBF12A3181B95674f0380C34732cA50e9c3476575",
        // 对应order中的 offer identifierOrCriteria 字段
        "offerIdentifier":"0",
        // offer资产数量
        "offerAmount":"1",
        // 这里对应 ConsiderationEnums.sol 中 BasicOrderType 的值
	    // 普通购买是 ETH_TO_ERC721_FULL_RESTRICTED 值为2
	    // 接offer是 ERC721_TO_ERC20_FULL_RESTRICTED 值为18
        "basicOrderType":"2",
        "startTime":"1655359388",
        "endTime":"1657951388",
        "zoneHash":"0x0000000000000000000000000000000000000000000000000000000000000000",
        "salt":"5321472219075503",
        // conduitKey
        "offererConduitKey":"0x0000007b02230091a7ed01230072f7006a004d60a8d4e71d599b8104250f0000",
        // conduitKey
        "fulfillerConduitKey":"0x0000007b02230091a7ed01230072f7006a004d60a8d4e71d599b8104250f0000",
        // fees数量
        "totalOriginalAdditionalRecipients":"2",
        // 各个fee信息，可以从consideration后两个元素中获取
        "additionalRecipients":[
            {
                "AdditionalRecipient":{
                    "amount":"50000000000000000",
				    // fee的接受者
                    "recipient":"0x8De9C5A032463C561423387a9648c5C7BCC5BC90"
                }
            },
            {
                "AdditionalRecipient":{
                    "amount":"50000000000000000",
                    "recipient":"0xBCcC2073ADfC46421308f62cfD9868dF00D339a8"
                }
            }
        ],
        // 挂单者签名
        "signature":"0x3c993f7ea36450f6f97324632073caf1b69867319612954dfba57b73e7a547244693751685ba9670d2ba64d7b00301da0e15b990f246438e9418eff1ae9af4651c"
    }
}
```