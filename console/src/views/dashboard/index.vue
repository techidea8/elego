<template>
  <div class="dashboard-container">
    <component :is="currentView" v-if="currentView!='empty'" />
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import adminDashboard from './admin'
import merchantDashboard from './merchant'
import noauth from "../error-page/401"   //没有权限
export default {
  name: 'Dashboard',
  components: { adminDashboard, merchantDashboard,noauth},
  data() {
    return {
      rolelist:[],
    }
  },
  computed: {
    ...mapGetters([
      'roles'
    ]),
    currentView(){
      
      if(this.roles.indexOf(4)>-1){
        return "adminDashboard"
      }else{
        return "noauth"  //"merchantDashboard"
      }
      
    },
  },
  created() {
    
  },
  methods:{
    
  }
}
</script>
