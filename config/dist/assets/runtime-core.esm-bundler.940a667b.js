import {g as w, i as H, N as m, p as k, r as x} from "./index.7fe3d568.js";

function f(e, r, o, t) {
    let s;
    try {
        s = t ? e(...t) : e()
    } catch (n) {
        h(n, r, o)
    }
    return s
}

function d(e, r, o, t) {
    if (H(e)) {
        const n = f(e, r, o, t);
        return n && w(n) && n.catch(c => {
            h(c, r, o)
        }), n
    }
    const s = [];
    for (let n = 0; n < e.length; n++) s.push(d(e[n], r, o, t));
    return s
}

function h(e, r, o, t = !0) {
    const s = r ? r.vnode : null;
    if (r) {
        let n = r.parent;
        const c = r.proxy, l = o;
        for (; n;) {
            const u = n.ec;
            if (u) {
                for (let a = 0; a < u.length; a++) if (u[a](e, c, l) === !1) return
            }
            n = n.parent
        }
        const p = r.appContext.config.errorHandler;
        if (p) {
            f(p, null, 10, [e, c, l]);
            return
        }
    }
    C(e, o, s, t)
}

function C(e, r, o, t = !0) {
    console.error(e)
}

function g(e, r, o = i, t = !1) {
    if (o) {
        const s = o[e] || (o[e] = []), n = r.__weh || (r.__weh = (...c) => {
            if (o.isUnmounted) return;
            k(), E(o);
            const l = d(r, o, e, c);
            return M(), x(), l
        });
        return t ? s.unshift(n) : s.push(n), n
    }
}

const v = e => (r, o = i) => g(e, (...t) => r(...t), o), W = v("m");

function I() {
    return {
        app: null,
        config: {
            isNativeTag: m,
            performance: !1,
            globalProperties: {},
            optionMergeStrategies: {},
            errorHandler: void 0,
            warnHandler: void 0,
            compilerOptions: {}
        },
        mixins: [],
        components: {},
        directives: {},
        provides: Object.create(null),
        optionsCache: new WeakMap,
        propsCache: new WeakMap,
        emitsCache: new WeakMap
    }
}

I();
let i = null;
const E = e => {
    i = e, e.scope.on()
}, M = () => {
    i && i.scope.off(), i = null
};
export {W as o};
