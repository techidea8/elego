<template>
  <div class="upload-container">
    <el-image  v-if="!!src" :src="src" :style="{width:size+'px',height:size+'px',borderradius:corner}"></el-image>
    <el-image v-if="!src" src="~@/assets/btnupload.png" :style="{width:size+'px',height:size+'px',borderradius:corner}"></el-image>
    <input v-if="upload" type="file"  ref="inputer" name="file" @change="uploadit" >
  </div>
</template>

<script>
import {upload as uploadApi} from "@/api/attach"

export default {
  name: 'upload',
  model:{
      prop:"src",
      event:"setvalue",
  },
  props: {
    corner:{
      type:[String],
      default:"0%",
    },
    size:{
      type:[Number,String],
      default:96
    },
    src: {
      type: String,
      default: ''
    },
    upload:{
      type:Boolean,
      default:true
    }
  },
  data() {
    return {
      tempUrl: '',
      dataObj: { token: '', key: '' }
    }
  },
  
  methods: {
    uploadit(){
            let inputDOM = this.$refs.inputer;
    　　　　this.file = inputDOM.files[0];// 通过DOM取文件数据
    　　　　let formData=new FormData();//new一个formData事件
    　　　　formData.append("file",this.file); //将file属性添加到formData里
    uploadApi(formData).then(res=>{
                   this.emitInput(res.data.url)
            })
    },
    rmImage() {
      this.emitInput('')
    },
    emitInput(val) {
      this.$emit('setvalue', val)
    },
    
  }
}
</script>

<style lang="scss" scoped>
    @import "~@/styles/mixin.scss";
    .upload-container {
        width: 100%;
        position: relative;
       input{
         position: absolute;
         left:0;
         right:0;
         top:0;
         bottom: 0;
         opacity:0
       }
        
    }

</style>
