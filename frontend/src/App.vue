<template>
    <div>
        <el-container class="main">
            <el-aside :width="tabWidth+'px'">
                <div>
                    <div class="isClossTab" @click="isClossTabFun">
                        <i :class="isCollapse?'el-icon-d-arrow-right':'el-icon-d-arrow-left'"></i>
                    </div>
                    <el-menu :class="'menu'"
                             default-active="1"
                             class="el-menu-vertical-demo"
                             @open="handleOpen"
                             @close="handleClose"
                             :collapse="isCollapse"
                             background-color="#545c64"
                             text-color="#fff"
                             active-text-color="#ffd04b"
                             active-background-color="#fff"
                    >

                        <el-menu-item index="1">
                            <i class="el-icon-box"></i>
                            <span slot="title">游戏版本</span>
                        </el-menu-item>
                    </el-menu>
                </div>
            </el-aside>
            <el-container>
                <el-header class="main-header">
                    <span  style="margin: 20px;text-align: center;height: 50px;line-height: 50px;">Minecraft Downloader</span>
                    <el-dropdown>
                        <span class="el-dropdown-link">
                            <img src="" alt="">
                        </span>
                        <el-dropdown-menu slot="dropdown">
                            <el-dropdown-item>退出登录</el-dropdown-item>
                        </el-dropdown-menu>
                    </el-dropdown>
                </el-header>
                <el-main>
                    <template>
                        <el-table
                                :data="list"
                                stripe
                                highlight-current-row
                                style="width: 100%">
                            <el-table-column
                                    type="index"
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
                    </template>
                </el-main>
                <el-footer class="main-footer" height="50px">
                    <p>Power By Enjoy</p>
                </el-footer>
            </el-container>
        </el-container>
    </div>
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
                isCollapse: false,
                tabWidth: 200,
                test1: 1,
                intelval: null,
                list: [],
            };
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
                    this.tabWidth = 64;
                } else {
                    this.tabWidth = 200;
                }
                this.isCollapse = !this.isCollapse;
            },
            getVersions() {
                window.backend.lists().then(result => {
                    this.list = result
                });
            }
        }
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