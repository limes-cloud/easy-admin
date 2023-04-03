import { createApp } from 'vue';
import ArcoVue from '@arco-design/web-vue';
import ArcoVueIcon from '@arco-design/web-vue/es/icon';
import globalComponents from '@/components';
import { parseTime, formatTime, hasPermission } from '@/utils';
import { densityList } from '@/utils/consts';
import router from './router';
import store, { useAppStore } from './store';
import i18n from './locale';
import directive from './directive';
// import './mock';
import App from './App.vue';
import '@arco-design/web-vue/dist/arco.css';
import '@/assets/style/global.less';
import '@/api/interceptor';

const app = createApp(App);

app.use(ArcoVue, {});
app.use(ArcoVueIcon);

app.use(router);
app.use(store);
app.use(i18n);
app.use(globalComponents);
app.use(directive);

document.title = useAppStore().title;

app.config.globalProperties.$staticUrl = `${
  import.meta.env.VITE_STATIC_BASE_URL
}/`;
app.config.globalProperties.$parseTime = parseTime;
app.config.globalProperties.$formatTime = formatTime;
app.config.globalProperties.$densityList = densityList;
app.config.globalProperties.$hasPermission = hasPermission;

app.mount('#app');
