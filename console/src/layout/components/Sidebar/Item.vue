<script>
export default {
  name: 'MenuItem',
  functional: true,
  props: {
    icon: {
      type: String,
      default: ''
    },
    title: {
      type: String,
      default: ''
    },
    svg:{
      type:Boolean,
      default:false
    }
  },
  render(h, context) {
    const { icon, title ,svg} = context.props
    const vnodes = []
     
    if (icon) {
       if(svg){
       vnodes.push(<svg-icon icon-class={icon} />)  
       }else{
         
         vnodes.push(<i  class={'el-icon ' + icon} />)
       }
       
    }
     

    if (title) {
      vnodes.push(<span slot='title'>{(title)}</span>)
    }
    return vnodes
  }
}
</script>
