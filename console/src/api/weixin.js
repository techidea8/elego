import request from '@/utils/request'

const prefix = "/weixin"



export function search(data) {
  return request({
    url: prefix +'/search',
    method: 'post',
    data
  })
}

export function create(data) {
  
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


export default {
  search,create,update,deleteIt
}
