package config

type WebSocketConfig struct {
	Name             string `json:"name,default=websocket"  mapstructure:"name"`                  // 当前服务器名称
	Version          string `json:"version,default=v1.0"    mapstructure:"version"`               // 当前框架版本号
	PingTime         int    `json:"pingTime" mapstructure:"pingTime"`                             // 心跳检测时间
	MaxConn          int    `json:"maxConn,default=10000"   mapstructure:"maxConn"`               // 当前服务器主机允许的最大链接个数
	WorkerPoolSize   uint32 `json:"workerPoolSize,default=10" mapstructure:"workerPoolSize"`      // 业务工作Worker池的数量
	MaxWorkerTaskLen uint32 `json:"maxWorkerTaskLen,default=200" mapstructure:"maxWorkerTaskLen"` // 业务工作Worker对应负责的任务队列最大任务存储数量
	MaxMsgChanLen    uint32 `json:"maxMsgChanLen,default=1024" mapstructure:"maxMsgChanLen"`      //SendBuffMsg发送消息的缓冲最大长度
	MessageType      int    `json:"messageType,default=1" mapstructure:"messageType"`             // 消息类型
}
