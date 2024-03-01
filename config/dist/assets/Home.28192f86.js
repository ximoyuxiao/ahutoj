import{u as D,a as N,_ as S}from "./index.7fe3d568.js";const f= V=>(Vue.pushScopeId("data-v-46cf24d6"),V=V(),Vue.popScopeId(),V),x={class:"Title Bold ArtFont",style:{"font-size":"22px"}},F={class:"dialog-footer"},I={class:"error"},U=f(()=>Vue.createElementVNode("div",{class:"Title Bold DarkGray"}," \u8FD1\u671F\u6BD4\u8D5B ",-1)),b={class:"contestsPreview"},P={class:"right",style:{height:"396px"}},$={class:"waitingContests"},z=["onClick"],A={class:"title"},H={class:"time"},M={class:"liveContests"},G=["onClick"],R={class:"title"},j={class:"time"},q={class:"overContests"},J=["onClick"],K=["onClick"],O={class:"time"},Q=f(()=>Vue.createElementVNode("div",{class:"Title Bold DarkGray"}," \u516C\u544A ",-1)),W={class:"notice",style:{height:"396px"}},X={class:"left"},Y={class:"noticeItem"},Z=["onClick"],ee={class:"rightTime"},te=Vue.defineComponent({__name:"Home",setup(V){const{proxy:s}=Vue.getCurrentInstance(),d=Vue.ref(!0),g=D(),v=N();var u=Vue.reactive({time:0,timer:null,isError:!1,loads:{getSystemTimeLoader:null,getContestLoader:null,init(){this.getSystemTimeLoader&&this.getSystemTimeLoader.close(),this.getContestLoader&&this.getContestLoader.close(),this.getSystemTimeLoader=null,this.getContestLoader=null}}}),i=Vue.reactive({liveList:[],waitingList:[],overList:[],showListIndex:1,init(){i.liveList=[],i.waitingList=[],i.overList=[]},show: o=>{i.showListIndex=o}}),n=Vue.reactive({noticeList:[],Selected:0,init(){n.noticeList=[],n.Selected=0},copy(o){n.Selected=0;for(let t in o){let a={ID:o[t].ID,UID:o[t].UID,Title:o[t].Title,Content:o[t].Content,CreateTime:o[t].CreatedTime,UpdateTime:o[t].UpdatedTime};n.noticeList.push(a)}console.log(n.noticeList)},SelectIdx(o){n.Selected=o}});async function C(){u.loads.init(),await E(),y(),L()}async function E(){u.loads.getSystemTimeLoader=s.elLoading({node:document.getElementsByClassName("Home")[0],text:"\u540C\u6B65\u7CFB\u7EDF\u65F6\u95F4"}),await s.getServerTime().then(o=>{let t=Date.now();if(o.time==null||Math.abs(o.time-t)<1500){u.time=t;return}u.time=o.time,u.time||(u.isError=!0),s.codeProcessor(o.code)}),u.loads.getSystemTimeLoader.close()}function y(){s.$get("api/contest/list",{Limit:50,Page:0}).then(o=>{var a,r;let t=o.data;if(t.code==0){let p=u.time;for(i.init(),t.Data.forEach(c=>{c.BeginTime>p?i.waitingList.length<5&&i.waitingList.push(c):c.EndTime>p?i.liveList.length<10&&i.liveList.push(c):i.overList.length<5&&i.overList.push(c)}); i.waitingList.length<5;)i.waitingList.push({Title:""});for(; i.liveList.length<5;)i.liveList.push({Title:""});for(; i.overList.length<5;)i.overList.push({Title:""})}s.codeProcessor((a=t==null?void 0:t.code)!=null?a:100001,(r=t==null?void 0:t.msg)!=null?r:"\u670D\u52A1\u5668\u9519\u8BEF\\\\error")})}function L(){s.$get("/api/notices").then(o=>{let t=o.data;t.code==0&&(console.log(t),n.copy(t.data))})}var m={contest: o=>{v.setContestRouterData(o.CID,o.IsPublic),s.$router.push({name:"Contest",params:{CID:o.CID}})}};return Vue.onMounted(()=>{C()}),(o, t)=>{const a=Vue.resolveComponent("Select"),r=Vue.resolveComponent("el-icon"),p=Vue.resolveComponent("el-button"),c=Vue.resolveComponent("el-dialog"),k=Vue.resolveComponent("el-empty"),h=Vue.resolveComponent("Lock"),B=Vue.resolveComponent("el-main"),w=Vue.resolveComponent("el-asider"),T=Vue.resolveComponent("el-container");return Vue.openBlock(),Vue.createElementBlock(Vue.Fragment,null,[Vue.createVNode(c,{modelValue:d.value,"onUpdate:modelValue":t[2]||(t[2]= e=>d.value=e),"close-icon":"false",center:"","append-to-body":!0,style:{"border-radius":"8px",width:"min(80%, 700px)"}},{title:Vue.withCtx(()=>[Vue.createElementVNode("div",x,Vue.toDisplayString(Vue.unref(n).noticeList[Vue.unref(n).Selected].Title),1)]),footer:Vue.withCtx(()=>[Vue.createElementVNode("span",F,[Vue.createVNode(p,{type:"primary",onClick:t[1]||(t[1]= e=>d.value=!1),style:{width:"80px"}},{default:Vue.withCtx(()=>[Vue.createVNode(r,null,{default:Vue.withCtx(()=>[Vue.createVNode(a)]),_:1})]),_:1})])]),default:Vue.withCtx(()=>[Vue.createVNode(Vue.unref(MdEditorV3),{class:"markDown",modelValue:Vue.unref(n).noticeList[Vue.unref(n).Selected].Content,"onUpdate:modelValue":t[0]||(t[0]= e=>Vue.unref(n).noticeList[Vue.unref(n).Selected].Content=e),theme:Vue.unref(g).theme>0?"light":"dark","preview-only":""},null,8,["modelValue","theme"])]),_:1},8,["modelValue"]),Vue.withDirectives(Vue.createElementVNode("div",I,[Vue.createVNode(k,{description:"\u6570\u636E\u540C\u6B65\u5931\u8D25\uFF0C\u53EF\u80FD\u662F\u7F51\u7EDC\u95EE\u9898\uFF0C\u8BF7\u7A0D\u540E\u91CD\u8BD5\uFF0C\u6216\u8005\u8054\u7CFB\u7F51\u7AD9\u8FD0\u7EF4\u4EBA\u5458\u3002"})],512),[[Vue.vShow,Vue.unref(u).isError]]),Vue.createVNode(T,{class:"Main"},{default:Vue.withCtx(()=>[Vue.createVNode(T,{class:"Top"},{default:Vue.withCtx(()=>[Vue.createVNode(B,{class:"Container",style:{height:"460px"}},{default:Vue.withCtx(()=>[U,Vue.withDirectives(Vue.createElementVNode("div",b,[Vue.createElementVNode("div",P,[Vue.createElementVNode("div",$,[(Vue.openBlock(!0),Vue.createElementBlock(Vue.Fragment,null,Vue.renderList(Vue.unref(i).waitingList,(e, l)=>(Vue.openBlock(),Vue.createElementBlock("div",{class:Vue.normalizeClass(["cursor_pointer",e.Title!=""?"item":"nothing"]),key:l,onClick: _=>{Vue.unref(m).contest(e)}},[Vue.createElementVNode("div",A,[e.IsPublic==-1&&e.Title?(Vue.openBlock(),Vue.createBlock(r,{key:0,color:"#ff3300",size:"22px"},{default:Vue.withCtx(()=>[Vue.createVNode(h)]),_:1})):Vue.createCommentVNode("",!0),Vue.createTextVNode(" "+Vue.toDisplayString(e.Title),1)]),Vue.createElementVNode("div",H,Vue.toDisplayString(e.BeginTime?Vue.unref(s).Utils.TimeTools.timestampToTime(e.BeginTime)+" - ":"")+" "+Vue.toDisplayString(e.EndTime?Vue.unref(s).Utils.TimeTools.timestampToTime(e.EndTime):""),1)],10,z))),128))]),Vue.createElementVNode("div",M,[(Vue.openBlock(!0),Vue.createElementBlock(Vue.Fragment,null,Vue.renderList(Vue.unref(i).liveList,(e, l)=>(Vue.openBlock(),Vue.createElementBlock("div",{class:Vue.normalizeClass(["cursor_pointer",e.Title!=""?"item":"nothing"]),key:l,onClick: _=>{Vue.unref(m).contest(e)}},[Vue.createElementVNode("div",R,[e.IsPublic==-1&&e.Title?(Vue.openBlock(),Vue.createBlock(r,{key:0,color:"#ff3300",size:"22px"},{default:Vue.withCtx(()=>[Vue.createVNode(h)]),_:1})):Vue.createCommentVNode("",!0),Vue.createTextVNode(" "+Vue.toDisplayString(e.Title),1)]),Vue.createElementVNode("div",j,Vue.toDisplayString(e.BeginTime?Vue.unref(s).Utils.TimeTools.timestampToTime(e.BeginTime)+" - ":"")+" "+Vue.toDisplayString(e.EndTime?Vue.unref(s).Utils.TimeTools.timestampToTime(e.EndTime):""),1)],10,G))),128))]),Vue.createElementVNode("div",q,[(Vue.openBlock(!0),Vue.createElementBlock(Vue.Fragment,null,Vue.renderList(Vue.unref(i).overList,(e, l)=>(Vue.openBlock(),Vue.createElementBlock("div",{class:Vue.normalizeClass(["cursor_pointer",e.Title!=""?"item":"nothing"]),key:l,onClick: _=>{Vue.unref(m).contest(e)}},[Vue.createElementVNode("div",{class:"title",onClick: _=>{Vue.unref(m).contest(e)}},[e.IsPublic==-1&&e.Title?(Vue.openBlock(),Vue.createBlock(r,{key:0,color:"#ff3300",size:"22px"},{default:Vue.withCtx(()=>[Vue.createVNode(h)]),_:1})):Vue.createCommentVNode("",!0),Vue.createTextVNode(" "+Vue.toDisplayString(e.Title),1)],8,K),Vue.createElementVNode("div",O,Vue.toDisplayString(e.BeginTime?Vue.unref(s).Utils.TimeTools.timestampToTime(e.BeginTime)+" - ":"")+" "+Vue.toDisplayString(e.EndTime?Vue.unref(s).Utils.TimeTools.timestampToTime(e.EndTime):""),1)],10,J))),128))])])],512),[[Vue.vShow,!Vue.unref(u).isError]])]),_:1}),Vue.createVNode(w,{class:"Container Left",style:{width:"min(40%, 400px)",height:"460px"}},{default:Vue.withCtx(()=>[Q,Vue.createElementVNode("div",W,[Vue.createElementVNode("div",X,[Vue.createElementVNode("div",Y,[(Vue.openBlock(!0),Vue.createElementBlock(Vue.Fragment,null,Vue.renderList(Vue.unref(n).noticeList,(e, l)=>(Vue.openBlock(),Vue.createElementBlock("div",{class:Vue.normalizeClass([e.Title!=""?"item":"nothing","cursor_pointer"]),key:l,onClick: _=>(Vue.unref(n).SelectIdx(l),d.value=!0)},[Vue.createElementVNode("div",null,[Vue.createTextVNode(Vue.toDisplayString(e.Title),1),Vue.createElementVNode("span",ee,Vue.toDisplayString(e.CreateTime?Vue.unref(s).Utils.TimeTools.timestampToTime(e.CreateTime):""),1)])],10,Z))),128))])])])]),_:1})]),_:1})]),_:1})],64)}}});const ie=S(te,[["__scopeId","data-v-46cf24d6"]]);export{ie as default};
