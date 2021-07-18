package args

import (
	"encoding/xml"
	"fmt"
	"testing"
	"time"
	"turingapp/utils"
)

func gmnuniordeReq() *UniordeReq {
	var req UniordeReq
	req.Appid = "wx71bcddd77824f86f"
	req.MchId = "1449761102"
	req.NonceStr = utils.FuncGenerateRandomString(16)
	req.OutTradeNo = "MKTT" + utils.FormatDateTime(time.Now(), "20060102150405") + utils.FuncGenerateRandomString(4)
	//req.Openid = "oa_rXvpo48JHICZpOgMVPUJ4Unbo" // jianlia
	req.Body="测试商品"
	req.NotifyIUrl="http://winlion.ngrok2.xiaomiqiu.cn/charge/notify"
	req.SpbillCreateIp = "220.170.56.113"
	req.TotalFee="10"
	req.TradeType="JSAPI"
	req.Openid = "oa_rXvsOT4zSOud1ymohd_pm9akc"
	return &req
}
func TestDispatch(t *testing.T) {
	req := gmnuniordeReq()
	resp, e := req.Dispatch("F2E21FF8E2aD74ABE7A4eA15A9307123")

	if e != nil {
		fmt.Printf(e.Error())
	} else {
		fmt.Println(resp)
	}
}
var xmltest=`<xml>
  <appid><![CDATA[wx2421b1c4370ec43b]]></appid>
  <attach><![CDATA[支付测试]]></attach>
  <bank_type><![CDATA[CFT]]></bank_type>
  <fee_type><![CDATA[CNY]]></fee_type>
  <is_subscribe><![CDATA[Y]]></is_subscribe>
  <mch_id><![CDATA[10000100]]></mch_id>
  <nonce_str><![CDATA[5d2b6c2a8db53831f7eda20af46e531c]]></nonce_str>
  <openid><![CDATA[oUpF8uMEb4qRXf22hE3X68TekukE]]></openid>
  <out_trade_no><![CDATA[1409811653]]></out_trade_no>
  <result_code><![CDATA[SUCCESS]]></result_code>
  <return_code><![CDATA[SUCCESS]]></return_code>
  <sign><![CDATA[B552ED6B279343CB493C5DD0D78AB241]]></sign>
  <time_end><![CDATA[20140903131540]]></time_end>
  <total_fee>1</total_fee>
  <coupon_fee><![CDATA[10]]></coupon_fee>
  <coupon_count><![CDATA[1]]></coupon_count>
  <coupon_type><![CDATA[CASH]]></coupon_type>
  <coupon_id><![CDATA[10000]]></coupon_id>
  <trade_type><![CDATA[JSAPI]]></trade_type>
  <transaction_id><![CDATA[1004400740201409030005092168]]></transaction_id>
</xml>`
func TestParseNotify(t *testing.T) {
	var req NotifyReq
	xml.Unmarshal([]byte(xmltest),&req)
	resp,e := req.ParseNotify("F2E21FF8E2aD74ABE7A4eA15A9307123")
	if e != nil {
		fmt.Printf(e.Error())
	} else {
		fmt.Println(resp)
	}
}
