## Key eviction
[https://redis.io/docs/reference/eviction/](https://redis.io/docs/reference/eviction/)<br>
Overview of Redis key eviction policies

### `maxmemory` configuration directive
redis.conf中修改配置
```bash
# 设置为0时代表不限制，这是64位系统默认行为，而32位系统使用隐式限制3GB
maxmemory 100mb
```

### Eviction policies
简称3大类 lru,lfu,random
细分策略分8种：<br>
1. noeviction 不驱逐，知道内存满了返回报错
2. allkeys-lru 驱逐最近最少使用
3. allkeys-lfu 驱逐最不频繁使用
4. allkeys-random 随机驱逐
5. volatile-lru 带过期时间lru
6. volatile-lfu
7. volatile-random
8. volatile-ttl 移除过期字段设置为 true 且剩余时间最短(TTL)值的键

```
The policies volatile-lru, volatile-lfu, volatile-random, and volatile-ttl behave like noeviction if there are no keys to evict matching the prerequisites.
以上是原文：我的理解是：如果这个过期时间没有到，他就不会触发eviction,就相当于noevication策略
```
选择一个合适的策略是很重要的，我们可以在应用程序时重新配置evication策略，使用Redis 的INFO命令来监视缓存丢失和命中次数以实现调优

### 一些使用策略的经验法则

### 驱逐策略是如何工作的？
步骤：<br>
1. 客户端运行新的命令，数据被添加
2. Redis 检查memory使用情况，如果大于`maxmemory`限制,就根据策略来驱逐key
3. 新命令被执行

### Approximated LRU algorithm 近似的LRU算法
+ 他不是一个精准的LRU算法
+ Redis LRU 算法的重要之处在于，您可以通过改变样本数来检查每次驱逐，从而调整算法的精度。此参数由以下配置指令控制:`maxmemory-samples 5`
+ Redis 不实用精准的LRU算法因为它需要消耗更多的内存
+ 在一个理论 LRU 实现中，我们预计，在旧密钥中，前半个密钥将过期。RedisLRU 算法只是概率性地使较旧的密钥过期。

### LFU evication mode 最少频繁使用的驱逐模式（也是一种近似的）
+ 从4.0版本开始，redis支持了这种模式
+ 它有一个计数器，称为 Morris 计数器，用每个对象只使用几个比特来估计对象访问频率，结合衰减周期，以便计数器随着时间的推移而减少。在某种程度上，我们不再希望把密钥看作是频繁访问的，即使它们过去曾经被访问过，这样算法就可以适应访问模式的变化。