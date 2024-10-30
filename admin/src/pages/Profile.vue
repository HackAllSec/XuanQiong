<template>
    <div>
        <el-card style="padding: 30px;" :header="t('app.webui.myprofile')">
            <div style="display: flex;">
                <el-upload accept=".png,.jpg" action="/api/v1/updateavatar" :on-success="handleSuccess" :headers="{'Authorization': `Bearer ${token}`}" :before-upload="beforeUpload">
                    <el-avatar size="large" :src="avatar"></el-avatar>
                </el-upload>
            </div>
            <div style="margin-top: 20px; font-size: 14px;">{{ t('app.webui.username') }}</div>
            <el-input v-model="userinfo.username" size="large" :placeholder="userinfo.username" style="margin-top: 10px;" @change="changeButtonstatus" />
            <div style="margin-top: 20px; font-size: 14px;">{{ t('app.webui.email') }}</div>
            <el-input v-model="userinfo.email" size="large" :placeholder="userinfo.email" style="margin-top: 10px;"  @change="checkEmail" />
            <div style="margin-top: 20px; font-size: 14px;">{{ t('app.webui.phone') }}</div>
            <el-input v-model="userinfo.phone" size="large" :placeholder="userinfo.phone" style="margin-top: 10px;"  @change="checkPhone" />
            <el-button type="primary" style="margin-top: 20px;" :disabled="buttonstatus" @click="modifyUserInfo">{{ t('app.webui.modify') }}</el-button>
        </el-card>
    </div>
</template>
<script lang="ts" setup>
    import { ref, onMounted } from 'vue'
    import { useI18n } from 'vue-i18n';
    import api from '../api'

    const { t } = useI18n();
    const token = sessionStorage.getItem('token')
    const userinfo = ref({})
    const buttonstatus = ref(true)
    const mountedFunctions = [getUserinfo]
    const avatar = ref(sessionStorage.getItem('avatar'))
    onMounted(() => {
        mountedFunctions.forEach(fn => {
            fn();
        });
    });
    async function getUserinfo() {
        try {
            const config = {
                headers: {
                    'Authorization': `Bearer ${token}`  // 使用Bearer schema
                }
            }
            const response = await api.get('/api/v1/userinfo', config)
            userinfo.value = response.data.data
            if (token && response.data.code == 0) {
                sessionStorage.removeItem('token')
                sessionStorage.removeItem('username')
                sessionStorage.removeItem('avatar')
                location.reload()
            }
        } catch (error) {
            // 处理请求错误
            //ElMessage.error(t('app.webui.loginerr2'));
        }
     }
    const beforeUpload = (file: File) => {
        avatar.value = URL.createObjectURL(file)
    }
    const handleSuccess = (response) => {
        ElMessage.success(t('app.webui.uploadsucc'))
        sessionStorage.setItem('avatar', '/download/file?id=' + response.data)
        location.reload()
    }
    const checkEmail = () => {
        if (!userinfo.value.email.match(/^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$/)) {
            ElMessage.error(t('app.webui.emailformat'))
            return
        }
        changeButtonstatus()
    }
    const checkPhone = () => {
        if (!userinfo.value.phone.match(/^[1][3,4,5,6,7,8,9][0-9]{9}$/)) {
            ElMessage.error(t('app.webui.phoneformat'))
            return
        }
        changeButtonstatus()
    }
    const changeButtonstatus = () => {
        buttonstatus.value = false
    }

    const modifyUserInfo = async () => {
        console.log(userinfo.value)
        try {
            const config = {
                headers: {
                    'Authorization': `Bearer ${token}`  // 使用Bearer schema
                }
            };
            const data = {
                "username": userinfo.value.username,
                "email": userinfo.value.email,
                "phone": userinfo.value.phone
            }
            const response = await api.post('/api/v1/updateuserinfo', data, config)
            if (response.data.code == 0) {
                sessionStorage.removeItem('token')
                sessionStorage.removeItem('username')
                sessionStorage.removeItem('avatar')
                location.reload()
            } else if (response.data.code == 1) {
                ElMessage.success(t('app.webui.modifysucc'))
            } else {
                ElMessage.error(t('app.webui.modifyerr'))
            }
        } catch (error) {
            // 处理请求错误
            console.error(error);
        }
    }
</script>