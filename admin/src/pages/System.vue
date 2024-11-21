<template>
    <el-card :header="t('app.webui.regandlogincnfg')">
        <el-form :disabled="editable" v-model="data.sysconf" label-width="auto" style="margin-left: 2%; max-width: 900px">
            <el-form-item :label="t('app.webui.userregister')">
                <el-switch v-model="data.sysconf.user_register" style="--el-switch-on-color: #13ce66;" />
            </el-form-item>
            <el-form-item :label="t('app.webui.userdisplay')">
                <el-input v-model="data.sysconf.user_display" />
            </el-form-item>
            <el-form-item :label="t('app.webui.maxattempts')" required>
                <el-input type="number" v-model="data.sysconf.max_attempts" />
            </el-form-item>
            <el-form-item :label="t('app.webui.lockoutduration')" required>
                <el-input type="number" v-model="data.sysconf.lockout_duration" />
            </el-form-item>
            <el-form-item>
                <span style="color: red;">{{ t('app.webui.needemail') }}</span>
            </el-form-item>
            <el-form-item :label="t('app.webui.emailserver')">
                <el-input v-model="data.emailconf.email_host" />
            </el-form-item>
            <el-form-item :label="t('app.webui.emailport')">
                <el-input type="number" v-model="data.emailconf.email_port" />
            </el-form-item>
            <el-form-item :label="t('app.webui.emailuser')">
                <el-input v-model="data.emailconf.email_user" />
            </el-form-item>
            <el-form-item :label="t('app.webui.emailpasswd')">
                <el-input type="password" show-password v-model="data.emailconf.email_password" />
            </el-form-item>
            <el-form-item :label="t('app.webui.emailsender')">
                <el-input v-model="data.emailconf.email_sender" />
            </el-form-item>
        </el-form>
    </el-card>
    <el-card :header="t('app.webui.jwtconfig')" style="margin-top: 2%;">
        <el-form :disabled="editable" :model="data.jwtconf" label-width="auto" style="margin-left: 2%; max-width: 900px">
            <el-form-item :label="t('app.webui.jwtkey')" required>
                <el-input type="password" show-password v-model="data.jwtconf.jwt_secret" />
            </el-form-item>
            <el-form-item :label="t('app.webui.jwtvalidity')" required>
                <el-input type="number" v-model="data.jwtconf.jwt_expires" />
            </el-form-item>
        </el-form>
    </el-card>
    <el-card :header="t('app.webui.noticeconfig')" style="margin-top: 2%;">
        <el-form :disabled="editable" :model="data.noticeconf" label-width="auto" style="margin-left: 2%; max-width: 900px">
            <el-form-item :label="t('app.webui.noticetype')">
                <el-select v-model="data.noticeconf.type">
                    <el-option v-for="item in noticetype" :key="item.key" :label="item.value" :value="item.key" />
                </el-select>
            </el-form-item>
            <el-form-item label="Webhook">
                <el-input v-model="data.noticeconf.webhook"/>
            </el-form-item>
            <el-form-item v-if="data.noticeconf.type == 1" label="Secret">
                <el-input type="password" show-password v-model="data.noticeconf.secret" />
            </el-form-item>
        </el-form>
    </el-card>
    <div v-if="editable" style="margin-top: 1%;">
        <el-button type="primary" @click="editable=false">{{ t('app.webui.edit') }}</el-button>
    </div>
    <div v-else style="margin-top: 1%;">
        <el-button type="primary" @click="Save">{{ t('app.webui.save') }}</el-button>
        <el-button @click="editable=true">{{ t('app.webui.cancel') }}</el-button>
    </div>
</template>
<script lang="ts" setup>
    import {ref, onMounted} from "vue";
    import {useI18n} from "vue-i18n";
    import api from '../api'
    
    const { t } = useI18n();
    const editable = ref(true)
    const token = sessionStorage.getItem("token")
    const data = ref({
        emailconf: {
            email_host: "",
            email_port: 0,
            email_user: "",
            email_password: "",
            email_sender: ""
        },
        jwtconf: {
            jwt_secret: "",
            jwt_expires: 0
        },
        noticeconf: {
            type: 0,
            secret: "",
            webhook: ""
        },
        sysconf: {
            user_register: false,
            user_display: "",
            max_attempts: 0,
            lockout_duration: 0
        }
    })
    const noticetype = ref([{"key": 0, "value": t('el.select.placeholder')}, {"key": 1, "value": t('app.webui.dingtalk')},{"key": 2, "value": t('app.webui.wxwork')}])
    
    onMounted(getSysConfig)
    async function getSysConfig() {
        try {
            const config = {
            headers: {
                    'Authorization': `Bearer ${token}`
                }
            };
            const response = await api.get('/api/v1/getsysconfig', config)
            if (response.data.code == 1) {
                data.value = response.data.data
            } else {
                sessionStorage.removeItem('token')
                sessionStorage.removeItem('username')
                sessionStorage.removeItem('avatar')
            }
        } catch (error) {
            //console.log(error)
        }
    }

    async function Save() {
        console.log(data.value)
        try {
            const config = {
            headers: {
                    'Authorization': `Bearer ${token}`
                }
            };
            data.value.emailconf.email_port = Number(data.value.emailconf.email_port)
            data.value.jwtconf.jwt_expires = Number(data.value.jwtconf.jwt_expires)
            data.value.noticeconf.type = Number(data.value.noticeconf.type)
            data.value.sysconf.max_attempts = Number(data.value.sysconf.max_attempts)
            data.value.sysconf.lockout_duration = Number(data.value.sysconf.lockout_duration)
            const response = await api.post('/api/v1/updatesysconfig', data.value, config)
            if (response.data.code == 1) {
                getSysConfig()
                ElMessage({
                    type: 'success',
                    message: t('app.webui.savesuccess'),
                })
                editable.value = true
            } else {
                ElMessage({
                    type: 'error',
                    message: t('app.webui.savefail'),
                })
            }
        } catch (error) {
            // 处理请求错误
            //ElMessage.error(t('app.webui.loginerr2'));
        }
    }
</script>
<style scoped>

</style>