<template>
    <div>
        <el-row :gutter="20">
            <el-col :span="12">
                <el-card shadow="hover">
                    <template #header>
                        <div class="clearfix">
                            <span>基础信息</span>
                        </div>
                    </template>
                    <div class="info">
                        <div class="info-image" @click="showDialog">
                            <img :src="user.headerImg" />
                            <span class="info-edit">
                                <i class="el-icon-lx-camerafill"></i>
                            </span>
                        </div>
                        <div class="info-name">{{ user.userName }}</div>
                        <div class="info-desc">{{ user.user_profile }}</div>
                    </div>
                </el-card>
            </el-col>
            <el-col :span="12">
                <el-card shadow="hover">
                    <template #header>
                        <div class="clearfix">
                            <span>账户编辑</span>
                        </div>
                    </template>
                    <el-form label-width="120px" :model="form" :rules="rules" ref="save">
                        <el-form-item label="用户名：" prop="username">{{ user.userName }}
                          <el-input style="display: none;" v-model="form.username"></el-input>
                        </el-form-item>
                        <el-form-item label="旧密码：" prop="old">
                            <el-input type="password"  v-model="form.old" show-password></el-input>
                        </el-form-item>
                        <el-form-item label="新密码：" prop="new">
                            <el-input type="password"  v-model="form.new" show-password></el-input>
                        </el-form-item>
                        <el-form-item prop="desc" label="个人简介：">
                            <el-input type="textarea" v-model="form.desc"></el-input>
                        </el-form-item>
                        <el-form-item>
                            <el-button type="primary" @click="onSubmit(user.headerImg)">保存</el-button>
                        </el-form-item>
                    </el-form>
                </el-card>
            </el-col>
        </el-row>
        <el-dialog title="裁剪图片" v-model="dialogVisible" width="600px">
            <vue-cropper ref="cropper" :src="imgSrc" :ready="cropImage" :zoom="cropImage" :cropmove="cropImage"
                style="width: 100%; height: 400px"></vue-cropper>

            <template #footer>
                <span class="dialog-footer">
                    <el-button class="crop-demo-btn" type="primary">选择图片
                        <input class="crop-input" type="file" name="image" accept="image/*" @change="setImage" />
                    </el-button>
                    <el-button type="primary" @click="saveAvatar">上传并保存</el-button>
                </span>
            </template>
        </el-dialog>
    </div>
</template>

<script>
import {computed, reactive, ref} from "vue";
import {useStore} from "vuex";
import VueCropper from "vue-cropperjs";
import "cropperjs/dist/cropper.css";
import avatar from "../assets/img/img.jpg";
import {ElMessage} from "element-plus";
import service from "../utils/request";
import {useRouter} from "vue-router";


export default {
    name: "user",
    components: {
        VueCropper,
    },
    setup() {
        const store =  useStore();
        const user = computed(() => store.state.user);
        const router = useRouter();
        const form = reactive({
            old: "",
            new: "",
            desc: "",
            username: "",
            headerImg: "",
        });

        const rules = {
          old: [
            {
              required: true,
              message: "请输入旧密码",
              trigger: "blur",
            },
          ],
          new: [
            { required: true, message: "请输入新密码", trigger: "blur" },
          ],
          desc: [
            { required: true, message: "请输入个人简介", trigger: "blur" },
          ],
        };

        const save = ref(null);
        const onSubmit = (a) => {
          save.value.validate((valid) => {
              if (valid) {
                service.post('user/setUserInfo',{"username":form.username, "password":form.old, "header_img": a, "user_profile": form.desc, "newPassword": form.new})
                    .then(function(res) {
                      if (res.code === -1 ){
                        ElMessage.error(res.msg);
                      } else {
                        ElMessage.success("修改成功");
                        sessionStorage.removeItem("Authorization");
                        sessionStorage.removeItem("user");
                        sessionStorage.removeItem("role");
                        sessionStorage.removeItem("userinfo");
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

        const avatarImg = ref(avatar);

        const imgSrc = ref("");
        const cropImg = ref("");
        const dialogVisible = ref(false);
        const cropper = ref(null);

        const showDialog = () => {
            dialogVisible.value = true;
            imgSrc.value = avatarImg.value;
        };

        const setImage = (e) => {
            const file = e.target.files[0];
            if (!file.type.includes("image/")) {
                return;
            }
            const reader = new FileReader();
            reader.onload = (event) => {
                dialogVisible.value = true;
                imgSrc.value = event.target.result;
                cropper.value && cropper.value.replace(event.target.result);
            };
            reader.readAsDataURL(file);
        };

        const cropImage = () => {
            cropImg.value = cropper.value.getCroppedCanvas().toDataURL();
        };

        const saveAvatar = () => {
            avatarImg.value = cropImg.value;
            dialogVisible.value = false;
        };

        return {
            name,
            form,

            cropper,
            avatarImg,
            imgSrc,
            cropImg,
            showDialog,
            dialogVisible,
            setImage,
            cropImage,
            saveAvatar,
            user,
            save,
            rules,
            onSubmit,
            router,
        };
    },
  mounted() {

  }
};
</script>

<style scoped>
.info {
    text-align: center;
    padding: 35px 0;
}
.info-image {
    position: relative;
    margin: auto;
    width: 100px;
    height: 100px;
    background: #f8f8f8;
    border: 1px solid #eee;
    border-radius: 50px;
    overflow: hidden;
}
.info-image img {
    width: 100%;
    height: 100%;
}
.info-edit {
    display: flex;
    justify-content: center;
    align-items: center;
    position: absolute;
    left: 0;
    top: 0;
    width: 100%;
    height: 100%;
    background: rgba(0, 0, 0, 0.5);
    opacity: 0;
    transition: opacity 0.3s ease;
}
.info-edit i {
    color: #eee;
    font-size: 25px;
}
.info-image:hover .info-edit {
    opacity: 1;
}
.info-name {
    margin: 15px 0 10px;
    font-size: 24px;
    font-weight: 500;
    color: #262626;
}
.crop-demo-btn {
    position: relative;
}
.crop-input {
    position: absolute;
    width: 100px;
    height: 40px;
    left: 0;
    top: 0;
    opacity: 0;
    cursor: pointer;
}
</style>