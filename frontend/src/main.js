import 'core-js/stable';
import 'regenerator-runtime/runtime';
import Vue from 'vue';
import App from './App.vue';

Vue.config.productionTip = false;
Vue.config.devtools = true;
import Antd from 'ant-design-vue';
import * as Wails from '@wailsapp/runtime';
import router from './router'
import 'ant-design-vue/dist/antd.css';

Vue.use(Antd);
Wails.Init(() => {
    new Vue({
        router,
        render: h => h(App),
        data() {
            return {
                from: 1
            }
        },
        template: '<App/>',
    }).$mount('#app');
});
