---
type: pull_request
number: 911
title: "fix(SPM-1802): fix remediation provider for package systems"
state: merged
author: mkholjuraev
created: 2022-11-10T11:59:50Z
updated: 2022-11-10T15:19:48Z
closed: 2022-11-10T15:00:23Z
merged: 2022-11-10T15:00:23Z
base_branch: master
head_branch: fix-package-systems-rem
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/911
---

# Pull Request #911: fix(SPM-1802): fix remediation provider for package systems

**Author**: @mkholjuraev
**Created**: November 10, 2022 at 11:59 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `fix-package-systems-rem`

## Description

resolves: https://issues.redhat.com/browse/SPM-1802.

to test:

1. visit pakcages pae.
2. open a package in the table
3. select updatable systems
4. click on remediate button
5. observe that remediation modal is loaded


---

## Discussion

### Comment by @codecov-commenter on November 10, 2022 at 12:06 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/911?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **72.04**% // Head: **71.98**% // Decreases project coverage by **`-0.05%`** :warning:
> Coverage data is based on head [(`46c217b`)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/911?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`9411070`)](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/94110709e4567e89b34b1cefb868243ee0c2c4ad?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch coverage: 10.00% of modified lines in pull request are covered.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master     #911      +/-   ##
==========================================
- Coverage   72.04%   71.98%   -0.06%     
==========================================
  Files         110      110              
  Lines        2572     2574       +2     
  Branches      655      656       +1     
==========================================
  Hits         1853     1853              
  Misses        657      657              
- Partials       62       64       +2     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/911?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...c/SmartComponents/PackageSystems/PackageSystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/911/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYWNrYWdlU3lzdGVtcy9QYWNrYWdlU3lzdGVtcy5qcw==) | `76.78% <10.00%> (-2.85%)` | :arrow_down: |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/911?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @jiridostal on November 10, 2022 at 03:19 PM UTC

:tada: This PR is included in version 1.55.1 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.55.1)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.55.1)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/911*
