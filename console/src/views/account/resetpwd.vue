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
            <span class="title">图灵应用-重置密码</span>
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
          <el-form-item prop="username">
            <el-input placeholder="请输入手机号" v-model="loginForm.username">
              <i class="el-input__icon el-icon-user" slot="suffix"></i
            ></el-input>
          </el-form-item>
          <el-form-item style="position: relative">
            <el-input
              v-model="loginForm.captcha"
              name="logVerify"
              placeholder="请输入验证码"
              
            >
              <template slot="append">
                 <el-button type="primary"   @click="loginVefify()">获取验证码</el-button> 
              </template>
            </el-input>
            
              
          
          </el-form-item>
          <el-form-item prop="password">
            <el-input
              :type="lock === 'lock' ? 'password' : 'text'"
              placeholder="请输入密码"
              v-model="loginForm.password"
            >
              <i
                :class="'el-input__icon el-icon-' + lock"
                @click="changeLock"
                slot="suffix"
              ></i>
            </el-input>
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
         <router-link to="login" >登录账号</router-link>
         <router-link to="resetpwd" >注册账号</router-link>
        </div>
        <div v-if="false" class="copyright">Copyright &copy; {{ curYear }} techidea8.com</div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapActions } from "vuex";
import { captcha } from "@/api/user";
export default {
  name: "Login",
  data() {
    const checkUsername = (rule, value, callback) => {
      if (value.length < 5 || value.length > 12) {
        return callback(new Error("请输入正确的用户名"));
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
    return {
      curYear: 0,
      lock: "lock",
      loginForm: {
        username: "admin",
        password: "123456",
        captcha: "",
        captchaId: "",
      },
      rules: {
        username: [{ validator: checkUsername, trigger: "blur" }],
        password: [{ validator: checkPassword, trigger: "blur" }],
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
      return await this.LoginIn(this.loginForm);
    },
    async submitForm() {
      this.$refs.loginForm.validate(async (v) => {
        if (v) {
          const flag = await this.login();
          if (!flag) {
            this.loginVefify();
          }
        } else {
          this.$error("请正确填写登录信息");
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
        this.picPath = ele.data.picPath;
        this.loginForm.captchaId = ele.data.captchaId;
      });
    },
  },
};
</script>

<style scoped lang="scss">
@import "@/styles/login.scss";
</style>
