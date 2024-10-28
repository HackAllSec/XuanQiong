import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import router from './router'
import ElementPlus from 'element-plus'
import { i18n } from './i18n'

// 读取 localStorage 中的语言设置
const savedLanguage = localStorage.getItem('selectedLanguage');
if (savedLanguage) {
  i18n.global.locale = savedLanguage;
}
const app = createApp(App);
app.use(ElementPlus);
app.use(i18n);
app.use(router);
app.mount('#app');