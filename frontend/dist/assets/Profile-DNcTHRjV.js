import{d as E,u as M,r as g,o as A,c as U,b as n,w as r,i,f as w,g as H,h as L,e,t as l,k as F,C as c,V as $,Q as D,m as N,F as Z,G as j}from"./index-BZXfLavj.js";/* empty css                *//* empty css                 *//* empty css                    */import{c as P}from"./utils-CYkVsYWI.js";/* empty css                   */const R={style:{width:"70%",height:"90vh",margin:"auto"}},G={style:{display:"flex","font-weight":"lighter","font-size":"14px"}},O={style:{"margin-left":"auto",display:"flex","align-items":"center"}},Q={style:{"margin-left":"20px"}},T={style:{"margin-left":"30px",display:"flex","align-items":"center"}},q={style:{"margin-left":"20px"}},J={style:{"margin-left":"30px",display:"flex","align-items":"center"}},K={style:{"margin-left":"20px"}},W={style:{"margin-top":"20px","font-weight":"lighter","font-size":"14px"}},X={style:{"margin-top":"20px","font-weight":"lighter","font-size":"14px"}},Y={style:{"margin-top":"20px","font-weight":"lighter","font-size":"14px"}},ie=E({__name:"Profile",setup(ee){const{t:o}=M(),d=sessionStorage.getItem("token"),a=g({}),f=g(!0),x=[P,y],_=g(sessionStorage.getItem("avatar"));A(()=>{x.forEach(s=>{s()})});async function y(){try{const s={headers:{Authorization:`Bearer ${d}`}},t=await w.get("/api/v1/userinfo",s);a.value=t.data.data,d&&t.data.code==0&&(sessionStorage.removeItem("token"),sessionStorage.removeItem("username"),sessionStorage.removeItem("avatar"),location.reload())}catch{}}const z=s=>{_.value=URL.createObjectURL(s)},b=s=>{c.success(o("app.webui.uploadsucc")),sessionStorage.setItem("avatar","/download/file?id="+s.data),location.reload()},V=()=>{if(!a.value.email.match(/^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$/)){c.error(o("app.webui.emailformat"));return}m()},C=()=>{if(!a.value.phone.match(/^[1][3,4,5,6,7,8,9][0-9]{9}$/)){c.error(o("app.webui.phoneformat"));return}m()},m=()=>{f.value=!1},k=async()=>{console.log(a.value);try{const s={headers:{Authorization:`Bearer ${d}`}},t={username:a.value.username,email:a.value.email,phone:a.value.phone},u=await w.post("/api/v1/updateuserinfo",t,s);u.data.code==0?(sessionStorage.removeItem("token"),sessionStorage.removeItem("username"),sessionStorage.removeItem("avatar"),location.reload()):u.data.code==1?c.success(o("app.webui.modifysucc")):c.error(o("app.webui.modifyerr"))}catch(s){console.error(s)}};return(s,t)=>{const u=$,I=D,h=N,v=Z,S=j,B=H;return L(),U("div",R,[n(B,{style:{padding:"30px","font-weight":"bold","font-size":"20px"},header:i(o)("app.webui.myprofile")},{default:r(()=>[e("div",G,[n(I,{accept:".png,.jpg",action:"/api/v1/updateavatar","on-success":b,headers:{Authorization:`Bearer ${i(d)}`},"before-upload":z},{default:r(()=>[n(u,{size:"large",src:_.value},null,8,["src"])]),_:1},8,["headers"]),e("div",O,[n(h,{size:"30"},{default:r(()=>t[3]||(t[3]=[e("svg",{t:"1728380528569",class:"icon",viewBox:"0 0 1024 1024",version:"1.1",xmlns:"http://www.w3.org/2000/svg","p-id":"14698",width:"200",height:"200"},[e("path",{d:"M752.006095 165.327238h-90.672762v69.339429h-58.660571V165.327238H421.302857v69.339429h-58.660571V165.327238h-90.672762v111.299048l39.936 45.616762-51.809524 29.915428-38.863238-44.446476a31.98781 31.98781 0 0 1-7.92381-21.065143v-147.992381c0-17.65181 14.336-31.98781 32.012191-31.987809h533.308952c17.67619 0 32.01219 14.336 32.012191 32.01219v147.968c0 7.753143-2.80381 15.238095-7.92381 21.065143l-38.863238 44.446476-51.809524-29.915428 39.936-45.616762V165.302857zM512 768a170.666667 170.666667 0 1 0 0-341.333333 170.666667 170.666667 0 0 0 0 341.333333z m0-58.660571a112.006095 112.006095 0 1 1 0-224.012191 112.006095 112.006095 0 0 1 0 224.012191z","p-id":"14699",fill:"#FFD700"}),e("path",{d:"M527.993905 265.240381a31.98781 31.98781 0 0 0-31.98781 0L232.399238 417.401905a32.01219 32.01219 0 0 0-15.993905 27.721143v304.371809c0 11.459048 6.095238 21.991619 15.993905 27.721143l263.606857 152.185905c9.898667 5.729524 22.089143 5.729524 31.98781 0l263.606857-152.185905c9.898667-5.729524 15.993905-16.286476 15.993905-27.721143V445.147429c0-11.459048-6.095238-21.991619-15.993905-27.721143L527.993905 265.264762z m-15.993905 58.514286l236.934095 136.777143v273.603047L512 870.887619l-236.934095-136.777143V460.53181L512 323.779048z","p-id":"14700",fill:"#FFD700"})],-1)])),_:1}),e("span",null,l(i(o)("app.webui.ranking")),1),e("span",Q,l(a.value.ranking),1)]),e("div",T,[n(h,{size:"30"},{default:r(()=>t[4]||(t[4]=[e("svg",{t:"1727430038190",class:"icon",viewBox:"0 0 1024 1024",version:"1.1",xmlns:"http://www.w3.org/2000/svg","p-id":"21943",width:"200",height:"200"},[e("path",{d:"M940 512H792V412c76.8 0 139-62.2 139-139 0-4.4-3.6-8-8-8h-60c-4.4 0-8 3.6-8 8 0 34.8-28.2 63-63 63H232c-34.8 0-63-28.2-63-63 0-4.4-3.6-8-8-8h-60c-4.4 0-8 3.6-8 8 0 76.8 62.2 139 139 139v100H84c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h148v96c0 6.5 0.2 13 0.7 19.3C164.1 728.6 116 796.7 116 876c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8 0-44.2 23.9-82.9 59.6-103.7 6 17.2 13.6 33.6 22.7 49 24.3 41.5 59 76.2 100.5 100.5S460.5 960 512 960s99.8-13.9 141.3-38.2c41.5-24.3 76.2-59 100.5-100.5 9.1-15.5 16.7-31.9 22.7-49C812.1 793.1 836 831.8 836 876c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8 0-79.3-48.1-147.4-116.7-176.7 0.4-6.4 0.7-12.8 0.7-19.3v-96h148c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8zM716 680c0 36.8-9.7 72-27.8 102.9-17.7 30.3-43 55.6-73.3 73.3-20.1 11.8-42 20-64.9 24.3V484c0-4.4-3.6-8-8-8h-60c-4.4 0-8 3.6-8 8v396.5c-22.9-4.3-44.8-12.5-64.9-24.3-30.3-17.7-55.6-43-73.3-73.3C317.7 752 308 716.8 308 680V412h408v268z","p-id":"21944",fill:"#d81e06"}),e("path",{d:"M304 280h56c4.4 0 8-3.6 8-8 0-28.3 5.9-53.2 17.1-73.5 10.6-19.4 26-34.8 45.4-45.4C450.9 142 475.7 136 504 136h16c28.3 0 53.2 5.9 73.5 17.1 19.4 10.6 34.8 26 45.4 45.4C650 218.9 656 243.7 656 272c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8 0-40-8.8-76.7-25.9-108.1-17.2-31.5-42.5-56.8-74-74C596.7 72.8 560 64 520 64h-16c-40 0-76.7 8.8-108.1 25.9-31.5 17.2-56.8 42.5-74 74C304.8 195.3 296 232 296 272c0 4.4 3.6 8 8 8z","p-id":"21945",fill:"#d81e06"})],-1)])),_:1}),e("span",null,l(i(o)("app.webui.totalvuln")),1),e("span",q,l(a.value.total),1)]),e("div",J,[n(h,{size:"30"},{default:r(()=>t[5]||(t[5]=[e("svg",{t:"1728382655697",class:"icon",viewBox:"0 0 1024 1024",version:"1.1",xmlns:"http://www.w3.org/2000/svg","p-id":"16992",width:"200",height:"200"},[e("path",{d:"M127.296 934.4h831.552v89.6H127.296v-89.6zM896.96 695.04H695.168a123.648 123.648 0 0 1-33.216-70.784c-8.96-108.352 111.232-160.32 111.232-292.928C774.4 218.496 685.248 128 542.656 128s-230.4 90.496-230.4 203.328c0 133.44 120.192 184.512 111.232 292.928A121.024 121.024 0 0 1 390.4 695.04H188.352a60.288 60.288 0 0 0-60.992 60.032v119.104h831.488V755.2a61.056 61.056 0 0 0-61.888-60.16z",fill:"#70C1AE","p-id":"16993"})],-1)])),_:1}),e("span",null,l(i(o)("app.webui.auditedvuln")),1),e("span",K,l(a.value.authed),1)])]),e("div",W,l(i(o)("app.webui.username")),1),n(v,{modelValue:a.value.username,"onUpdate:modelValue":t[0]||(t[0]=p=>a.value.username=p),size:"large",placeholder:a.value.username,style:{"margin-top":"10px"},onChange:m},null,8,["modelValue","placeholder"]),e("div",X,l(i(o)("app.webui.email")),1),n(v,{modelValue:a.value.email,"onUpdate:modelValue":t[1]||(t[1]=p=>a.value.email=p),size:"large",placeholder:a.value.email,style:{"margin-top":"10px"},onChange:V},null,8,["modelValue","placeholder"]),e("div",Y,l(i(o)("app.webui.phone")),1),n(v,{modelValue:a.value.phone,"onUpdate:modelValue":t[2]||(t[2]=p=>a.value.phone=p),size:"large",placeholder:a.value.phone,style:{"margin-top":"10px"},onChange:C},null,8,["modelValue","placeholder"]),n(S,{type:"primary",style:{"margin-top":"20px"},disabled:f.value,onClick:k},{default:r(()=>[F(l(i(o)("app.webui.modify")),1)]),_:1},8,["disabled"])]),_:1},8,["header"])])}}});export{ie as default};
