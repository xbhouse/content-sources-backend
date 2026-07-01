---
type: pull_request
number: 625
title: "SPM-866: make update_data unique"
state: merged
author: MichaelMraka
created: 2021-04-21T09:27:12Z
updated: 2021-04-21T12:34:30Z
closed: 2021-04-21T12:34:30Z
merged: 2021-04-21T12:34:29Z
base_branch: master
head_branch: pr2
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/625
---

# Pull Request #625: SPM-866: make update_data unique

**Author**: @MichaelMraka
**Created**: April 21, 2021 at 09:27 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pr2`

## Description

it will make system_package table smaller and faster



---

## Discussion

### Comment by @codecov-commenter on April 21, 2021 at 10:03 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/625?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#625](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/625?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (3b027cd) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/6090d29d8ff97260ba343d05255e553989358432?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (6090d29) will **increase** coverage by `0.02%`.
> The diff coverage is `100.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/625/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/625?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #625      +/-   ##
==========================================
+ Coverage   58.72%   58.74%   +0.02%     
==========================================
  Files          72       72              
  Lines        3244     3246       +2     
==========================================
+ Hits         1905     1907       +2     
  Misses       1062     1062              
  Partials      277      277              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `58.74% <100.00%> (+0.02%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/625?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [evaluator/evaluate\_packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/625/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX3BhY2thZ2VzLmdv) | `68.05% <100.00%> (+0.44%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/625?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/625?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [6090d29...3b027cd](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/625?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @josef-hak - Commented on April 21, 2021 at 09:55 AM UTC

### Review by @josef-hak - Approved on April 21, 2021 at 10:38 AM UTC

### Review by @semtexzv - Approved on April 21, 2021 at 12:21 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/625*
