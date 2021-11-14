package iface

/*
	将请求的一个消息封装到message中，定义抽象层接口
*/
type Message interface {
	GetMsgID() uint32     // 获取消息ID
	GetData() []byte // 获取消息内容

	SetMsgID(uint32)     // 设置消息ID
	SetData([]byte) // 设置消息内容
}
