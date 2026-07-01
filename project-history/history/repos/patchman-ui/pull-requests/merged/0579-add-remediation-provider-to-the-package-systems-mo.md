---
type: pull_request
number: 579
title: "Add remediation provider to the package-systems modal"
state: merged
author: jiridostal
created: 2021-07-01T16:24:14Z
updated: 2022-08-02T08:37:17Z
closed: 2021-07-22T10:15:51Z
merged: 2021-07-22T10:15:51Z
base_branch: master
head_branch: remediation-package-systems
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/579
---

# Pull Request #579: Add remediation provider to the package-systems modal

**Author**: @jiridostal
**Created**: July 01, 2021 at 04:24 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `remediation-package-systems`

## Description

Added a handler for remediations.

---

## Discussion

### Comment by @codecov-commenter on July 02, 2021 at 07:57 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/579?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#579](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/579?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (b04cdeb) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/e9faac1f686bfa1d5191c24a6d326e14ad923193?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (e9faac1) will **decrease** coverage by `0.26%`.
> The diff coverage is `5.26%`.

> :exclamation: Current head b04cdeb differs from pull request most recent head eaa75f7. Consider uploading reports for the commit eaa75f7 to get more accurate results
[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/579/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/579?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #579      +/-   ##
==========================================
- Coverage   53.90%   53.63%   -0.27%     
==========================================
  Files          76       76              
  Lines        1792     1801       +9     
  Branches      378      378              
==========================================
  Hits          966      966              
- Misses        750      758       +8     
- Partials       76       77       +1     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/579?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/SmartComponents/Advisories/Advisories.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/579/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yaWVzL0Fkdmlzb3JpZXMuanM=) | `82.97% <0.00%> (ø)` | |
| [...c/SmartComponents/PackageSystems/PackageSystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/579/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYWNrYWdlU3lzdGVtcy9QYWNrYWdlU3lzdGVtcy5qcw==) | `0.00% <0.00%> (ø)` | |
| [...rc/SmartComponents/Remediation/RemediationModal.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/579/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9SZW1lZGlhdGlvbi9SZW1lZGlhdGlvbk1vZGFsLmpz) | `13.63% <ø> (ø)` | |
| [src/SmartComponents/Systems/Systems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/579/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | `0.00% <ø> (ø)` | |
| [src/Utilities/DataMappers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/579/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9EYXRhTWFwcGVycy5qcw==) | `71.15% <ø> (ø)` | |
| [src/Utilities/Helpers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/579/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | `78.10% <33.33%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/579?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/579?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [e9faac1...eaa75f7](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/579?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @jiridostal on July 22, 2021 at 10:24 AM UTC

:tada: This PR is included in version 1.26.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.26.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.26.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @mkholjuraev - Commented on July 07, 2021 at 08:50 AM UTC

There is an issue with the number of selected systems. When you select a host and open the remediation button, create remediation, unselect the host and select another host, open remediation modal previously selected host(s) are still present 

### Review by @mkholjuraev - Approved on July 22, 2021 at 10:15 AM UTC

Looks great!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/579*
