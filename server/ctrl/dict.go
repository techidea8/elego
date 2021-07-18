
package ctrl

import (
	"github.com/techidea8/restgo"
	"github.com/techidea8/restgo/middleware"
	"net/http"
	"turingapp/args"
	"turingapp/model"
	"turingapp/service"
)

type DictCtrl struct {
	restgo.Ctrl
}
//注册路由,需要再initinal/router下添加路由注册
func (ctrl *DictCtrl) Router() {
	mod := restgo.Module("dict").Using(middleware.JwtAuth(),"search","getone")
	mod.Router("search", ctrl.Search)
	mod.Router("create", ctrl.Create)
	mod.Router("update", ctrl.Update)
	mod.Router("delete", ctrl.Delete)
	mod.Router("getone", ctrl.GetOne)

}

//搜索
func (ctrl *DictCtrl) Search(w http.ResponseWriter, req *http.Request) {
	arg := args.Dict{}
	restgo.Bind(req, &arg)
	if result, total, err := service.SearchDict(arg); err != nil {
		ctrl.RespFail(w, err.Error())
	} else {

		ctrl.RespList(w, result, total)
	}
}

//创建
func (ctrl *DictCtrl) Create(w http.ResponseWriter, req *http.Request) {
	property := model.Dict{}
	restgo.Bind(req, &property)
	if result, err := service.CreateDict(property); err != nil {
		ctrl.RespFail(w, err.Error())
	} else {
		ctrl.RespOk(w, result)
	}
}

//更新
func (ctrl *DictCtrl) Update(w http.ResponseWriter, req *http.Request) {
	property := model.Dict{}
	restgo.Bind(req, &property)
	result, err := service.UpdateDict(property)
	if err != nil {
		ctrl.RespFail(w, err.Error())
	} else {
		ctrl.RespOk(w, result)
	}
}

//删除一条记录
func (ctrl *DictCtrl) Delete(w http.ResponseWriter, req *http.Request) {
	property := model.Dict{}
	restgo.Bind(req, &property)
	if result, err := service.DeleteDict(property); err != nil {
		ctrl.RespFail(w, err.Error())
	} else {
		ctrl.RespOk(w, result)
	}
}

//获取一条记录
func (ctrl *DictCtrl) GetOne(w http.ResponseWriter, req *http.Request) {
	var property model.Dict
	restgo.Bind(req, &property)
	result := model.Dict{}
	var err error
	if(property.Id>0){
		result, err= service.FindDictById(property.Id);
	}else {
		result, err= service.FindDictByName(property.Name);
	}
	if err!=nil{
		ctrl.RespFail(w,err.Error())
	}else{
		ctrl.RespOk(w,result)
	}
}


