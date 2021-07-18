package ctrl

import (
	"net/http"
	"turingapp/args"
	"turingapp/model"
	"turingapp/service"

	"github.com/techidea8/restgo"
	"github.com/techidea8/restgo/middleware"
)

type UserCtrl struct {
	restgo.Ctrl
}

func (ctrl *UserCtrl) Router() {
	mod := restgo.Module("user").Using(middleware.JwtAuth())
	mod.Router("search", ctrl.Search)
	mod.Router("create", ctrl.Create)
	mod.Router("snsinfo", ctrl.SnsInfo)
	mod.Router("update", ctrl.Update)
	mod.Router("updatemobile", ctrl.Updatemobile)
	mod.Router("delete", ctrl.Delete)
	mod.Router("enable", ctrl.Enable)
	mod.Router("disable", ctrl.Disable)

}



//从图片列表中创建course
func (ctrl *UserCtrl) Search(w http.ResponseWriter, req *http.Request) {
	arg := args.User{}
	restgo.Bind(req, &arg)
	if result, total, err := service.SearchUser(arg); err != nil {
		ctrl.RespFail(w, err.Error())
	} else {
		ctrl.RespList(w, result, total)
	}
}

//从图片列表中创建course
func (ctrl *UserCtrl) Create(w http.ResponseWriter, req *http.Request) {
	property := args.User{}
	restgo.Bind(req, &property)
	if result, err := service.CreateUserFromMobile(property,property.Role); err != nil {
		ctrl.RespFail(w, err.Error())
	} else {
		ctrl.RespOk(w, result)
	}
}

//从图片列表中创建course
func (ctrl *UserCtrl) Update(w http.ResponseWriter, req *http.Request) {
	property := model.User{}
	restgo.Bind(req, &property)
	if result, err := service.UpdateUser(property); err != nil {
		ctrl.RespFail(w, err.Error())
	} else {
		ctrl.RespOk(w, result)
	}
}


//从图片列表中创建course
func (ctrl *UserCtrl) Updatemobile(w http.ResponseWriter, req *http.Request) {
	property := model.User{}
	restgo.Bind(req, &property)
	if result, err := service.UpdateMobile(property); err != nil {
		ctrl.RespFail(w, err.Error())
	} else {
		ctrl.RespOk(w, result)
	}
}


//从图片列表中创建course
func (ctrl *UserCtrl) Enable(w http.ResponseWriter, req *http.Request) {
	property := model.User{}
	restgo.Bind(req, &property)
	if result, err := service.EnableUser(property); err != nil {
		ctrl.RespFail(w, err.Error())
	} else {
		ctrl.RespOk(w, result)
	}
}//从图片列表中创建course
func (ctrl *UserCtrl) Disable(w http.ResponseWriter, req *http.Request) {
	property := model.User{}
	restgo.Bind(req, &property)
	if result, err := service.DisableUser(property); err != nil {
		ctrl.RespFail(w, err.Error())
	} else {
		ctrl.RespOk(w, result)
	}
}
//从图片列表中创建course
func (ctrl *UserCtrl) Delete(w http.ResponseWriter, req *http.Request) {
	property := model.User{}
	restgo.Bind(req, &property)
	if result, err := service.DeleteUser(property); err != nil {
		ctrl.RespFail(w, err.Error())
	} else {
		ctrl.RespOk(w, result)
	}
}

//从图片列表中创建course
func (ctrl *UserCtrl) SnsInfo(w http.ResponseWriter, req *http.Request) {
	property := model.User{}
	restgo.Bind(req, &property)
	result := model.User{}
	resp := make(map[string]interface{}, 0)

	if property.Id > 0 {
		resul, err := service.FindUserById(property.Id)
		result = resul
		if err != nil {
			resp["errmsg"] = err.Error()
		}
	}

	resp["nickName"] = result.NickName
	resp["headImage"] = result.HeadImage
	ctrl.RespOk(w, resp)
}
