<template>
  <div id="userLayout" class="user-layout-wrapper">
    <div class="container">
      <div class="top">
        <div class="desc">
          <img class="logo_login" src="@/assets/logo.png" alt="">
        </div>
        <div class="header">
          <a href="/">
            <!-- <img src="~@/assets/logo.png" class="logo" alt="logo" /> -->
            <span class="title">图灵应用开发平台</span>
          </a>
        </div>
      </div>
      <div class="main">
        <p>授权中...</p>
        <svg-icon icon-class="loading" />
      </div>

      <div class="footer">

        <div v-if="false" class="copyright">Copyright &copy; {{ curYear }} techidea8.com</div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Auth',
  data() {
    return {
      curYear: 0,
      lock: 'lock'
    }
  },
  created() {
    this.auth()
    this.curYear = new Date().getFullYear()
  },
  methods: {
    auth() {
      const token = this.$route.query.token
      this.$store.dispatch('account/authwithtoken', token).then(res => {
        this.$router.push('/')
      }).catch(err => {
        this.$router.push('/')
      })
    }
  }
}
</script>

<style scoped lang="scss">
@import "@/styles/login.scss";
</style>
