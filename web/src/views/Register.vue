<template>
    <div class="login-wrap">
        <div class="ms-login">
            <div class="ms-title">Tools</div>
            <el-form :model="param" :rules="rules" ref="login" label-width="0px" class="ms-content">
                <el-form-item prop="username">
                    <el-input v-model="param.username" placeholder="username">
                        <template #prepend>
                            <el-button icon="el-icon-user"></el-button>
                        </template>
                    </el-input>
                </el-form-item>
                <el-form-item prop="password">
                    <el-input type="password" placeholder="password" v-model="param.password"
                        @keyup.enter="submitForm()">
                        <template #prepend>
                            <el-button icon="el-icon-lock"></el-button>
                        </template>
                    </el-input>
                </el-form-item>
                <el-form-item prop="captcha">
                  <el-input v-model="param.captcha" placeholder="captcha">
                    <template #prepend>
                      <el-button icon="el-icon-lx-qrcode"></el-button>
                    </template>
                  </el-input>
                </el-form-item>
                <el-form-item prop="captcha">
                  <img :src="param.picPath" @click="getCaptcha()">
                </el-form-item>
                <div class="login-btn">
                    <el-button type="primary" @click="submitForm()">注册</el-button>
                </div>
                <p class="login-tips">Tips : 注册</p>
            </el-form>
        </div>
    </div>
</template>

<script>
import {reactive, ref} from "vue";
import {useStore} from "vuex";
import {useRouter} from "vue-router";
import {ElMessage} from "element-plus";
import service from "../utils/request";

export default {
    setup() {
        const router = useRouter();
        const param = reactive({
            username: "",
            password: "",
            captchaId: "",
            picPath: "",
            captcha: "",
            authority_id: 1,
        });

        const rules = {
            username: [
                {
                    required: true,
                    message: "请输入用户名",
                    trigger: "blur",
                },
            ],
            password: [
                { required: true, message: "请输入密码", trigger: "blur" },
            ],
            captcha: [
                { required: true, message: "请输入验证码", trigger: "blur" },
            ],
        };
        const login = ref(null);
        const submitForm = () => {
            login.value.validate((valid) => {
                if (valid) {
                    service.post('base/register',{"username":param.username, "password":param.password,"captcha":param.captcha,"captchaId":param.captchaId}).then(
                        function(res) {
                          if (res.code === -1 ){
                            ElMessage.error(res.msg);
                          } else {
                            ElMessage.success("注册成功");
                            sessionStorage.setItem("Authorization", res.data.token);
                            sessionStorage.setItem("user", res.data.user.userName)
                            sessionStorage.setItem("role", res.data.user.nickName)
                            router.push("/login");
                          }
                        }).catch(function (error) {
                            console.log(error)
                        });
                } else {
                    ElMessage.error("非法参数");
                    return false;
                }
            });
        };

        const store = useStore();
        store.commit("clearTags");
        const getCaptcha = function() {
          service.post('/base/captcha').then(
              function(res) {
                if (res.code === 0){
                  param.captchaId = res.data.captchaId;
                  param.picPath = res.data.picPath;
                } else {
                  ElMessage.error("获取验证码失败");
                }
              }
          ).catch(error => (
              console.log(error)
          ))
        }
        return {
            param,
            rules,
            login,
            submitForm,
            getCaptcha,
        };
    },
    mounted() {
      this.getCaptcha();
    },
};
</script>

<style scoped>
.login-wrap {
    position: relative;
    width: 100%;
    height: 100%;
    background-image: url(../assets/img/login-bg.jpg);
    background-size: 100%;
}
.ms-title {
    width: 100%;
    line-height: 50px;
    text-align: center;
    font-size: 20px;
    color: #fff;
    border-bottom: 1px solid #ddd;
}
.ms-login {
    position: absolute;
    left: 50%;
    top: 50%;
    width: 350px;
    margin: -190px 0 0 -175px;
    border-radius: 5px;
    background: rgba(255, 255, 255, 0.3);
    overflow: hidden;
}
.ms-content {
    padding: 30px 30px;
}
.login-btn {
    text-align: center;
}
.login-btn button {
    width: 100%;
    height: 36px;
    margin-bottom: 10px;
}
.login-tips {
    font-size: 12px;
    line-height: 30px;
}
</style>