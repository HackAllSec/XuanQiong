<template>
    <div>
        <el-card style="height: 100%; font-size: 20px; font-weight: bold;" shadow="always">
        <div style="font-weight: lighter; font-size: 17px;">
            <div style="display: flex; justify-content: space-around;">
                <div style="display: flex; width: 80%; gap: 1%;">
                    <el-input v-model="search" :placeholder="t('app.webui.search')" clearable style="width: 30%;" />
                </div>
                <div>
                    <el-button type="primary" @click="dialogVisibleAdd = true">{{ t('app.webui.add') }}</el-button>
                    <el-button :disabled="multideleteVisible" type="danger" @click="multiDeleteUsers">{{ t('app.webui.multidelete') }}</el-button>
                </div>
            </div>
            <el-table :data="currentData" @selection-change="handleSelectionChange">
                <el-table-column type="selection" width="55" />
                <el-table-column prop="id" :label="t('app.webui.id')" sortable />
                <el-table-column prop="username" :label="t('app.webui.username')" />
                <el-table-column
                    :filters="[
                        { text: t('app.webui.admin'), value: 1 },
                        { text: t('app.webui.user'), value: 0 }
                    ]"
                    :filter-method="rolefilterHandler"
                    :label="t('app.webui.role')">
                    <template #default="{ row }">
                        <span v-if="row.role != 0">{{ t('app.webui.admin') }}</span>
                        <span v-else>{{ t('app.webui.user') }}</span>
                    </template>
                </el-table-column>
                <el-table-column prop="email" :label="t('app.webui.email')" />
                <el-table-column prop="phone" :label="t('app.webui.phone')" />
                <el-table-column
                    :filters="[
                        { text: t('app.webui.enable'), value: 1 },
                        { text: t('app.webui.disable'), value: 0 }
                    ]"
                    :filter-method="statusfilterHandler"
                    :label="t('app.webui.status')">
                    <template #default="{ row }">
                        <div class="status">
                            <el-tag v-if="row.status != 0" type="success" effect="dark">{{ t('app.webui.enable') }}</el-tag>
                            <el-tag v-else type="info" effect="dark">{{ t('app.webui.disable') }}</el-tag>
                        </div>
                    </template>
                </el-table-column>
                <el-table-column :label="t('app.webui.createtime')"sortable>
                    <template #default="{ row }">
                        <span>{{ formatDate(row.create_time) }}</span>
                    </template>
                </el-table-column>
                <el-table-column :label="t('app.webui.operate')">
                    <template #default="scope">
                        <el-button size="small" type="primary" @click="handleEdit(scope.$index, scope.row)">{{ t('app.webui.edit') }}</el-button>
                        <el-button size="small" type="danger" @click="handleDelete(scope.$index, scope.row)">{{ t('app.webui.delete') }}</el-button>
                    </template>
                </el-table-column>
            </el-table>
        </div>
        <el-pagination
            :page-sizes="[15, 25, 50, 100, 200, 300]"
            v-model:page-size="pageSize"
            layout="total, sizes, prev, pager, next, jumper"
            :total="totalItems"
            @size-change="handleSizeChange"
            @current-change="handleCurrentChange"
            style="bottom: auto; margin-top: 2%; margin-left: 60%;"
            hide-on-single-page
        />
        </el-card>
    </div>
    <el-dialog :title="t('app.webui.adduser')" v-model="dialogVisibleAdd" width="30%" @close="resetForm">
        <el-form :model="userForm" label-width="100px">
            <el-form-item required :label="t('app.webui.role')">
                <el-select v-model="userForm.role" :placeholder="t('el.select.placeholder')">
                    <el-option v-for="item in rolelist" :key="item.id" :label="item.label" :value="item.id" />
                </el-select>
            </el-form-item>
            <el-form-item required :label="t('app.webui.username')">
                <el-input v-model="userForm.username" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item required :label="t('app.webui.password')">
                <el-input type="password" show-password v-model="userForm.password"></el-input>
            </el-form-item>
            <el-form-item required :label="t('app.webui.confirmpassword')">
                <el-input type="password" show-password v-model="confirmpassword"></el-input>
            </el-form-item>
            <el-form-item :label="t('app.webui.email')">
                <el-input v-model="userForm.email"></el-input>
            </el-form-item>
            <el-form-item :label="t('app.webui.phone')">
                <el-input v-model="userForm.phone"></el-input>
            </el-form-item>
        </el-form>
        <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisibleAdd = false">{{ t('app.webui.cancel') }}</el-button>
        <el-button type="primary" @click="addUser">{{ t('app.webui.confirm') }}</el-button>
        </span>
    </el-dialog>
    <el-dialog :title="t('app.webui.edituser')" v-model="dialogVisibleEdit" width="30%" @close="resetForm">
        <el-form :model="userForm" label-width="100px">
            <el-form-item :label="t('app.webui.role')">
                <el-select v-model="userForm.role">
                    <el-option v-for="item in rolelist" :key="item.id" :label="item.label" :value="item.id" />
                </el-select>
            </el-form-item>
            <el-form-item :label="t('app.webui.username')">
                <el-input v-model="userForm.username" autocomplete="off"></el-input>
            </el-form-item>
            <el-form-item :label="t('app.webui.password')">
                <el-input type="password" show-password v-model="userForm.password"></el-input>
            </el-form-item>
            <el-form-item :label="t('app.webui.confirmpassword')">
                <el-input type="password" show-password v-model="confirmpassword"></el-input>
            </el-form-item>
            <el-form-item :label="t('app.webui.email')">
                <el-input v-model="userForm.email"></el-input>
            </el-form-item>
            <el-form-item :label="t('app.webui.phone')">
                <el-input v-model="userForm.phone"></el-input>
            </el-form-item>
            <el-form-item :label="t('app.webui.status')">
                <el-switch v-model="userForm.status"
                    style="--el-switch-on-color: #13ce66;"/>
            </el-form-item>
        </el-form>
        <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisibleEdit = false">{{ t('app.webui.cancel') }}</el-button>
        <el-button type="primary" @click="editUser">{{ t('app.webui.confirm') }}</el-button>
        </span>
  </el-dialog>
</template>
<script lang="ts" setup>
import { ref, computed, onMounted } from 'vue'
import { useI18n } from 'vue-i18n';
import { formatDate } from '../utils'
import api from '../api'

const { t } = useI18n()
const data = ref({})
const rolelist = ref([
    { id: 1, label: t('app.webui.admin') },
    { id: 0, label: t('app.webui.user') }
])
const search = ref('')
const token = sessionStorage.getItem('token')
const dialogVisibleAdd = ref(false)
const dialogVisibleEdit = ref(false)
const userForm = ref({})
const mountedFunctions = [getUsers]
const currentPage = ref(1);
const pageSize = ref(15);
const totalItems = ref(0)
const multipleSelection = ref([])
const multideleteVisible = ref(true)
const confirmpassword = ref('')

const handleSelectionChange = (val) => {
    multipleSelection.value = val
    if (val.length > 0) {
        multideleteVisible.value = false
    } else {
        multideleteVisible.value = true
    }
}
// 计算当前页的数据
const currentData = computed(() => {
    const start = 0;
    const end = start + pageSize.value;
    //console.log(start, end)
    //console.log(data.value)
    if (search.value.trim() != '') {
        // 过滤数据
        return data.value.data.filter(item => {
            return (
                item.username.toLowerCase().includes(search.value.toLowerCase()) ||
                item.email.toLowerCase().includes(search.value.toLowerCase()) ||
                item.phone.includes(search.value)
            );
        });
    }
    try {
        return data.value.data.slice(start, end);
    } catch (error) {
        return [];
    }
});

// 处理每页条目数变化
function handleSizeChange(size: number) {
  pageSize.value = size;
  currentPage.value = 1; // 每次改变条目数时重置到第一页
  getUsers();
}

// 处理当前页变化
async function handleCurrentChange(page: number) {
    currentPage.value = page;
    await getUsers();
}

onMounted(() => {
  mountedFunctions.forEach(fn => {
    fn();
  });
});

async function getUsers() {
    try {
        const config = {
            headers: {
                'Authorization': `Bearer ${token}`  // 使用Bearer schema
            }
        };
        const response = await api.get(`/api/v1/getusers?page=${currentPage.value}&limit=${pageSize.value}`, config)
        if (response.data.code != 1) {
            // 清空token，返回登录页
            sessionStorage.removeItem("token")
            sessionStorage.removeItem("username")

        }
        data.value = response.data
        totalItems.value = response.data.total
    } catch (error) {
        // 处理请求错误
        //ElMessage.error(t('app.webui.loginerr2'));
    }
}
function multiDeleteUsers() {
    ElMessageBox.confirm(
        t('app.webui.deletenotice'),
        t('app.webui.confirmdelete'),
        {
            confirmButtonText: t('app.webui.confirm'),
            cancelButtonText: t('app.webui.cancel'),
            type: 'warning',
        }
    )
    .then(async () => {
        const data = {
            "ids": multipleSelection.value.map(item => item.id)
        }
        try {
            const config = {
                headers: {
                    'Authorization': `Bearer ${token}`
                }
            };
            const response = await api.post('/api/v1/multidelusers', data, config)
            if (response.data.code == 1) {
                ElMessage.success(t('app.webui.delsuccess'));
                getUsers()
                multideleteVisible.value = false
            } else if (response.data.code == 0) {
                ElMessage.error(t('app.webui.needlogin'));
            } else {
                ElMessage.error(t('app.webui.deletefail'));
            }
        } catch (error) {
            
        }
    })
    .catch(() => {
      
    })
}

const handleEdit = (index, row) => {
    userForm.value.id = row.id
    userForm.value.role = row.role
    userForm.value.username = row.username
    userForm.value.password = row.password
    userForm.value.email = row.email
    userForm.value.phone = row.phone
    userForm.value.status = row.status == 1
    dialogVisibleEdit.value = true
}
const rolefilterHandler = (value, row) => {
    return row.role === value
}
const statusfilterHandler = (value, row) => {
    return row.status === value
}
const handleDelete = (index, row) => {
    ElMessageBox.confirm(
        t('app.webui.deletenotice'),
        t('app.webui.confirmdelete'),
        {
            confirmButtonText: t('app.webui.confirm'),
            cancelButtonText: t('app.webui.cancel'),
            type: 'warning',
        }
    )
    .then(async () => {
        try {
            const config = {
                headers: {
                    'Authorization': `Bearer ${token}`
                }
            };
            const response = await api.post('/api/v1/deluser', {userid: row.id}, config)
            if (response.data.code == 1) {
                getUsers()
                ElMessage.success(t('app.webui.deletecomplete'));
            } else if (response.data.code == 0) {
                ElMessage.error(t('app.webui.needlogin'));
            } else {
                ElMessage.error(t('app.webui.deletefail'));
            }
        } catch (error) {
            
        }
    })
    .catch(() => {
      
    })
}
const resetForm = () => {
    userForm.value.id = null
    userForm.value.role = null
    userForm.value.username = ''
    userForm.value.password = ''
    userForm.value.email = ''
    userForm.value.phone = ''
    userForm.value.status = null
    confirmpassword.value = ''
}
async function addUser () {
    //console.log(userForm.value)
    if (userForm.value.role == null) {
        ElMessage.error(t('app.webui.roleisnull'));
        return
    }
    if (userForm.value.username == '') {
        ElMessage.error(t('app.webui.usernameempty'));
        return
    }
    if (userForm.value.password == '') {
        ElMessage.error(t('app.webui.passwordempty'));
        return
    }
    if (userForm.value.password != confirmpassword.value) {
        ElMessage.error(t('app.webui.passwordnotmatch'));
        return
    }
    try {
        const config = {
            headers: {
                'Authorization': `Bearer ${token}`
            }
        };
        const response = await api.post('/api/v1/adduser', userForm.value, config)
        if (response.data.code == 1) {
            getUsers()
            ElMessage.success(t('app.webui.addsuccess'));
            dialogVisibleAdd.value = false
            //console.log(userForm.value)
        } else if (response.data.code == 0) {
            ElMessage.error(t('app.webui.needlogin'));
        } else {
            ElMessage.error(t('app.webui.usernamealreadyexist'));
        }
    } catch (error) {
        // 处理请求错误
        //ElMessage.error(t('app.webui.loginerr2'));
    }
}
async function editUser() {
    //console.log(userForm.value)
    if (userForm.value.password == '') {
        if (userForm.value.password != confirmpassword.value) {
            ElMessage.error(t('app.webui.passwordnotmatch'));
            return
        }
    }
    userForm.value.status = userForm.value.status == true ? 1 : 0
    try {
        const config = {
            headers: {
                'Authorization': `Bearer ${token}`
            }
        };
        const response = await api.post('/api/v1/updateuser', userForm.value, config)
        if (response.data.code == 1) {
            getUsers()
            ElMessage.success(t('app.webui.modifysucc'));
        } else if (response.data.code == 0) {
            ElMessage.error(t('app.webui.needlogin'));
        } else {
            ElMessage.error(t('app.webui.modifyfail'));
        }
    } catch (error) {
        
    }
    dialogVisibleEdit.value = false
}
</script>
<style scoped>
    .status {
  display: flex;
  gap: 4px;
}
</style>