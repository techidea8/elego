package args

import (
	"github.com/techidea8/restgo"
	"gorm.io/gorm"
)

type Weixin struct {
	restgo.PageArg
	Total      int    `json:"total" form:"total"`

	AppId   string `json:"appId" form:"appId"`
}

//分页
func (p *Weixin) Condtions() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		if !p.Datefrom.IsZero() {
			db = db.Where("create_at  >= ?", p.Datefrom.String())
		}
		if !p.Dateto.IsZero() {
			db = db.Where("create_at  < ?", p.Dateto.String())
		}
		if p.Kword != "" {
			db = db.Where("(name  like )", "%"+p.Kword+"%")
		}
		if p.AppId !=""  {
			db = db.Where("app_id = ? ", p.AppId)
		}

		return db
	}
}
