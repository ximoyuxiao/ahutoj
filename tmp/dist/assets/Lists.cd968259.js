import{_ as h}from"./index.7fe3d568.js";const _=s=>(Vue.pushScopeId("data-v-7eec966c"),s=s(),Vue.popScopeId(),s),f=_(()=>Vue.createElementVNode("span",{class:"FontSize16 Bold DarkGray"},"\u9898\u5355 ID",-1)),C={class:"ProblemList"},v={class:"left"},x={class:"content"},N={class:"list"},B={class:"item cursor_pointer"},y=["onClick"],w={id:"LID",class:"FontSize14 Bold"},L={id:"Title",class:"FontSize18 Bold Bold"},E=_(()=>Vue.createElementVNode("div",{class:"FontSize14 ltype ltypeOffcial Bold",style:{"line-height":"24px"}}," \u5B98\u65B9 ",-1)),P={class:"Time"},D={class:"pagination"},F=Vue.defineComponent({__name:"Lists",setup(s){const{proxy:u}=Vue.getCurrentInstance();var e=Vue.reactive({search:null,Count:0,currentPage:1,limit:20,loading:null,init(){this.Count=0,this.currentPage=1,this.limit=20,this.loading=null},changePage:o=>{e.currentPage=o,p(),a.getData()}}),a=Vue.reactive({list:[],getData:()=>{u.$get("api/training/list",{Page:e.currentPage-1,Limit:e.limit}).then(o=>{var l,i;let t=o.data;t.code===0&&(console.log(t),e.Count=t.size,a.list=t.data),u.codeProcessor((l=t==null?void 0:t.code)!=null?l:100001,(i=t==null?void 0:t.msg)!=null?i:"\u670D\u52A1\u5668\u9519\u8BEF\\\\error")})}});function c(o){if(e.search=o!=null?o:e.search,!e.search){u.elMessage({message:"\u8BF7\u8F93\u5165\u6709\u6548\u7684\u9898\u5355ID\uFF01",type:"warning"});return}u.$router.push({name:"List",params:{LID:e.search}})}function p(){u.$router.replace({path:"/Lists",query:{Page:e.currentPage,Limit:e.limit}})}return a.getData(),(o,t)=>{const l=Vue.resolveComponent("el-button"),i=Vue.resolveComponent("el-input"),r=Vue.resolveComponent("el-row"),V=Vue.resolveComponent("el-main"),m=Vue.resolveComponent("el-aside"),g=Vue.resolveComponent("el-pagination"),d=Vue.resolveComponent("el-container");return Vue.openBlock(),Vue.createBlock(d,{class:"Main Bottom Top"},{default:Vue.withCtx(()=>[Vue.createVNode(m,null,{default:Vue.withCtx(()=>[Vue.createVNode(V,{class:"Container"},{default:Vue.withCtx(()=>[Vue.createVNode(r,{class:"Row"},{default:Vue.withCtx(()=>[f,Vue.createVNode(i,{modelValue:Vue.unref(e).search,"onUpdate:modelValue":t[1]||(t[1]=n=>Vue.unref(e).search=n),placeholder:"e.g. 1000",type:"text",class:"Left goToListInput input-with-select"},{append:Vue.withCtx(()=>[Vue.createVNode(l,{onClick:t[0]||(t[0]=n=>c()),type:"primary",class:"goToProblemButton Bold"},{default:Vue.withCtx(()=>[Vue.createTextVNode(" \u8DF3\u8F6C ")]),_:1})]),_:1},8,["modelValue"])]),_:1})]),_:1})]),_:1}),Vue.createVNode(d,{class:"Left"},{default:Vue.withCtx(()=>[Vue.createVNode(V,{class:"Container"},{default:Vue.withCtx(()=>[Vue.createElementVNode("div",C,[Vue.createElementVNode("div",v,[Vue.createElementVNode("div",x,[Vue.createElementVNode("div",N,[(Vue.openBlock(!0),Vue.createElementBlock(Vue.Fragment,null,Vue.renderList(Vue.unref(a).list,(n,S)=>(Vue.openBlock(),Vue.createElementBlock("div",B,[Vue.createElementVNode("div",{class:"left",onClick:()=>c(n.LID)},[Vue.createVNode(r,null,{default:Vue.withCtx(()=>[Vue.createElementVNode("div",w,"#"+Vue.toDisplayString(n.LID)+"\xA0-\xA0",1),Vue.createElementVNode("div",L,Vue.toDisplayString(n.Title),1),E]),_:2},1024),Vue.createElementVNode("div",P,Vue.toDisplayString(Vue.unref(u).Utils.TimeTools.timestampToTime(n.StartTime)),1)],8,y)]))),256))]),Vue.createElementVNode("div",D,[Vue.createVNode(g,{background:"",layout:"prev, pager, next","page-size":Vue.unref(e).limit,total:Vue.unref(e).Count,"current-page":Vue.unref(e).currentPage,onCurrentChange:Vue.unref(e).changePage},null,8,["page-size","total","current-page","onCurrentChange"])])])])])]),_:1})]),_:1})]),_:1})}}});const k=h(F,[["__scopeId","data-v-7eec966c"]]);export{k as default};
