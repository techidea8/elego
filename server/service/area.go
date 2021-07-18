package service

import (
	"errors"
	"turingapp/args"
	"turingapp/global"
	"turingapp/model"
)

//根据手机号查找用户
func FindAreaByCityCode(cityCode string) (result model.Area, err error) {
	instance := model.Area{}
	if err = global.DbEngin.Model(new(model.Area)).Where("city_code = ?", cityCode).First(&instance).Error; err != nil {
		return instance, err
	}
	return instance, nil
}

//根据手机号查找用户
func FindAreaByName(name string) (result model.Area, err error) {
	instance := model.Area{}
	if err = global.DbEngin.Model(new(model.Area)).Where("name = ?", name).First(&instance).Error; err != nil {
		return instance, err
	}
	return instance, nil
}

//根据手机号查找用户
func FindAreaById(id uint) (result model.Area, err error) {
	instance := model.Area{}
	if err = global.DbEngin.Model(new(model.Area)).Where("id = ?", id).First(&instance).Error; err != nil {
		return instance, err
	}
	return instance, nil
}

//创建新的设备
func CreateArea(property model.Area) (result model.Area, err error) {
	if property.Name == "" {
		err = errors.New("缺少区域名称")
		return
	}
	if property.Pid == 0 {
		err = errors.New("缺少父级区域")
		return
	}
	err = global.DbEngin.Create(&property).Error
	return property, nil
}

//更新设备
func UpdateArea(property model.Area) (result model.Area, err error) {
	err = global.DbEngin.Where("id = ?", property.Id).Updates(&property).Error
	return property, err
}

//删除
func DeleteArea(property model.Area) (result model.Area, err error) {
	err = global.DbEngin.Where("id = ?", property.Id).Delete(&property).Error
	return property, nil
}

//搜索支持分页
func SearchArea(arg args.Area) (result []model.Area, total int64, err error) {
	objs := make([]model.Area, 0)

	err = global.DbEngin.Model(new(model.Area)).Scopes(arg.Paginate(), arg.Condtions()).Order(arg.Sort()).Find(&objs).Error
	total = int64(arg.Total)
	if arg.Total == -1 {
		global.DbEngin.Model(new(model.Area)).Scopes(arg.Condtions()).Count(&total)
	}
	return objs, total, nil
}
