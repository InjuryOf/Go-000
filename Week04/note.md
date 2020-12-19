## 课程大纲

>- 工程项目结构
>- API设计
>- 配置管理
>- 包管理
>- 测试



## 课前疑问

> - 如何制定比较通用的目录结构？
> - API如何设计?
> - 配置管理的相对优解？
> - 项目的包管理？
> - 如果对项目进行完成的测试（白盒测试，性能测试）？

![image-20201210222715238](/Users/yanglei/Library/Application Support/typora-user-images/image-20201210222715238.png)

## 课程内容

> - 微服务中app的服务类型：
>   - Interface：对外的BFF服务
>   - service：对内的微服务
>   - Job：流式任务处理的服务
>   - Admin：区别于serfvice，面向运营侧的服务，数据权限不一致
>   - Task：定时任务，类似cronjob
> - Api设计
>   - Api Errors
>     - 错误传播：service error -> grpc error -> service  error
> - 配置
>   - function option
> - test
>   - Subtest + gomock



## 概念 & 知识点

> - 服务树
> - 阿兹卡班（定时任务调度平台）
> - Lint
> - Wrie

## 引用 & 扩展资料

> https://blog.golang.org/wire
>
> https://zouyx.github.io/posts/2016/2/21/Java-FlameGraph%E7%81%AB%E7%84%B0%E5%9B%BE.html
>
> https://zhuanlan.zhihu.com/p/110453784
>
> https://entgo.io/docs/getting-started/
>
> https://golang.org/pkg/testing/#T.Cleanup

