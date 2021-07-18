import axios from 'axios'
import { MessageBox, Message } from 'element-ui'
import store from '@/store'
import { getToken } from '@/utils/auth'

const service = axios.create({
  baseURL: process.env.VUE_APP_BASE_API,
  timeout: 5000,
  // transformRequest: [function (data) {
  //   let tmp =[]
  //   for(var i in data){
  //     if(typeof data[i] != "object"){
  //       tmp.push(i + "=" + encodeURIComponent(data[i]))
  //     }
  //   }
  //  return tmp.join("&")
  // // return data
  // }],
  headers: {"Content-Type":"application/json;charset=UTF-8","platform":"console"},
  validateStatus: function (status) {
    return status >= 200 && status < 600; 
  },
})
service.interceptors.request.use(
  config => {
    if (store.getters.token) {
      config.headers['X-Token'] = getToken()
    }
    return config
  },
  error => {
    console.log(error)
    return Promise.reject(error)
  }
)
let empptdata = {
  "code":200,
  "data":{},
  "rows":[],
  "total":0
}
service.interceptors.response.use(
  response => {
    
    let statuscode = response.status
    let respdata = response.data
    if(statuscode==403){
      MessageBox.confirm('你的登陆已经过期,请重新登陆', '登陆提醒', {
        confirmButtonText: '重新登陆',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        store.dispatch('user/resetToken').then(() => {
          location.reload()
        })
      })
      return empptdata
    }else if(statuscode == 200){
      return respdata
    }else {
      let msg = respdata.msg || "服务器繁忙请稍后"
      Message({
        message: msg,
        type: 'error',
        duration: 5 * 1000
      })
      return Promise.reject(new Error(respdata.msg || "服务器繁忙,请稍后"))
    }
  },
  error => {
    Message({
      message: error.message,
      type: 'error',
      duration: 5 * 1000
    })
    return Promise.reject(error)
  }
)

export default service
