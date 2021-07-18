<template>
  <div>
    <div ref="editor" style="text-align:left;margin: 5px;width: 100%;height: 100%;z-index: 98;"/>
  </div>
</template>

<script>

import E from 'wangeditor'
import { getToken } from '@/utils/auth'
import {hostUrl,upload} from "@/api/attach"
export default {
  name: 'Editor',
  model: {
    prop: 'value',
    event: 'change'
  },
  props: {
    value: {
      type: String,
      default: ''
    }
  },
  data() {
    return {
      headers: {
       'X-Token': getToken()
      },
      info: null,
      editor: null
    }
  },
  watch: {
    value: function(val) {
      this.editor.txt.html(val)
      // this.editor.txt.html(val)
    }
  },
  mounted() {
    // console.log(222)
    console.log(this.value)
    this.editor = new E(this.$refs.editor)
    //this.editor.customConfig.uploadImgShowBase64 = true // 使用 base64 保存图片
    // 不可修改
    this.editor.customConfig.uploadImgHeaders = this.headers
    // 自定义文件名，不可修改，修改后会上传失败
    this.editor.customConfig.uploadFileName = 'file'
    this.editor.customConfig.uploadImgServer = hostUrl() // 上传图片到服务器

    this.editor.customConfig.customUploadImg =async function (files, insert) {

        var i=0;
        for(i=0;i<files.length;i++){
          let formdata = new FormData()
          formdata.append("file",files[i])
          let res =await upload(formdata);
          insert(res.data.url)
        }


    }
    this.editor.customConfig.onchange = (html) => {
      this.info = html
      this.$emit('change', this.info)
      this.$emit('input', this.info)
    }
    this.editor.create()
    this.editor.txt.html(this.value)
    console.log("this.$refs.editor.offsetHeight",this.$refs.editor.offsetHeight)
  }
}
</script>

<style lang="scss">
  .editor-content{
    padding-left: 5px;
  }
  .w-e-text{

       img{
         width: 100%;
       }

  }
  .w-e-menu{
    z-index: 80 !important;
  }
  .w-e-toolbar .w-e-droplist{
    z-index: 99;
  }
  .w-e-toolbar{
    z-index: 99 !important;
  }
  .w-e-text-container{
      height: 100% !important;
      z-index: 99 !important;
  }
  .w-e-toolbar{
    display: flex;
    flex-direction: row;
    flex-wrap: nowrap;
    justify-content: space-between;
    position: sticky;
   
    z-index: 999;
  }
</style>
