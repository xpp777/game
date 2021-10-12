package iface

/*
	封包数据和拆包数据
*/
type Packet interface {
	Pack(msg Message) ([]byte, error) // 封包方法
	Unpack([]byte) (Message, error)   // 拆包方法
}
