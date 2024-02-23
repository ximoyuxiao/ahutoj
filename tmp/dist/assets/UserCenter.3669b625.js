import{_ as C,u as A,d as N,I as S,F as k,e as b,s as D}from"./index.de99a551.js";const U={class:"activityCalendar"},$={class:"left"},L={class:"right"},z=["onClick"],M={class:"levelFlagContent"},T={name:"ActivityCalendar"},P=Vue.defineComponent({...T,props:{data:null,endDate:null,width:null,height:null,cellLength:null,cellInterval:null,cellBorderRadius:null,header:null,showHeader:{type:Boolean},backgroundColor:null,colors:null,showWeekDayFlag:{type:Boolean},weekDayFlagText:null,levelMapper:null,showLevelFlag:{type:Boolean},levelFlagText:null,fontSize:null,fontColor:null,clickEvent:null,beginDate:null,levels:null,headerLength:null,weekDayFlagLength:null},setup(m){const o=m;var c=Vue.computed(()=>o);Vue.watch(c,(r,t)=>{i()},{deep:!0});var n=[];const e=Vue.reactive({data:[],beginDate:"",endDate:"",width:35,height:7,cellLength:16,cellInterval:6,cellBorderRadius:3,header:["\u4E00\u6708","\u4E8C\u6708","\u4E09\u6708","\u56DB\u6708","\u4E94\u6708","\u516D\u6708","\u4E03\u6708","\u516B\u6708","\u4E5D\u6708","\u5341\u6708","\u5341\u4E00\u6708","\u5341\u4E8C\u6708"],headerLength:[],showHeader:!0,backgroundColor:"#ffffff",colors:["#f5f5f5","#ebfaff","#e5f9ff","#c7f1ff","#86e0fe","#3ecefe"],showWeekDayFlag:!0,weekDayFlagText:["Mon","Tue","Wed","Thu","Fri","Sat","Sun"],weekDayFlagLength:[],levels:6,levelMapper:function(t){return t==0?0:t<=1?1:t<=3?2:t<=6?3:t<=9?4:5},showLevelFlag:!0,levelFlagText:["\u5C11","\u591A"],fontSize:12,fontColor:"#080808",clickEvent:function(t){}}),u=Vue.reactive({header(r){return"left:"+r.length+"px;font-size: "+e.fontSize+"px;color: "+e.fontColor},weekDay(r){return"top:"+r.length+"px;font-size: "+(e.fontSize-2)+"px;color: "+e.fontColor},content(){return"grid-template-columns: repeat("+(e.width+1)+","+(e.cellLength+e.cellInterval/2)+"px);grid-template-rows: repeat("+e.height+","+(e.cellLength+e.cellInterval/2)+"px);background-color:"+e.backgroundColor},item(r){return"width:"+e.cellLength+"px; height:"+e.cellLength+"px; background-color:"+(e==null?void 0:e.colors[e.levelMapper?e.levelMapper(r.count):0])+";border-radius:"+e.cellBorderRadius+"px;"+(r.index<0?"visibility: hidden;":"")},levelFlag(){return"grid-template-columns: repeat("+e.colors.length+","+(e.cellLength+e.cellInterval/2)+"px);grid-template-rows: repeat(1,"+(e.cellLength+e.cellInterval/2)+"px);background-color:"+e.backgroundColor},levelFlagItem(r){return"width:"+e.cellLength+"px; background-color:"+e.colors[r]+";border-radius:"+e.cellBorderRadius+"px;font-size: "+e.fontSize+"px;"}});function a(){let r=[0,31,28,31,30,31,30,31,31,30,31,30,31],t=e.width*e.height;if(!e.endDate)return;let s=e.endDate.split("-"),V=Number(s[0])-0,p=Number(s[1])-0,d=Number(s[2])-0,g=V,_=p,f=d;if(t<=d)f=1;else for(t-=d,_=(_-1+11)%12+1,f=1;t>0;){g%4==0&&g%100!=0||g%400==0?r[2]=29:r[2]=28;for(let v=_;v>0;v--)if(r[v]<=t)t-=r[v],_=(_-1+11)%12+1;else{f=r[v]-t,t=0;break}t>0&&(g-=1)}e.beginDate=g+"-"+(_<10?"0"+_:_)+"-"+(f<10?"0"+f:f)}function l(){if(!n)return;n.sort((h,F)=>h.date>F.date?1:h.date<F.date?-1:0);let r=[0,31,28,31,30,31,30,31,31,30,31,30,31],t=e.width*e.height,s=0;if(!e.beginDate)return;let V=e.beginDate.split("-"),p=Number(V[0])-0,d=Number(V[1])-0,g=Number(V[2])-0;(p%4==0&&p%100!=0||p%400==0)&&(r[2]=29),g++,s++,g>r[d]&&(g=1,d++),d>12&&(d=1,p++,p%4==0&&p%100!=0||p%400==0?r[2]=29:r[2]=28);let _=[],f=0;for(let h=0;h<t;h++){let F=p+"-"+(d<10?"0"+d:d)+"-"+(g<10?"0"+g:g),I={index:h,count:0,date:F};if(f<n.length){for(;f<n.length&&n[f].date<F;)f++;f<n.length&&n[f].date==F&&(I.count=n[f].count,f++)}_.push(I),g+=1,s++,g>r[d]&&(g=1,d+=1,d>12&&(d=1,p+=1,p%4==0&&p%100!=0||p%400==0?r[2]=29:r[2]=28),e.headerLength&&e.headerLength.push({length:(e.cellLength+e.cellInterval/2)*(s/e.height),text:e.header[d-1]}))}let v=new Date(p+"-"+d+"-"+(g-1)).getDay();for(let h=0;h<v;h++)_.unshift({index:h-v,date:"",count:0});if(e.data=_,e.showWeekDayFlag&&e.weekDayFlagText){e.weekDayFlagText.length>7&&(e.weekDayFlagText=e.weekDayFlagText.slice(0,7));for(let h=0;h<e.height;h++)h%2==0&&e.weekDayFlagLength.push({length:e.fontSize*4/2+20+e.cellInterval+h*(e.cellLength+e.cellInterval/2),text:e.weekDayFlagText[h%7]})}}function i(){if(o.endDate)e.endDate=o.endDate;else{let r=new Date;e.endDate=r.getFullYear()+"-"+(r.getMonth()+1<10?"0"+(r.getMonth()+1):r.getMonth()+1)+"-"+(r.getDate()+1<10?"0"+r.getDate():r.getDate())}o.data&&(n=o.data),o.height&&(e.height=o.height),o.width&&(e.width=o.width),o.cellLength&&(e.cellLength=o.cellLength),o.cellInterval&&(e.cellInterval=o.cellInterval),o.cellBorderRadius&&(e.cellBorderRadius=o.cellBorderRadius),o.header&&(e.header=o.header),o.showHeader&&(e.showHeader=o.showHeader),o.backgroundColor&&(e.backgroundColor=o.backgroundColor),o.colors&&(e.colors=o.colors),typeof o.showWeekDayFlag<"u"&&(e.showWeekDayFlag=o.showWeekDayFlag),o.weekDayFlagText&&(e.weekDayFlagText=o.weekDayFlagText),o.levelMapper&&(e.levelMapper=o.levelMapper),typeof o.showLevelFlag<"u"&&(e.showLevelFlag=o.showLevelFlag),o.levelFlagText&&(e.levelFlagText=o.levelFlagText),o.fontSize&&(e.fontSize=o.fontSize),o.fontColor&&(e.fontColor=o.fontColor),o.clickEvent&&(e.clickEvent=o.clickEvent),a(),l()}return Vue.onMounted(()=>{Vue.nextTick(()=>{i()})}),(r,t)=>(Vue.openBlock(),Vue.createElementBlock("div",U,[Vue.withDirectives(Vue.createElementVNode("div",$,[(Vue.openBlock(!0),Vue.createElementBlock(Vue.Fragment,null,Vue.renderList(e.weekDayFlagLength,(s,V)=>(Vue.openBlock(),Vue.createElementBlock("div",{key:V,style:Vue.normalizeStyle(u.weekDay(s))},Vue.toDisplayString(s.text),5))),128))],512),[[Vue.vShow,e.showWeekDayFlag]]),Vue.createElementVNode("div",L,[Vue.withDirectives(Vue.createElementVNode("div",{class:"header",style:Vue.normalizeStyle("height:"+e.fontSize+"px;")},[(Vue.openBlock(!0),Vue.createElementBlock(Vue.Fragment,null,Vue.renderList(e.headerLength,(s,V)=>(Vue.openBlock(),Vue.createElementBlock("div",{key:V,style:Vue.normalizeStyle(u.header(s))},Vue.toDisplayString(s.text),5))),128))],4),[[Vue.vShow,e.showHeader]]),Vue.createElementVNode("div",{class:"content",style:Vue.normalizeStyle(u.content())},[(Vue.openBlock(!0),Vue.createElementBlock(Vue.Fragment,null,Vue.renderList(e.data,(s,V)=>(Vue.openBlock(),Vue.createElementBlock("div",{class:"item",key:V,style:Vue.normalizeStyle(u.item(s)),onClick:p=>e.clickEvent?e.clickEvent(s):null},null,12,z))),128))],4),Vue.withDirectives(Vue.createElementVNode("div",M,[Vue.createElementVNode("div",{style:Vue.normalizeStyle("font-size:"+e.fontSize+"px;color: "+e.fontColor)},Vue.toDisplayString(e.levelFlagText?e.levelFlagText[0]:""),5),Vue.createElementVNode("div",{class:"levelFlag",style:Vue.normalizeStyle(u.levelFlag())},[(Vue.openBlock(!0),Vue.createElementBlock(Vue.Fragment,null,Vue.renderList(e.colors,(s,V)=>(Vue.openBlock(),Vue.createElementBlock("div",{key:V,style:Vue.normalizeStyle(u.levelFlagItem(V))},null,4))),128))],4),Vue.createElementVNode("div",{style:Vue.normalizeStyle("font-size:"+e.fontSize+"px;color: "+e.fontColor)},Vue.toDisplayString(e.levelFlagText?e.levelFlagText[1]:""),5)],512),[[Vue.vShow,e.showLevelFlag]])])]))}});const R=C(P,[["__scopeId","data-v-f7c27ab8"]]),E=m=>(Vue.pushScopeId("data-v-91e093a8"),m=m(),Vue.popScopeId(),m),j={class:"ChangeInfo"},H=E(()=>Vue.createElementVNode("div",{style:{height:"30px"}},"\u4FEE\u6539\u8D44\u6599",-1)),J=E(()=>Vue.createElementVNode("span",null,"\u6635\u79F0:\xA0",-1)),O=E(()=>Vue.createElementVNode("span",null,"\u5B66\u6821:\xA0",-1)),W=E(()=>Vue.createElementVNode("span",null,"\u73ED\u7EA7:\xA0",-1)),Y=E(()=>Vue.createElementVNode("span",null,"\u4E13\u4E1A:\xA0",-1)),K=E(()=>Vue.createElementVNode("span",null,"\u64C5\u957F:\xA0",-1)),Z=E(()=>Vue.createElementVNode("span",null,"\u90AE\u7BB1:\xA0",-1)),q=Vue.defineComponent({__name:"ChangeInfo",props:{userInfo:{default:{}},close:{type:Function,default:()=>{}}},setup(m){const o=m,{proxy:c}=Vue.getCurrentInstance();A();const n=N();var e=Vue.reactive({UID:"",UserName:"",School:"",Classes:"",Major:"",Adept:"",Email:"",init(){this.UserName=n.UserName,this.School=o.userInfo.School,this.Classes=o.userInfo.Classes,this.Major=o.userInfo.Major,this.Adept=o.userInfo.Adept,this.Email=o.userInfo.Email,u.AdeptArray=this.Adept==""?[]:this.Adept.split(";")},sendForm(){if(this.Adept=u.AdeptArray.length==0?"":u.AdeptArray.join(";"),this.UserName==""){c.elMessage({message:"\u6635\u79F0\u4E0D\u53EF\u4E3A\u7A7A",type:"warning"});return}if(!this.Email.match("^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(.[a-zA-Z0-9_-]+)+$")){c.elMessage({message:"\u8BF7\u8F93\u5165\u6B63\u786E\u7684\u90AE\u7BB1",type:"warning"});return}c.$axios.post("api/user/edit/",{UserName:this.UserName,School:this.School,Classes:this.Classes,Major:this.Major,Adept:this.Adept,Email:this.Email}).then(a=>{var i,r;let l=a.data;l.code==0?(n.UserName=this.UserName,o.userInfo.UserName=this.UserName,o.userInfo.School=this.School,o.userInfo.Classes=this.Classes,o.userInfo.Major=this.Major,o.userInfo.Adept=this.Adept,o.userInfo.Email=this.Email,o.userInfo.Vjid=this.Vjid,o.userInfo.AdeptArray=o.userInfo.Adept==""?[]:o.userInfo.Adept.split(";"),c.elNotification({message:"\u4FEE\u6539\u6210\u529F",type:"success",duration:1500}),o.close(0)):c.codeProcessor((i=l==null?void 0:l.code)!=null?i:100001,(r=l==null?void 0:l.msg)!=null?r:"\u670D\u52A1\u5668\u9519\u8BEF\\\\error")})}}),u=Vue.reactive({inputValue:"",AdeptArray:[],inputVisible:!1,handleClose(a){this.AdeptArray.splice(this.AdeptArray.indexOf(a),1)},showInput(){this.inputVisible=!0,Vue.nextTick(()=>{c.$refs.InputRef.focus()})},handleInputConfirm(){if(!!this.inputVisible){if(this.inputValue==""||this.AdeptArray.indexOf(this.inputValue)!=-1){this.inputVisible=!1,this.inputValue="",c.elMessage({message:"\u8BF7\u8F93\u5165\u6B63\u786E\u7684\u540D\u79F0\u5E76\u4E14\u4E0D\u5F97\u91CD\u590D"});return}this.AdeptArray.push(this.inputValue),this.inputVisible=!1,this.inputValue=""}},handleInputCancel(){this.inputVisible=!1,this.inputValue=""}});return Vue.onMounted(()=>{e.init()}),(a,l)=>{const i=Vue.resolveComponent("CircleClose"),r=Vue.resolveComponent("el-icon"),t=Vue.resolveComponent("el-input"),s=Vue.resolveComponent("el-tag"),V=Vue.resolveComponent("el-button"),p=Vue.resolveComponent("EditPen");return Vue.openBlock(),Vue.createElementBlock("div",j,[H,Vue.createVNode(r,{class:"close cursor_pointer",size:"30px",onClick:l[0]||(l[0]=d=>o.close(0))},{default:Vue.withCtx(()=>[Vue.createVNode(i)]),_:1}),Vue.createElementVNode("div",null,[J,Vue.createVNode(t,{modelValue:Vue.unref(e).UserName,"onUpdate:modelValue":l[1]||(l[1]=d=>Vue.unref(e).UserName=d)},null,8,["modelValue"])]),Vue.createElementVNode("div",null,[O,Vue.createVNode(t,{modelValue:Vue.unref(e).School,"onUpdate:modelValue":l[2]||(l[2]=d=>Vue.unref(e).School=d)},null,8,["modelValue"])]),Vue.createElementVNode("div",null,[W,Vue.createVNode(t,{modelValue:Vue.unref(e).Classes,"onUpdate:modelValue":l[3]||(l[3]=d=>Vue.unref(e).Classes=d)},null,8,["modelValue"])]),Vue.createElementVNode("div",null,[Y,Vue.createVNode(t,{modelValue:Vue.unref(e).Major,"onUpdate:modelValue":l[4]||(l[4]=d=>Vue.unref(e).Major=d)},null,8,["modelValue"])]),Vue.createElementVNode("div",null,[K,(Vue.openBlock(!0),Vue.createElementBlock(Vue.Fragment,null,Vue.renderList(Vue.unref(u).AdeptArray,d=>(Vue.openBlock(),Vue.createBlock(s,{key:d,style:{"min-width":"fit-content",margin:"0 1px"},closable:"","disable-transitions":!1,onClose:g=>Vue.unref(u).handleClose(d)},{default:Vue.withCtx(()=>[Vue.createTextVNode(Vue.toDisplayString(d),1)]),_:2},1032,["onClose"]))),128)),Vue.unref(u).inputVisible?(Vue.openBlock(),Vue.createBlock(t,{key:0,ref:"InputRef",modelValue:Vue.unref(u).inputValue,"onUpdate:modelValue":l[5]||(l[5]=d=>Vue.unref(u).inputValue=d),size:"small",class:"ml-1 w-20",style:{width:"100px"},onKeyup:l[6]||(l[6]=Vue.withKeys(d=>Vue.unref(u).handleInputConfirm(),["enter"])),onBlur:l[7]||(l[7]=d=>Vue.unref(u).handleInputCancel())},null,8,["modelValue"])):(Vue.openBlock(),Vue.createBlock(V,{key:1,class:"button-new-tag ml-1",size:"small",onClick:l[8]||(l[8]=d=>Vue.unref(u).showInput())},{default:Vue.withCtx(()=>[Vue.createTextVNode(" + \u6DFB\u52A0 ")]),_:1}))]),Vue.createElementVNode("div",null,[Z,Vue.createVNode(t,{modelValue:Vue.unref(e).Email,"onUpdate:modelValue":l[9]||(l[9]=d=>Vue.unref(e).Email=d)},null,8,["modelValue"])]),Vue.createElementVNode("div",{class:"btn cursor_noFocus cursor_pointer",onClick:l[10]||(l[10]=d=>Vue.unref(e).sendForm())},[Vue.createVNode(r,null,{default:Vue.withCtx(()=>[Vue.createVNode(p)]),_:1}),Vue.createTextVNode(" \xA0\u63D0\u4EA4\u4FEE\u6539 ")])])}}});const G=C(q,[["__scopeId","data-v-91e093a8"]]),y=m=>(Vue.pushScopeId("data-v-dadd75a6"),m=m(),Vue.popScopeId(),m),Q={class:"BindingCodeForce"},X=y(()=>Vue.createElementVNode("div",{class:"title"},"\u7ED1\u5B9ACodeForce",-1)),ee={key:0,class:"bindingNow"},te=y(()=>Vue.createElementVNode("span",null,"\u8D26\u53F7:\xA0",-1)),oe=y(()=>Vue.createElementVNode("span",null,"\u5BC6\u7801:\xA0",-1)),ue=Vue.defineComponent({__name:"BindingCodeForce",props:{CodeForceUser:{default:""},close:{type:Function,default:()=>{}}},setup(m){const o=m,{proxy:c}=Vue.getCurrentInstance();var n=Vue.reactive({CodeForceUser:o.CodeForceUser,CodeForcePass:"",send:()=>{if(n.CodeForceUser==""){c.elMessage({message:"ID\u4E0D\u80FD\u4E3A\u7A7A",type:"warning"});return}if(n.CodeForcePass==""){c.elMessage({message:"\u5BC6\u7801\u4E0D\u80FD\u4E3A\u7A7A",type:"warning"});return}c.$post("api/user/CodeForceBind/",{CodeForceUser:n.CodeForceUser,CodeForcePass:n.CodeForcePass}).then(e=>{var a,l;c.$log(e);let u=e.data;u.code==0&&(c.elNotification({message:"\u7ED1\u5B9A\u6210\u529F",type:"success",duration:1500}),o.close(1)),c.codeProcessor((a=u==null?void 0:u.code)!=null?a:100001,(l=u==null?void 0:u.msg)!=null?l:"\u670D\u52A1\u5668\u9519\u8BEF\\\\error")})}});return(e,u)=>{const a=Vue.resolveComponent("CircleClose"),l=Vue.resolveComponent("el-icon"),i=Vue.resolveComponent("el-input"),r=Vue.resolveComponent("Link");return Vue.openBlock(),Vue.createElementBlock("div",Q,[X,Vue.createVNode(l,{class:"close cursor_pointer",size:"30px",onClick:u[0]||(u[0]=t=>o.close(1))},{default:Vue.withCtx(()=>[Vue.createVNode(a)]),_:1}),o.CodeForceUser?(Vue.openBlock(),Vue.createElementBlock("div",ee," \u5F53\u524D\u5DF2\u7ED1\u5B9A\uFF1A"+Vue.toDisplayString(o.CodeForceUser),1)):Vue.createCommentVNode("",!0),Vue.createElementVNode("div",null,[te,Vue.createVNode(i,{modelValue:Vue.unref(n).CodeForceUser,"onUpdate:modelValue":u[1]||(u[1]=t=>Vue.unref(n).CodeForceUser=t),placeholder:"\u8F93\u5165CodeForce\u8D26\u53F7",autocomplete:"off"},null,8,["modelValue"])]),Vue.createElementVNode("div",null,[oe,Vue.createVNode(i,{modelValue:Vue.unref(n).CodeForcePass,"onUpdate:modelValue":u[2]||(u[2]=t=>Vue.unref(n).CodeForcePass=t),placeholder:"\u8F93\u5165CodeForce\u5BC6\u7801",autocomplete:"off",type:"password","show-password":""},null,8,["modelValue"])]),Vue.createElementVNode("div",{class:"btn cursor_noFocus cursor_pointer",onClick:u[3]||(u[3]=(...t)=>Vue.unref(n).send&&Vue.unref(n).send(...t))},[Vue.createVNode(l,null,{default:Vue.withCtx(()=>[Vue.createVNode(r)]),_:1}),Vue.createTextVNode(" \xA0\u7ED1\u5B9A ")])])}}});const ne=C(ue,[["__scopeId","data-v-dadd75a6"]]),x=m=>(Vue.pushScopeId("data-v-138e9ba1"),m=m(),Vue.popScopeId(),m),le={class:"BindingVJ","data-type":"form"},se=x(()=>Vue.createElementVNode("div",{class:"title"},"\u7ED1\u5B9AVJ\u8D26\u53F7",-1)),re={key:0,class:"bindingNow"},ae=x(()=>Vue.createElementVNode("span",null,"VJ ID:\xA0",-1)),ie=x(()=>Vue.createElementVNode("span",null,"VJ \u5BC6\u7801:\xA0",-1)),de=Vue.defineComponent({__name:"BindingVJudge",props:{Vjid:{default:""},close:{type:Function,default:()=>{}}},setup(m){const o=m,{proxy:c}=Vue.getCurrentInstance();var n=Vue.reactive({Vjid:o.Vjid,Vjpwd:"",isInChangeMode:!1,init(){this.Vjid=o.Vjid,this.Vjpwd="",this.isInChangeMode=!1},sendForm(){if(this.Vjid==""){c.elMessage({message:"ID\u4E0D\u80FD\u4E3A\u7A7A",type:"warning"});return}if(this.Vjpwd==""){c.elMessage({message:"\u5BC6\u7801\u4E0D\u80FD\u4E3A\u7A7A",type:"warning"});return}c.$post("api/user/vjudgeBind",{Vjid:n.Vjid,Vjpwd:n.Vjpwd}).then(e=>{var a,l;let u=e.data;u.code==0&&(c.elNotification({message:"\u7ED1\u5B9A\u6210\u529F",type:"success",duration:1500}),o.close(2)),c.codeProcessor((a=u==null?void 0:u.code)!=null?a:100001,(l=u==null?void 0:u.msg)!=null?l:"\u670D\u52A1\u5668\u9519\u8BEF\\\\error")})}});return(e,u)=>{const a=Vue.resolveComponent("CircleClose"),l=Vue.resolveComponent("el-icon"),i=Vue.resolveComponent("el-input"),r=Vue.resolveComponent("Link");return Vue.openBlock(),Vue.createElementBlock("div",le,[se,Vue.createVNode(l,{class:"close cursor_pointer",size:"30px",onClick:u[0]||(u[0]=t=>o.close(2))},{default:Vue.withCtx(()=>[Vue.createVNode(a)]),_:1}),o.Vjid?(Vue.openBlock(),Vue.createElementBlock("div",re," \u5F53\u524D\u5DF2\u7ED1\u5B9A\uFF1A"+Vue.toDisplayString(o.Vjid),1)):Vue.createCommentVNode("",!0),Vue.createElementVNode("div",null,[ae,Vue.createVNode(i,{modelValue:Vue.unref(n).Vjid,"onUpdate:modelValue":u[1]||(u[1]=t=>Vue.unref(n).Vjid=t),placeholder:"\u8F93\u5165VJudge\u8D26\u53F7"},null,8,["modelValue"])]),Vue.createElementVNode("div",null,[ie,Vue.createVNode(i,{modelValue:Vue.unref(n).Vjpwd,"onUpdate:modelValue":u[2]||(u[2]=t=>Vue.unref(n).Vjpwd=t),placeholder:"\u8F93\u5165VJudge\u5BC6\u7801",type:"password","show-password":""},null,8,["modelValue"])]),Vue.createElementVNode("div",{class:"btn cursor_noFocus cursor_pointer",onClick:u[3]||(u[3]=t=>Vue.unref(n).sendForm())},[Vue.createVNode(l,null,{default:Vue.withCtx(()=>[Vue.createVNode(r)]),_:1}),Vue.createTextVNode(" \xA0\u7ED1\u5B9A ")])])}}});const ce=C(de,[["__scopeId","data-v-138e9ba1"]]),w=m=>(Vue.pushScopeId("data-v-0a7a4c34"),m=m(),Vue.popScopeId(),m),Ve={class:"ChangePassword","data-type":"form"},pe=w(()=>Vue.createElementVNode("div",{style:{height:"30px"}},"\u4FEE\u6539\u5BC6\u7801",-1)),me=w(()=>Vue.createElementVNode("span",null,"\u65E7\u5BC6\u7801:\xA0",-1)),fe=w(()=>Vue.createElementVNode("span",null,"\u65B0\u5BC6\u7801:\xA0",-1)),ge=w(()=>Vue.createElementVNode("span",null,"\u65B0\u5BC6\u7801:\xA0",-1)),he=Vue.defineComponent({__name:"ChangePassword",props:{close:{type:Function,default:()=>{}}},setup(m){const o=m,{proxy:c}=Vue.getCurrentInstance();var n=Vue.reactive({OldPwd:"",Pwd:"",PwdAgain:"",init(){this.OldPwd="",this.Pwd="",this.PwdAgain=""},sendForm(){if(this.OldPwd==""){c.elMessage({message:"\u65E7\u5BC6\u7801\u4E0D\u80FD\u4E3A\u7A7A",type:"warning"});return}if(this.Pwd==""||this.PwdAgain==""){c.elMessage({message:"\u65B0\u5BC6\u7801\u4E0D\u80FD\u4E3A\u7A7A",type:"warning"});return}if(this.Pwd!=this.PwdAgain){c.elMessage({message:"\u4E24\u6B21\u8F93\u5165\u7684\u65B0\u5BC6\u7801\u4E0D\u540C",type:"warning"});return}c.$post("api/user/edit/pass/",{OldPwd:n.OldPwd,Pwd:n.Pwd}).then(e=>{var a,l;c.$log(e);let u=e.data;u.code==0&&(c.elNotification({message:"\u4FEE\u6539\u6210\u529F",type:"success",duration:1500}),o.close(10)),c.codeProcessor((a=u==null?void 0:u.code)!=null?a:100001,(l=u==null?void 0:u.msg)!=null?l:"\u670D\u52A1\u5668\u9519\u8BEF\\\\error")})}});return(e,u)=>{const a=Vue.resolveComponent("CircleClose"),l=Vue.resolveComponent("el-icon"),i=Vue.resolveComponent("el-input"),r=Vue.resolveComponent("EditPen");return Vue.openBlock(),Vue.createElementBlock("div",Ve,[pe,Vue.createVNode(l,{class:"close cursor_pointer",size:"30px",onClick:u[0]||(u[0]=t=>o.close(10))},{default:Vue.withCtx(()=>[Vue.createVNode(a)]),_:1}),Vue.createElementVNode("div",null,[me,Vue.createVNode(i,{modelValue:Vue.unref(n).OldPwd,"onUpdate:modelValue":u[1]||(u[1]=t=>Vue.unref(n).OldPwd=t),placeholder:"\u8F93\u5165\u65E7\u5BC6\u7801",autocomplete:"off"},null,8,["modelValue"])]),Vue.createElementVNode("div",null,[fe,Vue.createVNode(i,{modelValue:Vue.unref(n).Pwd,"onUpdate:modelValue":u[2]||(u[2]=t=>Vue.unref(n).Pwd=t),placeholder:"\u8F93\u5165\u65B0\u5BC6\u7801",autocomplete:"off",type:"password","show-password":""},null,8,["modelValue"])]),Vue.createElementVNode("div",null,[ge,Vue.createVNode(i,{modelValue:Vue.unref(n).PwdAgain,"onUpdate:modelValue":u[3]||(u[3]=t=>Vue.unref(n).PwdAgain=t),placeholder:"\u91CD\u590D\u4E00\u6B21\u65B0\u5BC6\u7801",autocomplete:"off",type:"password","show-password":""},null,8,["modelValue"])]),Vue.createElementVNode("div",{class:"btn cursor_noFocus cursor_pointer",onClick:u[4]||(u[4]=t=>Vue.unref(n).sendForm())},[Vue.createVNode(l,null,{default:Vue.withCtx(()=>[Vue.createVNode(r)]),_:1}),Vue.createTextVNode(" \xA0\u4FEE\u6539 ")])])}}});const _e=C(he,[["__scopeId","data-v-0a7a4c34"]]),B=m=>(Vue.pushScopeId("data-v-414ff7b0"),m=m(),Vue.popScopeId(),m),ve={class:"changeHeadImage"},Ce=B(()=>Vue.createElementVNode("div",{class:"title",style:{height:"30px"}}," \u4FEE\u6539\u5934\u50CF ",-1)),Ee={class:"uploadImage"},Fe={class:"preview"},we={class:"originImage"},Be={class:"img"},Ne=["src"],ye=B(()=>Vue.createElementVNode("div",{class:"hint"},"\u539F\u59CB\u56FE\u7247",-1)),xe={class:"hint"},Ie={class:"compressResult"},ke={class:"compressPercent"},De=B(()=>Vue.createElementVNode("div",null," \u538B\u7F29\u540E\u6700\u5927\u652F\u6301100KB\u5927\u5C0F\u56FE\u7247",-1)),Ae={class:"compressedImage"},Se={class:"img"},be=["src"],Ue=B(()=>Vue.createElementVNode("div",{class:"hint"},"\u81EA\u52A8\u538B\u7F29",-1)),$e={class:"hint"},Le={class:"btn cursor_pointer"},ze=Vue.defineComponent({__name:"ChangeHeadImage",props:["userInfo","close"],setup(m){const o=m,{proxy:c}=Vue.getCurrentInstance();N();var n=Vue.reactive({loading:null,originImage:null,originImageBlobURL:"",compressedImage:null,compressedImageBlobURL:"",compressPercent:"",newHeadURL:"",selectImage:e=>{if(!["image/jpeg","image/png","image/gif","image/webp"].includes(e.raw.type)){c.elMessage({message:"\u8BF7\u9009\u62E9\u6709\u6548\u7684\u56FE\u7247\u6587\u4EF6",type:"error"}),c.$refs.upload.clearFiles();return}n.loading=c.elLoading({node:document.getElementsByClassName("changeHeadImage")[0]}),c.$refs.upload.clearFiles();let a=e.raw;S.userHeadImageCompress(a).then(l=>{let i=l.data;n.originImageBlobURL=k.file2Blob(a),n.compressedImageBlobURL=k.file2Blob(i),n.originImage=a,n.compressedImage=i,n.compressPercent=((n.originImage.size-n.compressedImage.size)/n.originImage.size*100).toFixed(2),n.loading.close()})},uploadHeadImage:()=>{if(!n.originImageBlobURL||!n.compressedImageBlobURL){c.elMessage({message:"\u8BF7\u9009\u62E9\u56FE\u7247",type:"error"});return}b.uploadUserHeadImage(n.compressedImage).then(e=>{var a,l;let u=e.data;u.code==0&&(o.close(11),o.userInfo.HeadURL=u.ImageURL,c.elMessage({message:"\u4FEE\u6539\u6210\u529F",type:"success"})),c.codeProcessor((a=u==null?void 0:u.code)!=null?a:100001,(l=u==null?void 0:u.msg)!=null?l:"\u670D\u52A1\u5668\u9519\u8BEF\\\\error")})}});return Vue.onMounted(()=>{}),(e,u)=>{var p,d;const a=Vue.resolveComponent("CircleClose"),l=Vue.resolveComponent("el-icon"),i=Vue.resolveComponent("SemiSelect"),r=Vue.resolveComponent("el-divider"),t=Vue.resolveComponent("FolderOpened"),s=Vue.resolveComponent("Upload"),V=Vue.resolveComponent("el-upload");return Vue.openBlock(),Vue.createElementBlock("div",ve,[Ce,Vue.createVNode(l,{class:"close cursor_pointer",size:"30px",onClick:u[0]||(u[0]=g=>o.close(11))},{default:Vue.withCtx(()=>[Vue.createVNode(a)]),_:1}),Vue.createElementVNode("div",Ee,[Vue.createElementVNode("div",Fe,[Vue.createElementVNode("div",we,[Vue.createElementVNode("div",Be,[Vue.unref(n).originImageBlobURL?(Vue.openBlock(),Vue.createElementBlock("img",{key:0,src:Vue.unref(n).originImageBlobURL},null,8,Ne)):(Vue.openBlock(),Vue.createBlock(l,{key:1,size:"42px"},{default:Vue.withCtx(()=>[Vue.createVNode(i)]),_:1}))]),ye,Vue.createElementVNode("div",xe,Vue.toDisplayString(Vue.unref(n).originImageBlobURL?(((p=Vue.unref(n).originImage)==null?void 0:p.size)/1024).toFixed(2)+"KB":""),1)]),Vue.createElementVNode("div",Ie,[Vue.createElementVNode("div",ke,Vue.toDisplayString(Vue.unref(n).compressedImageBlobURL?"\u538B\u7F29\u7387\uFF1A"+Vue.unref(n).compressPercent+"%":""),1),Vue.createVNode(r),De]),Vue.createElementVNode("div",Ae,[Vue.createElementVNode("div",Se,[Vue.unref(n).compressedImageBlobURL?(Vue.openBlock(),Vue.createElementBlock("img",{key:0,src:Vue.unref(n).compressedImageBlobURL},null,8,be)):(Vue.openBlock(),Vue.createBlock(l,{key:1,size:"42px"},{default:Vue.withCtx(()=>[Vue.createVNode(i)]),_:1}))]),Ue,Vue.createElementVNode("div",$e,Vue.toDisplayString(Vue.unref(n).compressedImageBlobURL?(((d=Vue.unref(n).compressedImage)==null?void 0:d.size)/1024).toFixed(2)+"KB":""),1)])]),Vue.createVNode(V,{ref:"upload",class:"upload",limit:1,"auto-upload":!1,"show-file-list":!1,"on-change":Vue.unref(n).selectImage},{trigger:Vue.withCtx(()=>[Vue.createElementVNode("div",Le,[Vue.createVNode(l,null,{default:Vue.withCtx(()=>[Vue.createVNode(t)]),_:1}),Vue.createTextVNode(" \xA0\u9009\u62E9\u56FE\u7247 ")])]),default:Vue.withCtx(()=>[Vue.createElementVNode("div",{class:"btn cursor_pointer",onClick:u[1]||(u[1]=(...g)=>Vue.unref(n).uploadHeadImage&&Vue.unref(n).uploadHeadImage(...g))},[Vue.createVNode(l,null,{default:Vue.withCtx(()=>[Vue.createVNode(s)]),_:1}),Vue.createTextVNode(" \xA0\u786E\u5B9A ")])]),_:1},8,["on-change"])])])}}});const Me=C(ze,[["__scopeId","data-v-414ff7b0"]]),Te={class:"userCenter"},Pe={class:"infoBox"},Re=["src"],je={class:"user"},He=["src"],Je={class:"changImage"},Oe={class:"username"},We={class:"acStatus"},Ye={class:"acCount"},Ke={style:{color:"#00CC00"}},Ze={class:"submittedCount"},qe={style:{color:"#0000FF"}},Ge={class:"rating"},Qe={class:"userInfo"},Xe={class:"contentBox"},et={class:"leftBox"},tt={class:"rightBox"},ot={class:"activityCalendarBox"},ut=Vue.defineComponent({__name:"UserCenter",setup(m){const{proxy:o}=Vue.getCurrentInstance(),c=N(),n=A();var e=Vue.reactive({UID:"",HeadURL:"",UserName:"",School:"",Classes:"",Major:"",Adept:"",Email:"",CodeForceUser:"",Vjid:"",AdeptArray:[],Solved:0,Submited:0,Rating:0,copy:t=>{console.log(t),e.UID=t.UID,e.HeadURL=t.HeadURL,e.UserName=t.UserName,e.School=t.School,e.Classes=t.Classes,e.Major=t.Major,e.Adept=t.Adept,e.Email=t.Email,e.CodeForceUser=t.CodeForceUser,e.Vjid=t.Vjid,e.Submited=t.Submited,e.Solved=t.Solved,e.Rating=t.Rating,e.AdeptArray=e.Adept==""?[]:e.Adept.split(";"),console.log(e)}});function u(){o.$get("api/user/info?uid="+c.UID).then(t=>{var V,p;let s=t.data;s.code==0&&(console.log(s),e.copy(s),c.updateData(s)),o.codeProcessor((V=s==null?void 0:s.code)!=null?V:100001,(p=s==null?void 0:s.msg)!=null?p:"\u670D\u52A1\u5668\u9519\u8BEF\\\\error")})}var a=Vue.reactive({data:[],endDate:"",width:35,height:7,cellLength:20,cellInterval:6,cellBorderRadius:3,showHeader:!0,colors:["#f5f5f5","#ebfaff","#e5f9ff","#c7f3ff","#86e0fe","#3ecefe"],levelMapper:t=>t==0?0:t<=1?1:t<=3?2:t<=6?3:t<=9?4:5,showLevelFlag:!0,levelFlagText:["\u5C11","\u591A"],fontSize:12,fontColor:"#707070",clickEvent:function(s){},init(t){this.endDate=o.Utils.TimeTools.timestampToDate(Date.now(),2);let s=new Map;for(let V in t){let p=o.Utils.TimeTools.timestampToDate(t[V].SubmitTime,2);s.has(p)?s.set(p,s.get(p)+1):s.set(p,1)}s.forEach((V,p)=>{let d={date:p,count:V};this.data.push(d)})}});function l(){let t=sessionStorage.getItem("userSubmitData");if(t){let s=JSON.parse(t);if(s.UID==c.UID&&Date.now()-s.saveTime<6e5){a.init(s.data);return}}o.$get("api/submit/status",{UID:c.UID}).then(s=>{var p,d;let V=s.data;V.code==0&&(o.Buffer.UserCenter.submitData(V.Data,c.UID),a.init(V.Data)),o.codeProcessor((p=V==null?void 0:V.code)!=null?p:100001,(d=V==null?void 0:V.msg)!=null?d:"\u670D\u52A1\u5668\u9519\u8BEF\\\\error")})}n.$subscribe((t,s)=>{s.theme==1?(a.fontColor="#707070",a.colors=["#dde0e4","#c5f6fa","#99e9f2","#66d9e8","#3bc9db","#22b8cf"]):(a.fontColor="#cdcdcd",a.colors=["#d6d6d6","#e599f7","#da77f2","#be4bdb","#ae3ec9","#862e9c"])},{deep:!0,immediate:!0});var i=Vue.reactive({showChangeInfo:!1,showBindingCodeForce:!1,showBindingVJudge:!1,showChangePassword:!1,showChangeHeadImage:!1,init(){this.showChangeInfo=!1,this.showBindingCodeForce=!1,this.showBindingVJudge=!1,this.showChangePassword=!1,this.showChangeHeadImage=!1},show:t=>{switch(t){case 0:i.showChangeInfo=!0;break;case 1:i.showBindingCodeForce=!0;break;case 2:i.showBindingVJudge=!0;break;case 10:i.showChangePassword=!0;break;case 11:i.showChangeHeadImage=!0}},close:t=>{switch(t){case 0:i.showChangeInfo=!1;break;case 1:i.showBindingCodeForce=!1;break;case 2:i.showBindingVJudge=!1;break;case 10:i.showChangePassword=!1;break;case 11:i.showChangeHeadImage=!1}}});function r(t){if(t==0)return"color:#000000;";if(t<1200)return"color:#C0C0C0;";if(t>=1200&&t<1400)return"color:#00FF00;";if(t>=1400&&t<1600)return"color:#00FFFF;";if(t>=1600&&t<1900)return"color:#0000FF;";if(t>=1900&&t<2100)return"color:#FF00FF;";if(t>=2100&&t<2400)return"color:#FF8000;";if(t>=2400&&t<2600)return"color:#FF0000;";if(t>=2600)return"color:FFFF00;"}return Vue.onMounted(()=>{if(!c.isLogin){o.$router.replace({path:"/"});return}u(),l()}),(t,s)=>{const V=Vue.resolveComponent("Upload"),p=Vue.resolveComponent("el-icon"),d=Vue.resolveComponent("el-divider"),g=Vue.resolveComponent("el-tag"),_=Vue.resolveComponent("Setting");return Vue.openBlock(),Vue.createElementBlock("div",Te,[Vue.createElementVNode("div",Pe,[Vue.createElementVNode("img",{class:"filter cursor_noFocus",src:Vue.unref(e).HeadURL?Vue.unref(D)+Vue.unref(e).HeadURL:Vue.unref(o).Utils.DefaultHeadImage.show(Vue.unref(e).UID),alt:""},null,8,Re),Vue.createElementVNode("div",je,[Vue.createElementVNode("div",{id:"userImg",class:"cursor_noFocus cursor_pointer",onClick:s[0]||(s[0]=f=>Vue.unref(i).show(11))},[Vue.createElementVNode("img",{src:Vue.unref(e).HeadURL?Vue.unref(D)+Vue.unref(e).HeadURL:Vue.unref(o).Utils.DefaultHeadImage.show(Vue.unref(e).UID)},null,8,He),Vue.createElementVNode("div",Je,[Vue.createVNode(p,{size:"42px"},{default:Vue.withCtx(()=>[Vue.createVNode(V)]),_:1})])]),Vue.createElementVNode("div",Oe,Vue.toDisplayString(Vue.unref(c).UserName),1),Vue.createElementVNode("div",We,[Vue.createElementVNode("div",Ye,[Vue.createTextVNode(" AC:\xA0\xA0\xA0"),Vue.createElementVNode("span",Ke,Vue.toDisplayString(Vue.unref(e).Solved),1)]),Vue.createElementVNode("div",Ze,[Vue.createTextVNode(" Submit:\xA0\xA0\xA0"),Vue.createElementVNode("span",qe,Vue.toDisplayString(Vue.unref(e).Submited),1)]),Vue.createElementVNode("div",Ge,[Vue.createTextVNode(" Rating:\xA0\xA0\xA0"),Vue.createElementVNode("span",{style:Vue.normalizeStyle(r(Vue.unref(e).Rating?Vue.unref(e).Rating:0))},Vue.toDisplayString(Vue.unref(e).Rating?Vue.unref(e).Rating:0),5)])])]),Vue.createVNode(d,{style:{margin:"2px"}}),Vue.createElementVNode("div",Qe,[Vue.createElementVNode("div",null,"\u5B66\u6821:\xA0"+Vue.toDisplayString(Vue.unref(e).School),1),Vue.createElementVNode("div",null,"\u73ED\u7EA7:\xA0"+Vue.toDisplayString(Vue.unref(e).Classes),1),Vue.createElementVNode("div",null,"\u4E13\u4E1A:\xA0"+Vue.toDisplayString(Vue.unref(e).Major),1),Vue.createElementVNode("div",null,[Vue.createTextVNode(" \u64C5\u957F:\xA0 "),(Vue.openBlock(!0),Vue.createElementBlock(Vue.Fragment,null,Vue.renderList(Vue.unref(e).AdeptArray,f=>(Vue.openBlock(),Vue.createBlock(g,{key:f},{default:Vue.withCtx(()=>[Vue.createTextVNode(Vue.toDisplayString(f),1)]),_:2},1024))),128))]),Vue.createElementVNode("div",null,"\u90AE\u7BB1:\xA0"+Vue.toDisplayString(Vue.unref(e).Email),1),Vue.createElementVNode("div",{class:"set cursor_pointer",onClick:s[1]||(s[1]=f=>Vue.unref(i).show(0))},[Vue.createVNode(p,{size:"45px"},{default:Vue.withCtx(()=>[Vue.createVNode(_)]),_:1})])])]),Vue.createVNode(Vue.Transition,{"enter-active-class":"animate__animated animate__zoomIn","leave-active-class":"animate__animated animate__zoomOut"},{default:Vue.withCtx(()=>[Vue.unref(i).showChangeInfo?(Vue.openBlock(),Vue.createBlock(G,{key:0,userInfo:Vue.unref(e),close:Vue.unref(i).close},null,8,["userInfo","close"])):Vue.createCommentVNode("",!0)]),_:1}),Vue.createVNode(Vue.Transition,{"enter-active-class":"animate__animated animate__zoomIn","leave-active-class":"animate__animated animate__zoomOut"},{default:Vue.withCtx(()=>[Vue.unref(i).showChangeHeadImage?(Vue.openBlock(),Vue.createBlock(Me,{key:0,userInfo:Vue.unref(e),close:Vue.unref(i).close},null,8,["userInfo","close"])):Vue.createCommentVNode("",!0)]),_:1}),Vue.createElementVNode("div",Xe,[Vue.createElementVNode("div",et,[Vue.createElementVNode("div",{class:"functionBtn",onClick:s[2]||(s[2]=f=>Vue.unref(i).show(1))},Vue.toDisplayString(Vue.unref(e).CodeForceUser?`CodeForce:
`+Vue.unref(e).CodeForceUser:"\u7ED1\u5B9ACodeForce"),1),Vue.createElementVNode("div",{class:"functionBtn",onClick:s[3]||(s[3]=f=>Vue.unref(i).show(2))},Vue.toDisplayString(Vue.unref(e).Vjid?`VJudge:
`+Vue.unref(e).Vjid:"\u7ED1\u5B9AVJudge"),1),Vue.createVNode(d,{style:{margin:"2px"}}),Vue.createElementVNode("div",{class:"functionBtn",onClick:s[4]||(s[4]=f=>Vue.unref(i).show(10))}," \u4FEE\u6539\u5BC6\u7801 ")]),Vue.createElementVNode("div",tt,[Vue.createVNode(Vue.Transition,{"enter-active-class":"animate__animated animate__zoomIn","leave-active-class":"animate__animated animate__zoomOut"},{default:Vue.withCtx(()=>[Vue.unref(i).showBindingCodeForce?(Vue.openBlock(),Vue.createBlock(ne,{key:0,CodeForceUser:Vue.unref(e).CodeForceUser,close:Vue.unref(i).close},null,8,["CodeForceUser","close"])):Vue.createCommentVNode("",!0)]),_:1}),Vue.createVNode(Vue.Transition,{"enter-active-class":"animate__animated animate__zoomIn","leave-active-class":"animate__animated animate__zoomOut"},{default:Vue.withCtx(()=>[Vue.unref(i).showBindingVJudge?(Vue.openBlock(),Vue.createBlock(ce,{key:0,Vjid:Vue.unref(e).Vjid,close:Vue.unref(i).close},null,8,["Vjid","close"])):Vue.createCommentVNode("",!0)]),_:1}),Vue.createVNode(Vue.Transition,{"enter-active-class":"animate__animated animate__zoomIn","leave-active-class":"animate__animated animate__zoomOut"},{default:Vue.withCtx(()=>[Vue.unref(i).showChangePassword?(Vue.openBlock(),Vue.createBlock(_e,{key:0,close:Vue.unref(i).close},null,8,["close"])):Vue.createCommentVNode("",!0)]),_:1}),Vue.createElementVNode("div",ot,[Vue.createVNode(R,{data:Vue.unref(a).data,backgroundColor:"#ffffff00",width:Vue.unref(a).width,height:Vue.unref(a).height,cellLength:Vue.unref(a).cellLength,cellInterval:Vue.unref(a).cellInterval,cellBorderRadius:Vue.unref(a).cellBorderRadius,fontSize:Vue.unref(a).fontSize,fontColor:Vue.unref(a).fontColor,showLevelFlag:Vue.unref(a).showLevelFlag,colors:Vue.unref(a).colors,showWeekDayFlag:!0,levelMapper:Vue.unref(a).levelMapper,levelFlagText:Vue.unref(a).levelFlagText},null,8,["data","width","height","cellLength","cellInterval","cellBorderRadius","fontSize","fontColor","showLevelFlag","colors","levelMapper","levelFlagText"])])])])])}}});const lt=C(ut,[["__scopeId","data-v-c510ef7a"]]);export{lt as default};
