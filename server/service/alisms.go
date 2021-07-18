package service

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"turingapp/global"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

func SendRegisterSms(mobile, code string) (err error) {
	//5分钟内有效
	global.Redis.Set("code-"+mobile, code)
	global.Redis.SetKeyExpire("code-"+mobile, 5*60)
	ttl := global.Redis.Ttl("ttl-" + mobile)
	if 0 < ttl {
		return errors.New("请" + fmt.Sprintf("%d", ttl) + "秒后再试")
	}
	global.Redis.Set("ttl-"+mobile, "1")
	global.Redis.SetKeyExpire("ttl-"+mobile, 60)
	return alisms(mobile, global.Conf.Sms.Tplcode, global.Conf.Sms.SignName, global.Conf.Aliyun.AccessKey, global.Conf.Aliyun.SecretKey, "{\"code\":\""+code+"\"}")
}

func VerifySms(mobile, code string) (r bool, err error) {
	result, err := global.Redis.Get("code-" + mobile)
	if result == "" || err != nil {
		err = errors.New("验证码已过期")
		return
	}
	global.Redis.SetKeyExpire(mobile, -1)
	if result != code {
		err = errors.New("验证码不正确")
		return
	}
	return true, nil
}

func SendNotifySms(mobile, about string) (err error) {

	return alisms(mobile, global.Conf.Sms.Tplcode, global.Conf.Sms.SignName, global.Conf.Aliyun.AccessKey, global.Conf.Aliyun.SecretKey, "{\"about\":\""+about+"\"}")
}
func alisms(mobile, tpl, sign, ak, sk, param string) (err error) {
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", ak, sk)
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.PhoneNumbers = mobile
	request.SignName = sign
	request.TemplateCode = tpl
	request.TemplateParam = param
	response, err := client.SendSms(request)
	if err != nil {
		global.Log.Error("短信发送失败", zap.Error(err))
		return err
	}
	if response.Code == "OK" {
		return nil
	} else {
		global.Log.Error("短信发送失败", zap.String("sms", response.Message))
		return errors.New(response.Message)
	}
}
