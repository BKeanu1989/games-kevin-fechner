import{d as _,o as c,c as n,F as x,g as y,t as f,_ as k,r as i,h as E,i as b,a as d,e as p,f as B}from"./index-o4M0gs7U.js";const $={class:"first-deck"},D=_({__name:"FirstDeck",props:{cards:{}},setup(r){const s=r;return(o,u)=>(c(),n("section",$,[(c(!0),n(x,null,y(s.cards,t=>(c(),n("div",{class:"first-deck--single",key:t}," | "+f(t)+" | ",1))),128))]))}}),L=k(D,[["__scopeId","data-v-0ea41d2e"]]),M=["viewBox"],N=b('<path d="M 10 60 C 20 80, 40 80, 50 60" stroke="black" fill="transparent"></path><circle cx="20" cy="80" r="2" fill="gold"></circle><circle cx="40" cy="80" r="2" fill="gold"></circle><circle cx="50" cy="60" r="2" fill="green"></circle><circle cx="10" cy="60" r="2" fill="green"></circle><rect id="draggableRect" x="20" y="20" width="40" height="80" stroke="black" fill="transparent" stroke-width="2"></rect><text x="30" y="50" id="draggableText"> test </text>',7),S=[N],F=_({__name:"SimpleComposableCard",setup(r){const s=i(100),o=i(300),u=i(null);return E(()=>{const t=document.getElementById("svgCanvas"),v=document.getElementById("draggableRect");h(v);const w=document.getElementById("draggableText");h(w);function h(g){let e=null,m={x:0,y:0};g.addEventListener("mousedown",a=>{e=a.target,m.x=a.clientX-e.getBoundingClientRect().x,m.y=a.clientY-e.getBoundingClientRect().y}),t.addEventListener("mousemove",a=>{if(e){a.clientX-m.x,a.clientY-m.y;const l=C(a);console.log("🚀 ~ svg.addEventListener ~ coords:",l),(e.tagName==="rect"||e.tagName==="text")&&(e.setAttribute("x",l.x),e.setAttribute("y",l.y))}}),t.addEventListener("mouseup",()=>{e=null}),t.addEventListener("mouseleave",()=>{e=null})}function C(g){var e=t.getScreenCTM();return console.log("🚀 ~ getMousePosition ~ svg:",t),{x:(g.clientX-e.e)/e.a,y:(g.clientY-e.f)/e.d}}}),(t,v)=>(c(),n("svg",{ref_key:"svgElement",ref:u,id:"svgCanvas",viewBox:"0 0 "+s.value+" "+o.value,onload:"",xmlns:"http://www.w3.org/2000/svg","xmlns:xlink":"http://www.w3.org/1999/xlink"},S,8,M))}}),I={class:"second-deck"},T=_({__name:"SecondDeck",props:{cards:{}},setup(r){const s=r;return(o,u)=>(c(),n(x,null,[d("section",I,[(c(!0),n(x,null,y(s.cards,t=>(c(),n("div",{class:"second-deck--single",key:t.name}," | "+f(t)+" | ",1))),128))]),d("section",null,[p(F)])],64))}}),V=k(T,[["__scopeId","data-v-f5917ac4"]]),A=d("h1",{class:"container-v2"},"Card Game",-1),R={class:"cards"},G=d("h1",null,"First deck",-1),X=_({__name:"CardGameContainerV2",setup(r){const s=i(["a","2","3"]),o=i([{name:"3",type:""},{name:"2",type:""},{name:"a",type:""}]);return(u,t)=>(c(),n("div",null,[A,d("div",R,f(s.value),1),B(" --- "),G,p(V,{cards:o.value},null,8,["cards"]),p(L,{cards:s.value},null,8,["cards"])]))}}),Y={class:"page-view"},j=_({__name:"CardGame",setup(r){return(s,o)=>(c(),n("div",Y,[p(X)]))}});export{j as default};