package service

import (
	"errors"
	"time"
	"turingapp/args"
	"turingapp/global"
	"turingapp/model"

	"github.com/techidea8/restgo"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//根据手机号查找用户
func FindUserByMobile(mobile string) (result model.User, err error) {
	instance := model.User{}
	if err = global.DbEngin.Model(new(model.User)).Where("mobile = ?", mobile).First(&instance).Error; err != nil {
		return instance, err
	}
	return instance, nil
}

//根据手机号查找用户
func FindUserByOpenId(openId string) (result model.User, err error) {
	instance := model.User{}
	if err = global.DbEngin.Model(new(model.User)).Where("open_id = ?", openId).First(&instance).Error; err != nil {
		return instance, err
	}
	return instance, nil
}

//根据手机号查找用户
func FindUserById(id uint) (result model.User, err error) {
	instance := model.User{}
	if err = global.DbEngin.Model(new(model.User)).Where("id = ?", id).First(&instance).Error; err != nil {
		return instance, err
	}
	return instance, nil
}

//根据手机号查找用户
func LoginWithMobileAndPasswd(mobile, passwd string) (result model.User, err error) {
	instance := model.User{}
	eg := global.DbEngin
	if err = eg.Model(new(model.User)).Where("mobile = ?", mobile).First(&instance).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return instance, errors.New("用户不存在")
		} else {
			global.Log.Error(err.Error())
			return instance, err
		}

	}
	if err = ComparePasswd(instance.Passwd, passwd); err != nil {
		return instance, errors.New("密码不正确")
	}

	return instance, nil
}

//加密密码
func BuildPasswd(plainpasswd string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(plainpasswd), bcrypt.DefaultCost)
	return string(hash)
}

//密码比较
func ComparePasswd(encodedpwd, plainpasswd string) error {
	err := bcrypt.CompareHashAndPassword([]byte(encodedpwd), []byte(plainpasswd))
	return err
}

//创建新的设备
func CreateUserFromSnsUserInfo(property args.WxapiSnsUserinfo, role model.RoleType) (result model.User, err error) {
	var user model.User
	user.HeadImage = property.HeadImage
	user.OpenId = property.Openid
	user.NickName = property.Nickname
	user.Passwd = BuildPasswd(property.Openid)
	user.Score = 0
	user.Role = role
	user.ExpireAt = restgo.DateTimeFromTime(time.Now().AddDate(0, 0, 0))
	err = global.DbEngin.Create(&user).Error
	return user, err
}

//创建新的设备
func CreateUserFromMobile(property args.User, role model.RoleType) (result model.User, err error) {
	var user model.User
	user.Mobile = property.Mobile
	user.Passwd = BuildPasswd(property.Passwd)
	user.Score = 0
	user.Role = role
	user.ExpireAt = restgo.DateTimeFromTime(time.Now().AddDate(0, 0, 0))
	err = global.DbEngin.Create(&user).Error
	return user, err
}

//创建新的设备
func CreateUser(property model.User) (result model.User, err error) {
	if property.Role == 0 {
		err = errors.New("缺少用户角色")
		return
	}
	err = global.DbEngin.Create(&property).Error
	return property, nil
}

//创建新的设备
func GetOneUser(property model.User) (result model.User, err error) {
	err = global.DbEngin.Where("id = ?", property.Id).Find(&property).Error
	return property, nil
}


//更新设备
func UpdateMobile(property model.User) (result model.User, err error) {
	err = global.DbEngin.Where("id = ?", property.Id).Update("mobile",property.Mobile).Error
	return property, err
}


//更新设备
func UpdateUser(property model.User) (result model.User, err error) {
	err = global.DbEngin.Where("id = ?", property.Id).Updates(&property).Error
	return property, err
}

//删除
func EnableUser(property model.User) (result model.User, err error) {
	err = global.DbEngin.Where("id = ?", property.Id).Update("deleted",0).Error
	return property, nil
}

//删除
func DisableUser(property model.User) (result model.User, err error) {
	err = global.DbEngin.Where("id = ?", property.Id).Update("deleted",1).Error
	return property, nil
}

//删除
func DeleteUser(property model.User) (result model.User, err error) {
	return DisableUser(property)
}

//搜索支持分页
func SearchUser(arg args.User) (result []model.User, total int64, err error) {
	objs := make([]model.User, 0)

	err = global.DbEngin.Model(new(model.User)).Scopes(arg.Paginate(), arg.Condtions()).Order(arg.Sort()).Find(&objs).Error
	total = int64(arg.Total)
	if arg.Total == -1 {
		global.DbEngin.Model(new(model.User)).Scopes(arg.Condtions()).Count(&total)
	}
	return objs, total, nil
}
