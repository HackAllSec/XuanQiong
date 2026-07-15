<template>
  <div class="ops-layout">
    <el-card v-if="canReadAPIKeys || canManageAPIKeys" :header="t('app.webui.apikeymanager')">
      <div class="toolbar">
        <el-input v-model="apiKeyName" :placeholder="t('app.webui.apikeyname')" class="toolbar-input" />
        <el-date-picker
          v-model="apiKeyExpiresAt"
          type="datetime"
          :placeholder="t('app.webui.expiresat')"
          value-format="YYYY-MM-DDTHH:mm:ssZ"
          class="toolbar-input"
        />
        <el-button v-if="canManageAPIKeys" type="primary" @click="createAPIKey">{{ t('app.webui.add') }}</el-button>
      </div>
      <el-alert
        v-if="generatedKey"
        type="warning"
        show-icon
        :closable="false"
        :title="t('app.webui.apikeyshownonce')"
        class="key-alert"
      />
      <el-input v-if="generatedKey" v-model="generatedKey" readonly class="key-alert" />
      <el-table :data="apiKeys" border>
        <el-table-column prop="name" :label="t('app.webui.name')" />
        <el-table-column prop="key_prefix" :label="t('app.webui.apikeyprefix')" width="160" />
        <el-table-column :label="t('app.webui.expiresat')" width="180">
          <template #default="{ row }">{{ row.expires_at ? formatDate(row.expires_at) : '-' }}</template>
        </el-table-column>
        <el-table-column :label="t('app.webui.lastusedat')" width="180">
          <template #default="{ row }">{{ row.last_used_at ? formatDate(row.last_used_at) : '-' }}</template>
        </el-table-column>
        <el-table-column v-if="canManageAPIKeys" :label="t('app.webui.operate')" width="100">
          <template #default="{ row }">
            <el-button link type="danger" @click="deleteAPIKey(row.id)">{{ t('app.webui.delete') }}</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <el-card v-if="canImportVulns || canExportVulns" :header="t('app.webui.vulnimportexport')">
      <div class="action-row">
        <el-button v-if="canExportVulns" type="primary" @click="downloadFile('/api/v1/exportvulns', 'xuanqiong_vulns.csv')">
          {{ t('app.webui.exportvulns') }}
        </el-button>
        <el-upload
          v-if="canImportVulns"
          action="/api/v1/importvulns"
          accept=".csv"
          :headers="uploadHeaders"
          :show-file-list="false"
          :on-success="handleVulnImport"
        >
          <el-button>{{ t('app.webui.importvulns') }}</el-button>
        </el-upload>
      </div>
      <el-alert :title="t('app.webui.csvnotice')" type="info" show-icon :closable="false" />
    </el-card>

    <el-card v-if="canManageBackup" :header="t('app.webui.backuprestore')">
      <div class="action-row">
        <el-button type="primary" @click="downloadFile('/api/v1/exportbackup', 'xuanqiong_backup.json')">
          {{ t('app.webui.exportbackup') }}
        </el-button>
        <el-upload
          action="/api/v1/restorebackup"
          accept=".json"
          :headers="uploadHeaders"
          :show-file-list="false"
          :on-success="handleRestore"
        >
          <el-button type="warning">{{ t('app.webui.restorebackup') }}</el-button>
        </el-upload>
      </div>
      <el-alert :title="t('app.webui.restorenotice')" type="warning" show-icon :closable="false" />
    </el-card>
  </div>
</template>

<script lang="ts" setup>
import { computed, onMounted, ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useI18n } from 'vue-i18n'
import api from '../api'
import { getUploadHeaders, hasPermission } from '../auth'
import { formatDate } from '../utils'

const { t } = useI18n()

const apiKeys = ref<any[]>([])
const apiKeyName = ref('')
const apiKeyExpiresAt = ref('')
const generatedKey = ref('')
const uploadHeaders = getUploadHeaders()

const canReadAPIKeys = computed(() => hasPermission('api.key.read'))
const canManageAPIKeys = computed(() => hasPermission('api.key.manage'))
const canImportVulns = computed(() => hasPermission('vuln.import'))
const canExportVulns = computed(() => hasPermission('vuln.export'))
const canManageBackup = computed(() => hasPermission('backup.manage'))

async function loadAPIKeys() {
  if (!canReadAPIKeys.value) {
    return
  }
  const response = await api.get('/api/v1/apikeys')
  apiKeys.value = response.data.data || []
}

async function createAPIKey() {
  const response = await api.post('/api/v1/addapikey', {
    name: apiKeyName.value,
    expires_at: apiKeyExpiresAt.value,
  })
  if (response.data.code === 1) {
    generatedKey.value = response.data.api_key
    apiKeyName.value = ''
    apiKeyExpiresAt.value = ''
    ElMessage.success(t('app.webui.addsuccess'))
    await loadAPIKeys()
    return
  }
  ElMessage.error(response.data.msg || t('app.webui.addfail'))
}

async function deleteAPIKey(id: number) {
  await ElMessageBox.confirm(t('app.webui.confirmdelete'), t('app.webui.deletenotice'), { type: 'warning' })
  const response = await api.post('/api/v1/delapikey', { id })
  if (response.data.code === 1) {
    ElMessage.success(t('app.webui.delsuccess'))
    await loadAPIKeys()
    return
  }
  ElMessage.error(response.data.msg || t('app.webui.delfail'))
}

async function downloadFile(url: string, fallbackName: string) {
  const response = await api.get(url, { responseType: 'blob' })
  const disposition = response.headers['content-disposition'] || ''
  const filename = disposition.match(/filename=([^;]+)/)?.[1] || fallbackName
  const blobUrl = URL.createObjectURL(response.data)
  const link = document.createElement('a')
  link.href = blobUrl
  link.download = filename
  link.click()
  URL.revokeObjectURL(blobUrl)
}

function handleVulnImport(response: any) {
  if (response?.code === 1) {
    ElMessage.success(`${t('app.webui.imported')}: ${response.imported || 0}`)
    return
  }
  ElMessage.error(response?.msg || t('app.webui.modifyfail'))
}

function handleRestore(response: any) {
  if (response?.code === 1) {
    ElMessage.success(t('app.webui.restoresuccess'))
    return
  }
  ElMessage.error(response?.msg || t('app.webui.restorefail'))
}

onMounted(loadAPIKeys)
</script>

<style scoped>
.ops-layout {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.toolbar,
.action-row {
  display: flex;
  gap: 12px;
  align-items: center;
  margin-bottom: 16px;
}

.toolbar-input {
  max-width: 260px;
}

.key-alert {
  margin-bottom: 12px;
}
</style>
