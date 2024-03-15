import {_} from "./index.7fe3d568.js";

const V = d => (Vue.pushScopeId("data-v-5fd1bc98"), d = d(), Vue.popScopeId(), d),
    m = V(() => Vue.createElementVNode("span", null, "Run ID", -1)),
    D = V(() => Vue.createElementVNode("span", null, "\u9898\u76EEID", -1)),
    I = V(() => Vue.createElementVNode("span", null, "\u7528\u6237ID", -1)),
    f = V(() => Vue.createElementVNode("span", null, "\u7ADE\u8D5BID", -1)),
    C = V(() => Vue.createElementVNode("div", {class: "Top"}, '"0"\u6216\u8005\u7A7A\u4EE3\u8868\u4E0D\u586B', -1)),
    N = Vue.defineComponent({
        __name: "Rejudge", setup(d) {
            const {proxy: a} = Vue.getCurrentInstance();
            var e = Vue.reactive({
                SID: 0, PID: "", UID: "", CID: 0, init() {
                    this.SID = 0, this.PID = "", this.UID = "", this.CID = 0
                }
            });

            function r() {
                let l = {};
                e.SID && (l.SID = e.SID), e.PID && (l.PID = e.PID), e.UID && (l.UID = e.UID), e.CID && (l.CID = e.CID), a.$post("api/submit/rejudge/", l).then(t => {
                    var n, s;
                    console.log(t);
                    let u = t.data;
                    u.code == 0 && a.elMessage({
                        message: "\u91CD\u5224\u6210\u529F!",
                        type: "success"
                    }), a.codeProcessor((n = u == null ? void 0 : u.code) != null ? n : 100001, (s = u == null ? void 0 : u.msg) != null ? s : "\u670D\u52A1\u5668\u9519\u8BEF\\\\error")
                })
            }

            return (l, t) => {
                const u = Vue.resolveComponent("el-input-number"), n = Vue.resolveComponent("el-row"),
                    s = Vue.resolveComponent("el-input"), c = Vue.resolveComponent("el-main"),
                    p = Vue.resolveComponent("el-button"), i = Vue.resolveComponent("el-container");
                return Vue.openBlock(), Vue.createBlock(i, {direction: "vertical"}, {
                    default: Vue.withCtx(() => [Vue.createVNode(c, {class: "Container"}, {
                        default: Vue.withCtx(() => [Vue.createVNode(n, null, {
                            default: Vue.withCtx(() => [m, Vue.createVNode(u, {
                                modelValue: Vue.unref(e).SID,
                                "onUpdate:modelValue": t[0] || (t[0] = o => Vue.unref(e).SID = o),
                                min: 0,
                                class: "Left"
                            }, null, 8, ["modelValue"])]), _: 1
                        }), Vue.createVNode(n, {class: "Top"}, {
                            default: Vue.withCtx(() => [D, Vue.createVNode(s, {
                                modelValue: Vue.unref(e).PID,
                                "onUpdate:modelValue": t[1] || (t[1] = o => Vue.unref(e).PID = o),
                                style: {width: "200px"},
                                class: "Left"
                            }, null, 8, ["modelValue"])]), _: 1
                        }), Vue.createVNode(n, {class: "Top"}, {
                            default: Vue.withCtx(() => [I, Vue.createVNode(s, {
                                modelValue: Vue.unref(e).UID,
                                "onUpdate:modelValue": t[2] || (t[2] = o => Vue.unref(e).UID = o),
                                style: {width: "200px"},
                                class: "Left"
                            }, null, 8, ["modelValue"])]), _: 1
                        }), Vue.createVNode(n, {class: "Top"}, {
                            default: Vue.withCtx(() => [f, Vue.createVNode(u, {
                                modelValue: Vue.unref(e).CID,
                                "onUpdate:modelValue": t[3] || (t[3] = o => Vue.unref(e).CID = o),
                                min: 0,
                                class: "Left"
                            }, null, 8, ["modelValue"])]), _: 1
                        }), C]), _: 1
                    }), Vue.createVNode(p, {
                        type: "warning",
                        onClick: t[4] || (t[4] = o => r()),
                        class: "rejudgeButton Top"
                    }, {default: Vue.withCtx(() => [Vue.createTextVNode(" \u91CD\u5224 ")]), _: 1})]), _: 1
                })
            }
        }
    });
const x = _(N, [["__scopeId", "data-v-5fd1bc98"]]);
export {x as default};
