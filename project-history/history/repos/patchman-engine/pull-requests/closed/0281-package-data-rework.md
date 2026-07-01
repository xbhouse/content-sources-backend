---
type: pull_request
number: 281
title: "Package data rework"
state: closed
author: semtexzv
created: 2020-06-29T08:50:51Z
updated: 2020-07-13T08:19:12Z
closed: 2020-07-13T08:19:12Z
base_branch: master
head_branch: pkgdata
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/281
---

# Pull Request #281: Package data rework

**Author**: @semtexzv
**Created**: June 29, 2020 at 08:50 AM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `pkgdata`

## Description

- Adds general `package` table for storing data about packages
- Links packages to systems using `system_package` table
- Automatically updates package table when evaluating
- Adds summary & description for packages
- Adds search & filtering to packages

---

## Discussion

### Comment by @codecov-commenter on June 29, 2020 at 01:22 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/281?src=pr&el=h1) Report
> Merging [#281](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/281?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/acae0c620c69af0f54ebc8c28e97e9703cdccb29&el=desc) will **increase** coverage by `0.22%`.
> The diff coverage is `83.03%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/281/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/281?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #281      +/-   ##
==========================================
+ Coverage   62.20%   62.43%   +0.22%     
==========================================
  Files          51       51              
  Lines        2408     2465      +57     
==========================================
+ Hits         1498     1539      +41     
- Misses        722      732      +10     
- Partials      188      194       +6     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/281?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [base/database/query.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/281/diff?src=pr&el=tree#diff-YmFzZS9kYXRhYmFzZS9xdWVyeS5nbw==) | `80.24% <75.00%> (-2.47%)` | :arrow_down: |
| [manager/controllers/packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/281/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlcy5nbw==) | `58.06% <75.00%> (-5.94%)` | :arrow_down: |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/281/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `73.96% <84.81%> (+1.84%)` | :arrow_up: |
| [evaluator/remediations.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/281/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL3JlbWVkaWF0aW9ucy5nbw==) | `62.50% <100.00%> (ø)` | |
| [manager/controllers/package\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/281/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX3N5c3RlbXMuZ28=) | `70.00% <100.00%> (+2.14%)` | :arrow_up: |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/281/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `74.13% <0.00%> (-2.59%)` | :arrow_down: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/281?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/281?src=pr&el=footer). Last update [acae0c6...b5829b1](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/281?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @semtexzv on July 13, 2020 at 08:19 AM UTC

Closing, superseded by https://github.com/RedHatInsights/patchman-engine/pull/286

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/281*
