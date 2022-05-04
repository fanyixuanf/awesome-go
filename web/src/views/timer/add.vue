<template>
  <div>
    <div class="crumbs">
      <el-breadcrumb separator="/">
        <el-breadcrumb-item>
          <i class="el-icon-lx-calendar"></i> 定时任务
        </el-breadcrumb-item>
        <el-breadcrumb-item>添加定时任务</el-breadcrumb-item>
      </el-breadcrumb>
    </div>
    <div class="container">
      <div class="form-box">
        <el-form ref="formRef" :rules="rules" :model="form" label-width="120px">
          <el-form-item label="任务名称" prop="name">
            <el-input v-model="form.name"></el-input>
          </el-form-item>
          <el-form-item label="Key" prop="key">
            <el-input v-model="form.key"></el-input>
          </el-form-item>
          <el-form-item label="Url" prop="url">
            <el-input v-model="form.url"></el-input>
          </el-form-item>
          <el-form-item label="Cron表达式" prop="cron">
            <el-input v-model="form.cron"></el-input>
            <i class="el-icon-lx-warnfill"></i>定时任务详细配置参考:https://pkg.go.dev/github.com/robfig/cron/v3
          </el-form-item>
          <el-form-item label="消息" prop="msg">
            <el-input type="textarea" rows="5" v-model="form.msg"></el-input>
            <i class="el-icon-lx-warnfill"></i>消息内容, json格式
          </el-form-item>
          <el-form-item label="描述" prop="desc">
            <el-input type="textarea" rows="5" v-model="form.desc"></el-input>
          </el-form-item>

          <el-form-item>
            <el-button type="primary" @click="onSubmit">提交</el-button>
            <el-button @click="onReset">重置表单</el-button>
          </el-form-item>
        </el-form>
      </div>
    </div>
  </div>
</template>

<script>
import {reactive, ref} from "vue";
import {useRouter} from "vue-router";
import {ElMessage} from "element-plus";
import service from "../../utils/request";


export default {
  name: "addTimer",
  setup() {
    const router = useRouter();
    const rules = {
      name: [
        { required: true, message: "请输入任务名称", trigger: "blur" },
      ],
      key: [
        { required: true, message: '请输入Key', trigger: 'blur' }
      ],
      url: [
        { required: true, message: '请输入url', trigger: 'blur' }
      ],
      cron: [
        { required: true, message: '请输入cron', trigger: 'blur' }
      ],
      msg: [
        { required: true, message: '请输入msg', trigger: 'blur' }
      ],
    };
    const formRef = ref(null);
    const form = reactive({
      name: "",
      key:'',
      url:'',
      cron:'',
      msg: '',
      desc: '',
    });

    // 提交
    const onSubmit = () => {
      // 表单校验
      formRef.value.validate((valid) => {
        if (valid) {
          service.post('cronjob/addCronJob',{"name":form.name, "job_key":form.key, "url":form.url, "cron_expression":form.cron,"msg":form.msg, "desc":form.desc}).then(
              function(res){
                if (res.code == 0) {
                  ElMessage.success("添加成功！");
                  onReset();
                } else {
                  ElMessage.error("添加失败！");
                }
              }
          )

        } else {
          return false;
        }
      });
    };
    // 重置
    const onReset = () => {
      formRef.value.resetFields();
    };

    return {
      rules,
      formRef,
      form,
      onSubmit,
      onReset,
    };
  },
};
</script>
<style scoped>
.el-icon-lx-warnfill {
  color: #f56c6c;
}
</style>