#### PProf是什么？
分析go程序执行的性能工具
#### 有哪些采样方式？
+ rumtime/pprof:采集程序的指定区块的运行数据进行分析
+ net/http/pprof:基于HTTP Server运行，并且可以采集运行时数据进行分析
+ go test:通过测试用例，并指定所需标识来进行采集
#### 支持什么使用模式
+ 报告生成
+ 交互式终端使用
+ web界面
#### 都能采集哪些性能项？

+ CPU Profiling
+ Memory Profiling
+ Block Profiling
+ Mutex Profiling
+ Goroutine Profiling
  + Goroutine 分析，可以对当前应用程序正在运行的 Goroutine 进行堆栈跟踪和分析。这项功能在实际排查中会经常用到，因为很多问题出现时的表象就是 Goroutine 暴增，而这时候我们要做的事情之一就是查看应用程序中的 Goroutine 正在做什么事情，因为什么阻塞了
```go

```
