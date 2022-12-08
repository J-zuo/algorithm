### rabbitmq的使用场景
+ 秒杀业务，可靠性比较高

### rabbitmq基本架构组成

#### docker安装的rabbitmq,在客户端无法连接的原因？
```bash
# 启动容器
docker run -d -p 5672:5672 -p 15672:15672 --name rabbitmq rabbitmq

# 因为在docker run 方式启动的rabbitmq 没有开启插件
# 交互式进入容器
docker exec -it containId bash
# 然后执行如下命令
rabbitmq-plugins enable rabbitmq_management
# 安装好插件之后就可以访问了 
```
`localhost:15672是UI界面客户端，能看到一些基本信息`

`程序中需要用 amqp://guest:guest@localhost:5672/ 来建立连接`


#### rabbitmq的queue,exchange,binding的关系
1. 从AMQP协议可以看出，queue,exchange,binding是构成amqp协议的核心
    + Producer 消息生产者，投递消息的程序
    + Broker 消息队列服务器实体
      + Exchange消息交换机，它指定按什么规则，路由到哪个队列
      + Binding绑定，它的作用就是把Exchange和Queue按照路由规则绑定
      + Queue消息队列载体，每个消息都会被投入到一个或多个队列
    + Consumer 消费者，接收消费的程序
2. Exchange有三种主要类型：Fanout,Direct,Topic
   + FanoutExchange会忽略RoutingKey的设置，直接将消息广播到所有绑定的队列中
   + DirectExchange是RabbitMQ默认的Exchange,完全根据RoutingKey来路由消息。设置Exchange和Queue的Binding时需指定RoutingKey(一般为QueueName),发消息时也指定一样的RoutingKe，消息就会被路由到对应的Queue中
   + TopicExchange和DirectExchange类似，需要通过RoutingKey来路由消息，区别在于Topic支持模糊匹配。*表示匹配一个单词，#表示匹配没有或多个单词
   + HeadersExchange会忽略RoutingKey而根据消息中的Headers和创建绑定关系时指定的Arguments来匹配决定路由到哪些Queue,性能较差，可以被DirectExchange取代，不建议使用
   + DefaultExchange是一种特殊的DirectExchange,当你手动创建一个队列时，后台自动会将这个队列绑定到一个名称为空的DirectExchange上，绑定RoutingKey与队列名称相同。有了这个默认的交换机和绑定,使我们只关系队列这一层即可，适合一些简单的应用
   





#### 延伸：延时队列的解决方案
1. Redis Sorted Set
2. rabbitmq 利用`TTL(Time To Live)`和`DLX(Dead Letter Exchanges)`
3. 时间轮算法(比如kafka)




