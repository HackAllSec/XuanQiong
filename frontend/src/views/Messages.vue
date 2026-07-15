<template>
  <div class="page">
    <el-card :header="`${t('app.webui.messages')} (${t('app.webui.unread')}: ${unread})`">
      <template #header>
        <div class="card-header">
          <span>{{ t('app.webui.messages') }} ({{ t('app.webui.unread') }}: {{ unread }})</span>
          <el-button type="primary" @click="markAllRead">{{ t('app.webui.markallread') }}</el-button>
        </div>
      </template>
      <el-table :data="messages" border>
        <el-table-column prop="title" :label="t('app.webui.title')" min-width="180" />
        <el-table-column prop="content" :label="t('app.webui.content')" min-width="360" />
        <el-table-column :label="t('app.webui.status')" width="100">
          <template #default="{ row }">
            <el-tag :type="row.is_read ? 'info' : 'success'">
              {{ row.is_read ? t('app.webui.read') : t('app.webui.unread') }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column :label="t('app.webui.createtime')" width="180">
          <template #default="{ row }">{{ formatDate(row.create_time) }}</template>
        </el-table-column>
        <el-table-column :label="t('app.webui.operate')" width="110">
          <template #default="{ row }">
            <el-button v-if="!row.is_read" link type="primary" @click="markRead(row.id)">
              {{ t('app.webui.markread') }}
            </el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination
        v-model:current-page="currentPage"
        v-model:page-size="pageSize"
        :total="total"
        layout="total, prev, pager, next"
        class="pagination"
        @current-change="loadMessages"
      />
    </el-card>
  </div>
</template>

<script lang="ts" setup>
import { onMounted, ref } from 'vue'
import { ElMessage } from 'element-plus'
import { useI18n } from 'vue-i18n'
import api from '../api'
import { formatDate } from '../utils'

const { t } = useI18n()

const messages = ref<any[]>([])
const total = ref(0)
const unread = ref(0)
const currentPage = ref(1)
const pageSize = ref(15)

async function loadMessages() {
  const response = await api.get('/api/v1/messages', {
    params: { page: currentPage.value, limit: pageSize.value },
  })
  messages.value = response.data.data || []
  total.value = response.data.total || 0
  unread.value = response.data.unread || 0
}

async function markRead(id: number) {
  const response = await api.post('/api/v1/readmessage', { id })
  if (response.data.code === 1) {
    await loadMessages()
    return
  }
  ElMessage.error(response.data.msg || t('app.webui.modifyfail'))
}

async function markAllRead() {
  const response = await api.post('/api/v1/readallmessages')
  if (response.data.code === 1) {
    await loadMessages()
    return
  }
  ElMessage.error(response.data.msg || t('app.webui.modifyfail'))
}

onMounted(loadMessages)
</script>

<style scoped>
.page {
  width: 90%;
  margin: 24px auto;
}

.card-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.pagination {
  margin-top: 16px;
  justify-content: flex-end;
}
</style>
