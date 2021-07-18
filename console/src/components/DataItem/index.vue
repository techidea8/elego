<template>
    <div class="card-panel" @mouseover="setstat(true)" @mouseout="setstat(false)">
        <div  class="card-panel-icon-wrapper" :style="styleobject">
          <svg-icon :icon-class="svg" class-name="card-panel-icon" />
        </div>
        <div class="card-panel-description">
          <div class="card-panel-text" v-text="title"></div>
          <count-to :start-val="0" :decimals="decimals" :end-val="endval" :duration="2600" class="card-panel-num" />
        </div>
    </div>
</template>

<script>
import CountTo from 'vue-count-to'
import request from '@/utils/request'
export default {
  name:"DataItem",
  props:{
      svg:{
          type:String,
          default:"card-panel-icon"
      },
      decimals:{
        type:Number,
        default:0
      },
      title:{
          type:String,
          default:"参数名称",
      },
      color:{
          type:String,
          default:"#40c9c6",
      },
      bgcolor:{
          type:String,
          default:"#fff",
      },
      count:{
          type:Number,
          default:0,
      },
      duration:{
          type:Number,
          default:2600,
      },
      server:{
          type:String,
          default:"",
      },
      param:{ //如果定义了param参数需要解析数据
          type:String,
          default:"data",
      },
      cond:{
          type:Object,
          default:function(){
              return {
                  "datefrom":"",
                  "dateto":"",
              }
          }
      },
      overClass:{
          type:Object,
            default:function(){
              return {}
          }
      }  
  },
  data(){
      return {
          condition:{},
          activated:false,
          styleobject:{},
          endval:0,
      }
  },
  
  created() {
      this.condition = Object.assign({},this.cond)
      this.endval = this.count
      this.loaddata()
  },
  watch: {
      "cond":{
          deep:true,
          handler(n,o){
                if(!n){
                    this.condition = {}
                }else{
                    this.condition =Object.assign({},n)
                }
                this.loaddata()
          }
      }
  },
  components: {
    CountTo
  },
  methods: {
      setstat(r){
          this.activated = r
          if(r){
              
              this.styleobject = Object.assign({},{color:this.color,background:this.bgcolor},this.overClass)
          }else{
              this.styleobject = {}
          }
      },
      loaddata(){
            if(!this.server){
                return 
            }
           
           request( {
                        url: this.server,
                        method: 'post',
                        data:this.condition
                    }).then(res=>{
                        //
                        var tmp = this.param.split(".")
                        let r = res
                        for(var i=0;i<tmp.length;i++){
                            r = r[tmp[i]]
                        }
                        this.endval = r 
                    })
                
      }
  }
}
</script>

<style lang="scss" scoped>


  .card-panel-col {
    margin-bottom: 32px;
  }

  .card-panel {
    height: 108px;
    cursor: pointer;
    font-size: 12px;
    position: relative;
    overflow: hidden;
    color: #666;
    background: #fff;
    box-shadow: 4px 4px 40px rgba(0, 0, 0, .05);
    border-color: rgba(0, 0, 0, .05);

    &:hover {
      .card-panel-icon-wrapper {
        color: #fff;
      }

      .icon-1 {
        background: #40c9c6;
      }

      .icon-message {
        background: #36a3f7;
      }

      .icon-money {
        background: #f4516c;
      }

      .icon-shopping {
        background: #34bfa3
      }
    }

    .icon-people {
      color: #40c9c6;
    }

    .icon-message {
      color: #36a3f7;
    }

    .icon-money {
      color: #f4516c;
    }

    .icon-shopping {
      color: #34bfa3
    }

    .card-panel-icon-wrapper {
      float: left;
      margin: 14px 0 0 14px;
      padding: 16px;
      transition: all 0.38s ease-out;
      border-radius: 6px;
    }

    .card-panel-icon {
      float: left;
      font-size: 48px;
    }

    .card-panel-description {
      float: right;
      font-weight: bold;
      margin: 26px;
      margin-left: 0px;

      .card-panel-text {
        line-height: 18px;
        color: rgba(0, 0, 0, 0.45);
        font-size: 16px;
        margin-bottom: 12px;
      }

      .card-panel-num {
        font-size: 20px;
      }
    }
  }


@media (max-width:550px) {
  .card-panel-description {
    display: none;
  }

  .card-panel-icon-wrapper {
    float: none !important;
    width: 100%;
    height: 100%;
    margin: 0 !important;

    .svg-icon {
      display: block;
      margin: 14px auto !important;
      float: none !important;
    }
  }
}
</style>