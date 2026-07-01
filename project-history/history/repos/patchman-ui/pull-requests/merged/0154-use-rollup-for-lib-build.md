---
type: pull_request
number: 154
title: "Use rollup for lib build"
state: merged
author: karelhala
created: 2020-06-05T09:31:49Z
updated: 2020-06-08T11:20:43Z
closed: 2020-06-08T11:13:23Z
merged: 2020-06-08T11:13:23Z
base_branch: master
head_branch: introduce-rollup
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/154
---

# Pull Request #154: Use rollup for lib build

**Author**: @karelhala
**Created**: June 05, 2020 at 09:31 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `introduce-rollup`

## Description

This will improve libraru build to include multiple environments and allow easier treeshaking, once this library is used somewhere it will produce much smaller app bundle.

* [x] merge after #149

---

## Discussion

### Comment by @karelhala on June 05, 2020 at 09:32 AM UTC

ping @jiridostal 

### Comment by @jiridostal on June 08, 2020 at 07:50 AM UTC

Could you please resolve the conflicts?

### Comment by @karelhala on June 08, 2020 at 09:56 AM UTC

@jiridostal done. I've released beta version of this package, to test it in inventory and it went splendidly! With this change and vuln change as well, the inventory UI dropped from 10MB to 3MB. (compressed from 2MB to 700KB).

### Comment by @codecov-commenter on June 08, 2020 at 10:59 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/154?src=pr&el=h1) Report
> Merging [#154](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/154?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/6c59a6301ccc129d0218acc0fd589273d2cf2224&el=desc) will **decrease** coverage by `0.09%`.
> The diff coverage is `0.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/154/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/154?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #154      +/-   ##
==========================================
- Coverage   37.36%   37.27%   -0.10%     
==========================================
  Files          50       51       +1     
  Lines         819      821       +2     
  Branches       82       82              
==========================================
  Hits          306      306              
- Misses        468      470       +2     
  Partials       45       45              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/154?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [...SmartComponents/AffectedSystems/AffectedSystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/154/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZmZlY3RlZFN5c3RlbXMvQWZmZWN0ZWRTeXN0ZW1zLmpz) | `0.00% <ø> (ø)` | |
| [src/SmartComponents/Remediation/Remediation.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/154/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9SZW1lZGlhdGlvbi9SZW1lZGlhdGlvbi5qcw==) | `0.00% <ø> (ø)` | |
| [...rc/SmartComponents/Remediation/RemediationModal.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/154/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9SZW1lZGlhdGlvbi9SZW1lZGlhdGlvbk1vZGFsLmpz) | `0.00% <ø> (ø)` | |
| [src/SmartComponents/SystemDetail/InventoryPage.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/154/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1EZXRhaWwvSW52ZW50b3J5UGFnZS5qcw==) | `0.00% <ø> (ø)` | |
| [src/SmartComponents/Systems/Systems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/154/diff?src=pr&el=tree#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | `0.00% <ø> (ø)` | |
| [src/index.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/154/diff?src=pr&el=tree#diff-c3JjL2luZGV4Lmpz) | `0.00% <0.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/154?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/154?src=pr&el=footer). Last update [6c59a63...828fc84](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/154?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @jiridostal on June 08, 2020 at 11:13 AM UTC

Heh, that sounds really cool. Thank you for that!

### Comment by @jiridostal on June 08, 2020 at 11:20 AM UTC

:tada: This PR is included in version 0.12.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v0.12.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/0.12.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/154*
