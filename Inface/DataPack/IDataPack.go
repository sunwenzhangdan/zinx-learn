/*
 * @Date: 2021-04-22 16:24:37
 * @LastEditors: seven sun 
 * @LastEditTime: 2021-04-22 16:43:24
 * @FilePath: /zinx-learn/Inface/DataPack/IDataPack.go
 */

package infaceDataPack

import (
	infacemessage "zinx-learn/Inface/Message"
)

type IDataPack interface {
	GetHeadlen() uint32
	Pack(msg infacemessage.Imessage) ([]byte error)
	Unpack([]byte) (infacemessage.Imessage error)
}

func todo() {
	// TODO
	// 自动导入包
	//接口实现
}
