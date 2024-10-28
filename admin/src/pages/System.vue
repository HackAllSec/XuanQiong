<template>
    <el-card :header="t('app.webui.regandlogincnfg')">
        <el-form :disabled="editable" :model="data.login" label-width="auto" style="margin-left: 2%; max-width: 900px">
            <el-form-item :label="t('app.webui.userregister')">
                <el-switch v-model="data.login.user_register" style="--el-switch-on-color: #13ce66;" />
            </el-form-item>
            <el-form-item :label="t('app.webui.maxattempts')">
                <el-input v-model="data.login.max_attempts" />
            </el-form-item>
            <el-form-item :label="t('app.webui.lockoutduration')">
                <el-input v-model="data.login.lockout_duration" />
            </el-form-item>
            <el-form-item :label="t('app.webui.emailserver')">
                <el-input v-model="data.login.email.email_host" />
            </el-form-item>
            <el-form-item :label="t('app.webui.emailuser')">
                <el-input v-model="data.login.email.email_user" />
            </el-form-item>
            <el-form-item :label="t('app.webui.emailpasswd')">
                <el-input v-model="data.login.email.email_password" />
            </el-form-item>
        </el-form>
    </el-card>
    <el-card :header="t('app.webui.jwtconfig')" style="margin-top: 2%;">
        <el-form :disabled="editable" :model="data.jwt" label-width="auto" style="margin-left: 2%; max-width: 900px">
            <el-form-item :label="t('app.webui.jwtkey')">
                <el-input v-model="data.jwt.jwt_secret" />
            </el-form-item>
            <el-form-item :label="t('app.webui.jwtvalidity')">
                <el-input v-model="data.jwt.jwt_expire" />
            </el-form-item>
        </el-form>
    </el-card>
    <el-card :header="t('app.webui.noticeconfig')" style="margin-top: 2%;">
        <el-form :disabled="editable" :model="data.notice" label-width="auto" style="margin-left: 2%; max-width: 900px">
            <el-form-item :label="t('app.webui.noticetype')">
                <el-select v-model="data.notice.type">
                    <el-option v-for="item in noticetype" :key="item.key" :label="item.value" :value="item.key" />
                </el-select>
            </el-form-item>
            <el-form-item label="Webhook">
                <el-input />
            </el-form-item>
            <el-form-item v-if="data.notice.type == 1" label="Secret">
                <el-input />
            </el-form-item>
        </el-form>
    </el-card>
    <div style="margin-top: 1%;">
        <el-button v-if="editable" type="primary" @click="editable=false">{{ t('app.webui.edit') }}</el-button>
        <el-button v-else type="primary" @click="Save">{{ t('app.webui.save') }}</el-button>
    </div>
</template>
<script lang="ts" setup>
    import {ref, onMounted} from "vue";
    import {useI18n} from "vue-i18n";
    
    const { t } = useI18n();
    const editable = ref(true)
    const data = ref({
        "login": {
            "id": 1,
            "user_register": true,
            "max_attempts": 5,
            "lockout_duration": 3600,
            "email": {
                "email_host": "smtp.qq.com",
                "email_port": 465,
                "email_user": "xxx@qq.com",
                "email_password": "xxx",
                "email_sender": "xxx@qq.com"
            }
        },
        "jwt": {
            "id": 1,
            "jwt_secret": "xxx",
            "jwt_expire": 3600
        },
        "notice": {
            "id": 1,
            "type": 2,
            "secret": "xxx",
            "webhook": "xxx"
        }
    })
    const noticetype = ref([{"key": 1, "value": t('app.webui.dingtalk')},{"key": 2, "value": t('app.webui.wxwork')}])
</script>
<style scoped>

</style>