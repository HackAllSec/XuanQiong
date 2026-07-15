<template>
  <el-card :header="t('app.webui.regandlogincnfg')">
    <el-form :disabled="readonly" :model="data.sysconf" label-width="150px" class="form">
      <el-form-item :label="t('app.webui.userregister')">
        <el-switch v-model="data.sysconf.user_register" />
      </el-form-item>
      <el-form-item :label="t('app.webui.userdisplay')">
        <el-input v-model="data.sysconf.user_display" />
      </el-form-item>
      <el-form-item :label="t('app.webui.maxattempts')">
        <el-input v-model="data.sysconf.max_attempts" type="number" />
      </el-form-item>
      <el-form-item :label="t('app.webui.lockoutduration')">
        <el-input v-model="data.sysconf.lockout_duration" type="number" />
      </el-form-item>
      <el-alert :title="t('app.webui.needemail')" type="warning" :closable="false" />
      <el-form-item :label="t('app.webui.emailserver')">
        <el-input v-model="data.emailconf.email_host" />
      </el-form-item>
      <el-form-item :label="t('app.webui.emailport')">
        <el-input v-model="data.emailconf.email_port" type="number" />
      </el-form-item>
      <el-form-item :label="t('app.webui.emailuser')">
        <el-input v-model="data.emailconf.email_user" />
      </el-form-item>
      <el-form-item :label="t('app.webui.emailpasswd')">
        <el-input v-model="data.emailconf.email_password" type="password" show-password />
      </el-form-item>
      <el-form-item :label="t('app.webui.emailsender')">
        <el-input v-model="data.emailconf.email_sender" />
      </el-form-item>
    </el-form>
  </el-card>

  <el-card :header="t('app.webui.brandconfig')" class="section">
    <el-form :disabled="readonly" :model="data.sysconf" label-width="150px" class="form">
      <el-form-item :label="t('app.webui.sitename')">
        <el-input v-model="data.sysconf.site_name" />
      </el-form-item>
      <el-form-item :label="t('app.webui.frontendtitle')">
        <el-input v-model="data.sysconf.frontend_title" />
      </el-form-item>
      <el-form-item :label="t('app.webui.admintitle')">
        <el-input v-model="data.sysconf.admin_title" />
      </el-form-item>
      <el-form-item :label="t('app.webui.footertext')">
        <el-input v-model="data.sysconf.footer_text" />
      </el-form-item>
      <el-form-item :label="t('app.webui.helpurl')">
        <el-input v-model="data.sysconf.help_url" />
      </el-form-item>
      <el-form-item :label="t('app.webui.suggesturl')">
        <el-input v-model="data.sysconf.suggest_url" />
      </el-form-item>
      <el-form-item :label="t('app.webui.logoattachment')">
        <div class="upload-row">
          <el-input v-model="data.sysconf.logo_attachment_id" />
          <el-upload v-if="canUpdate" :headers="uploadHeaders" action="/api/v1/upload" :show-file-list="false" :on-success="handleLogoUpload">
            <el-button type="primary">{{ t('app.webui.uploadlogo') }}</el-button>
          </el-upload>
        </div>
      </el-form-item>
      <el-form-item :label="t('app.webui.faviconattachment')">
        <div class="upload-row">
          <el-input v-model="data.sysconf.favicon_attachment_id" />
          <el-upload v-if="canUpdate" :headers="uploadHeaders" action="/api/v1/upload" :show-file-list="false" :on-success="handleFaviconUpload">
            <el-button type="primary">{{ t('app.webui.uploadfavicon') }}</el-button>
          </el-upload>
        </div>
      </el-form-item>
    </el-form>
  </el-card>

  <el-card :header="t('app.webui.jwtconfig')" class="section">
    <el-form :disabled="readonly" :model="data.jwtconf" label-width="150px" class="form">
      <el-form-item :label="t('app.webui.jwtkey')">
        <el-input v-model="data.jwtconf.jwt_secret" type="password" show-password />
      </el-form-item>
      <el-form-item :label="t('app.webui.jwtvalidity')">
        <el-input v-model="data.jwtconf.jwt_expires" type="number" />
      </el-form-item>
    </el-form>
  </el-card>

  <el-card :header="t('app.webui.noticeconfig')" class="section">
    <el-form :disabled="readonly" :model="data.noticeconf" label-width="150px" class="form">
      <el-form-item :label="t('app.webui.noticetype')">
        <el-select v-model="data.noticeconf.type">
          <el-option v-for="item in noticeTypes" :key="item.key" :label="item.value" :value="item.key" />
        </el-select>
      </el-form-item>
      <el-form-item label="Webhook">
        <el-input v-model="data.noticeconf.webhook" />
      </el-form-item>
      <el-form-item v-if="data.noticeconf.type == 1" label="Secret">
        <el-input v-model="data.noticeconf.secret" type="password" show-password />
      </el-form-item>
    </el-form>
  </el-card>

  <div class="actions">
    <el-button v-if="readonly && canUpdate" type="primary" @click="readonly = false">{{ t('app.webui.edit') }}</el-button>
    <template v-else>
      <el-button type="primary" @click="save">{{ t('app.webui.save') }}</el-button>
      <el-button @click="cancelEdit">{{ t('app.webui.cancel') }}</el-button>
    </template>
  </div>
</template>

<script lang="ts" setup>
import { onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import api from '../api'
import { loadBranding } from '../branding'
import { getUploadHeaders, hasPermission } from '../auth'

const { t } = useI18n()
const readonly = ref(true)
const uploadHeaders = getUploadHeaders()
const canUpdate = hasPermission('system.config.update')

const data = ref({
  emailconf: {
    email_host: '',
    email_port: 0,
    email_user: '',
    email_password: '',
    email_sender: '',
  },
  jwtconf: {
    jwt_secret: '',
    jwt_expires: 0,
  },
  noticeconf: {
    type: 0,
    secret: '',
    webhook: '',
  },
  sysconf: {
    user_register: false,
    user_display: '',
    max_attempts: 0,
    lockout_duration: 0,
    site_name: '',
    frontend_title: '',
    admin_title: '',
    logo_attachment_id: '',
    favicon_attachment_id: '',
    footer_text: '',
    help_url: '',
    suggest_url: '',
  },
})

const noticeTypes = ref([
  { key: 0, value: t('el.select.placeholder') },
  { key: 1, value: t('app.webui.dingtalk') },
  { key: 2, value: t('app.webui.wxwork') },
])

async function loadSystemConfig() {
  const response = await api.get('/api/v1/getsysconfig')
  if (response.data.code === 1) {
    data.value = response.data.data
  }
}

async function save() {
  const payload = {
    ...data.value,
    emailconf: {
      ...data.value.emailconf,
      email_port: Number(data.value.emailconf.email_port),
    },
    jwtconf: {
      ...data.value.jwtconf,
      jwt_expires: Number(data.value.jwtconf.jwt_expires),
    },
    noticeconf: {
      ...data.value.noticeconf,
      type: Number(data.value.noticeconf.type),
    },
    sysconf: {
      ...data.value.sysconf,
      max_attempts: Number(data.value.sysconf.max_attempts),
      lockout_duration: Number(data.value.sysconf.lockout_duration),
    },
  }
  const response = await api.post('/api/v1/updatesysconfig', payload)
  if (response.data.code === 1) {
    ElMessage.success(t('app.webui.savesuccess'))
    readonly.value = true
    await Promise.all([loadSystemConfig(), loadBranding(true)])
    return
  }
  ElMessage.error(t('app.webui.savefail'))
}

async function cancelEdit() {
  readonly.value = true
  await loadSystemConfig()
}

function handleLogoUpload(response: any) {
  if (response?.code !== 1 || !response?.file_id) {
    ElMessage.error(t('app.webui.savefail'))
    return
  }
  data.value.sysconf.logo_attachment_id = response.file_id
}

function handleFaviconUpload(response: any) {
  if (response?.code !== 1 || !response?.file_id) {
    ElMessage.error(t('app.webui.savefail'))
    return
  }
  data.value.sysconf.favicon_attachment_id = response.file_id
}

onMounted(loadSystemConfig)
</script>

<style scoped>
.section {
  margin-top: 16px;
}

.form {
  max-width: 960px;
}

.actions {
  margin-top: 16px;
}

.upload-row {
  display: flex;
  gap: 12px;
  width: 100%;
}
</style>
