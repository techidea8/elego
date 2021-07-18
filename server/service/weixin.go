package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/silenceper/wechat/v2/util"
	"go.uber.org/zap"
	"turingapp/args"
	"turingapp/global"
	"turingapp/model"
)


//根据手机号查找用户
func FindWeixinByAppId(appId string) (result model.Weixin, err error) {
	instance := model.Weixin{}
	if err = global.DbEngin.Model(new(model.Weixin)).Where("app_id = ?", appId).First(&instance).Error; err != nil {
		return instance, err
	}
	return instance, nil
}

var weixinMap  map[string]*model.Weixin

func init()  {
	weixinMap = make(map[string]*model.Weixin)
}

func GetUserInfo(appId,code string) {

}
//根据手机号查找用户
func FindWeixinById(id uint) (result model.Weixin, err error) {
	instance := model.Weixin{}
	if err = global.DbEngin.Model(new(model.Weixin)).Where("id = ?", id).First(&instance).Error; err != nil {
		return instance, err
	}
	return instance, nil
}

func Sendredpack(){

}

//创建新的设备
func CreateWeixin(property model.Weixin) (result model.Weixin, err error) {
	if property.AppId == "" {
		err = errors.New("缺少APPID")
		return
	}
	if property.Secret == "" {
		err = errors.New("缺少密钥")
		return
	}
	err = global.DbEngin.Create(&property).Error
	return property, nil
}

//更新设备
func UpdateWeixin(property model.Weixin) (result model.Weixin, err error) {
	err = global.DbEngin.Where("id = ?", property.Id).Updates(&property).Error
	return property, err
}

//删除
func DeleteWeixin(property model.Weixin) (result model.Weixin, err error) {
	err = global.DbEngin.Where("id = ?", property.Id).Delete(&property).Error
	return property, nil
}

//搜索支持分页
func SearchWeixin(arg args.Weixin) (result []model.Weixin, total int64, err error) {
	objs := make([]model.Weixin, 0)

	err = global.DbEngin.Model(new(model.Weixin)).Scopes(arg.Paginate(), arg.Condtions()).Order(arg.Sort()).Find(&objs).Error
	total = int64(arg.Total)
	if arg.Total == -1 {
		global.DbEngin.Model(new(model.Weixin)).Scopes(arg.Condtions()).Count(&total)
	}
	return objs, total, nil
}

//
func GetAccessToken(appId string)(token string){
	accesstokenkey := fmt.Sprintf("%s-%s","accesstoken",appId)
	token,_ = global.Redis.Get(accesstokenkey)
	//那么去微信那边取
	if ""==token{
		wx,err := FindWeixinByAppId(appId)
		if err!=nil{
			global.Log.Error(err.Error())
			return
		}
		dsturl := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s",wx.AppId,wx.Secret)
		resp,err := util.HTTPGet(dsturl)
		if err!=nil{
			global.Log.Error(err.Error())
			return
		}
		respMap := make(map[string]interface{})
		json.Unmarshal(resp,&respMap)
		access_token,ok:= respMap["access_token"]
		if !ok{
			global.Log.Error("accesstoken result ",zap.Any("resp",respMap))
			return
		}
		global.Redis.Set(accesstokenkey,access_token)
		global.Redis.SetKeyExpireHourLater(accesstokenkey,1)
		token = access_token.(string)
	}
	return
}


//
func GetJsapiTicket(appId string)(token string){
	jsapiticketkey := fmt.Sprintf("%s-%s","jsapiticket",appId)
	token,_ = global.Redis.Get(jsapiticketkey)
	//那么去微信那边取
	if ""==token{
		accesstoken := GetAccessToken(appId)
		if accesstoken==""{
			global.Log.Error("accesstoen empty")
			return
		}
		dsturl := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=jsapi",accesstoken)
		resp,err := util.HTTPGet(dsturl)
		if err!=nil{
			global.Log.Error(err.Error())
			return
		}
		respMap := make(map[string]interface{})
		json.Unmarshal(resp,&respMap)
		ticket,ok:= respMap["ticket"]
		if !ok{
			global.Log.Error("jsticket result ",zap.Any("resp",respMap))
			return
		}
		global.Redis.Set(jsapiticketkey,ticket)
		global.Redis.SetKeyExpireHourLater(jsapiticketkey,1)
		token = ticket.(string)

	}
	return
}

//
func GetWxcardTicket(appId string)(token string){
	jsapiticketkey := fmt.Sprintf("%s-%s","wxcardticket",appId)
	token,_ = global.Redis.Get(jsapiticketkey)
	//那么去微信那边取
	if ""==token{
		accesstoken := GetAccessToken(appId)
		if accesstoken==""{
			global.Log.Error("accesstoen empty")
			return
		}
		dsturl := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=wx_card",accesstoken)
		resp,err := util.HTTPGet(dsturl)
		if err!=nil{
			global.Log.Error(err.Error())
			return
		}
		respMap := make(map[string]interface{})
		json.Unmarshal(resp,&respMap)
		ticket,ok:= respMap["ticket"]
		if !ok{
			global.Log.Error("wx_card result ",zap.Any("resp",respMap))
			return
		}
		global.Redis.Set(jsapiticketkey,ticket)
		global.Redis.SetKeyExpireHourLater(jsapiticketkey,1)
		token = ticket.(string)

	}
	return
}
