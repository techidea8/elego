package args

import (
	"crypto/md5"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"reflect"
	"sort"
	"strings"
	"time"
	"turingapp/global"
	"turingapp/utils"
)

type UniordeReq struct {
	Appid       string `json:"appid" xml:"appid"`
	MchId          string `json:"mch_id" xml:"mch_id"`
	NonceStr       string `json:"nonce_str" xml:"nonce_str"`
	Sign           string `json:"sign" xml:"sign"`
	Body      string `json:"body" xml:"body"`
	OutTradeNo string `json:"out_trade_no" xml:"out_trade_no"`
	TotalFee     string `json:"total_fee" xml:"total_fee"`
	SpbillCreateIp string `json:"spbill_create_ip" xml:"spbill_create_ip"`
	NotifyIUrl string `json:"notify_url" xml:"notify_url"`
	TradeType string `json:"trade_type" xml:"trade_type"`
	Openid string `json:"openid" xml:"openid"`

}

//
func (this *UniordeReq) BuildXml(apikey string) (result string, err error) {
	bytedata, err := json.Marshal(this)
	if err != nil {
		return
	}
	var mapResult map[string]string
	err = json.Unmarshal(bytedata, &mapResult)
	if err != nil {
		return
		//fmt.Println("JsonToMapDemo err: ", err)
	}
	sign,err := BuildSignForWxPay(mapResult,apikey)
	//fmt.Println("sign = " + sign)
	mapResult["sign"] = sign
	result, err = utils.Map2Xml(mapResult)
	return result, err
}

func BuildSignForWxPay(mapdata map[string]string,apikey string) (sign string,err error) {
	var keys []string
	for key, v := range mapdata {
		//空值和sign字段不参与签名

		if key != "sign" {
			if v != "" {
				keys = append(keys, key)
			}

		}
	}
	sort.Strings(keys)
	strArr := make([]string, 0)
	for _, k := range keys {
		strArr = append(strArr, k+"="+mapdata[k])
	}
	//fmt.Println("str = " + strings.Join(strArr, "&"))
	strArr = append(strArr, "key="+apikey)
	strA := strings.Join(strArr, "&")
	sign = utils.Md5Encode(strA)
	//fmt.Println("sign = " + sign)
	return sign,err
}

//发送红包
func (this *UniordeReq) Dispatch(apikey string) (result map[string]string, err error) {
	postxml,err := this.BuildXml(apikey)
	if err!=nil{
		return
	}
	xmlresp, err := utils.HTTPPost("https://api.mch.weixin.qq.com/pay/unifiedorder", postxml)

	if err!=nil{
		return
	}
	var resp UniordeResp
	err = xml.Unmarshal(xmlresp, &resp)
	if err != nil {
		return
	}
	if resp.ReturnMsg != "OK" {
		err = errors.New(resp.ReturnMsg)
		return
	}
	if resp.ResultCode != "SUCCESS" {
		err = errors.New(resp.ErrCodeDes)
		return
	}
	return resp.ForJsApi(apikey)
}

func (resp *UniordeResp)ForJsApi(apikey string)(result map[string]string,err error){
	mapresult := make(map[string]string,0)
	mapresult["appId"] = resp.Appid
	mapresult["timeStamp"] = fmt.Sprintf("%d",time.Now().Unix())
	mapresult["nonceStr"] = fmt.Sprintf("%d",time.Now().Unix())
	mapresult["package"] = "prepay_id="+resp.PrepayId
	mapresult["signType"] = "MD5"
	mapresult["paySign"],err = BuildSignForWxPay(mapresult,apikey)
	return mapresult,err
}

//WxPaySendRedPackResponse 微信发送红包返回值
type UniordeResp struct {
	XMLName    xml.Name `xml:"xml"`
	ReturnCode string   `json:"return_code" xml:"return_code"` // SUCCESS/FAIL 此字段是通信标识,非交易标识,交易是否成功需要查看result_code来判断
	ReturnMsg  string   `json:"return_msg" xml:"return_msg"`   // 返回信息,如非空,为错误原因:签名失败/参数格式校验错误

	//以下字段在return_code为SUCCESS的时候有返回

	Appid       string `json:"appid" xml:"appid"`
	MchId          string `json:"mch_id" xml:"mch_id"`
	NonceStr       string `json:"nonce_str" xml:"nonce_str"`
	Sign           string `json:"sign" xml:"sign"`
	ResultCode     string `json:"result_code" xml:"result_code"`   // SUCCESS/FAIL
	ErrCode        string `json:"err_code" xml:"err_code"`         // 详细参见第6节错误列表
	ErrCodeDes     string `json:"err_code_des" xml:"err_code_des"` // 错误返回的信息描述
	TradeType string `json:"trade_type" xml:"trade_type"`
	PrepayId string `json:"prepay_id" xml:"prepay_id"`
	IsSubscribe string  `json:"is_subscribe,omitempty" xml:"is_subscribe,omitempty"`
	BankType string  `json:"bank_type,omitempty" xml:"bank_type,omitempty"`
}

type NotifyReq struct {
	XMLName    xml.Name `xml:"xml"`
	ReturnCode string   `json:"return_code" xml:"return_code"` // SUCCESS/FAIL 此字段是通信标识,非交易标识,交易是否成功需要查看result_code来判断
	ReturnMsg  string   `json:"return_msg" xml:"return_msg"`   // 返回信息,如非空,为错误原因:签名失败/参数格式校验错误

	Appid       string `json:"appid" xml:"appid"`
	MchId          string `json:"mch_id" xml:"mch_id"`
	NonceStr       string `json:"nonce_str" xml:"nonce_str"`
	Sign           string `json:"sign" xml:"sign"`
	ResultCode     string `json:"result_code" xml:"result_code"`   // SUCCESS/FAIL
	ErrCode        string `json:"err_code" xml:"err_code"`         // 详细参见第6节错误列表
	ErrCodeDes     string `json:"err_code_des" xml:"err_code_des"` // 错误返回的信息描述
	TradeType string `json:"trade_type" xml:"trade_type"`
	PrepayId string `json:"prepay_id" xml:"prepay_id"`
	IsSubscribe string  `json:"is_subscribe" xml:"is_subscribe"`
	BankType string  `json:"bank_type" xml:"bank_type"`
	TotalFee int `json:"total_fee" xml:"total_fee"`
	CashFee int `json:"cash_fee" xml:"cash_fee"`
	TransactionId string `json:"transaction_id" xml:"transaction_id"`
	OutTradeNo string `json:"out_trade_no" xml:"out_trade_no"`
	TimeEnd string `json:"time_end" xml:"time_end"`
}
func (this *NotifyReq)ToMap()(map[string]interface{}){
 m := make(map[string]interface{})
 m["appid"] = this.Appid
 m["mch_id"] = this.MchId
 m["nonce_str"] = this.NonceStr
 m["sign"] = this.Sign
 m["result_code"] = this.ResultCode
 m["err_code"] = this.ErrCode
 m["err_code_des"] = this.ErrCodeDes
 m["trade_type"]  = this.TradeType
 m["prepay_id"] = this.PrepayId
 m["is_subscribe"] = this.IsSubscribe
 m["bank_type"] = this.BankType
 m["total_fee"] = this.TotalFee
 m["cash_fee"] = this.CashFee
 m["transaction_id"] =this.TransactionId
 m["out_trade_no"] = this.OutTradeNo
 m["time_end"] = this.TimeEnd
 return m

}


func (this *NotifyReq)ParseNotify(apikey string)(result NotifyReq,err error){
	mapdata := this.ToMap()
	if err!=nil{
		return result,err
	}
	mapdata2 := make(map[string]string)
	for k,v := range mapdata{
		if k=="total_fee" || k=="cash_fee"{
			mapdata2[k] = fmt.Sprintf("%d",v)
		}else{
			mapdata2[k] = fmt.Sprintf("%s",v)
		}
	}
	sign,err := BuildSignForWxPay(mapdata2,apikey)
	result = NotifyReq{}
	if err!=nil{
		return result,err
	}
	if mapdata["sign"] != sign{
		err = errors.New("签名不正确")
		return
	}else{
		mapdata1,err:=json.Marshal(mapdata)
		if err!=nil{
			return result,err
		}
		err =json.Unmarshal(mapdata1,&result)
		if err!=nil{
		   return result,err
		}
	}
	return result,err
}
func (this *NotifyReq)IsOk()bool{
	return this.ReturnCode =="SUCCESS"
}
//发送红包
func (this *UniordeResp) String() string {
	return utils.JsonString(this)
}

//发送红包
func (this *UniordeResp) IsOk() bool {
	return this.ReturnMsg == "OK"
}

//发送红包
func (this *UniordeResp) Xml() string {
	return utils.XmlString(this)
}



//微信支付回调

const (
	AckSuccess = `<xml><return_code><![CDATA[SUCCESS]]></return_code></xml>`
	AckFail    = `<xml><return_code><![CDATA[FAIL]]></return_code></xml>`
)



type WXPayNotify struct {
	ReturnCode    string `xml:"return_code"`
	ReturnMsg     string `xml:"return_msg"`
	Appid         string `xml:"appid"`
	MchID         string `xml:"mch_id"`
	DeviceInfo    string `xml:"device_info"`
	NonceStr      string `xml:"nonce_str"`
	Sign          string `xml:"sign"`
	ResultCode    string `xml:"result_code"`
	ErrCode       string `xml:"err_code"`
	ErrCodeDes    string `xml:"err_code_des"`
	Openid        string `xml:"openid"`
	IsSubscribe   string `xml:"is_subscribe"`
	TradeType     string `xml:"trade_type"`
	BankType      string `xml:"bank_type"`
	TotalFee      int64  `xml:"total_fee"`
	FeeType       string `xml:"fee_type"`
	CashFee       int64  `xml:"cash_fee"`
	CashFeeType   string `xml:"cash_fee_type"`
	CouponFee     int64  `xml:"coupon_fee"`
	CouponCount   int64  `xml:"coupon_count"`
	CouponID0     string `xml:"coupon_id_0"`
	CouponFee0    int64  `xml:"coupon_fee_0"`
	TransactionID string `xml:"transaction_id"`
	OutTradeNo    string `xml:"out_trade_no"`
	Attach        string `xml:"attach"`
	TimeEnd       string `xml:"time_end"`
}

/*
* 支付信息验证
* param data WXPayNotify
* reply true/false
 */
func (this *WXPayNotify)WXPayVerify(apikey string) bool {

	sign := WXmd5Sign(this,apikey)
	if this.Sign == sign {
		return true
	} else {

	}
	return false
}

/*
* md5签名
* param data interface{}
* reply sign 签名字符串
 */
func WXmd5Sign(data interface{},WXPApiKey string) (sign string) {
	val := make(map[string]string)
	datavalue := reflect.ValueOf(data)
	if datavalue.Kind() != reflect.Struct {
		global.Log.Warn("NOT A STRUCT ",zap.Any("kind",data))
		return ""
	}
	var keys []string
	for i := 0; i < datavalue.NumField(); i++ {
		k := datavalue.Type().Field(i)
		kl := k.Tag.Get("xml")
		v := fmt.Sprintf("%v", datavalue.Field(i))

		if v != "" && v != "0" && kl != "sign" {
			val[kl] = v
			keys = append(keys, kl)
		}
	}
	sort.Strings(keys)
	var stra string
	for _, v := range keys {
		stra = stra + v + "=" + val[v] + "&"
	}
	strb := stra + "key=" + WXPApiKey

	global.Log.Info("NOT A STRUCT ",zap.Any("kind",strb))
	hstr := md5.Sum([]byte(strb))

	sum := fmt.Sprintf("%x", hstr)
	sign = strings.ToUpper(sum)
	return sign
}