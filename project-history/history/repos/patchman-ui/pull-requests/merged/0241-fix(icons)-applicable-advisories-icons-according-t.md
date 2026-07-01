---
type: pull_request
number: 241
title: "fix(Icons): applicable advisories icons according to mockups"
state: merged
author: mkholjuraev
created: 2020-08-20T11:05:55Z
updated: 2021-08-09T06:55:31Z
closed: 2020-08-21T09:45:01Z
merged: 2020-08-21T09:45:01Z
base_branch: master
head_branch: fix-icon-space
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/241
---

# Pull Request #241: fix(Icons): applicable advisories icons according to mockups

**Author**: @mkholjuraev
**Created**: August 20, 2020 at 11:05 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `fix-icon-space`

## Description

fixes: [SPM-424](https://projects.engineering.redhat.com/browse/SPM-424) and [SPM-387](https://projects.engineering.redhat.com/browse/SPM-387)

1. Icon spacing in Inventory table ( could not run Inventory locally to check if actually fixed, but I believe it is fixed)
2. Security advisories icon should be black.
3. Order of the icons should be:  [1] security advisories [2] bug fixes [3] enhancements  
4. Script is added to update snapshots `"test:update": "jest --updateSnapshot --passWithNoTests",`

---

## Discussion

### Comment by @codecov-commenter on August 20, 2020 at 02:11 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/241?src=pr&el=h1) Report
> Merging [#241](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/241?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/984d4e80eabbe8a9d8884a58970800cae6cce2eb?el=desc) will **not change** coverage.
> The diff coverage is `n/a`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/241/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/241?src=pr&el=tree)

```diff
@@           Coverage Diff           @@
##           master     #241   +/-   ##
=======================================
  Coverage   48.34%   48.34%           
=======================================
  Files          55       55           
  Lines         937      937           
  Branches      101      101           
=======================================
  Hits          453      453           
  Misses        443      443           
  Partials       41       41           
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/241?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [...sentationalComponents/AdvisoryType/AdvisoryType.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/241/diff?src=pr&el=tree#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9BZHZpc29yeVR5cGUvQWR2aXNvcnlUeXBlLmpz) | `50.00% <ø> (ø)` | |
| [...ationalComponents/Snippets/AdvisorySeverityInfo.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/241/diff?src=pr&el=tree#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TbmlwcGV0cy9BZHZpc29yeVNldmVyaXR5SW5mby5qcw==) | `100.00% <ø> (ø)` | |
| [...ntationalComponents/Snippets/PortalAdvisoryLink.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/241/diff?src=pr&el=tree#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TbmlwcGV0cy9Qb3J0YWxBZHZpc29yeUxpbmsuanM=) | `100.00% <ø> (ø)` | |
| [src/Utilities/Helpers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/241/diff?src=pr&el=tree#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | `98.30% <ø> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/241?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/241?src=pr&el=footer). Last update [984d4e8...7275fc5](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/241?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @jiridostal on August 21, 2020 at 09:53 AM UTC

:tada: This PR is included in version 0.19.1 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v0.19.1)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/0.19.1)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @jiridostal - Commented on August 20, 2020 at 11:54 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/241*
