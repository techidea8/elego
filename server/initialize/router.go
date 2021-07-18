package initialize

import (
	"turingapp/ctrl"
	"turingapp/global"

	"github.com/techidea8/restgo"
)

//注册控制器
func InitRouter() {
	restgo.EnableCors()
	new(ctrl.AuthCtrl).Router()
	new(ctrl.OpenCtrl).Router()
	new(ctrl.CaptchaCtrl).Router()
	new(ctrl.AttachCtrl).Router()
	new(ctrl.UserCtrl).Router()
	new(ctrl.AreaCtrl).Router()
	new(ctrl.WeixinCtrl).Router()
	new(ctrl.WxmpApiCtrl).Router()
	new(ctrl.DictCtrl).Router()
	//ctrl.RegisterFileServer()
	global.Log.Debug("[√]完成控制器解析 ")
}
