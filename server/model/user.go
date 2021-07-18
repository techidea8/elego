package model

import "github.com/techidea8/restgo"

//用户基本信息表
type User struct {
	restgo.BaseModel
	HeadImage string          `json:"headImage" form:"headImage" gorm:"type:varchar(150);comment:用户头像"`
	NickName  string          `json:"nickName" form:"nickName" gorm:"type:varchar(32);comment:用户昵称"`
	Mobile    string          `json:"mobile" form:"mobile" gorm:"type:varchar(13);index:mobile;comment:用户手机号"` //激活区域`
	OpenId    string          `json:"openId" form:"openId" gorm:"type:varchar(40);index:openid;comment:用户的OPENID"`
	Passwd    string          `json:"-" form:"passwd" gorm:"type:varchar(120);comment:用户密码"`
	Role      RoleType        `json:"role" form:"tole" gorm:"type:int(11);comment:7超管;3店员;1会员"`
	Charge    int             `json:"charge" form:"charge" gorm:"type:int(11);default:0;comment:充钱金额"`
	Score     int             `json:"score" form:"score" gorm:"type:int(11);default:0;comment:积分"`
	ExpireAt  restgo.DateTime `json:"expireAt" form:"expireAt" gorm:"type:datetime;comment:过期时间"` //黄金会员停用时间
}

type RoleType int

const (
	RoleAdmin  RoleType = 4 //0b111
	RoleUser   RoleType = 3 //0b11
	RoleVip RoleType = 2 //黑金会员
	RoleMember RoleType = 1 //普通会员
)

//变成admin
func (this *User) AsAdmin() *User {
	this.Role = RoleAdmin
	return this
}

func (this *User) AsUser() *User {
	this.Role = RoleUser
	return this
}


func (this *User) AsVip() *User {
	this.Role = RoleVip
	return this
}

func (this *User) AsMember() *User {
	this.Role = RoleMember
	return this
}
