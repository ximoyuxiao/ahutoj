import{c as L,I as T,e as I,s as O,_ as v}from "./index.7fe3d568.js";const n= p=>(Vue.pushScopeId("data-v-780ee488"),p=p(),Vue.popScopeId(),p),y=n(()=>Vue.createElementVNode("span",null,"\u6807\u9898",-1)),D=n(()=>Vue.createElementVNode("span",null,"\u65F6\u95F4\u9650\u5236/ms",-1)),w=n(()=>Vue.createElementVNode("span",{style:{"margin-left":"50px"}},"\u5185\u5B58\u9650\u5236/MiB",-1)),P=n(()=>Vue.createElementVNode("span",null,"\u6807\u7B7E",-1)),k=n(()=>Vue.createElementVNode("span",null,"\u6765\u6E90",-1)),h=n(()=>Vue.createElementVNode("span",null,"\u53EF\u89C1",-1)),A=n(()=>Vue.createElementVNode("span",null,"\u6587\u672C\u7C7B\u578B",-1)),M={key:0},S={class:"markdown Top"},U={key:1},R={class:"normal"},z=n(()=>Vue.createElementVNode("span",null,"\u9898\u76EE\u63CF\u8FF0",-1)),H=n(()=>Vue.createElementVNode("span",null,"\u8F93\u5165\u63CF\u8FF0",-1)),$=n(()=>Vue.createElementVNode("span",null,"\u8F93\u51FA\u63CF\u8FF0",-1)),G=n(()=>Vue.createElementVNode("span",null,"\u8F93\u5165\u6837\u4F8B",-1)),K=n(()=>Vue.createElementVNode("span",null,"\u8F93\u51FA\u6837\u4F8B",-1)),q=n(()=>Vue.createElementVNode("span",null,"\u63D0\u793A",-1)),Y={class:"addProblem"},J=Vue.defineComponent({__name:"AddProblem",setup(p){const{proxy:V}=Vue.getCurrentInstance(),a=L(),m=Vue.ref(!1);var f={toolbar:["bold","underline","italic","-","strikeThrough","title","sub","sup","quote","unorderedList","orderedList","-","codeRow","code","link","image","table","mermaid","katex","-","revoke","next","save","=","preview","htmlPreview","catalog"]},e=Vue.reactive({PID:"",Title:"",Description:"",Input:"",Output:"",SampleInput:"",SampleOutput:"",LimitTime:1e3,LimitMemory:128,Hit:"",Label:"",Origin:-1,OriginPID:"",ContentType:1,Visible:1,ContentTypes:[{label:"PlainText",value:a.PROBLEM_CONTENTTYPE_NORMAL},{label:"MarkDown",value:a.PROBLEM_CONTENTTYPE_MARKDOWN}],Origins:[{label:"Local",value:a.PROBLEM_ORIGIN_LOCAL},{label:"CodeForce",value:a.PROBLEM_ORIGIN_CF},{label:"AtCoder",value:a.PROBLEM_ORIGIN_ATCODER},{label:"\u6D1B\u8C37",value:a.PROBLEM_ORIGIN_LUOGU}],OriginPlaceHolder:{[a.PROBLEM_ORIGIN_CF]:"\u7ADE\u8D5BID + \u9898\u76EE\u7F16\u53F7 \u5982\uFF1A1069A",[a.PROBLEM_ORIGIN_ATCODER]:"\u7ADE\u8D5BID + \u9898\u76EE\u7F16\u53F7 \u5982\uFF1Aabc277_e"},Visibles:[{label:"\u53EF\u89C1",value:a.PROBLEM_VISIBLE},{label:"\u4E0D\u53EF\u89C1",value:a.PROBLEM_UNVISIBLE}],init(){e.PID="",e.Title="",e.Description="",e.Input="",e.Output="",e.SampleInput="",e.SampleOutput="",e.LimitTime=1e3,e.LimitMemory=128,e.Hit="",e.Label="",e.Origin=0,e.OriginPID="",e.ContentType=-1,e.Visible=1},updateImg: r=>{if(r.length==0){V.elMessage({message:"\u4E0A\u4F20\u5185\u5BB9\u4E3A\u7A7A\uFF01",type:"warning"});return}ElementPlus.ElMessageBox.confirm("\u786E\u5B9A\u4E0A\u4F20\u5417\uFF1F","\u6CE8\u610F",{confirmButtonText:"\u786E\u5B9A",cancelButtonText:"\u53D6\u6D88",type:"warning"}).then(()=>{r.forEach(u=>{E(u)})})}});function E(r){T.problemImageCompress(r).then(u=>{u.code==0&&I.uploadProblemImage(u.data).then(i=>{var o,s;let l=i.data;if(l.code==0){let d=l.ImageURL;e.Description+=`
![](${O}${d})`,V.elMessage({message:`
              <strong>${r.name}\u4E0A\u4F20\u6210\u529F</strong><br/>
              <span>\u5DF2\u538B\u7F29:${(r.size/1024).toFixed(2)}KB->${(u.data.size/1024).toFixed(2)}KB</span>
            `,type:"success",duration:5e3,dangerouslyUseHTMLString:!0})}V.codeProcessor((o=l==null?void 0:l.code)!=null?o:100001,(s=l==null?void 0:l.msg)!=null?s:"\u670D\u52A1\u5668\u9519\u8BEF\\\\error")})})}function F(){V.$post("api/problem/add/",{Title:e.Title,Description:e.Description,Input:e.Input,Output:e.Output,SampleInput:e.SampleInput,SampleOutput:e.SampleOutput,LimitTime:e.LimitTime,LimitMemory:e.LimitMemory,Hit:e.Hit,Label:e.Label,Origin:e.Origin,OriginPID:e.OriginPID,ContentType:e.ContentType,Visible:e.Visible}).then(r=>{var i,l;let u=r.data;if(u.code==0){V.elMessage({message:"\u6DFB\u52A0\u6210\u529F!",type:"success"});let o=u.PID;e.Origin==-1&&V.$router.push({path:"/Admin/ProblemEdit/EditProblemJudgeFile",query:{PID:o}});return}V.codeProcessor((i=u==null?void 0:u.code)!=null?i:100001,(l=u==null?void 0:u.msg)!=null?l:"\u670D\u52A1\u5668\u9519\u8BEF\\\\error")})}return(r,u)=>{const i=Vue.resolveComponent("el-drawer"),l=Vue.resolveComponent("el-input"),o=Vue.resolveComponent("el-row"),s=Vue.resolveComponent("el-input-number"),d=Vue.resolveComponent("el-option"),c=Vue.resolveComponent("el-select"),N=Vue.resolveComponent("el-header"),g=Vue.resolveComponent("FullScreen"),b=Vue.resolveComponent("el-icon"),_=Vue.resolveComponent("el-button"),x=Vue.resolveComponent("el-main"),B=Vue.resolveComponent("el-container");return Vue.openBlock(),Vue.createElementBlock(Vue.Fragment,null,[Vue.createVNode(i,{modelValue:m.value,"onUpdate:modelValue":u[1]||(u[1]=t=>m.value=t),title:Vue.unref(e).Title,fullsreen:"true",size:"100%",direction:"btt"},{default:Vue.withCtx(()=>[Vue.createVNode(Vue.unref(MdEditorV3),{modelValue:Vue.unref(e).Description,"onUpdate:modelValue":u[0]||(u[0]=t=>Vue.unref(e).Description=t),toolbars:Vue.unref(f).toolbar,"on-upload-img":Vue.unref(e).updateImg,style:{height:"calc(100vh - 120px)"}},null,8,["modelValue","toolbars","on-upload-img"])]),_:1},8,["modelValue","title"]),Vue.createVNode(B,null,{default:Vue.withCtx(()=>[Vue.createVNode(N,{class:"Container"},{default:Vue.withCtx(()=>[Vue.createVNode(o,null,{default:Vue.withCtx(()=>[y,Vue.createVNode(l,{modelValue:Vue.unref(e).Title,"onUpdate:modelValue":u[2]||(u[2]=t=>Vue.unref(e).Title=t),class:"problemTitle Left"},null,8,["modelValue"])]),_:1}),Vue.createVNode(o,{class:"Top"},{default:Vue.withCtx(()=>[D,Vue.createVNode(s,{modelValue:Vue.unref(e).LimitTime,"onUpdate:modelValue":u[3]||(u[3]=t=>Vue.unref(e).LimitTime=t),min:200,max:1e4,step:200,"controls-position":"right",class:"Left"},null,8,["modelValue"]),w,Vue.createVNode(s,{modelValue:Vue.unref(e).LimitMemory,"onUpdate:modelValue":u[4]||(u[4]=t=>Vue.unref(e).LimitMemory=t),min:64,max:1024,step:64,"controls-position":"right",class:"Left"},null,8,["modelValue"])]),_:1}),Vue.createVNode(o,{class:"Top"},{default:Vue.withCtx(()=>[P,Vue.createVNode(l,{modelValue:Vue.unref(e).Label,"onUpdate:modelValue":u[5]||(u[5]=t=>Vue.unref(e).Label=t),placeholder:"\u6BCF\u4E2A\u6807\u7B7E\u7528\u82F1\u6587\u7684 ; \u9694\u5F00",class:"tagList Left"},null,8,["modelValue"])]),_:1}),Vue.createVNode(o,{class:"Top"},{default:Vue.withCtx(()=>[k,Vue.createVNode(c,{modelValue:Vue.unref(e).Origin,"onUpdate:modelValue":u[6]||(u[6]=t=>Vue.unref(e).Origin=t),class:"m-2 Left",placeholder:"Select"},{default:Vue.withCtx(()=>[(Vue.openBlock(!0),Vue.createElementBlock(Vue.Fragment,null,Vue.renderList(Vue.unref(e).Origins,(t,C)=>(Vue.openBlock(),Vue.createBlock(d,{key:t.label,label:t.label,value:t.value},null,8,["label","value"]))),128))]),_:1},8,["modelValue"]),Vue.withDirectives(Vue.createVNode(l,{modelValue:Vue.unref(e).OriginPID,"onUpdate:modelValue":u[7]||(u[7]=t=>Vue.unref(e).OriginPID=t),placeholder:Vue.unref(e).OriginPlaceHolder[Vue.unref(e).Origin],class:"originInfo"},null,8,["modelValue","placeholder"]),[[Vue.vShow,Vue.unref(e).Origin!=-1]])]),_:1}),Vue.createVNode(o,{class:"Top"},{default:Vue.withCtx(()=>[h,Vue.createVNode(c,{modelValue:Vue.unref(e).Visible,"onUpdate:modelValue":u[8]||(u[8]=t=>Vue.unref(e).Visible=t),class:"m-2 Left",placeholder:"Select"},{default:Vue.withCtx(()=>[(Vue.openBlock(!0),Vue.createElementBlock(Vue.Fragment,null,Vue.renderList(Vue.unref(e).Visibles,(t,C)=>(Vue.openBlock(),Vue.createBlock(d,{key:t.label,label:t.label,value:t.value},null,8,["label","value"]))),128))]),_:1},8,["modelValue"])]),_:1})]),_:1}),Vue.createVNode(x,{class:"Container"},{default:Vue.withCtx(()=>[Vue.createVNode(o,null,{default:Vue.withCtx(()=>[A,Vue.createVNode(c,{modelValue:Vue.unref(e).ContentType,"onUpdate:modelValue":u[9]||(u[9]=t=>Vue.unref(e).ContentType=t),class:"m-2 Left",placeholder:"Select"},{default:Vue.withCtx(()=>[(Vue.openBlock(!0),Vue.createElementBlock(Vue.Fragment,null,Vue.renderList(Vue.unref(e).ContentTypes,(t,C)=>(Vue.openBlock(),Vue.createBlock(d,{key:t.label,label:t.label,value:t.value},null,8,["label","value"]))),128))]),_:1},8,["modelValue"]),Vue.unref(e).ContentType==1?(Vue.openBlock(),Vue.createBlock(_,{key:0,class:"Left",type:"primary",onClick:u[10]||(u[10]=t=>m.value=!0)},{default:Vue.withCtx(()=>[Vue.createVNode(b,{size:"19px"},{default:Vue.withCtx(()=>[Vue.createVNode(g)]),_:1}),Vue.createTextVNode(" \xA0\u5168\u5C4F\u7F16\u8F91 ")]),_:1})):Vue.createCommentVNode("",!0)]),_:1}),Vue.unref(e).ContentType==1?(Vue.openBlock(),Vue.createElementBlock("div",M,[Vue.createElementVNode("div",S,[Vue.createVNode(Vue.unref(MdEditorV3),{modelValue:Vue.unref(e).Description,"onUpdate:modelValue":u[11]||(u[11]=t=>Vue.unref(e).Description=t),toolbars:Vue.unref(f).toolbar,"on-upload-img":Vue.unref(e).updateImg,"preview-only":"",style:{"max-height":"600px"}},null,8,["modelValue","toolbars","on-upload-img"])])])):(Vue.openBlock(),Vue.createElementBlock("div",U,[Vue.createElementVNode("div",R,[Vue.createVNode(o,{class:"Top"},{default:Vue.withCtx(()=>[z,Vue.createVNode(l,{modelValue:Vue.unref(e).Description,"onUpdate:modelValue":u[12]||(u[12]=t=>Vue.unref(e).Description=t),type:"textarea",autosize:"",class:"plainText Left"},null,8,["modelValue"])]),_:1}),Vue.createVNode(o,{class:"Top"},{default:Vue.withCtx(()=>[H,Vue.createVNode(l,{modelValue:Vue.unref(e).Input,"onUpdate:modelValue":u[13]||(u[13]=t=>Vue.unref(e).Input=t),type:"textarea",autosize:"",class:"plainText Left"},null,8,["modelValue"])]),_:1}),Vue.createVNode(o,{class:"Top"},{default:Vue.withCtx(()=>[$,Vue.createVNode(l,{modelValue:Vue.unref(e).Output,"onUpdate:modelValue":u[14]||(u[14]=t=>Vue.unref(e).Output=t),type:"textarea",autosize:"",class:"plainText Left"},null,8,["modelValue"])]),_:1}),Vue.createVNode(o,{class:"Top"},{default:Vue.withCtx(()=>[G,Vue.createVNode(l,{modelValue:Vue.unref(e).SampleInput,"onUpdate:modelValue":u[15]||(u[15]=t=>Vue.unref(e).SampleInput=t),type:"textarea",autosize:"",class:"plainText Left"},null,8,["modelValue"])]),_:1}),Vue.createVNode(o,{class:"Top"},{default:Vue.withCtx(()=>[K,Vue.createVNode(l,{modelValue:Vue.unref(e).SampleOutput,"onUpdate:modelValue":u[16]||(u[16]=t=>Vue.unref(e).SampleOutput=t),type:"textarea",autosize:"",class:"plainText Left"},null,8,["modelValue"])]),_:1}),Vue.createVNode(o,{class:"Top"},{default:Vue.withCtx(()=>[q,Vue.createVNode(l,{modelValue:Vue.unref(e).Hit,"onUpdate:modelValue":u[17]||(u[17]=t=>Vue.unref(e).Hit=t),type:"textarea",autosize:"",class:"plainText",style:{"margin-left":"50px"}},null,8,["modelValue"])]),_:1})])]))]),_:1})]),_:1}),Vue.createElementVNode("div",Y,[Vue.createVNode(_,{class:"addProblemButton Top",type:"primary",round:"",onClick:u[18]||(u[18]=t=>F())},{default:Vue.withCtx(()=>[Vue.createTextVNode(" \u6DFB \u52A0 ")]),_:1})])],64)}}});const j=v(J,[["__scopeId","data-v-780ee488"]]);export{j as default};