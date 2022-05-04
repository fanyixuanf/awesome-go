<template>
  <div>
    <div class="crumbs">
      <el-breadcrumb separator="/">
        <el-breadcrumb-item>
          <i class="el-icon-lx-calendar"></i> 定时任务
        </el-breadcrumb-item>
        <el-breadcrumb-item>定时任务列表</el-breadcrumb-item>
      </el-breadcrumb>
    </div>
    <div class="container">
<!--      <div class="handle-box">-->
<!--        <el-date-picker-->
<!--            v-model="value"-->
<!--            type="datetimerange"-->
<!--            :shortcuts="shortcuts"-->
<!--            range-separator="To"-->
<!--            value-format="YYYY-MM-DD HH:mm:ss"-->
<!--            start-placeholder="Start date"-->
<!--            end-placeholder="End date"-->
<!--        />-->
<!--        <el-button class="search" type="primary" icon="el-icon-search" @click="logSearch()">搜索</el-button>-->
<!--      </div>-->
      <el-table :data="tableData" border class="table" ref="multipleTable" header-cell-class-name="table-header">
        <el-table-column prop="ID" label="ID" width="55" align="center"></el-table-column>
        <el-table-column prop="CreatedAt" :formatter="formatter" label="添加时间" align="center"></el-table-column>
        <el-table-column label="Url" align="center" >
          <template #default="scope">{{ scope.row.url }}</template>
        </el-table-column>
        <el-table-column label="表达式" align="center">
          <template #default="scope">
            {{ scope.row.cron_expression }}
          </template>
        </el-table-column>
        <el-table-column prop="msg" label="消息"></el-table-column>
        <el-table-column prop="desc" label="描述"></el-table-column>
        <el-table-column prop="timer_id" label="任务ID"></el-table-column>
        <el-table-column label="操作" align="center">
          <template #default="scope">
            <el-button type="primary" icon="el-icon-search" @click="handleEdit(scope.$index, scope.row)">查看</el-button>
            <el-button type="success" v-if="scope.row.status === 0" icon="el-icon-lx-remind" @click="changeStatus(scope.row)">启动</el-button>
            <el-button type="danger" v-if="scope.row.status === 1" icon="el-icon-lx-remind" @click="changeStatus(scope.row)">暂停</el-button>
            <el-button type="danger"  icon="el-icon-delete" @click="handleEdit(scope.$index, scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="pagination">
        <el-pagination background layout="total, prev, pager, next" :current-page="query.pageIndex"
                       :page-size="query.pageSize" :total="pageTotal" @current-change="handlePageChange"></el-pagination>
      </div>
    </div>
    <!-- 编辑弹出框 -->
    <el-dialog title="查看" v-model="editVisible" width="30%">
      <el-form label-width="70px">
        <el-form-item label="PATH">
          {{ form.path }}
        </el-form-item>
        <el-form-item label="地址">
          {{ form.addr }}
        </el-form-item>
      </el-form>
      <template #footer>
                <span class="dialog-footer">
                    <el-button @click="editVisible = false">取 消</el-button>
                    <el-button type="primary" @click="editVisible = false">确 定</el-button>
                </span>
      </template>
    </el-dialog>
  </div>
</template>



<script>
import {reactive, ref} from "vue";
import {ElMessage, ElMessageBox} from "element-plus";
import service from "../../utils/request";

export default {
  name: "timerList",
  setup() {
    const query = reactive({
      pageIndex: 1,
      pageSize: 10,
    });
    const value = ref(null)
    const shortcuts = [
      {
        text: 'Last week',
        value: () => {
          const end = new Date()
          const start = new Date()
          start.setTime(start.getTime() - 3600 * 1000 * 24 * 7)
          return [start, end]
        },
      },
      {
        text: 'Last month',
        value: () => {
          const end = new Date()
          const start = new Date()
          start.setTime(start.getTime() - 3600 * 1000 * 24 * 30)
          return [start, end]
        },
      },
      {
        text: 'Last 3 months',
        value: () => {
          const end = new Date()
          const start = new Date()
          start.setTime(start.getTime() - 3600 * 1000 * 24 * 90)
          return [start, end]
        },
      },
    ]
    const tableData = ref([]);
    const pageTotal = ref(0);
    // 请求接口
    const logSearch = () => {
      service.post("cronjob/cronJobList", {"page": query.pageIndex, "pageSize": query.pageSize}).then(
          function(res) {
            tableData.value = res.data.list;
            pageTotal.value = res.data.total;
          }
      ).catch(function (err) {
        ElMessage.error("请求失败");
        return false;
      });
    }
    logSearch()
    // 分页导航
    const handlePageChange = (val) => {
      query.pageIndex = val;
      logSearch()
    };


    // 查看
    const editVisible = ref(false);
    const form = ref({
      path: "",
      addr: "",
    })
    const handleEdit = (index, row) => {
      form.value.path = row.path;
      form.value.addr = row.ip;
      editVisible.value = true;
    };
    const formatter = (row, column) => {
      let date = new Date(row.CreatedAt)
      let Y = date.getFullYear() + '-';
      let M = (date.getMonth()+1 < 10 ? '0'+(date.getMonth()+1) : date.getMonth()+1) + '-';
      let D = (date.getDate() < 10 ? '0' + date.getDate() : date.getDate()) + ' ';
      let h = (date.getHours() < 10 ? '0' + date.getHours() : date.getHours()) + ':';
      let m = (date.getMinutes() < 10 ? '0' + date.getMinutes() : date.getMinutes()) + ':';
      let s = date.getSeconds() < 10 ? '0' + date.getSeconds() : date.getSeconds();
      return Y+M+D+h+m+s;
    }
    const changeStatus = (row) => {
      let url = '';
      if (row.status == 0) {
        url = 'cronjob/startCronJob';
      } else {
        url = 'cronjob/stopCronJob'
      }
      service.post(url, {"id":row.ID}).then(
          function(res) {
            if (res.code == 0) {
              ElMessage.success("操作成功！");
            } else {
              ElMessage.error("操作失败！");
            }
            logSearch()
          }
      )
    }

    return {
      query,
      tableData,
      pageTotal,
      editVisible,
      value,
      handlePageChange,
      shortcuts,
      logSearch,
      handleEdit,
      form,
      formatter,
      changeStatus,
    };
  },
};
</script>

<style scoped>
.handle-box {
  margin-bottom: 20px;
}

.search {
  margin-left: 5px;
}

.table {
  width: 100%;
  font-size: 14px;
}

.table-td-thumb {
  display: block;
  margin: auto;
  width: 40px;
  height: 40px;
}
</style>
