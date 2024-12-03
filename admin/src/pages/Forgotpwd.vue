<template>
    <div class="forgot-container">
        <el-card class="forgot-card" :header="t('app.webui.forgot')" shadow="always">
            <p><el-text class="mx-1" type="danger">{{ t('app.webui.forgotnotice') }}</el-text></p>
            <p><el-text class="mx-1" type="danger">{{ t('app.webui.fstep1') }}</el-text></p>
            <p><el-text class="mx-1" type="danger">{{ t('app.webui.fstep2') }}</el-text></p>
            <p><el-text class="mx-1" type="danger">{{ t('app.webui.forgotnotice2') }}</el-text></p>
            <div style="display: flex; justify-content: center;">
                <el-button type="primary" @click="goback" auto-insert-space>{{ t('app.webui.confirm') }}</el-button>
            </div>
        </el-card>
    </div>
</template>

<script lang="ts" setup>
    import { useI18n } from 'vue-i18n';
    import { useRouter } from 'vue-router';

    const { t } = useI18n()
    const router = useRouter()
    const goback = () => {
        router.push('/')
    }
    
</script>

<style scoped>
  .forgot-container {
    display: flex;
    justify-content: center;
  }

  .forgot-card {
    width: 600px;
    margin-top: 10%;
    margin-bottom: 10%;
    font-size: 20px;
    font-weight: bold;
    /*background: #303030;*/
  }
</style>