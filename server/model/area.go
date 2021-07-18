package model

import "github.com/techidea8/restgo"

//产品信息表,绑定的红包信息
type Area struct {
	restgo.BaseModel
	CityCode string            `json:"cityCode" form:"cityCode" gorm:"type:varchar(50);comment:城市代码"`            //属于哪个产品
	TelCode  string            `json:"telCode" form:"telCode" gorm:"type:varchar(10);comment:电话代码"`              //哪个用户领奖了
	Level   AreaLevel          `json:"level" form:"level" gorm:"type:int(11);comment:城市等级"`                //哪个店员引导的
	Name      string             `json:"name" form:"name" gorm:"type:varchar(50);comment:城市名称"`                         //总金额
	Center        string          `json:"center" form:"center" gorm:"type:varchar(32);comment:城市中心经纬度"`                          //溯源码
	Pid      int64          `json:"pid" form:"pid" gorm:"type:bigint(20);index:pid;comment:父级ID"`     //单品序列号
}

//红包状态
type  AreaLevel int
const (
	AreaLevelProvice AreaLevel = 1     //省
	AreaLevelCity AreaLevel = 2      //市
	AreaLevelDistrict AreaLevel = 3   //区
	AreaLevelStreet AreaLevel = 4  //街道
)