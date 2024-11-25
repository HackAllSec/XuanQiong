<template>
    <div>
        <el-card style="height: 100%; font-size: 20px; font-weight: bold;" shadow="always">
        <div style="font-weight: lighter; font-size: 17px;">
            <div style="display: flex; justify-content: space-around;">
                <div style="display: flex; width: 80%; gap: 1%;">
                    <el-input v-model="search" :placeholder="t('app.webui.search')" clearable style="width: 30%;" />
                </div>
                <div>
                    <el-button type="primary" @click="dialogVisibleAdd=true">{{ t('app.webui.add') }}</el-button>
                    <el-button :disabled="multideleteVisible" type="danger" @click="multiDeleteVulnTypes">{{ t('app.webui.multidelete') }}</el-button>
                </div>
            </div>
            <el-table :data="currentData" @selection-change="handleSelectionChange">
                <el-table-column type="selection" width="55" />
                <el-table-column prop="id" :label="t('app.webui.id')" sortable />
                <el-table-column prop="name" :label="t('app.webui.name')" />
                <el-table-column :label="t('app.webui.createtime')" sortable :sort-method="sortDates">
                    <template #default="{ row }">
                        <span>{{ formatDate(row.create_time) }}</span>
                    </template>
                </el-table-column>
                <el-table-column :label="t('app.webui.updatetime')" sortable :sort-method="sortDates">
                    <template #default="{ row }">
                        <span>{{ formatDate(row.update_time) }}</span>
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
    <el-dialog :title="t('app.webui.addvulntype')" v-model="dialogVisibleAdd" width="30%" @close="resetForm">
        <el-form :model="vulntypeForm" label-width="100px">
            <el-form-item :label="t('app.webui.name')">
                <el-input v-model="vulntypeForm.name" autocomplete="off"></el-input>
            </el-form-item>
        </el-form>
        <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisibleAdd = false">{{ t('app.webui.cancel') }}</el-button>
        <el-button type="primary" @click="addVulnType">{{ t('app.webui.confirm') }}</el-button>
        </span>
    </el-dialog>
    <el-dialog :title="t('app.webui.editvulntype')" v-model="dialogVisibleEdit" width="30%" @close="resetForm">
        <el-form :model="vulntypeForm" label-width="100px">
            <el-form-item :label="t('app.webui.id')">
                <el-input v-model="vulntypeForm.id" autocomplete="off" disabled="true"></el-input>
            </el-form-item>
            <el-form-item :label="t('app.webui.name')">
                <el-input v-model="vulntypeForm.name" autocomplete="off"></el-input>
            </el-form-item>
        </el-form>
        <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisibleEdit = false">{{ t('app.webui.cancel') }}</el-button>
        <el-button type="primary" @click="editVulnType">{{ t('app.webui.confirm') }}</el-button>
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
const search = ref('')
const token = sessionStorage.getItem('token')
const mountedFunctions = [fetchVulnType]
const currentPage = ref(1);
const pageSize = ref(15);
const totalItems = ref(0)
const dialogVisibleAdd = ref(false)
const dialogVisibleEdit = ref(false)
const vulntypeForm = ref({})
const multipleSelection = ref([])
const multideleteVisible = ref(true)

const handleSelectionChange = (val) => {
    multipleSelection.value = val
    if (val.length > 0) {
        multideleteVisible.value = false
    } else {
        multideleteVisible.value = true
    }
}

const handleEdit = (index, row) => {
    vulntypeForm.value.id = row.id
    vulntypeForm.value.name = row.name
    dialogVisibleEdit.value = true
}
const resetForm = () => {
    vulntypeForm.value.id = null
    vulntypeForm.value.name = ''
}
function sortDates(a, b) {
    // 假设 create_time 和 update_time 是 Date 对象
    const dateA = new Date(a.create_time);
    const dateB = new Date(b.create_time);
    if (dateA < dateB) return -1;
    if (dateA > dateB) return 1;
    return 0;
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
                item.name.toLowerCase().includes(search.value.toLowerCase())
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
  fetchVulnType();
}

// 处理当前页变化
async function handleCurrentChange(page: number) {
    currentPage.value = page;
    await fetchVulnType();
}

onMounted(() => {
  mountedFunctions.forEach(fn => {
    fn();
  });
});

async function fetchVulnType() {
    //console.log(data.value)
    try {
        const response = await api.get(`/api/v1/getvulntype?page=${currentPage.value}&limit=${pageSize.value}`)
        data.value = response.data
        totalItems.value = response.data.total
        //console.log(data.value)
    } catch (error) {
        // 处理请求错误
        //ElMessage.error(t('app.webui.loginerr2'));
    }
}

function multiDeleteVulnTypes() {
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
                    Authorization: `Bearer ${token}`
                }
            };
            const response = await api.post('/api/v1/multidelvulntypes', data, config)
            if (response.data.code == 1) {
                ElMessage.success(t('app.webui.delsuccess'));
                fetchVulnType()
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
                    Authorization: `Bearer ${token}`
                }
            }
            const response = await api.post("/api/v1/delvulntype", {id: row.id}, config)
            if (response.data.code == 1) {
                ElMessage({
                    type: 'success',
                    message: t('app.webui.delsuccess'),
                })
                fetchVulnType()
                dialogVisibleAdd.value = false
                return
            } else if (response.data.code == 2) {
                ElMessage({
                    type: 'error',
                    message: t('app.webui.invalidinput'),
                })
                return
            }  else if (response.data.code == 3) {
                ElMessage({
                    type: 'error',
                    message: t('app.webui.delfail'),
                })
                return
            } else {
                // 返回登录界面
            }     
        } catch (error) {
            // 处理请求错误
            //ElMessage.error(t('app.webui.loginerr2'));
        }
    })
    .catch(() => {
      
    })
}
async function addVulnType() {
    //console.log(vulntypeForm.value)
    try {
        const config = {
            headers: {
                Authorization: `Bearer ${token}`
            }
        }
        const response = await api.post("/api/v1/addvulntype", vulntypeForm.value, config)
        //console.log(response)
        if (response.data.code == 1) {
            ElMessage({
                type: 'success',
                message: t('app.webui.addsuccess'),
            })
            fetchVulnType()
            dialogVisibleAdd.value = false
            return
        } else if (response.data.code == 2) {
            ElMessage({
                type: 'error',
                message: t('app.webui.invalidinput'),
            })
            return
        }  else if (response.data.code == 3) {
            ElMessage({
                type: 'error',
                message: t('app.webui.addfail'),
            })
            return
        } else {
            // 返回登录界面
        }
    } catch (error) {
        // 处理请求错误
        //ElMessage.error(t('app.webui.loginerr2'));
    }
}
const editVulnType = () => {
    //console.log(vulntypeForm.value)
    try {
        const config = {
            headers: {
                Authorization: `Bearer ${token}`
            }
        }
        const response = api.post("/api/v1/updatevulntype", vulntypeForm.value, config)
        //console.log(response)
        if (response.data.code == 1) {
            ElMessage({
                type: 'success',
                message: t('app.webui.modifysucc'),
            })
            fetchVulnType()
            dialogVisibleEdit.value = false
            return
        } else if (response.data.code == 2) {
            ElMessage({
                type: 'error',
                message: t('app.webui.invalidinput'),
            })
            return
        }  else if (response.data.code == 3) {
            ElMessage({
                type: 'error',
                message: t('app.webui.modifyfail'),
            })
            return
        } else {
            // 返回登录界面
        }
    } catch (error) {
        // 处理请求错误
        //ElMessage.error(t('app.webui.loginerr2'));
    }
}
</script>
<style scoped>
    .status {
  display: flex;
  gap: 4px;
}
</style>