## HAL 9000

hal9000是golang微服务的一个简单例子，涉及内容:

* go-kit框架
* grpc
* 服务注册发现etcd

项目名取自于电影**2001太空漫游**中的人工智能机器人hal9000

### go kit
Gokit是一系列的工具的集合，能够帮助你快速的构建健壮的，可靠的，可维护的微服务。它提供了一系列构建微服务的成熟的模式和惯用法，背后有着一群经验丰富的开发者支持，并且已经在生产环境中被广泛的使用。

#### gokit的架构

gokit不同于传统的MVC的框架，它只是一系列工具的组合，他有着自己的层次结构，主要有三层，分别是transport，endpoint和service层。

* transport层：这是一个抽象的层级，对应真实世界中的http/grpc/thrift等，通过gokit你可以在同一个微服务中同时支持http和grpc。

* endpoint层：endpoint层对应于controller中的action，主要是实现安全逻辑的地方，如果你要同时支持http和grpc，那么你将需要创建两个方法同时路由到同一个endpoint。

* service层：service是具体的业务逻辑实现的层级，在这里，你应该使用接口，并且通过实现这些接口来构建具体的业务逻辑。一个service通常聚合了多个endpoints，在service层，你应该使用clean architecture或者六边形模型，也就是说你的service层不需要知道enpoint以及transport层的具体实现，也不需要关心具体的http头部或者grpc的错误状态码。

* middleware：middleware实现了装饰器模式，通过middleware你可以包装你的service或者endpoint，通常你需要构建一个middleware链来实现如日志，rate limit，负载均衡和分布式追踪。

#### 缺点

* 太复杂：添加api的开销高，且大多数代码是重复的。要添加一个api，需要做:
  * 声明一个interface，并定义相关的方法
  * 实现这个interface
  * 实现endpoint的工厂方法
  * 实现transport方法
  * 实现request encoder，request decoder， response encoder response decoder
  * 把endpoint添加到server
  * 把endpoint添加到client
* 代码难理解
  * 虽然业务层，端点层，传输层分层清晰，有很好的抽象性，但是代码难以阅读，interface漫天飞