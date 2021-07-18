<template>
  <div class="dashboard-container">
    <user-list :role="3" />
  </div>
</template>

<script>
import userList from './list'
export default {
  name: 'Dashboard',
  components: { userList},
  data() {
    return {
     
    }
  },
  computed: {
    
  },
}
</script>
