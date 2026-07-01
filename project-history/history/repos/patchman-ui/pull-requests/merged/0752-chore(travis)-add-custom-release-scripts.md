---
type: pull_request
number: 752
title: "chore(Travis): add custom release scripts"
state: merged
author: mkholjuraev
created: 2022-03-15T13:47:19Z
updated: 2022-05-17T08:50:41Z
closed: 2022-03-16T10:33:32Z
merged: 2022-03-16T10:33:32Z
base_branch: master
head_branch: custom-release
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/752
---

# Pull Request #752: chore(Travis): add custom release scripts

**Author**: @mkholjuraev
**Created**: March 15, 2022 at 01:47 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `custom-release`

## Description

After Node.js upgrade to  version 16, Travis CI build is failing. @leSamo had already this situation and suggested writing custom release. The PR is inspired by [Vulnerability](https://github.com/RedHatInsights/vulnerability-ui/blob/master/.travis/custom_release.sh). I am not familiar with Travis, so any improvements are highly welcome.

---

## Discussion

### Comment by @codecov-commenter on March 15, 2022 at 02:38 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/752?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#752](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/752?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (c26a3d8) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/54839daeaa9815294b486ab2977bf45e3fb597ba?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (54839da) will **increase** coverage by `4.84%`.
> The diff coverage is `56.81%`.

> :exclamation: Current head c26a3d8 differs from pull request most recent head 6a3d275. Consider uploading reports for the commit 6a3d275 to get more accurate results

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/752/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/752?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #752      +/-   ##
==========================================
+ Coverage   66.47%   71.31%   +4.84%     
==========================================
  Files          99      100       +1     
  Lines        2246     2214      -32     
  Branches      562      555       -7     
==========================================
+ Hits         1493     1579      +86     
+ Misses        690      579     -111     
+ Partials       63       56       -7     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/752?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...Components/StatusReports/AdvisoriesStatusReport.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/752/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TdGF0dXNSZXBvcnRzL0Fkdmlzb3JpZXNTdGF0dXNSZXBvcnQuanM=) | `53.33% <ø> (+3.33%)` | :arrow_up: |
| [src/store/index.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/752/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL3N0b3JlL2luZGV4Lmpz) | `73.07% <ø> (ø)` | |
| [...rtComponents/Remediation/AsyncRemediationButton.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/752/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9SZW1lZGlhdGlvbi9Bc3luY1JlbWVkaWF0aW9uQnV0dG9uLmpz) | `33.33% <33.33%> (ø)` | |
| [...SmartComponents/AdvisorySystems/AdvisorySystems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/752/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yeVN5c3RlbXMvQWR2aXNvcnlTeXN0ZW1zLmpz) | `85.41% <37.50%> (-9.94%)` | :arrow_down: |
| [src/SmartComponents/Systems/Systems.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/752/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXMuanM=) | `66.66% <42.85%> (+66.66%)` | :arrow_up: |
| [src/Utilities/Helpers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/752/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | `77.22% <50.00%> (-0.78%)` | :arrow_down: |
| [src/Utilities/api.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/752/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9hcGkuanM=) | `54.00% <50.00%> (+0.46%)` | :arrow_up: |
| [src/Utilities/unitTestingUtilities.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/752/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy91bml0VGVzdGluZ1V0aWxpdGllcy5qcw==) | `18.75% <50.00%> (-9.83%)` | :arrow_down: |
| [...c/SmartComponents/Remediation/RemediationWizard.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/752/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9SZW1lZGlhdGlvbi9SZW1lZGlhdGlvbldpemFyZC5qcw==) | `66.66% <66.66%> (ø)` | |
| [...nalComponents/StatusReports/SystemsStatusReport.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/752/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TdGF0dXNSZXBvcnRzL1N5c3RlbXNTdGF0dXNSZXBvcnQuanM=) | `90.00% <100.00%> (+90.00%)` | :arrow_up: |
| ... and [28 more](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/752/diff?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/752?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/752?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [a610d4e...6a3d275](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/752?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @leSamo on March 16, 2022 at 09:54 AM UTC

Almost perfect, but in your script you are pushing to qa-stable env from master branch and to stage-stable env from stage-stable-branch. Both qa-stable and stage-stable map to stage-stable env and qa-beta and stage-beta map to stage-beta env as shown in this [release script](https://github.com/RedHatInsights/insights-frontend-builder-common/blob/9439bff6073515eb7f7fe8046ee1d1881c332870/src/Jenkinsfile#L16):
![Screenshot from 2022-03-16 10-50-23](https://user-images.githubusercontent.com/8426204/158563199-1554f0e2-4a44-4fce-a9a4-1dd1f62f4684.png)

Which would mean you are pushing to a single env from multiple branches which is undesirable and could cause confusion.

This PR will stop support for ci-beta env, I don't believe anyone is using it anymore, but better make sure.

BTW if you'd like to see the actual scripts running on Travis which will parse this custom release script you could find it here:
https://github.com/RedHatInsights/insights-frontend-builder-common

While we're at it, it would be nice to update these resources with up-to-date branch to env mappings:
https://github.com/RedHatInsights/patchman-ui#how-it-works
https://docs.engineering.redhat.com/pages/viewpage.action?spaceKey=SPM&title=Patch+UI+release+workflow

### Comment by @mkholjuraev on March 16, 2022 at 10:02 AM UTC

Do you suggest to  remove  pushing to qa-stable from master?

### Comment by @mkholjuraev on March 16, 2022 at 10:34 AM UTC

@leSamo thanks for the review!

### Comment by @jiridostal on March 18, 2022 at 02:07 PM UTC

:tada: This PR is included in version 1.40.4 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.40.4)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.40.4)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @leSamo - Approved on March 16, 2022 at 10:32 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/752*
