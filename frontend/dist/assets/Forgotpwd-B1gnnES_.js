import{d as k,u as z,r as c,c as B,b as s,w as f,h as r,f as E,g as A,y as C,V as U,W as Z,z as F,e as I,j as w,t as g,B as a,D as N,v as D,F as $,_ as j}from"./index-CD0ViaIB.js";import{a as L}from"./api-BG6rU4tp.js";/* empty css                *//* empty css                   */const M={class:"forgot-container"},S={class:"forgot-option"},T=k({__name:"Forgotpwd",setup(W){const{t:e}=z(),u=c(""),l=c(""),p=c(""),n=c("");function b(){u.value==""&&a.error(e("app.webui.loginerr1"))}function _(){if(n.value==""){a.error(e("app.webui.emailempty"));return}if(!n.value.match(/^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$/)){a.error(e("app.webui.emailformat"));return}}function m(){return l.value==""?(a.error(e("app.webui.passwordempty")),!1):l.value.length<8?(a.error(e("app.webui.passwordlength")),!1):l.value.match(/^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[^]{8,}$/)?!0:(a.error(e("app.webui.passwordcomplex")),!1)}function v(){if(l.value!=p.value){a.error(e("app.webui.passwordnotmatch"));return}}async function h(){if(u.value==""||l.value==""||p.value==""||n.value==""){a.error(e("app.webui.missingnotice"));return}if(l.value!=p.value){a.error(e("app.webui.passwordnotmatch"));return}if(m())try{const d={username:u.value,password:p.value,email:n.value},o=await L.post("/api/v1/forgot",d);o.data.code==0?a.error(e("app.webui.forgotnotice")):o.data.code==1?a.success(e("app.webui.forgotsuccess")):o.data.code==2?a.error(e("app.webui.usernamealreadyexist")):o.data.code==4?a.error(e("app.webui.emailformat")):o.data.code==5?a.error(e("app.webui.emailalreadyexist")):a.error(e("app.webui.forgotfail"))}catch{a.error(e("app.webui.forgotfail"))}}return(d,o)=>{const i=N,x=D,V=$,y=E;return A(),B("div",M,[s(y,{class:"forgot-card",header:r(e)("app.webui.forgot"),shadow:"always"},{default:f(()=>[s(i,{modelValue:u.value,"onUpdate:modelValue":o[0]||(o[0]=t=>u.value=t),maxlength:"50",placeholder:r(e)("app.webui.username"),size:"large","prefix-icon":r(C),clearable:"",onBlur:b},null,8,["modelValue","placeholder","prefix-icon"]),s(i,{modelValue:n.value,"onUpdate:modelValue":o[1]||(o[1]=t=>n.value=t),maxlength:"50",placeholder:r(e)("app.webui.email"),size:"large","prefix-icon":r(U),clearable:"",onBlur:_},null,8,["modelValue","placeholder","prefix-icon"]),s(i,{modelValue:l.value,"onUpdate:modelValue":o[2]||(o[2]=t=>l.value=t),maxlength:"50",type:"password","show-password":"",placeholder:r(e)("app.webui.password"),size:"large","prefix-icon":r(Z),clearable:"",onBlur:m},null,8,["modelValue","placeholder","prefix-icon"]),s(i,{modelValue:p.value,"onUpdate:modelValue":o[3]||(o[3]=t=>p.value=t),maxlength:"50",type:"password","show-password":"",placeholder:r(e)("app.webui.confirmpassword"),size:"large","prefix-icon":r(F),clearable:"",onBlur:v},null,8,["modelValue","placeholder","prefix-icon"]),I("div",S,[s(x,{href:"#/login"},{default:f(()=>[w(g(r(e)("app.webui.returnlogin")),1)]),_:1})]),s(V,{type:"success",size:"large",style:{width:"60%",margin:"2% 20%","margin-bottom":"20px","font-weight":"bold","font-size":"16px"},onClick:h,"auto-insert-space":""},{default:f(()=>[w(g(r(e)("app.webui.forgot")),1)]),_:1})]),_:1},8,["header"])])}}}),K=j(T,[["__scopeId","data-v-a7bc11e1"]]);export{K as default};