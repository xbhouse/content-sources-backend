---
type: pull_request
number: 805
title: "fix(PatchSetWizard): Adding missing message"
state: merged
author: gabipodolnikova
created: 2022-05-26T12:18:21Z
updated: 2022-05-31T07:16:16Z
closed: 2022-05-26T13:32:13Z
merged: 2022-05-26T13:32:13Z
base_branch: master
head_branch: assign-patchset-fix
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/805
---

# Pull Request #805: fix(PatchSetWizard): Adding missing message

**Author**: @gabipodolnikova
**Created**: May 26, 2022 at 12:18 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `assign-patchset-fix`

## Description

I am not sure what the title should be, but there was a missing message causing a blank page after assigning a patch set.

---

## Discussion

### Comment by @codecov-commenter on May 26, 2022 at 12:24 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/805?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#805](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/805?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (78661ff) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/fdbe8bc46429a3a7fc19794b756f9e15a8dbedf5?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (fdbe8bc) will **not change** coverage.
> The diff coverage is `n/a`.

```diff
@@           Coverage Diff           @@
##           master     #805   +/-   ##
=======================================
  Coverage   70.25%   70.25%           
=======================================
  Files         101      101           
  Lines        2414     2414           
  Branches      624      624           
=======================================
  Hits         1696     1696           
  Misses        648      648           
  Partials       70       70           
```



------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/805?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/805?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [fdbe8bc...78661ff](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/805?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @gabipodolnikova on May 26, 2022 at 12:54 PM UTC

> @gabipodolnikova thanks for the help. This message should have somehow disappeared while rebasing commits. Can I please ask to change the message from 'Patch set' to 'Create Patch set'.

yep, done :)



### Comment by @mkholjuraev on May 26, 2022 at 01:32 PM UTC

Merging if you do not mind.

### Comment by @jiridostal on May 31, 2022 at 07:16 AM UTC

:tada: This PR is included in version 1.47.5 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.47.5)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.47.5)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @mkholjuraev - Commented on May 26, 2022 at 12:25 PM UTC

@gabipodolnikova thanks for the help. This message should have somehow disappeared while rebasing commits. Can I please ask to change the message from 'Patch set' to 'Create Patch set'. 

### Review by @mkholjuraev - Approved on May 26, 2022 at 01:19 PM UTC

@gabipodolnikova  :+1: . I could not run this locally due to VPN, but I am sure that this is the fix.

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/805*
