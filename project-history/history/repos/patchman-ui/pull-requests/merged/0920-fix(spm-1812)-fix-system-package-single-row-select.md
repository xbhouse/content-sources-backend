---
type: pull_request
number: 920
title: "fix(SPM-1812): fix system package single row selection"
state: merged
author: mkholjuraev
created: 2022-11-28T22:40:06Z
updated: 2024-04-03T09:21:54Z
closed: 2022-11-30T08:33:19Z
merged: 2022-11-30T08:33:19Z
base_branch: master
head_branch: fix-system-package-select
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/920
---

# Pull Request #920: fix(SPM-1812): fix system package single row selection

**Author**: @mkholjuraev
**Created**: November 28, 2022 at 10:40 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `fix-system-package-select`

## Description

# Description

Associated Jira ticket: [SPM-1812](https://issues.redhat.com/browse/SPM-1812)

This fixes the single-row selection on the system packages table. 

# How to test the PR

1. Visit the systems page
2. Click on any system that has an upgradable package
3. Click on the Packages tab
4. Observe that bulks select is still working
5. Observe that a single row is also fixed. 
6. Observe that remediation works with selected row(s) that are selected using single-row selection
7. Observe that system advisory remediation is not broken

Please include steps to test your PR.

# Before the change
No snapshot is needed.
# After the change
No snapshot is needed.

# Dependent work link


# Checklist:

- [x] The commit message has the Jira ticket linked
- [x] PR has a short description
- [x] Tests for the changes have been added
- [ ] Screenshots before and after the change are added
- [ ] README.md is updated if necessary
- [ ] Needs additional dependent work


---

## Discussion

### Comment by @codecov-commenter on November 28, 2022 at 10:48 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/920?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **69.75**% // Head: **69.77**% // Increases project coverage by **`+0.02%`** :tada:
> Coverage data is based on head [(`b568eef`)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/920?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`b84932b`)](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/b84932b03075d19ec5ee603e38e0a9cb0767841d?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch coverage: 100.00% of modified lines in pull request are covered.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master     #920      +/-   ##
==========================================
+ Coverage   69.75%   69.77%   +0.02%     
==========================================
  Files         111      111              
  Lines        2605     2607       +2     
  Branches      677      679       +2     
==========================================
+ Hits         1817     1819       +2     
  Misses        712      712              
  Partials       76       76              
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/920?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/Utilities/RawDataForTesting.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/920/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9SYXdEYXRhRm9yVGVzdGluZy5qcw==) | `100.00% <ø> (ø)` | |
| [...artComponents/SystemAdvisories/SystemAdvisories.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/920/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1BZHZpc29yaWVzL1N5c3RlbUFkdmlzb3JpZXMuanM=) | `93.18% <100.00%> (+0.15%)` | :arrow_up: |
| [src/Utilities/useOnSelect.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/920/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy91c2VPblNlbGVjdC5qcw==) | `98.52% <100.00%> (+0.02%)` | :arrow_up: |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/920?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on November 30, 2022 at 08:22 AM UTC

@Fewwy thank you for the review! Snapshot testing is not that bad and also QE is writing e2e tests. Actually, this bug is detected by automation tests :).  I will ask them to cover them just to make sure and migrate cypress asap.

### Comment by @jiridostal on November 30, 2022 at 08:50 AM UTC

:tada: This PR is included in version 1.57.3 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.57.3)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.57.3)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @Fewwy - Approved on November 30, 2022 at 12:11 AM UTC

Tested it locally, and it works fine.
One thing that we should improve is writing a better test to check that it is fixed. Snapshot tests aren't a test.
You can check that the checkbox is "checked," but how do you go through a scenario of choosing the package and remediating it? Maybe we should ask QE to write an e2e test for that case?

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/920*
