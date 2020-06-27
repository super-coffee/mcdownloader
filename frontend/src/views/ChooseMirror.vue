<template>
  <a-card title="选择下载源">
    <a-form-model-item prop="region">
    <a-radio-group class="ant-radio-button" v-model="id">
      <a-radio-button :value="1">
        Mojang
      </a-radio-button>
      <a-radio-button :value="2">
        BMCL
      </a-radio-button>
      <a-radio-button :value="3">
        MCBBS
      </a-radio-button>
    </a-radio-group>
    </a-form-model-item>
    <a-form-model-item>
      <a-button type="primary" @click="setServerId(id)" >
        保存
      </a-button>
    </a-form-model-item>
  </a-card>
<!--  <el-main :show-overflow-tooltip="true">-->
<!--    <el-card class="box-card">-->
<!--      <div slot="header" class="clearfix">-->
<!--        <span>选择下载源</span>-->
<!--      </div>-->
<!--      <el-form>-->
<!--        <el-form-item>-->
<!--          <el-radio-group v-model="id">-->
<!--            <el-radio-button :label="1">Mojang</el-radio-button>-->
<!--            <el-radio-button :label="2">BMCL</el-radio-button>-->
<!--            <el-radio-button :label="3">MCBBS</el-radio-button>-->
<!--          </el-radio-group>-->
<!--        </el-form-item>-->
<!--        <el-form-item>-->
<!--          <el-button style="float: right;" @click="setServerId(id)" type="primary" round>保存</el-button>-->
<!--        </el-form-item>-->
<!--      </el-form>-->
<!--    </el-card>-->
<!--  </el-main>-->
</template>

<script>
export default {
  name: "ChooseMirror",
  mounted() {
    this.getServerId();
  },
  data() {
    return {
      id: 0
    };
  },
  methods: {
    getServerId() {
      window.backend.getServerId().then(result => {
        this.id = result;
      });
    },
    setServerId(id) {
      window.backend.setServerId(id).then(result => {
        if (result === true) {
          this.$message.success("设置成功");
        } else {
          this.$message.error("设置失败");
        }
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
<style>
* {
  padding: 0;
  margin: 0;
}
</style>
<style scoped lang="scss">
$header-height: 60px;
$background-color: #545c64;
$color: #fff;

.main {
  height: 100vh;
  min-width: 800px;
  min-height: 600px;
  overflow: hidden;

  aside {
    height: 100%;
    background-color: $background-color;
    color: $color;

    .isClossTab {
      width: 100%;
      height: $header-height;
      cursor: pointer;
      font-size: 25px;
      text-align: center;
      line-height: $header-height;
      font-weight: bold;
      border-right: 1px solid $background-color;
      box-sizing: border-box;
    }

    .menu {
      width: 100%;
      border-right: 0;
    }
  }

  .main-header {
    background-color: $background-color;
    color: $color;

    .el-dropdown {
      cursor: pointer;
      float: right;
    }

    .el-dropdown-link {
      img {
        $imgMargin: (($header-height - 50)/2);
        display: inline-block;
        width: 50px;
        height: 50px;
        border-radius: 25px;
        background-color: #fff;
        margin-top: $imgMargin;
      }
    }
  }

  .crumbs {
    margin-bottom: 20px;
  }

  .main-footer {
    text-align: center;
    background-color: $background-color;
    color: $color;
    line-height: 50px;
  }
}
</style>