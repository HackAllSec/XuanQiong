<template>
  <el-card :header="t('app.webui.rolemanager')">
    <div class="toolbar">
      <el-input v-model="keyword" :placeholder="t('app.webui.search')" clearable class="search-input" />
      <el-button v-if="canCreate" type="primary" @click="openCreateDialog">{{ t('app.webui.addrole') }}</el-button>
    </div>

    <el-table :data="filteredRoles" border>
      <el-table-column prop="id" :label="t('app.webui.id')" width="80" />
      <el-table-column prop="name" :label="t('app.webui.name')" min-width="160" />
      <el-table-column prop="code" label="Code" min-width="180" />
      <el-table-column prop="description" :label="t('app.webui.description')" min-width="220" />
      <el-table-column :label="t('app.webui.status')" width="100">
        <template #default="{ row }">
          <el-tag :type="row.status === 1 ? 'success' : 'info'">
            {{ row.status === 1 ? t('app.webui.enable') : t('app.webui.disable') }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column :label="t('app.webui.permissions')" min-width="260">
        <template #default="{ row }">
          <el-tag
            v-for="permission in row.permission_codes || []"
            :key="permission"
            class="permission-tag"
            size="small"
          >
            {{ permission }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column :label="t('app.webui.operate')" width="200">
        <template #default="{ row }">
          <el-button v-if="canUpdate" type="primary" size="small" @click="openEditDialog(row)">{{ t('app.webui.edit') }}</el-button>
          <el-button
            v-if="canDelete"
            type="danger"
            size="small"
            :disabled="row.is_system"
            @click="deleteRole(row)"
          >
            {{ t('app.webui.delete') }}
          </el-button>
        </template>
      </el-table-column>
    </el-table>
  </el-card>

  <el-dialog v-model="dialogVisible" :title="isEditing ? t('app.webui.editrole') : t('app.webui.addrole')" width="720px">
    <el-form :model="roleForm" label-width="110px">
      <el-form-item :label="t('app.webui.name')">
        <el-input v-model="roleForm.name" />
      </el-form-item>
      <el-form-item label="Code">
        <el-input v-model="roleForm.code" :disabled="isEditing && roleForm.is_system" />
      </el-form-item>
      <el-form-item :label="t('app.webui.description')">
        <el-input v-model="roleForm.description" type="textarea" :rows="3" />
      </el-form-item>
      <el-form-item :label="t('app.webui.status')">
        <el-switch v-model="roleForm.enabled" />
      </el-form-item>
      <el-form-item :label="t('app.webui.permissions')">
        <div class="permission-groups">
          <div v-for="group in permissionGroups" :key="group.category" class="permission-group">
            <div class="permission-group-title">{{ group.category }}</div>
            <el-checkbox-group v-model="roleForm.permission_codes">
              <el-checkbox v-for="item in group.items" :key="item.code" :label="item.code">
                {{ item.name }}
              </el-checkbox>
            </el-checkbox-group>
          </div>
        </div>
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button @click="dialogVisible = false">{{ t('app.webui.cancel') }}</el-button>
      <el-button type="primary" @click="submitRole">{{ t('app.webui.confirm') }}</el-button>
    </template>
  </el-dialog>
</template>

<script lang="ts" setup>
import { computed, onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import api from '../api'
import { hasPermission } from '../auth'

const { t } = useI18n()

const keyword = ref('')
const roles = ref<any[]>([])
const permissions = ref<any[]>([])
const dialogVisible = ref(false)
const isEditing = ref(false)
const roleForm = ref<any>(createRoleForm())

const canCreate = hasPermission('role.create')
const canUpdate = hasPermission('role.update')
const canDelete = hasPermission('role.delete')

function createRoleForm() {
  return {
    id: 0,
    name: '',
    code: '',
    description: '',
    enabled: true,
    is_system: false,
    permission_codes: [],
  }
}

const filteredRoles = computed(() => {
  const term = keyword.value.trim().toLowerCase()
  if (!term) {
    return roles.value
  }
  return roles.value.filter((item) =>
    [item.name, item.code, item.description].some((field) => String(field || '').toLowerCase().includes(term))
  )
})

const permissionGroups = computed(() => {
  const groups: Record<string, any[]> = {}
  for (const permission of permissions.value) {
    const category = permission.category || 'default'
    if (!groups[category]) {
      groups[category] = []
    }
    groups[category].push(permission)
  }
  return Object.keys(groups).map((category) => ({
    category,
    items: groups[category],
  }))
})

async function loadRoles() {
  const response = await api.get('/api/v1/getroles')
  roles.value = response.data.data || []
}

async function loadPermissions() {
  const response = await api.get('/api/v1/getpermissions')
  permissions.value = response.data.data || []
}

function openCreateDialog() {
  isEditing.value = false
  roleForm.value = createRoleForm()
  dialogVisible.value = true
}

function openEditDialog(role: any) {
  isEditing.value = true
  roleForm.value = {
    id: role.id,
    name: role.name,
    code: role.code,
    description: role.description,
    enabled: role.status === 1,
    is_system: role.is_system,
    permission_codes: [...(role.permission_codes || [])],
  }
  dialogVisible.value = true
}

async function submitRole() {
  if (!roleForm.value.name || !roleForm.value.code) {
    ElMessage.error(t('app.webui.invalidinput'))
    return
  }
  const payload = {
    id: roleForm.value.id,
    name: roleForm.value.name,
    code: roleForm.value.code,
    description: roleForm.value.description,
    status: roleForm.value.enabled ? 1 : 0,
    permission_codes: roleForm.value.permission_codes,
  }
  const url = isEditing.value ? '/api/v1/updaterole' : '/api/v1/addrole'
  const response = await api.post(url, payload)
  if (response.data.code === 1) {
    ElMessage.success(isEditing.value ? t('app.webui.modifysucc') : t('app.webui.addsuccess'))
    dialogVisible.value = false
    await loadRoles()
    return
  }
  ElMessage.error(response.data.msg || t('app.webui.savefail'))
}

async function deleteRole(role: any) {
  try {
    await ElMessageBox.confirm(t('app.webui.deletenotice'), t('app.webui.confirmdelete'), {
      type: 'warning',
      confirmButtonText: t('app.webui.confirm'),
      cancelButtonText: t('app.webui.cancel'),
    })
    const response = await api.post('/api/v1/delrole', { id: role.id })
    if (response.data.code === 1) {
      ElMessage.success(t('app.webui.delsuccess'))
      await loadRoles()
      return
    }
    ElMessage.error(response.data.msg || t('app.webui.delfail'))
  } catch {
    return
  }
}

onMounted(async () => {
  await Promise.all([loadRoles(), loadPermissions()])
})
</script>

<style scoped>
.toolbar {
  display: flex;
  justify-content: space-between;
  gap: 16px;
  margin-bottom: 16px;
}

.search-input {
  max-width: 320px;
}

.permission-tag {
  margin: 0 6px 6px 0;
}

.permission-groups {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(260px, 1fr));
  gap: 16px;
  width: 100%;
}

.permission-group {
  border: 1px solid var(--el-border-color-light);
  border-radius: 8px;
  padding: 12px;
}

.permission-group-title {
  font-weight: 600;
  margin-bottom: 10px;
  text-transform: capitalize;
}
</style>
