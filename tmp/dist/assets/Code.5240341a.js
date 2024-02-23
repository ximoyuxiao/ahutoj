import{c as _,b as y,_ as S}from"./index.de99a551.js";const V=d=>(Vue.pushScopeId("data-v-f1c247d1"),d=d(),Vue.popScopeId(),d),E={class:"code",ref:"code"},I={class:"submit"},g=Vue.createStaticVNode('<div class="header" data-v-f1c247d1><div style="width:140px;" data-v-f1c247d1>\u63D0\u4EA4ID</div><div style="width:120px;" data-v-f1c247d1>\u9898\u76EEID</div><div style="width:160px;" data-v-f1c247d1>\u7528\u6237</div><div style="width:120px;" data-v-f1c247d1>\u72B6\u6001</div><div style="width:140px;" data-v-f1c247d1>\u8BED\u8A00</div><div style="width:120px;" data-v-f1c247d1>\u7528\u65F6</div><div style="width:120px;" data-v-f1c247d1>\u5185\u5B58</div><div style="width:180px;" data-v-f1c247d1>\u63D0\u4EA4\u65F6\u95F4</div></div>',1),C={class:"item"},h={class:"SID",style:{width:"140px"}},F={class:"UID",style:{width:"160px"}},U={style:{width:"120px",display:"flex","justify-content":"center"}},w={style:{width:"140px"}},N={class:"submitTime",style:{width:"180px"}},B={key:0,class:"ce"},T=V(()=>Vue.createElementVNode("div",{class:"title"},"\u9519\u8BEF\u4FE1\u606F",-1)),x={key:1,class:"pe"},R=V(()=>Vue.createElementVNode("div",{class:"title"},"\u683C\u5F0F\u6709\u8BEF",-1)),L={key:2,class:"failed"},b=V(()=>Vue.createElementVNode("div",{class:"title"},"\u610F\u5916\u7684\u9519\u8BEF",-1)),k=V(()=>Vue.createElementVNode("div",{class:"title"},"\u4EE3\u7801",-1)),P={id:"judging"},M={class:"notFound"},A=Vue.defineComponent({__name:"Code",setup(d){const{proxy:o}=Vue.getCurrentInstance(),a=_(),p=y();var m=Vue.ref(!0),r=Vue.ref(!1),c=null,i=Vue.reactive({SID:-1,loading:null,init(){i.SID=-1,i.loading=null}}),e=Vue.reactive({UID:"",Lang:1,PID:0,Result:"",SID:1,Source:"",SubmitTime:0,UseMemory:0,UseTime:0,CeInfo:"",hasCeInfo:!1,autoUpdate:null,updateTimeStep:1,init(){e.UID="",e.Lang=1,e.PID=0,e.Result="",e.SID=1,e.Source="",e.SubmitTime=0,e.UseMemory=0,e.UseTime=0,e.CeInfo="",e.hasCeInfo=!1,clearTimeout(e.autoUpdate)},copy(u){e.UID=u.UID,e.Lang=u.Lang,e.PID=u.PID,e.Result=u.Result,e.SID=u.SID,e.Source=u.Source,e.SubmitTime=u.SubmitTime,e.UseMemory=u.UseMemory,e.UseTime=u.UseTime,e.CeInfo=u.CeInfo,u.CeInfo&&u.CeInfo!=""&&(e.hasCeInfo=!0),u.Result=="JUDGING"||u.Result=="REJUDGING"||u.Result=="PENDING"?(r.value=!0,c||(c=o.elLoading({node:document.getElementById("judging"),text:"\u6B63\u5728\u5224\u9898\u4E2D..."})),e.autoUpdate=setTimeout(()=>{e.updateTimeStep<15?(f(),e.updateTimeStep++):(p.needPing(),o.elMessage({message:"\u7F51\u7EDC\u53EF\u80FD\u51FA\u73B0\u95EE\u9898\uFF0C\u6216\u8005\u670D\u52A1\u5668\u7E41\u5FD9\uFF0C\u8BF7\u5237\u65B0\u6216\u7A0D\u540E\u518D\u8BD5\u3002",type:"warning"}))},e.updateTimeStep/3*1e3+500)):(r.value&&u.Result=="AC"&&o.elNotification({message:"\u606D\u559C\u4F60\uFF0C\u901A\u8FC7\u4E86\u8BE5\u9898\uFF01",type:"success"}),r.value=!1,c&&c.close())}});function f(){i.loading=o.elLoading({node:o.$refs.code}),i.SID!=-1&&o.$get("api/submit/"+i.SID).then(u=>{var s,l;let t=u.data;t.code==0&&(e.init(),e.copy(t),m.value=!1),i.loading.close(),o.codeProcessor((s=t==null?void 0:t.code)!=null?s:100001,(l=t==null?void 0:t.msg)!=null?l:"\u670D\u52A1\u5668\u9519\u8BEF\\\\error")})}function v(){let u={};if(e.SID)u.SID=e.SID;else{o.elMessage({message:"\u6570\u636E\u5F02\u5E38\uFF0C\u8BF7\u5237\u65B0\u540E\u91CD\u8BD5",type:"error"});return}o.$post("api/submit/rejudge/",u).then(t=>{var l,n;console.log(t);let s=t.data;s.code==0&&(f(),o.elNotification({message:"\u91CD\u5224\u6210\u529F!",type:"success"})),o.codeProcessor((l=s==null?void 0:s.code)!=null?l:100001,(n=s==null?void 0:s.msg)!=null?n:"\u670D\u52A1\u5668\u9519\u8BEF\\\\error")})}function D(u){let t=o.$route.query.CID?"ContestProblem":"Problem";o.$router.push({name:t,params:{PID:u,CID:o.$route.query.CID?o.$route.query.CID:void 0}})}return Vue.onMounted(()=>{o.$route.query.SID&&(i.SID=Number(o.$route.query.SID)),f()}),Vue.onUnmounted(()=>{clearTimeout(e.autoUpdate)}),(u,t)=>{const s=Vue.resolveComponent("el-input"),l=Vue.resolveComponent("el-empty");return Vue.openBlock(),Vue.createElementBlock("div",E,[Vue.withDirectives(Vue.createElementVNode("div",I,[g,Vue.createElementVNode("div",C,[Vue.createElementVNode("div",h,Vue.toDisplayString(Vue.unref(e).SID),1),Vue.createElementVNode("div",{class:"PID cursor_pointer",style:{width:"120px"},onClick:t[0]||(t[0]=n=>D(Vue.unref(e).PID))},Vue.toDisplayString(Vue.unref(e).PID),1),Vue.createElementVNode("div",F,Vue.toDisplayString(Vue.unref(e).UID.length>15?Vue.unref(e).UID.slice(0,15)+"...":Vue.unref(e).UID),1),Vue.createElementVNode("div",U,[Vue.createElementVNode("div",{class:"res cursor_pointer",style:Vue.normalizeStyle("color: #ffffff; background-color:"+Vue.unref(o).Utils.StatusConstValManager.getStatusColor(Vue.unref(e).Result))},Vue.toDisplayString(Vue.unref(e).Result),5)]),Vue.createElementVNode("div",w,Vue.toDisplayString(Vue.unref(o).Utils.StatusConstValManager.getLangString(Vue.unref(e).Lang)),1),Vue.createElementVNode("div",{class:"useTime",style:Vue.normalizeStyle("width: 120px;"+(Vue.unref(e).Result=="TLE"?"color: #ff381e;":""))},Vue.toDisplayString(Vue.unref(e).UseTime)+"\xA0ms ",5),Vue.createElementVNode("div",{class:"useMemory",style:Vue.normalizeStyle("width: 120px;"+(Vue.unref(e).Result=="MLE"?"color: #ff381e;":""))},Vue.toDisplayString((Vue.unref(e).UseMemory/1024).toFixed(0))+"\xA0KB ",5),Vue.createElementVNode("div",N,Vue.toDisplayString(Vue.unref(o).Utils.TimeTools.timestampToTime(Vue.unref(e).SubmitTime)),1)]),Vue.unref(r)?Vue.createCommentVNode("",!0):(Vue.openBlock(),Vue.createElementBlock(Vue.Fragment,{key:0},[Vue.unref(e).hasCeInfo?(Vue.openBlock(),Vue.createElementBlock("div",B,[T,Vue.createVNode(s,{modelValue:Vue.unref(e).CeInfo,"onUpdate:modelValue":t[1]||(t[1]=n=>Vue.unref(e).CeInfo=n),autosize:{minRows:5},readonly:"",resize:"none","show-word-limit":"",type:"textarea"},null,8,["modelValue"])])):Vue.createCommentVNode("",!0),Vue.unref(e).Result=="PE"?(Vue.openBlock(),Vue.createElementBlock("div",x,[R,Vue.createVNode(s,{modelValue:Vue.unref(a).SUBMIT_RESULT_PE,"onUpdate:modelValue":t[2]||(t[2]=n=>Vue.unref(a).SUBMIT_RESULT_PE=n),autosize:{minRows:5},readonly:"",resize:"none","show-word-limit":"",type:"textarea"},null,8,["modelValue"])])):Vue.createCommentVNode("",!0),Vue.unref(e).Result=="FAILED"?(Vue.openBlock(),Vue.createElementBlock("div",L,[b,Vue.createVNode(s,{modelValue:Vue.unref(a).SUBMIT_RESULT_FAILED,"onUpdate:modelValue":t[3]||(t[3]=n=>Vue.unref(a).SUBMIT_RESULT_FAILED=n),autosize:{minRows:3},readonly:"",resize:"none","show-word-limit":"",type:"textarea"},null,8,["modelValue"]),Vue.createElementVNode("div",{class:"rejudge cursor_pointer",onClick:v}," \u91CD\u5224 ")])):Vue.createCommentVNode("",!0),k,Vue.createVNode(s,{modelValue:Vue.unref(e).Source,"onUpdate:modelValue":t[4]||(t[4]=n=>Vue.unref(e).Source=n),autosize:{minRows:5},readonly:"",resize:"none","show-word-limit":"",type:"textarea"},null,8,["modelValue"])],64)),Vue.withDirectives(Vue.createElementVNode("div",P,null,512),[[Vue.vShow,Vue.unref(r)]])],512),[[Vue.vShow,!Vue.unref(m)]]),Vue.withDirectives(Vue.createElementVNode("div",M,[Vue.createVNode(l,{description:"\u65E0\u7ED3\u679C"})],512),[[Vue.vShow,Vue.unref(m)]])],512)}}});const $=S(A,[["__scopeId","data-v-f1c247d1"]]);export{$ as default};
