---
type: pull_request
number: 1495
title: "[Snyk] Security upgrade axios from 1.12.2 to 1.13.5"
state: merged
author: jdobes
created: 2026-02-10T10:23:23Z
updated: 2026-02-16T17:37:22Z
closed: 2026-02-16T17:37:06Z
merged: 2026-02-16T17:37:06Z
base_branch: master
head_branch: snyk-fix-55f3816366cfdcba07e077ac6f658c8f
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1495
---

# Pull Request #1495: [Snyk] Security upgrade axios from 1.12.2 to 1.13.5

**Author**: @jdobes
**Created**: February 10, 2026 at 10:23 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `snyk-fix-55f3816366cfdcba07e077ac6f658c8f`

## Description

![snyk-top-banner](https://res.cloudinary.com/snyk/image/upload/r-d/scm-platform/snyk-pull-requests/pr-banner-default.svg)

### Snyk has created this PR to fix 1 vulnerabilities in the npm dependencies of this project.

#### Snyk changed the following file(s):

- `package.json`
- `package-lock.json`




#### Vulnerabilities that will be fixed with an upgrade:

|  | Issue |  
:-------------------------:|:-------------------------
![high severity](https://res.cloudinary.com/snyk/image/upload/w_20,h_20/v1561977819/icon/h.png 'high severity') | Prototype Pollution <br/>[SNYK-JS-AXIOS-15252993](https://snyk.io/vuln/SNYK-JS-AXIOS-15252993) 




---

> [!IMPORTANT]
>
> - Check the changes in this PR to ensure they won't cause issues with your project.
> - Max score is 1000. Note that the real score may have changed since the PR was raised.
> - This PR was automatically created by Snyk using the credentials of a real user.

---

**Note:** _You are seeing this because you or someone else with access to this repository has authorized Snyk to open fix PRs._

For more information: <img src="https://api.segment.io/v1/pixel/track?data=eyJ3cml0ZUtleSI6InJyWmxZcEdHY2RyTHZsb0lYd0dUcVg4WkFRTnNCOUEwIiwiYW5vbnltb3VzSWQiOiJjNGJiMjYxYi01OGNhLTRiYTYtYmJiNC0yNjNkYmU5ZjJiMDAiLCJldmVudCI6IlBSIHZpZXdlZCIsInByb3BlcnRpZXMiOnsicHJJZCI6ImM0YmIyNjFiLTU4Y2EtNGJhNi1iYmI0LTI2M2RiZTlmMmIwMCJ9fQ==" width="0" height="0"/>
🧐 [View latest project report](https://app.snyk.io/org/hybrid-cloud-experience-red-hat-insights/project/abdec9e9-4fb8-4f10-8528-02a6b74b6b33?utm_source&#x3D;github&amp;utm_medium&#x3D;referral&amp;page&#x3D;fix-pr)
📜 [Customise PR templates](https://docs.snyk.io/scan-using-snyk/pull-requests/snyk-fix-pull-or-merge-requests/customize-pr-templates?utm_source=github&utm_content=fix-pr-template)
🛠 [Adjust project settings](https://app.snyk.io/org/hybrid-cloud-experience-red-hat-insights/project/abdec9e9-4fb8-4f10-8528-02a6b74b6b33?utm_source&#x3D;github&amp;utm_medium&#x3D;referral&amp;page&#x3D;fix-pr/settings)
📚 [Read about Snyk's upgrade logic](https://docs.snyk.io/scan-with-snyk/snyk-open-source/manage-vulnerabilities/upgrade-package-versions-to-fix-vulnerabilities?utm_source=github&utm_content=fix-pr-template)

---

**Learn how to fix vulnerabilities with free interactive lessons:**

🦉 [Prototype Pollution](https://learn.snyk.io/lesson/prototype-pollution/?loc&#x3D;fix-pr)

[//]: # 'snyk:metadata:{"breakingChangeRiskLevel":null,"FF_showPullRequestBreakingChanges":false,"FF_showPullRequestBreakingChangesWebSearch":false,"customTemplate":{"variablesUsed":[],"fieldsUsed":[]},"dependencies":[{"name":"axios","from":"1.12.2","to":"1.13.5"}],"env":"prod","issuesToFix":["SNYK-JS-AXIOS-15252993"],"prId":"c4bb261b-58ca-4ba6-bbb4-263dbe9f2b00","prPublicId":"c4bb261b-58ca-4ba6-bbb4-263dbe9f2b00","packageManager":"npm","priorityScoreList":[null],"projectPublicId":"abdec9e9-4fb8-4f10-8528-02a6b74b6b33","projectUrl":"https://app.snyk.io/org/hybrid-cloud-experience-red-hat-insights/project/abdec9e9-4fb8-4f10-8528-02a6b74b6b33?utm_source=github&utm_medium=referral&page=fix-pr","prType":"fix","templateFieldSources":{"branchName":"default","commitMessage":"default","description":"default","title":"default"},"templateVariants":["updated-fix-title"],"type":"auto","upgrade":["SNYK-JS-AXIOS-15252993"],"vulns":["SNYK-JS-AXIOS-15252993"],"patch":[],"isBreakingChange":false,"remediationStrategy":"vuln"}'


---

## Discussion

### Comment by @swadeley on February 10, 2026 at 11:07 AM UTC

@dependabot rebase

### Comment by @swadeley on February 10, 2026 at 01:10 PM UTC

@dependabot rebase

### Comment by @codecov-commenter on February 16, 2026 at 02:59 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1495?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 62.53%. Comparing base ([`1c1d35a`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/1c1d35a84b44bdf268ec90be59583f6a1b12fc06?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`36d30b1`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/36d30b133f14f2a05159f2192685c0a88b322604?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 1 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1495   +/-   ##
=======================================
  Coverage   62.53%   62.53%           
=======================================
  Files         127      127           
  Lines        3310     3310           
  Branches      892      899    +7     
=======================================
  Hits         2070     2070           
  Misses       1120     1120           
  Partials      120      120           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1495/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [combined](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1495/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `?` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1495?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

---

## Reviews

### Review by @TenSt - Approved on February 12, 2026 at 11:11 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1495*
