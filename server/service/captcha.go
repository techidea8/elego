package service

import (
	"github.com/mojocn/base64Captcha"
)

var store = base64Captcha.DefaultMemStore

func GenerateCaptcha(w, h, l int) (id, base64 string, err error) {
	driver := base64Captcha.NewDriverDigit(h, w, l, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, base64, err = cp.Generate()
	return id, base64, err
}

//验证验证码
func VerifyCaptcha(catchaId, code string) bool {
	return store.Verify(catchaId, code, true)
}
