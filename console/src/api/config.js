import request from '@/utils/request'

const prefix = "/config"
export function getAll(corpId) {
  return request({
    url: prefix+'/getAll',
    method: 'post',
    data:{
      corpId:0
    }
  })
}

export function search(arg) {
  return request({
    url: prefix + '/search',
    method: 'post',
    data:arg
  })
}

export function create(data) {
    console.log("create",data)
  return request({
    url: prefix+'/create',
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
