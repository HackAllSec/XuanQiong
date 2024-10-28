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
                <el-link href="#/register" style="margin-right: 10px;">{{ t('app.webui.forgot') }}</el-link>
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
    import { jwtDecode } from 'jwt-decode'
    import { useI18n } from 'vue-i18n';

    const { t, locale } = useI18n();
    function performAction() {
        let token = sessionStorage.getItem('token');
        if (token) {
            try {
            const decodedToken = jwtDecode(token)
            let currentTime = Math.floor(Date.now() / 1000)
            if (currentTime > decodedToken.exp) {
                sessionStorage.removeItem('token')
                sessionStorage.removeItem('username')
                location.reload();
                return;
            }
            } catch (error) {
                sessionStorage.removeItem('token')
                sessionStorage.removeItem('username')
                location.reload();
                return;
            }
            router.push('/');
        }
    }

    onMounted(performAction);
    const router = useRouter();
    const uname = ref('');
    const passwd = ref('');

    async function Login() {
        if (uname.value == '' || passwd.value == '') {
            ElMessage.error('账号或密码不能为空');
            return;
        }
        try {
            const response = await api.post<LoginResponse>('/api/v1/login', {
                username: uname.value,
                password: passwd.value,
            } as LoginPayload);
            console.log(response.data);
            // 检查登录是否成功
            if (response.data.token) {
                ElNotification({
                    title: '登录成功',
                    message: response.data.username + ', 欢迎您' ,
                    type: 'success',
                });
                await new Promise((resolve) => setTimeout(resolve, 1000))
                sessionStorage.setItem('token', response.data.token);
                sessionStorage.setItem('username', response.data.username);
                router.push('/')
            } else {
                ElMessage.error(response.data.msg);
            }
            console.log(response.data.username);
            console.log(response.data.token);
        } catch (error) {
            // 处理请求错误
            console.error(error);
            ElMessage.error('登录请求失败');
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