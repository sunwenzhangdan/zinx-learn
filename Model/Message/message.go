package messagemodel



type Message struct{
	Id uint32
	DataLen uint32
	Data []byte
}


/**
 * @description: New Message
 * @param {*} id
 * @param {uint32} dataLen
 * @param {[]byte} data
 * @return {*}
 */
func NewMessage(id,dataLen uint32,data []byte) *Message{
	return &Message{
		Id: id,
		DataLen: dataLen,
		Data: data,
	}
}


func (m *Message) GetMessageId() uint32{
    return m.Id
}

func (m *Message) GetDataLen() uint32{
     return m.DataLen	
}
func (m *Message) GetData() []byte{
	return m.Data
}
func (m *Message) SetDataLen(datalen uint32){
	m.DataLen=datalen

}
func (m *Message) SetMessageId(id uint32){
	m.Id=id
	
}
func (m *Message) SetData(data []byte){
      m.Data=data
}


