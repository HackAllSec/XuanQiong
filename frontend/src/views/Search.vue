<template>
    <el-empty :description="t('app.webui.empty')" />
</template>
<script lang="ts" setup>
import { useI18n } from 'vue-i18n';

const { t } = useI18n();
</script>
<style scoped>

</style>