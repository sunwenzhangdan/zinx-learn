/*
 * @Date: 2021-04-23 07:58:49
 * @LastEditors: seven sun
 * @LastEditTime: 2021-04-23 08:02:38
 * @FilePath: /zinx-learn/Inface/connection/Iconnection.go
 */
package iconnection

import "net"

type Iconnection interface {
	Start()
	Stop()
	GetConnection() *net.TCPConn
	GetConnId() uint32
	RemoteAddr() *net.Addr
	SendMsg(msgid int32, data []byte) error
}
