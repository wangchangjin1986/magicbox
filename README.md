# magicbox
魔法盒，积累多个通用的工具包
### routine pool
goroutine pool,目前支持
* （1）普通的pool，即每次add()都会临时创建一个goroutine
* （2）固定大小的pool(FixPool）,创建时，根据参数size初始化对应的goroutine，bufferSize为接收job的缓冲区大小，单位为Job的个数
### msync
自定义的同步功能结构体和函数，目前有：
* （1）mwaitgroup，主要是由于官方sync包中的state为私有，不能返回当前尚未done的routine的数量，因此override了一个，能够返回当前未完成的routine的个数
### structure
常用的数据结构
* (1)PriorityQueue，优先级队列，按照指定优先级出队，目前是按照priority的值从小到大的顺序出队列
* (2)BlockingDelayQueue，无界阻塞队列，基于PriorityQueue实现的阻塞队列，到执行时间的队列元素统一输出到chan C中，如果没有到期的元素，则队列处于阻塞状态，直到下一个到期元素的时间到达；（目前待优化，缩短sleep状态时新增元素处理的等待时长）
* (3)TimeWheel，时间轮，拷贝自https://www.luozhiyun.com/archives/444
### util
* (1)Interceptor, 提供通过反射实现的类似AOP切面能力，在函数前后提供增强能力

### go语言守则
#### Go 箴言
        不要通过共享内存进行通信，通过通信共享内存
        并发不是并行
        管道用于协调；互斥量（锁）用于同步
        接口越大，抽象就越弱
        利用好零值
        空接口 interface{} 没有任何类型约束
        Gofmt 的风格不是人们最喜欢的，但 gofmt 是每个人的最爱
        允许一点点重复比引入一点点依赖更好
        系统调用必须始终使用构建标记进行保护
        必须始终使用构建标记保护 Cgo
        Cgo 不是 Go
        使用标准库的 unsafe 包，不能保证能如期运行
        清晰比聪明更好
        反射永远不清晰
        错误是值
        不要只检查错误，还要优雅地处理它们
        设计架构，命名组件，（文档）记录细节
        文档是供用户使用的
        不要（在生产环境）使用 panic()
Author: Rob Pike See more: https://go-proverbs.github.io/
#### Go 之禅
        每个 package 实现单一的目的
        显式处理错误
        尽早返回，而不是使用深嵌套
        让调用者处理并发（带来的问题）
        在启动一个 goroutine 时，需要知道何时它会停止
        避免 package 级别的状态
        简单很重要
        编写测试以锁定 package API 的行为
        如果你觉得慢，先编写 benchmark 来证明
        适度是一种美德
        可维护性
Author: Dave Cheney See more: https://the-zen-of-go.netlify.com/