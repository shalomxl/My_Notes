# IPFS 笔记

## 节点配置文件

来自官方：

```json
{
  // 节点身份信息
  "Identity": {
    "PeerID": "12D3KooWQzEtKqG5c8tppyvTYfRoYDRi8XFdNGxahR1MN57V542m",
    "PrivKey": "CAESQHX7c8uLWJdfrN1N6nythJi8/qtPb20Xh+k6fEvegJjD4WbggXSoAzWQuCT1uht5hQgDsZXntMGieQJ7Ka7TuXI="
  },
  // 存储配置
  "Datastore": {
    "StorageMax": "10GB",
    "StorageGCWatermark": 90,
    "GCPeriod": "1h",
    "Spec": {
      "mounts": [
        {
          "child": {
            "path": "blocks",
            "shardFunc": "/repo/flatfs/shard/v1/next-to-last/2",
            "sync": true,
            "type": "flatfs"
          },
          "mountpoint": "/blocks",
          "prefix": "flatfs.datastore",
          "type": "measure"
        },
        {
          "child": {
            "compression": "none",
            "path": "datastore",
            "type": "levelds"
          },
          "mountpoint": "/",
          "prefix": "leveldb.datastore",
          "type": "measure"
        }
      ],
      "type": "mount"
    },
    "HashOnRead": false,
    "BloomFilterSize": 0
  },
  // 节点网络通信配置
  "Addresses": {
    "Swarm": [
      "/ip4/0.0.0.0/tcp/4001",
      "/ip6/::/tcp/4001",
      "/ip4/0.0.0.0/udp/4001/quic",
      "/ip6/::/udp/4001/quic"
    ],
    "Announce": [],
    "AppendAnnounce": [],
    "NoAnnounce": [],
    "API": "/ip4/127.0.0.1/tcp/5001",
    "Gateway": "/ip4/127.0.0.1/tcp/8080"
  },
  // 文件系统挂载配置
  "Mounts": {
    "IPFS": "/ipfs",
    "IPNS": "/ipns",
    "FuseAllowOther": false
  },
  // LibP2P 节点发现配置
  "Discovery": {
    "MDNS": {
      "Enabled": true,
      "Interval": 10
    }
  },
  "Routing": {
    "Type": "dht"
  },
  // IPNS配置
  "Ipns": {
    "RepublishPeriod": "",
    "RecordLifetime": "",
    "ResolveCacheSize": 128
  },
  // 中继节点配置
  "Bootstrap": [
    "/dnsaddr/bootstrap.libp2p.io/p2p/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN",
    "/dnsaddr/bootstrap.libp2p.io/p2p/QmQCU2EcMqAqQPR2i9bChDtGNJchTbq5TbXJJ16u19uLTa",
    "/dnsaddr/bootstrap.libp2p.io/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb",
    "/dnsaddr/bootstrap.libp2p.io/p2p/QmcZf59bWwK5XFi76CZX8cbJ4BhTzzA3gU1ZjYZcYW3dwt",
    "/ip4/104.131.131.82/tcp/4001/p2p/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ",
    "/ip4/104.131.131.82/udp/4001/quic/p2p/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ"
  ],
  // HTTP 网关配置
  "Gateway": {
    "HTTPHeaders": {
      "Access-Control-Allow-Headers": [
        "X-Requested-With",
        "Range",
        "User-Agent"
      ],
      "Access-Control-Allow-Methods": [
        "GET"
      ],
      "Access-Control-Allow-Origin": [
        "*"
      ]
    },
    "RootRedirect": "",
    "Writable": false,
    "PathPrefixes": [],
    "APICommands": [],
    "NoFetch": false,
    "NoDNSLink": false,
    "PublicGateways": null
  },
  // 节点允许的API配置
  "API": {
    "HTTPHeaders": {}
  },
  // P2P 蜂巢网络配置
  "Swarm": {
    "AddrFilters": null,
    "DisableBandwidthMetrics": false,
    "DisableNatPortMap": false,
    "RelayClient": {},
    "RelayService": {},
    "Transports": {
      "Network": {},
      "Security": {},
      "Multiplexers": {}
    },
    "ConnMgr": {
      "Type": "basic",
      "LowWater": 600,
      "HighWater": 900,
      "GracePeriod": "20s"
    }
  },
  "AutoNAT": {},
  "Pubsub": {
    "Router": "",
    "DisableSigning": false
  },
  "Peering": {
    "Peers": null
  },
  "DNS": {
    "Resolvers": {}
  },
  "Migration": {
    "DownloadSources": [],
    "Keep": ""
  },
  "Provider": {
    "Strategy": ""
  },
  "Reprovider": {
    "Interval": "12h",
    "Strategy": "all"
  },
  // 实验性功能配置
  "Experimental": {
    "FilestoreEnabled": false,
    "UrlstoreEnabled": false,
    "GraphsyncEnabled": false,
    "Libp2pStreamMounting": false,
    "P2pHttpProxy": false,
    "StrategicProviding": false,
    "AcceleratedDHTClient": false
  },
  "Plugins": {
    "Plugins": null
  },
  "Pinning": {
    "RemoteServices": {}
  },
  "Internal": {}
}
```