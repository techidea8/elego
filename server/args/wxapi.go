package args

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
	"time"
)

type WxapiReq struct {

	Url      string    `json:"url" form:"url"`
	Code    string   `json:"code" form:"code"`
	AppId   string   `json:"appId" form:"appId"`
	Scope string `json:"scope" form:"scope"
`
}

const (
	SNSAPIBASE = "snsapi_base"
	SNAAPIUSERINFO ="snsapi_userinfo"
)

/**
{
  "openid": "OPENID",
  "nickname": NICKNAME,
  "sex": 1,
  "province":"PROVINCE",
  "city":"CITY",
  "country":"COUNTRY",
  "headimgurl":"https://thirdwx.qlogo.cn/mmopen/g3MonUZtNHkdmzicIlibx6iaFqAc56vxLSUfpb6n5WKSYVY0ChQKkiaJSgQ1dZuTOgvLLrhJbERQQ4eMsv84eavHiaiceqxibJxCfHe/46",
  "privilege":[ "PRIVILEGE1" "PRIVILEGE2"     ],
  "unionid": "o6_bmasdasdsad6_2sgVt7hMZOPfL"
}
 */
//{
//"access_token":"ACCESS_TOKEN",
//"expires_in":7200,
//"refresh_token":"REFRESH_TOKEN",
//"openid":"OPENID",
//"scope":"SCOPE"
//}
type WxapiBaseResp struct {
	ErrCode int `json:"errcode"`
	ErrMsg string `json:"errmsg"`

}
func (this *WxapiBaseResp)IsOk()bool{
	return this.ErrCode == 0
}
func (this *WxapiBaseResp)Error()error{
	return errors.New(this.ErrMsg)
}
type WxapiOauth2AccessTokenResp struct {
	WxapiBaseResp
	AccessToken string `json:"access_token"`
	OpenId string `json:"openid"`
	Scope string `json:"scope"`
}
func (this *WxapiOauth2AccessTokenResp)IsSnsapiUserinfo ()bool{
	return this.Scope == "snsapi_userinfo"
}


type WxapiSnsUserinfo struct {
	WxapiBaseResp
	Openid string `json:"openid"`
	Nickname string `json:"nickname"`
	Sex int `json:"sex"`
	Province string `json:"province"`
	City string `json:"city"`
	Country string `json:"country"`
	HeadImage string `json:"headimgurl"`
	Unionid string `json:"unionid"`
}

func (this *WxapiSnsUserinfo)IsMan ()bool{
	return this.Sex == 1
}

func (this *WxapiSnsUserinfo)IsWoman ()bool{
	return this.Sex == 2
}



type JssdkReq struct{
	AppId string `json:"appId" form:"appId"`
	Timestamp string `json:"timestamp" form:"timestamp"`
	NonceStr string `json:"nonceStr" form:"nonceStr"`
	Signature string `json:"signature" form:"signature"`
	Url string `json:"url" form:"url"`
}


type WxCardReq struct{

	CardId string `json:"cardId" form:"cardId"`
	Timestamp string `json:"timestamp" form:"timestamp"`
	NonceStr string `json:"nonceStr" form:"nonceStr"`
	SignType string `json:"signType" form:"signType"`
	CardSign string `json:"cardSign" form:"cardSign"`
}

func NewWxCardReq()*WxCardReq{
	tmp:=&WxCardReq{}
	tmp.Timestamp = fmt.Sprintf("%d",time.Now().Unix())
	tmp.NonceStr = fmt.Sprintf("%d",time.Now().Unix())
	tmp.SignType = "SHA1"
	return  tmp
}
func (this*WxCardReq)BuildPackage(ticket string)(*WxCardReq){

	h := sha1.New()
	arr := make([]string,0)
	arr = append(arr,this.Timestamp,this.NonceStr,this.CardId,ticket)
	sort.Strings(arr)

	io.WriteString(h,strings.Join(arr,""))
	sum := h.Sum(nil)
	this.CardSign = fmt.Sprintf("%x",sum)

	return this
}