<template>
  <div class="page">
    <el-card :header="t('app.webui.apikeymanager')">
      <div class="toolbar">
        <el-input v-model="name" :placeholder="t('app.webui.apikeyname')" class="toolbar-input" />
        <el-date-picker
          v-model="expiresAt"
          type="datetime"
          :placeholder="t('app.webui.expiresat')"
          value-format="YYYY-MM-DDTHH:mm:ssZ"
          class="toolbar-input"
        />
        <el-button type="primary" @click="createKey">{{ t('app.webui.add') }}</el-button>
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
        <el-table-column :label="t('app.webui.operate')" width="100">
          <template #default="{ row }">
            <el-button link type="danger" @click="deleteKey(row.id)">{{ t('app.webui.delete') }}</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<script lang="ts" setup>
import { onMounted, ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useI18n } from 'vue-i18n'
import api from '../api'
import { formatDate } from '../utils'

const { t } = useI18n()

const apiKeys = ref<any[]>([])
const name = ref('')
const expiresAt = ref('')
const generatedKey = ref('')

async function loadKeys() {
  const response = await api.get('/api/v1/apikeys')
  apiKeys.value = response.data.data || []
}

async function createKey() {
  const response = await api.post('/api/v1/addapikey', {
    name: name.value,
    expires_at: expiresAt.value,
  })
  if (response.data.code === 1) {
    generatedKey.value = response.data.api_key
    name.value = ''
    expiresAt.value = ''
    ElMessage.success(t('app.webui.addsuccess'))
    await loadKeys()
    return
  }
  ElMessage.error(response.data.msg || t('app.webui.addfail'))
}

async function deleteKey(id: number) {
  await ElMessageBox.confirm(t('app.webui.confirmdelete'), t('app.webui.delete'), { type: 'warning' })
  const response = await api.post('/api/v1/delapikey', { id })
  if (response.data.code === 1) {
    ElMessage.success(t('app.webui.delsuccess'))
    await loadKeys()
    return
  }
  ElMessage.error(response.data.msg || t('app.webui.delfail'))
}

onMounted(loadKeys)
</script>

<style scoped>
.page {
  width: 90%;
  margin: 24px auto;
}

.toolbar {
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
