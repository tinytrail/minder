"use strict";(self.webpackChunkminder_docs=self.webpackChunkminder_docs||[]).push([[9688],{28149:(e,n,r)=>{r.r(n),r.d(n,{assets:()=>l,contentTitle:()=>o,default:()=>d,frontMatter:()=>i,metadata:()=>a,toc:()=>c});var s=r(74848),t=r(28453);const i={title:"Changelog",sidebar_position:30},o="Changelog",a={id:"about/changelog",title:"Changelog",description:"Profile selectors* - Sep 9, 2024",source:"@site/docs/about/changelog.md",sourceDirName:"about",slug:"/about/changelog",permalink:"/about/changelog",draft:!1,unlisted:!1,tags:[],version:"current",sidebarPosition:30,frontMatter:{title:"Changelog",sidebar_position:30},sidebar:"minder",previous:{title:"Minder CLI configuration",permalink:"/ref/cli_configuration"},next:{title:"Roadmap",permalink:"/about/roadmap"}},l={},c=[];function u(e){const n={a:"a",br:"br",code:"code",h1:"h1",header:"header",li:"li",p:"p",pre:"pre",strong:"strong",ul:"ul",...(0,t.R)(),...e.components};return(0,s.jsxs)(s.Fragment,{children:[(0,s.jsx)(n.header,{children:(0,s.jsx)(n.h1,{id:"changelog",children:"Changelog"})}),"\n",(0,s.jsxs)(n.ul,{children:["\n",(0,s.jsxs)(n.li,{children:["\n",(0,s.jsxs)(n.p,{children:[(0,s.jsx)(n.strong,{children:"Profile selectors"})," - Sep 9, 2024",(0,s.jsx)(n.br,{}),"\n","You can now specify which repositories a profile applies to using a Common Expression Language (CEL) grammar."]}),"\n"]}),"\n",(0,s.jsxs)(n.li,{children:["\n",(0,s.jsxs)(n.p,{children:[(0,s.jsx)(n.strong,{children:"Rule evaluation history"})," - Sep 4, 2024",(0,s.jsx)(n.br,{}),"\n","You can now see how your security rules have applied to your repositories, pull requests, and artifacts throughout time, in addition to their current state."]}),"\n"]}),"\n",(0,s.jsxs)(n.li,{children:["\n",(0,s.jsxs)(n.p,{children:[(0,s.jsx)(n.strong,{children:"User management"})," - Aug 5, 2024",(0,s.jsx)(n.br,{}),"\n","Minder organization administrators can now invite additional users to the organization, and can set users permissions."]}),"\n"]}),"\n",(0,s.jsxs)(n.li,{children:["\n",(0,s.jsxs)(n.p,{children:[(0,s.jsx)(n.strong,{children:"Manage all GitHub repositories"})," - Jul 17, 2024",(0,s.jsx)(n.br,{}),"\n","Minder can now (optionally) manage all repositories within a GitHub organization, including new repositories that are created. Administrators can continue to select individual repositories to manage."]}),"\n"]}),"\n",(0,s.jsxs)(n.li,{children:["\n",(0,s.jsxs)(n.p,{children:[(0,s.jsx)(n.strong,{children:"Built-in rules"})," - Apr 6, 2024",(0,s.jsx)(n.br,{}),"\n","Minder now includes all the rules in our ",(0,s.jsx)(n.a,{href:"https://github.com/mindersec/minder-rules-and-profiles/",children:"sample rules repository"})," in your new projects automatically. This means that you do not need to clone that repository or add those rule types to make use of them."]}),"\n",(0,s.jsxs)(n.p,{children:["To use them, prefix the rule name as it exists in the sample rules repository with ",(0,s.jsx)(n.code,{children:"stacklok/"}),". For example:"]}),"\n",(0,s.jsx)(n.pre,{children:(0,s.jsx)(n.code,{className:"language-yaml",children:"---\nversion: v1\ntype: profile\nname: uses-builtin-rules\ncontext:\n  provider: github\nrepository:\n- type: stacklok/secret_scanning\n  def:\n    enabled: true\n"})}),"\n",(0,s.jsxs)(n.p,{children:["You can still define custom rules, or continue to use the rules that exist in the ",(0,s.jsx)(n.a,{href:"https://github.com/mindersec/minder-rules-and-profiles",children:"sample rules repository"}),"."]}),"\n"]}),"\n",(0,s.jsxs)(n.li,{children:["\n",(0,s.jsxs)(n.p,{children:[(0,s.jsx)(n.strong,{children:"User roles"})," - Jan 30, 2024",(0,s.jsx)(n.br,{}),"\n","You can now provide access control for users (eg: administrator, editor, viewer) in your project using ",(0,s.jsx)(n.a,{href:"/user_management/user_roles",children:"built-in roles"}),"."]}),"\n"]}),"\n"]})]})}function d(e={}){const{wrapper:n}={...(0,t.R)(),...e.components};return n?(0,s.jsx)(n,{...e,children:(0,s.jsx)(u,{...e})}):u(e)}},28453:(e,n,r)=>{r.d(n,{R:()=>o,x:()=>a});var s=r(96540);const t={},i=s.createContext(t);function o(e){const n=s.useContext(i);return s.useMemo((function(){return"function"==typeof e?e(n):{...n,...e}}),[n,e])}function a(e){let n;return n=e.disableParentContext?"function"==typeof e.components?e.components(t):e.components||t:o(e.components),s.createElement(i.Provider,{value:n},e.children)}}}]);