import{d as M,u as L,a as E,r as H,o as B,c as F,b as t,w as e,e as l,f as A,E as j,g as D,h as d,t as n,i as a,j as r,k as h,l as P,m as I,n as S,p as T,q as N,s as R,v as q,x as G,_ as J}from"./index-JDjDOS_H.js";/* empty css                *//* empty css                        *//* empty css               *//* empty css                *//* empty css                     */import{f as K}from"./utils-DPxaPrc8.js";const O={style:{height:"90vh"}},Q={style:{display:"flex","font-size":"13px","margin-top":"6%",gap:"10px"}},U={style:{display:"flex","font-size":"13px","margin-top":"6%",gap:"10px"}},W={style:{display:"flex","font-size":"13px","margin-top":"6%",gap:"10px"}},X={style:{display:"flex","font-size":"13px","margin-top":"6%",gap:"10px"}},Y={style:{width:"84%",margin:"auto"}},Z={class:"status"},$={style:{display:"flex","margin-left":"85%","margin-top":"1%","align-items":"center"}},e1=M({__name:"Index",setup(l1){const{t:s}=L(),g=E(),o=H({}),b=[z,x];B(()=>{b.forEach(v=>{v()})});function z(){g.redirectedFrom.path==="/login"&&location.reload()}async function x(){try{const v=await A.get("/api/v1/getvulnabs");o.value=v.data}catch{}}function k(v,i){const c=new Date(v.create_time),u=new Date(i.create_time);return c<u?-1:c>u?1:0}return(v,i)=>{const c=I,u=j,m=S,y=T,_=D,w=N,f=R,V=q,C=G;return d(),F("div",O,[t(u,{justify:"space-evenly"},{default:e(()=>[t(_,{shadow:"always",style:{width:"15%",height:"30%"}},{header:e(()=>[t(u,{justify:"space-between",style:{"font-weight":"bold","font-size":"18px"}},{default:e(()=>[l("span",null,n(a(s)("app.webui.totalvuln")),1),t(c,{size:"25"},{default:e(()=>i[0]||(i[0]=[l("svg",{t:"1727430038190",class:"icon",viewBox:"0 0 1024 1024",version:"1.1",xmlns:"http://www.w3.org/2000/svg","p-id":"21943",width:"200",height:"200"},[l("path",{d:"M940 512H792V412c76.8 0 139-62.2 139-139 0-4.4-3.6-8-8-8h-60c-4.4 0-8 3.6-8 8 0 34.8-28.2 63-63 63H232c-34.8 0-63-28.2-63-63 0-4.4-3.6-8-8-8h-60c-4.4 0-8 3.6-8 8 0 76.8 62.2 139 139 139v100H84c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h148v96c0 6.5 0.2 13 0.7 19.3C164.1 728.6 116 796.7 116 876c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8 0-44.2 23.9-82.9 59.6-103.7 6 17.2 13.6 33.6 22.7 49 24.3 41.5 59 76.2 100.5 100.5S460.5 960 512 960s99.8-13.9 141.3-38.2c41.5-24.3 76.2-59 100.5-100.5 9.1-15.5 16.7-31.9 22.7-49C812.1 793.1 836 831.8 836 876c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8 0-79.3-48.1-147.4-116.7-176.7 0.4-6.4 0.7-12.8 0.7-19.3v-96h148c4.4 0 8-3.6 8-8v-56c0-4.4-3.6-8-8-8zM716 680c0 36.8-9.7 72-27.8 102.9-17.7 30.3-43 55.6-73.3 73.3-20.1 11.8-42 20-64.9 24.3V484c0-4.4-3.6-8-8-8h-60c-4.4 0-8 3.6-8 8v396.5c-22.9-4.3-44.8-12.5-64.9-24.3-30.3-17.7-55.6-43-73.3-73.3C317.7 752 308 716.8 308 680V412h408v268z","p-id":"21944",fill:"#d81e06"}),l("path",{d:"M304 280h56c4.4 0 8-3.6 8-8 0-28.3 5.9-53.2 17.1-73.5 10.6-19.4 26-34.8 45.4-45.4C450.9 142 475.7 136 504 136h16c28.3 0 53.2 5.9 73.5 17.1 19.4 10.6 34.8 26 45.4 45.4C650 218.9 656 243.7 656 272c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8 0-40-8.8-76.7-25.9-108.1-17.2-31.5-42.5-56.8-74-74C596.7 72.8 560 64 520 64h-16c-40 0-76.7 8.8-108.1 25.9-31.5 17.2-56.8 42.5-74 74C304.8 195.3 296 232 296 272c0 4.4 3.6 8 8 8z","p-id":"21945",fill:"#d81e06"})],-1)])),_:1})]),_:1})]),default:e(()=>[t(y,null,{default:e(()=>[t(m,{value:o.value.total},null,8,["value"]),l("div",Q,[l("span",null,n(a(s)("app.webui.weeklyadditions")),1),l("span",null,n(o.value.weeklyAdditionsVuln),1)])]),_:1})]),_:1}),t(_,{shadow:"always",style:{width:"15%",height:"30%"}},{header:e(()=>[t(u,{justify:"space-between",style:{"font-weight":"bold","font-size":"18px"}},{default:e(()=>[l("span",null,n(a(s)("app.webui.totalpoc")),1),t(c,{size:"25"},{default:e(()=>i[1]||(i[1]=[l("svg",{t:"1727429978264",class:"icon",viewBox:"0 0 1161 1024",version:"1.1",xmlns:"http://www.w3.org/2000/svg","p-id":"20805",width:"200",height:"200"},[l("path",{d:"M258.205538 354.461538c10.436923 0 18.924308 8.467692 18.924308 18.944v624.836924c0 10.456615-8.467692 18.944-18.924308 18.944H177.289846a18.944 18.944 0 0 1-18.924308-18.944v-624.836924c0-10.476308 8.467692-18.944 18.924308-18.944h80.915692z m158.365539 120.497231c10.436923 0 18.924308 8.467692 18.924308 18.924308v504.359385c0 10.456615-8.467692 18.944-18.924308 18.944h-80.915692a18.944 18.944 0 0 1-18.924308-18.944V493.883077c0-10.436923 8.467692-18.924308 18.924308-18.924308h80.915692z m-316.731077 0c10.456615 0 18.924308 8.467692 18.924308 18.924308v504.359385c0 10.456615-8.467692 18.944-18.924308 18.944H18.944A18.944 18.944 0 0 1 0 998.242462V493.883077c0-10.436923 8.467692-18.924308 18.944-18.924308H99.84z m139.421538-82.628923H196.214154v586.988308h43.047384V392.329846z m158.365539 120.477539h-43.027692V979.298462h43.027692V512.787692z m-316.731077 0H37.868308V979.298462h43.027692V512.787692zM1063.384615 0a98.461538 98.461538 0 0 1 98.461539 98.461538v748.307693a98.461538 98.461538 0 0 1-98.461539 98.461538H492.307692V433.230769a39.384615 39.384615 0 0 0-36.430769-39.286154L452.923077 393.846154h-98.461539v-59.076923a39.384615 39.384615 0 0 0-36.430769-39.286154L315.076923 295.384615h-98.461538V98.461538a98.461538 98.461538 0 0 1 98.461538-98.461538h748.307692z m-142.867692 520.664615c-44.898462 0-78.769231 14.572308-102.4 44.504616-21.267692 26.387692-31.507692 60.258462-31.507692 102.006154 0 42.535385 9.846154 76.406154 30.326154 101.612307 23.236923 29.144615 57.895385 44.110769 103.975384 44.11077 30.326154 0 56.32-8.664615 77.981539-25.993847 23.236923-18.510769 37.809231-44.110769 44.110769-76.8h-41.747692l-1.476923 5.395693c-5.710769 19.377231-15.064615 34.126769-28.061539 44.228923-13.390769 10.24-30.72 15.36-51.2 15.36-31.507692 0-54.744615-9.846154-69.710769-29.538462-14.178462-18.510769-21.267692-44.504615-21.267692-78.375384 0-32.689231 7.089231-58.683077 21.661538-77.587693 15.36-20.873846 38.203077-31.113846 68.529231-31.113846 20.48 0 37.021538 4.332308 50.412307 13.390769 13.390769 9.058462 22.449231 23.236923 27.175385 42.141539h41.747692l-0.984615-5.592616c-5.060923-25.796923-17.329231-46.710154-36.430769-62.148923-21.267692-17.329231-48.443077-25.6-81.132308-25.6z m-30.956308-413.538461c-42.929231 0-76.8 13.784615-101.218461 42.141538-23.630769 26.781538-35.052308 61.44-35.052308 104.369231 0 42.535385 11.421538 77.193846 35.052308 103.975385 24.418462 27.569231 58.289231 41.747692 101.218461 41.747692 42.535385 0 76.406154-13.784615 101.218462-41.353846 23.630769-26.781538 35.446154-61.44 35.446154-104.369231 0-43.323077-11.815385-78.375385-35.446154-104.763077-24.812308-27.963077-58.683077-41.747692-101.218462-41.747692z m-292.489846 5.513846H480.492308V393.846154h42.92923v-109.489231h72.861539l5.395692-0.059077c63.763692-1.516308 95.822769-30.247385 95.822769-86.193231 0-57.107692-33.476923-85.464615-100.430769-85.464615z m292.470154 32.295385c29.932308 0 53.169231 9.452308 69.316923 29.144615 15.753846 19.298462 24.024615 45.686154 24.024616 79.556923 0 33.476923-8.270769 59.864615-24.024616 78.769231-16.147692 19.298462-39.384615 29.144615-69.316923 29.144615s-53.169231-10.24-69.710769-29.932307c-15.753846-19.298462-23.630769-45.292308-23.630769-77.981539 0-33.083077 7.876923-59.076923 23.630769-78.375385 16.541538-20.48 39.778462-30.326154 69.710769-30.326153z m-296.014769 4.332307c20.873846 0 36.233846 3.938462 46.08 11.815385 9.846154 7.089231 14.966154 19.692308 14.966154 37.021538 0 17.329231-5.12 29.932308-14.572308 37.809231-9.846154 7.876923-25.206154 11.815385-46.473846 11.815385h-70.104616v-98.461539z","p-id":"20806",fill:"#ed5565"})],-1)])),_:1})]),_:1})]),default:e(()=>[t(y,null,{default:e(()=>[t(m,{value:o.value.hasPoc},null,8,["value"]),l("div",U,[l("span",null,n(a(s)("app.webui.weeklyadditionspoc")),1),l("span",null,n(o.value.weeklyAdditionsPoc),1)])]),_:1})]),_:1}),t(_,{shadow:"always",style:{width:"15%",height:"30%"}},{header:e(()=>[t(u,{justify:"space-between",style:{"font-weight":"bold","font-size":"18px"}},{default:e(()=>[l("span",null,n(a(s)("app.webui.totalexp")),1),t(c,{size:"25"},{default:e(()=>i[2]||(i[2]=[l("svg",{t:"1727429727880",class:"icon",viewBox:"0 0 1024 1024",version:"1.1",xmlns:"http://www.w3.org/2000/svg","p-id":"13778",width:"200",height:"200"},[l("path",{d:"M676.408889 303.407407l-101.262222-175.976296c-132.437333-40.201481-114.801778 130.085926-203.207111 168.846222L303.407407 246.708148C347.477333 160.274963 387.337481 83.057778 429.283556 0h206.582518L758.518519 235.747556 676.408889 303.407407zM127.393185 493.037037l80.327111 61.477926c-28.899556 61.060741-55.48563 117.380741-82.223407 173.018074 29.430519 138.657185 149.238519 47.824593 215.836444 103.917037V910.222222H94.966519L0 735.156148 127.393185 493.037037zM1024 735.004444L925.999407 910.222222H682.666667v-79.568592c68.342519-48.279704 171.956148 24.045037 212.574814-97.583408l-83.133629-184.395852L895.089778 493.037037 1024 735.004444zM0 387.640889L101.793185 189.62963 303.407407 422.874074l-12.818963 32.237037L0 387.640889zM1024 358.362074L739.745185 417.185185 720.592593 392.836741 919.210667 189.62963 1024 358.362074zM417.185185 1024l80.402963-263.661037 26.699852-1.820444L606.814815 1024h-189.62963z",fill:"#1296db","p-id":"13779"}),l("path",{d:"M606.814815 455.111111h-189.62963v-75.851852h189.62963v75.851852zM682.666667 568.888889H341.333333v-75.851852h341.333334v75.851852zM796.444444 682.666667H227.555556v-75.851852h568.888888v75.851852z",fill:"#1296db","p-id":"13780"})],-1)])),_:1})]),_:1})]),default:e(()=>[t(y,null,{default:e(()=>[t(m,{value:o.value.hasExp},null,8,["value"]),l("div",W,[l("span",null,n(a(s)("app.webui.weeklyadditionsexp")),1),l("span",null,n(o.value.weeklyAdditionsExp),1)])]),_:1})]),_:1}),t(_,{shadow:"always",style:{width:"15%",height:"30%"}},{header:e(()=>[t(u,{justify:"space-between",style:{"font-weight":"bold","font-size":"18px"}},{default:e(()=>[l("span",null,n(a(s)("app.webui.totalproduct")),1),t(c,{size:"25"},{default:e(()=>i[3]||(i[3]=[l("svg",{t:"1727429889519",class:"icon",viewBox:"0 0 1025 1024",version:"1.1",xmlns:"http://www.w3.org/2000/svg","p-id":"17925",width:"200",height:"200"},[l("path",{d:"M302.32811 504.78l-146.944-88.063c-11.776-7.168-27.136-3.328-34.048 8.448-7.168 11.776-3.328 27.136 8.448 34.304l146.944 88.064c11.776 7.168 27.136 3.328 34.048-8.704 7.168-11.802 3.328-27.162-8.448-34.048zM459.00011 1020.878c5.888 1.792 12.544 1.28 18.688-2.304-5.632 3.072-12.288 4.096-18.688 2.304zM302.32811 504.78l-146.944-88.064c-11.776-7.168-27.136-3.328-34.048 8.448-7.168 11.776-3.328 27.136 8.448 34.304l146.944 88.064c11.776 7.168 27.136 3.328 34.048-8.704 7.168-11.802 3.328-27.162-8.448-34.048z m0 0l-146.944-88.064c-11.776-7.168-27.136-3.328-34.048 8.448-7.168 11.776-3.328 27.136 8.448 34.304l146.944 88.064c11.776 7.168 27.136 3.328 34.048-8.704 7.168-11.802 3.328-27.162-8.448-34.048z m0 0l-146.944-88.064c-11.776-7.168-27.136-3.328-34.048 8.448-7.168 11.776-3.328 27.136 8.448 34.304l146.944 88.064c11.776 7.168 27.136 3.328 34.048-8.704 7.168-11.802 3.328-27.162-8.448-34.048z m0 0l-146.944-88.064c-11.776-7.168-27.136-3.328-34.048 8.448-7.168 11.776-3.328 27.136 8.448 34.304l146.944 88.064c11.776 7.168 27.136 3.328 34.048-8.704 7.168-11.802 3.328-27.162-8.448-34.048z m0 0l-146.944-88.064c-11.776-7.168-27.136-3.328-34.048 8.448-7.168 11.776-3.328 27.136 8.448 34.304l146.944 88.064c11.776 7.168 27.136 3.328 34.048-8.704 7.168-11.802 3.328-27.162-8.448-34.048z",fill:"#3CBF71","p-id":"17926"}),l("path",{d:"M477.68811 1018.547c-6.144 3.584-12.8 4.096-18.688 2.304 6.4 1.818 13.056 0.794 18.688-2.304z",fill:"#3CBF71","p-id":"17927"}),l("path",{d:"M490.64111 954.112V511.744L880.27311 290.56l39.885-24.166c7.987-5.197 9.498-11.802 10.547-16.538 1.024-9.728-3.584-19.968-13.568-25.088L476.81711 4.608c-7.168-3.584-15.36-3.584-22.528 0L13.97011 224.768C-2.41489 235.52 0.14511 259.328 0.14511 261.888v469.76c0 8.704 4.608 16.896 12.032 21.504 0 0 430.848 260.352 438.272 263.68 15.104 7.424 24.576 5.376 31.744 1.792l143.36-82.944a25.56 25.56 0 0 0 9.216-35.072c-6.912-12.288-22.784-16.384-34.816-9.472l-99.072 57.6v-0.256l-10.24 5.632m-50.176-442.368v441.088l-390.4-235.264v-427.52l390.4 221.696z m25.088-43.52L78.22511 248.576 465.55311 55.04l387.328 193.536-387.328 219.648z",fill:"#3CBF71","p-id":"17928"}),l("path",{d:"M812.94511 917.376l201.984-121.344c5.632-3.584 9.216-9.728 9.216-16.64V556.928c0-7.834-5.888-13.312-12.288-16.384l-5.12-2.56c-0.512-0.256-1.28-0.256-1.792-0.256h1.536L810.13011 439.68c-5.376-2.56-11.776-2.56-17.152 0l-196.096 98.048c-14.336 5.504-17.92 17.792-17.92 24.704v216.96c0 6.912 3.584 13.056 9.472 16.64 0 0 200.167 120.09 201.19 120.346 0 0.025 11.93 8.857 23.322 0.998z m-30.464-49.152l-165.12-99.584V589.696l165.12 93.44v185.088z m19.2-218.112L640.65711 558.72l161.024-80.384L962.45011 558.72l-160.768 91.392zM985.74511 768.64l-164.864 99.328V683.136l164.864-93.44V768.64z",fill:"#3CBF71","p-id":"17929"})],-1)])),_:1})]),_:1})]),default:e(()=>[t(y,null,{default:e(()=>[t(m,{value:o.value.affectedProduct},null,8,["value"]),l("div",X,[l("span",null,n(a(s)("app.webui.weeklyadditionsproduct")),1),l("span",null,n(o.value.weeklyAdditionsProduct),1)])]),_:1})]),_:1})]),_:1}),l("div",Y,[t(_,{style:{"margin-top":"2%"},shadow:"always"},{default:e(()=>[t(V,{data:o.value.data},{default:e(()=>[t(w,{prop:"id",label:a(s)("app.webui.id"),width:"180",sortable:""},null,8,["label"]),t(w,{prop:"vuln_name",label:a(s)("app.webui.name")},null,8,["label"]),t(w,{prop:"vuln_type",label:a(s)("app.webui.type"),width:"180"},null,8,["label"]),t(w,{prop:"vuln_level",label:a(s)("app.webui.level"),width:"80"},{default:e(({row:p})=>[p.vuln_level==="Critical"?(d(),r(f,{key:0,type:"danger",effect:"dark",color:"#CD0000"},{default:e(()=>[h(n(a(s)("app.webui.critical")),1)]),_:1})):p.vuln_level==="High"?(d(),r(f,{key:1,type:"danger",effect:"dark",color:"#EE2C2C"},{default:e(()=>[h(n(a(s)("app.webui.high")),1)]),_:1})):p.vuln_level==="Medium"?(d(),r(f,{key:2,type:"warning",effect:"dark",color:"#FF6600"},{default:e(()=>[h(n(a(s)("app.webui.medium")),1)]),_:1})):(d(),r(f,{key:3,type:"primary",effect:"dark"},{default:e(()=>[h(n(a(s)("app.webui.low")),1)]),_:1}))]),_:1},8,["label"]),t(w,{prop:"cvss",label:"CVSS",width:"80"}),t(w,{label:a(s)("app.webui.status"),width:"120"},{default:e(({row:p})=>[l("div",Z,[p.poc!=""&&p.poc!="0"?(d(),r(f,{key:0,type:"success",effect:"dark"},{default:e(()=>i[4]||(i[4]=[h("Poc")])),_:1})):(d(),r(f,{key:1,type:"info",effect:"dark"},{default:e(()=>i[5]||(i[5]=[h("Poc")])),_:1})),p.exp!=""&&p.exp!="0"?(d(),r(f,{key:2,type:"success",effect:"dark"},{default:e(()=>i[6]||(i[6]=[h("Exp")])),_:1})):(d(),r(f,{key:3,type:"info",effect:"dark"},{default:e(()=>i[7]||(i[7]=[h("Exp")])),_:1}))])]),_:1},8,["label"]),t(w,{label:a(s)("app.webui.createtime"),width:"120",sortable:"","sort-method":k},{default:e(({row:p})=>[l("span",null,n(a(K)(p.create_time)),1)]),_:1},8,["label"])]),_:1},8,["data"]),l("div",$,[t(C,{href:"#/vulnlist",style:{"font-size":"18px"}},{default:e(()=>[l("span",null,n(a(s)("app.webui.readmore")),1),t(c,null,{default:e(()=>[t(a(P))]),_:1})]),_:1})])]),_:1})])])}}}),d1=J(e1,[["__scopeId","data-v-fbe6ec4c"]]);export{d1 as default};