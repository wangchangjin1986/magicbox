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
* (3)TimeWheel，时间轮，拷贝自 https://www.luozhiyun.com/archives/444
