package iface

import (
	"context"
	"net"

	"github.com/gorilla/websocket"
)

/**
定义连接接口
*/
type Connection interface {
	Start()                                       // 启动连接，让当前连接开始工作
	Stop()                                        // 停止连接，结束当前连接状态M
	Context() context.Context                     // 返回ctx，用于用户自定义的go程获取连接退出状态
	GetConnection() *websocket.Conn               // 从当前连接获取原始的socket Conn
	GetConnID() int64                             // 获取当前连接ID
	RemoteAddr() net.Addr                         // 获取远程客户端地址信息
	SendMsg(msgID uint32, data interface{}) error // 直接将Message数据发送数据给远程的客户端
	SetPingTime()                                 // 设置心跳
	IsHeartbeatTimeout() (timeout bool)           // 检测心跳

	SetProperty(key string, value interface{}) 		//设置链接属性
	GetProperty(key string) (interface{}, error)	//获取链接属性
	RemoveProperty(key string) 						//移除链接属性
}
