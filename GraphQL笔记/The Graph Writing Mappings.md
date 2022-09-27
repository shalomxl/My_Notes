# Writing Mappings
mapping.ts是使用Typescript的子集AssemblyScript语法来编写的，用来将以太坊的数据转换为Entity。
## event的handler
```ts
import { NewGravatar, UpdatedGravatar } from '../generated/Gravity/Gravity'
import { Gravatar } from '../generated/schema'

export function handleNewGravatar(event: NewGravatar): void {
  let gravatar = new Gravatar(event.params.id)
  gravatar.owner = event.params.owner
  gravatar.displayName = event.params.displayName
  gravatar.imageUrl = event.params.imageUrl
  gravatar.save()
}

export function handleUpdatedGravatar(event: UpdatedGravatar): void {
  let id = event.params.id
  let gravatar = Gravatar.load(id)
  if (gravatar == null) {
    gravatar = new Gravatar(id)
  }
  gravatar.owner = event.params.owner
  gravatar.displayName = event.params.displayName
  gravatar.imageUrl = event.params.imageUrl
  gravatar.save()
}
```
每个handle方法都需要与yaml文件中定义的名字相同。并且需要接收一个名为`event`的参数，这个参数类型与合约的事件名是一一对应的。


mapping api 文档：[AssemblyScript API - The Graph Docs](https://thegraph.com/docs/en/developing/assemblyscript-api/)

