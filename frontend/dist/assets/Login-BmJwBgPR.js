import{d as h,u as v,x,r as w,o as y,c as k,b as s,w as l,h as o,f as V,g as I,y as E,z as S,A as z,e as B,j as u,t as d,B as r,C,D as L,v as N,F as D,_ as K}from"./index-CD0ViaIB.js";import{a as M}from"./api-BG6rU4tp.js";/* empty css                *//* empty css                   */const T={class:"login-container"},U={class:"login-option"},j=h({__name:"Login",setup(A){const{t:e}=v(),c=x(),n=w(""),i=w("");y(()=>{sessionStorage.getItem("token")!=null&&c.push("/")});async function m(){if(n.value==""||i.value==""){r.error(e("app.webui.loginerr1"));return}try{const a=await M.post("/api/v1/login",{username:n.value,password:i.value});a.data.code==0&&r.error(e("app.webui.loginerr3")),a.data.code==1&&(C({title:e("app.webui.loginsucc"),message:a.data.username+", "+e("app.webui.welcome"),type:"success"}),await new Promise(t=>setTimeout(t,1e3)),sessionStorage.setItem("token",a.data.token),sessionStorage.setItem("username",a.data.username),a.data.avatar!=""?sessionStorage.setItem("avatar","/download/file?id="+a.data.avatar):sessionStorage.setItem("avatar","/avatar.svg"),c.push("/")),a.data.code==2&&r.error(e("app.webui.inputformaterror")),a.data.code==3&&r.error(e("app.webui.loginerr4")+" "+a.data.times+" "+e("app.webui.times"))}catch{r.error(e("app.webui.loginerr2"))}}return(a,t)=>{const g=L,f=N,_=D,b=V;return I(),k("div",T,[s(b,{class:"login-card",header:o(e)("app.webui.login"),shadow:"always"},{default:l(()=>[s(g,{modelValue:n.value,"onUpdate:modelValue":t[0]||(t[0]=p=>n.value=p),maxlength:"50",placeholder:o(e)("app.webui.username"),size:"large","prefix-icon":o(E),clearable:""},null,8,["modelValue","placeholder","prefix-icon"]),s(g,{modelValue:i.value,"onUpdate:modelValue":t[1]||(t[1]=p=>i.value=p),maxlength:"50",type:"password","show-password":"",placeholder:o(e)("app.webui.password"),size:"large","prefix-icon":o(S),clearable:"",onKeyup:z(m,["enter","native"])},null,8,["modelValue","placeholder","prefix-icon"]),B("div",U,[s(f,{href:"#/forgotpwd",style:{"margin-right":"10px"}},{default:l(()=>[u(d(o(e)("app.webui.forgot")),1)]),_:1}),s(f,{href:"#/register"},{default:l(()=>[u(d(o(e)("app.webui.register")),1)]),_:1})]),s(_,{type:"success",size:"large",style:{width:"60%",margin:"2% 20%","margin-bottom":"20px","font-weight":"bold","font-size":"16px"},onClick:m,"auto-insert-space":""},{default:l(()=>[u(d(o(e)("app.webui.login")),1)]),_:1})]),_:1},8,["header"])])}}}),G=K(j,[["__scopeId","data-v-1a1aac66"]]);export{G as default};