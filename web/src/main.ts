import { createApp } from 'vue';
import ArcoVue from '@arco-design/web-vue';
import ArcoVueIcon from '@arco-design/web-vue/es/icon';
import globalComponents from '@/components';
import { parseTime, formatTime, hasPermission } from '@/utils';
import { densityList, noticeList } from '@/utils/consts';
import systemConfig from '@/utils/config';
import icons from '@/icons/index.vue'; // 全局注册svg-icon组件
import router from './router';
import store from './store';
import i18n from './locale';
import directive from './directive';
// import './mock';
import App from './App.vue';
import '@arco-design/web-vue/dist/arco.css';
import '@/assets/style/global.less';
import '@/api/interceptor';

// eslint-disable-next-line import/no-unresolved
import 'virtual:svg-icons-register';

const app = createApp(App);

app.use(ArcoVue, {});
app.use(ArcoVueIcon);

app.use(router);
app.use(store);
app.use(i18n);
app.use(globalComponents);
app.use(directive);

app.component('SvgIcon', icons);
systemConfig().then((res) => {
  app.config.globalProperties.$parseTime = parseTime;
  app.config.globalProperties.$formatTime = formatTime;
  app.config.globalProperties.$densityList = densityList;
  app.config.globalProperties.$hasPermission = hasPermission;
  app.config.globalProperties.$noticeList = noticeList;
  app.config.globalProperties.$staticUrl = `${res.staticUrl}/`;
  app.config.globalProperties.$logo = res.logo;
  app.mount('#app');
});
