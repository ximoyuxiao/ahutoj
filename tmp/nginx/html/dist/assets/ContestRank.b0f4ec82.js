import{a as B,_ as T}from"./index.de99a551.js";const N=D=>(Vue.pushScopeId("data-v-30c9ffea"),D=D(),Vue.popScopeId(),D),x={class:"ContestRank",ref:"ContestRank"},E={class:"contestInfo"},w={class:"title"},A={class:"content"},L={class:"time"},F=N(()=>Vue.createElementVNode("br",null,null,-1)),U={class:"header"},M=Vue.createStaticVNode('<div style="width:70px;" data-v-30c9ffea>\u6392\u540D</div><div style="width:300px;" data-v-30c9ffea>\u7528\u6237\u540D</div><div style="width:120px;" data-v-30c9ffea>\u73ED\u7EA7</div><div style="width:90px;" data-v-30c9ffea>AC\u6570</div><div style="width:80px;" data-v-30c9ffea>\u7F5A\u65F6</div>',5),O={key:0,style:{width:"70px"}},R={class:"medalIcon"},z={key:1,style:{width:"70px"}},$={style:{width:"300px","text-align":"start"}},X={style:{width:"120px"}},j={style:{width:"90px"}},q={style:{width:"80px"}},J={key:0},G={key:1},H={class:"header"},K=Vue.createStaticVNode('<div style="width:70px;" data-v-30c9ffea>\u6392\u540D</div><div style="width:300px;" data-v-30c9ffea>\u7528\u6237\u540D</div><div style="width:120px;" data-v-30c9ffea>\u73ED\u7EA7</div><div style="width:90px;" data-v-30c9ffea>\u603B\u5F97\u5206</div><div style="width:80px;" data-v-30c9ffea>\u603B\u7528\u65F6</div>',5),Q={key:0,style:{width:"70px"}},W={class:"medalIcon"},Y={key:1,style:{width:"70px"}},Z={style:{width:"300px","text-align":"start"}},ee={style:{width:"120px"}},te={style:{width:"90px"}},oe={style:{width:"80px"}},ne={key:0},re=Vue.defineComponent({__name:"ContestRank",setup(D){const{proxy:c}=Vue.getCurrentInstance(),b=B();var k=Vue.ref(1),u=Vue.reactive({rankList:[],rankListOfOI:[],calculateOfACM(t){var n;let s=new Map,m=new Array;for(let e in l.Data)s[l.Data[e].PID]={index:-1,Time:268435455},m.push(l.Data[e].PID);for(let e in t){let o={ACNumber:0,AllSubmit:0,CENumber:0,Uclass:"",Problems:[],Uname:"",UserID:"",TimePenalty:0,ProblemsMap:new Map,medal:0};o.ACNumber=t[e].ACNumber,o.AllSubmit=t[e].AllSubmit,o.CENumber=t[e].CENumber,o.Uname=t[e].Uname,o.UserID=t[e].UserID,o.Uclass=t[e].Uclass;let V=0;for(let d in t[e].Problems){let p=t[e].Problems[d];if(p.PID=="")continue;p.Status=="AC"&&(V+=Number((p.Time/1e3).toFixed(0)),p.Time<((n=s[p.PID])==null?void 0:n.Time)&&(s[p.PID]={index:Number(e),Time:p.Time}));let f={Time:p.Time,SubmitNumber:p.SubmitNumber,Status:p.Status,Pioneer:!1};o.ProblemsMap[p.PID]=f}o.TimePenalty=(t[e].AllSubmit-t[e].CENumber-t[e].ACNumber)*900+V,u.rankList.push(o)}for(let e in s)s[e].index!=-1&&(u.rankList[s[e].index].ProblemsMap[e].Pioneer=!0);u.rankList.sort((e,o)=>e.ACNumber<o.ACNumber?1:e.ACNumber==o.ACNumber?e.TimePenalty<o.TimePenalty?-1:1:-1);let r=Number((u.rankList.length*.1).toFixed(0)),a=Number((u.rankList.length*.3).toFixed(0)),i=Number((u.rankList.length*.6).toFixed(0));r=r<1?1:r,a=a<2?2:a,i=i<3?3:i;for(let e=0;e<u.rankList.length;e++)if(e<r)u.rankList[e].medal=1;else if(e<a)u.rankList[e].medal=2;else if(e<i)u.rankList[e].medal=3;else break;c.Buffer.ContestRank.rankData(u.rankList,l.CID)},calculateOfOI(t){var n;console.log(t);let s=new Map,m=new Array;for(let e in l.Data)s[l.Data[e].PID]={index:-1,Time:268435455},m.push(l.Data[e].PID);for(let e in t){let o={Solved:0,Score:0,Problems:[],Uname:"",UserID:"",Uclass:"",AllTime:0,ProblemsMap:new Map,medal:0};o.Solved=t[e].AcNumber,o.Uname=t[e].Uname,o.UserID=t[e].UserID,o.Uclass=t[e].Uclass;for(let V in t[e].Problems){let d=t[e].Problems[V];if(d.PID=="")continue;d.Score==100&&d.Time<((n=s[d.PID])==null?void 0:n.Time)&&(s[d.PID]={index:Number(e),Time:d.Time});let p={PID:d.PID,Time:d.Time,Score:d.Score,Submited:d.Submited,Pioneer:!1};o.ProblemsMap[d.PID]=p,o.AllTime+=d.Time,o.Score+=d.Score}u.rankListOfOI.push(o)}for(let e in s)s[e].index!=-1&&(u.rankListOfOI[s[e].index].ProblemsMap[e].Pioneer=!0);u.rankListOfOI.sort((e,o)=>e.Score<o.Score?1:e.Score==o.Score?e.AllTime<o.AllTime?-1:1:-1);let r=Number((u.rankListOfOI.length*.1).toFixed(0)),a=Number((u.rankListOfOI.length*.3).toFixed(0)),i=Number((u.rankListOfOI.length*.6).toFixed(0));r=r<1?1:r,a=a<2?2:a,i=i<3?3:i;for(let e=0;e<u.rankList.length;e++)if(e<r)u.rankListOfOI[e].medal=1;else if(e<a)u.rankListOfOI[e].medal=2;else if(e<i)u.rankListOfOI[e].medal=3;else break;c.Buffer.ContestRank.rankData(u.rankListOfOI,l.CID),console.log(u.rankListOfOI)}});function y(){var t;const s=XLSX.utils.book_new(),m="rank";var r;if(l.Type==1){console.log(u.rankList);var a=["\u6392\u540D","\u7528\u6237\u540D","\u73ED\u7EA7","AC\u6570","\u7F5A\u65F6"];for(var i in l.Data)a.push(c.Utils.TSBaseTools.numberToAlpha(Number(i)+1));var n=[];console.log(a),console.log(n),n.push(a);for(var i in u.rankList){var e=u.rankList[i],o=[];o.push(Number(i)+1),o.push(e.Uname+"("+e.UserID+")"),o.push(e.Uclass),o.push(e.ACNumber),o.push(c.Utils.TimeTools.timestampToInterval(e.TimePenalty*1e3,2));for(var V in l.Data){var d=e.ProblemsMap[l.Data[V].PID];if(!d){o.push("");continue}if(d.Status=="AC"){o.push(d.Status+"("+d.Time+")");continue}o.push("("+d.SubmitNumber+")")}n.push(o)}console.log(n),r=XLSX.utils.aoa_to_sheet(n),console.log(r)}if(l.Type==2){console.log(u.rankListOfOI);var a=["\u6392\u540D","\u7528\u6237\u540D","\u73ED\u7EA7","\u603B\u5F97\u5206","\u603B\u7528\u65F6"];for(var i in l.Data)a.push(c.Utils.TSBaseTools.numberToAlpha(Number(i)+1));var n=[];n.push(a);for(var i in u.rankListOfOI){var p=u.rankListOfOI[i],o=[];o.push(Number(i)+1),o.push(p.Uname+"("+p.UserID+")"),o.push(p.Uclass),o.push(p.Score),o.push(p.AllTime);for(var V in l.Data){var f=p.ProblemsMap[l.Data[V].PID];if(!f){o.push("");continue}o.push(f.Score+"("+f.Time+"ms)")}n.push(o)}console.log(n),r=XLSX.utils.aoa_to_sheet(n),console.log(r)}XLSX.utils.book_append_sheet(s,r,m);const h=XLSX.write(s,{bookType:"xlsx",type:"binary"});console.log(h),t=new Blob([v(h)],{type:"application/octet-stream"}),console.log(t),_(l.CID+":"+l.Title+"\u7684\u6392\u540D.xlsx",t)}function v(t){const s=new ArrayBuffer(t.length),m=new Uint8Array(s);for(let r=0;r!==t.length;++r)m[r]=t.charCodeAt(r)&255;return s}function _(t,s){const m=document.createElement("a");m.href=window.URL.createObjectURL(s),m.download=t,m.click()}async function g(){if(!c.$route.params.CID){c.elMessage({message:"\u8DF3\u8F6C\u5730\u5740\u9519\u8BEF\uFF0C\u8BF7\u91CD\u8BD5",type:"warning"});return}let t=c.elLoading({node:c.$refs.ContestRank}),s=c.$route.params.CID,m=b.getContestRouterData(s),r=m==null?void 0:m.Pass;await I(s,r);let a=sessionStorage.getItem("contestRankData"+l.CID);if(a&&(a=JSON.parse(a),a.CID==l.CID&&Date.now()-a.saveTime<1e4)){u.rankList=a.data,t.close();return}c.$get("api/contest/"+s+"/rank",{Pass:r}).then(i=>{c.$log(i.data),l.Type==1?u.calculateOfACM(i.data.Data):l.Type==2&&u.calculateOfOI(i.data.Data),t.close()})}var l=Vue.reactive({Data:[],length:0,BeginTime:0,CID:null,Type:1,Description:"",EndTime:0,IsPublic:1,Pass:null,Size:0,Title:"",UID:"",copy(t){let s=t.Problems.split(",");for(let m in s)for(let r in t.Data)if(s[m]==t.Data[r].PID){l.Data.push({PID:t.Data[r].PID,Title:t.Data[r].Title,ACNum:t.Data[r].ACNum,SubmitNum:t.Data[r].SubmitNum});break}l.length=t.length,l.BeginTime=t.BeginTime,l.CID=t.CID,l.Type=t.Type,l.Description=t.Description,l.EndTime=t.EndTime,l.IsPublic=t.IsPublic,l.Pass=t.Pass,l.Size=t.Size,l.Title=t.Title,l.UID=t.UID}});async function I(t,s){await c.$get("api/contest/"+t,{Pass:s}).then(m=>{var a,i;let r=m.data;r.code==0?(l.copy(r),k.value=l.Type):r.code==160504?c.$router.push({path:"/Contests"}):r.code==160503&&c.$router.push({path:"/Contests"}),c.codeProcessor((a=r==null?void 0:r.code)!=null?a:100001,(i=r==null?void 0:r.msg)!=null?i:"\u670D\u52A1\u5668\u9519\u8BEF\\\\error")})}function P(t){if(t){if(t.Pioneer)return"background-color:#2f9e44 !important; color: #eaeaea;";if(t.SubmitNumber>0)return t.Status=="AC"?"background-color:#7ace27 !important; color: #FEFEFE;":"background-color:#e74a23 !important; color: #eaeaea;"}}function S(t){if(t){if(t.Pioneer)return"background-color:#2f9e44 !important; color: #eaeaea;";if(t.Submited)return t.Score==100?"background-color:#7ace27 !important; color: #FEFEFE;":"background-color:#e74a23 !important; color: #eaeaea;"}}function C(){let t={};if(c.$route.params.CID)t.CID=c.$route.params.CID;else{c.elMessage({message:"\u6570\u636E\u5F02\u5E38\uFF0C\u8BF7\u91CD\u65B0\u8FDB\u5165\u6BD4\u8D5B\u754C\u9762",type:"warning"});return}c.$router.push({name:"Contest",params:t})}return Vue.onMounted(()=>{g()}),(t,s)=>{const m=Vue.resolveComponent("Back"),r=Vue.resolveComponent("el-icon"),a=Vue.resolveComponent("el-button"),i=Vue.resolveComponent("Medal");return Vue.openBlock(),Vue.createElementBlock("div",x,[Vue.createElementVNode("div",{class:"contestID cursor_pointer",onClick:C},[Vue.createVNode(r,{size:"32px"},{default:Vue.withCtx(()=>[Vue.createVNode(m)]),_:1}),Vue.createElementVNode("div",null," #"+Vue.toDisplayString(Vue.unref(l).CID),1)]),Vue.createElementVNode("div",E,[Vue.createElementVNode("div",w,Vue.toDisplayString(Vue.unref(l).Title),1),Vue.createElementVNode("div",A,Vue.toDisplayString(Vue.unref(l).Description),1),Vue.createElementVNode("div",L,Vue.toDisplayString(Vue.unref(c).Utils.TimeTools.timestampToTime(Vue.unref(l).BeginTime))+" - "+Vue.toDisplayString(Vue.unref(c).Utils.TimeTools.timestampToTime(Vue.unref(l).EndTime)),1)]),Vue.createVNode(a,{style:{margin:"auto"},onClick:s[0]||(s[0]=n=>y())},{default:Vue.withCtx(()=>[Vue.createTextVNode("\u4E0B\u8F7D\u7ADE\u8D5B\u6392\u540D")]),_:1}),F,Vue.unref(k)==1?(Vue.openBlock(),Vue.createElementBlock(Vue.Fragment,{key:0},[Vue.createElementVNode("div",U,[M,(Vue.openBlock(!0),Vue.createElementBlock(Vue.Fragment,null,Vue.renderList(Vue.unref(l).Data,(n,e)=>(Vue.openBlock(),Vue.createElementBlock("div",{key:e,style:{width:"80px"}},Vue.toDisplayString(Vue.unref(c).Utils.TSBaseTools.numberToAlpha(e+1)+"("+n.ACNum+"/"+n.SubmitNum+")"),1))),128))]),(Vue.openBlock(!0),Vue.createElementBlock(Vue.Fragment,null,Vue.renderList(Vue.unref(u).rankList,(n,e)=>(Vue.openBlock(),Vue.createElementBlock("div",{class:"item",key:e},[n.medal!=0?(Vue.openBlock(),Vue.createElementBlock("div",O,[Vue.createElementVNode("div",R,[n.medal==1?(Vue.openBlock(),Vue.createBlock(r,{key:0,size:"36px",color:"#fabd08"},{default:Vue.withCtx(()=>[Vue.createVNode(i)]),_:1})):n.medal==2?(Vue.openBlock(),Vue.createBlock(r,{key:1,size:"36px",color:"#d6d6d6"},{default:Vue.withCtx(()=>[Vue.createVNode(i)]),_:1})):(Vue.openBlock(),Vue.createBlock(r,{key:2,size:"36px",color:"#c57120"},{default:Vue.withCtx(()=>[Vue.createVNode(i)]),_:1}))]),Vue.createTextVNode(" "+Vue.toDisplayString(e+1),1)])):(Vue.openBlock(),Vue.createElementBlock("div",z,Vue.toDisplayString(e+1),1)),Vue.createElementVNode("div",$,Vue.toDisplayString(n.Uname)+"("+Vue.toDisplayString(n.UserID)+")",1),Vue.createElementVNode("div",X,Vue.toDisplayString(n.Uclass),1),Vue.createElementVNode("div",j,Vue.toDisplayString(n.ACNumber),1),Vue.createElementVNode("div",q,Vue.toDisplayString(Vue.unref(c).Utils.TimeTools.timestampToInterval(n.TimePenalty*1e3,2)),1),(Vue.openBlock(!0),Vue.createElementBlock(Vue.Fragment,null,Vue.renderList(Vue.unref(l).Data,(o,V)=>(Vue.openBlock(),Vue.createElementBlock("div",{class:"problemStatus",key:V,style:Vue.normalizeStyle("width:80px;"+P(n.ProblemsMap[o.PID]))},[n.ProblemsMap[o.PID]&&n.ProblemsMap[o.PID].Status=="AC"?(Vue.openBlock(),Vue.createElementBlock("div",J,Vue.toDisplayString((n.ProblemsMap[o.PID].Time/1e3).toFixed(0)),1)):Vue.createCommentVNode("",!0),n.ProblemsMap[o.PID]&&n.ProblemsMap[o.PID].SubmitNumber>1||n.ProblemsMap[o.PID]&&n.ProblemsMap[o.PID].SubmitNumber>=1&&n.ProblemsMap[o.PID].Status!="AC"?(Vue.openBlock(),Vue.createElementBlock("div",G," (-"+Vue.toDisplayString(n.ProblemsMap[o.PID].Status=="AC"?n.ProblemsMap[o.PID].SubmitNumber-1:n.ProblemsMap[o.PID].SubmitNumber)+") ",1)):Vue.createCommentVNode("",!0)],4))),128))]))),128))],64)):Vue.createCommentVNode("",!0),Vue.unref(k)==2?(Vue.openBlock(),Vue.createElementBlock(Vue.Fragment,{key:1},[Vue.createElementVNode("div",H,[K,(Vue.openBlock(!0),Vue.createElementBlock(Vue.Fragment,null,Vue.renderList(Vue.unref(l).Data,(n,e)=>(Vue.openBlock(),Vue.createElementBlock("div",{key:e,style:{width:"80px"}},Vue.toDisplayString(Vue.unref(c).Utils.TSBaseTools.numberToAlpha(e+1)),1))),128))]),(Vue.openBlock(!0),Vue.createElementBlock(Vue.Fragment,null,Vue.renderList(Vue.unref(u).rankListOfOI,(n,e)=>(Vue.openBlock(),Vue.createElementBlock("div",{class:"item",key:e},[n.medal!=0?(Vue.openBlock(),Vue.createElementBlock("div",Q,[Vue.createElementVNode("div",W,[n.medal==1?(Vue.openBlock(),Vue.createBlock(r,{key:0,size:"36px",color:"#fabd08"},{default:Vue.withCtx(()=>[Vue.createVNode(i)]),_:1})):n.medal==2?(Vue.openBlock(),Vue.createBlock(r,{key:1,size:"36px",color:"#d6d6d6"},{default:Vue.withCtx(()=>[Vue.createVNode(i)]),_:1})):(Vue.openBlock(),Vue.createBlock(r,{key:2,size:"36px",color:"#c57120"},{default:Vue.withCtx(()=>[Vue.createVNode(i)]),_:1}))]),Vue.createTextVNode(" "+Vue.toDisplayString(e+1),1)])):(Vue.openBlock(),Vue.createElementBlock("div",Y,Vue.toDisplayString(e+1),1)),Vue.createElementVNode("div",Z,Vue.toDisplayString(n.Uname)+"("+Vue.toDisplayString(n.UserID)+")",1),Vue.createElementVNode("div",ee,Vue.toDisplayString(n.Uclass),1),Vue.createElementVNode("div",te,Vue.toDisplayString(n.Score),1),Vue.createElementVNode("div",oe,Vue.toDisplayString(n.AllTime),1),(Vue.openBlock(!0),Vue.createElementBlock(Vue.Fragment,null,Vue.renderList(Vue.unref(l).Data,(o,V)=>(Vue.openBlock(),Vue.createElementBlock("div",{class:"problemStatus",key:V,style:Vue.normalizeStyle("width:80px;"+S(n.ProblemsMap[o.PID]))},[n.ProblemsMap[o.PID]?(Vue.openBlock(),Vue.createElementBlock("div",ne,Vue.toDisplayString(n.ProblemsMap[o.PID].Score)+" ("+Vue.toDisplayString(n.ProblemsMap[o.PID].Time)+"ms) ",1)):Vue.createCommentVNode("",!0)],4))),128))]))),128))],64)):Vue.createCommentVNode("",!0)],512)}}});const ue=T(re,[["__scopeId","data-v-30c9ffea"]]);export{ue as default};
