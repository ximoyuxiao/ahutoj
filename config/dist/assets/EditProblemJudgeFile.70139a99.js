import{f as w,_ as y}from "./index.7fe3d568.js";const p= V=>(Vue.pushScopeId("data-v-917d811d"),V=V(),Vue.popScopeId(),V),P=p(()=>Vue.createElementVNode("span",null,"\u9898\u53F7",-1)),A={class:"preview Center"},k={class:"fileList"},I=p(()=>Vue.createElementVNode("div",{class:"header"},[Vue.createElementVNode("div",null,"\u6587\u4EF6\u540D"),Vue.createElementVNode("div",null,"\u7C7B\u578B"),Vue.createElementVNode("div",null,"\u5927\u5C0F"),Vue.createElementVNode("div",null,"\u64CD\u4F5C")],-1)),S={id:"notFound"},b=p(()=>Vue.createElementVNode("div",{class:"el-upload__text"}," \u4EE5 zip/in/out \u683C\u5F0F\u6587\u4EF6\u4E0A\u4F20\uFF0C\u70B9\u51FB\u6216\u8005\u62D6\u62FD ",-1)),M=Vue.defineComponent({__name:"EditProblemJudgeFile",setup(V){const{proxy:l}=Vue.getCurrentInstance();var d=Vue.ref([]),i=Vue.ref([]),n=Vue.reactive({PID:"",isSearched:!0,onFocus(){n.isSearched=!0},getProblem(u){u&&(n.PID=u),l.$get("api/file/"+n.PID).then(t=>{i.value=[],d.value=[];let e=t.data;e.code==0&&(i.value=e.Data),e.Data||l.elMessage({message:"\u6682\u65E0\u6587\u4EF6",type:"info"}),n.isSearched=!0})}});function F(u, t){d.value=[],t.forEach(a=>{d.value.push(a)});let e=t[t.length-1].name;i.value.forEach(a=>{if(e==a.Filename){l.elMessage({message:a.Filename+" \u5DF2\u5B58\u5728\uFF0C\u5C06\u4F1A\u8986\u76D6\u539F\u6587\u4EF6!",type:"warning"});return}})}function f(u, t){d.value=[],t.forEach(e=>{d.value.push(e)}),l.elMessage({message:"\u53D6\u6D88\u4E0A\u4F20 "+u.name,type:"success"})}function v(){if(d.value.length==0){l.elMessage({message:"\u4E0A\u4F20\u5217\u8868\u4E3A\u7A7A\uFF01",type:"warning"});return}ElementPlus.ElMessageBox.confirm("\u786E\u5B9A\u4E0A\u4F20\u5417\uFF1F","\u6CE8\u610F",{confirmButtonText:"\u786E\u5B9A",cancelButtonText:"\u53D6\u6D88",type:"warning"}).then(()=>{d.value.forEach(u=>{u=u.raw,w.uploadProblemJudgeFile(u,n.PID).then(t=>{var a,o;let e=t.data;if(e.code==0){for(let s in i.value)if(i.value[s].Filename==u.name){i.value.splice(Number(s),1);break}i.value.push({Filename:u.name,FileType:u.name.split(".")[1],FileSize:u.size}),l.elMessage({message:u.name+" \u4E0A\u4F20\u6210\u529F!",type:"success"})}l.codeProcessor((a=e==null?void 0:e.code)!=null?a:100001,(o=e==null?void 0:e.msg)!=null?o:"\u670D\u52A1\u5668\u9519\u8BEF\\\\error")})}),l.$refs.upload.clearFiles()})}function E(u){let t=i.value[u].Filename;ElementPlus.ElMessageBox.confirm("\u786E\u5B9A\u89E3\u538B "+t+" \u5417\uFF1F","\u6CE8\u610F",{confirmButtonText:"\u786E\u5B9A",cancelButtonText:"\u53D6\u6D88",type:"warning"}).then(()=>{let e=new FormData;e.append("file",t),l.$post("api/file/unzip/"+n.PID,e,2).then(a=>{var s,r;console.log(a);let o=a.data;o.code==0&&(n.getProblem(n.PID),l.elMessage({message:"\u89E3\u538B\u6210\u529F",type:"success"})),l.codeProcessor((s=o==null?void 0:o.code)!=null?s:100001,(r=o==null?void 0:o.msg)!=null?r:"\u670D\u52A1\u5668\u9519\u8BEF\\\\error")})})}function h(u){let t=i.value[u].Filename;ElementPlus.ElMessageBox.confirm("\u786E\u5B9A\u5220\u9664 "+t+" \u5417\uFF1F","\u6CE8\u610F",{confirmButtonText:"\u786E\u5B9A",cancelButtonText:"\u53D6\u6D88",type:"warning"}).then(()=>{let e=new FormData;e.append("file",t),l.$delete("api/file/"+n.PID,e).then(a=>{var s,r;let o=a.data;o.code==0&&(i.value.splice(u,1),l.elMessage({message:"\u5220\u9664\u6210\u529F",type:"success"})),l.codeProcessor((s=o==null?void 0:o.code)!=null?s:100001,(r=o==null?void 0:o.msg)!=null?r:"\u670D\u52A1\u5668\u9519\u8BEF\\\\error")})})}return Vue.onMounted(()=>{var u=l.$route.query.PID;u&&n.getProblem(u)}),(u, t)=>{const e=Vue.resolveComponent("el-button"),a=Vue.resolveComponent("el-input"),o=Vue.resolveComponent("el-row"),s=Vue.resolveComponent("el-header"),r=Vue.resolveComponent("el-container"),g=Vue.resolveComponent("el-aside"),C=Vue.resolveComponent("upload-filled"),B=Vue.resolveComponent("el-icon"),D=Vue.resolveComponent("el-upload"),m=Vue.resolveComponent("el-main"),N=Vue.resolveComponent("el-conatiner");return Vue.openBlock(),Vue.createBlock(r,{direction:"horizontal"},{default:Vue.withCtx(()=>[Vue.createVNode(g,{style:{width:"287px",height:"100%"}},{default:Vue.withCtx(()=>[Vue.createVNode(r,null,{default:Vue.withCtx(()=>[Vue.createVNode(s,{class:"Container"},{default:Vue.withCtx(()=>[Vue.createVNode(o,null,{default:Vue.withCtx(()=>[P,Vue.createVNode(a,{modelValue:Vue.unref(n).PID,"onUpdate:modelValue":t[1]||(t[1]= c=>Vue.unref(n).PID=c),placeholder:"P1000",style:{width:"203px"},class:"Left"},{append:Vue.withCtx(()=>[Vue.createVNode(e,{icon:Vue.unref(ElementPlusIconsVue.Search),onClick:t[0]||(t[0]= c=>Vue.unref(n).getProblem(null))},null,8,["icon"])]),_:1},8,["modelValue"])]),_:1})]),_:1})]),_:1})]),_:1}),Vue.createVNode(m,{class:"Left",style:{padding:"0 0 0 0"}},{default:Vue.withCtx(()=>[Vue.createVNode(N,{style:{width:"620px"}},{default:Vue.withCtx(()=>[Vue.createVNode(m,{class:"Container"},{default:Vue.withCtx(()=>[Vue.withDirectives(Vue.createElementVNode("div",A,[Vue.createElementVNode("div",k,[I,Vue.withDirectives(Vue.createElementVNode("div",S," \u6682\u65E0\u6587\u4EF6 ",512),[[Vue.vShow,Vue.unref(i).length==0]]),(Vue.openBlock(!0),Vue.createElementBlock(Vue.Fragment,null,Vue.renderList(Vue.unref(i),(c, _)=>(Vue.openBlock(),Vue.createElementBlock("div",{class:"fileItem",key:c.Filename},[Vue.createElementVNode("div",null,Vue.toDisplayString(c.Filename),1),Vue.createElementVNode("div",null,Vue.toDisplayString(c.Filename.split(".")[1]),1),Vue.createElementVNode("div",null,Vue.toDisplayString((c.FileSize/1024/1024).toFixed(2))+" MB",1),Vue.createElementVNode("div",null,[Vue.createVNode(e,{type:"danger",plain:"",size:"small",onClick: x=>h(_)},{default:Vue.withCtx(()=>[Vue.createTextVNode(" \u5220\u9664 ")]),_:2},1032,["onClick"]),c.Filename.split(".")[1]=="zip"?(Vue.openBlock(),Vue.createBlock(e,{key:0,type:"warning",plain:"",size:"small",onClick: x=>E(_)},{default:Vue.withCtx(()=>[Vue.createTextVNode(" \u89E3\u538B ")]),_:2},1032,["onClick"])):Vue.createCommentVNode("",!0)])]))),128))]),Vue.createVNode(D,{ref:"upload",class:"uploadJson",drag:"",accept:".zip,.in,.out",multiple:!0,"auto-upload":!1,"on-change":F,"on-remove":f},{default:Vue.withCtx(()=>[Vue.createVNode(B,{class:"el-icon--upload"},{default:Vue.withCtx(()=>[Vue.createVNode(C)]),_:1}),b]),_:1},512)],512),[[Vue.vShow,Vue.unref(n).isSearched]])]),_:1}),Vue.createVNode(e,{onClick:v,type:"primary",class:"uploadButton Top"},{default:Vue.withCtx(()=>[Vue.createTextVNode(" \u4E0A\u4F20 ")]),_:1})]),_:1})]),_:1})]),_:1})}}});const T=y(M,[["__scopeId","data-v-917d811d"]]);export{T as default};
