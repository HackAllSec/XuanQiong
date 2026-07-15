<template>
  <el-card shadow="always" class="user-card">
    <div class="toolbar">
      <el-input v-model="search" :placeholder="t('app.webui.search')" clearable class="search-input" />
      <div class="toolbar-actions">
        <el-button v-if="canCreate" type="primary" @click="openCreateDialog">{{ t('app.webui.add') }}</el-button>
        <el-button v-if="canDelete" :disabled="selectedRows.length === 0" type="danger" @click="multiDeleteUsers">
          {{ t('app.webui.multidelete') }}
        </el-button>
      </div>
    </div>

    <el-table :data="filteredUsers" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" />
      <el-table-column prop="id" :label="t('app.webui.id')" width="80" />
      <el-table-column prop="username" :label="t('app.webui.username')" min-width="140" />
      <el-table-column :label="t('app.webui.role')" min-width="220">
        <template #default="{ row }">
          <el-tag v-for="role in row.roles || []" :key="role" class="role-tag" size="small">{{ role }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="email" :label="t('app.webui.email')" min-width="180" />
      <el-table-column prop="phone" :label="t('app.webui.phone')" min-width="140" />
      <el-table-column :label="t('app.webui.status')" width="100">
        <template #default="{ row }">
          <el-tag :type="row.status === 1 ? 'success' : 'info'">
            {{ row.status === 1 ? t('app.webui.enable') : t('app.webui.disable') }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column :label="t('app.webui.createtime')" min-width="140">
        <template #default="{ row }">
          {{ formatDate(row.create_time) }}
        </template>
      </el-table-column>
      <el-table-column :label="t('app.webui.operate')" width="180">
        <template #default="{ row }">
          <el-button v-if="canUpdate" size="small" type="primary" @click="openEditDialog(row)">{{ t('app.webui.edit') }}</el-button>
          <el-button v-if="canDelete" size="small" type="danger" @click="deleteUser(row)">{{ t('app.webui.delete') }}</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
      v-model:current-page="currentPage"
      v-model:page-size="pageSize"
      :page-sizes="[15, 25, 50, 100]"
      :total="totalItems"
      layout="total, sizes, prev, pager, next, jumper"
      class="pagination"
      @size-change="loadUsers"
      @current-change="loadUsers"
    />
  </el-card>

  <el-dialog v-model="dialogVisible" :title="isEditing ? t('app.webui.edituser') : t('app.webui.adduser')" width="520px" @close="resetForm">
    <el-form :model="userForm" label-width="110px">
      <el-form-item :label="t('app.webui.role')">
        <el-select v-model="userForm.role_ids" multiple filterable style="width: 100%">
          <el-option v-for="role in roleList" :key="role.id" :label="role.name" :value="role.id" />
        </el-select>
      </el-form-item>
      <el-form-item :label="t('app.webui.username')">
        <el-input v-model="userForm.username" />
      </el-form-item>
      <el-form-item :label="t('app.webui.password')">
        <el-input v-model="userForm.password" type="password" show-password />
      </el-form-item>
      <el-form-item :label="t('app.webui.confirmpassword')">
        <el-input v-model="confirmPassword" type="password" show-password />
      </el-form-item>
      <el-form-item :label="t('app.webui.email')">
        <el-input v-model="userForm.email" />
      </el-form-item>
      <el-form-item :label="t('app.webui.phone')">
        <el-input v-model="userForm.phone" />
      </el-form-item>
      <el-form-item v-if="isEditing" :label="t('app.webui.status')">
        <el-switch v-model="userForm.enabled" />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button @click="dialogVisible = false">{{ t('app.webui.cancel') }}</el-button>
      <el-button type="primary" @click="submitUser">{{ t('app.webui.confirm') }}</el-button>
    </template>
  </el-dialog>
</template>

<script lang="ts" setup>
import { computed, onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { formatDate } from '../utils'
import api from '../api'
import { hasPermission } from '../auth'

const { t } = useI18n()
const canCreate = hasPermission('user.create')
const canUpdate = hasPermission('user.update')
const canDelete = hasPermission('user.delete')

const search = ref('')
const users = ref<any[]>([])
const roleList = ref<any[]>([])
const currentPage = ref(1)
const pageSize = ref(15)
const totalItems = ref(0)
const selectedRows = ref<any[]>([])
const dialogVisible = ref(false)
const isEditing = ref(false)
const confirmPassword = ref('')
const userForm = ref(createUserForm())

function createUserForm() {
  return {
    id: 0,
    username: '',
    password: '',
    email: '',
    phone: '',
    role_ids: [],
    enabled: true,
  }
}

const filteredUsers = computed(() => {
  const term = search.value.trim().toLowerCase()
  if (!term) {
    return users.value
  }
  return users.value.filter((item) =>
    [item.username, item.email, item.phone, (item.roles || []).join(' ')].some((field) =>
      String(field || '').toLowerCase().includes(term)
    )
  )
})

function handleSelectionChange(rows: any[]) {
  selectedRows.value = rows
}

async function loadUsers() {
  const response = await api.get('/api/v1/getusers', {
    params: { page: currentPage.value, limit: pageSize.value },
  })
  users.value = response.data.data || []
  totalItems.value = response.data.total || 0
}

async function loadRoles() {
  const response = await api.get('/api/v1/getroles')
  roleList.value = response.data.data || []
}

function openCreateDialog() {
  isEditing.value = false
  resetForm()
  dialogVisible.value = true
}

function openEditDialog(row: any) {
  isEditing.value = true
  userForm.value = {
    id: row.id,
    username: row.username,
    password: '',
    email: row.email,
    phone: row.phone,
    role_ids: [...(row.role_ids || [])],
    enabled: row.status === 1,
  }
  confirmPassword.value = ''
  dialogVisible.value = true
}

function resetForm() {
  userForm.value = createUserForm()
  confirmPassword.value = ''
}

async function submitUser() {
  if (!userForm.value.username) {
    ElMessage.error(t('app.webui.usernameempty'))
    return
  }
  if (!isEditing.value && !userForm.value.password) {
    ElMessage.error(t('app.webui.passwordempty'))
    return
  }
  if (userForm.value.password && userForm.value.password !== confirmPassword.value) {
    ElMessage.error(t('app.webui.passwordnotmatch'))
    return
  }
  const payload: any = {
    id: userForm.value.id,
    username: userForm.value.username,
    password: userForm.value.password,
    email: userForm.value.email,
    phone: userForm.value.phone,
    role_ids: userForm.value.role_ids,
    status: userForm.value.enabled ? 1 : 0,
  }
  const url = isEditing.value ? '/api/v1/updateuser' : '/api/v1/adduser'
  const response = await api.post(url, payload)
  if (response.data.code === 1) {
    ElMessage.success(isEditing.value ? t('app.webui.modifysucc') : t('app.webui.addsuccess'))
    dialogVisible.value = false
    await loadUsers()
    return
  }
  ElMessage.error(response.data.msg || t('app.webui.modifyfail'))
}

async function deleteUser(row: any) {
  try {
    await ElMessageBox.confirm(t('app.webui.deletenotice'), t('app.webui.confirmdelete'), {
      type: 'warning',
      confirmButtonText: t('app.webui.confirm'),
      cancelButtonText: t('app.webui.cancel'),
    })
    const response = await api.post('/api/v1/deluser', { userid: row.id })
    if (response.data.code === 1) {
      ElMessage.success(t('app.webui.delsuccess'))
      await loadUsers()
      return
    }
    ElMessage.error(response.data.msg || t('app.webui.delfail'))
  } catch {
    return
  }
}

async function multiDeleteUsers() {
  try {
    await ElMessageBox.confirm(t('app.webui.deletenotice'), t('app.webui.confirmdelete'), {
      type: 'warning',
      confirmButtonText: t('app.webui.confirm'),
      cancelButtonText: t('app.webui.cancel'),
    })
    const response = await api.post('/api/v1/multidelusers', {
      ids: selectedRows.value.map((item) => item.id),
    })
    if (response.data.code === 1) {
      ElMessage.success(t('app.webui.delsuccess'))
      selectedRows.value = []
      await loadUsers()
      return
    }
    ElMessage.error(response.data.msg || t('app.webui.delfail'))
  } catch {
    return
  }
}

onMounted(async () => {
  await Promise.all([loadUsers(), loadRoles()])
})
</script>

<style scoped>
.user-card {
  height: 100%;
}

.toolbar {
  display: flex;
  justify-content: space-between;
  gap: 16px;
  margin-bottom: 16px;
}

.toolbar-actions {
  display: flex;
  gap: 12px;
}

.search-input {
  max-width: 320px;
}

.role-tag {
  margin: 0 6px 6px 0;
}

.pagination {
  margin-top: 16px;
  justify-content: flex-end;
}
</style>
