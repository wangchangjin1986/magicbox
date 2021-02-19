# magicbox
魔法盒，积累多个通用的工具包
### routine pool
goroutine pool,目前支持
    （1）普通的pool，即每次add()都会临时创建一个goroutine
    （2）固定大小的pool(FixPool）,创建时，根据参数size初始化对应的goroutine，bufferSize为接收job的缓冲区大小，单位为Job的个数
### msync
自定义的同步功能结构体和函数，目前有：
    （1）mwaitgroup，主要是由于官方sync包中的state为私有，不能返回当前尚未done的routine的数量，因此override了一个，能够返回当前未完成的routine的个数

    