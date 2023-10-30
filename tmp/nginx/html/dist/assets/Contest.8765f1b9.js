import{a as L,d as M,c as z,_ as R}from"./index.de99a551.js";import{I as q}from"./Input.a729bc6b.js";const d=p=>(Vue.pushScopeId("data-v-346e9fbe"),p=p(),Vue.popScopeId(),p),H={class:"contest"},O={class:"notFound"},j={class:"infoBox",ref:"infoBox"},G={class:"artFont bold",style:{"margin-top":"24px"}},J={class:"title bold artFont"},K={key:0,class:"ctype ctypeICPC bold"},Q={key:1,class:"ctype ctypeOI bold"},W={class:"text"},X={class:"text"},Y={key:0,class:"status"},Z=d(()=>Vue.createElementVNode("div",{class:"point",style:{"background-color":"#5ebd00"}},null,-1)),ee={key:1,class:"status"},te=d(()=>Vue.createElementVNode("div",{class:"point",style:{"background-color":"#ff3300"}},null,-1)),oe={class:"left_time"},ne={class:"time"},se={class:"begin_time"},ue={class:"end_time"},ie={class:"process"},le={class:"functionBox"},re={class:"problemList",ref:"problemList"},ce=d(()=>Vue.createElementVNode("div",{class:"nav"},[Vue.createElementVNode("div",{style:{width:"90px"}},"\u5E8F\u53F7"),Vue.createElementVNode("div",{style:{width:"calc(100% - 190px)"}},"\u9898\u76EE"),Vue.createElementVNode("div",{style:{width:"100px"}},"\u901A\u8FC7\u60C5\u51B5")],-1)),ae=["onClick"],Ve=["onClick"],me={class:"status"},de={class:"needPass"},pe=d(()=>Vue.createElementVNode("div",{class:"title"},"\u9A8C\u8BC1",-1)),fe={class:"input"},_e=d(()=>Vue.createElementVNode("div",{class:"label"},"\u5BC6\u7801",-1)),Ce=Vue.defineComponent({__name:"Contest",setup(p){const{proxy:n}=Vue.getCurrentInstance(),D=L(),h=M(),N=z();var m=Vue.ref(!1),c=Vue.ref(""),f=Vue.ref(!0),T=Vue.ref(!1),_=Vue.reactive({contestInfo:null,problemList:null,init(){this.contestInfo&&this.contestInfo.close(),this.problemList&&this.problemList.close(),this.contestInfo=null,this.problemList=null}}),t=Vue.reactive({CID:null,Data:[],length:0,BeginTime:0,Type:1,Description:"",EndTime:0,IsPublic:1,Pass:null,Size:0,Title:"",UID:"",copy(e){let o=e.Problems.split(",");for(let r in o)for(let i in e.Data)if(o[r]==e.Data[i].PID){t.Data.push({PID:e.Data[i].PID,Title:e.Data[i].Title,ACNum:e.Data[i].ACNum,SubmitNum:e.Data[i].SubmitNum});break}t.length=e.length,t.BeginTime=e.BeginTime,t.CID=e.CID,t.Type=e.Type,t.Description=e.Description,t.EndTime=e.EndTime,t.IsPublic=e.IsPublic,t.Pass=e.Pass,t.Size=e.Size,t.Title=e.Title,t.UID=e.UID}}),u=Vue.reactive({status:1,percent:0,color:"#2cbbfe",timmer:null,lostTime:0,allTime:0,timeDistance:0,init(){u.status=1,u.percent=0,u.color="#2cbbfe",u.timmer=null,u.lostTime=0,u.allTime=0,u.timeDistance=0,clearInterval(u.timmer)},begin(){this.init();let e=Date.now()-this.timeDistance;if(e>=t.EndTime){this.percent=100,this.color="#9e9e9e",this.status=0;return}this.lostTime=e-t.BeginTime,this.allTime=Math.abs(t.EndTime-t.BeginTime),this.timmer=setInterval(()=>{if(this.lostTime+=1e3,this.percent=Math.floor(this.lostTime/this.allTime*100),this.percent>=100){this.percent=100,this.color="#9e9e9e",this.status=0,clearInterval(u.timmer);return}else this.percent>=90?this.color="#f03e3e":this.percent>=60?this.color="#ff8c00":this.percent>=40?this.color="#bcee68":this.color="#66cd00"},1e3)}});async function y(){var s,a;let e=n.$route.params.CID;if(t.CID=e,!e){n.elMessage({message:"\u9875\u9762\u8DF3\u8F6C\u5F02\u5E38\uFF0C\u8BF7\u91CD\u8BD5\u3002",type:"error"}),m.value=!1;return}let o=D.getContestRouterData(e),r=(s=o==null?void 0:o.IsPublic)!=null?s:1,i=(a=o==null?void 0:o.Pass)!=null?a:"";r==-1?g(e,i):r==1&&C(e,i),T.value=I()}function I(){return h.isLogin?(h.PermissionMap&N.ContestAdminBit)!=0||(h.PermissionMap&N.SuperAdminBit)!=0:!1}async function g(e,o){o?C(e,o):(c.value="",m.value=!0)}async function C(e,o){m.value=!1,f.value=!0,_.init(),_.contestInfo=n.elLoading({node:n.$refs.infoBox}),_.problemList=n.elLoading({node:n.$refs.problemList});let r={Pass:o};await n.$get("api/contest/"+e,r).then(i=>{var a,V;let s=i.data;s.code==0?(t.copy(s),D.setContestRouterData(e,o?-1:1,o),f.value=!1):s.code==160504?g(e,""):s.code==160503&&n.$router.push({path:"/Contests"}),n.codeProcessor((a=s==null?void 0:s.code)!=null?a:100001,(V=s==null?void 0:s.msg)!=null?V:"\u670D\u52A1\u5668\u9519\u8BEF\\\\error")}),_.init(),await n.getServerTime().then(i=>{let s=Date.now();if(i.time==null||Math.abs(i.time-s)<1500){u.begin();return}u.timeDistance=s-i.time,u.begin()})}function E(e){n.$router.push({name:"ContestProblem",params:{PID:e,CID:t.CID}})}function b(){let e=t.CID;n.$router.push({name:"ContestRank",params:{CID:e}})}function k(){let e=t.CID,o;n.$route.params.Pass&&(o=n.$route.params.Pass),n.$router.push({path:"/ContestStatus",query:{CID:e,Pass:o}})}function x(){let e=t.CID;n.$router.push({path:"/Admin/ContestEdit/UpdateContest",query:{CID:e}})}return Vue.onMounted(()=>{y()}),Vue.onUnmounted(()=>{clearInterval(u.timmer)}),(e,o)=>{const r=Vue.resolveComponent("el-empty"),i=Vue.resolveComponent("el-row"),s=Vue.resolveComponent("el-progress"),a=Vue.resolveComponent("Histogram"),V=Vue.resolveComponent("el-icon"),v=Vue.resolveComponent("el-button"),F=Vue.resolveComponent("DataAnalysis"),S=Vue.resolveComponent("Edit"),w=Vue.resolveComponent("el-divider"),A=Vue.resolveComponent("Unlock"),P=Vue.resolveComponent("el-main"),$=Vue.resolveComponent("el-container");return Vue.openBlock(),Vue.createBlock($,{class:"mainContainer"},{default:Vue.withCtx(()=>[Vue.createVNode(P,{class:"main"},{default:Vue.withCtx(()=>[Vue.createElementVNode("div",H,[Vue.unref(m)?Vue.createCommentVNode("",!0):(Vue.openBlock(),Vue.createElementBlock(Vue.Fragment,{key:0},[Vue.withDirectives(Vue.createElementVNode("div",O,[Vue.createVNode(r,{description:"\u80A5\u80A0\u62B1\u6B49\uFF0C\u6728\u6709\u627E\u5230\u8BE5\u6BD4\u8D5B\uFF0C\u8FD4\u56DE\u91CD\u8BD5\u5427\u3002"})],512),[[Vue.vShow,Vue.unref(f)]]),Vue.unref(f)?Vue.createCommentVNode("",!0):(Vue.openBlock(),Vue.createElementBlock(Vue.Fragment,{key:0},[Vue.createElementVNode("div",j,[Vue.createVNode(i,null,{default:Vue.withCtx(()=>[Vue.createElementVNode("div",G,"\xA0#"+Vue.toDisplayString(Vue.unref(t).CID)+"\xA0",1),Vue.createElementVNode("div",J,Vue.toDisplayString(Vue.unref(t).Title),1),Vue.unref(t).Type==1?(Vue.openBlock(),Vue.createElementBlock("div",K," ICPC ")):(Vue.openBlock(),Vue.createElementBlock("div",Q," OI "))]),_:1}),Vue.createElementVNode("div",W,"\u521B\u5EFA\u8005\uFF1A"+Vue.toDisplayString(Vue.unref(t).UID),1),Vue.createElementVNode("div",X,"\u63CF\u8FF0\uFF1A"+Vue.toDisplayString(Vue.unref(t).Description),1),Vue.unref(u).status==1?(Vue.openBlock(),Vue.createElementBlock("div",Y,[Z,Vue.createTextVNode(" \u8FDB\u884C\u4E2D ")])):(Vue.openBlock(),Vue.createElementBlock("div",ee,[te,Vue.createTextVNode(" \u5DF2\u7ED3\u675F ")])),Vue.createElementVNode("div",oe," \u5269\u4F59\u65F6\u95F4:"+Vue.toDisplayString(Vue.unref(n).Utils.TimeTools.timestampToInterval(Vue.unref(u).allTime-Vue.unref(u).lostTime,2)),1),Vue.createElementVNode("div",ne,[Vue.createElementVNode("div",se,Vue.toDisplayString(Vue.unref(n).Utils.TimeTools.timestampToTime(Vue.unref(t).BeginTime)),1),Vue.createElementVNode("div",ue,Vue.toDisplayString(Vue.unref(n).Utils.TimeTools.timestampToTime(Vue.unref(t).EndTime)),1)]),Vue.createElementVNode("div",ie,[Vue.createVNode(s,{"text-inside":!0,percentage:Vue.unref(u).percent,"stroke-width":25,striped:"","striped-flow":"",duration:15,color:Vue.unref(u).color},null,8,["percentage","color"])]),Vue.createElementVNode("div",le,[Vue.createVNode(v,{class:"contestButton",onClick:o[0]||(o[0]=l=>b())},{default:Vue.withCtx(()=>[Vue.createVNode(V,{size:"16px"},{default:Vue.withCtx(()=>[Vue.createVNode(a)]),_:1}),Vue.createTextVNode("\xA0\xA0\u6392 \u540D ")]),_:1}),Vue.createVNode(v,{class:"contestButton",onClick:o[1]||(o[1]=l=>k())},{default:Vue.withCtx(()=>[Vue.createVNode(V,{size:"16px"},{default:Vue.withCtx(()=>[Vue.createVNode(F)]),_:1}),Vue.createTextVNode("\xA0\xA0\u72B6 \u6001 ")]),_:1}),Vue.unref(T)?(Vue.openBlock(),Vue.createBlock(v,{key:0,class:"contestButton",onClick:o[2]||(o[2]=l=>x())},{default:Vue.withCtx(()=>[Vue.createVNode(V,{size:"16px"},{default:Vue.withCtx(()=>[Vue.createVNode(S)]),_:1}),Vue.createTextVNode("\xA0\xA0\u7F16 \u8F91 ")]),_:1})):Vue.createCommentVNode("",!0)])],512),Vue.createVNode(w),Vue.createElementVNode("div",re,[ce,(Vue.openBlock(!0),Vue.createElementBlock(Vue.Fragment,null,Vue.renderList(Vue.unref(t).Data,(l,B)=>(Vue.openBlock(),Vue.createElementBlock("div",{class:"item",key:B},[Vue.createElementVNode("div",{class:"flag cursor_pointer",onClick:U=>E(l.PID)},Vue.toDisplayString(Vue.unref(n).Utils.TSBaseTools.numberToAlpha(B+1)),9,ae),Vue.createElementVNode("div",{class:"title cursor_pointer",onClick:U=>E(l.PID)},Vue.toDisplayString(l.Title),9,Ve),Vue.createElementVNode("div",me,[Vue.createVNode(s,{type:"circle",width:22,"stroke-width":3,percentage:l.SubmitNum==0?0:l.ACNum/l.SubmitNum*100,"show-text":!1,style:{margin:"0 10px"}},null,8,["percentage"]),Vue.createTextVNode(" "+Vue.toDisplayString(l.ACNum+"/"+l.SubmitNum),1)])]))),128))],512)],64))],64)),Vue.withDirectives(Vue.createElementVNode("div",de,[pe,Vue.createElementVNode("div",fe,[_e,Vue.createVNode(q,{modelValue:Vue.unref(c),"onUpdate:modelValue":o[3]||(o[3]=l=>Vue.isRef(c)?c.value=l:c=l),onClick:o[4]||(o[4]=l=>C(Vue.unref(t).CID,Vue.unref(c))),type:"text"},null,8,["modelValue"])]),Vue.createElementVNode("div",{class:"btn cursor_pointer",onClick:o[5]||(o[5]=l=>C(Vue.unref(t).CID,Vue.unref(c)))},[Vue.createVNode(V,null,{default:Vue.withCtx(()=>[Vue.createVNode(A)]),_:1}),Vue.createTextVNode(" \xA0\u786E\u5B9A ")])],512),[[Vue.vShow,Vue.unref(m)]])])]),_:1})]),_:1})}}});const De=R(Ce,[["__scopeId","data-v-346e9fbe"]]);export{De as default};