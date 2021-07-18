package ctrl

import (
	"net/http"

	"github.com/techidea8/restgo"
	"turingapp/service"
)

type AttachCtrl struct {
	restgo.Ctrl
}

//路由注册
func (ctrl *AttachCtrl) Router() {
	grouter := restgo.Module("attach")
	grouter.Router("upload", ctrl.upload)
}

//上传
//退出系统
func (ctrl *AttachCtrl) upload(w http.ResponseWriter, req *http.Request) {
	oss := service.NewOss()
	_, header, err := req.FormFile("file")
	if err != nil {
		restgo.RespFail(w, err.Error())
		return
	}
	url, key, err := oss.UploadFile(header)
	if err != nil {
		restgo.RespFail(w, err.Error())
		return
	}
	restgo.RespOkMap(w, map[string]interface{}{"url": url, "key": key})
}
