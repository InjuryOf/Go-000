## 课程大纲

> - Goroutine
> - Memory model
> - Package Sync：
> - Chan
> - Package Context



## 课程笔记

> - Goroutine：
>   - 并行：多个处理器同时处理多个任务
>   - 并发：单个处理器分配不同时间片处理多个任务
>   - 用户态线程：
>   - 内核态线程：
>   - 使用原则：
>     - 必须管理启动的Goroutine的生命周期
>     - 由调用者控制Goroutine的启动
>     - 调用者必须知道Goroutine什么时候结束
>     - Goroutine超时控制
> - Memory model
>   - Happen-Before：程序顺序执行
>   - Memory Reordering ：CPU重排
>   - cacheline
>   - 内存屏障
>   - MESI
>   - Memory order
>   - Store buffer
> - Package Sync：
> - Chan
> - Package Context
> - 注意点：
>   - log.Fatal内部会调用os.Exit，会导致defer不执行



## 课程问题



## 课程资料引用

> - Go 内存模型：https://www.jianshu.com/p/5e44168f47a3
> - https://golang.org/ref/mem