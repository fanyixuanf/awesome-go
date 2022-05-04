<template>
    <div>
        <el-row :gutter="20">
            <el-col :span="8">
                <el-card shadow="hover" class="mgb20" style="height:400px;">
                    <div class="user-info">
                        <img :src="user.headerImg" class="user-avator" alt />
                        <div class="user-info-cont">
                            <div class="user-info-name">{{ name }}</div>
                            <div>{{ role }}</div>
                            <div class="user-profile">{{ user.user_profile }}</div>
                        </div>
                    </div>
                    <div class="user-info-list">
                        客户端IP：
                        <span>{{ params.ip }}</span>
                    </div>
                    <div class="user-info-list">
                        登录时间：
                        <span>{{ params.login_date }}</span>
                    </div>
                    <div class="user-info-list">
                        登录地点：
                        <span>{{ params.address }}</span>
                    </div>
                    <div class="user-info-list">
                        系统信息：
                        <span>{{ params.goos }}</span>
                    </div>
                    <div class="user-info-list">
                        环境信息：
                        <span>{{ params.goVersion }}</span>
                    </div>
                    <div class="user-info-list">
                        经纬信息：
                        <span>{{ params.loc }}</span>
                    </div>
                    <div class="user-info-list">
                        时区信息：
                        <span>{{ params.timezone }}</span>
                    </div>
                </el-card>
            </el-col>
             <el-col :span="16">
                <el-row :gutter="20" class="mgb20">
                    <el-col :span="8">
                        <el-card shadow="hover" :body-style="{ margin: '10px' }">
                            <div class="grid-content grid-con-1">
                                <el-progress type="dashboard" :percentage="params.ramUsedPercent">
                                  <template #default="{ percentage }">
                                    <span class="percentage-value">{{ params.ramUsedPercent }}%</span>
                                    <span class="percentage-label">UsedPercent</span>
                                  </template>
                                </el-progress>
                                <div class="grid-cont-right">
                                    <div class="grid-num">{{ params.ramTotalMb }}</div>
                                    <div>RAM</div>
                                </div>
                            </div>
                        </el-card>
                    </el-col>
                    <el-col :span="8">
                        <el-card shadow="hover" :body-style="{ margin: '10px' }">
                            <div class="grid-content grid-con-2">
                                <el-progress type="dashboard" :percentage="params.diskUsedPercent">
                                  <template #default="{ percentage }">
                                    <span class="percentage-value">{{ params.diskUsedPercent }}%</span>
                                    <span class="percentage-label">UsedPercent</span>
                                  </template>
                                </el-progress>
                                <div class="grid-cont-right">
                                    <div class="grid-num">{{ params.diskTotalGb }}</div>
                                    <div>DISK</div>
                                </div>
                            </div>
                        </el-card>
                    </el-col>
                    <el-col :span="8">
                        <el-card shadow="hover" :body-style="{ margin: '10px' }">
                            <div class="grid-content grid-con-3">
                                <el-progress type="dashboard" :percentage="params.cpus">
                                  <template #default="{ percentage }">
                                    <span class="percentage-value">{{ params.cpus }}%</span>
                                    <span class="percentage-label">CPU</span>
                                  </template>
                                </el-progress>
                                <div class="grid-cont-right">
                                    <div class="grid-num">{{ params.numGoroutine }}</div>
                                    <div>Goroutine</div>
                                </div>
                            </div>
                        </el-card>
                    </el-col>
                </el-row>

            </el-col>
        </el-row>

    </div>
</template>

<script>
import Schart from "vue-schart";
import {computed, reactive} from "vue";
import {useStore} from "vuex";
import service from "../utils/request";
import {ElMessage} from "element-plus";


export default {
    name: "dashboard",
    components: { Schart },
    setup() {
        const name = sessionStorage.getItem("user");
        const role = sessionStorage.getItem("role");
        const params = reactive({
          login_date: "",
          ip: "",
          address: "",
          goos: "",
          ramUsedPercent: 0,
          ramTotalMb: 0,
          diskUsedPercent: 0,
          diskTotalGb: 0,
          numGoroutine: 0,
          cpus: 0,
          goVersion: "",
          timezone: "",
          loc: "",
        });
        const initData = function() {
          service.post('/dashboard/init').then(
              function(res){
                if (res.code === 0) {
                  params.login_date = res.data.login_date;
                  params.address = res.data.other.country + " · " +
                                   res.data.other.region + " · " +
                                   res.data.other.city;
                  params.ip = res.data.other.ip;
                  params.goos = res.data.server.os.goos;
                  params.ramUsedPercent = res.data.server.ram.usedPercent;
                  params.ramTotalMb = res.data.server.ram.totalMb;
                  params.diskUsedPercent = res.data.server.disk.usedPercent;
                  params.diskTotalGb = res.data.server.disk.totalGb;
                  params.numGoroutine = res.data.server.os.numGoroutine;
                  params.cpus = Math.floor(res.data.server.cpu.cpus[0]);
                  params.goVersion = res.data.server.os.goVersion;
                  params.loc = res.data.other.loc;
                  params.timezone = res.data.other.timezone;
                } else {
                  ElMessage.error("请求失败, 请联系管理员")
                }
              }
          ).catch(
              err => (console.log(err))
          )
        };
        const store = useStore();
        const user = computed(()=> store.state.user);
        return {
            name,
            role,
            initData,
            params,
            user,
        };
    },
  mounted() {
      this.initData()
  },
};
</script>

<style scoped>
.grid-content {
    display: flex;
    align-items: center;
    height: 100px;
}

.grid-cont-right {
    flex: 1;
    text-align: center;
    font-size: 14px;
    color: #999;
}

.grid-num {
    font-size: 30px;
    font-weight: bold;
}

.grid-con-icon {
    font-size: 50px;
    width: 100px;
    height: 100px;
    text-align: center;
    line-height: 100px;
    color: #fff;
}

.grid-con-1 .grid-con-icon {
    background: rgb(45, 140, 240);
}

.grid-con-1 .grid-num {
    color: rgb(45, 140, 240);
}

.grid-con-2 .grid-con-icon {
    background: rgb(100, 213, 114);
}

.grid-con-2 .grid-num {
    color: rgb(45, 140, 240);
}

.grid-con-3 .grid-con-icon {
    background: rgb(242, 94, 67);
}

.grid-con-3 .grid-num {
    color: rgb(242, 94, 67);
}

.user-info {
    display: flex;
    align-items: center;
    padding-bottom: 20px;
    border-bottom: 2px solid #ccc;
    margin-bottom: 20px;
}

.user-avator {
    width: 120px;
    height: 120px;
    border-radius: 50%;
}

.user-info-cont {
    padding-left: 50px;
    flex: 1;
    font-size: 14px;
    color: #999;
}

.user-info-cont div:first-child {
    font-size: 30px;
    color: #222;
}

 .user-profile {
  font-size: 20px;
  color: #409eff;
}

.user-info-list {
    font-size: 14px;
    color: #999;
    line-height: 25px;
}

.user-info-list span {
    margin-left: 70px;
}

.mgb20 {
    margin-bottom: 20px;
}

.percentage-value {
  display: block;
  margin-top: 10px;
  font-size: 28px;
}
.percentage-label {
  display: block;
  margin-top: 10px;
  font-size: 12px;
}
.demo-progress .el-progress--line {
  margin-bottom: 15px;
  width: 350px;
}
.demo-progress .el-progress--circle {
  margin-right: 15px;
}
</style>
