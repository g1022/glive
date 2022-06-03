# alive

check alive

# 使用方法

一些定时任务，或者长驻任务，可以主协程启用一个子协程，然后主协程阻塞。
外界可以通过此 http 来探活监测

# 测试示例

启动代码

```
(p3) ➜  alive git:(main) ✗ cd test
(p3) ➜  test git:(main) ✗ go test -test.v -test.run="^TestServer$"
=== RUN   TestServer
2022/06/03 10:22:30 http://127.0.0.1:8090/api/ping
```

看到结果：

```
(p3) ➜  test git:(main) ✗ curl http://127.0.0.1:8090/api/ping
{code:0,msg:"alive"}
```

# 代码嵌入示例

```
go get github.com/g1022/glive
```

在你的阻塞代码之前运行：

```
go func() { 
    alive.RunServer(":8080")
}()
```

注意：

1. 里面的端口是自己的可用端口
2. 仅适用于定时任务场景
3. 如果是 web 服务，自己可以弄

