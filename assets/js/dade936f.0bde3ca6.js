"use strict";(self.webpackChunkminder_docs=self.webpackChunkminder_docs||[]).push([[2019],{90992:(e,n,i)=>{i.r(n),i.d(n,{assets:()=>d,contentTitle:()=>o,default:()=>h,frontMatter:()=>t,metadata:()=>r,toc:()=>a});var s=i(74848),l=i(28453);const t={title:"GitHub Actions",sidebar_position:70},o="GitHub Actions Configuration Rules",r={id:"ref/rules/github_actions",title:"GitHub Actions",description:"There are several rule types that can be used to configure GitHub Actions.",source:"@site/docs/ref/rules/github_actions.md",sourceDirName:"ref/rules",slug:"/ref/rules/github_actions",permalink:"/ref/rules/github_actions",draft:!1,unlisted:!1,tags:[],version:"current",sidebarPosition:70,frontMatter:{title:"GitHub Actions",sidebar_position:70},sidebar:"minder",previous:{title:"Known vulnerabilities",permalink:"/ref/rules/pr_vulnerability_check"},next:{title:"Presence of a license file",permalink:"/ref/rules/license"}},d={},a=[{value:"<code>github_actions_allowed</code> - Which actions are allowed to be used",id:"github_actions_allowed---which-actions-are-allowed-to-be-used",level:2},{value:"Entity",id:"entity",level:3},{value:"Type",id:"type",level:3},{value:"Rule parameters",id:"rule-parameters",level:3},{value:"Rule definition options",id:"rule-definition-options",level:3},{value:"<code>allowed_selected_actions</code> - Verifies that only allowed actions are used",id:"allowed_selected_actions---verifies-that-only-allowed-actions-are-used",level:2},{value:"Entity",id:"entity-1",level:3},{value:"Type",id:"type-1",level:3},{value:"Rule parameters",id:"rule-parameters-1",level:3},{value:"Rule definition options",id:"rule-definition-options-1",level:3},{value:"<code>default_workflow_permissions</code> - Sets the default permissions granted to the <code>GITHUB_TOKEN</code> when running workflows",id:"default_workflow_permissions---sets-the-default-permissions-granted-to-the-github_token-when-running-workflows",level:2},{value:"Entity",id:"entity-2",level:3},{value:"Type",id:"type-2",level:3},{value:"Rule parameters",id:"rule-parameters-2",level:3},{value:"Rule definition options",id:"rule-definition-options-2",level:3},{value:"<code>actions_check_pinned_tags</code> - Verifies that any actions use pinned tags",id:"actions_check_pinned_tags---verifies-that-any-actions-use-pinned-tags",level:2},{value:"Entity",id:"entity-3",level:3},{value:"Type",id:"type-3",level:3},{value:"Rule parameters",id:"rule-parameters-3",level:3},{value:"Rule definition options",id:"rule-definition-options-3",level:3}];function c(e){const n={code:"code",h1:"h1",h2:"h2",h3:"h3",li:"li",p:"p",ul:"ul",...(0,l.R)(),...e.components};return(0,s.jsxs)(s.Fragment,{children:[(0,s.jsx)(n.h1,{id:"github-actions-configuration-rules",children:"GitHub Actions Configuration Rules"}),"\n",(0,s.jsx)(n.p,{children:"There are several rule types that can be used to configure GitHub Actions."}),"\n",(0,s.jsxs)(n.h2,{id:"github_actions_allowed---which-actions-are-allowed-to-be-used",children:[(0,s.jsx)(n.code,{children:"github_actions_allowed"})," - Which actions are allowed to be used"]}),"\n",(0,s.jsxs)(n.p,{children:["This rule allows you to limit the actions that are allowed to run for a repository.\nIt is recommended to use the ",(0,s.jsx)(n.code,{children:"selected"})," option for allowed actions, and then\nselect the actions that are allowed to run."]}),"\n",(0,s.jsx)(n.h3,{id:"entity",children:"Entity"}),"\n",(0,s.jsxs)(n.ul,{children:["\n",(0,s.jsx)(n.li,{children:(0,s.jsx)(n.code,{children:"repository"})}),"\n"]}),"\n",(0,s.jsx)(n.h3,{id:"type",children:"Type"}),"\n",(0,s.jsxs)(n.ul,{children:["\n",(0,s.jsx)(n.li,{children:(0,s.jsx)(n.code,{children:"github_actions_allowed"})}),"\n"]}),"\n",(0,s.jsx)(n.h3,{id:"rule-parameters",children:"Rule parameters"}),"\n",(0,s.jsxs)(n.ul,{children:["\n",(0,s.jsx)(n.li,{children:"None"}),"\n"]}),"\n",(0,s.jsx)(n.h3,{id:"rule-definition-options",children:"Rule definition options"}),"\n",(0,s.jsxs)(n.p,{children:["The ",(0,s.jsx)(n.code,{children:"github_actions_allowed"})," rule supports the following options:"]}),"\n",(0,s.jsxs)(n.ul,{children:["\n",(0,s.jsxs)(n.li,{children:[(0,s.jsx)(n.code,{children:"allowed_actions (enum)"})," - Which actions are allowed to be used","\n",(0,s.jsxs)(n.ul,{children:["\n",(0,s.jsxs)(n.li,{children:[(0,s.jsx)(n.code,{children:"all"})," - Any action or reusable workflow can be used, regardless of who authored it or where it is defined."]}),"\n",(0,s.jsxs)(n.li,{children:[(0,s.jsx)(n.code,{children:"local_only"})," - Only actions and reusable workflows that are defined in the repository or organization can be used."]}),"\n",(0,s.jsxs)(n.li,{children:[(0,s.jsx)(n.code,{children:"selected"})," - Only the actions and reusable workflows that are explicitly listed are allowed. Use the ",(0,s.jsx)(n.code,{children:"allowed_selected_actions"})," ",(0,s.jsx)(n.code,{children:"rule_type"})," to set the list of allowed actions."]}),"\n"]}),"\n"]}),"\n"]}),"\n",(0,s.jsxs)(n.h2,{id:"allowed_selected_actions---verifies-that-only-allowed-actions-are-used",children:[(0,s.jsx)(n.code,{children:"allowed_selected_actions"})," - Verifies that only allowed actions are used"]}),"\n",(0,s.jsxs)(n.p,{children:["To use this rule, the repository profile for ",(0,s.jsx)(n.code,{children:"github_actions_allowed"})," must\nbe configured to ",(0,s.jsx)(n.code,{children:"selected"}),"."]}),"\n",(0,s.jsx)(n.h3,{id:"entity-1",children:"Entity"}),"\n",(0,s.jsxs)(n.ul,{children:["\n",(0,s.jsx)(n.li,{children:(0,s.jsx)(n.code,{children:"repository"})}),"\n"]}),"\n",(0,s.jsx)(n.h3,{id:"type-1",children:"Type"}),"\n",(0,s.jsxs)(n.ul,{children:["\n",(0,s.jsx)(n.li,{children:(0,s.jsx)(n.code,{children:"allowed_selected_actions"})}),"\n"]}),"\n",(0,s.jsx)(n.h3,{id:"rule-parameters-1",children:"Rule parameters"}),"\n",(0,s.jsxs)(n.ul,{children:["\n",(0,s.jsx)(n.li,{children:"None"}),"\n"]}),"\n",(0,s.jsx)(n.h3,{id:"rule-definition-options-1",children:"Rule definition options"}),"\n",(0,s.jsxs)(n.p,{children:["The ",(0,s.jsx)(n.code,{children:"allowed_selected_actions"})," rule supports the following options:"]}),"\n",(0,s.jsxs)(n.ul,{children:["\n",(0,s.jsxs)(n.li,{children:[(0,s.jsx)(n.code,{children:"github_owner_allowed (boolean)"})," - Whether GitHub-owned actions are allowed. For example, this includes the actions in the ",(0,s.jsx)(n.code,{children:"actions"})," organization."]}),"\n",(0,s.jsxs)(n.li,{children:[(0,s.jsx)(n.code,{children:"verified_allowed (boolean)"})," - Whether actions that are verified by GitHub are allowed."]}),"\n",(0,s.jsxs)(n.li,{children:[(0,s.jsx)(n.code,{children:"patterns_allowed (boolean)"})," - Specifies a list of string-matching patterns to allow specific action(s) and reusable workflow(s). Wildcards, tags, and SHAs are allowed."]}),"\n"]}),"\n",(0,s.jsxs)(n.h2,{id:"default_workflow_permissions---sets-the-default-permissions-granted-to-the-github_token-when-running-workflows",children:[(0,s.jsx)(n.code,{children:"default_workflow_permissions"})," - Sets the default permissions granted to the ",(0,s.jsx)(n.code,{children:"GITHUB_TOKEN"})," when running workflows"]}),"\n",(0,s.jsx)(n.p,{children:"Verifies the default workflow permissions granted to the GITHUB_TOKEN\nwhen running workflows in a repository, as well as if GitHub Actions\ncan submit approving pull request reviews."}),"\n",(0,s.jsx)(n.h3,{id:"entity-2",children:"Entity"}),"\n",(0,s.jsxs)(n.ul,{children:["\n",(0,s.jsx)(n.li,{children:(0,s.jsx)(n.code,{children:"repository"})}),"\n"]}),"\n",(0,s.jsx)(n.h3,{id:"type-2",children:"Type"}),"\n",(0,s.jsxs)(n.ul,{children:["\n",(0,s.jsx)(n.li,{children:(0,s.jsx)(n.code,{children:"default_workflow_permissions"})}),"\n"]}),"\n",(0,s.jsx)(n.h3,{id:"rule-parameters-2",children:"Rule parameters"}),"\n",(0,s.jsxs)(n.ul,{children:["\n",(0,s.jsx)(n.li,{children:"None"}),"\n"]}),"\n",(0,s.jsx)(n.h3,{id:"rule-definition-options-2",children:"Rule definition options"}),"\n",(0,s.jsxs)(n.p,{children:["The ",(0,s.jsx)(n.code,{children:"default_workflow_permissions"})," rule supports the following options:"]}),"\n",(0,s.jsxs)(n.ul,{children:["\n",(0,s.jsxs)(n.li,{children:[(0,s.jsx)(n.code,{children:"default_workflow_permissions (boolean)"})," - Whether GitHub-owned actions are allowed. For example, this includes the actions in the ",(0,s.jsx)(n.code,{children:"actions"})," organization."]}),"\n",(0,s.jsxs)(n.li,{children:[(0,s.jsx)(n.code,{children:"can_approve_pull_request_reviews (boolean)"})," - Whether the ",(0,s.jsx)(n.code,{children:"GITHUB_TOKEN"})," can approve pull request reviews."]}),"\n"]}),"\n",(0,s.jsxs)(n.h2,{id:"actions_check_pinned_tags---verifies-that-any-actions-use-pinned-tags",children:[(0,s.jsx)(n.code,{children:"actions_check_pinned_tags"})," - Verifies that any actions use pinned tags"]}),"\n",(0,s.jsx)(n.p,{children:"Verifies that actions use pinned tags as opposed to floating tags."}),"\n",(0,s.jsx)(n.h3,{id:"entity-3",children:"Entity"}),"\n",(0,s.jsxs)(n.ul,{children:["\n",(0,s.jsx)(n.li,{children:(0,s.jsx)(n.code,{children:"repository"})}),"\n"]}),"\n",(0,s.jsx)(n.h3,{id:"type-3",children:"Type"}),"\n",(0,s.jsxs)(n.ul,{children:["\n",(0,s.jsx)(n.li,{children:(0,s.jsx)(n.code,{children:"actions_check_pinned_tags"})}),"\n"]}),"\n",(0,s.jsx)(n.h3,{id:"rule-parameters-3",children:"Rule parameters"}),"\n",(0,s.jsxs)(n.ul,{children:["\n",(0,s.jsx)(n.li,{children:"None"}),"\n"]}),"\n",(0,s.jsx)(n.h3,{id:"rule-definition-options-3",children:"Rule definition options"}),"\n",(0,s.jsxs)(n.ul,{children:["\n",(0,s.jsx)(n.li,{children:"None"}),"\n"]})]})}function h(e={}){const{wrapper:n}={...(0,l.R)(),...e.components};return n?(0,s.jsx)(n,{...e,children:(0,s.jsx)(c,{...e})}):c(e)}},28453:(e,n,i)=>{i.d(n,{R:()=>o,x:()=>r});var s=i(96540);const l={},t=s.createContext(l);function o(e){const n=s.useContext(t);return s.useMemo((function(){return"function"==typeof e?e(n):{...n,...e}}),[n,e])}function r(e){let n;return n=e.disableParentContext?"function"==typeof e.components?e.components(l):e.components||l:o(e.components),s.createElement(t.Provider,{value:n},e.children)}}}]);