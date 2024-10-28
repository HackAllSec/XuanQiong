<template>
    <div class="forgot-container">
        <el-card class="forgot-card" :header="t('app.webui.forgot')" shadow="always">
            <el-input
                v-model="uname"
                maxlength="50"
                :placeholder="t('app.webui.username')"
                size="large"
                :prefix-icon="User"
                clearable
                @blur="checkuname"
            />
            <el-input
                v-model="email"
                maxlength="50"
                :placeholder="t('app.webui.email')"
                size="large"
                :prefix-icon="Message"
                clearable
                @blur="checkemail"
            />
            <el-input
                v-model="passwd"
                maxlength="50"
                type="password"
                show-password
                :placeholder="t('app.webui.password')"
                size="large"
                :prefix-icon="Unlock"
                clearable
                @blur="checkpasswd"
            />
            <el-input
                v-model="cfmpasswd"
                maxlength="50"
                type="password"
                show-password
                :placeholder="t('app.webui.confirmpassword')"
                size="large"
                :prefix-icon="Lock"
                clearable
                @blur="checkcfmpasswd"
            />
            <div class="forgot-option">
                <el-link href="#/login">{{ t('app.webui.returnlogin') }}</el-link>
            </div>
            <el-button
                type="success"
                size="large"
                style="width: 60%; margin: 2% 20%; margin-bottom: 20px; font-weight: bold; font-size: 16px;"
                @click="forgot"
                auto-insert-space
                >{{ t('app.webui.forgot') }}
            </el-button>
        </el-card>
    </div>
</template>

<script lang="ts" setup>
    import { ref } from 'vue'
    import { User, Lock, Unlock, Message } from '@element-plus/icons-vue'
    import { useI18n } from 'vue-i18n';
    import api from '../api'

    const { t } = useI18n()
    const uname = ref('');
    const passwd = ref('');
    const cfmpasswd = ref('');
    const email = ref('');
    function checkuname() {
        if (uname.value == '') {
            ElMessage.error(t('app.webui.loginerr1'));
        }
    }
    function checkemail() {
        if (email.value == '') {
            ElMessage.error(t('app.webui.emailempty'));
            return;
        }
        if (!email.value.match(/^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$/)) {
            ElMessage.error(t('app.webui.emailformat'));
            return;
        }
    }
    function checkpasswd() {
        if (passwd.value == '') {
            ElMessage.error(t('app.webui.passwordempty'));
            return false
        }
        if (passwd.value.length < 8) {
            ElMessage.error(t('app.webui.passwordlength'))
            return false
        }
        if (!passwd.value.match(/^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[^]{8,}$/)) {
            ElMessage.error(t('app.webui.passwordcomplex'))
            return false
        }
        return true
    }
    function checkcfmpasswd() {
        if (passwd.value != cfmpasswd.value) {
            ElMessage.error(t('app.webui.passwordnotmatch'));
            return;
        }
    }
    async function forgot() {
        if (uname.value == '' || passwd.value == '' || cfmpasswd.value == '' || email.value == '') {
            ElMessage.error(t('app.webui.missingnotice'));
            return;
        }
        if (passwd.value != cfmpasswd.value) {
            ElMessage.error(t('app.webui.passwordnotmatch'));
            return;
        }
        if (!checkpasswd()) {
            return;
        }
        try {
            const data = {
                "username": uname.value,
                "password": cfmpasswd.value,
                "email": email.value,
            }
            const response = await api.post('/api/v1/forgot', data)
            if (response.data.code == 0) {
                ElMessage.error(t('app.webui.forgotnotice'))
            } else if (response.data.code == 1) {
                ElMessage.success(t('app.webui.forgotsuccess'))
            } else if (response.data.code == 2) {
                ElMessage.error(t('app.webui.usernamealreadyexist'))
            } else if (response.data.code == 4) {
                ElMessage.error(t('app.webui.emailformat'))
            } else if (response.data.code == 5) {
                ElMessage.error(t('app.webui.emailalreadyexist'))
            } else {
                ElMessage.error(t('app.webui.forgotfail'))
            }
        } catch (error){
            ElMessage.error(t('app.webui.forgotfail'))
        }
    }
    
</script>

<style scoped>
  .forgot-container {
    display: flex;
    justify-content: center;
  }

  .forgot-card {
    width: 500px;
    margin-top: 10%;
    margin-bottom: 10%;
    font-size: 20px;
    font-weight: bold;
    /*background: #303030;*/
  }
  
  .el-input {
    width: 80%;
    margin: 0 10%;
    padding: 15px;
  }
  .forgot-option {
    display: flex;
    justify-content: right;
    margin: 2% 10%;
  }
</style>