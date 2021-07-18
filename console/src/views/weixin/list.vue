<template>
  <div class="app-container">
      <el-container>

      
   <el-row>
          <el-form :inline="true" v-model="cond"  class="demo-form-inline">
    <el-form-item label="关键字">
      <el-input type="text" placeholder="名称或内容" v-model="cond.kword"></el-input>
    </el-form-item>
    <el-form-item label="创建时间">
      <el-date-picker
                    v-model="cond.daterange"
                    type="datetimerange"
                    align="right"
                    start-placeholder="开始日期"
                    end-placeholder="结束日期"
                    :default-time="['12:00:00', '08:00:00']">
                  </el-date-picker>
    </el-form-item>
    <el-form-item>
    <el-button-group>
    <el-button type="primary" @click="search">搜索</el-button>
     <el-button type="primary" @click="dlgvisible=true">添加</el-button>
    </el-button-group>
    </el-form-item>
          </el-form>
    </el-row>
    
    <el-table  :data="dataList" :row-class-name="tableRowClassName" style="width: 100%;" border>
     
      <el-table-column align="header-center" label="ID" width="100">
        <template slot-scope="scope">
         
          {{ scope.row.id }} 
        </template>
      </el-table-column>
      <el-table-column align="header-center" label="名称">
        <template slot-scope="scope">
          {{ scope.row.name }}
        </template>
      </el-table-column> 
      <el-table-column align="header-center" label="APPID">
        <template slot-scope="scope">
            {{ scope.row.appId }}
        </template>
      </el-table-column> 
      <el-table-column align="header-center" label="密钥">
        <template slot-scope="scope">
            {{ scope.row.secret }}
        </template>
      </el-table-column> 


      <el-table-column align="header-center" label="接入时间">
        <template slot-scope="scope">
          {{ scope.row.create_at }}
        </template>
      </el-table-column>
      
     
      
      <el-table-column align="header-center" label="创建时间" v-if="false">
        <template slot-scope="scope">
         {{scope.row.create_at}}
        </template>
      </el-table-column>
     
     
      <el-table-column align="center" label="操作">
        <template slot-scope="scope">
         <el-button-group>
          <el-button type="primary" size="small" @click="handleEdit(scope)">修改</el-button>
         
          <el-popconfirm title="确定要删除该记录吗?" @onConfirm="handleDelete(scope)">
               <el-button slot="reference" type="danger" size="small" >删除</el-button>
          </el-popconfirm>
         </el-button-group>
        </template>
      </el-table-column>
    </el-table>
       
        
 <el-footer>
     <el-pagination
     background
      @size-change="handleSizeChange"
      @current-change="handleCurrentChange"
      :current-page="cond.pagefrom"
      :page-sizes="[5,8,10,20,100, 200, 300, 400]"
      :page-size="cond.pagesize"
      layout="total, sizes, prev, pager, next"
      :total="cond.total">
    </el-pagination>
   </el-footer>
 



  <el-dialog
  :title="dialogType==='edit'?'修改':'添加'"
  :visible.sync="dlgvisible"
  >
    
      <el-form :model="dataobject" label-width="80px" style="margin-left:20px;margin-right:20px" label-position="left">

        <el-form-item label="产品名称">
                  <el-input v-model="dataobject.name" placeholder="输入名称"></el-input>
        </el-form-item> 
        
       <el-form-item label="AppId">
                  <el-input v-model="dataobject.appId" placeholder="输入appid"></el-input>
        </el-form-item> 

         <el-form-item label="Secret">
                  <el-input v-model="dataobject.secret" placeholder="输入secret"></el-input>
        </el-form-item> 

         <el-form-item label="token">
                  <el-input v-model="dataobject.token" placeholder="输入token"></el-input>
        </el-form-item> 

      
       <el-form-item label="商户号">
                  <el-input v-model="dataobject.mchId" placeholder="输入mchId"></el-input>
        </el-form-item> 


       <el-form-item label="apikey">
                  <el-input v-model="dataobject.apikey" placeholder="输入apikey"></el-input>
        </el-form-item> 


         <el-form-item label="cert">
                  <el-input v-model="dataobject.cert" placeholder="输入cert路径"></el-input>
        </el-form-item>  

        
        <div style="text-align:right;">
        <el-button type="danger" @click="dlgvisible=false">取消</el-button>
        <el-button type="primary" @click="confirmData">确认并提交</el-button>
      </div>

      </el-form>
      
     
    </el-dialog>
    

      </el-container>
        </div>
</template>

<script>

import {deepClone} from "@/utils/index" 
import { search,update,create ,deleteIt} from '@/api/weixin'

const defaultObject = {
  id:0,
  name: "",
  appId:"",secret:"",
  token:"",
  apitoken:"",
  cert:"",
}

export default {
 
  data() {
    let tmp = []
    for(var i=1;i<10;i++){
      tmp.push({name:`theme${i}`,label:`${i}分屏`})
    }
    return {
      dlgvisible:false,
      dataobject: Object.assign({}, defaultObject),
      
      dataList: [],
      dialogVisible: false,
      dialogType: 'new',
      checkStrictly: false,
     
      cond:{
        pagefrom:1,
        pagesize:10,
        appId:"",
        total:-1,
        daterange:[],
        }
    }
  },
  created() {
    // Mock: get all routes and roles list from server
    this.search()
    
  },
  methods: {
    
    handleSizeChange(n){
        this.cond.pagesize = n
        this.search()
    },
    handleCurrentChange(p){
        this.cond.pagefrom = p
        this.search()
    },
    tableRowClassName({row, rowIndex}) {
        if (row.requireAuth) {
          return 'warning-row';
        }
        return '';
      },
    async search() {
      let datearg = {
        
      } 
      
      if(this.cond.daterange && this.cond.daterange.length>0){
        datearg.datefrom = this.cond.daterange[0].format("yyyy-MM-dd hh:mm:ss")
        datearg.dateto = this.cond.daterange[1].format("yyyy-MM-dd hh:mm:ss")
      }
      const res = await search(Object.assign(Object.assign({},this.cond,datearg)))
      
      this.dataList = res.rows
      this.cond.total = res.total
    },
    
    handleAddOne() {
      this.dataobject = Object.assign({}, defaultObject)
      this.dialogType = 'new'
      this.dlgvisible = true
    },
    handleEdit(scope) {
      this.dialogType = 'edit'
      this.dlgvisible = true
    
      this.dataobject = deepClone(scope.row)

      //this.dataobject.daterange = [new Date(scope.row.startat),new Date(scope.row.expireat)]
    },
    handleDelete({ $index, row }) {
         deleteIt(row.id).then(res=>{
           this.dataList.splice($index, 1)
          this.$message({
            type: 'success',
            message: '你已成功删除数据'
          })
         }).catch(err => { console.error(err) })
    },
    
    async confirmData() {
      const isEdit = this.dialogType === 'edit'
     
      this.dataobject.probability = JSON.stringify(this.dataobject.probabilityObj)
      if (isEdit) {
        update(Object.assign({},this.dataobject)).then(res=>{
            this.search()
        })
      } else {
        const { data } = await create(Object.assign({},this.dataobject))
        this.dataobject.id = data.id
        this.dataList.push(data)
      }
      this.dlgvisible = false
      this.$ok( "数据操作成功")
    }
  }
}
</script>

<style lang="scss" scoped>

</style>
