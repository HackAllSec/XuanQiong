<template>
    <div style="width: 70%; height: 90vh; margin: auto;">
        <el-card style="padding: 30px; font-weight: bold; font-size: 20px;" :header="t('app.webui.modifypassword')">
            <div style="margin-top: 20px; font-weight: lighter; font-size: 14px;">{{ t('app.webui.oldpassword') }}</div>
            <el-input v-model="oldpassword" size="large" :placeholder="oldpassword" style="margin-top: 10px;" type="password" show-password />
            <div style="margin-top: 20px; font-weight: lighter; font-size: 14px;">{{ t('app.webui.newpassword') }}</div>
            <el-input v-model="newpassword" size="large" :placeholder="newpassword" style="margin-top: 10px;" type="password" show-password @change="checkNewPassword" />
            <div style="margin-top: 20px; font-weight: lighter; font-size: 14px;">{{ t('app.webui.confirmpassword') }}</div>
            <el-input v-model="confirmpassword" size="large" :placeholder="confirmpassword" style="margin-top: 10px;" type="password" show-password @change="checkConfirmPassword" />
            <el-button type="primary" :disabled="buttonstatus" style="margin-top: 20px;" @click="changePassword">{{ t('app.webui.submit') }}</el-button>
        </el-card>
    </div>
</template>
<script lang="ts" setup>
import { ref,onMounted } from 'vue'
import { GoldMedal } from '@element-plus/icons-vue'
import { useI18n } from 'vue-i18n';
import { checkLogin } from '../utils';
import api from '../api';

const { t } = useI18n();
const buttonstatus = ref(true)
const token = sessionStorage.getItem('token')
const oldpassword = ref('')
const newpassword = ref('')
const confirmpassword = ref('')
const mountedFunctions = [checkLogin]

onMounted(() => {
        mountedFunctions.forEach(fn => {
            fn();
        });
    });

const checkNewPassword = () => {
    if (newpassword.value.length < 8) {
        ElMessage.error(t('app.webui.passwordlength'))
        return false
    }
    if (!newpassword.value.match(/^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[^]{8,}$/)) {
        ElMessage.error(t('app.webui.passwordcomplex'))
        return false
    }
    return true
}
const checkConfirmPassword = () => {
    if (newpassword.value != confirmpassword.value) {
        ElMessage.error(t('app.webui.passwordnotmatch'))
        return
    }
    if (checkNewPassword()) {
        changeButtonstatus()
    }
}

const changeButtonstatus = () => {
    buttonstatus.value = false
}
    
async function changePassword() {
    console.log(oldpassword,newpassword,confirmpassword)
    try {
        const config = {
            headers: {
                'Authorization': `Bearer ${token}`  // 使用Bearer schema
            }
        };
        const data = {
            "oldpassword": oldpassword.value,
            "newpassword": newpassword.value
        }
        const response = await api.post('/api/v1/updatepassword', data, config)
        if (response.data.code == 0) {
                sessionStorage.removeItem('token')
                sessionStorage.removeItem('username')
                sessionStorage.removeItem('avatar')
                location.reload()
        } else if (response.data.code == 1) {
            ElMessage.success(t('app.webui.modifysucc'))
        } else if (response.data.code == 3) {
            ElMessage.error(t('app.webui.oldpassworderr'))
        } else {
            ElMessage.error(t('app.webui.modifyerr'))
        }
    } catch (error){
        ElMessage.error(t('app.webui.modifyfail'))
    }
}
</script>