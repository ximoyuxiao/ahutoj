import{_ as d}from"./index.de99a551.js";const s={class:"input"},r=["placeholder","onKeydown","type"],_=Vue.defineComponent({__name:"Input",props:["modelValue","placeholder","type"],emits:["update:modelValue","click"],setup(a,{emit:n}){const o=a;var e=Vue.ref(null);function p(t){n("update:modelValue",t.target.value)}function c(t){n("click",t)}return Vue.onMounted(()=>{o.modelValue&&(e.value=o.modelValue)}),(t,l)=>(Vue.openBlock(),Vue.createElementBlock("div",s,[Vue.withDirectives(Vue.createElementVNode("input",{"onUpdate:modelValue":l[0]||(l[0]=u=>Vue.isRef(e)?e.value=u:e=u),onInput:p,placeholder:o.placeholder,onKeydown:Vue.withKeys(c,["enter"]),type:o.type},null,40,r),[[Vue.vModelDynamic,Vue.unref(e)]])]))}});const V=d(_,[["__scopeId","data-v-2e043a17"]]);export{V as I};
