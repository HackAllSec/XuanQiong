<template>
    <div class="register-container">
        <el-card class="register-card" :header="t('app.webui.register')" shadow="always">
            <el-input
                style="width: 75%; margin-left: 10%; padding: 15px;"
                v-model="uname"
                maxlength="50"
                :placeholder="t('app.webui.username')"
                size="large"
                :prefix-icon="User"
                clearable
                @blur="checkuname"
            />
            <el-input
                style="width: 75%; margin-left: 10%; padding: 15px;"
                v-model="phone"
                maxlength="50"
                :placeholder="t('app.webui.phone')"
                size="large"
                :prefix-icon="Phone"
                clearable
            />
            <div style="display: flex; align-items: center;">
                <el-input
                style="width: 75%; margin-left: 10%; padding: 15px;"
                v-model="email"
                maxlength="50"
                :placeholder="t('app.webui.email')"
                size="large"
                :prefix-icon="Message"
                clearable
                @blur="checkemail"
                />
                <el-button v-if="showcaptcha" style="width: 15%;" type="primary" size="small" @click="getCaptcha">{{ t('app.webui.getcaptcha') }}</el-button>
                <el-countdown v-else format="ss" :value="remaintime" @finish="showcaptcha=true" value-style="font-size: 14px;" />
            </div>
            <el-input
                style="width: 75%; margin-left: 10%; padding: 15px;"
                v-model="captcha"
                maxlength="50"
                :placeholder="t('app.webui.captcha')"
                size="large"
                :prefix-icon="Camera"
                clearable
            />
            <el-input
                style="width: 75%; margin-left: 10%; padding: 15px;"
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
                style="width: 75%; margin-left: 10%; padding: 15px;"
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
            <div class="register-option">
                <el-link href="#/login">{{ t('app.webui.returnlogin') }}</el-link>
            </div>
            <el-button
                type="success"
                size="large"
                style="width: 60%; margin: 2% 20%; margin-bottom: 20px; font-weight: bold; font-size: 16px;"
                @click="Register"
                auto-insert-space
                >{{ t('app.webui.register') }}
            </el-button>
        </el-card>
    </div>
</template>

<script lang="ts" setup>
    import { ref } from 'vue'
    import { User, Lock, Unlock, Message, Camera, Phone } from '@element-plus/icons-vue'
    import { useI18n } from 'vue-i18n';
    import api from '../api'

    const { t } = useI18n()
    const uname = ref('');
    const phone = ref('');
    const passwd = ref('');
    const cfmpasswd = ref('');
    const email = ref('');
    const captcha = ref('');
    const showcaptcha = ref(true);
    const remaintime = ref(0);
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
    async function getCaptcha() {
        if (email.value == '') {
            ElMessage.error(t('app.webui.emailempty'))
            return;
        }
        try {
            const response = await api.get('/api/v1/getcaptcha?email=' + email.value)
            if (response.data.code == 1) {
                ElMessage.success(t('app.webui.captchasucc'))
                remaintime.value = Date.now() + 1000 * 120
                showcaptcha.value = false
            } else {
                ElMessage.error(t('app.webui.captchafail'))
            }
        } catch (error){
            //错误处理
        }
    }
    async function Register() {
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
                "phone": phone.value,
                "password": cfmpasswd.value,
                "email": email.value,
                "captcha": captcha.value,
            }
            const response = await api.post('/api/v1/register', data)
            if (response.data.code == 0) {
                ElMessage.error(t('app.webui.registernotice'))
            } else if (response.data.code == 1) {
                ElMessage.success(t('app.webui.registersuccess'))
            } else if (response.data.code == 2) {
                ElMessage.error(t('app.webui.usernamealreadyexist'))
            } else if (response.data.code == 4) {
                ElMessage.error(t('app.webui.emailformat'))
            } else if (response.data.code == 5) {
                ElMessage.error(t('app.webui.emailalreadyexist'))
            } else if (response.data.code == 6) {
                ElMessage.error(t('app.webui.captchaerr'))
            } else {
                ElMessage.error(t('app.webui.registerfail'))
            }
        } catch (error){
            ElMessage.error(t('app.webui.registerfail'))
        }
    }
    
</script>

<style scoped>
  .register-container {
    display: flex;
    justify-content: center;
  }

  .register-card {
    width: 500px;
    margin-top: 10%;
    margin-bottom: 10%;
    font-size: 20px;
    font-weight: bold;
    /*background: #303030;*/
  }
  
  .register-option {
    display: flex;
    justify-content: right;
    margin: 2% 10%;
  }
</style>