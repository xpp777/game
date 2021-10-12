package netw

import "game-ws/iface"

//Request 请求
type Request struct {
	conn iface.Connection //已经和客户端建立好的 链接
	msg  iface.Message    //客户端请求的数据
}

//GetConnection 获取请求连接信息
func (r *Request) GetConnection() iface.Connection {
	return r.conn
}

//GetData 获取请求消息的数据
func (r *Request) GetData() interface{} {
	return r.msg.GetData()
}

//GetMsgID 获取请求的消息的ID
func (r *Request) GetMsgID() uint32 {
	return r.msg.GetMsgID()
}
