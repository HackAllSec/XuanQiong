<template>
    <div v-if="showedit">
        <el-card shadow="always" :header="t('app.webui.editvuln')">
        <div>
            <el-form inline="true" :model="form" label-width="auto" size="large">
            <el-form-item :label="t('app.webui.vulnname')" prop="vuln_name" style="width: 45%" :rules="[{ required: true, message: t('app.webui.required') },]">
                <el-input v-model="form.vuln_name" />
            </el-form-item>
            <el-form-item :label="t('app.webui.vulntype')" style="width: 45%" required>
                <el-select v-model="form.vuln_type_id" :placeholder="t('el.select.placeholder')">
                    <el-option v-for="item in vulntype" :key="item.id" :label="item.name" :value="item.id" />
                </el-select>
            </el-form-item>
            <el-form-item :label="t('app.webui.affectedproduct')" prop="affected_product" style="width: 45%" :rules="[{ required: true, message: t('app.webui.required') },]">
                <el-input v-model="form.affected_product" />
            </el-form-item>
            <el-form-item :label="t('app.webui.productversion')" style="width: 45%" prop="affected_product_version" :rules="[{ required: true, message: t('app.webui.required') },]">
                <el-input v-model="form.affected_product_version" />
            </el-form-item>
            <el-form-item label="CVSS" style="width: 29%" required>
                <el-input v-model="form.cvss" type="number" step="0.1" @change="ShowVlunLevel" />
            </el-form-item>
            <el-form-item :label="t('app.webui.vulnlevel')" style="width: 29%" required>
                <el-input v-model="vulnlevel" readonly disabled />
            </el-form-item>
            <el-form-item :label="t('app.webui.ispublic')" style="width: 29%">
                <el-switch v-model="form.is_public" />
            </el-form-item>
            <el-form-item :label="t('app.webui.cveid')" style="width: 29%">
                <el-input v-model="form.cve" />
            </el-form-item>
            <el-form-item :label="t('app.webui.nvdid')" style="width: 29%">
                <el-input v-model="form.nvd" />
            </el-form-item>
            <el-form-item :label="t('app.webui.edbid')" style="width: 29%">
                <el-input v-model="form.edbid" />
            </el-form-item>
            <el-form-item :label="t('app.webui.cnnvdid')" style="width: 29%">
                <el-input v-model="form.cnnvd" />
            </el-form-item>
            <el-form-item :label="t('app.webui.cnvdid')" style="width: 29%">
                <el-input v-model="form.cnvd" />
            </el-form-item>
            <el-form-item :label="`Fofa ${t('app.webui.searchquery')}`" style="width: 45%">
                <el-input v-model="form.fofa_query" />
            </el-form-item>
            <el-form-item :label="`ZoomEye ${t('app.webui.searchquery')}`" style="width: 45%">
                <el-input v-model="form.zoom_eye_query" />
            </el-form-item>
            <el-form-item :label="`Quake ${t('app.webui.searchquery')}`" style="width: 45%">
                <el-input v-model="form.quake_query" />
            </el-form-item>
            <el-form-item :label="`Hunter ${t('app.webui.searchquery')}`" style="width: 45%">
                <el-input v-model="form.hunter_query" />
            </el-form-item>
            <el-form-item :label="`Google ${t('app.webui.searchquery')}`" style="width: 45%">
                <el-input v-model="form.google_query" />
            </el-form-item>
            <el-form-item :label="`Shodan ${t('app.webui.searchquery')}`" style="width: 45%">
                <el-input v-model="form.shodan_query" />
            </el-form-item>
            <el-form-item :label="`Censys ${t('app.webui.searchquery')}`" style="width: 45%">
                <el-input v-model="form.censys_query" />
            </el-form-item>
            <el-form-item :label="`Greynoise ${t('app.webui.searchquery')}`" style="width: 45%">
                <el-input v-model="form.greynoise_query" />
            </el-form-item>
            <el-form-item :label="t('app.webui.vulndesc')" style="width: 90%" prop="description" :rules="[{ required: true, message: t('app.webui.required') },]">
                <el-input v-model="form.description" type="textarea" autosize />
            </el-form-item>
            
            <el-form-item :label="t('app.webui.vulnsuggest')" style="width: 90%" prop="repair_suggestion" :rules="[{ required: true, message: t('app.webui.required') },]">
                <el-input v-model="form.repair_suggestion" type="textarea" autosize />
            </el-form-item>
            <el-form-item label="Poc" style="width: 90%" @change="ShowPocInput">
                <el-radio-group v-model="poc" size="large">
                    <el-radio-button label="Xray" value="xray" />
                    <el-radio-button label="Nuclei" value="nuclei" />
                    <el-radio-button label="Goby" value="goby" />
                    <el-radio-button :label="t('app.webui.other')" value="other" />
                </el-radio-group>
                <el-input v-model="form.poc" type="textarea" :placeholder="pocplaceholder" autosize style="margin-top: 2%;" />
            </el-form-item>
            <el-form-item label="Exp" style="width: 90%"  @change="ShowExpInput">
                <el-radio-group v-model="exp" size="large">
                    <el-radio-button label="Xray" value="xray" />
                    <el-radio-button label="Nuclei" value="nuclei" />
                    <el-radio-button label="Goby" value="goby" />
                    <el-radio-button :label="t('app.webui.other')" value="other" />
                </el-radio-group>
                <el-input v-model="form.exp" type="textarea" :placeholder="expplaceholder" autosize style="margin-top: 2%;" />
            </el-form-item>
            
            <el-form-item :label="t('app.webui.attachfile')" style="width: 100%">
                <el-upload class="upload-demo" drag accept=".zip,.doc,.docx,.pdf,.txt" action="/api/v1/upload" :headers="{'Authorization': `Bearer ${token}`}" :on-success="handleSuccess" :on-remove="handleRemove" style="width: 90%">
                    <el-icon class="el-icon--upload"><upload-filled /></el-icon>
                    <div class="el-upload__text">
                        {{ t('app.webui.draguplaod') }} <em>{{ t('app.webui.clickupload') }}</em>
                    </div>
                    <template #tip>
                    <div class="el-upload__tip">
                        {{  t('app.webui.uploadnotice1') }}
                    </div>
                    </template>
                </el-upload>
            </el-form-item>
            <el-form-item style="width: 100%;">
                <el-button @click="showedit=false" size="default">{{ t('app.webui.cancel') }}</el-button>
                <el-button type="primary" @click="onSubmit" size="default">{{ t('app.webui.confirm') }}</el-button>
            </el-form-item>
            </el-form>
        </div>
    </el-card>
    </div>
    <div v-else-if="showvulndetail">
        <el-card shadow="always">
            <template #header>
                <div style="display: flex; justify-content: space-between; align-items: center;">
                    <span style="display: flex; align-items: center;">
                        <el-icon size="25" style="margin-right: 10px;">
                            <svg t="1727430038190" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="21943" width="200" height="200"><path d="M940 512H792V412c76.8 0 139-62.2 139-139 0-4.4-3.6-8-8-8h-60c-4.4 0-8 3.6-8 8 0 34.8-28.2 63-63 63H232c-34.8 0-63-28.2-63-63 0-4.4-3.6-8-8-8h-60c-4.4 0-8 3.6-8 8 0 76.8 62.2 139 139 139v100H84c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h148v96c0 6.5 0.2 13 0.7 19.3C164.1 728.6 116 796.7 116 876c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8 0-44.2 23.9-82.9 59.6-103.7 6 17.2 13.6 33.6 22.7 49 24.3 41.5 59 76.2 100.5 100.5S460.5 960 512 960s99.8-13.9 141.3-38.2c41.5-24.3 76.2-59 100.5-100.5 9.1-15.5 16.7-31.9 22.7-49C812.1 793.1 836 831.8 836 876c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8 0-79.3-48.1-147.4-116.7-176.7 0.4-6.4 0.7-12.8 0.7-19.3v-96h148c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8zM716 680c0 36.8-9.7 72-27.8 102.9-17.7 30.3-43 55.6-73.3 73.3-20.1 11.8-42 20-64.9 24.3V484c0-4.4-3.6-8-8-8h-60c-4.4 0-8 3.6-8 8v396.5c-22.9-4.3-44.8-12.5-64.9-24.3-30.3-17.7-55.6-43-73.3-73.3C317.7 752 308 716.8 308 680V412h408v268z" p-id="21944" fill="#d81e06"></path><path d="M304 280h56c4.4 0 8-3.6 8-8 0-28.3 5.9-53.2 17.1-73.5 10.6-19.4 26-34.8 45.4-45.4C450.9 142 475.7 136 504 136h16c28.3 0 53.2 5.9 73.5 17.1 19.4 10.6 34.8 26 45.4 45.4C650 218.9 656 243.7 656 272c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8 0-40-8.8-76.7-25.9-108.1-17.2-31.5-42.5-56.8-74-74C596.7 72.8 560 64 520 64h-16c-40 0-76.7 8.8-108.1 25.9-31.5 17.2-56.8 42.5-74 74C304.8 195.3 296 232 296 272c0 4.4 3.6 8 8 8z" p-id="21945" fill="#d81e06"></path></svg>
                        </el-icon>
                        {{ vulndetail.data.vuln_name }}
                    </span>
                    <el-button type="primary" @click="goBack">{{ t('el.pageHeader.title') }}</el-button>
                </div>
            </template>
            <div >
                <el-descriptions :column="5">
                    <el-descriptions-item :label="`${t('app.webui.id')}:`">
                        <el-tag>{{vulndetail.data.id}}</el-tag>
                    </el-descriptions-item>
                    <el-descriptions-item :label="`${t('app.webui.type')}:`">{{vulndetail.data.vuln_type}}</el-descriptions-item>
                    <el-descriptions-item :label="`${t('app.webui.level')}:`">
                        <el-tag v-if="vulndetail.data.vuln_level === 'Critical'" type="danger" color="#CD0000" effect="dark" size="small">{{ t('app.webui.critical') }}</el-tag>
                        <el-tag v-else-if="vulndetail.data.vuln_level === 'High'" type="danger" color="#EE2C2C" effect="dark" size="small">{{ t('app.webui.high') }}</el-tag>
                        <el-tag v-else-if="vulndetail.data.vuln_level === 'Medium'" type="warning" color="#FF6600" effect="dark" size="small">{{ t('app.webui.medium') }}</el-tag>
                        <el-tag v-else type="primary" effect="dark" size="small">{{ t('app.webui.low') }}</el-tag>
                    </el-descriptions-item>
                    <el-descriptions-item :label="`${t('app.webui.createtime')}:`">{{formatDate(vulndetail.data.create_time)}}</el-descriptions-item>
                    <el-descriptions-item :label="`${t('app.webui.submitter')}:`">{{vulndetail.data.submitter}}</el-descriptions-item>
                    <el-descriptions-item :label="`${t('app.webui.cveid')}:`"><el-tag>{{vulndetail.data.cve}}</el-tag></el-descriptions-item>
                    <el-descriptions-item :label="`${t('app.webui.nvdid')}:`"><el-tag>{{vulndetail.data.nvd}}</el-tag></el-descriptions-item>
                    <el-descriptions-item :label="`${t('app.webui.edbid')}:`"><el-tag>{{vulndetail.data.edb}}</el-tag></el-descriptions-item>
                    <el-descriptions-item :label="`${t('app.webui.cnnvdid')}:`"><el-tag>{{vulndetail.data.cnnvd}}</el-tag></el-descriptions-item>
                    <el-descriptions-item :label="`${t('app.webui.cnvdid')}:`"><el-tag>{{vulndetail.data.cnvd}}</el-tag></el-descriptions-item>
                    <el-descriptions-item span="2" :label="`${t('app.webui.affectedproduct')}:`">{{vulndetail.data.affected_product}}</el-descriptions-item>
                    <el-descriptions-item span="2" :label="`${t('app.webui.productversion')}:`">{{vulndetail.data.affected_product_version}}</el-descriptions-item>
                    <el-descriptions-item label="CVSS:">{{vulndetail.data.cvss}}</el-descriptions-item>
                    <el-descriptions-item span="5" :label="`${t('app.webui.vulndesc')}:`">
                        <div style="margin-top: 10px; padding-inline: 25px; white-space: pre-wrap;">{{vulndetail.data.description}}</div>
                    </el-descriptions-item>
                    <el-descriptions-item span="5" :label="`${t('app.webui.vulnsuggest')}:`">
                        <div style="margin-top: 10px; padding-inline: 25px; white-space: pre-wrap;">{{vulndetail.data.repair_suggestion}}</div>
                    </el-descriptions-item>
                    <el-descriptions-item v-if="vulndetail.data.attachment_id" span="5" :label="`${t('app.webui.attachfile')}:`">
                        <div style="margin-top: 10px; padding-inline: 25px;">
                            <a :href="`/download/file?id=${vulndetail.data.attachment_id}`">{{ vulndetail.data.attachment_name }}</a>
                        </div>
                    </el-descriptions-item>
                </el-descriptions>
            </div>
        </el-card>
        <el-card style="margin-top: 20px;" shadow="always" :header="t('app.webui.searchquery')">
            <div>
                <el-descriptions :column="5">
                    <el-descriptions-item span="2" :label="`Fofa ${t('app.webui.searchquery')}:`">{{vulndetail.data.fofa_query}}</el-descriptions-item>
                    <el-descriptions-item span="2" :label="`ZoomEye ${t('app.webui.searchquery')}:`">{{vulndetail.data.zoom_eye_query}}</el-descriptions-item>
                    <el-descriptions-item :label="`Quake ${t('app.webui.searchquery')}:`">{{vulndetail.data.quake_query}}</el-descriptions-item>
                    <el-descriptions-item span="2" :label="`Hunter ${t('app.webui.searchquery')}:`">{{vulndetail.data.hunter_query}}</el-descriptions-item>
                    <el-descriptions-item span="2" :label="`Google ${t('app.webui.searchquery')}:`">{{vulndetail.data.google_query}}</el-descriptions-item>
                    <el-descriptions-item :label="`Shodan ${t('app.webui.searchquery')}:`">{{vulndetail.data.shodan_query}}</el-descriptions-item>
                    <el-descriptions-item span="2" :label="`Censys ${t('app.webui.searchquery')}:`">{{vulndetail.data.censys_query}}</el-descriptions-item>
                    <el-descriptions-item :label="`Greynoise ${t('app.webui.searchquery')}:`">{{vulndetail.data.greynoise_query}}</el-descriptions-item>
                    
                </el-descriptions>
            </div>
        </el-card>
        <el-card style="margin-top: 20px;" shadow="always" header="Poc">
            <div v-if="vulndetail.data.poc != ''" style="position: relative;">
                <el-button type="primary" :icon="DocumentCopy" circle style="position: absolute; right: 10px; top: 10px; z-index: 1000;" @click="copyToClipboard(vulndetail.data.poc)" />
                <el-input v-model="vulndetail.data.poc" type="textarea" autosize readonly input-style="background-color: #515151; color: #fff;" />
            </div>
            <el-empty v-else style="height: 30vh;" :description="t('app.webui.empty')" />
        </el-card>
        <el-card style="margin-top: 20px;" shadow="always" header="Exp">
            <div v-if="vulndetail.data.exp != ''" style="position: relative;">
                <el-button type="primary" :icon="DocumentCopy" circle style="position: absolute; right: 10px; top: 10px; z-index: 1000;" @click="copyToClipboard(vulndetail.data.exp)" />
                <el-input v-model="vulndetail.data.exp" type="textarea" autosize readonly input-style="background-color: #515151; color: #fff;" />
            </div>
            <el-empty v-else style="height: 30vh;" :description="t('app.webui.empty')" />
        </el-card>
    </div>
    <div v-else>
        <el-card style="height: 100%; font-size: 20px; font-weight: bold;" shadow="always">
        <div style="font-weight: lighter; font-size: 17px;">
            <div style="display: flex; justify-content: space-around;">
                <div style="display: flex; width: 80%; gap: 1%;">
                    <el-input v-model="search" :placeholder="t('app.webui.search')" clearable style="width: 30%;" />
                </div>
                <el-button :disabled="multideleteVisible" type="danger" @click="multiDeleteUser">{{ t('app.webui.multidelete') }}</el-button>
            </div>
            <el-table :data="currentData" @selection-change="handleSelectionChange" @cell-mouse-enter="cellMouseEnter" @cell-mouse-leave="cellMouseLeave" @cell-click="cellClick">
                <el-table-column type="selection" width="55" />
                <el-table-column prop="id" :label="t('app.webui.id')" sortable width="180" />
                <el-table-column prop="vuln_name" :label="t('app.webui.name')" />
                <el-table-column prop="vuln_type" :filters="typefilter" :filter-method="typefilterHandler" :label="t('app.webui.type')" width="180" />
                <el-table-column
                    prop="vuln_level"
                    :filters="[
                        { text: t('app.webui.critical'), value: 'Critical' },
                        { text: t('app.webui.high'), value: 'High' },
                        { text: t('app.webui.medium'), value: 'Medium' },
                        { text: t('app.webui.low'), value: 'Low' }
                    ]"
                    :filter-method="levelfilterHandler"
                    :label="t('app.webui.level')"
                    width="85" >
                    <template #default="{ row }">
                        <el-tag v-if="row.vuln_level === 'Critical'" type="danger" effect="dark" color="#CD0000">{{ t('app.webui.critical') }}</el-tag>
                        <el-tag v-else-if="row.vuln_level === 'High'" type="danger" effect="dark" color="#EE2C2C">{{ t('app.webui.high') }}</el-tag>
                        <el-tag v-else-if="row.vuln_level === 'Medium'" type="warning" effect="dark" color="#FF6600">{{ t('app.webui.medium') }}</el-tag>
                        <el-tag v-else type="primary" effect="dark">{{ t('app.webui.low') }}</el-tag>
                    </template>
                </el-table-column>
                <el-table-column prop="cvss" label="CVSS" sortable width="85" />
                <el-table-column
                    :filters="[
                        { text: 'Poc', value: 'Poc' },
                        { text: 'Exp', value: 'Exp' }
                    ]"
                    :filter-method="statusfilterHandler"
                    label="Poc&Exp"
                    width="120">
                    <template #default="{ row }">
                        <div class="status">
                            <el-tag v-if="row.poc != '' && row.poc != '0'" type="success" effect="dark">Poc</el-tag>
                            <el-tag v-else type="info" effect="dark">Poc</el-tag>
                            <el-tag v-if="row.exp != '' && row.exp != '0'" type="success" effect="dark">Exp</el-tag>
                            <el-tag v-else type="info" effect="dark">Exp</el-tag>
                        </div>
                    </template>
                </el-table-column>
                <el-table-column :label="t('app.webui.ispublic')" sortable :sort-method="sortIsPublic" width="110">
                    <template #default="{ row }">
                        <span v-if="row.is_public">{{ t('app.webui.yes') }}</span>
                        <span v-else>{{ t('app.webui.no') }}</span>
                    </template>
                </el-table-column>
                <el-table-column :label="t('app.webui.createtime')" width="135" sortable :sort-method="sortDates">
                    <template #default="{ row }">
                        <span>{{ formatDate(row.create_time) }}</span>
                    </template>
                </el-table-column>
                <el-table-column :label="t('app.webui.updatetime')" width="135" sortable :sort-method="sortDates">
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
    
</template>
<script lang="ts" setup>
import { ref, computed, onMounted } from 'vue'
import { useI18n } from 'vue-i18n';
import { formatDate } from '../utils'
import api from '../api'
import { DocumentCopy, UploadFilled } from '@element-plus/icons-vue'
import { useRoute } from 'vue-router'


const { t } = useI18n()
const poc = ref('xray')
const exp = ref('xray')
const vulntype = ref([])
const vulnlevel = ref('--')
const data = ref({})
const vulndetail = ref({})
const typefilter = ref([])
const search = ref('')
const mountedFunctions = [getVulnTypes, fetchVulnList]
const currentPage = ref(1);
const pageSize = ref(15);
const totalItems = ref(0)
const multideleteVisible = ref(true)
const showvulndetail = ref(false)
const showedit = ref(false)
const form = ref({
        vuln_name: '',
        vuln_type_id: 1,
        cvss: 0.1,
        vuln_level: '',
        cve: '',
        nvd: '',
        edb: '',
        cnnvd: '',
        cnvd: '',
        affected_product: '',
        repair_suggestion: '',
        poc: '',
        exp: '',
        poc_type: poc.value,
        exp_type: exp.value,
        is_public: false,
        description: '',
        fofa_query: '',
        zoom_eye_query: '',
        quake_query: '',
        hunter_query: '',
        google_query: '',
        shodan_query: '',
        censys_query: '',
        greynoise_query: '',
        attachment_id: '',
        affected_product_version: '',
    })

    const ShowVlunLevel = () => {
        form.value.cvss = Number(form.value.cvss)
        if (form.value.cvss > 0 && form.value.cvss < 4) {
            form.value.vuln_level = 'Low'
            vulnlevel.value = t('app.webui.low')
            return
        } else if (form.value.cvss >= 4 && form.value.cvss < 7) {
            form.value.vuln_level = 'Medium'
            vulnlevel.value = t('app.webui.medium')
            return
        } else if (form.value.cvss >= 7 && form.value.cvss < 9) {
            form.value.vuln_level = 'High'
            vulnlevel.value = t('app.webui.high')
            return
        } else if (form.value.cvss >= 9 && form.value.cvss <= 10) {
            form.value.vuln_level = 'Critical'
            vulnlevel.value = t('app.webui.critical')
            return
        } else {
            ElMessage.error(t('app.webui.cvsserror'))
        }
    }
    
    async function getVulnTypes() {
        try {
            const response = await api.get('/api/v1/getvulntypes')
            vulntype.value = response.data.data
        } catch (error) {
            //console.error(error)
        }
    }
const handleSelectionChange = (val) => {
    multideleteVisible.value = val.id
    if (val.length > 0) {
        multideleteVisible.value = false
    } else {
        multideleteVisible.value = true
    }
}
function sortIsPublic(a, b) {
    if (a.is_public && !b.is_public) return -1;
    if (!a.is_public && b.is_public) return 1;
    return 0;
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
    if (search.value.trim() != '') {
        // 过滤数据
        return data.value.data.filter(item => {
            return (
                item.id.toLowerCase().includes(search.value.toLowerCase()) ||
                item.vuln_name.toLowerCase().includes(search.value.toLowerCase())
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

const cellMouseEnter = (row, column, cell, event) => {
    if (column.property === 'id' || column.property === 'vuln_name') {
        event.target.style.cursor = 'pointer'
        event.target.style.color = '#409EFF'
    }
}
const cellMouseLeave = (row, column, cell, event) => {
    if (column.property === 'id' || column.property === 'vuln_name') {
        event.target.style.cursor = ''
        event.target.style.color = ''
    }
}
const goBack = () => {
    showvulndetail.value = false
    vulndetail.value = {}
}

async function getVulnDetail(id) {
    const token = sessionStorage.getItem('token')
        try {
            const config = {
                headers: {
                    'Authorization': `Bearer ${token}`  // 使用Bearer schema
                }
            };
            const response = await api.get('/api/v1/getvulndtl?id=' + id, config)
            if (token && response.data.code == 0) {
                sessionStorage.removeItem('token')
                sessionStorage.removeItem('username')
                sessionStorage.removeItem('avatar')
                location.reload()
            }
            vulndetail.value = response.data
        } catch (error) {
            // 处理请求错误
            //ElMessage.error(t('app.webui.loginerr2'));
        }
}
const cellClick = (row, cell) => {
    if (cell.no == 1 || cell.no == 2) {
        getVulnDetail(row.id)
        showvulndetail.value = true
    }
}

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
const handleEdit = async (index, row) => {
    console.log(row.id)
    await getVulnDetail(row.id)
    console.log(vulndetail.value)
    form.value = vulndetail.value.data
    showedit.value = true
    console.log(form.value)
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

function onSubmit() {
    console.log(form.value)
}
const typefilterHandler = (value: string, row: any) => {
    return row.vuln_type === value
}
const levelfilterHandler = (value: string, row: any) => {
    return row.level === value
}
const statusfilterHandler = (value: string, row: any) => {
    if (value === 'Poc') {
        return row.poc != ''
    } else {
        return row.exp != ''
    }
}
</script>
<style scoped>
    .status {
  display: flex;
  gap: 4px;
}
</style>