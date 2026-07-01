---
type: pull_request
number: 1499
title: "[Snyk] Security upgrade grafana/grafana from 11.2.1 to 11.2.2"
state: closed
author: chambridge
created: 2024-10-02T08:02:53Z
updated: 2024-10-02T13:45:57Z
closed: 2024-10-02T13:45:56Z
base_branch: master
head_branch: snyk-fix-a1fe9404b307609e49b23e9d5fa1c48c
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1499
---

# Pull Request #1499: [Snyk] Security upgrade grafana/grafana from 11.2.1 to 11.2.2

**Author**: @chambridge
**Created**: October 02, 2024 at 08:02 AM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `snyk-fix-a1fe9404b307609e49b23e9d5fa1c48c`

## Description

![snyk-top-banner](https://github.com/andygongea/OWASP-Benchmark/assets/818805/c518c423-16fe-447e-b67f-ad5a49b5d123)

### Snyk has created this PR to fix 4 vulnerabilities in the dockerfile dependencies of this project.

Keeping your Docker base image up-to-date means you’ll benefit from security fixes in the latest version of your chosen image.

#### Snyk changed the following file(s):

- `dev/grafana/Dockerfile`

We recommend upgrading to `grafana/grafana:11.2.2`, as this image has only **0** known vulnerabilities. To do this, merge this pull request, then verify your application still works as expected.



#### Vulnerabilities that will be fixed with an upgrade:

|  | Issue | Score | 
:-------------------------:|:-------------------------|:-------------------------
![medium severity](https://res.cloudinary.com/snyk/image/upload/w_20,h_20/v1561977819/icon/m.png 'medium severity') | Out-of-bounds Write <br/>[SNYK-ALPINE319-BUSYBOX-6913413](https://snyk.io/vuln/SNYK-ALPINE319-BUSYBOX-6913413) | &nbsp;&nbsp;**514**&nbsp;&nbsp; 
![medium severity](https://res.cloudinary.com/snyk/image/upload/w_20,h_20/v1561977819/icon/m.png 'medium severity') | Out-of-bounds Write <br/>[SNYK-ALPINE319-BUSYBOX-6913413](https://snyk.io/vuln/SNYK-ALPINE319-BUSYBOX-6913413) | &nbsp;&nbsp;**514**&nbsp;&nbsp; 
![medium severity](https://res.cloudinary.com/snyk/image/upload/w_20,h_20/v1561977819/icon/m.png 'medium severity') | Use After Free <br/>[SNYK-ALPINE319-BUSYBOX-6928845](https://snyk.io/vuln/SNYK-ALPINE319-BUSYBOX-6928845) | &nbsp;&nbsp;**514**&nbsp;&nbsp; 
![medium severity](https://res.cloudinary.com/snyk/image/upload/w_20,h_20/v1561977819/icon/m.png 'medium severity') | Use After Free <br/>[SNYK-ALPINE319-BUSYBOX-6928846](https://snyk.io/vuln/SNYK-ALPINE319-BUSYBOX-6928846) | &nbsp;&nbsp;**514**&nbsp;&nbsp; 
![medium severity](https://res.cloudinary.com/snyk/image/upload/w_20,h_20/v1561977819/icon/m.png 'medium severity') | Use After Free <br/>[SNYK-ALPINE319-BUSYBOX-6928847](https://snyk.io/vuln/SNYK-ALPINE319-BUSYBOX-6928847) | &nbsp;&nbsp;**514**&nbsp;&nbsp; 



---

> [!IMPORTANT]
>
> - Check the changes in this PR to ensure they won't cause issues with your project.
> - Max score is 1000. Note that the real score may have changed since the PR was raised.
> - This PR was automatically created by Snyk using the credentials of a real user.

---

**Note:** _You are seeing this because you or someone else with access to this repository has authorized Snyk to open fix PRs._

For more information: <img src="https://api.segment.io/v1/pixel/track?data=eyJ3cml0ZUtleSI6InJyWmxZcEdHY2RyTHZsb0lYd0dUcVg4WkFRTnNCOUEwIiwiYW5vbnltb3VzSWQiOiIwYTYyNGI1NS1kODhkLTQ1OWYtOGQ0My0wZjdjNzM5ZjBiZmEiLCJldmVudCI6IlBSIHZpZXdlZCIsInByb3BlcnRpZXMiOnsicHJJZCI6IjBhNjI0YjU1LWQ4OGQtNDU5Zi04ZDQzLTBmN2M3MzlmMGJmYSJ9fQ==" width="0" height="0"/>
🧐 [View latest project report](https://app.snyk.io/org/hybrid-cloud-experience-red-hat-insights/project/f4635cb3-382d-48ac-93bb-f4b105d4dcf4?utm_source&#x3D;github&amp;utm_medium&#x3D;referral&amp;page&#x3D;fix-pr)
📜 [Customise PR templates](https://docs.snyk.io/scan-using-snyk/pull-requests/snyk-fix-pull-or-merge-requests/customize-pr-templates)
🛠 [Adjust project settings](https://app.snyk.io/org/hybrid-cloud-experience-red-hat-insights/project/f4635cb3-382d-48ac-93bb-f4b105d4dcf4?utm_source&#x3D;github&amp;utm_medium&#x3D;referral&amp;page&#x3D;fix-pr/settings)
📚 [Read about Snyk's upgrade logic](https://support.snyk.io/hc/en-us/articles/360003891078-Snyk-patches-to-fix-vulnerabilities)

---

**Learn how to fix vulnerabilities with free interactive lessons:**

🦉 [Use After Free](https://learn.snyk.io/lesson/use-after-free/?loc&#x3D;fix-pr)

[//]: # 'snyk:metadata:{"customTemplate":{"variablesUsed":[],"fieldsUsed":[]},"dependencies":[{"name":"grafana/grafana","from":"11.2.1","to":"11.2.2"}],"env":"prod","issuesToFix":[{"exploit_maturity":"No Known Exploit","id":"SNYK-ALPINE319-BUSYBOX-6913413","priority_score":514,"priority_score_factors":[{"type":"fixability","label":true,"score":214},{"type":"severity","label":"medium","score":300},{"type":"scoreVersion","label":"v1","score":1}],"severity":"medium","title":"Out-of-bounds Write"},{"exploit_maturity":"No Known Exploit","id":"SNYK-ALPINE319-BUSYBOX-6928845","priority_score":514,"priority_score_factors":[{"type":"fixability","label":true,"score":214},{"type":"severity","label":"medium","score":300},{"type":"scoreVersion","label":"v1","score":1}],"severity":"medium","title":"Use After Free"},{"exploit_maturity":"No Known Exploit","id":"SNYK-ALPINE319-BUSYBOX-6928846","priority_score":514,"priority_score_factors":[{"type":"fixability","label":true,"score":214},{"type":"severity","label":"medium","score":300},{"type":"scoreVersion","label":"v1","score":1}],"severity":"medium","title":"Use After Free"},{"exploit_maturity":"No Known Exploit","id":"SNYK-ALPINE319-BUSYBOX-6928847","priority_score":514,"priority_score_factors":[{"type":"fixability","label":true,"score":214},{"type":"severity","label":"medium","score":300},{"type":"scoreVersion","label":"v1","score":1}],"severity":"medium","title":"Use After Free"},{"exploit_maturity":"No Known Exploit","id":"SNYK-ALPINE319-BUSYBOX-6913413","priority_score":514,"priority_score_factors":[{"type":"fixability","label":true,"score":214},{"type":"severity","label":"medium","score":300},{"type":"scoreVersion","label":"v1","score":1}],"severity":"medium","title":"Out-of-bounds Write"}],"prId":"0a624b55-d88d-459f-8d43-0f7c739f0bfa","prPublicId":"0a624b55-d88d-459f-8d43-0f7c739f0bfa","packageManager":"dockerfile","priorityScoreList":[514,514,514,514],"projectPublicId":"f4635cb3-382d-48ac-93bb-f4b105d4dcf4","projectUrl":"https://app.snyk.io/org/hybrid-cloud-experience-red-hat-insights/project/f4635cb3-382d-48ac-93bb-f4b105d4dcf4?utm_source=github&utm_medium=referral&page=fix-pr","prType":"fix","templateFieldSources":{"branchName":"default","commitMessage":"default","description":"default","title":"default"},"templateVariants":["updated-fix-title","priorityScore"],"type":"auto","upgrade":["SNYK-ALPINE319-BUSYBOX-6913413","SNYK-ALPINE319-BUSYBOX-6913413","SNYK-ALPINE319-BUSYBOX-6928845","SNYK-ALPINE319-BUSYBOX-6928846","SNYK-ALPINE319-BUSYBOX-6928847"],"vulns":["SNYK-ALPINE319-BUSYBOX-6913413","SNYK-ALPINE319-BUSYBOX-6928845","SNYK-ALPINE319-BUSYBOX-6928846","SNYK-ALPINE319-BUSYBOX-6928847"],"patch":[],"isBreakingChange":false,"remediationStrategy":"vuln"}'


---

## Discussion

### Comment by @jira-linking on October 02, 2024 at 08:02 AM UTC

Commits missing Jira IDs:
5ec38e4fac47a7b0e1ef5abfadc0f3ea2912aa34


### Comment by @codecov-commenter on October 02, 2024 at 08:08 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1499?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 65.03%. Comparing base [(`20de9cd`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/20de9cd5ee7f2ccee530d25c578ee07995e5f0fb?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`5ec38e4`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/5ec38e4fac47a7b0e1ef5abfadc0f3ea2912aa34?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Report is 1 commits behind head on master.

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1499   +/-   ##
=======================================
  Coverage   65.03%   65.03%           
=======================================
  Files         114      114           
  Lines        7880     7880           
=======================================
  Hits         5125     5125           
  Misses       2216     2216           
  Partials      539      539           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1499/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1499/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `65.03% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1499?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @MichaelMraka on October 02, 2024 at 01:45 PM UTC

Duplicate of #1498

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1499*
