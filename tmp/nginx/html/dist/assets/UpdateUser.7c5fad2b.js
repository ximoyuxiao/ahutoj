import{_ as c}from"./index.de99a551.js";const a=o=>(Vue.pushScopeId("data-v-a38b164c"),o=o(),Vue.popScopeId(),o),p={class:"UpdateUser"},i={class:"uid"},V=a(()=>Vue.createElementVNode("label",{class:"loginLabel"}," \u8D26\u53F7 ",-1)),_={class:"password"},m=a(()=>Vue.createElementVNode("label",{class:"loginLabel",for:"username"}," \u5BC6\u7801 ",-1)),f=Vue.defineComponent({__name:"UpdateUser",setup(o){const{proxy:u}=Vue.getCurrentInstance();var t=Vue.reactive({UID:"",Password:"",init(){this.UID="",this.Pass=""}});function d(){u.$post("api/admin/user/edit/password/",{UID:t.UID,Password:t.Password}).then(n=>{let e=n.data;e.code==0&&u.elMessage({message:"\u4FEE\u6539\u6210\u529F!",type:"success"}),u.codeProcessor(e.code,e.msg)})}return(n,e)=>{const l=Vue.resolveComponent("el-button"),r=Vue.resolveComponent("el-row");return Vue.openBlock(),Vue.createElementBlock("div",p,[Vue.createElementVNode("div",i,[V,Vue.withDirectives(Vue.createElementVNode("input",{class:"loginInput",name:"UID",type:"text",autocomplete:"off",maxlength:"20","onUpdate:modelValue":e[0]||(e[0]=s=>Vue.unref(t).UID=s)},null,512),[[Vue.vModelText,Vue.unref(t).UID]])]),Vue.createElementVNode("div",_,[m,Vue.withDirectives(Vue.createElementVNode("input",{class:"loginInput",name:"password",type:"text",autocomplete:"off",maxlength:"20","onUpdate:modelValue":e[1]||(e[1]=s=>Vue.unref(t).Password=s)},null,512),[[Vue.vModelText,Vue.unref(t).Password]])]),Vue.createVNode(r,null,{default:Vue.withCtx(()=>[Vue.createVNode(l,{class:"confirm",type:"primary",round:"",shoudAddSpace:"",onClick:e[2]||(e[2]=s=>d())},{default:Vue.withCtx(()=>[Vue.createTextVNode("\u91CD\u7F6E\u5BC6\u7801")]),_:1})]),_:1})])}}});const v=c(f,[["__scopeId","data-v-a38b164c"]]);export{v as default};
