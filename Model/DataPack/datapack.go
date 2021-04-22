/*
 * @Date: 2021-04-22 16:45:35
 * @LastEditors: seven sun 
 * @LastEditTime: 2021-04-22 18:41:16
 * @FilePath: /zinx-learn/Model/DataPack/datapack.go
 */
package DataPack

import (
	messagemodel "zinx-learn/Model/Message"
	"bytes"
	"encoding/binary"
	infacemessage "zinx-learn/Inface/Message"
)

type Datapack struct {
}

func NewDataPack() *Datapack {
	return &Datapack{}
}
func (dp *Datapack) GetHeadlen() uint32 {
	return 8
}

/**
 * @description: 所有的数据，只通过get set 方法对外提供 不能让其直接拿属性
 * @param {infacemessage.Imessage} msg
 * @return {*}
 */
func (dp *Datapack) Pack(msg infacemessage.Imessage) ([]byte  error){
   dataBuff:=bytes.NewBuffer([]byte{})
   if err:=binary.Write(dataBuff,binary.LittleEndian,msg.GetDataLen());err!=nil{
	   return nil,err
   }
   if err:=binary.Write(dataBuff,binary.LittleEndian,msg.GetMessageId());err!=nil{
	   return nil,err
   }
   if err:=binary.Write(dataBuff,binary.LittleEndian,msg.GetData());err!=nil{
	   return nil,err
   }
   return dataBuff.Bytes(),nil
}

func (dp *Datapack) Unpack(data []byte) (infacemessage.Imessage,error){
	 databuf:=bytes.NewReader(data)
	 msg:=messagemodel.Message{}
	 if err:=binary.Read(databuf,binary.LittleEndian,msg.DataLen);err!=nil{
		 return nil,err
	 }

	 if err!=binary.Read(databuf,binary.LittleEndian,msg.Id);err!=nil{
		 return nil,err
	 }
	 return &msg,nil
}





