# websocket 轻量级游戏框架

## 使用方式:

```
	global.Config = config.C
	s := netw.NewServer(netw.WithPacket(pack.NewDataPack()))
	s.AddRouter(0, &router.PingRouter{})
	gin.SetMode(gin.ReleaseMode)
	g := gin.New()
	g.Use(gin.Recovery(), middleware.Logs)
	g.GET("/ws", s.Start)
	g.Run(":8080")
```
