import{S as f,a as c}from"./StatusSearch.23f9c4df.js";import{_ as D}from"./index.de99a551.js";import"./Input.a729bc6b.js";const g={class:"status"},I=Vue.defineComponent({__name:"Status",setup(d){const{proxy:u}=Vue.getCurrentInstance();var n={list:null},s=Vue.reactive({search:e=>{t.Page=1,t.PID=e.PID,t.UID=e.UID,t.Lang=e.Lang,t.Result=e.Result,s.update()},update:()=>{n.list=u.elLoading({node:document.getElementsByClassName("statusList")[0]});let e={};t.Page&&(e.Page=t.Page-1),t.Limit&&(e.Limit=t.Limit),t.PID&&(e.PID=t.PID),t.UID&&t.UID!=""&&(e.UID=t.UID),t.Lang&&t.Lang>0&&(e.Lang=t.Lang),t.Result&&t.Result!="\u4E0D\u9650"&&(e.Result=t.Result),u.$get("api/submit/status",e).then(a=>{var r,l;let i=a.data;i.code==0&&(t.Count=i.Count,o.list=i.Data,s.syncUrl(t)),u.codeProcessor((r=i==null?void 0:i.code)!=null?r:100001,(l=i==null?void 0:i.msg)!=null?l:"\u670D\u52A1\u5668\u9519\u8BEF\\\\error"),n.list.close()})},syncUrl:e=>{let a={};e.Page&&(a.Page=e.Page),e.Limit&&(a.Limit=e.Limit),e.PID&&(a.PID=e.PID),e.UID&&e.UID!=""&&(a.UID=e.UID),e.Lang&&e.Lang>0&&(a.Lang=e.Lang),e.Result&&e.Result!="\u4E0D\u9650"&&(a.Result=e.Result),u.$router.replace({path:"/Status",query:a})}}),t=Vue.reactive({PID:null,UID:"",Lang:-1,Result:"\u4E0D\u9650",Count:0,Page:1,Limit:20,PIDSetter:e=>{t.PID=e},UIDSetter:e=>{t.UID=e}}),o=Vue.reactive({list:[]});return Vue.provide("config",s),Vue.provide("query",t),Vue.onBeforeMount(()=>{u.$route.query.Page&&(t.Page=Number(u.$route.query.Page)),u.$route.query.Limit&&(t.Limit=Number(u.$route.query.Limit)),typeof u.$route.query.PID<"u"&&(t.PID=u.$route.query.PID),typeof u.$route.query.UID<"u"&&(t.UID=u.$route.query.UID),typeof u.$route.query.Result<"u"&&(t.Result=u.$route.query.Result),typeof u.$route.query.Lang<"u"&&(t.Lang=Number(u.$route.query.Lang))}),Vue.onMounted(()=>{s.update()}),(e,a)=>(Vue.openBlock(),Vue.createElementBlock("div",g,[Vue.createVNode(f,{query:Vue.unref(t)},null,8,["query"]),Vue.createVNode(c,{data:Vue.unref(o).list},null,8,["data"])]))}});const p=D(I,[["__scopeId","data-v-1204f058"]]);export{p as default};