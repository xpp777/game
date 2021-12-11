# websocket 轻量级游戏框架

## 使用方式:

```
	netw.SetConfig(&iface.Config{
		PingTime:       30,
		MaxConn:        100000,
		WorkerPoolSize: 10,
		MessageType:    2,
	})
	s := netw.NewServer(netw.WithPacket(pack.NewDataPack()))
	s.AddRouter(0, &router.PingRouter{})
	gin.SetMode(gin.ReleaseMode)
	g := gin.New()
	g.Use(gin.Recovery(), middleware.Logs)
	g.GET("/ws", s.Start)
	g.Run(":8080")
```
