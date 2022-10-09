## Order
|字段名|解释|类型|
|---|---|---|
|created_date|用户listing时签名的时间|string|
|closing_date|挂单的过期时间|string / null|
|listing_time|created_date对应的Unix时间戳|number|
|expiration_time|closing_date对应的Unix时间戳|number|
|order_hash|:{这个Hash计算方式是怎样的？:(style="color:#cc0000"):}:|string / null|
|protocol_data|卖家挂单时签名的数据|[Order Parameters](brain://api.thebrain.com/zBJh6lmfMEWrNhIe90B2qA/dLp2dyRVNkuuj4wvOhfcaA/OrderParameters)|
|protocol_address|这里是Seaport合约地址|string / null|
|maker|cellection创建者信息|[Account](brain://api.thebrain.com/zBJh6lmfMEWrNhIe90B2qA/5P5TYq2orkeraweslJ_dPA/Account)|
|taker||[Account](brain://api.thebrain.com/zBJh6lmfMEWrNhIe90B2qA/5P5TYq2orkeraweslJ_dPA/Account)|
|current_price|当前价格，单位wei|string|
|maker_fees|这里包含OS平台费和版税|[Fees](brain://api.thebrain.com/zBJh6lmfMEWrNhIe90B2qA/sV0MDCBrOkmCHdY9H0nBWQ/Fees)|
|taker_fees||[Fees](brain://api.thebrain.com/zBJh6lmfMEWrNhIe90B2qA/sV0MDCBrOkmCHdY9H0nBWQ/Fees)|
|side||string|
|order_type|order类型|string|
|canceled|是否取消|bool|
|finalized|是否成交|bool|
|marked_invalid||bool|
|client_signature|这段签名数据与protocol_data里的签名数据一样|string / null|
|relay_id||string|
|maker_asset_bundle|||
|taker_asset_bundle|||

案例数据：
```js
{
	"created_date":"2022-06-16T02:25:30.630424", 
	"closing_date":"2022-06-19T02:22:29",
	"listing_time":1655346330,
	"expiration_time":1655605349,
	"order_hash":"0xd3e311f64897d2c55cf4c2dc64926485a6ae568b40e2a5a3ec60664c92ee4a47",
	"protocol_data":Object{...},
	"protocol_address":"0x00000000006c3852cbef3e08e8df289169ede581",
	"maker":Object{...},
	"taker":null,
	"current_price":"10000000000000000",
	"maker_fees":Array[0],
	"taker_fees":Array[1],
	"side":"bid",
	"order_type":"basic",
	"cancelled":false,
	"finalized":false,
	"marked_invalid":true,
	"client_signature":"0xd97b9c15d86c6e3eef536fe8a46263ce1dd9aea91a53a26a7cda5fa54ce9e577277484faab8bbbfa7a65ad0b1033016e3cfa180e810620d7c544df27e233083d1b",
	"relay_id":"T3JkZXJWMlR5cGU6NzQxNjE5",
	"maker_asset_bundle":Object{...},
	"taker_asset_bundle":Object{...}
}
```

## Account

|字段名|解释|类型|
|---|---|---|
|user|用户编号|number|
|profile_img_url|用户头像图片链接|string|
|address|用户地址|string|
|config|附加信息|string|

案例数据
```
{
	"user":3241,
	"profile_img_url":"https://storage.googleapis.com/opensea-static/opensea-profile/30.png",
	"address":"0xbccc2073adfc46421308f62cfd9868df00d339a8",
	"config":""
}
```

## Order Parameters

这段数据就是offerer在挂单时签名的数据，以及offerer的签名
## parameters
|字段名|解释|类型|
|---|---|---|
|offerer|||
|zone|||
|zone_hash|||
|start_time|||
|end_time|||
|order_type|||
|salt|||
|conduitKey|||
|nonce|||
|offer||[Offer](brain://api.thebrain.com/zBJh6lmfMEWrNhIe90B2qA/jXNszlC86ES0BrFmqTY-EQ/Offer)|
|consideration||[Consideration](brain://PMRc9FO0hkW2v1fEq-afaQ/Consideration)|


## signature


## 案例数据
```
{
    "parameters":{
        "offerer":"0xbccc2073adfc46421308f62cfd9868df00d339a8",
        "offer":[
            {
                "itemType":2,
                "token":"0xBF12A3181B95674f0380C34732cA50e9c3476575",
                "identifierOrCriteria":"0",
                "startAmount":"1",
                "endAmount":"1"
            }
        ],
        "consideration":[
            {
                "itemType":0,
                "token":"0x0000000000000000000000000000000000000000",
                "identifierOrCriteria":"0",
                "startAmount":"950000000000000000",
                "endAmount":"950000000000000000",
                "recipient":"0xBCcC2073ADfC46421308f62cfD9868dF00D339a8"
            },
            {
                "itemType":0,
                "token":"0x0000000000000000000000000000000000000000",
                "identifierOrCriteria":"0",
                "startAmount":"25000000000000000",
                "endAmount":"25000000000000000",
                "recipient":"0x8De9C5A032463C561423387a9648c5C7BCC5BC90"
            },
            {
                "itemType":0,
                "token":"0x0000000000000000000000000000000000000000",
                "identifierOrCriteria":"0",
                "startAmount":"25000000000000000",
                "endAmount":"25000000000000000",
                "recipient":"0xBCcC2073ADfC46421308f62cfD9868dF00D339a8"
            }
        ],
        "startTime":"1655359388",
        "endTime":"1657951388",
        "orderType":2,
        "zone":"0x00000000E88FE2628EbC5DA81d2b3CeaD633E89e",
        "zoneHash":"0x0000000000000000000000000000000000000000000000000000000000000000",
        "salt":"5321472219075503",
        "conduitKey":"0x0000007b02230091a7ed01230072f7006a004d60a8d4e71d599b8104250f0000",
        "totalOriginalConsiderationItems":3,
        "counter":0
    },
    "signature":"0x66e29de5c1df4a513e1d8a786cfcc668ef11eb860961bd0976a1794f9598c9096e4f88b64b72854925d03d7a97bb644b78ea49fdcbb8f71a6ea68f5a730660991c"
}
```