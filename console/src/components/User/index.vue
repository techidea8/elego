<template>
  <div :class="{row:flex=='row',col:flex=='col'}" v-if="id>0">
    
    <el-image :src="user.headImage">
      <div slot="error" class="image-slot">
        <i class="el-icon-user"></i>
      </div>
    </el-image>
    <label v-text="user.nickName"></label>
  </div>
</template>

<script>
import {snsinfo} from "@/api/user"

export default {
  props: {
     id:{
       Type:[String,Number],
       require:true,
     },
     flex:{
       Type:String,
       default:'row'
     }
  },
  data() {
    return {
        user:{
          headImage:"",
          nickName:""
        }
    }
  },
  created(){
      this.loadata(+this.id)
  },
  methods: {
    loadata(id){  
      if(!id){
        return 
      }
      snsinfo(id).then(res=>{
            //console.log("res",res)
            this.user.headImage = res.headImage
            this.user.nickName = res.nickName
      })
    }
  }
}
</script>

<style scoped>
.row{
  display: flex;
  flex-direction: row;
  justify-content: flex-start;
  align-items: center;
}
.col{
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  align-items: center;
}

</style>
