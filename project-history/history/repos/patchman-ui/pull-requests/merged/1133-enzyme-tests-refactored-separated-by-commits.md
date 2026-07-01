---
type: pull_request
number: 1133
title: "Enzyme tests refactored separated by commits"
state: merged
author: Fewwy
created: 2023-10-20T12:03:47Z
updated: 2023-12-13T11:16:05Z
closed: 2023-11-24T12:47:47Z
merged: 2023-11-24T12:47:47Z
base_branch: master
head_branch: test-refactor
labels: ["released", "refactor"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1133
---

# Pull Request #1133: Enzyme tests refactored separated by commits

**Author**: @Fewwy
**Created**: October 20, 2023 at 12:03 PM UTC
**Status**: Merged
**Labels**: `released`, `refactor`
**Base**: `master` ← **Head**: `test-refactor`

## Description

This is the same as https://github.com/RedHatInsights/patchman-ui/pull/1123 PR but separated by commits/file changed

---

## Discussion

### Comment by @codecov-commenter on October 20, 2023 at 12:12 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1133?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:

[see 4 files with indirect coverage changes](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1133/indirect-changes?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)


:loudspeaker: Thoughts on this report? [Let us know!](https://about.codecov.io/pull-request-comment-report/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)


### Comment by @Fewwy on October 20, 2023 at 03:04 PM UTC

/retest

### Comment by @Fewwy on October 24, 2023 at 02:36 PM UTC

/retest

### Comment by @Fewwy on October 26, 2023 at 07:46 AM UTC

/retest

### Comment by @Fewwy on October 26, 2023 at 08:26 AM UTC

/retest

### Comment by @niyazRedhat on October 27, 2023 at 02:52 PM UTC

/retest

### Comment by @Fewwy on October 27, 2023 at 04:41 PM UTC

/retest

### Comment by @Fewwy on November 24, 2023 at 12:47 PM UTC

@johnsonm325 As far as I understand - yes

### Comment by @mkholjuraev on December 13, 2023 at 11:15 AM UTC

:tada: This PR is included in version 1.65.2 :tada:

The release is available on [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.65.2)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @johnsonm325 - Approved on November 13, 2023 at 08:16 PM UTC

Approved. But I do have one question. It looks like the snapshots now remove the names of the child components and replaces them with pure html like `<h5>` and `<div>` rather than `<Breadcrumb>` and the like. Is this the kind of output we expect to see from RTL?

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1133*
