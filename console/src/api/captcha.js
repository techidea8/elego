const prefix = "/captcha"
import request from '@/utils/request'
export function captcha(data){
 
  return request({
    url: prefix+'/context',
    method: 'post',
    data
  })
}
