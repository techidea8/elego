package ctrl

import (
	"encoding/json"
	"fmt"
	"github.com/techidea8/restgo/utils"
	"net/http"
	"turingapp/args"
	"turingapp/global"

	"github.com/techidea8/restgo"
)

type OpenCtrl struct {
	restgo.Ctrl
}

func (this *OpenCtrl) Router() {
	mod := restgo.Module("open")
	mod.Router("appinfo", this.appinfo)
	mod.Router("geocoder",this.Geocoder)
	mod.Router("ipcoder",this.Ipcoder)
}

//发送短信
func (ctrl *OpenCtrl) appinfo(w http.ResponseWriter, req *http.Request) {
	ctrl.RespOk(w, global.Conf.System)
}


//地址解析
func (ctrl *OpenCtrl) Geocoder(w http.ResponseWriter, req *http.Request) {
	var arg args.Geocoder
	restgo.Bind(req,&arg)
	url := fmt.Sprintf("https://apis.map.qq.com/ws/geocoder/v1/?location=%s,%s&key=%s",arg.Lat,arg.Lng,global.Conf.System.QQmapKey)
	bts,err := utils.HttpGet(url)
	if err!=nil{
		restgo.RespFail(w,err.Error())
		return
	}
	var resp args.GeocoderResp
	err = json.Unmarshal(bts,&resp)
	if err!=nil{
		restgo.RespFail(w,err.Error())
		return
	}
	if resp.IsOk(){
		resp.Result.AddressComponent.Lat = arg.Lat
		resp.Result.AddressComponent.Lng = arg.Lng
		resp.Result.AddressComponent.Ip = utils.ClientIp(req)

		restgo.RespOk(w,resp.Result.AddressComponent)
	}else{
		restgo.RespFail(w,resp.Error())
	}


}


//通过Ip定位,目前看起来非常准
func (ctrl *OpenCtrl) Ipcoder(w http.ResponseWriter, req *http.Request) {
	var arg args.Geocoder
	restgo.Bind(req,&arg)
	ip := utils.ClientIp(req)
	url := fmt.Sprintf("https://apis.map.qq.com/ws/location/v1/ip?ip=%s&key=%s",ip,global.Conf.System.QQmapKey)
	bts,err := utils.HttpGet(url)
	if err!=nil{
		restgo.RespFail(w,err.Error())
		return
	}
	var resp args.IpcoderResp
	err = json.Unmarshal(bts,&resp)
	if err!=nil{
		restgo.RespFail(w,err.Error())
		return
	}
	if resp.IsOk(){
		resp.Result.AdInfo.Ip = ip
		restgo.RespOk(w,resp.Result.AdInfo)
	}else{
		restgo.RespFail(w,resp.Error())
	}


}