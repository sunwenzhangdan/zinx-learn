/*
 * @Date: 2021-04-23 08:03:58
 * @LastEditors: seven sun
 * @LastEditTime: 2021-04-26 11:04:40
 * @FilePath: /zinx-learn/Model/Connection/Connect.go
 */
package connection

import "net"

type Connection struct {
	Conn     *net.TCPConn
	ConnId   uint32
	isClosed bool
	handleAPI
}

func Start() {

}
func Stop() {

}
func GetConnection() *net.TCPConn {

}
func GetConnId() uint32 {

}
func RemoteAddr() *net.Addr {

}
func SendMsg(msgid int32, data []byte) error {

}
