<template>
  <el-card :header="t('app.webui.auditlogs')">
    <div class="toolbar">
      <el-input v-model="keyword" :placeholder="t('app.webui.search')" clearable class="toolbar-input" />
      <el-input v-model="action" :placeholder="t('app.webui.auditaction')" clearable class="toolbar-input" />
      <el-button type="primary" @click="loadLogs">{{ t('app.webui.search') }}</el-button>
    </div>

    <el-table :data="logs" border>
      <el-table-column prop="username" :label="t('app.webui.username')" width="120" />
      <el-table-column prop="action" :label="t('app.webui.auditaction')" width="180" />
      <el-table-column prop="method" label="Method" width="100" />
      <el-table-column prop="path" label="Path" min-width="220" />
      <el-table-column prop="client_ip" label="IP" width="140" />
      <el-table-column prop="result_message" :label="t('app.webui.result')" min-width="180" />
      <el-table-column :label="t('app.webui.createtime')" width="180">
        <template #default="{ row }">
          {{ formatDate(row.create_time) }}
        </template>
      </el-table-column>
      <el-table-column :label="t('app.webui.detail')" width="100">
        <template #default="{ row }">
          <el-button link type="primary" @click="showDetail(row)">{{ t('app.webui.detail') }}</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
      v-model:current-page="currentPage"
      v-model:page-size="pageSize"
      :total="total"
      layout="total, prev, pager, next"
      class="pagination"
      @current-change="loadLogs"
    />
  </el-card>

  <el-dialog v-model="detailVisible" :title="t('app.webui.auditdetail')" width="900px">
    <el-descriptions :column="1" border v-if="currentLog">
      <el-descriptions-item :label="t('app.webui.username')">{{ currentLog.username || '-' }}</el-descriptions-item>
      <el-descriptions-item :label="t('app.webui.auditaction')">{{ currentLog.action }}</el-descriptions-item>
      <el-descriptions-item label="Path">{{ currentLog.path }}</el-descriptions-item>
      <el-descriptions-item label="IP">{{ currentLog.client_ip }}</el-descriptions-item>
      <el-descriptions-item :label="t('app.webui.requestbody')">
        <pre class="payload-block">{{ currentLog.request_body || '-' }}</pre>
      </el-descriptions-item>
      <el-descriptions-item :label="t('app.webui.responsebody')">
        <pre class="payload-block">{{ currentLog.response_body || '-' }}</pre>
      </el-descriptions-item>
    </el-descriptions>
  </el-dialog>
</template>

<script lang="ts" setup>
import { onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import api from '../api'
import { formatDate } from '../utils'

const { t } = useI18n()

const keyword = ref('')
const action = ref('')
const logs = ref<any[]>([])
const total = ref(0)
const currentPage = ref(1)
const pageSize = ref(20)
const detailVisible = ref(false)
const currentLog = ref<any | null>(null)

async function loadLogs() {
  const response = await api.get('/api/v1/getauditlogs', {
    params: {
      page: currentPage.value,
      limit: pageSize.value,
      keyword: keyword.value,
      action: action.value,
    },
  })
  logs.value = response.data.data || []
  total.value = response.data.total || 0
}

function showDetail(log: any) {
  currentLog.value = log
  detailVisible.value = true
}

onMounted(loadLogs)
</script>

<style scoped>
.toolbar {
  display: flex;
  gap: 12px;
  margin-bottom: 16px;
}

.toolbar-input {
  max-width: 260px;
}

.pagination {
  margin-top: 16px;
  justify-content: flex-end;
}

.payload-block {
  margin: 0;
  white-space: pre-wrap;
  word-break: break-all;
  font-size: 12px;
}
</style>
