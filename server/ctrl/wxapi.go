package ctrl

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"github.com/silenceper/wechat/v2/util"
	"github.com/techidea8/restgo"
	"github.com/techidea8/restgo/middleware"
	"io"
	"net/http"
	url2 "net/url"
	"time"
	"turingapp/args"
	"turingapp/model"
	"turingapp/service"
	"turingapp/utils"
)

type WxmpApiCtrl struct {
	restgo.Ctrl
}

func (this *WxmpApiCtrl) Router() {
	mod := restgo.Module("wxmpapi")
	mod.Router("authorizeurl", this.Authorizeurl)
	mod.Router("userinfo", this.Userinfo)
	mod.Router("initjssdk", this.Initjssdk)
	mod.Router("initcard", this.Initcard)
}



//jsapi_ticket=sM4AOVdWfPE4DxkXGEs8VMCPGGVi4C3VM0P37wVUCFvkVAy_90u5h9nbSlYy3-Sl-HhTdfl2fzFy1AOcHKP7qg&
//noncestr=Wm3WZYTPz0wzccnW&
//timestamp=1414587457&
//url=http://mp.weixin.qq.com?params=value
func (ctrl *WxmpApiCtrl) Initjssdk(w http.ResponseWriter, req *http.Request) {
	var arg args.JssdkReq
	restgo.Bind(req, &arg)
	arg.NonceStr = utils.RandomString(10)
	arg.Timestamp = fmt.Sprintf("%d",time.Now().Unix())
	jsticket := service.GetJsapiTicket(arg.AppId)
    signdata := fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%s&url=%s",jsticket,arg.NonceStr,arg.Timestamp,arg.Url)
    h := sha1.New()
    io.WriteString(h,signdata)
    arg.Signature = fmt.Sprintf("%x",h.Sum(nil))
	ctrl.RespOk(w,arg)
}
type CardReq struct {
	AppId string `json:"appId" form:"appId"`
	CardId string `json:"cardId" form:"cardId"`
}
func (ctrl *WxmpApiCtrl) Initcard(w http.ResponseWriter, req *http.Request) {
	var arg CardReq
	restgo.Bind(req, &arg)
	jsticket := service.GetWxcardTicket(arg.AppId)
	wxcardReq := args.NewWxCardReq()
	wxcardReq.CardId = arg.CardId
	fmt.Printf("%v   %s",wxcardReq,jsticket)
	result := wxcardReq.BuildPackage(jsticket)
	ctrl.RespOk(w,result)
}
//鉴权
func (ctrl *WxmpApiCtrl) Authorizeurl(w http.ResponseWriter, req *http.Request) {
	var arg args.WxapiReq
	restgo.Bind(req, &arg)
	if arg.Scope==""{
		arg.Scope = args.SNAAPIUSERINFO
	}
	url := fmt.Sprintf("https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=%s&state=1#wechat_redirect",arg.AppId,url2.QueryEscape(arg.Url),arg.Scope)
	ctrl.RespOk(w,url)
}


//微信、小程序的snscode
func (ctrl *WxmpApiCtrl) Userinfo(w http.ResponseWriter, req *http.Request) {
	var arg args.WxapiReq
	restgo.Bind(req, &arg)
	weixin,err := service.FindWeixinByAppId(arg.AppId)
	if err!=nil{
		ctrl.RespFail(w,err.Error())
		return
	}
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code",arg.AppId,weixin.Secret,arg.Code)
	/**
	{
	  "access_token":"ACCESS_TOKEN",
	  "expires_in":7200,
	  "refresh_token":"REFRESH_TOKEN",
	  "openid":"OPENID",
	  "scope":"SCOPE"
	}
	 */
	result,err := util.HTTPGet(url)
	if err!=nil{
		ctrl.RespFail(w,err.Error())
		return
	}
	var resp args.WxapiOauth2AccessTokenResp
	json.Unmarshal(result,&resp)
	if !resp.IsOk(){
		ctrl.RespFail(w,err.Error())
		return
	}
	//如果没有进一步的意愿,那么就直接返回吧
	if !resp.IsSnsapiUserinfo(){
		ctrl.RespOk(w,resp.OpenId)
		return
	}
	url = fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN",resp.AccessToken,resp.OpenId)
	result,err = util.HTTPGet(url)
	if err!=nil{
		ctrl.RespFail(w,err.Error())
		return
	}
	var userInfo args.WxapiSnsUserinfo
	json.Unmarshal(result,&userInfo)
	if !resp.IsOk(){
		ctrl.RespFail(w,err.Error())
		return
	}
	ctrl.RespOk(w,userInfo)
}


func (ctrl *WxmpApiCtrl) respUserInfo(w http.ResponseWriter, user model.User) {

	token, err := middleware.CreateToken(user.Id)
	if err != nil {
		ctrl.RespFail(w, err.Error())
		return
	}
	userMap := make(map[string]interface{}, 0)
	userMap["token"] = token
	//id, role, name, avatar ,token
	userMap["id"] = user.Id
	userMap["name"] = user.NickName
	userMap["avatar"] = user.HeadImage
	userMap["role"] = user.Role
	ctrl.RespOk(w, userMap)

}
