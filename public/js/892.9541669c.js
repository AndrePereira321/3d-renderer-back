"use strict";(self["webpackChunk_3d_renderer"]=self["webpackChunk_3d_renderer"]||[]).push([[892],{5117:function(r,e,s){var o=s(7139);const i=(0,o.Q_)("forms",{state:()=>({forms:{}}),getters:{getFormFields:r=>e=>r.forms[e]?.fields,getFormRawValues:r=>e=>{if(!r.forms[e])return;const s={};for(const o in r.forms[e].fields)s[o]=r.forms[e].fields[o].value;return s},getErrorsStatus:r=>e=>{const s=[];if(r.forms[e]){for(const o in r.forms[e].fields){const i=r.forms[e].fields[o];1==i.error&&s.push({fieldName:o,fieldCurrentValue:i.value,fieldOriginalValue:i.originalValue,errorDescription:i.errorDescription??"UNKNOWN"})}return s}},getField:r=>(e,s)=>r.forms[e]?.fields[s],getFieldValue:r=>(e,s)=>r.forms[e]?.fields[s]?.value,getFieldErrorStatus:r=>(e,s)=>r.forms[e]?.fields[s]?.error,getFieldErrorMessage:r=>(e,s)=>r.forms[e]?.fields[s]?.errorDescription},actions:{setFieldValue(r,e,s,o,i,t){this.forms[r]||(this.forms[r]={fields:{}}),this.forms[r].fields[e]||(this.forms[r].fields[e]={originalValue:s});const n=this.forms[r].fields[e];"undefined"!=typeof i&&(n.error=i,n.errorDescription=t),n.value=s,n.required=o},setFieldError(r,e,s,o){const i=this.forms[r].fields[e];i&&(i.error=s,i.errorDescription=o)},validateForm(r){const e=this.forms[r].fields??{};for(const s in e){const r=e[s];r.required&&!r.value&&(r.error=!0,r.errorDescription="fieldRequired")}},resetFormErrors(r){const e=this.forms[r].fields??{};for(const s in e)e[s].error=void 0,e[s].errorDescription=void 0},resetFormValues(r){const e=this.forms[r].fields??{};for(const s in e)e[s].value=e[s].originalValue}}});e["Z"]=i},892:function(r,e,s){s.r(e),s.d(e,{default:function(){return d}});var o=s(3396),i=s(4870),t=s(2268),n=s(5117),l=s(4329);const f={key:0};var a=(0,o.aZ)({name:"app-form",props:{formId:null,btnText:{default:"confirm"},btnClass:null,btnContainerClass:null,btnIcon:null,btnIconSize:null,displayErrors:{type:Boolean,default:!0}},emits:["submit","update:displayErrors"],setup(r,{emit:e}){const s=r,a=(0,n.Z)(),{t:u}=(0,l.QT)(),d=(0,o.Fl)((()=>a.getErrorsStatus(s.formId))),c=()=>{a.validateForm(s.formId),d.value&&d.value.length>0?e("submit",!0,d.value):e("submit",!1)};return(e,s)=>{const n=(0,o.up)("app-button");return(0,o.wg)(),(0,o.iD)("div",null,[(0,o._)("div",null,[(0,o.WI)(e.$slots,"default")]),(0,o._)("div",null,[r.displayErrors&&(0,i.SU)(d)?((0,o.wg)(),(0,o.iD)("div",f,[((0,o.wg)(!0),(0,o.iD)(o.HY,null,(0,o.Ko)((0,i.SU)(d),((r,e)=>((0,o.wg)(),(0,o.iD)("p",{key:e,class:"text-error"},(0,t.zw)((0,i.SU)(u)(r.fieldName))+" - "+(0,t.zw)(r.errorDescription),1)))),128))])):(0,o.kq)("",!0),(0,o._)("div",{class:(0,t.C_)(["responsive-flex justify-end gap-4 my-4 mx-2",r.btnContainerClass])},[(0,o.Wm)(n,{icon:r.btnIcon,iconSize:r.btnIconSize,classes:r.btnClass,onClick:c},{default:(0,o.w5)((()=>[(0,o.Uk)((0,t.zw)(r.btnText),1)])),_:1},8,["icon","iconSize","classes"]),(0,o.WI)(e.$slots,"btnZone")],2)])])}}});const u=a;var d=u}}]);
//# sourceMappingURL=892.9541669c.js.map