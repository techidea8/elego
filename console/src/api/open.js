import request from '@/utils/request'

const prefix = "/open"

export function sendsms(data) {
  
  return request({
    url: prefix+'/sendsms',
    method: 'post',
    data
  })
}


export function apps(data) {
  
  return request({
    url: prefix+'/apps',
    method: 'post',
    data
  })
}


export function auth(data) {
  
  return request({
    url: prefix+'/auth',
    method: 'post',
    data
  })
}