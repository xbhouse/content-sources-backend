---
type: pull_request
number: 1079
title: "chore: verify preview env better in prod webpack config"
state: merged
author: mkholjuraev
created: 2023-06-28T12:21:19Z
updated: 2023-07-03T16:56:42Z
closed: 2023-06-28T13:04:22Z
merged: 2023-06-28T13:04:22Z
base_branch: master
head_branch: prod-webpack
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1079
---

# Pull Request #1079: chore: verify preview env better in prod webpack config

**Author**: @mkholjuraev
**Created**: June 28, 2023 at 12:21 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `prod-webpack`

## Description

# Description

Associated Jira ticket: # [(issue)](https://issues.redhat.com/browse/RHINENG-502)

As the app containerized build start getting used, we are required to precisely differentiate preview and stable builds. This makes it possible by adjusting the prod webpack config a bit.


# How to test the PR

Testing happens after gets merged

# Before the change


# After the change


# Dependent work link


# Checklist:

- [ ] The commit message has the Jira ticket linked
- [ ] PR has a short description
- [ ] Screenshots before and after the change are added
- [ ] Tests for the changes have been added
- [ ] README.md is updated if necessary
- [ ] Needs additional dependent work


---

## Discussion

### Comment by @mkholjuraev on June 28, 2023 at 12:30 PM UTC

> Could we use `preview` instead of `beta` for future-proofness?

I do not think we are there yet. As resources are still fetched from /beta prefix

![image](https://github.com/RedHatInsights/patchman-ui/assets/59481011/5bba383c-f118-4ea5-a1ef-80c47aca3d50)


### Comment by @codecov-commenter on June 28, 2023 at 12:30 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1079?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Patch coverage has no change and project coverage change: **`+0.07`** :tada:
> Comparison is base [(`9916b41`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/9916b414c9a8e0abbc951263f5999158f4fe0de4?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.80% compared to head [(`55a2ad7`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1079?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) 62.87%.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1079      +/-   ##
==========================================
+ Coverage   62.80%   62.87%   +0.07%     
==========================================
  Files         119      119              
  Lines        2960     2966       +6     
  Branches      758      761       +3     
==========================================
+ Hits         1859     1865       +6     
  Misses       1101     1101              
```


[see 3 files with indirect coverage changes](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1079/indirect-changes?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1079?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on July 03, 2023 at 04:56 PM UTC

:tada: This PR is included in version 1.63.4 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.63.4)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.63.4)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @leSamo - Commented on June 28, 2023 at 12:27 PM UTC

Could we use `preview` instead of `beta` for future-proofness?

### Review by @leSamo - Approved on June 28, 2023 at 12:31 PM UTC

Ah, I see, good then

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1079*
