package args

import (
	"github.com/techidea8/restgo"
	"gorm.io/gorm"
	"turingapp/model"
)

type Area struct {
	restgo.PageArg
	Total  int  `json:"total" form:"total"`
	Name string `json:"mobile" form:"mobile"`
	Level model.AreaLevel `json:"level" form:"level"`
	Pid int `json:"pid" form:"pid"`
	Levels []int `json:"levels" form:"levels"`
}

//分页
func (p *Area) Condtions() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		if !p.Datefrom.IsZero() {
			db = db.Where("create_at  >= ?", p.Datefrom.String())
		}
		if !p.Dateto.IsZero() {
			db = db.Where("create_at  < ?", p.Dateto.String())
		}
		if p.Kword != "" {
			db = db.Where("(name  like  ? )", "%"+p.Kword+"%")
		}
		if p.Name != "" {
			db = db.Where("name = ?", p.Name)
		}
		if p.Level>0 {
			db = db.Where("level = ?", p.Level)
		}
		if p.Pid>0 {
			db = db.Where("pid = ?", p.Pid)
		}
		if len(p.Levels)>0{
			db = db.Where("level in  ?  ",p.Levels)
		}
		return db
	}
}
