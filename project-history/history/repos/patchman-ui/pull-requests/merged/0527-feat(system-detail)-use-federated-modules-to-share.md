---
type: pull_request
number: 527
title: "feat(system detail): use federated modules to share system detail"
state: merged
author: karelhala
created: 2021-05-06T11:15:14Z
updated: 2021-05-11T09:50:07Z
closed: 2021-05-11T09:41:43Z
merged: 2021-05-11T09:41:43Z
base_branch: master
head_branch: modularize-system-detail
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/527
---

# Pull Request #527: feat(system detail): use federated modules to share system detail

**Author**: @karelhala
**Created**: May 06, 2021 at 11:15 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `modularize-system-detail`

## Description

### Use federated modules to share system detail

Since we can now use federated modules in every environment let's share the system detail with it. This way we don't have to maintain the version of it in all environments, it's the same for all users on one environment. If a bug is found the fix is deployed just in this application and propagated to inventory via federetad modules.

### Usage
```JSX
import React, { useContext } from 'react';
import AsyncComponent from '@redhat-cloud-services/frontend-components/AsyncComponent';
import { RegistryContext } from '../../store';
import { useRouteMatch } from 'react-router-dom';

const PatchTab = () => (
  <AsyncComponent
      appName="patch"
      module="./SystemDetail"
      getRegistry={useContext(RegistryContext).getRegistry}
  />);
```

---

## Discussion

### Comment by @karelhala on May 06, 2021 at 11:15 AM UTC

ping @mkholjuraev @jiridostal 

### Comment by @karelhala on May 06, 2021 at 02:30 PM UTC

@mkholjuraev you'll have to run patch app on port 8002 and inventory on port 8003 and change your config to serve `inventory`. I'll add this to your spandx file

### Comment by @codecov-commenter on May 06, 2021 at 02:46 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/527?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#527](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/527?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (2a65052) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/24347032afee3d5ed59953d8ce8f56cbe060bce2?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (2434703) will **decrease** coverage by `0.24%`.
> The diff coverage is `0.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/527/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/527?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #527      +/-   ##
==========================================
- Coverage   55.25%   55.01%   -0.25%     
==========================================
  Files          78       78              
  Lines        1817     1825       +8     
  Branches      383      387       +4     
==========================================
  Hits         1004     1004              
- Misses        754      760       +6     
- Partials       59       61       +2     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/527?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/index.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/527/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL2luZGV4Lmpz) | `0.00% <0.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/527?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/527?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [2434703...2a65052](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/527?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @jiridostal on May 11, 2021 at 09:50 AM UTC

:tada: This PR is included in version 1.19.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.19.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.19.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @mkholjuraev - Approved on May 06, 2021 at 12:15 PM UTC

The PR looks good to me, but I was not able to run it locally and try it out. Inventory app breaks with the error: **Error while loading component!**. 

![image](https://user-images.githubusercontent.com/59481011/117296537-76a71480-ae75-11eb-953c-a3719c01209f.png)


### Review by @mkholjuraev - Approved on May 11, 2021 at 09:41 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/527*
