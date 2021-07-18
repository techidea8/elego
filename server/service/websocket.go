package service

import (
	"sync"
	"turingapp/global"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

var deviceMap sync.Map

type CmdType int

const (
	CmdRefreshDevice CmdType = 0x11
	CmdUpgradeDevice CmdType = 0x12
)

type MsgType struct {
	Cmd     CmdType     `json:"cmd"`
	Content interface{} `json:"content"`
}

func PublishMessage(deviceNo string, msg MsgType) {
	go publishMessage(deviceNo, msg)
}
func BradCastMessage(msg MsgType) {
	deviceMap.Range(func(k, v interface{}) bool {

		con := v.(*websocket.Conn)
		con.WriteJSON(msg)
		return true
	})
}

//从图片列表中中创建课程
func publishMessage(deviceNo string, msg MsgType) {
	v, ok := deviceMap.Load(deviceNo)
	if !ok || v == nil {
		global.Log.Error("con 无效", zap.String("deviceNo", deviceNo))
		return
	}
	con := v.(*websocket.Conn)
	con.WriteJSON(msg)
}

//从图片列表中中创建课程
func StoreConnect(deviceNo string, con *websocket.Conn) {
	deviceMap.Store(deviceNo, con)
}

//从图片列表中中创建课程
func RemoveConnect(deviceNo string) {
	deviceMap.Delete(deviceNo)
}
