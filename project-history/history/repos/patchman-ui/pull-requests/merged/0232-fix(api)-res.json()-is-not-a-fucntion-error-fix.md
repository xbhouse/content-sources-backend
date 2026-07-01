---
type: pull_request
number: 232
title: "fix(api) res.json() is not a fucntion error fix"
state: merged
author: mkholjuraev
created: 2020-08-13T10:24:00Z
updated: 2021-08-09T06:55:33Z
closed: 2020-08-13T12:11:22Z
merged: 2020-08-13T12:11:22Z
base_branch: master
head_branch: fix-api-error-catch
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/232
---

# Pull Request #232: fix(api) res.json() is not a fucntion error fix

**Author**: @mkholjuraev
**Created**: August 13, 2020 at 10:24 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `fix-api-error-catch`

## Description

fixes errors in console: 
        1. res.json() is not a function error
        2. Text input: Text input requires either an id or aria-label to be specified  error 

Before:
![Screenshot from 2020-08-13 11-38-55](https://user-images.githubusercontent.com/59481011/90123905-d5fd4b80-dd5f-11ea-9a43-f7d14c93e91a.png)

After: 
![Screenshot from 2020-08-13 12-18-05](https://user-images.githubusercontent.com/59481011/90123915-dac1ff80-dd5f-11ea-9870-a013ab9527ef.png)


---

## Discussion

### Comment by @codecov-commenter on August 13, 2020 at 10:34 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/232?src=pr&el=h1) Report
> Merging [#232](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/232?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/9557d7fad734287a030091515ea64a1330a36358&el=desc) will **decrease** coverage by `0.08%`.
> The diff coverage is `0.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/232/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/232?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #232      +/-   ##
==========================================
- Coverage   36.91%   36.82%   -0.09%     
==========================================
  Files          51       51              
  Lines         848      850       +2     
  Branches       95       95              
==========================================
  Hits          313      313              
- Misses        481      483       +2     
  Partials       54       54              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/232?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [...c/PresentationalComponents/Filters/SearchFilter.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/232/diff?src=pr&el=tree#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9GaWx0ZXJzL1NlYXJjaEZpbHRlci5qcw==) | `0.00% <ø> (ø)` | |
| [src/Utilities/api.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/232/diff?src=pr&el=tree#diff-c3JjL1V0aWxpdGllcy9hcGkuanM=) | `17.64% <0.00%> (-0.73%)` | :arrow_down: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/232?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/232?src=pr&el=footer). Last update [9557d7f...f18b0a4](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/232?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @jiridostal on August 18, 2020 at 11:31 AM UTC

:tada: This PR is included in version 0.19.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v0.19.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/0.19.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/232*
