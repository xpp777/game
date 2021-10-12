package iface

import (
	"github.com/gin-gonic/gin"
)

/**
定义服务接口
*/
type Server interface {
	Start(c *gin.Context)                  // 启动服务器方法
	Stop()                                 // 停止服务器方法
	Serve(c *gin.Context)                  // 开启业务服务方法
	AddRouter(msgID uint32, router Router) // 路由功能：给当前服务注册一个路由业务方法，供客户端链接处理使用

	GetConnMgr() ConnManager // 得到链接管理

	SetOnConnStart(func(Connection)) // 设置该Server的连接创建时Hook函数
	SetOnConnStop(func(Connection))  // 设置该Server的连接断开时的Hook函数

	CallOnConnStart(conn Connection) // 调用连接OnConnStart Hook函数
	CallOnConnStop(conn Connection)  // 调用连接OnConnStop Hook函数

	Packet() Packet
}
