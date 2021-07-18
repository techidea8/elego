<template>
  <div class="app-container">
    <el-row :gutter="24">
      <el-header>
          <el-form :inline="true" v-model="cond"  class="demo-form-inline">
        <el-form-item label="时间范围">
          <el-date-picker
              v-model="daterange"
              type="daterange"
              align="right"
              unlink-panels
              range-separator="至"
              start-placeholder="开始日期"
              end-placeholder="结束日期"
              :default-time="['00:00:00', '23:59:59']"
              :picker-options="pickerOptions">
          </el-date-picker>
        </el-form-item>
         <el-form-item>

        </el-form-item>


          <el-form-item>
            <el-button type="primary" @click="resetcond" >搜索</el-button>
          </el-form-item>

          <el-form-item  id="tipsite" >
                 <el-button type="primary" @click="showscript = true" >网站接入参数</el-button>
          </el-form-item>

          <el-form-item  id="tiplite" >
                 <el-button type="primary" @click="resetcond" >小程序接入指南</el-button>
          </el-form-item>

          <el-form-item  id="tipwap" >
                 <el-button type="primary" @click="showwx=true" >微网站接入参数</el-button>
          </el-form-item>
                       <router-link to="/corp/kefu">
                         <el-button type="primary" id="tipkefu" >配置客服</el-button>
                         </router-link>
                          <router-link to="/corp/charge">
                         <el-button type="primary" id="tipcharge" >充值续费</el-button>
                         </router-link>
                         <el-button type="danger" @click="$navigate(`http://kefu.techidea8.com/html/wiki/part1/kfapp.html`)">下载客服APP</el-button>

        </el-form>

           </el-header>
    </el-row>
      <el-row :gutter="40">
        <el-col :xs="12" :sm="12" :lg="6" class="card-panel-col">
          <router-link to="/merchant/routine">
          <card-item svg="zixun" bgcolor="#40c9c6"  color="#fff" server="report/routinenum" :cond="cond" title="新增咨询个数" ></card-item>
          </router-link>
        </el-col>  
        <el-col :xs="12" :sm="12" :lg="6" class="card-panel-col">
             <router-link to="/merchant/forms">
          <card-item svg="form" bgcolor="#36a3f7" color="#fff" server="report/formnum" :cond="cond"  title="新增表单个数" ></card-item>
             </router-link>
        </el-col>  
        <el-col :xs="12" :sm="12" :lg="6" class="card-panel-col">
             <router-link to="/merchant/routine">
          <card-item svg="zixun" bgcolor="#f4516c" color="#fff" server="report/routinenum"  title="累计咨询个数" ></card-item>
             </router-link>
          </el-col>  
        <el-col :xs="12" :sm="12" :lg="6" class="card-panel-col">
             <router-link to="/merchant/forms">
          <card-item svg="form" bgcolor="#34bfa3" color="#fff" server="report/formnum" title="累计表单个数" ></card-item>
             </router-link>
         </el-col>  
     
      </el-row>
      

      <el-row :gutter="16">
      
          <el-card class="box-card">
            <div slot="header" class="clearfix">
              <span>最新咨询</span>
              <router-link :to="`/merchant/routine`">
              <el-button style="float: right; padding: 3px 0" type="text">咨询管理</el-button>
              </router-link>
            </div>


           <el-table  :data="dataList" :stripe="true" border>
      <el-table-column align="header-center" fixed="left" label="#ID">
        <template slot-scope="scope">
          {{ scope.row.id }}

        </template>
      </el-table-column>
      <el-table-column align="header-center" fixed="left" label="描述" >
        <template slot-scope="scope">
          {{ scope.row.memo }}

        </template>
      </el-table-column>
      <el-table-column align="header-center" fixed="left" label="访客IP">
        <template slot-scope="scope">
          {{ scope.row.ip }}
        </template>
      </el-table-column>

       <el-table-column align="header-center" fixed="left" label="访客来源">
        <template slot-scope="scope">
          {{ scope.row.city }}
        </template>
      </el-table-column>
      <el-table-column align="header-center" fixed="left" label="kefu">
        <template slot-scope="scope">
             <div v-if="!!scope.row.kfid">
          <kefu :mid="scope.row.kfid"></kefu>
          </div>
        </template>
      </el-table-column>
     <el-table-column align="header-center" fixed="left" label="咨询时长(秒)">
        <template slot-scope="scope">
          {{ scope.row.ttl |timetip }}
        </template>
      </el-table-column>
      <el-table-column align="header-center" label="咨询状态">
        <template slot-scope="scope">
          {{ scope.row.stat | getStat }}
        </template>
      </el-table-column>
      <el-table-column align="header-center" label="创建时间">
        <template slot-scope="scope">
          {{ scope.row.createat }}
        </template>
      </el-table-column>


    </el-table>

          </el-card>

      
      
      </el-row>
    <el-dialog title="请复制如下内容" :visible.sync="showscript" center="" width="660px">
           <el-input
          type="textarea"
          :autosize="{ minRows: 4, maxRows: 10}"
          placeholder="请输入内容"
          v-model="txtcode">
        </el-input>
            <span slot="footer" class="dialog-footer">
               <el-button type="primary" v-clipboard:copy="txtcode"
                        v-clipboard:success="onCopy"
                        v-clipboard:error="onError" 
                        >复制</el-button>
                <el-button type="primary" @click="showscript = false">关闭</el-button>
            </span>
    </el-dialog>

    <el-dialog title="请复制如下内容" :visible.sync="showwx" center="" width="660px">
           <el-input
          type="textarea"
          :autosize="{ minRows: 2, maxRows: 4}"
          placeholder="请输入内容"
          v-model="txturl">
        </el-input>
            <span slot="footer" class="dialog-footer">
               <el-button type="primary" v-clipboard:copy="txturl"
                        v-clipboard:success="onCopy"
                        v-clipboard:error="onError" 
                        >复制</el-button>
                <el-button type="primary" @click="showwx = false">关闭</el-button>
            </span>
    </el-dialog>

    <el-dialog title="请扫码" :visible.sync="showqrcode" center="" width="315px">
            <div id="qrcode"></div>
            <span slot="footer" class="dialog-footer">
                <el-button type="primary" @click="showqrcode = false">关闭</el-button>
            </span>
    </el-dialog>

 


  </div>
</template>

<script>




import {qrcodeUrl,getOne,activate,kefuIndex} from "@/api/account"
import {search} from "@/api/account"
import {dateFormat,pickerOptions } from "@/utils"
import QRCode from 'qrcodejs2'

import Driver from 'driver.js' // import driver.js
import 'driver.js/dist/driver.min.css' // import driver.js css
import steps from './steps'

const stats=[
  {"name":1,"label":"待指派"},
  {"name":2,"label":"处理中"},
  {"name":3,"label":"已完成"},
  {"name":99,"label":"订单已删除"},
]
export default {
  name: 'DashboardEditor',
  data() {
    let corpid = this.$store.getters.corpId
    return {

      emptyGif: 'https://wpimg.wallstcn.com/0e03b7da-db9e-4819-ba10-9016ddfdaed3',
      daterange:[new Date(dateFormat(new Date(),"yyyy-MM-dd 00:00:00")),new Date(dateFormat(new Date(),"yyyy-MM-dd 23:59:59"))],
      driver: {},
       dataList:[],
      pickerOptions:pickerOptions,
     
       showqrcode:false,
       showdownloadqrcode:false,
       shopUrl:'',
       corpInfo:{logo:""},
      cond:{
        datefrom:dateFormat(new Date(),"yyyy-MM-dd 00:00:00"),
        dateto:dateFormat(new Date(),"yyyy-MM-dd 23:59:59")
      },
      showscript:false,
      showwx:false,
      txturl:"",
      txtcode:"",
       sortArg:{
        order:"desc",
        prop:"numsaled"
      },
    }
  },
  filters:{
  getStat(v){
        let tmp = {}
        stats.forEach(o=>{
            if(o.name==v){
              tmp = o
            }
        })
        return tmp.label
    },
    timetip(ttl){
        if(ttl<60){
          return ttl +"秒"
        }else if(ttl<3600){
          return (ttl - ttl%60)/60+"分"+ ttl%60 +"秒"
        }else{
          return ttl /3600 +"小时"
        }
    },


  },
  created(){

     
     this.shopUrl = kefuIndex(this.$store.getters.corpId)
     getOne(this.$store.getters.corpId).then(res=>{
       this.corpInfo = res.data
     })
     this.initscript()
     this. loadroutines()
  },
  mounted() {
    let version = this.$version()
    let process = localStorage.getItem("version"+version)
    if(!!process){
      return 
    }
      this.driver = new Driver({
        doneBtnText: '确定',              // Text on the final button
      closeBtnText: '关闭',            // Text on the close button for this step
     onReset: (el) => {
       localStorage.setItem("version"+version,1)
     },
      nextBtnText: '下一项',              // Next button text for this step
      prevBtnText: '上一项',          // Previous button text for this step
      })
      this.driver.defineSteps(steps)
      this.driver.start()
  },
 methods:{
   loadroutines(){
      search(this.cond).then(res=>{
        this.dataList = res.rows
      })
   },
   initscript(){
      let src = location.protocol +"//kefu.techidea8.com/asset/js/kefu.js?siteid="+this.$store.getters.corpId

      let txt1  = "<script type=\"text/javascript\" id=\"techidea8kefujs\" src=\"" + src + "\">"+"<\/script>" 
      let txt2  =   `<div style="display: flex;position: fixed;z-index: 999;right: 0px;top: 320px;">
                      <div onclick="showkefu()" style="display: flex;justify-content: center;align-items: center;width: 60px;height: 60px;border: 1px solid rgba(0, 0, 0, .1);border-radius: 31px;box-shadow: 0 0 14px 0 rgba(0, 0, 0, .16);cursor: pointer;text-decoration: none;background-color: #00a6f1;;">
                          <img style="width:42px;height:42px;" src="data:image/png;base64,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" />
                      </div>
                  </div>`
         this.txtcode =txt1+txt2   
         this.txturl =       location.protocol +"//visitor.kefu.techidea8.com/pages/index/wap?siteid="+this.$store.getters.corpId
   },
  
    sortChange({ prop, order}){
       console.log(order,prop)
       if(order==""){
         order = "descending"
         prop="sortIndex"
       }
       if(order=="descending"){
         this.sortArg.order="desc"
         this.sortArg.prop=prop
       }
       if(order=="ascending"){
         this.sortArg.order="asc"
         this.sortArg.prop=prop
       }
     
    },
   onCopy(){
      this.$ok('已经复制到你的剪切板')
   },
   onError(e){
     this.$error(e)
   },
  qrcode(scope){

        if(this.corpInfo.numkefu<1){
          this.$confirm("当前可用客服数量为0,请先配置客服").then(res=>{
            
            this.$router.push("/merchant/kefu")
          })
          return 
        }
        //这里弹出微信展示
        this.showqrcode = true
        //绑定的url内容 在哪里?
        let url = this.shopUrl
        console.log("url",url)
        this.$nextTick(function () {
                    document.getElementById("qrcode").innerHTML = "";
                    let qrcode = new QRCode("qrcode", {
                        width: 250,
                        height: 250,
                        text: url,
                        colorDark: "#109dff",
                        colorLight: "#d9d9d9"
                    });
                });

    },

   
   resetcond(){
      let cond = {}
      if(this.daterange.length==2){
          cond.datefrom = dateFormat(this.daterange[0],"yyyy-MM-dd 00:00:00")
          cond.dateto = dateFormat(this.daterange[1],"yyyy-MM-dd 00:00:00")
          this.cond =Object.assign({},cond)
      }
      this.loadroutines()

   }
 }

}
</script>

