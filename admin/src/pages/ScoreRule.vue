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
                    <el-button :disabled="multideleteVisible" type="danger" @click="multiDeleteScoreRules">{{ t('app.webui.multidelete') }}</el-button>
                </div>
            </div>
            <el-table :data="currentData" @selection-change="handleSelectionChange">
                <el-table-column type="selection" width="55" />
                <el-table-column prop="id" :label="t('app.webui.id')" sortable width="80" />
                <el-table-column
                :label="t('app.webui.type')"
                :filters="[
                        { text: t('app.webui.pocscorerule'), value: 1 },
                        { text: t('app.webui.exprule'), value: 2 },
                        { text: t('app.webui.incidencerule'), value: 3 },
                        { text: t('app.webui.otherrule'), value: 4 }
                    ]"
                :filter-method="typefilterHandler"
                width="140">
                    <template #default="{ row }">
                        <span v-if="row.type == 1">{{ t('app.webui.pocscorerule') }}</span>
                        <span v-else-if="row.type == 2">{{ t('app.webui.exprule') }}</span>
                        <span v-else-if="row.type == 3">{{ t('app.webui.incidencerule') }}</span>
                        <span v-else>{{ t('app.webui.otherrule') }}</span>
                    </template>
                </el-table-column>
                <el-table-column prop="rule" :label="t('app.webui.scorerule')" />
                <el-table-column prop="score" :label="t('app.webui.score')" width="120" />
                <el-table-column prop="coefficient" :label="t('app.webui.coefficient')" width="120" />
                <el-table-column :label="t('app.webui.createtime')" sortable :sort-method="sortDates" width="140">
                    <template #default="{ row }">
                        <span>{{ formatDate(row.create_time) }}</span>
                    </template>
                </el-table-column>
                <el-table-column :label="t('app.webui.updatetime')" sortable :sort-method="sortDates" width="140">
                    <template #default="{ row }">
                        <span>{{ formatDate(row.update_time) }}</span>
                    </template>
                </el-table-column>
                <el-table-column :label="t('app.webui.operate')" width="140">
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
    <el-dialog :title="t('app.webui.addscore')" v-model="dialogVisibleAdd" width="30%" @close="resetForm">
        <el-form :model="scoreruleForm" label-width="100px">
            <el-form-item :label="t('app.webui.type')">
                <el-select v-model="scoreruleForm.type" :placeholder="t('el.select.placeholder')">
                    <el-option v-for="item in typelist" :key="item.id" :label="item.type" :value="item.id" />
                </el-select>
            </el-form-item>
            <el-form-item :label="t('app.webui.scorerule')">
                <el-input v-model="scoreruleForm.rule" type="textarea" autocomplete="off" />
            </el-form-item>
            <el-form-item :label="t('app.webui.score')">
                <el-input v-model="scoreruleForm.score" type="number" step="0.1" autocomplete="off" />
            </el-form-item>
            <el-form-item :label="t('app.webui.coefficient')">
                <el-input v-model="scoreruleForm.coefficient" type="number" step="0.1" autocomplete="off" />
            </el-form-item>
        </el-form>
        <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisibleAdd = false">{{ t('app.webui.cancel') }}</el-button>
        <el-button type="primary" @click="addScorerule">{{ t('app.webui.confirm') }}</el-button>
        </span>
    </el-dialog>
    <el-dialog :title="t('app.webui.editscore')" v-model="dialogVisibleEdit" width="30%" @close="resetForm">
        <el-form :model="scoreruleForm" label-width="100px">
            <el-form-item :label="t('app.webui.type')">
                <el-select v-model="scoreruleForm.type" :placeholder="t('el.select.placeholder')">
                    <el-option v-for="item in typelist" :key="item.id" :label="item.type" :value="item.id" />
                </el-select>
            </el-form-item>
            <el-form-item :label="t('app.webui.scorerule')">
                <el-input v-model="scoreruleForm.rule" type="textarea" autocomplete="off" />
            </el-form-item>
            <el-form-item :label="t('app.webui.score')">
                <el-input v-model="scoreruleForm.score" type="number" step="0.1" autocomplete="off" />
            </el-form-item>
            <el-form-item :label="t('app.webui.coefficient')">
                <el-input v-model="scoreruleForm.coefficient" type="number" step="0.1" autocomplete="off" />
            </el-form-item>
        </el-form>
        <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisibleEdit = false">{{ t('app.webui.cancel') }}</el-button>
        <el-button type="primary" @click="editScorerule">{{ t('app.webui.confirm') }}</el-button>
        </span>
    </el-dialog>
</template>
<script lang="ts" setup>
import { ref, computed, onMounted } from 'vue'
import { useI18n } from 'vue-i18n';
import { formatDate } from '../utils'
import api from '../api'

const { t } = useI18n()
const token = sessionStorage.getItem('token')
const data = ref({})
const search = ref('')
const mountedFunctions = [fetchScoreRules]
const currentPage = ref(1);
const pageSize = ref(15);
const totalItems = ref(0)
const dialogVisibleAdd = ref(false)
const dialogVisibleEdit = ref(false)
const multipleSelection = ref([])
const scoreruleForm = ref({})
const typelist = ref([
    {"id": 1, "type": t('app.webui.pocscorerule')},
    {"id": 2, "type": t('app.webui.exprule')},
    {"id": 3, "type": t('app.webui.incidencerule')},
    {"id": 4, "type": t('app.webui.otherrule')}
])
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
    scoreruleForm.value.id = row.id
    scoreruleForm.value.type = row.type
    scoreruleForm.value.rule = row.rule
    scoreruleForm.value.score = row.score
    scoreruleForm.value.coefficient = row.coefficient
    dialogVisibleEdit.value = true
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
  fetchScoreRules();
}

// 处理当前页变化
async function handleCurrentChange(page: number) {
    currentPage.value = page;
    await fetchScoreRules();
}

onMounted(() => {
  mountedFunctions.forEach(fn => {
    fn();
  });
});

async function fetchScoreRules() {
    try {
        const config = {
            headers: {
                'Authorization': `Bearer ${token}`  // 使用Bearer schema
            }
        };
        const response = await api.get(`/api/v1/getscorerules?page=${currentPage.value}&limit=${pageSize.value}`, config)
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

function multiDeleteScoreRules() {
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
            const response = await api.post('/api/v1/multidelscorerules', data, config)
            if (response.data.code == 1) {
                ElMessage.success(t('app.webui.delsuccess'));
                fetchScoreRules()
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
            }
            const response = await api.post("/api/v1/delscorerule", {id: row.id}, config)
            if (response.data.code == 1) {
                ElMessage({
                    type: 'success',
                    message: t('app.webui.delsuccess'),
                })
                fetchScoreRules()
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
    
const typefilterHandler = (value, row) => {
    return row.type === value
}
const resetForm = () => {
    scoreruleForm.value = {
        "id": null,
        "type": 1,
        "rule": '',
        "score": null,
        "coefficient": 1
    }
}
const addScorerule = async () => {
    //console.log(scoreruleForm.value)
    try {
        const config = {
            headers: {
                'Authorization': `Bearer ${token}`
            }
        }
        scoreruleForm.value.score = Number(scoreruleForm.value.score)
        scoreruleForm.value.coefficient = Number(scoreruleForm.value.coefficient)
        const response = await api.post('/api/v1/addscorerule', scoreruleForm.value, config)
        if (response.data.code == 1) {
            fetchScoreRules()
            ElMessage.success(t('app.webui.addsuccess'));
            dialogVisibleAdd.value = false
        } else if (response.data.code == 2) {
            ElMessage.error(t('app.webui.invalidinput'))
            return
        }  else if (response.data.code == 3) {
            ElMessage.error(t('app.webui.addfail'))
            return
        } else {
            // 返回登录界面
        }
    } catch (error) {
        // 处理请求错误
        //ElMessage.error(t('app.webui.loginerr2'));
    }
}
const editScorerule = async () => {
    //console.log(scoreruleForm.value)
    try {
        const config = {
            headers: {
                'Authorization': `Bearer ${token}`
            }
        }
        scoreruleForm.value.score = Number(scoreruleForm.value.score)
        scoreruleForm.value.coefficient = Number(scoreruleForm.value.coefficient)
        const response = await api.post("/api/v1/editscorerule", scoreruleForm.value, config)
        //console.log(response)
        if (response.data.code == 1) {
            fetchScoreRules()
            ElMessage.success(t('app.webui.addsuccess'));
            dialogVisibleEdit.value = false
            return
        } else if (response.data.code == 2) {
            ElMessage.error(t('app.webui.invalidinput'))
            return
        }  else if (response.data.code == 3) {
            ElMessage.error(t('app.webui.modifyfail'))
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