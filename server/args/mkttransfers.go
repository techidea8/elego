package args

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"net/http"
	"sort"
	"strings"
	"turingapp/utils"
)

type MkttransfersReq struct {
	MchAppid       string `json:"mch_appid" xml:"mch_appid"`
	Mchid          string `json:"mchid" xml:"mchid"`
	DeviceInfo     string `json:"device_info,omitempty" xml:"device_info,omitempty"`
	NonceStr       string `json:"nonce_str" xml:"nonce_str"`
	Sign           string `json:"sign" xml:"sign"`
	PartnerTradeNo string `json:"partner_trade_no" xml:"partner_trade_no"`
	Openid         string `json:"openid" xml:"openid"`
	CheckName      string `json:"check_name" xml:"check_name"`
	ReUserName     string `json:"re_user_name,omitempty" xml:"re_user_name,omitempty"`
	Amount         string `json:"amount" xml:"amount"`
	Desc           string `json:"desc" xml:"desc"`
	SpbillCreateIp string `json:"spbill_create_ip,omitempty" xml:"spbill_create_ip,omitempty"`
}

//
func (this *MkttransfersReq) BuildXml(apikey string) (result string, err error) {
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
	var keys []string
	for key, v := range mapResult {
		//空值和sign字段不参与签名

		if key != "sign" {
			if v == "" {
				err = errors.New(key + "字段为空")
				return
			} else {
				keys = append(keys, key)
			}

		}
	}
	sort.Strings(keys)
	strArr := make([]string, 0)
	for _, k := range keys {
		strArr = append(strArr, k+"="+mapResult[k])
	}
	//fmt.Println("str = " + strings.Join(strArr, "&"))
	strArr = append(strArr, "key="+apikey)
	strA := strings.Join(strArr, "&")
	sign := utils.Md5Encode(strA)
	//fmt.Println("sign = " + sign)
	mapResult["sign"] = sign
	result, err = utils.Map2Xml(mapResult)
	return result, err
}

//发送红包
func (this *MkttransfersReq) DispatchWithp12(apikey string, p12file string) (result MkttransfersResp, err error) {
	postxml, err := this.BuildXml(apikey)
	if err != nil {
		return
	}
	client, err := utils.WxPayGetCertHttpClientUsep12(p12file, this.Mchid)
	if err != nil {
		return
	}
	return this.dispatch(apikey, postxml, client)
}

//发送红包
func (this *MkttransfersReq) dispatch(apikey string, postxml string, client *http.Client) (result MkttransfersResp, err error) {

	xmlresp, err := utils.HttpPostXmlWithCert("https://api.mch.weixin.qq.com/mmpaymkttransfers/promotion/transfers", postxml, client)

	var resp MkttransfersResp
	err = xml.Unmarshal(xmlresp, &resp)
	if err != nil {
		return
	}
	if resp.ReturnCode != "SUCCESS" {
		err = errors.New(resp.ReturnMsg)
		return
	}
	if resp.ResultCode != "SUCCESS" {
		err = errors.New(resp.ErrCodeDes)
		return
	}

	return resp, nil
}

//发送红包
func (this *MkttransfersReq) DispatchWithkey(apikey string, certfile, keypath, rootca string) (result MkttransfersResp, err error) {
	postxml, err := this.BuildXml(apikey)
	if err != nil {
		return
	}
	client, err := utils.WxPayGetCertHttpClientUseKey(certfile, keypath, rootca)
	if err != nil {
		return
	}
	return this.dispatch(apikey, postxml, client)
}

//WxPaySendRedPackResponse 微信发送红包返回值
type MkttransfersResp struct {
	XMLName    xml.Name `xml:"xml"`
	ReturnCode string   `json:"return_code" xml:"return_code"` // SUCCESS/FAIL 此字段是通信标识,非交易标识,交易是否成功需要查看result_code来判断
	ReturnMsg  string   `json:"return_msg" xml:"return_msg"`   // 返回信息,如非空,为错误原因:签名失败/参数格式校验错误

	//以下字段在return_code为SUCCESS的时候有返回
	ResultCode     string `json:"result_code" xml:"result_code"`   // SUCCESS/FAIL
	ErrCode        string `json:"err_code" xml:"err_code"`         // 详细参见第6节错误列表
	ErrCodeDes     string `json:"err_code_des" xml:"err_code_des"` // 错误返回的信息描述
	MchAppid       string `json:"mch_appid" xml:"mch_appid"`
	Mchid          string `json:"mchid" xml:"mchid"`
	DeviceInfo     string `json:"device_info,omitempty" xml:"device_info,omitempty"`
	NonceStr       string `json:"nonce_str" xml:"nonce_str"`
	Sign           string `json:"sign" xml:"sign"`
	PartnerTradeNo string `json:"partner_trade_no" xml:"partner_trade_no"`
	PaymentNo      string `json:"payment_no" xml:"payment_no"`
	PaymentTime    string `json:"payment_time" xml:"payment_time"`
}

//发送红包
func (this *MkttransfersResp) String() string {
	return utils.JsonString(this)
}

//发送红包
func (this *MkttransfersResp) IsOk() bool {
	return this.ResultCode == "SUCCESS"
}

//发送红包
func (this *MkttransfersResp) Xml() string {
	return utils.XmlString(this)
}
