
package model
import "github.com/techidea8/restgo"
//数据库模型
type Dict struct{
	
        restgo.BaseModel
    
        Name string  `json:"name" form:"name" gorm:"comment:'所属分类';type:varchar(30)"` 
    
        Label string  `json:"label" form:"label" gorm:"comment:'字段名称';type:varchar(255)"` 
    
        Value string  `json:"value" form:"value" gorm:"comment:'状态值';type:varchar(1024)"` 
    
}

