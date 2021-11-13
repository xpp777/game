package iface

type Config struct {
	PingTime         int    // 心跳检测时间
	MaxConn          int    // 当前服务器主机允许的最大链接个数
	WorkerPoolSize   uint32 // 业务工作Worker池的数量
	MessageType      int    // 消息类型
}
