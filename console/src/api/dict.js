
import request from '@/utils/request'

const prefix = "/dict"

//查询,搜索
export function search(data) {
  return request({
    url: prefix +'/search',
    method: 'post',
    data:data
  })
}
//创建对象
export function create(data) {

  return request({
    url: prefix+'/create',
    method: 'post',
    data
  })
}
//更新
export function update(data) {

  return request({
    url: prefix+'/update',
    method: 'post',
    data
  })
}
const valuemap = {}
//获取一条记录
export function getOne(kword) {

  if(typeof valuemap[kword]=="undefined"){
    let data = {}
    if (/^\d+$/.test(kword)){
      data.id = +kword
    }else{
      data.name = kword
    }
    return new Promise((resolve,reject)=>{
      request({
        url: prefix+'/getone',
        method: 'post',
        data
      }).then(res=>{
        valuemap[kword] = res.data
        resolve(res.data)
      })
    })
    
  }else{
    return Promise.resolve(valuemap[kword])
  }
}

//删除某一条记录
export function deleteIt(id) {

  return request({
    url: prefix+'/delete',
    method: 'post',
    data:{
      id
    }
  })
}

//默认输出api对象
export default {search,create,update,deleteIt,getOne}

