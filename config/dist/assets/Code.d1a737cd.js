import{c as U,b as A,_ as b}from "./index.7fe3d568.js";const n= m=>(Vue.pushScopeId("data-v-c2c84920"),m=m(),Vue.popScopeId(),m),R={key:0},L={key:1,class:"Main",style:{color:"white"}},z=n(()=>Vue.createElementVNode("div",{class:"Title Bold ArtFont Left"}," Judging ",-1)),P=n(()=>Vue.createElementVNode("div",{class:"Title Bold ArtFont Left"}," Rejudging ",-1)),M=n(()=>Vue.createElementVNode("div",{class:"Title Bold ArtFont Left"}," Pending ",-1)),G=n(()=>Vue.createElementVNode("div",{class:"Title Bold ArtFont Left"}," Failed ",-1)),$=n(()=>Vue.createElementVNode("div",{class:"Title Bold ArtFont Left"}," Compile Error ",-1)),j=n(()=>Vue.createElementVNode("div",{class:"Title Bold ArtFont Left"}," Wrong Answer ",-1)),J={class:"Left",style:{"font-size":"16px","margin-top":"4px"}},q=n(()=>Vue.createElementVNode("div",{class:"Title Bold ArtFont Left"}," Time Limit Exceeded ",-1)),W={class:"Left",style:{"font-size":"16px","margin-top":"4px"}},O=n(()=>Vue.createElementVNode("div",{class:"Title Bold ArtFont Left"}," Memory Limit Exceeded ",-1)),Q={class:"Left",style:{"font-size":"16px","margin-top":"4px"}},H=n(()=>Vue.createElementVNode("div",{class:"Title Bold ArtFont Left"}," Partical Error ",-1)),K={class:"Left",style:{"font-size":"16px","margin-top":"4px"}},X=n(()=>Vue.createElementVNode("div",{class:"Title Bold ArtFont Left"}," Output Limit Exceeded ",-1)),Y={class:"Left",style:{"font-size":"16px","margin-top":"4px"}},Z=n(()=>Vue.createElementVNode("div",{class:"Title Bold ArtFont Left"}," Runtime Error ",-1)),ee={class:"Left",style:{"font-size":"16px","margin-top":"4px"}},te=n(()=>Vue.createElementVNode("div",{class:"Title Bold ArtFont Left"}," \u672C\u9898\u6570\u636E\u7F3A\u5931\uFF0C\u8BF7\u8054\u7CFB\u7BA1\u7406\u5458 ",-1)),ue=n(()=>Vue.createElementVNode("div",{class:"Title Bold ArtFont Left"}," Accepted ",-1)),oe={class:"Left",style:{"font-size":"16px","margin-top":"4px"}},le=n(()=>Vue.createElementVNode("div",{class:"Title Bold ArtFont Left"}," Unknown Return Value ",-1)),ne={class:"Left",style:{"font-size":"16px","margin-top":"4px"}},ae=n(()=>Vue.createElementVNode("div",{class:"title"},"\u9519\u8BEF\u4FE1\u606F",-1)),se=n(()=>Vue.createElementVNode("div",{class:"title"},"\u683C\u5F0F\u6709\u8BEF",-1)),re=n(()=>Vue.createElementVNode("div",{class:"title"},"\u610F\u5916\u7684\u9519\u8BEF",-1)),Ve={class:"Main Bottom"},ie={id:"judging",class:"Container",style:{width:"100%",height:"200px"}},ce=Vue.defineComponent({__name:"Code",setup(m){const{proxy:t}=Vue.getCurrentInstance(),p=U(),x=A();var y=Vue.ref(!0),f=Vue.ref(!1),C=null,c=Vue.reactive({SID:-1,loading:null,init(){c.SID=-1,c.loading=null}}),e=Vue.reactive({UID:"",Lang:1,PID:0,Result:"",SID:1,Source:"",SubmitTime:0,UseMemory:0,UseTime:0,CeInfo:"",hasCeInfo:!1,PassSample:0,SampleNumber:0,autoUpdate:null,updateTimeStep:1,init(){e.UID="",e.Lang=1,e.PID=0,e.Result="",e.SID=1,e.Source="",e.SubmitTime=0,e.UseMemory=0,e.UseTime=0,e.CeInfo="",e.hasCeInfo=!1,e.PassSample=0,e.SampleNumber=0,clearTimeout(e.autoUpdate)},copy(u){e.UID=u.UID,e.Lang=u.Lang,e.PID=u.PID,e.Result=u.Result,e.SID=u.SID,e.Source=u.Source,e.SubmitTime=u.SubmitTime,e.UseMemory=u.UseMemory,e.UseTime=u.UseTime,e.CeInfo=u.CeInfo,e.PassSample=u.PassSample,e.SampleNumber=u.SampleNumber,u.CeInfo&&u.CeInfo!=""&&(e.hasCeInfo=!0),u.Result=="JUDGING"||u.Result=="REJUDGING"||u.Result=="PENDING"?(f.value=!0,C||(C=t.elLoading({node:document.getElementById("judging"),text:"\u6B63\u5728\u5224\u9898\u4E2D..."})),e.autoUpdate=setTimeout(()=>{e.updateTimeStep<15?(N(),e.updateTimeStep++):(x.needPing(),t.elMessage({message:"\u7F51\u7EDC\u53EF\u80FD\u51FA\u73B0\u95EE\u9898\uFF0C\u6216\u8005\u670D\u52A1\u5668\u7E41\u5FD9\uFF0C\u8BF7\u5237\u65B0\u6216\u7A0D\u540E\u518D\u8BD5\u3002",type:"warning"}))},e.updateTimeStep/3*1e3+500)):(f.value&&u.Result=="AC"&&t.elNotification({message:"\u606D\u559C\u4F60\uFF0C\u901A\u8FC7\u4E86\u8BE5\u9898\uFF01",type:"success"}),f.value=!1,C&&C.close())}});function N(){c.loading=t.elLoading({node:t.$refs.code}),c.SID!=-1&&t.$get("api/submit/"+c.SID).then(u=>{var r,V;let o=u.data;o.code==0&&(e.init(),e.copy(o),y.value=!1),c.loading.close(),t.codeProcessor((r=o==null?void 0:o.code)!=null?r:100001,(V=o==null?void 0:o.msg)!=null?V:"\u670D\u52A1\u5668\u9519\u8BEF\\\\error")})}function B(){let u={};if(e.SID)u.SID=e.SID;else{t.elMessage({message:"\u6570\u636E\u5F02\u5E38\uFF0C\u8BF7\u5237\u65B0\u540E\u91CD\u8BD5",type:"error"});return}t.$post("api/submit/rejudge/",u).then(o=>{var V,l;console.log(o);let r=o.data;r.code==0&&(N(),t.elNotification({message:"\u91CD\u5224\u6210\u529F!",type:"success"})),t.codeProcessor((V=r==null?void 0:r.code)!=null?V:100001,(l=r==null?void 0:r.msg)!=null?l:"\u670D\u52A1\u5668\u9519\u8BEF\\\\error")})}function h(u){let o=t.$route.query.CID?"ContestProblem":"Problem";t.$router.push({name:o,params:{PID:u,CID:t.$route.query.CID?t.$route.query.CID:void 0}})}Vue.onMounted(()=>{t.$route.query.SID&&(c.SID=Number(t.$route.query.SID)),N()}),Vue.onUnmounted(()=>{clearTimeout(e.autoUpdate)});const w=[{}];return(u, o)=>{const r=Vue.resolveComponent("el-empty"),V=Vue.resolveComponent("Loading"),l=Vue.resolveComponent("el-icon"),s=Vue.resolveComponent("el-row"),a=Vue.resolveComponent("el-header"),_=Vue.resolveComponent("WarningFilled"),E=Vue.resolveComponent("CircleCloseFilled"),D=Vue.resolveComponent("RemoveFilled"),F=Vue.resolveComponent("DocumentDelete"),k=Vue.resolveComponent("SuccessFilled"),I=Vue.resolveComponent("QuestionFilled"),d=Vue.resolveComponent("el-table-column"),v=Vue.resolveComponent("el-table"),S=Vue.resolveComponent("el-main"),g=Vue.resolveComponent("el-input"),T=Vue.resolveComponent("el-container");return Vue.openBlock(),Vue.createElementBlock(Vue.Fragment,null,[Vue.unref(y)?(Vue.openBlock(),Vue.createElementBlock("div",R,[Vue.createVNode(r,{description:"\u65E0\u7ED3\u679C"})])):(Vue.openBlock(),Vue.createElementBlock("div",L,[Vue.createVNode(T,{class:"Top Bottom Main"},{default:Vue.withCtx(()=>[Vue.unref(e).Result=="JUDGING"?(Vue.openBlock(),Vue.createBlock(a,{key:0,class:"Container StatusBarJUDGING",style:Vue.normalizeStyle("background-color:"+Vue.unref(t).Utils.StatusConstValManager.getStatusColor(Vue.unref(e).Result))},{default:Vue.withCtx(()=>[Vue.createVNode(s,null,{default:Vue.withCtx(()=>[Vue.createVNode(l,{size:"26px",class:"is-loading"},{default:Vue.withCtx(()=>[Vue.createVNode(V)]),_:1}),z]),_:1})]),_:1},8,["style"])):Vue.unref(e).Result=="REJUDGING"?(Vue.openBlock(),Vue.createBlock(a,{key:1,class:"Container StatusBarJUDGING",style:Vue.normalizeStyle("background-color:"+Vue.unref(t).Utils.StatusConstValManager.getStatusColor(Vue.unref(e).Result))},{default:Vue.withCtx(()=>[Vue.createVNode(s,null,{default:Vue.withCtx(()=>[Vue.createVNode(l,{size:"26px",class:"is-loading"},{default:Vue.withCtx(()=>[Vue.createVNode(V)]),_:1}),P]),_:1})]),_:1},8,["style"])):Vue.unref(e).Result=="PENDING"?(Vue.openBlock(),Vue.createBlock(a,{key:2,class:"Container StatusBarPENDING",style:Vue.normalizeStyle("background-color:"+Vue.unref(t).Utils.StatusConstValManager.getStatusColor(Vue.unref(e).Result))},{default:Vue.withCtx(()=>[Vue.createVNode(s,null,{default:Vue.withCtx(()=>[Vue.createVNode(l,{size:"26px",class:"is-loading"},{default:Vue.withCtx(()=>[Vue.createVNode(V)]),_:1}),M]),_:1})]),_:1},8,["style"])):Vue.unref(e).Result=="FAILED"?(Vue.openBlock(),Vue.createBlock(a,{key:3,class:"Container StatusBarAC",style:Vue.normalizeStyle("background-color:"+Vue.unref(t).Utils.StatusConstValManager.getStatusColor(Vue.unref(e).Result))},{default:Vue.withCtx(()=>[Vue.createVNode(s,null,{default:Vue.withCtx(()=>[Vue.createVNode(l,{size:"26px"},{default:Vue.withCtx(()=>[Vue.createVNode(_)]),_:1}),G]),_:1})]),_:1},8,["style"])):Vue.unref(e).Result=="CE"?(Vue.openBlock(),Vue.createBlock(a,{key:4,class:"Container StatusBarWA",style:Vue.normalizeStyle("background-color:"+Vue.unref(t).Utils.StatusConstValManager.getStatusColor(Vue.unref(e).Result))},{default:Vue.withCtx(()=>[Vue.createVNode(s,null,{default:Vue.withCtx(()=>[Vue.createVNode(l,{size:"26px"},{default:Vue.withCtx(()=>[Vue.createVNode(_)]),_:1}),$]),_:1})]),_:1},8,["style"])):Vue.unref(e).Result=="WA"?(Vue.openBlock(),Vue.createBlock(a,{key:5,class:"Container",style:Vue.normalizeStyle("background-color:"+Vue.unref(t).Utils.StatusConstValManager.getStatusColor(Vue.unref(e).Result))},{default:Vue.withCtx(()=>[Vue.createVNode(s,null,{default:Vue.withCtx(()=>[Vue.createVNode(l,{size:"26px"},{default:Vue.withCtx(()=>[Vue.createVNode(E)]),_:1}),j,Vue.createElementVNode("div",J,"\u901A\u8FC7\u4E86 "+Vue.toDisplayString(Vue.unref(e).PassSample)+" / "+Vue.toDisplayString(Vue.unref(e).SampleNumber)+" \u4E2A\u6D4B\u8BD5\u6570\u636E",1)]),_:1})]),_:1},8,["style"])):Vue.unref(e).Result=="TLE"?(Vue.openBlock(),Vue.createBlock(a,{key:6,class:"Container",style:Vue.normalizeStyle("background-color:"+Vue.unref(t).Utils.StatusConstValManager.getStatusColor(Vue.unref(e).Result))},{default:Vue.withCtx(()=>[Vue.createVNode(s,null,{default:Vue.withCtx(()=>[Vue.createVNode(l,{size:"26px"},{default:Vue.withCtx(()=>[Vue.createVNode(_)]),_:1}),q,Vue.createElementVNode("div",W,"\u901A\u8FC7\u4E86 "+Vue.toDisplayString(Vue.unref(e).PassSample)+" / "+Vue.toDisplayString(Vue.unref(e).SampleNumber)+" \u4E2A\u6D4B\u8BD5\u6570\u636E",1)]),_:1})]),_:1},8,["style"])):Vue.unref(e).Result=="MLE"?(Vue.openBlock(),Vue.createBlock(a,{key:7,class:"Container",style:Vue.normalizeStyle("background-color:"+Vue.unref(t).Utils.StatusConstValManager.getStatusColor(Vue.unref(e).Result))},{default:Vue.withCtx(()=>[Vue.createVNode(s,null,{default:Vue.withCtx(()=>[Vue.createVNode(l,{size:"26px"},{default:Vue.withCtx(()=>[Vue.createVNode(_)]),_:1}),O,Vue.createElementVNode("div",Q,"\u901A\u8FC7\u4E86 "+Vue.toDisplayString(Vue.unref(e).PassSample)+" / "+Vue.toDisplayString(Vue.unref(e).SampleNumber)+" \u4E2A\u6D4B\u8BD5\u6570\u636E",1)]),_:1})]),_:1},8,["style"])):Vue.unref(e).Result=="PE"?(Vue.openBlock(),Vue.createBlock(a,{key:8,class:"Container",style:Vue.normalizeStyle("background-color:"+Vue.unref(t).Utils.StatusConstValManager.getStatusColor(Vue.unref(e).Result))},{default:Vue.withCtx(()=>[Vue.createVNode(s,null,{default:Vue.withCtx(()=>[Vue.createVNode(l,{size:"26px"},{default:Vue.withCtx(()=>[Vue.createVNode(E)]),_:1}),H,Vue.createElementVNode("div",K,"\u901A\u8FC7\u4E86 "+Vue.toDisplayString(Vue.unref(e).PassSample)+" / "+Vue.toDisplayString(Vue.unref(e).SampleNumber)+" \u4E2A\u6D4B\u8BD5\u6570\u636E",1)]),_:1})]),_:1},8,["style"])):Vue.unref(e).Result=="OLE"?(Vue.openBlock(),Vue.createBlock(a,{key:9,class:"Container",style:Vue.normalizeStyle("background-color:"+Vue.unref(t).Utils.StatusConstValManager.getStatusColor(Vue.unref(e).Result))},{default:Vue.withCtx(()=>[Vue.createVNode(s,null,{default:Vue.withCtx(()=>[Vue.createVNode(l,{size:"26px"},{default:Vue.withCtx(()=>[Vue.createVNode(D)]),_:1}),X,Vue.createElementVNode("div",Y,"\u901A\u8FC7\u4E86 "+Vue.toDisplayString(Vue.unref(e).PassSample)+" / "+Vue.toDisplayString(Vue.unref(e).SampleNumber)+" \u4E2A\u6D4B\u8BD5\u6570\u636E",1)]),_:1})]),_:1},8,["style"])):Vue.unref(e).Result=="RE"?(Vue.openBlock(),Vue.createBlock(a,{key:10,class:"Container",style:Vue.normalizeStyle("background-color:"+Vue.unref(t).Utils.StatusConstValManager.getStatusColor(Vue.unref(e).Result))},{default:Vue.withCtx(()=>[Vue.createVNode(s,null,{default:Vue.withCtx(()=>[Vue.createVNode(l,{size:"26px"},{default:Vue.withCtx(()=>[Vue.createVNode(D)]),_:1}),Z,Vue.createElementVNode("div",ee,"\u901A\u8FC7\u4E86 "+Vue.toDisplayString(Vue.unref(e).PassSample)+" / "+Vue.toDisplayString(Vue.unref(e).SampleNumber)+" \u4E2A\u6D4B\u8BD5\u6570\u636E",1)]),_:1})]),_:1},8,["style"])):Vue.unref(e).SampleNumber==0?(Vue.openBlock(),Vue.createBlock(a,{key:11,class:"Container StatusBarJUDGING",style:{"background-color":"#409EFF"}},{default:Vue.withCtx(()=>[Vue.createVNode(s,null,{default:Vue.withCtx(()=>[Vue.createVNode(l,{size:"26px"},{default:Vue.withCtx(()=>[Vue.createVNode(F)]),_:1}),te]),_:1})]),_:1})):Vue.unref(e).Result=="AC"?(Vue.openBlock(),Vue.createBlock(a,{key:12,class:"Container StatusBarAC",style:Vue.normalizeStyle("background-color:"+Vue.unref(t).Utils.StatusConstValManager.getStatusColor(Vue.unref(e).Result))},{default:Vue.withCtx(()=>[Vue.createVNode(s,null,{default:Vue.withCtx(()=>[Vue.createVNode(l,{size:"26px"},{default:Vue.withCtx(()=>[Vue.createVNode(k)]),_:1}),ue,Vue.createElementVNode("div",oe,"\u901A\u8FC7\u4E86 "+Vue.toDisplayString(Vue.unref(e).PassSample)+" / "+Vue.toDisplayString(Vue.unref(e).SampleNumber)+" \u4E2A\u6D4B\u8BD5\u6570\u636E",1)]),_:1})]),_:1},8,["style"])):(Vue.openBlock(),Vue.createBlock(a,{key:13,class:"Container",style:{"background-color":"#1E1E1E"}},{default:Vue.withCtx(()=>[Vue.createVNode(s,null,{default:Vue.withCtx(()=>[Vue.createVNode(l,{size:"26px"},{default:Vue.withCtx(()=>[Vue.createVNode(I)]),_:1}),le,Vue.createElementVNode("div",ne,"\u901A\u8FC7\u4E86 "+Vue.toDisplayString(Vue.unref(e).PassSample)+" / "+Vue.toDisplayString(Vue.unref(e).SampleNumber)+" \u4E2A\u6D4B\u8BD5\u6570\u636E",1)]),_:1})]),_:1})),Vue.createVNode(a,{class:"Container Top"},{default:Vue.withCtx(()=>[Vue.createVNode(v,{data:w,style:{width:"100%"}},{default:Vue.withCtx(()=>[Vue.createVNode(d,{prop:"SID",label:"\u63D0\u4EA4 ID",align:"center"},{default:Vue.withCtx(()=>[Vue.createTextVNode(Vue.toDisplayString(Vue.unref(e).SID),1)]),_:1}),Vue.createVNode(d,{prop:"UID",label:"\u7528\u6237",align:"center"},{default:Vue.withCtx(()=>[Vue.createTextVNode(Vue.toDisplayString(Vue.unref(e).UID.length>15?Vue.unref(e).UID.slice(0,15)+"...":Vue.unref(e).UID),1)]),_:1}),Vue.createVNode(d,{prop:"SubmitTime",label:"\u63D0\u4EA4\u65F6\u95F4","min-width":"170px",align:"center"},{default:Vue.withCtx(()=>[Vue.createTextVNode(Vue.toDisplayString(Vue.unref(t).Utils.TimeTools.timestampToTime(Vue.unref(e).SubmitTime)),1)]),_:1}),Vue.createVNode(d,{prop:"PID",label:"\u9898\u76EE",align:"center"},{default:Vue.withCtx(()=>[Vue.createElementVNode("div",{onClick:o[0]||(o[0]= i=>h(Vue.unref(e).PID)),class:"cursor_pointer Bold ArtFont",style:{color:"#569CD6"}},Vue.toDisplayString(Vue.unref(e).PID),1)]),_:1}),Vue.createVNode(d,{prop:"Time",label:"\u7528\u65F6",align:"center",style:Vue.normalizeStyle("width: 120px;"+(Vue.unref(e).Result=="TLE"?"color: #ff381e;":""))},{default:Vue.withCtx(()=>[Vue.createTextVNode(Vue.toDisplayString(Vue.unref(e).UseTime)+"\xA0ms ",1)]),_:1},8,["style"]),Vue.createVNode(d,{prop:"Mem",label:"\u5185\u5B58",align:"center",style:Vue.normalizeStyle("width: 120px;"+(Vue.unref(e).Result=="MLE"?"color: #ff381e;":""))},{default:Vue.withCtx(()=>[Vue.createTextVNode(Vue.toDisplayString((Vue.unref(e).UseMemory/1024/1024).toFixed(0))+"\xA0MB ",1)]),_:1},8,["style"]),Vue.createVNode(d,{prop:"Lang",label:"\u8BED\u8A00",align:"center"},{default:Vue.withCtx(()=>[Vue.createTextVNode(Vue.toDisplayString(Vue.unref(t).Utils.StatusConstValManager.getLangString(Vue.unref(e).Lang)),1)]),_:1})]),_:1})]),_:1}),Vue.unref(y)?(Vue.openBlock(),Vue.createBlock(S,{key:14,class:"Container"},{default:Vue.withCtx(()=>[Vue.createVNode(r,{description:"\u65E0\u7ED3\u679C"})]),_:1})):Vue.unref(e).hasCeInfo?(Vue.openBlock(),Vue.createBlock(S,{key:15,class:"Container"},{default:Vue.withCtx(()=>[ae,Vue.createVNode(g,{modelValue:Vue.unref(e).CeInfo,"onUpdate:modelValue":o[1]||(o[1]= i=>Vue.unref(e).CeInfo=i),autosize:{minRows:5},readonly:"",resize:"none","show-word-limit":"",type:"textarea"},null,8,["modelValue"])]),_:1})):Vue.unref(e).Result=="PE"?(Vue.openBlock(),Vue.createBlock(S,{key:16,class:"Container"},{default:Vue.withCtx(()=>[se,Vue.createVNode(g,{modelValue:Vue.unref(p).SUBMIT_RESULT_PE,"onUpdate:modelValue":o[2]||(o[2]= i=>Vue.unref(p).SUBMIT_RESULT_PE=i),autosize:{minRows:5},readonly:"",resize:"none","show-word-limit":"",type:"textarea"},null,8,["modelValue"])]),_:1})):Vue.unref(e).Result=="FAILED"?(Vue.openBlock(),Vue.createBlock(S,{key:17,class:"Container"},{default:Vue.withCtx(()=>[re,Vue.createVNode(g,{modelValue:Vue.unref(p).SUBMIT_RESULT_FAILED,"onUpdate:modelValue":o[3]||(o[3]= i=>Vue.unref(p).SUBMIT_RESULT_FAILED=i),autosize:{minRows:3},readonly:"",resize:"none","show-word-limit":"",type:"textarea"},null,8,["modelValue"]),Vue.createElementVNode("div",{class:"rejudge cursor_pointer",onClick:B}," \u91CD\u5224 ")]),_:1})):Vue.createCommentVNode("",!0),Vue.createVNode(g,{class:"Top",modelValue:Vue.unref(e).Source,"onUpdate:modelValue":o[4]||(o[4]= i=>Vue.unref(e).Source=i),autosize:{minRows:5},readonly:"",resize:"none","show-word-limit":"",type:"textarea",style:{"border-radius":"0px"}},null,8,["modelValue"])]),_:1})])),Vue.createElementVNode("div",Ve,[Vue.withDirectives(Vue.createElementVNode("div",ie,null,512),[[Vue.vShow,Vue.unref(f)]])])],64)}}});const me=b(ce,[["__scopeId","data-v-c2c84920"]]);export{me as default};