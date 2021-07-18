<template>
  <div class="app-container">
    <el-container>

        <el-main>
          <el-header>
                  <el-form :inline="true" v-model="cond"  class="demo-form-inline">
                <el-form-item label="创建时间">
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
                    <el-input v-model="cond.kword" placeholder="请输名称"></el-input>
                  </el-form-item>

                  <el-form-item>
                    <el-button type="primary" @click="search" >搜索</el-button>
                     <el-button type="primary" @click="handleAddOne"  v-if="role==3" >添加店员账号</el-button>
                  </el-form-item>

                </el-form>
          </el-header>



 <el-main>

    <el-table height="640" :data="dataList"  border>
      <el-table-column align="header-center" label="#ID" width="100">
        <template slot-scope="scope">
          {{ scope.row.id }}
        </template>
      </el-table-column>

      <el-table-column align="header-center" label="头像">
        <template slot-scope="scope">
             <el-image style="width: 64px; height: 64px" :src="scope.row.headImage" fit="fit"></el-image>
        </template>
      </el-table-column>
      <el-table-column align="header-center" label="姓名" >
        <template slot-scope="scope">
          {{ scope.row.nickName }}
        </template>
      </el-table-column>

       <el-table-column align="header-center" label="手机号">
        <template slot-scope="scope">
          {{ scope.row.mobile }}
        </template>
      </el-table-column>


 
<el-table-column align="header-center" label="状态" width="220">
        <template slot-scope="scope">
          {{ scope.row.deleted===0?"在用":"停用" }}

        </template>
      </el-table-column>
      <el-table-column align="header-center" label="类型">
        <template slot-scope="scope">
          <state group="role" :code="scope.row.role"></state>
        </template>
      </el-table-column>
      <el-table-column align="center" label="操作">
        <template slot-scope="scope">
        <el-popconfirm title="确定要停用该账户吗?"  >   </el-popconfirm>
            
          
          <el-button-group>
          <el-button type="primary" v-if="false" size="small" @click="editpwd(scope)">重置密码</el-button>
              <el-button type="danger" size="small" v-if="scope.row.deleted==0" @click="disable(scope)" >停用账号</el-button>
          <el-button type="danger" size="small" v-if="scope.row.deleted===1" @click="enable(scope)">重新启用</el-button>
          </el-button-group>
        </template>
      </el-table-column>
    </el-table>

</el-main>

 <el-footer>
     <el-pagination
     background
      @size-change="handleSizeChange"
      @current-change="handleCurrentChange"
      :current-page="cond.pagefrom"
      :page-sizes="[20,100, 200, 300, 400]"
      :page-size="cond.pagesize"
      layout="total, sizes, prev, pager, next"
      :total="cond.total">
    </el-pagination>
   </el-footer>
  </el-main>
  </el-container>

<el-dialog
    title="添加店员账号"
    :visible.sync="dialogVisible"
    
  >
         <el-form  label-position="top">
         <el-form-item label="手机号">
           <el-input v-model="dataobject.mobile" placeholder="手机号码"></el-input>
         </el-form-item>
         <el-form-item label="密码">
          <el-input v-model="dataobject.passwd" placeholder="至少长6位"></el-input>
         </el-form-item>
         </el-form>
       
        <div slot="footer" style="text-align:right;">
         
        <el-button type="danger" @click="confirmData">提交</el-button>
      </div>
      

      
     
    </el-dialog>

  </div>
</template>

<script>
import {deepClone,dateFormat,pickerOptions} from "@/utils/index"
import { search ,update,create,disable,enable,deleteIt,editPwd} from '@/api/user'


const defaultObject = {
  id:0,
  name: '',
  logo:'',
  memo:"",
  createAt:""
}

export default {
  props:{
      role:{
        type:[Number,String],
        require:true
      }
  },
  data() {
    return {
       
      dataList: [],
      dialogVisible: false,
      dataobject:{},
      daterange:[],
     pickerOptions:Object.assign({},pickerOptions),
      cond:{
        kword:"",
        pagefrom:1,
        pagesize:20,
        total:0
      },
    }
  },
  created() {
    this.search()
  },
  methods: {
    
    search(){
        let tmp = Object.assign({},this.cond)
        tmp.role = this.role
            if(this.daterange.length>0){
                tmp.datefrom = dateFormat(this.daterange[0],"yyyy-MM-dd 00:00:00")
                tmp.dateto = dateFormat(this.daterange[1],"yyyy-MM-dd 23:59:59")
            }
        search(tmp).then(res=>{
          this.dataList = res.rows
          this.cond.total = res.total
        })
    },
    handleSizeChange(n){
        this.cond.pagesize = n
        this.search()
    },
    handleCurrentChange(p){
        this.cond.pagefrom = p
        this.search()
    },

    handleAddOne() {
      this.dataobject = Object.assign({}, defaultObject)
      this.dialogType = 'new'
      this.dialogVisible = true
    },
    handleEdit(scope) {
      this.dialogType = 'edit'
      this.dialogVisible = true
      this.dataobject = deepClone(scope.row)

    },
    handleDelete({ $index, row }) {
      this.$confirm('确定要伤处该配置吗?', '警告', {
        confirmButtonText: '确定删除',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(async() => {
          await deleteIt(row.id)
          this.dataList.splice($index, 1)
          this.$message({
            type: 'success',
            message: '你已成功删除数据'
          })
        })
        .catch(err => { console.error(err) })
    },
    enable(scope){
        enable({id:scope.row.id}).then(res=>{
          scope.row.deleted=0
          this.$ok(res.msg)
        }) .catch(err => { this.$error(err) })
    },
    disable(scope){
        disable({id:scope.row.id}).then(res=>{
          scope.row.deleted=1
          this.$ok(res.msg)
        }).catch(err => { this.$error(err) })
    },
    async confirmData() {
      const isEdit = this.dialogType === 'edit'
      if (isEdit) {
        update(this.dataobject).then(res=>{
          for (let index = 0; index < this.dataList.length; index++) {
          if (this.dataList[index].id === this.dataobject.id) {
            this.dataList.splice(index, 1, Object.assign({}, this.dataobject))
            break
          }
        }
        this.$message({
            type: 'success',
            message: '你已成功修改'
          })
          this.dialogVisible = false
        }).catch(e=>{
            this.$notify({
              title: '提示',
              message: e.message,
              type: 'error'
            })
             this.dialogVisible = false
        })

      } else {

        create(Object.assign({},this.dataobject,{role:this.role})).then(res=>{
          this.dataobject.id = res.data.id
        this.dataList.push(this.dataobject)
        this.dialogVisible = false
        }).catch(e=>{
            this.$notify({
              title: '提示',
              message: e.message,
              type: 'error'
            })
             this.dialogVisible = false
        })

      }


    },
     editpwd(scope){
        this.$prompt('请输入新的密码', '提示', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            inputErrorMessage: '输入不能为空',
             inputPlaceholder:'输入内容',
            inputValidator: (value)=>{
              if(value.trim().length < 1){
                    return '输入不能为空'
                }
            },
            beforeClose: (action, instance, done)=>{
                if(action == 'confirm'){
                    let value = instance.inputValue
                    if(value && value != ''){
                      let data = {}
                      data.id = scope.row.id
                      data.passwd = value
                      editPwd(data).then(res=>{
                         done()
                        this.$ok("密码修改成功")
                      }).catch(e=>{
                        console.log(e)
                          this.$error(e)
                      })
                    }else{

                    }
                }else{
                  done()
                }
              }
      })
    },
  }
}
</script>

<style lang="scss" scoped>

</style>
