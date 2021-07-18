<template>
  <div>
      <el-select v-model="chosecode">
          <el-option :label="item.label" :value="item.value" :key="index" v-for="(item,index) in optdata"></el-option>
      </el-select>
  </div>
</template>

<script>
import dictApi from "@/api/dict"

export default {
  model:{
      prop:"code",
      event:"setvalue",
  },
  props: {
     group:{
         type:[String],require:false
     },
     code:{
         type:[String,Number],require:true
     }
  },
  watch:{
       chosecode(n,v){
           this.$emit("setvalue",n)
       },
       code(){
        //   this.initstate(this.group)
       }
  },
  data() {
    return {
        chosecode:"",
        labelstate:"",
        optdata:[],
    }
  },
  created(){
      this.initstate(this.group)
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
                        this.optdata = arr
                        this.chosecode = this.code 
            })
      }
  }
}
</script>

<style scoped>

</style>
