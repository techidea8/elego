package model

import (
	"fmt"
	"gitee.com/chunanyong/gowe"
	"github.com/techidea8/restgo"
)

//用户基本信息表
type Weixin struct {
	restgo.BaseModel
	AppId string  `json:"appId" form:"appId" gorm:"type:varchar(40);index:appid;comment:应用ID"`
	Secret  string   `json:"secret" form:"secret" gorm:"type:varchar(60);comment:密码"`
	Token    string  `json:"token" form:"token" gorm:"type:varchar(120);comment:密钥"`               //激活区域`
	MchId    string `json:"mchId" form:"mchId" gorm:"type:varchar(40);comment:商户号"`
	Cert    string `json:"cert" form:"cert" gorm:"type:varchar(120);comment:密钥路径"`
	Apikey    string `json:"apikey" form:"apikey" gorm:"type:varchar(120);comment:操作密码"`
	Name    string `json:"name" form:"name" gorm:"type:varchar(120);comment:微信名称"`
}

func (wxConfig *Weixin) GetToken() string {
	return  wxConfig.Token
}

func (wxConfig *Weixin) GetAesKey() string {
	return ""
}

func (wxConfig *Weixin) GetOauth2() bool {
	panic("implement me")
}

func (wxConfig *Weixin) GetId() string {
	return fmt.Sprint("%d",wxConfig.Id)
}

func (wxConfig *Weixin) GetAppId() string {
	return wxConfig.AppId
}

func (wxConfig *Weixin) GetAccessToken() string {
	//从缓存中获取wxAccessToken,这里只是演示
	wxAccessToken, err := gowe.GetAccessToken(wxConfig)
	if err == nil && wxAccessToken.ErrCode == 0 {
		return wxAccessToken.AccessToken
	}
	return ""
}

func (wxConfig *Weixin) GetSecret() string {
	return wxConfig.Secret
}