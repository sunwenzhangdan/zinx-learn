/*
 * @Date: 2021-04-22 13:08:25
 * @LastEditors: seven sun
 * @LastEditTime: 2021-04-22 16:13:38
 * @FilePath: /zinx-learn/Inface/Message/Imessage.go
 */
package infacemessage

type Imessage interface {
	GetDataLen() uint32
	GetMessageId() uint32
	GetData() []byte
	SetDataLen(uint32)
	SetMessageId(uint32)
	SetData([]byte)
}
