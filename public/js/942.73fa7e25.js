"use strict";(self["webpackChunk_3d_renderer"]=self["webpackChunk_3d_renderer"]||[]).push([[942],{5117:function(e,l,o){var r=o(7139);const s=(0,r.Q_)("forms",{state:()=>({forms:{}}),getters:{getFormFields:e=>l=>e.forms[l]?.fields,getFormRawValues:e=>l=>{if(!e.forms[l])return;const o={};for(const r in e.forms[l].fields)o[r]=e.forms[l].fields[r].value;return o},getErrorsStatus:e=>l=>{const o=[];if(e.forms[l]){for(const r in e.forms[l].fields){const s=e.forms[l].fields[r];1==s.error&&o.push({fieldName:r,fieldCurrentValue:s.value,fieldOriginalValue:s.originalValue,errorDescription:s.errorDescription??"UNKNOWN"})}return o}},getField:e=>(l,o)=>e.forms[l]?.fields[o],getFieldValue:e=>(l,o)=>e.forms[l]?.fields[o]?.value,getFieldErrorStatus:e=>(l,o)=>e.forms[l]?.fields[o]?.error,getFieldErrorMessage:e=>(l,o)=>e.forms[l]?.fields[o]?.errorDescription},actions:{setFieldValue(e,l,o,r,s,t){this.forms[e]||(this.forms[e]={fields:{}}),this.forms[e].fields[l]||(this.forms[e].fields[l]={originalValue:o});const i=this.forms[e].fields[l];"undefined"!=typeof s&&(i.error=s,i.errorDescription=t),i.value=o,i.required=r},setFieldError(e,l,o,r){const s=this.forms[e].fields[l];s&&(s.error=o,s.errorDescription=r)},validateForm(e){const l=this.forms[e].fields??{};for(const o in l){const e=l[o];e.required&&!e.value&&(e.error=!0,e.errorDescription="fieldRequired")}},resetFormErrors(e){const l=this.forms[e].fields??{};for(const o in l)l[o].error=void 0,l[o].errorDescription=void 0},resetFormValues(e){const l=this.forms[e].fields??{};for(const o in l)l[o].value=l[o].originalValue}}});l["Z"]=s},1041:function(e,l,o){var r=o(1158),s=o(7139);const t=(0,s.Q_)("references",{state:()=>({references:{}}),getters:{},actions:{async fetchReference(e){if(this.references[e])return this.references[e];const l=await r.Z.references.getReference(e);return l&&(this.references[e]=l),l}}});l["Z"]=t},3942:function(e,l,o){o.r(l),o.d(l,{default:function(){return m}});var r=o(3396),s=o(2268),t=o(4870),i=o(5117),n=o(1041),a=o(4329);const u={class:"flex-centered gap-2"},f=["for"],c={key:0};var d=(0,r.aZ)({name:"app-input-select",props:{valueField:null,formId:null,name:null,refTable:null,options:null,modelValue:null,labelField:{default:"text1"},multiple:{type:Boolean},clearable:{type:Boolean},selectable:null,filter:null,autoComplete:null,autoScroll:{type:Boolean},sameLine:{type:Boolean},icon:null,iconSize:null,labelClass:null,required:{type:Boolean},responsive:{type:Boolean},containerClass:null,loadData:{type:Boolean}},emits:["update:modelValue"],setup(e,{emit:l}){const o=e,{t:d}=(0,a.QT)(),p=(0,n.Z)(),m=(0,i.Z)(),v=(0,t.Fl)((()=>o.refTable?p.references[o.refTable]?.values:o.options)),b=(0,t.iH)(o.modelValue),F=(0,t.Fl)((()=>o.formId?m.getFieldValue(o.formId,o.name):void 0)),g=e=>{const l=e??{};return o.valueField?l[o.valueField]:l["code"]??l};(0,r.YP)((()=>o.modelValue),(e=>b.value=e)),(0,r.YP)(b,(e=>{o.formId&&m.setFieldValue(o.formId,o.name,g(e),o.required),l("update:modelValue",e)}));const h=(0,t.Fl)((()=>{const e=[""];return o.sameLine?o.responsive?e.push("responsive-flex-centered gap-2"):e.push("flex-centered gap-2"):e.push("flex-column gap-2"),e}));return(0,r.wF)((()=>{o.refTable&&p.fetchReference(o.refTable),o.formId&&(o.loadData&&F.value?b.value=o.options?.find((e=>b.value===g(e))):m.setFieldValue(o.formId,o.name,g(b.value),o.required))})),(l,o)=>{const i=(0,r.up)("font-awesome-icon"),n=(0,r.up)("v-select");return(0,r.wg)(),(0,r.iD)("div",{class:(0,s.C_)(["w-full",(0,t.SU)(h),e.containerClass])},[(0,r._)("div",u,[e.icon?((0,r.wg)(),(0,r.j4)(i,{key:0,icon:e.icon,size:e.iconSize},null,8,["icon","size"])):(0,r.kq)("",!0),(0,r._)("label",{for:e.name,class:(0,s.C_)(["label",e.labelClass])},[(0,r.Uk)((0,s.zw)((0,t.SU)(d)(e.name))+" ",1),e.required?((0,r.wg)(),(0,r.iD)("span",c,"*")):(0,r.kq)("",!0)],10,f)]),(0,r.Wm)(n,{modelValue:b.value,"onUpdate:modelValue":o[0]||(o[0]=e=>b.value=e),options:(0,t.SU)(v),label:e.labelField,multiple:e.multiple,clearable:e.clearable,selectable:e.selectable,filter:e.filter,"auto-complete":e.autoComplete,"auto-scroll":e.autoScroll},{"list-header":(0,r.w5)((()=>[(0,r.WI)(l.$slots,"list-footer")])),"list-footer":(0,r.w5)((()=>[(0,r.WI)(l.$slots,"list-footer")])),"no-options":(0,r.w5)((()=>[(0,r.WI)(l.$slots,"no-options")])),option:(0,r.w5)((e=>[(0,r.WI)(l.$slots,"option",{option:e})])),"selected-option":(0,r.w5)((e=>[(0,r.WI)(l.$slots,"selected-option",{selectedOption:e})])),"selected-option-container":(0,r.w5)((({option:e,deselect:o,multiple:s,disabled:t})=>[(0,r.WI)(l.$slots,"selected-option-container",{selectedOption:{option:e,deselect:o,multiple:s,disabled:t}})])),_:3},8,["modelValue","options","label","multiple","clearable","selectable","filter","auto-complete","auto-scroll"])],2)}}});const p=d;var m=p}}]);
//# sourceMappingURL=942.73fa7e25.js.map