# 如何检查graph是否健康？


* Find your `Deployment ID` ("Qm....") 
* Go to https://graphiql-online.com/ 
* Enter this API: `https://api.thegraph.com/index-node/graphql` 
* Run Query: 
```
{
  indexingStatuses(subgraphs: ["Qm..."]) {
	subgraph
	synced
	health
	entityCount
	fatalError {
	  handler
	  message
	  deterministic
	  block {
		hash
		number
	  }
	}
	chains {
	  chainHeadBlock {
		number
	  }
	  earliestBlock {
		number
	  }
	  latestBlock {
		number
	  }
	}
  }
}
``` 