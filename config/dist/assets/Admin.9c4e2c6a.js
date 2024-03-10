import{d as v,c as b,_ as S}from "./index.7fe3d568.js";const r= c=>(Vue.pushScopeId("data-v-578e4ad1"),c=c(),Vue.popScopeId(),c),g=r(()=>Vue.createElementVNode("span",null,"\u62A5\u8868",-1)),D=r(()=>Vue.createElementVNode("span",null,"\u9898\u76EE",-1)),L=r(()=>Vue.createElementVNode("span",null,"\u6570\u636E\u751F\u6210\u5668",-1)),I=r(()=>Vue.createElementVNode("span",null,"\u6BD4\u8D5B",-1)),U=r(()=>Vue.createElementVNode("span",null,"\u9898\u5355",-1)),T=r(()=>Vue.createElementVNode("span",null,"\u7528\u6237",-1)),P=r(()=>Vue.createElementVNode("span",null,"\u6743\u9650",-1)),O=r(()=>Vue.createElementVNode("span",null,"\u516C\u544A",-1)),j={class:"admin Bottom"},H=Vue.defineComponent({__name:"Admin",setup(c){const{proxy:s}=Vue.getCurrentInstance(),V=v(),a=b();var e=Vue.reactive({Administrator:!1,ProblemAdmin:!1,ContestAdmin:!1,SourceBorwser:!1,ListAdmin:!1,SuperAdmin:!1});function m(){let o=localStorage.getItem("ahutOjToken"),u=localStorage.getItem("ahutOjUserUid");if(!V.isLogin){s.elMessage({message:"\u767B\u5F55\u72B6\u6001\u65E0\u6548!",type:"warning"}),i();return}if(V.UID!=u||!o){s.elMessage({message:"\u767B\u5F55\u8EAB\u4EFD\u9A8C\u8BC1\u4E0D\u901A\u8FC7!",type:"warning"}),i();return}s.$axios.get("api/admin/permission/"+u).then(d=>{var n,l;let t=d.data;t.code==0&&_(t.PermissionMap),s.codeProcessor((n=t==null?void 0:t.code)!=null?n:100001,(l=t==null?void 0:t.msg)!=null?l:"\u670D\u52A1\u5668\u9519\u8BEF\\\\error")})}function _(o){if(o<=3){s.elMessage({message:"\u6743\u9650\u9A8C\u8BC1\u672A\u901A\u8FC7\uFF0C\u4F60\u65E0\u6743\u8BBF\u95EE",type:"error"}),i();return}e.Administrator=!0;let u="<span>\u4F60\u7684\u6743\u9650\u8868:</span><br/>        <span style='color: #00ccff'>\u7BA1\u7406\u5458</span>";(o&a.SuperAdminBit)!=0&&(e.SuperAdmin=!0,e.ProblemAdmin=!0,e.ContestAdmin=!0,e.ListAdmin=!0,e.SourceBorwser=!0),u+=e.SuperAdmin?"<br /><span style='color: #6741d9'>\u8D85\u7EA7\u7BA1\u7406\u5458</span>":"<br /><span style='color: #ff3300'>\u8D85\u7EA7\u7BA1\u7406\u5458</span>",(o&a.ProblemAdminBit)!=0&&(e.ProblemAdmin=!0),u+=e.ProblemAdmin?"<br /><span style='color: #5ebd00'>\u9898\u76EE\u7F16\u8F91</span>":"<br /><span style='color: #ff3300'>\u9898\u76EE\u7F16\u8F91</span>",(o&a.ContestAdminBit)!=0&&(e.ContestAdmin=!0),u+=e.ContestAdmin?"<br /><span style='color: #5ebd00'>\u7ADE\u8D5B\u7F16\u8F91</span>":"<br /><span style='color: #ff3300'>\u7ADE\u8D5B\u7F16\u8F91</span>",(o&a.SourceBorwserBit)!=0&&(e.SourceBorwser=!0),u+=e.SourceBorwser?"<br /><span style='color: #5ebd00'>\u8D44\u6E90\u7BA1\u7406</span>":"<br /><span style='color: #ff3300'>\u8D44\u6E90\u7BA1\u7406</span>",(o&a.ListAdminBit)!=0&&(e.ListAdmin=!0),u+=e.ListAdmin?"<br /><span style='color: #5ebd00'>\u9898\u5355\u7F16\u8F91</span>":"<br /><span style='color: #ff3300'>\u9898\u5355\u7F16\u8F91</span>",s.elNotification({title:"\u6B22\u8FCE\u6765\u5230\u7BA1\u7406\u754C\u9762",message:u,type:"success",duration:3e3,dangerouslyUseHTMLString:!0,offset:100})}function i(){s.$router.replace({path:"Home"})}Vue.onMounted(()=>{m()});const p=(o, u)=>{console.log(o,u)},f=(o, u)=>{console.log(o,u)};return(o, u)=>{const d=Vue.resolveComponent("DataAnalysis"),t=Vue.resolveComponent("el-icon"),n=Vue.resolveComponent("el-menu-item"),l=Vue.resolveComponent("Tickets"),C=Vue.resolveComponent("SwitchFilled"),A=Vue.resolveComponent("Files"),B=Vue.resolveComponent("User"),h=Vue.resolveComponent("Lock"),N=Vue.resolveComponent("Notification"),x=Vue.resolveComponent("el-menu"),E=Vue.resolveComponent("el-aside"),w=Vue.resolveComponent("router-view"),k=Vue.resolveComponent("el-main"),y=Vue.resolveComponent("el-container");return Vue.unref(e).Administrator?(Vue.openBlock(),Vue.createBlock(y,{key:0,class:"main"},{default:Vue.withCtx(()=>[Vue.createVNode(E,{class:"main-aside"},{default:Vue.withCtx(()=>[Vue.createVNode(x,{"active-text-color":"#ffd04b","background-color":"#262626",class:"menu","default-active":"/Admin/Stastics","text-color":"#fff",onOpen:p,onClose:f,router:""},{default:Vue.withCtx(()=>[Vue.createVNode(n,{index:"/Admin/Stastics"},{default:Vue.withCtx(()=>[Vue.createVNode(t,null,{default:Vue.withCtx(()=>[Vue.createVNode(d)]),_:1}),g]),_:1}),Vue.unref(e).ProblemAdmin?(Vue.openBlock(),Vue.createBlock(n,{key:0,index:"/Admin/ProblemEdit"},{default:Vue.withCtx(()=>[Vue.createVNode(t,null,{default:Vue.withCtx(()=>[Vue.createVNode(l)]),_:1}),D]),_:1})):Vue.createCommentVNode("",!0),Vue.createVNode(n,{index:"/Admin/DataGenerator"},{default:Vue.withCtx(()=>[Vue.createVNode(t,null,{default:Vue.withCtx(()=>[Vue.createVNode(l)]),_:1}),L]),_:1}),Vue.unref(e).ContestAdmin?(Vue.openBlock(),Vue.createBlock(n,{key:1,index:"/Admin/ContestEdit"},{default:Vue.withCtx(()=>[Vue.createVNode(t,null,{default:Vue.withCtx(()=>[Vue.createVNode(C)]),_:1}),I]),_:1})):Vue.createCommentVNode("",!0),Vue.unref(e).ListAdmin?(Vue.openBlock(),Vue.createBlock(n,{key:2,index:"/Admin/ListEdit"},{default:Vue.withCtx(()=>[Vue.createVNode(t,null,{default:Vue.withCtx(()=>[Vue.createVNode(A)]),_:1}),U]),_:1})):Vue.createCommentVNode("",!0),Vue.unref(e).Administrator?(Vue.openBlock(),Vue.createBlock(n,{key:3,index:"/Admin/UserEdit"},{default:Vue.withCtx(()=>[Vue.createVNode(t,null,{default:Vue.withCtx(()=>[Vue.createVNode(B)]),_:1}),T]),_:1})):Vue.createCommentVNode("",!0),Vue.unref(e).SuperAdmin?(Vue.openBlock(),Vue.createBlock(n,{key:4,index:"/Admin/AdminEdit"},{default:Vue.withCtx(()=>[Vue.createVNode(t,null,{default:Vue.withCtx(()=>[Vue.createVNode(h)]),_:1}),P]),_:1})):Vue.createCommentVNode("",!0),Vue.unref(e).Administrator?(Vue.openBlock(),Vue.createBlock(n,{key:5,index:"/Admin/NoticeEdit"},{default:Vue.withCtx(()=>[Vue.createVNode(t,null,{default:Vue.withCtx(()=>[Vue.createVNode(N)]),_:1}),O]),_:1})):Vue.createCommentVNode("",!0)]),_:1})]),_:1}),Vue.createVNode(k,{class:"mainComponent"},{default:Vue.withCtx(()=>[Vue.createElementVNode("div",j,[Vue.createVNode(w,null,{default:Vue.withCtx(({Component:F})=>[Vue.createVNode(Vue.Transition,{"enter-active-class":"animate__animated SlideInDown"},{default:Vue.withCtx(()=>[(Vue.openBlock(),Vue.createBlock(Vue.resolveDynamicComponent(F)))]),_:2},1024)]),_:1})])]),_:1})]),_:1})):Vue.createCommentVNode("",!0)}}});const G=S(H,[["__scopeId","data-v-578e4ad1"]]);export{G as default};