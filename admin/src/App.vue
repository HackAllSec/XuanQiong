<template>
  <el-config-provider :locale="elementPlusLocale">
      <router-view />
  </el-config-provider>
  <el-backtop :right="100" :bottom="100" />
</template>

<script lang="ts">
import { defineComponent, computed } from 'vue';
import { useI18n } from 'vue-i18n';
import zhCn from 'element-plus/es/locale/lang/zh-cn';
import en from 'element-plus/es/locale/lang/en';

export default defineComponent({
  name: 'App',
  components: {
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
</style>