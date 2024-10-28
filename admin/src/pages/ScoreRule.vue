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
                    <el-button :disabled="multideleteVisible" type="danger" @click="multiDeleteUser">{{ t('app.webui.multidelete') }}</el-button>
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
import { DocumentCopy } from '@element-plus/icons-vue'

const { t } = useI18n()
const data = ref({
    "data":[{
            "id": 1,
            "type": 1,
            "rule": "Xray、Nuclei、Goby等完整Poc，误报低",
            "score": 20,
            "coefficient": 1,
            "create_time": "2024-10-23T16:20:32.31+08:00",
            "update_time": "0001-01-01T00:00:00Z"
        },
        {
            "id": 2,
            "type": 1,
            "rule": "Xray、Nuclei、Goby等完整Poc，误报较高",
            "score": 10,
            "coefficient": 1,
            "create_time": "2024-10-23T16:20:32.31+08:00",
            "update_time": "0001-01-01T00:00:00Z"
        },
        {
            "id": 3,
            "type": 1,
            "rule": "仅包含Payload或无法工具化的Poc",
            "score": 5,
            "coefficient": 1,
            "create_time": "2024-10-23T16:20:32.31+08:00",
            "update_time": "0001-01-01T00:00:00Z"
        },
        {
            "id": 4,
            "type": 2,
            "rule": "Xray、Nuclei、Goby等完整Exp，误报低",
            "score": 30,
            "coefficient": 1,
            "create_time": "2024-10-23T16:20:32.31+08:00",
            "update_time": "0001-01-01T00:00:00Z"
        },
        {
            "id": 5,
            "type": 2,
            "rule": "Xray、Nuclei、Goby等完整Exp，误报较高",
            "score": 15,
            "coefficient": 1,
            "create_time": "2024-10-23T16:20:32.31+08:00",
            "update_time": "0001-01-01T00:00:00Z"
        },
        {
            "id": 6,
            "type": 2,
            "rule": "Xray、Nuclei、Goby等完整Exp，误报较高",
            "score": 5,
            "coefficient": 1,
            "create_time": "2024-10-23T16:20:32.31+08:00",
            "update_time": "0001-01-01T00:00:00Z"
        },
        {
            "id": 7,
            "type": 3,
            "rule": "互联网资产数大于 5000",
            "score": 30,
            "coefficient": 1,
            "create_time": "2024-10-23T16:20:32.31+08:00",
            "update_time": "0001-01-01T00:00:00Z"
        }
    ]})
const vulndetail = ref({})
const search = ref('')
const mountedFunctions = []//fetchVulnList]
const currentPage = ref(1);
const pageSize = ref(15);
const totalItems = ref(0)
const dialogVisibleAdd = ref(false)
const dialogVisibleEdit = ref(false)
const scoreruleForm = ref({})
const typelist = ref([
    {"id": 1, "type": t('app.webui.pocscorerule')},
    {"id": 2, "type": t('app.webui.exprule')},
    {"id": 3, "type": t('app.webui.incidencerule')},
    {"id": 4, "type": t('app.webui.otherrule')}
])
const multideleteVisible = ref(true)

const handleSelectionChange = (val) => {
    multideleteVisible.value = val.id
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
    console.log(data.value)
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
  fetchVulnList();
}

// 处理当前页变化
async function handleCurrentChange(page: number) {
    currentPage.value = page;
    await fetchVulnList();
}

onMounted(() => {
  mountedFunctions.forEach(fn => {
    fn();
  });
});

async function fetchVulnList() {
    try {
        const response = await api.get(`/api/v1/getvulnlist?page=${currentPage.value}&limit=${pageSize.value}`)
        data.value = response.data
        totalItems.value = response.data.total
        typefilter.value = response.data.data.reduce((acc, item) => {
            if (!acc.some(i => i.value === item.vuln_type)) {
                acc.push({ text: item.vuln_type, value: item.vuln_type });
            }
            return acc;
        }, []);
    } catch (error) {
        // 处理请求错误
        //ElMessage.error(t('app.webui.loginerr2'));
    }
}

function multiDeleteUser() {
    ElMessageBox.confirm(
    t('app.webui.deletenotice'),
    t('app.webui.confirmdelete'),
    {
      confirmButtonText: t('app.webui.confirm'),
      cancelButtonText: t('app.webui.cancel'),
      type: 'warning',
    }
  )
    .then(() => {
      ElMessage({
        type: 'success',
        message: t('app.webui.deletecomplete'),
      })
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
    .then(() => {
        ElMessage({
        type: 'success',
        message: t('app.webui.deletecomplete'),
      })
      console.log(index,row)
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
const addScorerule = () => {
    console.log(scoreruleForm.value)
    dialogVisibleAdd.value = false
}
const editScorerule = () => {
    console.log(scoreruleForm.value)
    dialogVisibleEdit.value = false
}
</script>
<style scoped>
    .status {
  display: flex;
  gap: 4px;
}
</style>