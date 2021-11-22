# sync.Pool笔记

## 学习资料  

[sync包英文文档](https://pkg.go.dev/sync#Pool)  

[sync包中文文档](https://studygolang.com/pkgdoc)  

[码农桃花源深度解析sync.Pool](https://www.cnblogs.com/qcrao-2018/p/12736031.html)  

[深度分析 Golang sync.Pool 底层原理](https://www.cyhone.com/articles/think-in-sync-pool/)

## 知识点大纲，画出自己的知识点地图

## 实践案例

### 初学小案例

```golang
type GoPOOL struct {
	a string
	b int
}

var goPoolSync = sync.Pool{
	New: func() interface{} { return new(GoPOOL) },
}

func TestSyncPool(t *testing.T) {
	goPoolSync.Put(&GoPOOL{
		a: "aaaa",
		b: 3,
	})
	pool := goPoolSync.Get().(*GoPOOL)

	fmt.Println(pool.a, pool.b)
}
```

### 适用的场景是什么？

#### 业务场景

无

#### 技术场景

当多个 goroutine 都需要创建同⼀个对象的时候，如果 goroutine 数过多，导致对象的创建数⽬剧增，进⽽导致 GC 压⼒增大。形成 “并发⼤ －> 占⽤内存⼤ －> GC 缓慢 －> 处理并发能⼒降低 －> 并发更⼤”这样的恶性循环。在这个时候，需要有⼀个对象池，每个 goroutine 不再⾃⼰单独创建对象，⽽是从对象池中获取出⼀个对象（如果池中已经有的话）。

## 底层技术和原理

### sync.Pool要解决什么问题？

Golang 1.3版本引入，避免因频繁建立临时对象所带来的消耗以及对 GC 造成的压力。

### 使用到了哪些底层技术？技术组成部分有哪些？

### 解决问题的原理是什么？关键的实现方案是什么？（sync.Pool是如何解决这个问题的？这个问题有更好的方法吗？）

### 与其他实现方案的不同在哪里，侧重点是什么？（sync.Pool有什么优缺点？trade-off是什么？）

## 拓展思维

### sync.Pool能够让你联想到什么？

### 能否从中提取出一些通用的事物，放之四海而皆准的事物？举出几个案例

### 给出以上所有的答案的反面观点，是否有漏洞？有没有不好的地方？

## 总结归纳（这里才是一篇正式出炉的概括性文章）
