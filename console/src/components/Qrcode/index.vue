<template>
    <vue-qr style="text-align:center" :logoSrc="config.imagePath" :text="config.value" :size="size" :margin="margin"></vue-qr>
</template>

<script>
import VueQr from 'vue-qr'

export default {
  name: 'qrcode',
  props:{
        url:{
            require:true
        },logo:{
            default:process.env.VUE_APP_LOGO
        },size:{
            default:200
        },margin:{
            default:0
        }
  },
  components: {

     VueQr
 },
    computed:{
          config(){
              return {
                  value:this.url,
                  imagePath:this.logo
              }
      	  }
    }
}
</script>

<style scoped>

</style>
