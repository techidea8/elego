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
            <span class="title">图灵应用-注册账号</span>
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
            <el-input placeholder="请输入手机号" v-model="loginForm.mobile">
              <i class="el-input__icon el-icon-user" slot="suffix"></i
            ></el-input>
          </el-form-item>
          <el-form-item prop="code">
            <el-input
              v-model="loginForm.code"
              name="logVerify"
              placeholder="请输入验证码"
              
            >
              <template slot="append">
                 <el-button type="primary"   @click="sendsms()" v-text="wait"></el-button> 
              </template>
            </el-input>
            
              
          
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
          
          <el-form-item>
            <el-button type="primary" @click="submitForm" style="width: 100%"
              >注册</el-button
            >
          </el-form-item>
        </el-form>
      </div>

      <div class="footer">
        <div class="links" >
         <router-link to="login" >登录账号</router-link>
         <router-link to="resetpwd" >重置密码</router-link>
        </div>
        <div v-if="false" class="copyright">Copyright &copy; {{ curYear }} techidea8.com</div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapActions } from "vuex";
import { register } from "@/api/account";
import { sendsms } from "@/api/open";
import { setInterval, clearInterval } from 'timers';
export default {
  name: "Login",
  data() {
    const checkUsername = (rule, value, callback) => {
      if (/^1[3456789]d{9}$/.test(value)) {
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
      if (!/^\d{6}$/.test(value)) {
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
      },
      rules: {
        mobile: [{ validator: checkUsername, trigger: "blur" }],
        passwd: [{ validator: checkPassword, trigger: "blur" }],
        code: [{ validator: checkCode, trigger: "blur" }],
      },
      logVerify: "",
      picPath: "",
      wait:"获取验证码",
      timer:0,
      ttl:0,
    };
  },
  created() {
    this.curYear = new Date().getFullYear();
    
  },
  methods: {
   register() {
      this.$store.dispatch("account/register",this.loginForm).then(res=>{
          this.$success("注册成功")
          this.$router.push("/")
      })
    },
    async submitForm() {
      this.$refs.loginForm.validate(async (v) => {
        if (v) {
          this.register()
        } else {
          this.$error("请正确填写登录信息");
        
          return false;
        }
      });
    },
    changeLock() {
      this.lock === "lock" ? (this.lock = "unlock") : (this.lock = "lock");
    },
    updatettl(){
      if(this.ttl>0){
        this.ttl = this.ttl - 1
        this.wait = this.ttl + "秒后再试"
      }else{
        this.ttl = 0
        this.wait = "获取验证码"
        clearInterval(this.timer)
      }
    },
    sendsms() {
      if(this.ttl>0){
        return 
      }
      
      sendsms({mobile:this.loginForm.mobile}).then((res) => {
        this.ttl = 60
        this.timer = setInterval(this.updatettl,1000)
          this.$success(res.msg)
      });
    },
  },
};
</script>

<style scoped lang="scss">
@import "@/styles/login.scss";
</style>
