<template>
    <div class="login-container">
        <el-card class="login-card" :header="t('app.webui.login')" shadow="always">
            <el-input
                v-model="uname"
                maxlength="50"
                :placeholder="t('app.webui.username')"
                size="large"
                :prefix-icon="User"
                clearable
            />
            <el-input
                v-model="passwd"
                maxlength="50"
                type="password"
                show-password
                :placeholder="t('app.webui.password')"
                size="large"
                :prefix-icon="Lock"
                clearable
                @keyup.enter.native="Login"
            />
            <div class="login-option">
                <el-link href="#/forgotpwd" style="margin-right: 10px;">{{ t('app.webui.forgot') }}</el-link>
                <el-link href="#/register">{{ t('app.webui.register') }}</el-link>
            </div>
            <el-button
                type="success"
                size="large"
                style="width: 60%; margin: 2% 20%; margin-bottom: 20px; font-weight: bold; font-size: 16px;"
                @click="Login"
                auto-insert-space
                >{{t('app.webui.login')}}
            </el-button>
        </el-card>
    </div>
</template>

<script lang="ts" setup>
    import { ref } from 'vue'
    import { User, Lock } from '@element-plus/icons-vue'
    import api from '../api'
    import { LoginPayload, LoginResponse } from '../types'
    import { useRouter } from 'vue-router'
    import { onMounted } from 'vue';
    import { useI18n } from 'vue-i18n';
    const { t } = useI18n();
    const router = useRouter();
    const uname = ref('');
    const passwd = ref('');
    
    onMounted(
        () => {
            if (sessionStorage.getItem('token') != null) {
                router.push('/')
            }
        }
    )

    async function Login() {
        if (uname.value == '' || passwd.value == '') {
            ElMessage.error(t('app.webui.loginerr1'));
            return;
        }
        try {
            const response = await api.post<LoginResponse>('/api/v1/login', {
                username: uname.value,
                password: passwd.value,
            } as LoginPayload);
            // 检查登录是否成功
            if (response.data.code == 0) {
                ElMessage.error(t('app.webui.loginerr3'));
            }
            if (response.data.code == 1) {
                ElNotification({
                    title: t('app.webui.loginsucc'),
                    message: response.data.username + ', ' + t('app.webui.welcome') ,
                    type: 'success',
                });
                await new Promise((resolve) => setTimeout(resolve, 1000))
                sessionStorage.setItem('token', response.data.token);
                sessionStorage.setItem('username', response.data.username);
                if (response.data.avatar != '') {
                    sessionStorage.setItem('avatar', '/download/file?id=' + response.data.avatar);
                } else {
                    sessionStorage.setItem('avatar', '/avatar.svg');
                }
                router.push('/')
            }
            if (response.data.code == 2) {
                ElMessage.error(t('app.webui.inputformaterror'));
            }
            if (response.data.code == 3) {
                ElMessage.error(t('app.webui.loginerr4') + ' ' + response.data.times + ' ' + t('app.webui.times'))
            }
        } catch (error) {
            // 处理请求错误
            ElMessage.error(t('app.webui.loginerr2'));
        }
    }
    
</script>

<style scoped>
  .login-container {
    display: flex;
    justify-content: center;
  }

  .login-card {
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

  .login-option {
    display: flex;
    justify-content: right;
    margin: 2% 10%;
  }
</style>