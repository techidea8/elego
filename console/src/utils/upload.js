import axios from 'axios'
import { MessageBox, Message } from 'element-ui'
import store from '@/store'
import { getToken } from '@/utils/auth'

// create an axios instance
const service = axios.create({
  baseURL: process.env.VUE_APP_BASE_API, // url = base url + request url
  // withCredentials: true, // send cookies when cross-domain requests
  timeout: 5000,// request timeout,
  headers: {"Content-Type":"multipart/form-data","platform":"console"},
  validateStatus: function (status) {
    return status >= 200 && status < 600; // 默认的
  },
})

// request interceptor
service.interceptors.request.use(
  config => {
    // do something before request is sent

    if (store.getters.token) {
      // let each request carry token
      // ['X-Token'] is a custom headers key
      // please modify it according to the actual situation
      config.headers['X-Token'] = getToken()
    }
    return config
  },
  error => {
    // do something with request error
    console.log(error) // for debug
    return Promise.reject(error)
  }
)
let empptdata = {
  "code":200,
  "data":{},
  "rows":[],
  "total":0
}
// response interceptor
service.interceptors.response.use(
  /**
   * If you want to get http information such as headers or status
   * Please return  response => response
  */

  /**
   * Determine the request status by custom code
   * Here is just an example
   * You can also judge the status by HTTP Status Code
   */
  response => {
    
    let statuscode = response.status

    if (statuscode> 400 && statuscode <= 403) {
        // to re-login
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
    }else if(statuscode<=300 && statuscode>=200){
      let res = response.data
        if(!res){
          res = empptdata
    }
    // if the custom code is not 20000, it is judged as an error.
    if (res.code == 200) {
      return res
    }else{
      return Promise.reject(new Error(res.msg))
    }
    }else{
      return Promise.reject(new Error("服务器繁忙,请稍后"))
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
