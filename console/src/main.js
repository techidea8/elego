import Vue from 'vue'

import Cookies from 'js-cookie'

import 'normalize.css/normalize.css' // a modern alternative to CSS resets

import Element from 'element-ui'
import './styles/element-variables.scss'

import 'video.js/dist/video-js.css'
import '@/styles/index.scss' // global css

import App from './App'
import store from './store'
import router from './router'

import './icons' // icon
import './permission' // permission control
import './utils/error-log' // error log

import VueClipboard from 'vue-clipboard2'

Vue.use(VueClipboard)

import * as filters from './filters' // global filters

/**
 * If you don't want to use mock-server
 * you want to use MockJs for mock api
 * you can execute: mockXHR()
 *
 * Currently MockJs will be used in the production environment,
 * please remove it before going online ! ! !
 */
if (process.env.NODE_ENV === 'production') {
  const { mockXHR } = require('../mock')
  mockXHR()
}

Vue.use(Element, {
  size: Cookies.get('size') || 'medium' // set element-ui default size
})

// register global utility filters
Object.keys(filters).forEach(key => {
  Vue.filter(key, filters[key])
})

Vue.config.productionTip = false
import { Message } from 'element-ui';
Vue.prototype.$error=(msg)=>{
  Message.close()
  Message(
    {
      message: msg,
      type: 'error'
    }
  )
}
Vue.prototype.$success=(msg)=>{
  Message(
    {
      message: msg,
      type: 'success'
    }
  )
}
Vue.prototype.$ok=(msg)=>{
  Message(
    {
      message: msg,
      type: 'success'
    }
  )
}

Date.prototype.format = function(fmt) {//author: meizz
  var o = {
    "M+" : this.getMonth()+1,                 //月份
    "d+" : this.getDate(),                    //日
    "h+" : this.getHours(),                   //小时
    "m+" : this.getMinutes(),                 //分
    "s+" : this.getSeconds(),                 //秒
    "q+" : Math.floor((this.getMonth()+3)/3), //季度
    "S"  : this.getMilliseconds()             //毫秒
  };
  if(/(y+)/.test(fmt))
    fmt=fmt.replace(RegExp.$1, (this.getFullYear()+"").substr(4 - RegExp.$1.length));
  for(var k in o)
    if(new RegExp("("+ k +")").test(fmt))
  fmt = fmt.replace(RegExp.$1, (RegExp.$1.length==1) ? (o[k]) : (("00"+ o[k]).substr((""+ o[k]).length)));
  return fmt;
}

import VDistpicker from 'v-distpicker'

Vue.component('city-picker', VDistpicker)

import DataItem from "@/components/DataItem"
Vue.component('data-item',DataItem)

import LineChart from "@/components/LineChart"
Vue.component('line-chart',LineChart)


import PickImage from "@/components/PickImage/index"
Vue.component('pick-image',PickImage)

import StateSelect from "@/components/StateSelect/index"
Vue.component('state-select',StateSelect)

import State from "@/components/State/index"
Vue.component('state',State)

import Block from "@/components/Block/block"
Vue.component('block',Block)

import User from "@/components/User/index"
Vue.component('user',User)

new Vue({
  el: '#app',
  router,
  store,
  render: h => h(App)
})
