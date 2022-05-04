<template>
    <div>
        <div class="crumbs">
            <el-breadcrumb separator="/">
                <el-breadcrumb-item>
                    <i class="el-icon-lx-cascades"></i> 日志列表
                </el-breadcrumb-item>
            </el-breadcrumb>
        </div>
        <div class="container">
            <div class="handle-box">
              <el-date-picker
                  v-model="value"
                  type="datetimerange"
                  :shortcuts="shortcuts"
                  range-separator="To"
                  value-format="YYYY-MM-DD HH:mm:ss"
                  start-placeholder="Start date"
                  end-placeholder="End date"
              />
              <el-button class="search" type="primary" icon="el-icon-search" @click="logSearch()">搜索</el-button>
            </div>
            <el-table :data="tableData" border class="table" ref="multipleTable" header-cell-class-name="table-header">
                <el-table-column prop="id" label="ID" width="55" align="center"></el-table-column>
                <el-table-column prop="CreatedAt" :formatter="formatter" label="操作时间" align="center"></el-table-column>
                <el-table-column label="PATH" align="center" >
                    <template #default="scope">{{ scope.row.path }}</template>
                </el-table-column>
                <el-table-column label="错误信息" align="center">
                    <template #default="scope">
                        {{ scope.row.error_message }}
                    </template>
                </el-table-column>
                <el-table-column prop="ip" label="地址"></el-table-column>
                <el-table-column label="状态" align="center">
                    <template #default="scope">
                        <el-tag :type="
                                scope.row.status === 200
                                    ? 'success'
                                    : scope.row.state === '失败'
                                    ? 'danger'
                                    : ''
                            ">{{ scope.row.status }}</el-tag>
                    </template>
                </el-table-column>
                <el-table-column label="操作" width="180" align="center">
                    <template #default="scope">
                        <el-button type="text" icon="el-icon-edit" @click="handleEdit(scope.$index, scope.row)">查看
                        </el-button>
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
import service from "../utils/request";

export default {
    name: "logs",
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
        service.post("logs/getLogsList", {"page": query.pageIndex, "pageSize": query.pageSize, "start_date": value.value !== null ? value.value[0]:"", "end_date":value.value !== null ?value.value[1]:''}).then(
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
