<template>
  <div id="app">
  <div>
    <el-container  class="main">
      <el-aside :width="tabWidth+'px'">
        <div>
          <div class="isClossTab" @click="isClossTabFun">
            <i :class="isCollapse?'el-icon-d-arrow-right':'el-icon-d-arrow-left'"></i>
          </div>
          <el-menu :class="'menu'"
                   :default-active="$route.path"
                   class="el-menu-vertical-demo"
                   @open="handleOpen"
                   @close="handleClose"
                   router
                   :collapse="isCollapse"
                   background-color="#545c64"
                   text-color="#fff"
                   active-text-color="#ffd04b"
                   active-background-color="#fff"
          >

            <el-menu-item index="/">
              <i class="el-icon-box"></i>
              <span slot="title">游戏版本</span>
            </el-menu-item>
            <el-menu-item index="/about">
              <i class="el-icon-download"></i>
              <span slot="title">下载设置</span>
            </el-menu-item>
          </el-menu>
        </div>
      </el-aside>
      <el-container>
        <el-header class="main-header">
          <span  style="margin: 20px;text-align: center;height: 50px;line-height: 50px;">Minecraft Downloader</span>

        </el-header>

          <router-view/>
        <el-footer class="main-footer" height="50px">
          <p>Power By Enjoy</p>
        </el-footer>
      </el-container>
    </el-container>
  </div>
  </div>
</template>
<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

#nav {
  padding: 30px;

  a {
    font-weight: bold;
    color: #2c3e50;

    &.router-link-exact-active {
      color: #42b983;
    }
  }
}
</style>
<script>
  export default {
    name: 'app',
    mounted() {

    },
    data() {
      return{
        isCollapse: false,
        tabWidth: 200,
      }
    },
    methods: {

      handleOpen(key, keyPath) {
        console.log(key, keyPath);
      },
      handleClose(key, keyPath) {
        console.log(key, keyPath);
      },
      isClossTabFun() {
        clearInterval(this.intelval);
        if (!this.isCollapse) {
          clearInterval(this.intelval);
          if (!this.isCollapse) {
            this.intelval = setInterval(() => {
              if (this.tabWidth <= 64)
                clearInterval(this.intelval);
              this.tabWidth -= 8;
            }, 1);
          }
        } else {
          this.tabWidth = 200;
        }
        this.isCollapse = !this.isCollapse;
      },
    },
  }
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
  $color: #FFF;

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
          background-color: #FFF;
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