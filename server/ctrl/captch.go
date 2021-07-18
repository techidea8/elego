package ctrl

import (
	"turingapp/global"
	"turingapp/service"

	"github.com/techidea8/restgo"
	"strconv"

	"net/http"

	log "github.com/cihub/seelog"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

type CaptchaCtrl struct {
	restgo.Ctrl
}

func (ctrl *CaptchaCtrl) Router() {
	mod := restgo.Module("captcha")
	mod.Router("context", ctrl.context)
}

//图片验证码
type CaptchArg struct {
	Ticket string `json:"ticket" form:"ticket"`
	Width  int    `json:"width" form:"width"`
	Height int    `json:"height" form:"height"`
	Code   string `json:"code" form:"code"`
	Len    int    `json:"len" form:"len"`
}

var store = base64Captcha.DefaultMemStore

//刷新验证码
func (ctrl *CaptchaCtrl) context(w http.ResponseWriter, req *http.Request) {
	var arg CaptchArg

	width := req.URL.Query().Get("width")
	height := req.URL.Query().Get("height")
	arg.Width, _ = strconv.Atoi(width)
	arg.Height, _ = strconv.Atoi(height)
	if arg.Width == 0 {
		arg.Width = global.Conf.Captcha.ImgWidth
	}
	if arg.Height == 0 {
		arg.Height = global.Conf.Captcha.ImgHeight
	}

	if id, b64s, err := service.GenerateCaptcha(arg.Width, arg.Height, global.Conf.Captcha.KeyLong); err != nil {
		log.Error("验证码获取失败!", zap.Any("err", err))
		ctrl.RespFail(w, "验证码获取失败")
	} else {
		ctrl.RespOkMap(w, map[string]interface{}{
			"captchaId":  id,
			"base64data": b64s,
		})
	}
}
