import 'core-js/stable';
import 'regenerator-runtime/runtime';
import Vue from 'vue';
import App from './App.vue';

Vue.config.productionTip = false;
Vue.config.devtools = true;

import * as Wails from '@wailsapp/runtime';
import './plugins/element.js'

Wails.Init(() => {
	new Vue({
		render: h => h(App)
	}).$mount('#app');
});
