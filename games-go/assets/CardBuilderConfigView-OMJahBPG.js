import{d as _,r as u,o as n,c as a,a as t,F as r,w as d,v as p,t as c,e as m}from"./index-o4M0gs7U.js";const x={class:"card-builder"},h={action:""},f=t("p",null,"Todo",-1),b=t("i",null,"in the near future",-1),g={class:"group"},w=t("label",{for:"text-input"},"Text",-1),C={class:"group"},k=t("label",{for:"text-input-2"},"Text",-1),y={class:"result relative"},B={class:"res-inp-1 absolute top-2 left-2 text-2xl"},T={class:"res-inp-2 absolute top-2 right-2 text-2xl"},V=_({__name:"CardBuilder",setup(v){const s=u(""),l=u(""),i=u(!1);return(U,e)=>(n(),a("div",x,[t("form",h,[i.value?(n(),a(r,{key:0},[f,t("button",{type:"button",onClick:e[0]||(e[0]=o=>i.value=!1)},"Switch to input")],64)):(n(),a(r,{key:1},[t("button",{type:"button",onClick:e[1]||(e[1]=o=>i.value=!0),disabled:""},"Switch to Image Upload"),b,t("div",g,[w,d(t("input",{class:"mt-1 block w-full px-3 py-2 bg-white border border-slate-300 rounded-md text-sm shadow-sm placeholder-slate-400",id:"text-input","onUpdate:modelValue":e[2]||(e[2]=o=>s.value=o)},null,512),[[p,s.value]])]),t("div",C,[k,d(t("input",{class:"",id:"text-input-2","onUpdate:modelValue":e[3]||(e[3]=o=>l.value=o)},null,512),[[p,l.value]])])],64))]),t("div",y,[t("div",B,c(s.value),1),t("div",T,c(l.value),1)])]))}}),I=t("h1",null,"Create Your Card here",-1),S=_({__name:"CardBuilderConfigView",setup(v){return(s,l)=>(n(),a("main",null,[I,m(V)]))}});export{S as default};
