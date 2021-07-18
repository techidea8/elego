<template>
  <div >
    <el-header >
        <el-form :inline="true">
        <el-form-item>
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
          <el-radio-group v-model="datatype">
            <el-radio-button  label="d">日曲线</el-radio-button>
            <el-radio-button disabled   label="m">月曲线</el-radio-button>

          </el-radio-group>
        </el-form-item>
        </el-form>
   </el-header> 
  <div :class="className" ref="chartdom" :style="{height:height,width:width}"></div>
  </div>
</template>

<script>
import echarts from 'echarts'
require('echarts/theme/macarons') // echarts theme
import resize from './mixins/resize'
import {pickerOptions,dateFormat} from "@/utils/index"
import request from '@/utils/request'
//默认的配置
const defaultoptions ={
        xAxis: {
          data: [],
          boundaryGap: false,
          axisTick: {
            show: false
          }
        },
        grid: {
          left: 10,
          right: 10,
          bottom: 20,
          top: 30,
          containLabel: true
        },
        tooltip: {
          trigger: 'axis',
          axisPointer: {
            type: 'cross'
          },
          padding: [5, 10]
        },
        yAxis: {axisTick: {show: false}},
        legend: {data: []},series: []
      };
export default {
  mixins: [resize],
  props: {
    className: {
      type: String,
      default: 'chart'
    },
    width: {
      type: String,
      default: '100%'
    },
    height: {
      type: String,
      default: '350px'
    },
    autoResize: {
      type: Boolean,
      default: true
    },
    server:{
        type:String,
        required:true
    },
    cond:{
        type:Object,
        defult(){
            return {
              
            }
        }
    }
  },
  data() {
    return {
      daterange:[],
      pickerOptions:Object.assign({},pickerOptions),
      chart: null,
      datatype:"d",
    }
  },
  computed: {
    condition(){
        let tmp = {}
        if(!!this.daterange && this.daterange.length>0){
            tmp.datefrom = dateFormat(this.daterange[0],"yyyy-MM-dd 00:00:01")
            tmp.dateto = dateFormat(this.daterange[1],"yyyy-MM-dd 23:59:59")
        }  
        tmp.datatype = this.datatype
        tmp.asc = "id"
        return Object.assign({},tmp,this.cond)
    }
  },
  watch: {
    condition: {
      deep: true,
      handler(val) {
        if(!val){
          return 
        }
        this.initOption()
      }
    },
    server: {
      handler(val) {
        this.initOption()
      }
    },
  },
  mounted() {
    this.$nextTick(() => {
      this.initChart()
    })
  },
  beforeDestroy() {
    if (!this.chart) {
      return
    }
    this.chart.dispose()
    this.chart = null
  },
  methods: {
    color(){
             return "hsb(" + Math.random()  + ", 1, 1)";
    },
    buildOption(xdata,ydata){
            var opts = Object.assign({},defaultoptions)
            opts.xAxis.data = xdata
            opts.legend.data=[]
            opts.series=[]
            for(var i in ydata){
                let y = ydata[i]
                if(!y.title){
                    y.title = "序列" + i
                }
                if(!y.colora){
                    y.colora = this.color()
                }
                if(!y.colorl){
                    y.colora = this.color()
                }
                opts.legend.data.push(y.title)
                var serie = {
                        name: y.title, 
                        itemStyle: {
                            normal: {
                            color:  y.colorl,
                            lineStyle: {
                                color:y.colorl,
                                width:y.widthl || 2
                            },
                            /*areaStyle: {
                                    color:y.colora
                                }*/
                            }
                        },
                        smooth: true,
                        type: 'line',
                        data: y.data,
                        animationDuration: 2800,
                        animationEasing: 'cubicInOut'
                 }
                opts.series.push(serie)        
            }
          return opts
    },
    initOption(){
        
         if(!this.server){
            return 
        }
       
        let cond = Object.assign({},this.cond||{})
        request({
            url:this.server,
            method:"POST",
            data:this.condition
        }).then(res=>{
            let xdata = res.data.xdata
            let ydata = res.data.ydata
            this.chart.clear()
            if (xdata.length==0){
              return 
            }
            let opt =    this.buildOption(xdata,ydata)
            
            this.chart.setOption(opt)
            
        }).catch(err=>{
            return 
            console.error("测试数据让大家跟有成就感",err.message)
            let xdata=["A1","A2","A3"]
            let ydata=[{"data":[13,25,13],"title":"测试数据1"},{"data":[23,12,15],"title":"测试数据2"}]
            let opt =    this.buildOption(xdata,ydata)
            this.chart.setOption(opt)
        })
        
        
    }  ,
    initChart() {
      this.chart = echarts.init(this.$refs.chartdom, 'macarons')
      this.initOption()  
    }
  }
}
</script>
