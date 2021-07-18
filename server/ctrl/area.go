package ctrl

import (
	"net/http"
	"turingapp/args"
	"turingapp/model"
	"turingapp/service"

	"github.com/techidea8/restgo"
	"github.com/techidea8/restgo/middleware"
)

type AreaCtrl struct {
	restgo.Ctrl
}

func (ctrl *AreaCtrl) Router() {
	mod := restgo.Module("area").Using(middleware.JwtAuth(),"search")
	mod.Router("search", ctrl.Search)
	mod.Router("create", ctrl.Create)
	mod.Router("update", ctrl.Update)
	mod.Router("delete", ctrl.Delete)
}

//从图片列表中创建course
func (ctrl *AreaCtrl) Search(w http.ResponseWriter, req *http.Request) {
	arg := args.Area{}
	restgo.Bind(req, &arg)
	if result, total, err := service.SearchArea(arg); err != nil {
		ctrl.RespFail(w, err.Error())
	} else {
		ctrl.RespList(w, result, total)
	}
}

//从图片列表中创建course
func (ctrl *AreaCtrl) Create(w http.ResponseWriter, req *http.Request) {
	property := model.Area{}
	restgo.Bind(req, &property)
	if result, err := service.CreateArea(property); err != nil {
		ctrl.RespFail(w, err.Error())
	} else {
		ctrl.RespOk(w, result)
	}
}

//从图片列表中创建course
func (ctrl *AreaCtrl) Update(w http.ResponseWriter, req *http.Request) {
	property := model.Area{}
	restgo.Bind(req, &property)
	if result, err := service.UpdateArea(property); err != nil {
		ctrl.RespFail(w, err.Error())
	} else {
		ctrl.RespOk(w, result)
	}
}

//从图片列表中创建course
func (ctrl *AreaCtrl) Delete(w http.ResponseWriter, req *http.Request) {
	property := model.Area{}
	restgo.Bind(req, &property)
	if result, err := service.DeleteArea(property); err != nil {
		ctrl.RespFail(w, err.Error())
	} else {
		ctrl.RespOk(w, result)
	}
}