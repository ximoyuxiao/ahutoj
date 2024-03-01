import{h as C,_ as w}from "./index.7fe3d568.js";const f= p=>(Vue.pushScopeId("data-v-72f934ff"),p=p(),Vue.popScopeId(),p),x={class:"updateList"},N={class:"search"},y=f(()=>Vue.createElementVNode("span",{style:{width:"100px"}},"\u9898\u5355\u53F7\uFF1A",-1)),L={key:0,class:"table"},P={class:"addlist"},I=f(()=>Vue.createElementVNode("span",null,"\u6BD4\u8D5B\u6807\u9898\uFF1A",-1)),B=f(()=>Vue.createElementVNode("span",null,"\u6BD4\u8D5B\u63CF\u8FF0\uFF1A",-1)),S={class:"count"},A=f(()=>Vue.createElementVNode("div",{style:{display:"flex"}},null,-1)),b={style:{display:"flex","justify-content":"flex-end",padding:"10px 0"}},k={class:"showList"},T={class:"list"},$=["checked","onChange"],M=["onClick"],U={class:"pagination"},z=Vue.defineComponent({__name:"UpdateList",setup(p){const{proxy:l}=Vue.getCurrentInstance(),h=200;var r=Vue.reactive({LID:1e3,isSearched:!1,onFocus(){r.isSearched=!1},getListInfo(t){t&&(r.LID=t),l.$axios.get("api/training/"+r.LID).then(e=>{var o,a;let u=e.data;u.code==0?(i.copy(u),d.uptdate(u.Data),r.isSearched=!0):r.isSearched=!1,l.codeProcessor((o=u==null?void 0:u.code)!=null?o:100001,(a=u==null?void 0:u.msg)!=null?a:"\u670D\u52A1\u5668\u9519\u8BEF\\\\error")})}}),i=Vue.reactive({Title:"",Description:"",Problems:[],init(){i.Title="",i.Description="",i.Problems=[]},copy(t){i.Title=t.Title,i.Description=t.Description,t.Data.forEach(e=>{i.Problems.push(e.LID)})}}),d=Vue.reactive({data:[],searchPid:"",search(){if(this.data.length==h){l.elMessage({message:"\u6DFB\u52A0\u5931\u8D25,\u4E00\u4E2A\u9898\u5355\u6700\u591A\u53EA\u80FD\u6DFB\u52A0"+h+"\u9053\u9898\u76EE",type:"warning"});return}for(let t in this.data)if(this.data[t].PID==this.searchPid){l.elMessage({message:"\u6DFB\u52A0\u5931\u8D25\uFF0C\u8BE5\u9898\u76EE\u5DF2\u7ECF\u5B58\u5728\uFF01",type:"warning"});return}l.$axios.get("api/problem/"+this.searchPid).then(t=>{var u,o;let e=t.data;if(e.code==0){let a={PID:e.PID,Title:e.Title};this.data.push(a),l.elMessage({message:"\u6DFB\u52A0\u6210\u529F!",type:"success"})}else l.codeProcessor((u=e==null?void 0:e.code)!=null?u:100001,(o=e==null?void 0:e.msg)!=null?o:"\u670D\u52A1\u5668\u9519\u8BEF\\\\error")})},delete(t, e){for(let u in this.data)t==u&&this.data.splice(u,1);l.elMessage({message:"\u5220\u9664\u9898\u53F7<"+e.PID+">\u6210\u529F!",type:"success"})},uptdate(t){this.searchPid="",this.data=[],t.forEach(e=>{d.data.push({PID:e.PID,Title:e.Ptitle})})}});function D(){if(i.Problems=[],d.data.forEach(t=>{i.Problems.push(t.PID)}),i.Title==""){C({message:"\u8BF7\u8F93\u5165\u9898\u5355\u6807\u9898!",type:"warning"});return}l.$post("/api/training/edit/",{LID:r.LID,Title:i.Title,Description:i.Description,Problems:i.Problems.join(",")}).then(t=>{var u,o;let e=t.data;e.code==0?l.elMessage({message:"\u4FEE\u6539\u6210\u529F!",type:"success"}):l.codeProcessor((u=e==null?void 0:e.code)!=null?u:100001,(o=e==null?void 0:e.msg)!=null?o:"\u670D\u52A1\u5668\u9519\u8BEF\\\\error")})}function _(t){l.$post("/api/training/delete/",{LID:t}).then(e=>{var o,a;let u=e.data;u.code==0?l.elMessage({message:"\u5220\u9664\u6210\u529F!",type:"success"}):l.codeProcessor((o=u==null?void 0:u.code)!=null?o:100001,(a=u==null?void 0:u.msg)!=null?a:"\u670D\u52A1\u5668\u9519\u8BEF\\\\error")})}var c=Vue.reactive({Count:0,currentPage:1,limit:20,init(){c.Count=0,c.currentPage=1,c.limit=20}}),s=Vue.reactive({Data:[],isShowed:!1,selectAlled:!1,init(){this.Data=[]},showList:()=>{s.isShowed?s.isShowed=!1:(s.isShowed=!0,s.search())},search:()=>{l.$get("api/training/list?Page="+(c.currentPage-1)+"&Limit="+c.limit).then(t=>{var u,o;let e=t.data;e.code==0&&(c.Count=e.Count,s.Data=e.data,s.Data.forEach(a=>{a.selected=!1}),s.isShowed=!0,s.selectAlled=!1),l.codeProcessor((u=e==null?void 0:e.code)!=null?u:100001,(o=e==null?void 0:e.msg)!=null?o:"\u670D\u52A1\u5668\u9519\u8BEF\\\\error")})},changePage: t=>{c.currentPage=t,s.search()},selectList: t=>{s.Data[t].selected=!s.Data[t].selected},selectAll:()=>{s.selectAlled=!s.selectAlled,s.Data.forEach(t=>{t.selected=s.selectAlled})},batchDelete:()=>{let t=[],e="";if(s.Data.forEach(u=>{u.selected&&(t.push(u.LID),e+=u.LID+" ")}),t.length<=0){l.elMessage({message:"\u672A\u9009\u62E9\u4EFB\u4F55\u9898\u5355\uFF01",type:"warning"});return}ElementPlus.ElMessageBox.confirm("\u786E\u5B9A\u8981\u6279\u91CF\u5220\u9664\u9898\u5355 "+e+" \u9898\u5355\u5417\uFF1F","\u6CE8\u610F",{confirmButtonText:"\u786E\u5B9A",cancelButtonText:"\u53D6\u6D88",type:"warning"}).then(()=>{l.$axios.post("/api/training/delete/",{LID:t}).then(u=>{var a,V;let o=u.data;o.code==0&&(i.init(),l.elMessage({message:"\u6279\u91CF\u5220\u9664\u6210\u529F!",type:"success"})),l.codeProcessor((a=o==null?void 0:o.code)!=null?a:100001,(V=o==null?void 0:o.msg)!=null?V:"\u670D\u52A1\u5668\u9519\u8BEF\\\\error")}),s.isShowed=!1})}});return Vue.onMounted(()=>{var t=l.$route.query.LID;t&&(r.LID=t,r.getListInfo(t))}),(t, e)=>{const u=Vue.resolveComponent("el-input-number"),o=Vue.resolveComponent("el-button"),a=Vue.resolveComponent("el-divider"),V=Vue.resolveComponent("el-input"),m=Vue.resolveComponent("el-table-column"),v=Vue.resolveComponent("el-table"),F=Vue.resolveComponent("el-pagination");return Vue.openBlock(),Vue.createElementBlock("div",x,[Vue.createElementVNode("div",N,[y,Vue.createVNode(u,{modelValue:Vue.unref(r).LID,"onUpdate:modelValue":e[0]||(e[0]= n=>Vue.unref(r).LID=n),min:1e3,style:{width:"200px"},onFocus:e[1]||(e[1]= n=>Vue.unref(r).onFocus())},null,8,["modelValue"]),Vue.createVNode(o,{plain:"",onClick:e[2]||(e[2]= n=>Vue.unref(r).getListInfo(null))},{default:Vue.withCtx(()=>[Vue.createTextVNode("\u641C\u7D22")]),_:1})]),Vue.createVNode(a),Vue.unref(r).isSearched?(Vue.openBlock(),Vue.createElementBlock("div",L,[Vue.createElementVNode("div",P,[Vue.createElementVNode("div",null,[I,Vue.createVNode(V,{modelValue:Vue.unref(i).Title,"onUpdate:modelValue":e[3]||(e[3]= n=>Vue.unref(i).Title=n)},null,8,["modelValue"])]),Vue.createElementVNode("div",null,[B,Vue.createVNode(V,{modelValue:Vue.unref(i).Description,"onUpdate:modelValue":e[4]||(e[4]= n=>Vue.unref(i).Description=n),type:"textarea",autosize:""},null,8,["modelValue"])]),Vue.createVNode(v,{data:Vue.unref(d).data,style:{width:"100%",margin:"20px 0"}},{default:Vue.withCtx(()=>[Vue.createVNode(m,{label:"\u9898\u53F7",prop:"PID"}),Vue.createVNode(m,{label:"\u6807\u9898",prop:"Title"}),Vue.createVNode(m,{align:"right"},{header:Vue.withCtx(()=>[Vue.createElementVNode("div",S,Vue.toDisplayString(Vue.unref(d).data.length)+"/"+Vue.toDisplayString(h),1),Vue.createVNode(V,{modelValue:Vue.unref(d).searchPid,"onUpdate:modelValue":e[5]||(e[5]= n=>Vue.unref(d).searchPid=n),style:{width:"200px"}},null,8,["modelValue"]),Vue.createVNode(o,{type:"primary",onClick:e[6]||(e[6]= n=>Vue.unref(d).search())},{default:Vue.withCtx(()=>[Vue.createTextVNode(" \u6DFB\u52A0\u9898\u76EE ")]),_:1}),A]),default:Vue.withCtx(n=>[Vue.createVNode(o,{type:"danger",onClick: g=>Vue.unref(d).delete(n.$index,n.row)},{default:Vue.withCtx(()=>[Vue.createTextVNode(" \u5220\u9664 ")]),_:2},1032,["onClick"])]),_:1})]),_:1},8,["data"]),Vue.createElementVNode("div",b,[Vue.createVNode(o,{plain:"",onClick:e[7]||(e[7]= n=>D())},{default:Vue.withCtx(()=>[Vue.createTextVNode("\u66F4\u65B0")]),_:1}),Vue.createVNode(o,{plain:"",onClick:e[8]||(e[8]= n=>_([Vue.unref(r).LID]))},{default:Vue.withCtx(()=>[Vue.createTextVNode("\u5220\u9664")]),_:1})])])])):Vue.createCommentVNode("",!0),Vue.createElementVNode("div",k,[Vue.createVNode(o,{plain:"",onClick:e[9]||(e[9]= n=>Vue.unref(s).showList())},{default:Vue.withCtx(()=>[Vue.createTextVNode(Vue.toDisplayString(Vue.unref(s).isShowed?"\u5173\u95ED\u5217\u8868":"\u663E\u793A\u5217\u8868"),1)]),_:1}),Vue.withDirectives(Vue.createVNode(o,{plain:"",onClick:e[10]||(e[10]= n=>Vue.unref(s).selectAll())},{default:Vue.withCtx(()=>[Vue.createTextVNode(" \u5F53\u9875\u5168\u9009 ")]),_:1},512),[[Vue.vShow,Vue.unref(s).isShowed]]),Vue.withDirectives(Vue.createVNode(o,{plain:"",type:"warning",onClick:e[11]||(e[11]= n=>Vue.unref(s).batchDelete())},{default:Vue.withCtx(()=>[Vue.createTextVNode(" \u6279\u91CF\u5220\u9664 ")]),_:1},512),[[Vue.vShow,Vue.unref(s).isShowed]])]),Vue.withDirectives(Vue.createElementVNode("div",T,[(Vue.openBlock(!0),Vue.createElementBlock(Vue.Fragment,null,Vue.renderList(Vue.unref(s).Data,(n, g)=>(Vue.openBlock(),Vue.createElementBlock("div",{class:Vue.normalizeClass(n.selected?"item itemSelected":"item"),key:g},[Vue.createElementVNode("input",{type:"checkbox",checked:n.selected,onChange:Vue.withModifiers(E=>Vue.unref(s).selectList(g),["stop"])},null,40,$),Vue.createElementVNode("div",{class:"title cursor_pointer",onClick:Vue.withModifiers(E=>Vue.unref(r).getListInfo(n.LID),["stop"])},Vue.toDisplayString(n.LID)+"\xA0-\xA0"+Vue.toDisplayString(n.Title),9,M)],2))),128))],512),[[Vue.vShow,Vue.unref(s).isShowed]]),Vue.withDirectives(Vue.createElementVNode("div",U,[Vue.createVNode(F,{background:"",layout:"prev, pager, next","page-size":Vue.unref(c).limit,total:Vue.unref(c).Count,"current-page":Vue.unref(c).currentPage,onCurrentChange:Vue.unref(s).changePage},null,8,["page-size","total","current-page","onCurrentChange"])],512),[[Vue.vShow,Vue.unref(s).isShowed]])])}}});const q=w(z,[["__scopeId","data-v-72f934ff"]]);export{q as default};
