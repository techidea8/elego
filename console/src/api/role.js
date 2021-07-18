import request from '@/utils/request'
import { deepClone } from '@/utils/index.js'


export function getRoutes(){
  return new Promise((resolve,reject)=>{
        resolve({
          "code":200,
          "data":routes,
          "msg":"加载成功"
        })
  })
}
export function getAuth(token) {
  return request({
    url: '/role/getAuth',
    method: 'post',
    data:{
      token
    }
  })
}

export function getRoles() {
  return request({
    url: '/role/getAll',
    method: 'post'
  })
}

export function addRole(data) {
  let tmp = Object.assign({},data)
  delete tmp.id
  tmp.routes = JSON.stringify(data.routes)
  return request({
    url: '/role/create',
    method: 'post',
    data:tmp
  })
}

export function updateRole( data) {
  let tmp = Object.assign({},data)
  tmp.routes = JSON.stringify(data.routes)
  return request({
    url: '/role/update',
    method: 'post',
    data:tmp
  })
}
