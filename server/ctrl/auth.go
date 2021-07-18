package ctrl

import (
	"encoding/json"
	"fmt"
	"net/http"
	"turingapp/args"
	"turingapp/model"
	"turingapp/service"
	"turingapp/utils"

	"gorm.io/gorm"

	"github.com/techidea8/restgo"
	"github.com/techidea8/restgo/middleware"
)

type AuthCtrl struct {
	restgo.Ctrl
}

func (this *AuthCtrl) Router() {
	mod := restgo.Module("account").Using(middleware.JwtAuth(), "authwithpwd", "authwithcode")
	mod.Router("authwithtoken", this.authwithtoken)
	mod.Router("authwithpwd", this.authwithpwd)
	mod.Router("authwithcode", this.authwithcode)
	mod.Router("getinfo", this.getinfo)
}

type reqArg struct {
	Token     string `json:"token" form:"token"`
	Mobile    string `json:"mobile" form:"mobile"`
	Passwd    string `json:"passwd" form:"passwd"`
	Code      string `json:"code" form:"code"`
	CaptchaId string `json:"captchaId" form:"captchaId"`
}

//发送短信
func (ctrl *AuthCtrl) authwithpwd(w http.ResponseWriter, req *http.Request) {
	var arg reqArg
	restgo.Bind(req, &arg)
	if !service.VerifyCaptcha(arg.CaptchaId, arg.Code) {
		ctrl.RespFail(w, "验证码不正确,请重试")
		return
	}
	user, err := service.LoginWithMobileAndPasswd(arg.Mobile, arg.Passwd)
	if err != nil {
		ctrl.RespFail(w, err.Error())
		return
	}
	ctrl.respUserInfo(w, user)
}

//微信通过code换取token
func (ctrl *AuthCtrl) authwithcode(w http.ResponseWriter, req *http.Request) {
	var arg args.WxapiReq
	restgo.Bind(req, &arg)
	weixin, err := service.FindWeixinByAppId(arg.AppId)
	if err != nil {
		ctrl.RespFail(w, err.Error())
		return
	}
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code", arg.AppId, weixin.Secret, arg.Code)

	result, err := utils.HTTPGet(url)
	if err != nil {
		ctrl.RespFail(w, err.Error())
		return
	}
	var resp args.WxapiOauth2AccessTokenResp
	json.Unmarshal(result, &resp)
	if !resp.IsOk() {
		ctrl.RespFail(w, resp.Error().Error())
		return
	}
	//如果没有进一步的意愿,那么就直接返回吧
	if !resp.IsSnsapiUserinfo() {
		ctrl.RespOk(w, resp.OpenId)
		return
	}
	url = fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN", resp.AccessToken, resp.OpenId)
	result, err = utils.HTTPGet(url)
	if err != nil {
		ctrl.RespFail(w, err.Error())
		return
	}
	var userInfo args.WxapiSnsUserinfo
	json.Unmarshal(result, &userInfo)
	if !resp.IsOk() {
		ctrl.RespFail(w, err.Error())
		return
	}
	//这里处理用户信息
	user, err := service.FindUserByOpenId(userInfo.Openid)
	if err == nil {
		ctrl.respUserInfo(w, user)
		return
	}
	//如果记录不存在
	if err == gorm.ErrRecordNotFound {
		user, err = service.CreateUserFromSnsUserInfo(userInfo, model.RoleUser)
		if err == nil {
			ctrl.RespFail(w, err.Error())
			return
		}
		ctrl.RespOk(w, user)
	} else {
		ctrl.RespFail(w, err.Error())
		return
	}

}

//发送短信
func (ctrl *AuthCtrl) authwithtoken(w http.ResponseWriter, req *http.Request) {
	var arg reqArg
	restgo.Bind(req, &arg)
	userId, err := middleware.ParseToken(arg.Token)
	if err != nil {
		ctrl.RespFail(w, err.Error())
		return
	}
	user, err := service.FindUserById(userId)
	if err != nil {
		ctrl.RespFail(w, err.Error())
	} else {
		ctrl.respUserInfo(w, user)
	}
}

func (ctrl *AuthCtrl) getinfo(w http.ResponseWriter, req *http.Request) {
	var arg reqArg
	restgo.Bind(req, &arg)
	userId, err := middleware.CurrentUserId(req)
	if err != nil {
		ctrl.RespFail(w, err.Error())
		return
	}
	user, err := service.FindUserById(userId)
	if err != nil {
		ctrl.RespFail(w, err.Error())
	} else {
		ctrl.respUserInfo(w, user)
	}
}

func (ctrl *AuthCtrl) respUserInfo(w http.ResponseWriter, user model.User) {

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
	userMap["username"] = user.OpenId
	userMap["expireAt"] = user.ExpireAt.Time().Unix()
	userMap["mobile"] = user.Mobile
	if  len(user.Mobile)>10{
		userMap["mobile"] =  user.Mobile[:4]+"****" + user.Mobile[8:]
	}
	ctrl.RespOk(w, userMap)

}
