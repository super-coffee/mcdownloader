<template>
        <el-main>
          <template>
            <el-table
                    :data="list.slice((currentPage-1)*pageSize,currentPage*pageSize)"
                    stripe
                    highlight-current-row
                    style="width: 100%">
              <el-table-column
                      prop="id"
                      label="#"
                      width="50">
              </el-table-column>
              <el-table-column
                      prop="versionId"
                      label="版本名称"
                      width="200">
              </el-table-column>

              <el-table-column
                      prop="releaseTime"
                      label="发布时间">
              </el-table-column>
              <el-table-column
                      label="操作">
                <el-button
                        size="mini"
                        @click="handleEdit(scope.$index, scope.row)">修复</el-button>
                <el-button
                        size="mini"
                        type="success"
                        plain
                        @click="handleDelete(scope.$index, scope.row)">下载</el-button>
              </el-table-column>
            </el-table>
            <el-pagination
                    @size-change="handleSizeChange"
                    @current-change="handleCurrentChange"
                    :current-page="currentPage"
                    :page-sizes="[10, 20, 30, 50, 100]"
                    :page-size="pageSize"
                    layout="total, sizes, prev, pager, next, jumper"
                    :total="list.length">
            </el-pagination>
          </template>
        </el-main>
</template>

<script>

  export default {
    name: 'app',
    mounted() {
      this.$message({
        message: '界面框架载入成功！恭喜你',
        type: "success",
      });
      this.getVersions();

    },
    data() {
      return {

        intelval: null,
        list: [],
        currentPage: 1,
        pageSize: 10,

      };
    },
    methods: {
      getVersions() {
        window.backend.lists().then(result => {
          this.list = result
        });
      },
      // 每页多少条
      handleSizeChange(val) {
        this.pageSize = val;
      },
      // 当前页
      handleCurrentChange(val) {
        this.currentPage = val;
      },
    }
  }
</script>
