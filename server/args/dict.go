package args

import (
	"github.com/techidea8/restgo"
	"gorm.io/gorm"
)

type Dict struct {
	restgo.PageArg
	Total int `json:"total" form:"total"`

	restgo.BaseModel

	Name string `json:"name" form:"name" gorm:"comment:'所属分类';type:varchar(30)"`

	Label string `json:"label" form:"label" gorm:"comment:'字段名称';type:varchar(255)"`

	Value string `json:"value" form:"value" gorm:"comment:'状态值';type:varchar(1024)"`
}

//分页
func (p *Dict) Condtions() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if !p.Datefrom.IsZero() {
			db = db.Where("create_at  >= ?", p.Datefrom.String())
		}
		if !p.Dateto.IsZero() {
			db = db.Where("create_at  &lt; ?", p.Dateto.String())
		}
		return db
	}
}

//结束
