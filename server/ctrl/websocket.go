package ctrl

import (
	"net/http"
	"sync"
	"turingapp/global"
	"turingapp/service"

	"github.com/gorilla/websocket"
	"github.com/techidea8/restgo"
	"go.uber.org/zap"
)

type WebsocketCtrl struct {
	restgo.Ctrl
}

func (ctrl *WebsocketCtrl) Router() {
	mod := restgo.Module("websocket")
	mod.Router("entry", ctrl.WebsocketEntry)
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // use default options
var deviceMap sync.Map

//前端通过/mqtt/websocket?clientid=deviceno链接过来
func (ctrl *WebsocketCtrl) WebsocketEntry(w http.ResponseWriter, req *http.Request) {
	c, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		global.Log.Error(err.Error())
		return
	}
	deviceNo := req.URL.Query().Get("clientid")
	service.StoreConnect(deviceNo, c)
	//service.Online(deviceNo)
	global.Log.Debug("online", zap.String("no", deviceNo))

	defer func(key string) {
		service.RemoveConnect(key)
		//service.Offline(key)
		global.Log.Debug("offline", zap.String("no", key))
		c.Close()
	}(deviceNo)

	for {
		//读取前端的数据
		_, msg, err := c.ReadMessage()
		if err != nil {
			global.Log.Debug("error when readmsg", zap.Error(err))
			return
		}
		global.Log.Debug("<=", zap.Binary("msg", msg))

	}
	//我们现在要发布信息到前端
}
