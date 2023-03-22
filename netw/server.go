package netw

import (
	"github.com/xpp777/game/iface"
	"github.com/xpp777/ztimer"
	"net/http"
	"sync/atomic"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	Upgrader = websocket.Upgrader{
		ReadBufferSize:    4096,
		WriteBufferSize:   4096,
		EnableCompression: true,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	GlobalServer iface.Server
	ZTimer       = ztimer.NewAutoExecTimerScheduler()
)

// Server 接口实现，定义一个Server服务类
type Server struct {
	sesIDGen int64 // 记录已经生成的会话ID流水号
	// 当前Server的消息管理模块，用来绑定MsgID和对应的处理方法
	msgHandler iface.MsgHandle
	// 当前Server的链接管理器
	ConnMgr iface.ConnManager
	// 该Server的连接创建时Hook函数
	OnConnStart func(conn iface.Connection)
	// 该Server的连接断开时的Hook函数
	OnConnStop func(conn iface.Connection)
	packet     iface.Packet
}

// NewServer 创建一个服务器句柄
func NewServer(opt ...Option) iface.Server {
	s := &Server{
		msgHandler: NewMsgHandle(),
		ConnMgr:    NewConnManager(),
		packet:     NewDataPack(),
	}
	for _, option := range opt {
		option(s)
	}
	s.msgHandler.StartWorkerPool()
	GlobalServer = s
	return s
}

// ============== 实现 iface.Server 里的全部接口方法 ========

// Start 开启网络服务
func (s *Server) Start(c *gin.Context) {
	// 等待客户端建立连接请求
	var (
		err      error
		wsSocket *websocket.Conn
	)
	if wsSocket, err = Upgrader.Upgrade(c.Writer, c.Request, nil); err != nil {
		return
	}
	if s.ConnMgr.Len() >= config.MaxConn {
		wsSocket.Close()
		return
	}
	// 处理该新连接请求的 业务 方法， 此时应该有 handler 和 conn是绑定的
	dealConn := NewConnection(s, wsSocket, atomic.AddInt64(&s.sesIDGen, 1), s.msgHandler)
	// 启动当前链接的处理业务
	dealConn.Start()
}

// Stop 停止服务
func (s *Server) Stop() {
	zap.S().Info("[STOP] server...")
	// 将其他需要清理的连接信息或者其他信息 也要一并停止或者清理
	s.ConnMgr.ClearConn()
}

// Serve 运行服务
func (s *Server) Serve(c *gin.Context) {
	s.Start(c)
	select {}
}

// AddRouter 路由功能：给当前服务注册一个路由业务方法，供客户端链接处理使用
func (s *Server) AddRouter(msgID uint32, router iface.Router) {
	s.msgHandler.AddRouter(msgID, router)
}

// GetConnMgr 得到链接管理
func (s *Server) GetConnMgr() iface.ConnManager {
	return s.ConnMgr
}

// SetOnConnStart 设置该Server的连接创建时Hook函数
func (s *Server) SetOnConnStart(hookFunc func(iface.Connection)) {
	s.OnConnStart = hookFunc
}

// SetOnConnStop 设置该Server的连接断开时的Hook函数
func (s *Server) SetOnConnStop(hookFunc func(iface.Connection)) {
	s.OnConnStop = hookFunc
}

// CallOnConnStart 调用连接OnConnStart Hook函数
func (s *Server) CallOnConnStart(conn iface.Connection) {
	if s.OnConnStart != nil {
		s.OnConnStart(conn)
	}
}

// CallOnConnStop 调用连接OnConnStop Hook函数
func (s *Server) CallOnConnStop(conn iface.Connection) {
	if s.OnConnStop != nil {
		s.OnConnStop(conn)
	}
}

func (s *Server) Packet() iface.Packet {
	return s.packet
}
