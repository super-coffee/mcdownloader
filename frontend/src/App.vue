<template>
    <div id="app" :style="{height:'100%'}">
        <a-layout id="components-layout-demo-custom-trigger" :style="{height:'100%'}">
            <a-layout-sider v-model="collapsed" collapsible :collapsedWidth=80>
                <div class="logo">MC Download</div>
                <a-menu theme="dark" mode="inline" :default-selected-keys="['/download']">
                    <a-menu-item key="/download">
                        <router-link to="/download">
                            <a-icon type="download"/>
                            <span>游戏下载</span>
                        </router-link>
                    </a-menu-item>
                    <a-menu-item key="/chooseMirror">
                        <router-link to="/chooseMirror">
                            <a-icon type="cloud"/>
                            <span>下载源</span>
                        </router-link>
                    </a-menu-item>
                </a-menu>
            </a-layout-sider>
            <a-layout>
                <a-layout-header style="background: #fff; padding: 0">
                    <a-icon
                            class="trigger"
                            :type="collapsed ? 'menu-unfold' : 'menu-fold'"
                            @click="() => (collapsed = !collapsed)"
                    />
                </a-layout-header>
                <a-layout-content
                        :style="{ margin: '24px 16px', padding: '24px', background: '#fff', minHeight: '280px' }"
                >
                    <router-view/>
                </a-layout-content>
            </a-layout>
        </a-layout>
    </div>
</template>
<style>
    #components-layout-demo-custom-trigger .trigger {
        font-size: 18px;
        line-height: 64px;
        padding: 0 24px;
        cursor: pointer;
        transition: color 0.3s;
    }

    #components-layout-demo-custom-trigger .trigger:hover {
        color: #1890ff;
    }

    #components-layout-demo-custom-trigger .logo {
        height: 32px;
        background: rgba(255, 255, 255, 0.2);
        margin: 16px;
        color: #eeeeee;
        text-align: center;
    }
</style>
<script>
    export default {
        name: "app",
        mounted() {

        },
        created () {
            this.getBreadcrumb()
        },
        data() {
            return {
                collapsed: false,
                isCollapse: false,
                tabWidth: 200,
                drawer: false,
                breadList: [],
                name: '',
            };
        },
        methods: {
            getBreadcrumb () {
                this.breadList = []
                this.name = this.$route.name
                this.$route.matched.forEach(item => {
                    this.breadList.push(item)
                })
            },
        },
        watch: {
            $route () {
                this.getBreadcrumb()
            }
        }

    };

</script>