---
type: pull_request
number: 956
title: "perf(RHIF-171): Get rid of activeApps in redux"
state: merged
author: mkholjuraev
created: 2023-02-06T19:46:32Z
updated: 2024-04-03T09:22:47Z
closed: 2023-02-16T16:12:01Z
merged: 2023-02-16T16:12:01Z
base_branch: master
head_branch: system-detail-refactor
labels: ["released", "refactor"]
url: https://github.com/RedHatInsights/patchman-ui/pull/956
---

# Pull Request #956: perf(RHIF-171): Get rid of activeApps in redux

**Author**: @mkholjuraev
**Created**: February 06, 2023 at 07:46 PM UTC
**Status**: Merged
**Labels**: `released`, `refactor`
**Base**: `master` ← **Head**: `system-detail-refactor`

## Description

# Description

Associated Jira ticket: # [RHIF-171](https://issues.redhat.com/browse/RHIF-171)

This PR removes AppInfo federated module usage together with rendering the SystemDetail component using activeApps in redux. Thus, making the application's complexity and performance.

NEEDS TO BE MERGED AFTER: https://github.com/RedHatInsights/insights-inventory-frontend/pull/1766

# How to test the PR

Run both PRs together and make sure no app breaks. Basic functionality should not change. Observe that the system details page works as expected.

# Before the change


# After the change


# Dependent work link


# Checklist:

- [x] The commit message has the Jira ticket linked
- [x] PR has a short description
- [ ] Screenshots before and after the change are added
- [ ] Tests for the changes have been added
- [ ] README.md is updated if necessary
- [ ] Needs additional dependent work


---

## Discussion

### Comment by @mkholjuraev on February 10, 2023 at 11:24 AM UTC

oh, there is the inventory app that has an inventory id different in useParams. Will do something about it.

### Comment by @codecov-commenter on February 10, 2023 at 03:08 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/956?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **72.35**% // Head: **72.17**% // Decreases project coverage by **`-0.18%`** :warning:
> Coverage data is based on head [(`529a25d`)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/956?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`6c20611`)](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/6c20611f57f3cea5b078416b2f79a2bfe4bf9054?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch coverage: 78.57% of modified lines in pull request are covered.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master     #956      +/-   ##
==========================================
- Coverage   72.35%   72.17%   -0.18%     
==========================================
  Files         110      110              
  Lines        2615     2609       -6     
  Branches      658      657       -1     
==========================================
- Hits         1892     1883       -9     
- Misses        723      726       +3     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/956?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...rc/SmartComponents/SystemDetail/InventoryDetail.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/956?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1EZXRhaWwvSW52ZW50b3J5RGV0YWlsLmpz) | `83.33% <0.00%> (-0.54%)` | :arrow_down: |
| [src/store/Reducers/SystemDetailStore.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/956?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL1JlZHVjZXJzL1N5c3RlbURldGFpbFN0b3JlLmpz) | `60.00% <ø> (-23.34%)` | :arrow_down: |
| [...artComponents/SystemAdvisories/SystemAdvisories.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/956?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1BZHZpc29yaWVzL1N5c3RlbUFkdmlzb3JpZXMuanM=) | `93.02% <100.00%> (-0.16%)` | :arrow_down: |
| [src/SmartComponents/SystemDetail/SystemDetail.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/956?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1EZXRhaWwvU3lzdGVtRGV0YWlsLmpz) | `71.42% <100.00%> (-1.91%)` | :arrow_down: |
| [...c/SmartComponents/SystemPackages/SystemPackages.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/956?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1QYWNrYWdlcy9TeXN0ZW1QYWNrYWdlcy5qcw==) | `100.00% <100.00%> (ø)` | |
| [src/store/index.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/956?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL2luZGV4Lmpz) | `100.00% <100.00%> (ø)` | |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/956?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on February 16, 2023 at 04:29 PM UTC

:tada: This PR is included in version 1.59.2 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.59.2)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.59.2)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @gkarat - Changes Requested on February 10, 2023 at 10:51 AM UTC

Hey @mkholjuraev, I took a look and run the PR together with Inventory locally. I am afraid that the inventory details page doesn't load the Patch tab properly and throws errors (see the screenshot). I compared stage-stable (stage.foo.redhat.com:1337/insights/inventory vs console.stage.redhat.com/insights/inventory).

a message from the error that most probably caused the crash:
```
Uncaught (in promise) 
Object { title: "There was an error getting data", detail: "incorrect inventory_id format", status: 400 }
vendors-node_modules_redux-promise-middleware_dist_es_index_js.32f0fccd2880bfa38c39.js:173:11
```
![image](https://user-images.githubusercontent.com/31385370/218073891-4f6a5644-851c-4f47-b08d-bb7d2ff6db8c.png)

### Review by @adonispuente - Dismissed on February 14, 2023 at 05:57 PM UTC

Sorry , I didnt mean to approve, I was still looking at it

### Review by @gkarat - Approved on February 16, 2023 at 12:12 PM UTC

LGTM! works fine locally with inventory deployed together. nothing to call out from the code TBH.

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/956*
