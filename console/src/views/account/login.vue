<template>
  <div id="userLayout" class="user-layout-wrapper">
    <div class="container">
      <div class="top">
        <div class="desc">
          <img class="logo_login" src="@/assets/logo.png" alt="" />
        </div>
        <div class="header">
          <a href="/">
            <!-- <img src="~@/assets/logo.png" class="logo" alt="logo" /> -->
            <span class="title">图灵应用开发平台</span>
          </a>
        </div>
      </div>
      <div class="main">
        <el-form
          :model="loginForm"
          :rules="rules"
          ref="loginForm"
          @keyup.enter.native="submitForm"
        >
          <el-form-item prop="mobile">
            <el-input placeholder="请输入用户名" v-model="loginForm.mobile">
              <i class="el-input__icon el-icon-user" slot="suffix"></i
            ></el-input>
          </el-form-item>
          <el-form-item prop="passwd">
            <el-input
              :type="lock === 'lock' ? 'password' : 'text'"
              placeholder="请输入密码"
              v-model="loginForm.passwd"
            >
              <i
                :class="'el-input__icon el-icon-' + lock"
                @click="changeLock"
                slot="suffix"
              ></i>
            </el-input>
          </el-form-item>
          <el-form-item prop="code">
            <el-input
              v-model="loginForm.code"
              name="logVerify"
              placeholder="请输入验证码"
              style="width: 60%"
            />
            <div class="vPic">
              <img
                v-if="picPath"
                :src="picPath"
                width="100%"
                height="100%"
                alt="请输入验证码"
                @click="loginVefify()"
              />
            </div>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="submitForm" style="width: 100%"
              >登 录</el-button
            >
          </el-form-item>
        </el-form>
      </div>

      <div class="footer">
        <div class="links" >
         <router-link to="register" >注册新账号</router-link>
         <router-link to="resetpwd" >重置密码</router-link>
        </div>
        <div v-if="false" class="copyright">Copyright &copy; {{ curYear }} techidea8.com</div>
      </div>
    </div>
  </div>
</template>

<script>
import { captcha } from "@/api/captcha";
export default {
  name: "Login",
  data() {
    const checkUsername = (rule, value, callback) => {
      if (!/^1[3456789]\d{9}$/.test(value)) {
        return callback(new Error("请输入正确的手机号"));
      } else {
        callback();
      }
    };
    const checkPassword = (rule, value, callback) => {
      if (value.length < 6 || value.length > 12) {
        return callback(new Error("请输入正确的密码"));
      } else {
        callback();
      }
    };
     const checkCode = (rule, value, callback) => {
      if (value.length !=6 ) {
        return callback(new Error("请输入正确的验证码"));
      } else {
        callback();
      }
    };
    return {
      curYear: 0,
      lock: "lock",
      loginForm: {
        mobile: "",
        passwd: "",
        code: "",
        captchaId: "",
      },
      rules: {
        mobile: [{ validator: checkUsername, trigger: "blur" }],
        passwd: [{ validator: checkPassword, trigger: "blur" }],
        code: [{ validator: checkCode, trigger: "blur" }],
      },
      logVerify: "",
      picPath: "",
    };
  },
  created() {
    this.loginVefify();
    this.curYear = new Date().getFullYear();
  },
  methods: {
    async login() {
      console.log(this.loginForm)
      return await this.$store.dispatch("account/login",this.loginForm).then(res=>{
         this.$router.push("/")
      })
    },
    async submitForm() {
      this.$refs.loginForm.validate(async (v) => {
        console.log("v",v)
        if (v) {
          const flag = await this.login();
          if (!flag) {
            this.loginVefify();
          }
        } else {
          this.$message({
            type: "error",
            message: "请正确填写登录信息",
            showClose: true,
          });
          this.loginVefify();
          return false;
        }
      });
    },
    changeLock() {
      this.lock === "lock" ? (this.lock = "unlock") : (this.lock = "lock");
    },
    loginVefify() {
      captcha({}).then((ele) => {
        this.picPath = ele.data.base64data;
        this.loginForm.captchaId = ele.data.captchaId;
      });
    },
  },
};
</script>

<style scoped lang="scss">
@import "@/styles/login.scss";
</style>
