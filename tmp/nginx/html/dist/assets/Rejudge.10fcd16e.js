import{_ as i}from"./index.de99a551.js";const d=s=>(Vue.pushScopeId("data-v-5d82ef84"),s=s(),Vue.popScopeId(),s),p={class:"Rejudge"},m=d(()=>Vue.createElementVNode("span",null,"\u63D0\u4EA4ID\uFF1A",-1)),c=d(()=>Vue.createElementVNode("span",null,"\u9898\u76EEID\uFF1A",-1)),D=d(()=>Vue.createElementVNode("span",null,"\u7528\u6237ID\uFF1A",-1)),_=d(()=>Vue.createElementVNode("span",null,"\u7ADE\u8D5BID\uFF1A",-1)),I=d(()=>Vue.createElementVNode("span",null,'"0"\u6216\u8005\u7A7A\u4EE3\u8868\u4E0D\u586B',-1)),f={style:{display:"flex","justify-content":"flex-start",padding:"10px 0"}},E=Vue.defineComponent({__name:"Rejudge",setup(s){const{proxy:r}=Vue.getCurrentInstance();var e=Vue.reactive({SID:0,PID:"",UID:"",CID:0,init(){this.SID=0,this.PID="",this.UID="",this.CID=0}});function a(){let o={};e.SID&&(o.SID=e.SID),e.PID&&(o.PID=e.PID),e.UID&&(o.UID=e.UID),e.CID&&(o.CID=e.CID),r.$post("api/submit/rejudge/",o).then(u=>{var l,V;console.log(u);let t=u.data;t.code==0&&r.elMessage({message:"\u91CD\u5224\u6210\u529F!",type:"success"}),r.codeProcessor((l=t==null?void 0:t.code)!=null?l:100001,(V=t==null?void 0:t.msg)!=null?V:"\u670D\u52A1\u5668\u9519\u8BEF\\\\error")})}return(o,u)=>{const t=Vue.resolveComponent("el-input-number"),l=Vue.resolveComponent("el-input"),V=Vue.resolveComponent("el-button");return Vue.openBlock(),Vue.createElementBlock("div",p,[Vue.createElementVNode("div",null,[m,Vue.createVNode(t,{modelValue:Vue.unref(e).SID,"onUpdate:modelValue":u[0]||(u[0]=n=>Vue.unref(e).SID=n),min:0},null,8,["modelValue"])]),Vue.createElementVNode("div",null,[c,Vue.createVNode(l,{modelValue:Vue.unref(e).PID,"onUpdate:modelValue":u[1]||(u[1]=n=>Vue.unref(e).PID=n),style:{width:"200px"}},null,8,["modelValue"])]),Vue.createElementVNode("div",null,[D,Vue.createVNode(l,{modelValue:Vue.unref(e).UID,"onUpdate:modelValue":u[2]||(u[2]=n=>Vue.unref(e).UID=n),style:{width:"200px"}},null,8,["modelValue"])]),Vue.createElementVNode("div",null,[_,Vue.createVNode(t,{modelValue:Vue.unref(e).CID,"onUpdate:modelValue":u[3]||(u[3]=n=>Vue.unref(e).CID=n),min:0},null,8,["modelValue"])]),I,Vue.createElementVNode("div",f,[Vue.createVNode(V,{plain:"",type:"warning",onClick:u[4]||(u[4]=n=>a())},{default:Vue.withCtx(()=>[Vue.createTextVNode(" \u91CD\u5224 ")]),_:1})])])}}});const C=i(E,[["__scopeId","data-v-5d82ef84"]]);export{C as default};
