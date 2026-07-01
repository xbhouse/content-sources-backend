---
type: pull_request
number: 675
title: "feat(AdvisoryType): SPM-1197 add advisory type name"
state: merged
author: mkholjuraev
created: 2021-10-25T18:27:40Z
updated: 2022-05-17T08:49:53Z
closed: 2021-11-01T13:00:16Z
merged: 2021-11-01T13:00:16Z
base_branch: master
head_branch: advisory-type
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/675
---

# Pull Request #675: feat(AdvisoryType): SPM-1197 add advisory type name

**Author**: @mkholjuraev
**Created**: October 25, 2021 at 06:27 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `advisory-type`

## Description

*No description provided*

---

## Discussion

### Comment by @jiridostal on October 26, 2021 at 12:56 PM UTC

![obrazek](https://user-images.githubusercontent.com/6286045/138883056-dedd7778-65a5-44e9-b24a-bfc075fd9cb3.png)

I see advisory types just as simple text?

### Comment by @mkholjuraev on October 26, 2021 at 01:04 PM UTC

Ups, I forgot to commit the new changes :)

### Comment by @mkholjuraev on October 26, 2021 at 02:06 PM UTC

I have made some changes to the design. Now advisory type is displayed on the right side, above security and reboot
 required info. The new design is approved by UX.

![image](https://user-images.githubusercontent.com/59481011/138895642-09a3175a-6933-4b9d-beb7-16562d957f9a.png)


### Comment by @codecov-commenter on October 26, 2021 at 02:11 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/675?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#675](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/675?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (158eb30) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/3550fb90dee95b44d67817f93fe15de66ae88568?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (3550fb9) will **decrease** coverage by `0.03%`.
> The diff coverage is `80.00%`.

> :exclamation: Current head 158eb30 differs from pull request most recent head 362852a. Consider uploading reports for the commit 362852a to get more accurate results
[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/675/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/675?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #675      +/-   ##
==========================================
- Coverage   83.09%   83.05%   -0.04%     
==========================================
  Files          83       83              
  Lines        1857     1859       +2     
  Branches      460      463       +3     
==========================================
+ Hits         1543     1544       +1     
- Misses        277      278       +1     
  Partials       37       37              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/675?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...ntationalComponents/PackageHeader/PackageHeader.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/675/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9QYWNrYWdlSGVhZGVyL1BhY2thZ2VIZWFkZXIuanM=) | `83.33% <75.00%> (-16.67%)` | :arrow_down: |
| [...ationalComponents/AdvisoryHeader/AdvisoryHeader.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/675/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9BZHZpc29yeUhlYWRlci9BZHZpc29yeUhlYWRlci5qcw==) | `85.71% <100.00%> (+0.71%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/675?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/675?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [3550fb9...362852a](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/675?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @jiridostal on November 01, 2021 at 01:15 PM UTC

:tada: This PR is included in version 1.36.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.36.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.36.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/675*
