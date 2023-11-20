## 目录结构
```
├──  bootstrap                 项目启动引导服务
│   └──  gateway_intit         网关初始化
├──  cmd                       项目启动服务
│   └──  main.go               项目入口文件         
│       ├──  ymal         ymal基本配置文件 （建议使用env结合阿波罗服务进行配置文件发布）
│       └──  lua          lua脚本 构建网关限流熔断器 以及分布式锁
├── pkg                  功能工具包库（不可使用外部的自建服务包，不可写入业务代码！！！）
│   ├── app              请求返回处理包库，比如格式返回和请求参数校验等
│   ├── cache            缓存库，比如redis连接池的初始化
│   ├── conf             配置库，配置初始化
│   ├── enums            标准常量库，即类型、以及返回code的常量设置包
│   ├── file             文件处理库
│   ├── logging          日志处理库
│   └── util             工具库，类似php的helpers
├── runtime              运行输出文件夹
│   └── logs             程序里面添加的输出日志
├── service              服务请求数据层，封装对各种第三方服务的请求（比如权益黑名单等）处理库（只可引入PKG工具包）
├── vendor               加载的库包，类似php的
```

#### 项目启动
````

go build -o  h5-gateway.exe main.go
h5-gateway.exe start --port 8080 --env local
````

#### 改造
该框架基于gin go1.18以下1.15以上版本构建

-- 基于lua构建的限流 熔断器
-- 监控组件 可监控内部阻塞协程调用栈，防止协程泄漏导致系统异常
-- 服务提供者 可以体面上大部分的中间的调用端 如:mysql（基于gorm) ，redis , elastic ，kafka
-- grpc功能(待构建中)
-- 较好的日志记录
-- 可构建基础的网关服务和基于gin框架搭建web服务

