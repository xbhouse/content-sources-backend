---
type: pull_request
number: 500
title: "fix(Systems): SPM-847 remove extra mounted remediation modal"
state: closed
author: mkholjuraev
created: 2021-04-13T22:25:54Z
updated: 2021-08-09T06:54:15Z
closed: 2021-04-14T11:30:59Z
base_branch: master
head_branch: remove-extra-modal
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/500
---

# Pull Request #500: fix(Systems): SPM-847 remove extra mounted remediation modal

**Author**: @mkholjuraev
**Created**: April 13, 2021 at 10:25 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `remove-extra-modal`

## Description

It looks like the useEffect with empty dependencies is being triggered twice which contradicts its design. I was not able to find the cause of it. I suggest using this temporary fix for now. We can dig it more after the summit. 

What I tried so far:
1. most likely it is triggered due to parent unmount remount events. I have tried tracing which parent, but I could not trace.
2. create a ref to remediation modal and try triggering the effect if the modal already exists
2. create a state [mounted, setMounted] = useState(false);


---

## Discussion

### Comment by @codecov-io on April 13, 2021 at 10:30 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/500?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#500](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/500?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (e42b2c4) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/d201c6a4ffabe39d60fb6a4d3841e24477ffe21b?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (d201c6a) will **decrease** coverage by `0.06%`.
> The diff coverage is `0.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/500/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/500?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #500      +/-   ##
==========================================
- Coverage   54.49%   54.42%   -0.07%     
==========================================
  Files          78       78              
  Lines        1558     1560       +2     
  Branches      183      183              
==========================================
  Hits          849      849              
- Misses        647      649       +2     
  Partials       62       62              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/500?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...rc/SmartComponents/Remediation/RemediationModal.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/500/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9SZW1lZGlhdGlvbi9SZW1lZGlhdGlvbk1vZGFsLmpz) | `16.66% <0.00%> (-2.09%)` | :arrow_down: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/500?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/500?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [d201c6a...e42b2c4](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/500?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/500*
