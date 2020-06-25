<template>
    <el-main>
        <el-card class="box-card">
            <div slot="header" class="clearfix">
                <span>下载源选择</span>

            </div>
            <el-form>
                <el-form-item>
                    <el-radio-group v-model="id">
                        <el-radio :label="1">MC官方源</el-radio>
                        <el-radio :label="2">BMCL镜像源</el-radio>
                        <el-radio :label="3">MCBBS镜像源</el-radio>
                    </el-radio-group>
                </el-form-item>
                <el-form-item>
                    <el-button style="float: right;" @click="setServerId(id)" type="primary" round>保存</el-button>
                </el-form-item>
            </el-form>
        </el-card>
    </el-main>

</template>

<script>

    export default {
        name: 'app',
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
                    this.id = result
                });
            },
            setServerId(id) {
                window.backend.setServerId(id).then(result => {
                    if (result === true) {
                        this.$message({
                            message: "设置成功",
                            type: "success",
                        });
                    } else {
                        this.$message({
                            message: "设置失败",
                            type: "danger",
                        });
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
            },
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