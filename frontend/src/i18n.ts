// i18n.ts
import { createI18n } from 'vue-i18n';
import zhCn from 'element-plus/es/locale/lang/zh-cn'
import en from 'element-plus/es/locale/lang/en'
import zhCN from './lang/zh-CN';
import enUS from './lang/en-US';

export const i18n = createI18n({
  locale: 'zh-CN', // 默认语言
  messages: {
    'zh-CN': {
      ...zhCn,
      ...zhCN
    },
    'en-US': {
      ...en,
      ...enUS
    }
  }
});