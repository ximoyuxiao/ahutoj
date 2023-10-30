import{c as F,I as B,e as b,s as g,_ as N}from"./index.de99a551.js";const o=p=>(Vue.pushScopeId("data-v-a3cca2df"),p=p(),Vue.popScopeId(),p),I={class:"addProblem"},O=o(()=>Vue.createElementVNode("span",null,"\u6807\u9898",-1)),C=o(()=>Vue.createElementVNode("span",null,"\u65F6\u95F4\u9650\u5236",-1)),v=o(()=>Vue.createElementVNode("span",null,"\xA0\u6BEB\u79D2",-1)),L=o(()=>Vue.createElementVNode("span",null,"\u5185\u5B58\u9650\u5236",-1)),x=o(()=>Vue.createElementVNode("span",null,"\xA0MiB",-1)),T=o(()=>Vue.createElementVNode("span",null,"\u6765\u6E90",-1)),D=o(()=>Vue.createElementVNode("span",null,"\u53EF\u89C1",-1)),y=o(()=>Vue.createElementVNode("span",null,"\u6807\u7B7E",-1)),P=o(()=>Vue.createElementVNode("br",null,null,-1)),A=o(()=>Vue.createElementVNode("span",null,"\u6587\u672C\u7C7B\u578B",-1)),k={key:0,class:"markdown"},M={key:1,class:"normal"},S=o(()=>Vue.createElementVNode("span",null,"\u9898\u76EE\u63CF\u8FF0",-1)),U=o(()=>Vue.createElementVNode("span",null,"\u8F93\u5165\u63CF\u8FF0",-1)),R=o(()=>Vue.createElementVNode("span",null,"\u8F93\u51FA\u63CF\u8FF0",-1)),w=o(()=>Vue.createElementVNode("span",null,"\u8F93\u5165\u6837\u4F8B",-1)),h=o(()=>Vue.createElementVNode("span",null,"\u8F93\u51FA\u6837\u4F8B",-1)),H=o(()=>Vue.createElementVNode("span",null,"\u63D0\u793A",-1)),z=Vue.defineComponent({__name:"AddProblem",setup(p){const{proxy:i}=Vue.getCurrentInstance(),a=F();var c={toolbar:["bold","underline","italic","-","strikeThrough","title","sub","sup","quote","unorderedList","orderedList","-","codeRow","code","link","image","table","mermaid","katex","-","revoke","next","save","=","preview","htmlPreview","catalog"]},e=Vue.reactive({PID:"",Title:"",Description:"",Input:"",Output:"",SampleInput:"",SampleOutput:"",LimitTime:1e3,LimitMemory:128,Hit:"",Label:"",Origin:-1,OriginPID:"",ContentType:1,Visible:1,ContentTypes:[{label:"PlainText",value:a.PROBLEM_CONTENTTYPE_NORMAL},{label:"MarkDown",value:a.PROBLEM_CONTENTTYPE_MARKDOWN}],Origins:[{label:"Local",value:a.PROBLEM_ORIGIN_LOCAL},{label:"CodeForce",value:a.PROBLEM_ORIGIN_CF},{label:"AtCoder",value:a.PROBLEM_ORIGIN_ATCODER},{label:"\u6D1B\u8C37",value:a.PROBLEM_ORIGIN_LUOGU}],OriginPlaceHolder:{[a.PROBLEM_ORIGIN_CF]:"\u7ADE\u8D5BID + \u9898\u76EE\u7F16\u53F7 \u5982\uFF1A1069A",[a.PROBLEM_ORIGIN_ATCODER]:"\u7ADE\u8D5BID + \u9898\u76EE\u7F16\u53F7 \u5982\uFF1Aabc277_e"},Visibles:[{label:"\u53EF\u89C1",value:a.PROBLEM_VISIBLE},{label:"\u4E0D\u53EF\u89C1",value:a.PROBLEM_UNVISIBLE}],init(){e.PID="",e.Title="",e.Description="",e.Input="",e.Output="",e.SampleInput="",e.SampleOutput="",e.LimitTime=1e3,e.LimitMemory=128,e.Hit="",e.Label="",e.Origin=0,e.OriginPID="",e.ContentType=-1,e.Visible=1},updateImg:r=>{if(r.length==0){i.elMessage({message:"\u4E0A\u4F20\u5185\u5BB9\u4E3A\u7A7A\uFF01",type:"warning"});return}ElementPlus.ElMessageBox.confirm("\u786E\u5B9A\u4E0A\u4F20\u5417\uFF1F","\u6CE8\u610F",{confirmButtonText:"\u786E\u5B9A",cancelButtonText:"\u53D6\u6D88",type:"warning"}).then(()=>{r.forEach(u=>{E(u)})})}});function E(r){B.problemImageCompress(r).then(u=>{u.code==0&&b.uploadProblemImage(u.data).then(n=>{var V,d;let t=n.data;if(t.code==0){let s=t.ImageURL;e.Description+=`
![](${g}${s})`,i.elMessage({message:`
              <strong>${r.name}\u4E0A\u4F20\u6210\u529F</strong><br/>
              <span>\u5DF2\u538B\u7F29:${(r.size/1024).toFixed(2)}KB->${(u.data.size/1024).toFixed(2)}KB</span>
            `,type:"success",duration:5e3,dangerouslyUseHTMLString:!0})}i.codeProcessor((V=t==null?void 0:t.code)!=null?V:100001,(d=t==null?void 0:t.msg)!=null?d:"\u670D\u52A1\u5668\u9519\u8BEF\\\\error")})})}function _(){i.$post("api/problem/add/",{Title:e.Title,Description:e.Description,Input:e.Input,Output:e.Output,SampleInput:e.SampleInput,SampleOutput:e.SampleOutput,LimitTime:e.LimitTime,LimitMemory:e.LimitMemory,Hit:e.Hit,Label:e.Label,Origin:e.Origin,OriginPID:e.OriginPID,ContentType:e.ContentType,Visible:e.Visible}).then(r=>{var n,t;let u=r.data;if(u.code==0){i.elMessage({message:"\u6DFB\u52A0\u6210\u529F!",type:"success"});let V=u.PID;e.Origin==-1&&i.$router.push({path:"/Admin/ProblemEdit/EditProblemJudgeFile",query:{PID:V}});return}i.codeProcessor((n=u==null?void 0:u.code)!=null?n:100001,(t=u==null?void 0:u.msg)!=null?t:"\u670D\u52A1\u5668\u9519\u8BEF\\\\error")})}return(r,u)=>{const n=Vue.resolveComponent("el-input"),t=Vue.resolveComponent("el-row"),V=Vue.resolveComponent("el-input-number"),d=Vue.resolveComponent("el-option"),s=Vue.resolveComponent("el-select"),f=Vue.resolveComponent("el-button");return Vue.openBlock(),Vue.createElementBlock("div",I,[Vue.createVNode(t,null,{default:Vue.withCtx(()=>[O,Vue.createVNode(n,{modelValue:Vue.unref(e).Title,"onUpdate:modelValue":u[0]||(u[0]=l=>Vue.unref(e).Title=l),class:"problemTitle"},null,8,["modelValue"])]),_:1}),Vue.createVNode(t,null,{default:Vue.withCtx(()=>[C,Vue.createVNode(V,{modelValue:Vue.unref(e).LimitTime,"onUpdate:modelValue":u[1]||(u[1]=l=>Vue.unref(e).LimitTime=l),min:200,max:1e4,step:200,"controls-position":"right"},null,8,["modelValue"]),v,L,Vue.createVNode(V,{modelValue:Vue.unref(e).LimitMemory,"onUpdate:modelValue":u[2]||(u[2]=l=>Vue.unref(e).LimitMemory=l),min:64,max:1024,step:64,"controls-position":"right"},null,8,["modelValue"]),x]),_:1}),Vue.createVNode(t,null,{default:Vue.withCtx(()=>[T,Vue.createVNode(s,{modelValue:Vue.unref(e).Origin,"onUpdate:modelValue":u[3]||(u[3]=l=>Vue.unref(e).Origin=l),class:"m-2",placeholder:"Select"},{default:Vue.withCtx(()=>[(Vue.openBlock(!0),Vue.createElementBlock(Vue.Fragment,null,Vue.renderList(Vue.unref(e).Origins,(l,m)=>(Vue.openBlock(),Vue.createBlock(d,{key:l.label,label:l.label,value:l.value},null,8,["label","value"]))),128))]),_:1},8,["modelValue"]),Vue.withDirectives(Vue.createVNode(n,{modelValue:Vue.unref(e).OriginPID,"onUpdate:modelValue":u[4]||(u[4]=l=>Vue.unref(e).OriginPID=l),placeholder:Vue.unref(e).OriginPlaceHolder[Vue.unref(e).Origin],class:"originInfo"},null,8,["modelValue","placeholder"]),[[Vue.vShow,Vue.unref(e).Origin!=-1]])]),_:1}),Vue.createVNode(t,null,{default:Vue.withCtx(()=>[D,Vue.createVNode(s,{modelValue:Vue.unref(e).Visible,"onUpdate:modelValue":u[5]||(u[5]=l=>Vue.unref(e).Visible=l),class:"m-2",placeholder:"Select"},{default:Vue.withCtx(()=>[(Vue.openBlock(!0),Vue.createElementBlock(Vue.Fragment,null,Vue.renderList(Vue.unref(e).Visibles,(l,m)=>(Vue.openBlock(),Vue.createBlock(d,{key:l.label,label:l.label,value:l.value},null,8,["label","value"]))),128))]),_:1},8,["modelValue"])]),_:1}),Vue.createVNode(t,null,{default:Vue.withCtx(()=>[y,Vue.createVNode(n,{modelValue:Vue.unref(e).Label,"onUpdate:modelValue":u[6]||(u[6]=l=>Vue.unref(e).Label=l),placeholder:"\u8BF7\u8F93\u5165\u7684\u6BCF\u4E2A\u6807\u7B7E\u4E4B\u95F4\u7528';'\u9694\u5F00",class:"tagList"},null,8,["modelValue"])]),_:1}),P,Vue.createVNode(t,null,{default:Vue.withCtx(()=>[A,Vue.createVNode(s,{modelValue:Vue.unref(e).ContentType,"onUpdate:modelValue":u[7]||(u[7]=l=>Vue.unref(e).ContentType=l),class:"m-2",placeholder:"Select"},{default:Vue.withCtx(()=>[(Vue.openBlock(!0),Vue.createElementBlock(Vue.Fragment,null,Vue.renderList(Vue.unref(e).ContentTypes,(l,m)=>(Vue.openBlock(),Vue.createBlock(d,{key:l.label,label:l.label,value:l.value},null,8,["label","value"]))),128))]),_:1},8,["modelValue"])]),_:1}),Vue.unref(e).ContentType==1?(Vue.openBlock(),Vue.createElementBlock("div",k,[Vue.createVNode(Vue.unref(MdEditorV3),{modelValue:Vue.unref(e).Description,"onUpdate:modelValue":u[8]||(u[8]=l=>Vue.unref(e).Description=l),toolbars:Vue.unref(c).toolbar,"on-upload-img":Vue.unref(e).updateImg},null,8,["modelValue","toolbars","on-upload-img"])])):(Vue.openBlock(),Vue.createElementBlock("div",M,[Vue.createElementVNode("div",null,[S,Vue.createVNode(n,{modelValue:Vue.unref(e).Description,"onUpdate:modelValue":u[9]||(u[9]=l=>Vue.unref(e).Description=l),type:"textarea",autosize:"",class:"plainText"},null,8,["modelValue"])]),Vue.createElementVNode("div",null,[U,Vue.createVNode(n,{modelValue:Vue.unref(e).Input,"onUpdate:modelValue":u[10]||(u[10]=l=>Vue.unref(e).Input=l),type:"textarea",autosize:"",class:"plainText"},null,8,["modelValue"])]),Vue.createElementVNode("div",null,[R,Vue.createVNode(n,{modelValue:Vue.unref(e).Output,"onUpdate:modelValue":u[11]||(u[11]=l=>Vue.unref(e).Output=l),type:"textarea",autosize:"",class:"plainText"},null,8,["modelValue"])]),Vue.createElementVNode("div",null,[w,Vue.createVNode(n,{modelValue:Vue.unref(e).SampleInput,"onUpdate:modelValue":u[12]||(u[12]=l=>Vue.unref(e).SampleInput=l),type:"textarea",autosize:"",class:"plainText"},null,8,["modelValue"])]),Vue.createElementVNode("div",null,[h,Vue.createVNode(n,{modelValue:Vue.unref(e).SampleOutput,"onUpdate:modelValue":u[13]||(u[13]=l=>Vue.unref(e).SampleOutput=l),type:"textarea",autosize:"",class:"plainText"},null,8,["modelValue"])]),Vue.createElementVNode("div",null,[H,Vue.createVNode(n,{modelValue:Vue.unref(e).Hit,"onUpdate:modelValue":u[14]||(u[14]=l=>Vue.unref(e).Hit=l),type:"textarea",autosize:"",class:"plainText"},null,8,["modelValue"])])])),Vue.createVNode(f,{class:"addProblemButton",type:"primary",round:"",onClick:u[15]||(u[15]=l=>_())},{default:Vue.withCtx(()=>[Vue.createTextVNode(" \u6DFB \u52A0 ")]),_:1})])}}});const G=N(z,[["__scopeId","data-v-a3cca2df"]]);export{G as default};