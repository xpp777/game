package netw

import (
	"bytes"
	"encoding/binary"
	"game-ws/iface"
)

//DataPack 封包拆包类实例，暂时不需要成员
type DataPack struct{}

//NewDataPack 封包拆包实例初始化方法
func NewDataPack() iface.Packet {
	return &DataPack{}
}

//Pack 封包方法(压缩数据)
func (dp *DataPack) Pack(msg iface.Message) ([]byte, error) {
	//创建一个存放bytes字节的缓冲
	dataBuff := bytes.NewBuffer([]byte{})
	//写msgID
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgID()); err != nil {
		return nil, err
	}
	//写data数据
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}
	return dataBuff.Bytes(), nil
}

//Unpack 拆包方法(解压数据)
func (dp *DataPack) Unpack(binaryData []byte) (iface.Message, error) {
	//创建一个从输入二进制数据的ioReader
	dataBuff := bytes.NewReader(binaryData)
	//只解压head的信息，得到dataLen和msgID
	msg := &Message{}
	//读msgID
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.ID); err != nil {
		return nil, err
	}
	//读dataLen
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Data); err != nil {
		return nil, err
	}
	//这里只需要把head的数据拆包出来就可以了，然后再通过head的长度，再从conn读取一次数据
	return msg, nil
}
