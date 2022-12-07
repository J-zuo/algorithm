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

#### 延时队列的解决方案
1. Redis Sorted Set
2. rabbitmq 利用`TTL(Time To Live)`和`DLX(Dead Letter Exchanges)`
3. 时间轮算法(比如kafka)




