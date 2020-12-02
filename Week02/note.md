## GO语言实践—error

> - 问题：
>   - golang 中的异常机制及工作原理
>   - panic是什么？什么情况下使用panic
>   - 如何正确的抛出error，以及处理异常
>   - golang 中error的类型结构
>   - 封装实现一个标准的异常处理组件
>   - 不同服务之间请求的的异常处理



## 课程笔记

> - error返回指针的意义：防止异常定义内容一致
> - panic是一个fatal error，当抛出panic，则代表系统出现了某种问题，直接中断程序运行。注意不要catch panic之后，再尝试重试处理，例如数据库或资源初始化、栈溢出、索引越界等，根据依赖关系（强依赖 、弱依赖）来觉得是否panic
> - Request-Driven：在框架或服务顶层捕捉panic，进行错误信息输出，终止当前请求，不关闭当前进程。保证后续请求能够正常访问。
> - Error处理方式：
>   - Sentinel Error：预定义的特定错误，处理措施不灵活，调用方必须进行相等性判断，当想提供更多的错误上下文时，就无法进行相等性判断。
>     - 缺点：
>       - 容易成为API公共部分
>       - 在两个包之间创建了依赖
>     - 结论：尽可能避免使用Sentinel error
>   - Error types：使用switch断言错误
>   - Opaque errors：最灵活的错误处理策略，不透明错误处理。只能够知道是否出现错误
>   - Opaque优化操作（内部断言）：Assert errors for behaviour , no type
> - Error Handle：
>   - Indented flow is for errors：避免错误处理代码循环缩进
>   - Eliminate error handling by eliminating errors
>   - Wrap Error
>     - 使用wrap的场景：
>       - 
>     - 不应该使用的场景
>     - 使用总结：
>       - 自己开发的基础库(kit)，只返回根错误，不使用error.wrap返回
>       - 如果函数不打算处理错误，用足够的上下文通过wrap.error返回到调用堆栈中，并返回出去
>       -  函数内如果已经处理过错误，则不应该继续将错误往上抛
> - Go 1.13 errors
>   - errors标准库新增Is、As、UnWrap方法，用于处理异常
>   - fmt增加  %w 占位符，用于返回wrap error信息
>   - 存在的问题：未记录堆栈信息，可以使用Pkg/errors包中的wrap方法来处理异常
> - Go 2 Error
> - Go 泛型及替代工具：go generate 



## 课程作业

> - 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？



### 课程资料引用

> - GO学习文档：https://golang.org/doc/effective_go.html
> - Warp Error处理：https://github.com/pkg/errors
> - 课程文章资料：https://shimo.im/docs/GYvDrQT8qW8RgkGY
> - 错误资料引用：
>   - https://dave.cheney.net/2012/01/18/why-go-gets-exceptions-right
>     https://dave.cheney.net/2015/01/26/errors-and-exceptions-redux
>     https://dave.cheney.net/2014/11/04/error-handling-vs-exceptions-redux
>     https://rauljordan.com/2020/07/06/why-go-error-handling-is-awesome.html
>     https://morsmachine.dk/error-handling
>     https://blog.golang.org/error-handling-and-go
>     https://www.ardanlabs.com/blog/2014/10/error-handling-in-go-part-i.html
>     https://www.ardanlabs.com/blog/2014/11/error-handling-in-go-part-ii.html
>     https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully
>     https://commandcenter.blogspot.com/2017/12/error-handling-in-upspin.html
>     https://blog.golang.org/errors-are-values
>     https://dave.cheney.net/2016/06/12/stack-traces-and-the-errors-package
>   - https://www.ardanlabs.com/blog/2017/05/design-philosophy-on-logging.html
>     https://crawshaw.io/blog/xerrors
>     https://blog.golang.org/go1.13-errors
>     https://medium.com/gett-engineering/error-handling-in-go-53b8a7112d04
>     https://medium.com/gett-engineering/error-handling-in-go-1-13-5ee6d1e0a55c
> - Go 2的错误检查展望：https://go.googlesource.com/proposal/+/master/design/29934-error-values.md

