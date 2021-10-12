package netw

//Message 消息
type Message struct {
	ID   uint32      `json:"msgId"`          //消息的ID
	Data interface{} `json:"data,omitempty"` //消息的内容
}

//NewMsgPackage 创建一个Message消息包
func NewMsgPackage(ID uint32, data interface{}) *Message {
	return &Message{
		ID:   ID,
		Data: data,
	}
}

//GetMsgID 获取消息ID
func (msg *Message) GetMsgID() uint32 {
	return msg.ID
}

//GetData 获取消息内容
func (msg *Message) GetData() interface{} {
	return msg.Data
}

//SetMsgID 设计消息ID
func (msg *Message) SetMsgID(msgID uint32) {
	msg.ID = msgID
}

//SetData 设计消息内容
func (msg *Message) SetData(data interface{}) {
	msg.Data = data
}
