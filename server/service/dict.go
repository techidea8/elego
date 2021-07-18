
package service

import (
	"sync"
	"turingapp/args"
	"turingapp/global"
	"turingapp/model"
	
)

var dictCache sync.Map

//创建新记录
func CreateDict(property model.Dict) (result model.Dict, err error) {
	err = global.DbEngin.Create(&property).Error
	if(err==nil){
		dictCache.Store(result.Name,result)
	}
	return property, nil
}

//更新某一条记录
func UpdateDict(property model.Dict) (result model.Dict, err error) {
	err = global.DbEngin.Where("id = ?", property.Id).Updates(&property).Error
	if(err==nil){
		dictCache.Store(result.Name,result)
	}
	return property, err
}

//删除某一条记录
func DeleteDict(property model.Dict) (result model.Dict, err error) {
	err = global.DbEngin.Where("id = ?", property.Id).Updates(&property).Error
	if err!=nil{
		return
	}
	err = global.DbEngin.Where("id = ?", property.Id).Delete(&model.Dict{}).Error
	if(err==nil){
		dictCache.Delete(property.Name)
	}
	return property, nil
}


//查询一条记录
func FindDictById(id uint) (result model.Dict, err error) {
	result = model.Dict{}
	err = global.DbEngin.Model(new(model.Dict)).Where("id = ?", id).First(&result).Error
	if(err==nil){
		dictCache.Store(result.Name,result)
	}
	return result, err
}

//查询一条记录
func FindDictByName(name string) (result model.Dict, err error) {
	result1,ok:=dictCache.Load(name)
	if ok{
		return  result1.(model.Dict),nil
	}
	result = model.Dict{}
	err = global.DbEngin.Model(new(model.Dict)).Where("name = ?", name).First(&result).Error
	if(err==nil){
		dictCache.Store(result.Name,result)
	}
	return result, err
}
//搜索支持分页
func SearchDict(arg args.Dict) (result []model.Dict, total int64, err error) {
	objs := make([]model.Dict, 0)
	err = global.DbEngin.Model(new(model.Dict)).Scopes(arg.Paginate(), arg.Condtions()).Order(arg.Sort()).Find(&objs).Error
	total = 0
	if arg.Total == -1 {
		global.DbEngin.Model(new(model.Dict)).Scopes(arg.Condtions()).Count(&total)
	}
	return objs, total, nil
}





