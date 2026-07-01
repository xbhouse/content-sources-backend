---
type: pull_request
number: 214
title: "Pass pf react core to inventory to properly render inventory"
state: merged
author: karelhala
created: 2020-08-05T14:52:10Z
updated: 2020-08-18T11:32:01Z
closed: 2020-08-10T10:21:09Z
merged: 2020-08-10T10:21:09Z
base_branch: master
head_branch: pass-pf-react-core
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/214
---

# Pull Request #214: Pass pf react core to inventory to properly render inventory

**Author**: @karelhala
**Created**: August 05, 2020 at 02:52 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `pass-pf-react-core`

## Description

There is a new feature brewing in chrome, building chrome with webpack [1]. This PR passes newly required dependencies. In future we'll use Federated modules [2] so there won't be a need for these changes, but this is required as a mid-step solution.

Also, for enabling A/B testing of inventory detail we are working on new inventory component, for now this component will be undefined so use `React.Fragment` as fallback.

[1] https://github.com/RedHatInsights/insights-chrome/pull/842
[2] https://webpack.js.org/concepts/module-federation/

---

## Discussion

### Comment by @codecov-commenter on August 05, 2020 at 03:04 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/214?src=pr&el=h1) Report
> Merging [#214](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/214?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/41558009439e9ac79e1665eeaecc91c2bb7f0b1f&el=desc) will **decrease** coverage by `0.13%`.
> The diff coverage is `0.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/214/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/214?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #214      +/-   ##
==========================================
- Coverage   37.04%   36.91%   -0.14%     
==========================================
  Files          51       51              
  Lines         845      848       +3     
  Branches       94       95       +1     
==========================================
  Hits          313      313              
- Misses        479      481       +2     
- Partials       53       54       +1     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/214?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [...SmartComponents/AffectedSystems/AffectedSystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/214/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZmZlY3RlZFN5c3RlbXMvQWZmZWN0ZWRTeXN0ZW1zLmpz) | `0.00% <ø> (ø)` | |
| [src/SmartComponents/SystemDetail/InventoryPage.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/214/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1EZXRhaWwvSW52ZW50b3J5UGFnZS5qcw==) | `0.00% <0.00%> (ø)` | |
| [src/SmartComponents/Systems/Systems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/214/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | `0.00% <ø> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/214?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/214?src=pr&el=footer). Last update [4155800...91d178b](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/214?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @karelhala on August 05, 2020 at 03:13 PM UTC

ping @jiridostal 

### Comment by @jiridostal on August 18, 2020 at 11:31 AM UTC

:tada: This PR is included in version 0.19.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v0.19.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/0.19.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/214*
