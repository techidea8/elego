package args

import (
	"fmt"
	"testing"
	"time"
	"turingapp/utils"
)

func gmnmkttransferreq() *MkttransfersReq {
	var req MkttransfersReq
	req.MchAppid = "wx71bcddd77824f86f"
	req.Mchid = "1449761102"
	req.NonceStr = utils.FuncGenerateRandomString(16)
	req.PartnerTradeNo = "MKTT" + utils.FormatDateTime(time.Now(), "20060102150405") + utils.FuncGenerateRandomString(4)
	//req.Openid = "oa_rXvpo48JHICZpOgMVPUJ4Unbo" // jianlian
	req.Openid = "oa_rXvsOT4zSOud1ymohd_pm9akc" //winlion
	req.CheckName = "NO_CHECK"
	req.Amount = "100"
	req.Desc = "高培会员福利活动"
	return &req
}
func TestDispatchmmttGmn(t *testing.T) {
	req := gmnmkttransferreq()
	resp, e := req.DispatchWithp12("F2E21FF8E2aD74ABE7A4eA15A9307123", "E:\\prj002\\code\\cert\\apiclient_cert.p12")

	if e != nil {
		fmt.Printf(e.Error())
	} else {
		fmt.Println(resp.Xml())
	}
}
