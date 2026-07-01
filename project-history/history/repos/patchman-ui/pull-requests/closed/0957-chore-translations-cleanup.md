---
type: pull_request
number: 957
title: "chore: Translations cleanup"
state: closed
author: leSamo
created: 2023-02-14T00:57:01Z
updated: 2023-08-14T13:14:48Z
closed: 2023-07-26T09:31:46Z
base_branch: master
head_branch: translations-cleanup
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/957
---

# Pull Request #957: chore: Translations cleanup

**Author**: @leSamo
**Created**: February 14, 2023 at 12:57 AM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `translations-cleanup`

## Description

# Description
Fixed translations not generating. Added a new package which implements command to detect unused translation keys. This command is set to run on every translation generation (see the yellow text in the bottom of the screenshot. This information can be used to cleanup and improve navigability of the `Messages.js` file.

![Screenshot from 2023-02-14 01-53-50](https://user-images.githubusercontent.com/8426204/218610025-1096355b-a3cf-4f83-b848-733dd2841331.png)

# How to test the PR/Before/After
There should be not change to functionality. Only affects dev env.

# Checklist:

- [ ] The commit message has the Jira ticket linked
- [x] PR has a short description
- [x] Screenshots before and after the change are added
- [ ] Tests for the changes have been added
- [x] README.md is updated if necessary
- [ ] Needs additional dependent work


---

## Discussion

### Comment by @leSamo on February 14, 2023 at 01:00 AM UTC

Hmm, I found some false negatives with the unused keys detection, I'll add update the package to address this.

### Comment by @codecov-commenter on February 14, 2023 at 01:04 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/957?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **72.35**% // Head: **72.35**% // No change to project coverage :thumbsup:
> Coverage data is based on head [(`fea9997`)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/957?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`89bb482`)](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/89bb482e0410ddb3417b969d4b0379b08944535d?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch has no changes to coverable lines.

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master     #957   +/-   ##
=======================================
  Coverage   72.35%   72.35%           
=======================================
  Files         110      110           
  Lines        2615     2615           
  Branches      658      658           
=======================================
  Hits         1892     1892           
  Misses        723      723           
```



Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/957?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/957*
