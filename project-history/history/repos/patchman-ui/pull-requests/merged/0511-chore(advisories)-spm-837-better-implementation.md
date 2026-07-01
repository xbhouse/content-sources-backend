---
type: pull_request
number: 511
title: "chore(Advisories): SPM-837 better implementation"
state: merged
author: mkholjuraev
created: 2021-04-22T10:58:12Z
updated: 2021-04-22T11:36:49Z
closed: 2021-04-22T11:28:17Z
merged: 2021-04-22T11:28:17Z
base_branch: master
head_branch: make-advisories-type
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/511
---

# Pull Request #511: chore(Advisories): SPM-837 better implementation

**Author**: @mkholjuraev
**Created**: April 22, 2021 at 10:58 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `make-advisories-type`

## Description

I  was not aware that Patternfly flex component has a no-wrap style option. Also applicable advisories column is made no-wrap on scroll 

---

## Discussion

### Comment by @codecov-commenter on April 22, 2021 at 11:02 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/511?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#511](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/511?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (6677eff) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/1de5992f18cba1c5b2688909427e47e56d8d2cd3?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (1de5992) will **not change** coverage.
> The diff coverage is `n/a`.

> :exclamation: Current head 6677eff differs from pull request most recent head bdca79b. Consider uploading reports for the commit bdca79b to get more accurate results
[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/511/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/511?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@           Coverage Diff           @@
##           master     #511   +/-   ##
=======================================
  Coverage   54.49%   54.49%           
=======================================
  Files          78       78           
  Lines        1558     1558           
  Branches      183      183           
=======================================
  Hits          849      849           
  Misses        647      647           
  Partials       62       62           
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/511?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...sentationalComponents/AdvisoryType/AdvisoryType.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/511/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9BZHZpc29yeVR5cGUvQWR2aXNvcnlUeXBlLmpz) | `100.00% <ø> (ø)` | |
| [...resentationalComponents/Snippets/AdvisoriesIcon.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/511/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TbmlwcGV0cy9BZHZpc29yaWVzSWNvbi5qcw==) | `66.66% <ø> (ø)` | |
| [src/Utilities/Helpers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/511/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | `86.39% <ø> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/511?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/511?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [1de5992...bdca79b](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/511?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @jiridostal on April 22, 2021 at 11:36 AM UTC

:tada: This PR is included in version 1.17.10 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.17.10)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.17.10)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/511*
