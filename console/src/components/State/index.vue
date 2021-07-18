<template>
  <div>
      <el-tag v-text="labelstate"></el-tag>
  </div>
</template>

<script>
import dictApi from "@/api/dict"

export default {
  props: {
     group:{
         type:[String],require:false
     },
     code:{
         type:[String,Number],require:true
     }
  },
  data() {
    return {
        labelstate:""
    }
  },
  created(){
      this.initstate(this.group)
  },
  watch:{
      code(){
          this.initstate(this.group)
      }
  },
  methods: {
      initstate(group){
            dictApi.getOne(group).then(res=>{
                        let arr = JSON.parse(res.value)
                        arr.forEach(element => {
                                if(element.value==this.code){
                                    this.labelstate = element.label
                                }
                        }); 
            })
      }
  }
}
</script>

<style scoped>

</style>
