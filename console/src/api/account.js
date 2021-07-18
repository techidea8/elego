import request from '@/utils/request'

let prefix = "/account"
export function login(data) {
 
  return request({
    url: prefix + '/authwithpwd',
    method: 'post',
    data
  })
}
export function register(data) {
  return request({
    url:  prefix + '/register',
    method: 'post',
    data
  })
}
export function getInfo(token) {
  return request({
    url: prefix + '/getinfo',
    method: 'post',
    data: { token }
  })
}

export function logout() {
  return request({
    url: prefix + '/logout',
    method: 'post'
  })
}