<template>
  <el-config-provider :locale="elementPlusLocale">
    <el-container>
    <el-header>
      <Navbar />
    </el-header>
    <el-main>
      <router-view />
    </el-main>
    <Footer />
  </el-container>
  <el-backtop :right="100" :bottom="100" />
  </el-config-provider>
</template>

<script lang="ts">
import { defineComponent, computed } from 'vue';
import Navbar from './components/Navbar.vue';
import Footer from './components/Footer.vue';

import { useI18n } from 'vue-i18n';
import zhCn from 'element-plus/es/locale/lang/zh-cn';
import en from 'element-plus/es/locale/lang/en';

export default defineComponent({
  name: 'App',
  components: {
    Navbar,
    Footer
  },
  setup() {
    const { locale } = useI18n();
    // 计算属性，根据vue-i18n的locale值返回对应的Element Plus语言包
    const elementPlusLocale = computed(() => {
      if (locale.value === 'zh-CN') {
        return zhCn;
      } else {
        return en;
      }
    });

    return {
      elementPlusLocale
    };
  }
});
</script>

<style scoped>
.el-container {
  /*background-color: #383737;*/
}
.el-header {
  --el-header-padding: 0 0;
}
</style>