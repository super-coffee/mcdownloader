<template>
  <el-main v-loading.lock="isLoading" :show-overflow-tooltip="true">
    <el-progress v-show="first" :status="status" :text-inside="true" :stroke-width="26" :percentage="step"></el-progress>
    <el-table
      stripe
      highlight-current-row
      style="width: 100%"
      :data="list.slice((currentPage-1)*pageSize, currentPage*pageSize)"
    >

      <el-table-column prop="id" label="#" width="50"></el-table-column>
      <el-table-column prop="versionId" label="版本名称" width="200"></el-table-column>
      <el-table-column prop="releaseTime" label="发布时间"></el-table-column>
      <el-table-column label="操作">
        <template slot-scope="scope">
        <el-button
                size="mini"
                type="primary"
                icon="el-icon-download"
                plain
                @click="mcDownload(scope.row.versionId)"
        >下载</el-button>
        <el-button
          size="mini"
          icon="el-icon-magic-stick"
          @click="mcDownload(scope.row.versionId)"
        >修复</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-divider></el-divider>
    <el-pagination
      layout="total, sizes, prev, pager, next, jumper"
      :current-page="currentPage"
      :page-sizes="[8, 20, 30, 50, 100]"
      :page-size="pageSize"
      :total="list.length"
      background
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
      first:false,
      downloading:false,
      step:0,
      isLoading: true,
      intelval: null,
      list: [],
      currentPage: 1,
      pageSize: 8,
      loading: null,
      status: null,
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
    },
    mcDownload(a){
      if (this.downloading ===true){
        this.$message.error('有任务正在下载，请稍后！');
        return
      }
      this.first = true;
      this.status = null;
      this.step = 1;
      this.downloading = true
      let that = this
      setTimeout(async function() {
        for (let i=0; i<=100; i++){
          that.step = i;
          await sleep(100)
          console.log(i)
        }
        that.downloading = false
        that.status = 'success'
        that.$message.success('下载成功!');

      }, 0)

      console.log(a)
    }
  }
};
function sleep(ms) {
  return new Promise(resolve => setTimeout(resolve, ms))
}
</script>
