
<template>
    <div class="app-container">
        <el-container> 
     <el-row>
            <el-form :inline="true" v-model="cond"  class="demo-form-inline">
      <el-form-item label="关键字">
        <el-input type="text" placeholder="名称或内容" v-model="cond.kword"></el-input>
      </el-form-item>
      
      <el-form-item>
      <el-button-group>
      <el-button type="primary" @click="search">搜索</el-button>
      <el-button type="primary" @click="handleAddOne">添加</el-button>
      </el-button-group>
      </el-form-item>
  
            </el-form>
      </el-row>
      
      <el-table  :data="dataList" style="width: 100%;" border>
        
        <el-table-column align="header-center" label="编号">
          <template slot-scope="scope" >
				      <p v-text="scope.row.id"></p>
          </template>
        </el-table-column>
       
        <el-table-column align="header-center" label="所属分类">
          <template slot-scope="scope" >
				      <p v-text="scope.row.name"></p>
          </template>
        </el-table-column>
       
        <el-table-column align="header-center" label="字段名称">
          <template slot-scope="scope" >
				        <p v-text="scope.row.label"></p>
               
          </template>
        </el-table-column>
       
        <el-table-column align="header-center" label="状态值">
          <template slot-scope="scope" >
              <el-badge :value="item.value" class="item" :key="index" v-for="(item,index) in scope.row.valueObj">
                <el-tag type="success" v-text="item.label"></el-tag>
            </el-badge>
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
    :visible.sync="dialogVisible"
    :direction="dialogType==='edit'?'rtl':'ltr'"
    ref="drawer"
    >
      
    <el-form :model="dataobject" label-width="80px" style="margin-left:20px;margin-right:20px" label-position="left">
          
                  
                    
          
                  
                        
                        <el-form-item label="名称">
                            <el-input v-model="dataobject.name" type="text" placeholder="输入"></el-input>
                        </el-form-item>
                        
                   
          
                  
                        
                            <el-form-item label="描述">
                                <el-input v-model="dataobject.label" type="text" placeholder="输入"></el-input>
                            </el-form-item>
                        
                   
          
                             <el-form-item :label="`状态值${index+1}`" v-for="(item,index) in dataobject.valueObj" :key="index">
                                <div style="display:flex;flex-direction:row;">
                                <el-input type="number" v-model="item.value" placeholder="状态值"></el-input>
                                <el-input type="text" v-model="item.label" placeholder="状态名"></el-input>
                                <el-button-group style="display:flex;flex-direct:row">
                                    <el-button type="danger" icon="el-icon-delete" size="mini" @click="removeitem(index)">删除</el-button>
                                    <el-button type="success" icon="el-icon-plus"  size="mini" @click="addnewitem(index+1)">添加</el-button>
                                </el-button-group>
                                </div>
                      </el-form-item> 
                   
          
         
          <div style="text-align:right;">
          <el-button type="danger" @click="dialogVisible=false">取消</el-button>
          <el-button type="primary" @click="confirmData">确认并提交</el-button>
        </div>
  
        </el-form>
        
       
    </el-dialog>
      
  
    </el-container>
  </div>
</template>
  
<script>
  
  import dictApi from '@/api/dict'

  const defaultObject = {  "id": 0 , "name": "" , "label": "" , "value": "" , valueObj:[{"value":1,"label":"正常"}] }
  
  export default {
   
    data() {
    
      return {
        dataobject: Object.assign({}, defaultObject),
        dataList: [],
        dialogVisible: false,
        dialogType: 'new',
        checkStrictly: false,
		chosedIds:[],
        cond:{
          pagefrom:1,
          pagesize:10,
          kword:"",
          total:-1,
          }
      }
    },
    created() {
      this.search()
    },
    methods: {
      removeitem(index){
          if(this.dataobject.valueObj.length<2){
            this.$error('至少要保留1个')
            return 
          }
          this.dataobject.valueObj.splice(index,1)
      },
      addnewitem(index){
            this.dataobject.valueObj.push({value:index+1,label:""})
      },
      datechange(a){ 
		  if(!a){
			this.dataobject.daterange = []
			this.cond.total = -1	
		  }else{
			this.dataobject.daterange = a
			this.cond.total = -1	
		  }
      },
      handleSizeChange(n){ 
          this.cond.pagesize = n
		  this.cond.total = -1
          this.search()
      },
      handleCurrentChange(p){
          this.cond.pagefrom = p
          this.search()
      },
      async search() { 
        const res = await dictApi.search(Object.assign(Object.assign({},this.cond)))
        res.rows.forEach(element => {
          element.valueObj = JSON.parse(element.value)
        });
        this.dataList = res.rows
        this.cond.total = res.total
      },
      
      handleAddOne() {
        this.dataobject = Object.assign({}, defaultObject)
        this.dialogType = 'new'
        this.dialogVisible = true
      },
      handleEdit(scope) {  
        this.dialogType = 'edit'
        this.dialogVisible = true
        this.dataobject = Object.assign({},scope.row)
      },
      handleDelete({ $index, row }) {  
          dictApi.deleteIt(row.id).then(res=>{
             this.dataList.splice($index, 1)
            this.$notify.success('你已成功删除数据')
           }).catch(err => { this.$notify.error(err.message) })
      },
      
      async confirmData() { 
        const isEdit = this.dialogType === 'edit'
        this.dataobject.value = JSON.stringify(this.dataobject.valueObj)
        if (isEdit) {
            dictApi.update(Object.assign({},this.dataobject)).then(res=>{
              this.search()
			        this.$notify.success(res.msg || "数据修改成功")
          })
        } else {
          const { data,msg } = await dictApi.create(Object.assign({},this.dataobject))
            this.dataobject.id = data.id
            this.dataList.push(Object.assign({},this.dataobject,data))
            this.$notify.success(msg || "数据添加成功")
        }
        this.dialogVisible = false
        
      }
    }
  }
  </script>
<style lang="scss" scoped>
.item {
  margin-top: 10px;
  margin-right: 40px;
}
  </style>
