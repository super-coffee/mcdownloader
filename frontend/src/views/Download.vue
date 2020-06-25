<template>
  <el-main>
    <el-table
      stripe
      highlight-current-row
      style="width: 100%"
      :data="list.slice((currentPage-1)*pageSize, currentPage*pageSize)"
      :v-loading="isLoading"
    >
      <el-table-column prop="id" label="#" width="50"></el-table-column>
      <el-table-column prop="versionId" label="版本名称" width="200"></el-table-column>
      <el-table-column prop="releaseTime" label="发布时间"></el-table-column>
      <el-table-column label="操作">
        <el-button
          size="mini"
          icon="el-icon-magic-stick"
          @click="handleEdit(scope.$index, scope.row)"
        >修复</el-button>
        <el-button
          size="mini"
          type="primary"
          icon="el-icon-download"
          plain
          @click="handleDelete(scope.$index, scope.row)"
        >下载</el-button>
      </el-table-column>
    </el-table>
    <el-pagination
      layout="total, sizes, prev, pager, next, jumper"
      :current-page="currentPage"
      :page-sizes="[10, 20, 30, 50, 100]"
      :page-size="pageSize"
      :total="list.length"
      @size-change="handleSizeChange"
      @current-change="handleCurrentChange"
    ></el-pagination>
  </el-main>
</template>

<script>
export default {
  name: "Download",
  mounted() {
    this.getVersions();
  },
  data() {
    return {
      isLoading: true,
      intelval: null,
      list: [],
      currentPage: 1,
      pageSize: 10
    };
  },
  methods: {
    getVersions() {
      window.backend.lists().then(result => {
        this.list = result;
        this.isLoading = false;
      });
    },
    // 每页多少条
    handleSizeChange(val) {
      this.pageSize = val;
    },
    // 当前页
    handleCurrentChange(val) {
      this.currentPage = val;
    }
  }
};
</script>
