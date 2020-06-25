import 'core-js/stable';
import 'regenerator-runtime/runtime';
import Vue from 'vue';
import App from './App.vue';
Vue.config.productionTip = false;
Vue.config.devtools = true;

import * as Wails from '@wailsapp/runtime';
import './plugins/element.js'
import router from './router'

Wails.Init(() => {
    new Vue({
        router,
        render: h => h(App),
        data() {
            return {
                from: 1
            }
        }
    }).$mount('#app');
});
