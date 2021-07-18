package args

import (
	"github.com/techidea8/restgo"
	"gorm.io/gorm"
	"turingapp/model"
)

type User struct {
	restgo.PageArg
	Total  int  `json:"total" form:"total"`
	Mobile string `json:"mobile" form:"mobile"`
	Passwd string `json:"passwd" form:"passwd"`
	Role model.RoleType `json:"role" form:"role"`
	UserId uint `json:"userId" form:"userId"`
}

//分页
func (p *User) Condtions() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		if !p.Datefrom.IsZero() {
			db = db.Where("create_at  >= ?", p.Datefrom.String())
		}
		if !p.Dateto.IsZero() {
			db = db.Where("create_at  < ?", p.Dateto.String())
		}
		if p.Kword != "" {
			db = db.Where("nick_name  like  ? ", "%"+p.Kword+"%")
		}
		if p.Mobile != "" {
			db = db.Where("mobile = ?", p.Mobile)
		}
		if p.Role>0 {
			db = db.Where("role = ?", p.Role)
		}
		return db
	}
}
