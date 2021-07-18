package args

import (
	"github.com/techidea8/restgo"
	"gorm.io/gorm"
)

type Account struct {
	restgo.PageArg
	Name   string `json:"name" form:"name"`
	Mobile string `json:"mobile" form:"mobile"`
}

//分页
func (p *Account) Conditions() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		if p.Mobile != "" {
			db = db.Where("mobile = ?", p.Mobile)
		}
		if p.Name != "" {
			db = db.Where("name like ?", "%"+p.Name+"%")
		}
		return db
	}
}
