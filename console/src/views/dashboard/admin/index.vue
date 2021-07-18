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

                       <router-link to="/admin/member">
                         <el-button type="primary" id="tipkefu" >会员管理</el-button>
                         </router-link>
                          <router-link to="/admin/staff">
                         <el-button type="primary" id="tipcharge" >店员管理</el-button>
                         </router-link>
                          <router-link to="/admin/product">
                         <el-button type="primary" id="tipkefu" >红包配置</el-button>
                         </router-link>
                           <router-link to="/admin/item">
                         <el-button type="primary" id="tipkefu" >红包管理</el-button>
                         </router-link>
                         
        </el-form>
   </el-header>
    </el-row>
      <el-row :gutter="40" class="card-panel-row">
        <el-col :xs="12" :sm="12" :lg="6" class="card-panel-col">
          <router-link to="/admin/member">
          <data-item svg="pthy" bgcolor="#40c9c6"  color="#fff" server="report/membernum" :cond="Object.assign({},cond,{role:1})" title="新增用户(个)" ></data-item>
          </router-link>
        </el-col>  
        <el-col :xs="12" :sm="12" :lg="6" class="card-panel-col">
             <router-link to="/admin/member">
          <data-item svg="pthy" bgcolor="#36a3f7" color="#fff" server="report/membernum"  title="累计用户(个)" ></data-item>
             </router-link>
        </el-col>  
        
          <el-col :xs="12" :sm="12" :lg="6" class="card-panel-col">
             <router-link to="/admin/staff">
          <data-item svg="hjhy" bgcolor="#f4516c" color="#fff" server="report/membernum" :cond="Object.assign({},cond,{role:2})" title="新增加黑金会员(个)" ></data-item>
             </router-link>
          </el-col>  
        <el-col :xs="12" :sm="12" :lg="6" class="card-panel-col">
             <router-link to="/admin/staff">
                 <data-item svg="hjhy" bgcolor="#34bfa3" color="#fff" server="report/membernum"  :cond="Object.assign({},{role:2})" title="累计黑金会员(个)" ></data-item>
             </router-link>
         </el-col> 
        <el-col :xs="12" :sm="12" :lg="6" class="card-panel-col">
             <router-link to="/admin/item">
          <data-item svg="hbgs" bgcolor="#f4516c" color="#fff" server="report/redpkgnum" :cond="cond"  title="发放红包个数" ></data-item>
             </router-link>
          </el-col>  
        <el-col :xs="12" :sm="12" :lg="6" class="card-panel-col">
             <router-link to="/admin/item">
                 <data-item svg="hbgs" bgcolor="#34bfa3" color="#fff" server="report/redpkgnum" title="累计发放红包个数" ></data-item>
             </router-link>
         </el-col>  
          <el-col :xs="12" :sm="12" :lg="6" class="card-panel-col">
             <router-link to="/admin/item">
          <data-item svg="hbje" bgcolor="#f4516c" color="#fff" server="report/redpkgamt"  :cond="cond"   title="发放红包金额" ></data-item>
             </router-link>
          </el-col>  
        <el-col :xs="12" :sm="12" :lg="6" class="card-panel-col">
             <router-link to="/admin/item">
                 <data-item svg="hbje" bgcolor="#34bfa3" color="#fff" server="report/redpkgamt" title="累计发放红包金额" ></data-item>
             </router-link>
         </el-col> 
          <el-col :xs="12" :sm="12" :lg="6" class="card-panel-col">
             <router-link to="/admin/appraise">
          <data-item svg="pinggu" bgcolor="#f4516c" color="#fff" server="report/appraisenum"  title="新增评估(份)" ></data-item>
             </router-link>
          </el-col>  
        <el-col :xs="12" :sm="12" :lg="6" class="card-panel-col">
             <router-link to="/admin/appraise">
                 <data-item svg="pinggu" bgcolor="#34bfa3" color="#fff" server="report/appraisenum"  :cond="cond" title="累计评估(份)" ></data-item>
             </router-link>
         </el-col> 
          <el-col :xs="12" :sm="12" :lg="6" class="card-panel-col">
             <router-link to="/admin/routine">
          <data-item svg="jiance" bgcolor="#f4516c" color="#fff" server="report/detectnum" :cond="cond"  title="新增检测(份)" ></data-item>
             </router-link>
          </el-col>  
        <el-col :xs="12" :sm="12" :lg="6" class="card-panel-col">
             <router-link to="/admin/forms">
                 <data-item svg="jiance" bgcolor="#34bfa3" color="#fff" server="report/detectnum" title="累计检测(份)" ></data-item>
             </router-link>
         </el-col> 
      </el-row>



       <el-row :gutter="40" class="card-panel-row">
         <el-col :xs="48" :sm="48" :lg="24" class="card-panel-col">
         <el-card class="box-card">
            <!--企业增长曲线-->
            <line-chart  server="report/curve"></line-chart>
         </el-card>
         </el-col>
      </el-row>
      <el-row :gutter="40" v-if="false">
        <el-col :xs="24" :sm="24" :lg="12"  >
          <el-card class="box-card">
            <div slot="header" class="clearfix">
              <span>新用户</span>
              <router-link :to="`/admin/corp`">
              <el-button style="float: right; padding: 3px 0" type="text">累计用户</el-button>
              </router-link>
            </div>
           <el-table  :data="corpList" :stripe="true" >
      <el-table-column align="header-center" fixed="left" label="#ID">
        <template slot-scope="scope">
          {{ scope.row.id }}

        </template>
      </el-table-column>
      <el-table-column align="header-center" fixed="left" label="企业名称" >
        <template slot-scope="scope">
          {{ scope.row.name }}

        </template>
      </el-table-column>

      <el-table-column align="header-center" fixed="left" label="客服数目" >
        <template slot-scope="scope">
          {{ scope.row.numkefu }}
        </template>
      </el-table-column>
     
      <el-table-column align="header-center" fixed="left" label="过期时间" >
        <template slot-scope="scope">
          {{ scope.row.expireat }}
        </template>
      </el-table-column>
       

       <el-table-column align="header-center" fixed="left" label="注册时间">
        <template slot-scope="scope">
          {{ scope.row.createat }}
        </template>
      </el-table-column>
     

    </el-table>

          </el-card>
        </el-col>


       <el-col :xs="24" :sm="24" :lg="12"  >
          <el-card class="box-card">
            <div slot="header" class="clearfix">
              <span>最新红包</span>
              <router-link :to="`/admin/charge`">
              <el-button style="float: right; padding: 3px 0" type="text">红包列表</el-button>
              </router-link>
            </div>
           <el-table  :data="chargeList" :stripe="true" >
      <el-table-column align="header-center" fixed="left" label="#ID">
        <template slot-scope="scope">
          {{ scope.row.id }}

        </template>
      </el-table-column>
      <el-table-column align="header-center" fixed="left" label="产品名称" >
        <template slot-scope="scope">
          {{ scope.row.title }}

        </template>
      </el-table-column>

      <el-table-column align="header-center" fixed="left" label="原价" >
        <template slot-scope="scope">
          {{ scope.row.price }}
        </template>
      </el-table-column>
     
      <el-table-column align="header-center" fixed="left" label="购买天数" >
        <template slot-scope="scope">
          {{ scope.row.num }}
        </template>
      </el-table-column>
       

       <el-table-column align="header-center" fixed="left" label="付款金额">
        <template slot-scope="scope">
          {{ scope.row.cost }}
        </template>
      </el-table-column>
      <el-table-column align="header-center" fixed="left" label="付款金额">
        <template slot-scope="scope">
          {{ scope.row.stat==2?"已付":(scope.row.stat==1?"待付":"无效") }}
        </template>
      </el-table-column>
      <el-table-column align="header-center" fixed="left" label="付款时间">
        <template slot-scope="scope">
          {{ scope.row.createat }}
        </template>
      </el-table-column>

    </el-table>

          </el-card>
        </el-col>

      
      
      </el-row>
   

 


  </div>
</template>
<style scoped>
.card-panel-col{
  margin-top: 5px;
  margin-bottom: 5px;
}
</style>
<script>




import {search as corps,charges} from "@/api/account"
import {dateFormat,pickerOptions } from "@/utils"

const stats=[
  {"name":1,"label":"待指派"},
  {"name":2,"label":"处理中"},
  {"name":3,"label":"已完成"},
  {"name":99,"label":"订单已删除"},
]
export default {
  name: 'DashboardEditor',
  data() {
    
    return {

      daterange:[new Date(dateFormat(new Date(),"yyyy-MM-dd 00:00:00")),new Date(dateFormat(new Date(),"yyyy-MM-dd 23:59:59"))],
     
      corpList:[],
      chargeList:[],
      pickerOptions:pickerOptions,
      cond:{
        datefrom:dateFormat(new Date(),"yyyy-MM-dd 00:00:00"),
        dateto:dateFormat(new Date(),"yyyy-MM-dd 23:59:59")
      },
      condamtxz:{stat:2,datefrom:dateFormat(new Date(),"yyyy-MM-dd 00:00:00"), dateto:dateFormat(new Date(),"yyyy-MM-dd 23:59:59")},
      condamttotal:{stat:2},
       sortArg:{
        order:"desc",
        prop:"numsaled"
      },
    }
  },
 created() {
   
 },
 methods:{
  
   sortChange({ prop, order}){
      
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

   
   resetcond(){
      let cond = {}
      if(this.daterange.length==2){
          cond.datefrom = dateFormat(this.daterange[0],"yyyy-MM-dd 00:00:00")
          cond.dateto = dateFormat(this.daterange[1],"yyyy-MM-dd 00:00:00")
          this.cond =Object.assign({},cond)
      }

   }
 }

}
</script>
