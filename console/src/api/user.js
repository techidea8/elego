import request from '@/utils/request'

const prefix = "/user"



export function auth(data) {
  
  return request({
    url: prefix+'/auth',
    method: 'post',
    data
  })
}

export function search(name) {
  let cond = {} 
  if(typeof name=="object"){
      cond = name
  }else{
    cond.name = name
  }
  return request({
    url: prefix +'/search',
    method: 'post',
    data:cond
  })
}

export function create(data) {
  
  return request({
    url: prefix+'/create',
    method: 'post',
    data
  })
}

export function disable(data) {
  
  return request({
    url: prefix+'/disable',
    method: 'post',
    data
  })
}

export function enable(data) {
  
  return request({
    url: prefix+'/enable',
    method: 'post',
    data
  })
}

export function editPwd(data) {
  
  return request({
    url: prefix+'/editPwd',
    method: 'post',
    data
  })
}

export function update(data) {
  
  return request({
    url: prefix+'/update',
    method: 'post',
    data
  })
}


export function deleteIt(id) {
  
  return request({
    url: prefix+'/delete',
    method: 'post',
    data:{
      id
    }
  })
}

export function register(data) {
  
  return request({
    url: prefix+'/register',
    method: 'post',
    data
  })
}

const userMap = {}
export function snsinfo(id) {
  return new Promise((resolve,reject)=>{
    if(!userMap[id]){
      request({
        url: prefix+'/snsinfo',
        method: 'post',
        data:{id}
      }).then(res=>{
        if(res.data.nickName){
          userMap[id] = res.data
        }
        resolve(res.data)
      })  
    }else{
      resolve(userMap[id])
    }    
  })
  
  
}
