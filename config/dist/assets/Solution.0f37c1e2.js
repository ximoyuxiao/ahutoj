import{d as z,u as M,s as S,_ as H}from "./index.7fe3d568.js";const U= C=>(Vue.pushScopeId("data-v-148c444d"),C=C(),Vue.popScopeId(),C),R={class:"notFound"},q={class:"topBox"},j={class:"count"},G={class:"userInfo"},J={class:"headImg"},K=["src"],O={class:"info"},Q={class:"userName"},W={class:"updateTime"},X={class:"updateTime"},Y={class:"foldCover"},Z=["onClick"],ee={class:"status"},te={class:"left"},oe={class:"thumbsUp"},ue=["onClick"],ne=["onClick"],se={class:"comment"},re={class:"item"},le={class:"left"},ae=["src"],ce={class:"right"},ie={class:"userName"},Ve={class:"content"},de={class:"buttom"},me=["onClick"],pe={class:"time"},Ce={key:0,class:"notFound"},_e={key:1,class:"pagination"},he={class:"myComment"},ge=["onClick"],fe={key:0,class:"pagination"},ve={class:"publish"},De=U(()=>Vue.createElementVNode("h1",null,"\u6807\u9898",-1)),Ee=U(()=>Vue.createElementVNode("div",{class:"right"},[Vue.createElementVNode("div",{class:"hint"},[Vue.createElementVNode("div",{class:"title"}," \u9898\u89E3\u4EC5\u4F9B\u5B66\u4E60\u53C2\u8003\u4F7F\u7528\uFF01 "),Vue.createElementVNode("div",{class:"text"}," \u6284\u88AD\u3001\u590D\u5236\u9898\u89E3\uFF0C\u4EE5\u8FBE\u5230\u5237 AC \u7387/AC \u6570\u91CF\u6216\u5176\u4ED6\u76EE\u7684\u7684\u884C\u4E3A\uFF0C\u5728\u672C\u5E73\u53F0\u662F\u4E25\u683C\u7981\u6B62\u7684\u3002 "),Vue.createElementVNode("div",{class:"text"},[Vue.createTextVNode(" \u5E73\u53F0\u975E\u5E38\u91CD\u89C6\u5B66\u672F\u8BDA\u4FE1\u3002\u6B64\u7C7B\u884C\u4E3A\u5C06\u4F1A\u5BFC\u81F4\u60A8\u7684\u8D26\u53F7\u88AB"),Vue.createElementVNode("strong",null,"\u5C01\u7981"),Vue.createTextVNode("\u3002 ")])])],-1)),Ne=Vue.defineComponent({__name:"Solution",setup(C){const d=z(),_=M(),{proxy:r}=Vue.getCurrentInstance();var h=Vue.ref(!0),V=Vue.reactive({currentPage:1,Page:1,Limit:10,Count:0,changePage: o=>{V.currentPage=o,A(),e.getSolutions(e.PID)}}),w={toolbar:["bold","underline","italic","-","strikeThrough","title","sub","sup","quote","unorderedList","orderedList","-","codeRow","code","link","image","table","mermaid","katex","-","revoke","next","save","=","preview","htmlPreview","catalog"]},e=Vue.reactive({PID:"",data:[],myComment:"",title:"",mySolution:"",getSolutions: o=>{h.value=!0;let l=d.isLogin?d.UID:null;r.$get("api/solution/solutions",{PID:o,UID:l,Page:V.currentPage-1,Limit:V.Limit},0,2).then(n=>{var c,s;let u=n.data;if((u==null?void 0:u.code)==0){e.data=u.SolutionList,V.Count=u.count;for(let a of e.data)a.CanFold=!1,a.Fold=!1,a.Page=1,a.Limit=10,a.count=0,a.Comments=[],a.ShowComments=!1,a.ThumbsUp=a.FavoriteCount,a.IThumbsUp=a.isFavorite,a.CommentCount=0,a.CommentChangePage= m=>{a.Page=m,e.getComments(a)},e.getComments(a);h.value=e.data.length==0}r.codeProcessor((c=u==null?void 0:u.code)!=null?c:100001,(s=u==null?void 0:u.msg)!=null?s:"\u670D\u52A1\u5668\u9519\u8BEF\\\\error")})},fold:(o, l)=>{e.data[o].Fold=l},thumbUp:(o, l, n)=>{if(!d.isLogin||!d.UID){r.elMessage({message:"\u60A8\u8FD8\u672A\u767B\u5F55",type:"warning"});return}let u=d.UID;r.$post("api/favorite/",{SID:l,UID:u,ActionType:n},0,2).then(c=>{var a,m;let s=c.data;(s==null?void 0:s.code)==0&&(s.data,e.data[o].IThumbsUp=s.IsFavorite,e.data[o].ThumbsUp=s.Count,r.elNotification({message:s.IsFavorite==!0?"\u6210\u529F\u70B9\u8D5E":"\u53D6\u6D88\u70B9\u8D5E",type:"success"})),r.codeProcessor((a=s==null?void 0:s.code)!=null?a:100001,(m=s==null?void 0:s.msg)!=null?m:"\u670D\u52A1\u5668\u9519\u8BEF\\\\error")})},openComment: o=>{if(e.data[o].ShowComment){e.data[o].ShowComment=!1;return}e.myComment="",e.data[o].ShowComment=!0,e.getComments(e.data[o])},getComments: o=>{r.$get("api/comment/comments",{SID:o.SID,Page:o.Page-1,Limit:o.Limit},0,2).then(l=>{var u,c;let n=l.data;(n==null?void 0:n.code)==0&&(o.Count=n.Count,o.Comments=n.Data,o.CommentCount=n.Count),r.codeProcessor((u=n==null?void 0:n.code)!=null?u:100001,(c=n==null?void 0:n.msg)!=null?c:"\u670D\u52A1\u5668\u9519\u8BEF\\\\error")})},closeComment: o=>{e.data[o].ShowComment=!1},publishComment: o=>{if(!d.isLogin||!d.UID){r.elMessage({message:"\u60A8\u8FD8\u672A\u767B\u5F55",type:"warning"});return}if(!e.myComment){r.elMessage({message:"\u8BF7\u8F93\u5165\u5185\u5BB9",type:"warning"});return}let l=e.myComment;r.$post("/api/comment/",{SID:e.data[o].SID,ActionType:1,UID:d.UID,Text:l,FCID:-1},0,2).then(n=>{var c,s;let u=n.data;(u==null?void 0:u.code)==0&&(e.getComments(e.data[o]),e.myComment="",r.elNotification({message:"\u8BC4\u8BBA\u6210\u529F\uFF01",type:"success"})),r.codeProcessor((c=u==null?void 0:u.code)!=null?c:100001,(s=u==null?void 0:u.msg)!=null?s:"\u670D\u52A1\u5668\u9519\u8BEF\\\\error")})},deleteComment:(o, l, n, u)=>{r.$delete("forum/solution/comment/delete/"+l,{UID:u,SLTID:n},0,2).then(c=>{var a,m;let s=c.data;(s==null?void 0:s.code)==0&&(e.getComments(e.data[o]),r.elNotification({message:"\u5220\u9664\u6210\u529F",type:"success"})),r.codeProcessor((a=s==null?void 0:s.code)!=null?a:100001,(m=s==null?void 0:s.msg)!=null?m:"\u670D\u52A1\u5668\u9519\u8BEF\\\\error")})},publishSolution:()=>{if(!d.isLogin||!d.UID){r.elMessage({message:"\u60A8\u8FD8\u672A\u767B\u5F55",type:"warning"});return}if(!e.mySolution){r.elMessage({message:"\u8BF7\u8F93\u5165\u5185\u5BB9",type:"warning"});return}let o=d.UID;r.$post("api/solution/",{ActionType:1,PID:e.PID,UID:o,Title:e.title,Text:e.mySolution},0,2).then(l=>{var u,c;let n=l.data;n.code==0&&(e.mySolution="",window.scroll(0,0),r.elNotification({message:"\u53D1\u5E03\u6210\u529F\uFF01",type:"success"})),r.codeProcessor((u=n==null?void 0:n.code)!=null?u:100001,(c=n==null?void 0:n.msg)!=null?c:"\u670D\u52A1\u5668\u9519\u8BEF\\\\error")}),e.getSolutions(e.PID)}});function A(){r.$router.replace({path:"/Solution",query:{PID:e.PID,Page:V.currentPage,Limit:V.Limit}})}return Vue.onMounted(()=>{let o=r.$route.query.PID;e.PID=o,e.getSolutions(o)}),(o, l)=>{const n=Vue.resolveComponent("el-empty"),u=Vue.resolveComponent("ArrowDown"),c=Vue.resolveComponent("el-icon"),s=Vue.resolveComponent("CaretTop"),a=Vue.resolveComponent("ArrowUpBold"),m=Vue.resolveComponent("ChatLineSquare"),x=Vue.resolveComponent("ArrowUp"),f=Vue.resolveComponent("el-pagination"),v=Vue.resolveComponent("el-input"),D=Vue.resolveComponent("el-button"),T=Vue.resolveComponent("Hide"),g=Vue.resolveComponent("el-radio-button"),P=Vue.resolveComponent("el-radio-group"),b=Vue.resolveComponent("el-main"),L=Vue.resolveComponent("el-footer"),E=Vue.resolveComponent("el-container"),$=Vue.resolveComponent("el-asider");return Vue.openBlock(),Vue.createBlock(E,{class:"Main"},{default:Vue.withCtx(()=>[Vue.createVNode(E,null,{default:Vue.withCtx(()=>[Vue.createVNode(b,{class:"Container"},{default:Vue.withCtx(()=>[Vue.withDirectives(Vue.createElementVNode("div",R,[Vue.createVNode(n,{description:"\u8BE5\u9898\u8FD8\u6728\u6709\u9898\u89E3"})],512),[[Vue.vShow,Vue.unref(h)]]),Vue.unref(e).data.length>0?(Vue.openBlock(),Vue.createElementBlock(Vue.Fragment,{key:0},[Vue.createElementVNode("div",q,[Vue.createElementVNode("div",j,"\u5171\u53D1\u73B0\xA0"+Vue.toDisplayString(Vue.unref(V).Count)+"\xA0\u4E2A\u9898\u89E3",1)]),(Vue.openBlock(!0),Vue.createElementBlock(Vue.Fragment,null,Vue.renderList(Vue.unref(e).data,(t, p)=>{var N,F,B,y,k,I;return Vue.openBlock(),Vue.createElementBlock("div",{class:Vue.normalizeClass(t.IThumbsUp==1?"item itemThumbsUp":"item")},[Vue.createElementVNode("div",G,[Vue.createElementVNode("div",J,[Vue.createElementVNode("img",{src:t.HeadURL?Vue.unref(S)+t.HeadURL:Vue.unref(r).Utils.DefaultHeadImage.show(t.UID),style:{height:"80px",width:"80px"}},null,8,K)]),Vue.createElementVNode("div",O,[Vue.createElementVNode("div",Q,Vue.toDisplayString(t.UID),1),Vue.createElementVNode("div",W," \u6700\u540E\u66F4\u65B0\u65F6\u95F4\uFF1A"+Vue.toDisplayString(Vue.unref(r).Utils.TimeTools.timestampToTime(t.UpdateTime)),1),Vue.createElementVNode("div",X," \u521B\u5EFA\u65F6\u95F4:"+Vue.toDisplayString(Vue.unref(r).Utils.TimeTools.timestampToTime(t.CreateTime)),1)])]),Vue.createElementVNode("div",{class:Vue.normalizeClass((N=t==null?void 0:t.Fold)==null||N?"contentFold":"content")},[Vue.createElementVNode("h3",null,Vue.toDisplayString(t.Title),1),Vue.createVNode(Vue.unref(MdEditorV3),{class:"mdEditor",modelValue:t.text,"onUpdate:modelValue": i=>t.text=i,theme:Vue.unref(_).theme>0?"light":"dark","preview-only":""},null,8,["modelValue","onUpdate:modelValue","theme"]),Vue.withDirectives(Vue.createElementVNode("div",Y,[Vue.createElementVNode("div",{onClick: i=>Vue.unref(e).fold(p,!1)},[Vue.createTextVNode(" \u5C55\u5F00 \xA0 "),Vue.createVNode(c,{size:"26px"},{default:Vue.withCtx(()=>[Vue.createVNode(u)]),_:1})],8,Z)],512),[[Vue.vShow,((F=t==null?void 0:t.Fold)!=null?F:!0)&&t.CanFold]])],2),Vue.createElementVNode("div",ee,[Vue.createElementVNode("div",te,[Vue.createElementVNode("div",oe,[t.IThumbsUp==!0?(Vue.openBlock(),Vue.createBlock(c,{key:0,class:"iconThumbsUp cursor_pointer",size:"20px",onClick: i=>Vue.unref(e).thumbUp(p,t.SID,3)},{default:Vue.withCtx(()=>[Vue.createVNode(s)]),_:2},1032,["onClick"])):(Vue.openBlock(),Vue.createBlock(c,{key:1,class:"icon cursor_pointer",size:"20px",onClick: i=>Vue.unref(e).thumbUp(p,t.SID,1)},{default:Vue.withCtx(()=>[Vue.createVNode(a)]),_:2},1032,["onClick"])),Vue.createTextVNode(" \xA0"+Vue.toDisplayString(t.ThumbsUp)+"\xA0\u4E2A\u4EBA\u9876\u4E86\u8FD9\u4E2A\u5E16\u5B50 ",1)]),Vue.createElementVNode("div",{class:"conmment cursor_pointer",onClick: i=>Vue.unref(e).openComment(p)},[Vue.createVNode(c,{class:"icon cursor_pointer",size:"20px"},{default:Vue.withCtx(()=>[Vue.createVNode(m)]),_:1}),Vue.createTextVNode(" \xA0"+Vue.toDisplayString(t.CommentCount)+"\xA0\u6761\u8BC4\u8BBA ",1)],8,ue)]),!(t!=null&&t.Fold)&&t.CanFold?(Vue.openBlock(),Vue.createElementBlock("div",{key:0,class:"right cursor_pointer",onClick: i=>Vue.unref(e).fold(p,!0)},[Vue.createTextVNode(" \u6536\u8D77\xA0 "),Vue.createVNode(c,{size:"26px"},{default:Vue.withCtx(()=>[Vue.createVNode(x)]),_:1})],8,ne)):Vue.createCommentVNode("",!0)]),Vue.withDirectives(Vue.createElementVNode("div",se,[(Vue.openBlock(!0),Vue.createElementBlock(Vue.Fragment,null,Vue.renderList((B=t==null?void 0:t.Comments)!=null?B:[],(i, Fe)=>(Vue.openBlock(),Vue.createElementBlock("div",re,[Vue.createElementVNode("div",le,[Vue.createElementVNode("img",{src:i.HeadURL?Vue.unref(S)+i.HeadURL:Vue.unref(r).Utils.DefaultHeadImage.show(i.UID),alt:""},null,8,ae)]),Vue.createElementVNode("div",ce,[Vue.createElementVNode("div",ie,Vue.toDisplayString(i.UID),1),Vue.createElementVNode("div",Ve,Vue.toDisplayString(i.Text),1),Vue.createElementVNode("div",de,[Vue.unref(d).isLogin&&Vue.unref(d).UID==i.UID?(Vue.openBlock(),Vue.createElementBlock("div",{key:0,class:"delete cursor_pointer",onClick: Be=>Vue.unref(e).deleteComment(p,i.SLTCMTID,t.SLTID,Vue.unref(d).UID)}," \u5220\u9664 ",8,me)):Vue.createCommentVNode("",!0),Vue.createElementVNode("div",pe,Vue.toDisplayString(Vue.unref(r).Utils.TimeTools.timestampToDate(i.UpdateTime,2)),1)])])]))),256)),((y=t.Count)!=null?y:0)==0?(Vue.openBlock(),Vue.createElementBlock("div",Ce,[Vue.createVNode(n,{description:"\u5F53\u524D\u8FD8\u6CA1\u6709\u8BC4\u8BBA\u54E6~"})])):Vue.createCommentVNode("",!0),((k=t.Count)!=null?k:0)>=10?(Vue.openBlock(),Vue.createElementBlock("div",_e,[Vue.createVNode(f,{background:"",layout:"prev, pager, next","page-size":t.Limit,total:t.Count,"current-page":t.Page,onCurrentChange:t.CommentChangePage},null,8,["page-size","total","current-page","onCurrentChange"])])):Vue.createCommentVNode("",!0),Vue.createElementVNode("div",he,[Vue.createVNode(v,{modelValue:Vue.unref(e).myComment,"onUpdate:modelValue":l[0]||(l[0]= i=>Vue.unref(e).myComment=i),autosize:"",type:"textarea",placeholder:"\u7559\u4E0B\u4F60\u53CB\u597D\u7684\u8BC4\u8BBA\u5427\uFF01"},null,8,["modelValue"]),Vue.createVNode(D,{type:"primary",plain:Vue.unref(_).theme>0,onClick: i=>Vue.unref(e).publishComment(p)},{default:Vue.withCtx(()=>[Vue.createTextVNode(" \u53D1\u8868 ")]),_:2},1032,["plain","onClick"])]),Vue.createElementVNode("div",{class:"hideComment cursor_pointer",onClick: i=>Vue.unref(e).closeComment(p)},[Vue.createTextVNode(" \u9690\u85CF\u8BC4\u8BBA\xA0 "),Vue.createVNode(c,{size:"24px"},{default:Vue.withCtx(()=>[Vue.createVNode(T)]),_:1})],8,ge)],512),[[Vue.vShow,(I=t==null?void 0:t.ShowComment)!=null?I:!1]])],2)}),256)),Vue.unref(V).Count>5?(Vue.openBlock(),Vue.createElementBlock("div",fe,[Vue.createVNode(f,{background:"",layout:"prev, pager, next","page-size":Vue.unref(V).Limit,total:Vue.unref(V).Count,"current-page":Vue.unref(V).currentPage,onCurrentChange:Vue.unref(V).changePage},null,8,["page-size","total","current-page","onCurrentChange"]),Vue.createVNode(P,{modelValue:Vue.unref(V).Limit,"onUpdate:modelValue":l[1]||(l[1]= t=>Vue.unref(V).Limit=t),onChange:l[2]||(l[2]= t=>Vue.unref(V).changePage(1)),style:{margin:"0 20px"}},{default:Vue.withCtx(()=>[Vue.createVNode(g,{label:5}),Vue.createVNode(g,{label:10}),Vue.createVNode(g,{label:15})]),_:1},8,["modelValue"])])):Vue.createCommentVNode("",!0)],64)):Vue.createCommentVNode("",!0)]),_:1}),Vue.createVNode(L,{class:"Container"},{default:Vue.withCtx(()=>[Vue.createElementVNode("div",ve,[De,Vue.createVNode(v,{modelValue:Vue.unref(e).title,"onUpdate:modelValue":l[3]||(l[3]= t=>Vue.unref(e).title=t),placeholder:"Please input",clearable:"",style:{width:"500px"}},null,8,["modelValue"]),Vue.createVNode(Vue.unref(MdEditorV3),{class:"mdEditor",modelValue:Vue.unref(e).mySolution,"onUpdate:modelValue":l[4]||(l[4]= t=>Vue.unref(e).mySolution=t),toolbars:Vue.unref(w).toolbar,theme:Vue.unref(_).theme>0?"light":"dark"},null,8,["modelValue","toolbars","theme"]),Vue.createVNode(D,{type:"primary",plain:Vue.unref(_).theme>0,onClick:Vue.unref(e).publishSolution},{default:Vue.withCtx(()=>[Vue.createTextVNode(" \u53D1\u5E03\u9898\u89E3 ")]),_:1},8,["plain","onClick"])])]),_:1})]),_:1}),Vue.createVNode($,{class:"Container Left warning"},{default:Vue.withCtx(()=>[Ee]),_:1})]),_:1})}}});const ke=H(Ne,[["__scopeId","data-v-148c444d"]]);export{ke as default};
