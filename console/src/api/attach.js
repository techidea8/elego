import request from '@/utils/upload'

const prefix = "/attach"


export function upload(formdata) {
  return request({
    url: prefix +'/upload',
    method: 'post',
    data:formdata,
    timeout:1000*60,
  })
}

export function hostUrl() {
  return   process.env.VUE_APP_BASE_API + prefix +'/upload'
}