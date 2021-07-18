package ctrl

import (
	"net/http"
	"turingapp/args"
	"turingapp/model"
	"turingapp/service"

	"github.com/techidea8/restgo"
	"github.com/techidea8/restgo/middleware"
)

type WeixinCtrl struct {
	restgo.Ctrl
}

func (ctrl *WeixinCtrl) Router() {
	mod := restgo.Module("weixin").Using(middleware.JwtAuth())
	mod.Router("search", ctrl.Search)
	mod.Router("create", ctrl.Create)
	mod.Router("update", ctrl.Update)
	mod.Router("delete", ctrl.Delete)
}

//从图片列表中创建course
func (ctrl *WeixinCtrl) Search(w http.ResponseWriter, req *http.Request) {
	arg := args.Weixin{}
	restgo.Bind(req, &arg)
	if result, total, err := service.SearchWeixin(arg); err != nil {
		ctrl.RespFail(w, err.Error())
	} else {
		ctrl.RespList(w, result, total)
	}
}

//从图片列表中创建course
func (ctrl *WeixinCtrl) Create(w http.ResponseWriter, req *http.Request) {
	property := model.Weixin{}
	restgo.Bind(req, &property)
	if result, err := service.CreateWeixin(property); err != nil {
		ctrl.RespFail(w, err.Error())
	} else {
		ctrl.RespOk(w, result)
	}
}

//从图片列表中创建course
func (ctrl *WeixinCtrl) Update(w http.ResponseWriter, req *http.Request) {
	property := model.Weixin{}
	restgo.Bind(req, &property)
	if result, err := service.UpdateWeixin(property); err != nil {
		ctrl.RespFail(w, err.Error())
	} else {
		ctrl.RespOk(w, result)
	}
}

//从图片列表中创建course
func (ctrl *WeixinCtrl) Delete(w http.ResponseWriter, req *http.Request) {
	property := model.Weixin{}
	restgo.Bind(req, &property)
	if result, err := service.DeleteWeixin(property); err != nil {
		ctrl.RespFail(w, err.Error())
	} else {
		ctrl.RespOk(w, result)
	}
}