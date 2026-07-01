---
type: pull_request
number: 1219
title: "[Snyk] Security upgrade @redhat-cloud-services/host-inventory-client from 1.5.2 to 2.0.1"
state: merged
author: jiridostal
created: 2024-11-08T04:18:40Z
updated: 2025-02-11T09:26:19Z
closed: 2025-02-11T09:26:15Z
merged: 2025-02-11T09:26:15Z
base_branch: master
head_branch: snyk-fix-02ce97e9d34c4b6236ea357a4c504530
labels: ["dependencies"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1219
---

# Pull Request #1219: [Snyk] Security upgrade @redhat-cloud-services/host-inventory-client from 1.5.2 to 2.0.1

**Author**: @jiridostal
**Created**: November 08, 2024 at 04:18 AM UTC
**Status**: Merged
**Labels**: `dependencies`
**Base**: `master` ← **Head**: `snyk-fix-02ce97e9d34c4b6236ea357a4c504530`

## Description

![snyk-top-banner](https://github.com/andygongea/OWASP-Benchmark/assets/818805/c518c423-16fe-447e-b67f-ad5a49b5d123)

### Snyk has created this PR to fix 2 vulnerabilities in the npm dependencies of this project.

#### Snyk changed the following file(s):

- `package.json`
- `package-lock.json`




#### Vulnerabilities that will be fixed with an upgrade:

|  | Issue | Score | 
:-------------------------:|:-------------------------|:-------------------------
![high severity](https://res.cloudinary.com/snyk/image/upload/w_20,h_20/v1561977819/icon/h.png 'high severity') | Cross-site Request Forgery (CSRF) <br/>[SNYK-JS-AXIOS-6032459](https://snyk.io/vuln/SNYK-JS-AXIOS-6032459) | &nbsp;&nbsp;**676**&nbsp;&nbsp; 
![medium severity](https://res.cloudinary.com/snyk/image/upload/w_20,h_20/v1561977819/icon/m.png 'medium severity') | Regular Expression Denial of Service (ReDoS) <br/>[SNYK-JS-AXIOS-6124857](https://snyk.io/vuln/SNYK-JS-AXIOS-6124857) | &nbsp;&nbsp;**586**&nbsp;&nbsp; 




---

> [!IMPORTANT]
>
> - Check the changes in this PR to ensure they won't cause issues with your project.
> - Max score is 1000. Note that the real score may have changed since the PR was raised.
> - This PR was automatically created by Snyk using the credentials of a real user.

---

**Note:** _You are seeing this because you or someone else with access to this repository has authorized Snyk to open fix PRs._

For more information: <img src="https://api.segment.io/v1/pixel/track?data=eyJ3cml0ZUtleSI6InJyWmxZcEdHY2RyTHZsb0lYd0dUcVg4WkFRTnNCOUEwIiwiYW5vbnltb3VzSWQiOiIxNzM4YjYwNi1jZjI1LTQ2ZjUtYWIxYi04ZTExMWM0NWE4NzMiLCJldmVudCI6IlBSIHZpZXdlZCIsInByb3BlcnRpZXMiOnsicHJJZCI6IjE3MzhiNjA2LWNmMjUtNDZmNS1hYjFiLThlMTExYzQ1YTg3MyJ9fQ==" width="0" height="0"/>
🧐 [View latest project report](https://app.snyk.io/org/hybrid-cloud-experience-red-hat-insights/project/abdec9e9-4fb8-4f10-8528-02a6b74b6b33?utm_source&#x3D;github&amp;utm_medium&#x3D;referral&amp;page&#x3D;fix-pr)
📜 [Customise PR templates](https://docs.snyk.io/scan-using-snyk/pull-requests/snyk-fix-pull-or-merge-requests/customize-pr-templates)
🛠 [Adjust project settings](https://app.snyk.io/org/hybrid-cloud-experience-red-hat-insights/project/abdec9e9-4fb8-4f10-8528-02a6b74b6b33?utm_source&#x3D;github&amp;utm_medium&#x3D;referral&amp;page&#x3D;fix-pr/settings)
📚 [Read about Snyk's upgrade logic](https://support.snyk.io/hc/en-us/articles/360003891078-Snyk-patches-to-fix-vulnerabilities)

---

**Learn how to fix vulnerabilities with free interactive lessons:**

🦉 [Cross-site Request Forgery (CSRF)](https://learn.snyk.io/lesson/csrf-attack/?loc&#x3D;fix-pr)
🦉 [Regular Expression Denial of Service (ReDoS)](https://learn.snyk.io/lesson/redos/?loc&#x3D;fix-pr)

[//]: # 'snyk:metadata:{"customTemplate":{"variablesUsed":[],"fieldsUsed":[]},"dependencies":[{"name":"@redhat-cloud-services/host-inventory-client","from":"1.5.2","to":"2.0.1"}],"env":"prod","issuesToFix":[{"exploit_maturity":"Proof of Concept","id":"SNYK-JS-AXIOS-6032459","priority_score":676,"priority_score_factors":[{"type":"exploit","label":"Proof of Concept","score":107},{"type":"fixability","label":true,"score":214},{"type":"cvssScore","label":"7.1","score":355},{"type":"scoreVersion","label":"v1","score":1}],"severity":"high","title":"Cross-site Request Forgery (CSRF)"},{"exploit_maturity":"Proof of Concept","id":"SNYK-JS-AXIOS-6124857","priority_score":586,"priority_score_factors":[{"type":"exploit","label":"Proof of Concept","score":107},{"type":"fixability","label":true,"score":214},{"type":"cvssScore","label":"5.3","score":265},{"type":"scoreVersion","label":"v1","score":1}],"severity":"medium","title":"Regular Expression Denial of Service (ReDoS)"}],"prId":"1738b606-cf25-46f5-ab1b-8e111c45a873","prPublicId":"1738b606-cf25-46f5-ab1b-8e111c45a873","packageManager":"npm","priorityScoreList":[676,586],"projectPublicId":"abdec9e9-4fb8-4f10-8528-02a6b74b6b33","projectUrl":"https://app.snyk.io/org/hybrid-cloud-experience-red-hat-insights/project/abdec9e9-4fb8-4f10-8528-02a6b74b6b33?utm_source=github&utm_medium=referral&page=fix-pr","prType":"fix","templateFieldSources":{"branchName":"default","commitMessage":"default","description":"default","title":"default"},"templateVariants":["updated-fix-title","priorityScore"],"type":"auto","upgrade":["SNYK-JS-AXIOS-6032459","SNYK-JS-AXIOS-6124857"],"vulns":["SNYK-JS-AXIOS-6032459","SNYK-JS-AXIOS-6124857"],"patch":[],"isBreakingChange":true,"remediationStrategy":"vuln"}'


---

## Discussion

### Comment by @gkarat on February 07, 2025 at 10:01 AM UTC

/retest

### Comment by @gkarat on February 07, 2025 at 10:03 AM UTC

@RedHatInsights/team-interact, this must be ready. I have manually tested the two affected endpoints and compared both of the responses before/after these changes, and the responses are equal. I assume no further changes required with this upgrade.

### Comment by @codecov-commenter on February 07, 2025 at 10:33 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1219?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: Patch coverage is `66.66667%` with `1 line` in your changes missing coverage. Please review.
> Project coverage is 63.59%. Comparing base [(`bf6464b`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/bf6464b42a611f168757d7b5e8ec3bae4816b169?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`7fdf2dd`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/7fdf2dd555f491950d90efcf8768e11934f17227?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1219?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [src/Utilities/api.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1219?src=pr&el=tree&filepath=src%2FUtilities%2Fapi.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9hcGkuanM=) | 66.66% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1219?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1219      +/-   ##
==========================================
- Coverage   63.59%   63.59%   -0.01%     
==========================================
  Files         124      124              
  Lines        3269     3266       -3     
  Branches      860      860              
==========================================
- Hits         2079     2077       -2     
+ Misses       1190     1189       -1     
```

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1219?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @bastilian - Approved on February 10, 2025 at 11:41 AM UTC

LGTM! OS filter appears to still work as expected! Thank you @jiridostal & @gkarat!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1219*
