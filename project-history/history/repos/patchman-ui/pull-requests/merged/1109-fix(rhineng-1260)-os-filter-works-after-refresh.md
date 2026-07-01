---
type: pull_request
number: 1109
title: "fix(RHINENG-1260): OS filter works after refresh"
state: merged
author: AsToNlele
created: 2023-08-17T11:18:20Z
updated: 2023-08-22T08:35:42Z
closed: 2023-08-22T08:16:16Z
merged: 2023-08-22T08:16:16Z
base_branch: master
head_branch: RHINENG-1260
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1109
---

# Pull Request #1109: fix(RHINENG-1260): OS filter works after refresh

**Author**: @AsToNlele
**Created**: August 17, 2023 at 11:18 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `RHINENG-1260`

## Description

# Description

Associated Jira ticket: [RHINENG-1260](https://issues.redhat.com/browse/RHINENG-1260)

When reloading the page the `currentValue` is loaded from the URL and is an `array`. But when you apply an OS filter through the UI it is an `string`. I've added a check for an array so the page doesn't crash.

# How to test the PR

In the systems table apply an OS filter and refresh, the load should be successful.


# Checklist:

- [x] The commit message has the Jira ticket linked
- [x] PR has a short description
- [ ] Screenshots before and after the change are added
- [ ] Tests for the changes have been added
- [ ] README.md is updated if necessary
- [ ] Needs additional dependent work


---

## Discussion

### Comment by @AsToNlele on August 21, 2023 at 02:35 PM UTC

Thank you! Removed the comment and included some tests :smile: 

### Comment by @AsToNlele on August 22, 2023 at 06:32 AM UTC

/retest

### Comment by @mkholjuraev on August 22, 2023 at 08:35 AM UTC

:tada: This PR is included in version 1.63.12 :tada:

The release is available on [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.63.12)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @bastilian - Approved on August 21, 2023 at 11:19 AM UTC

LGTM! Thank you @AsToNlele!

Side note: The comment does feel a bit redundant. :) It states the already obvious, but doesn't explain why it should not split in case it is an array, or why it can be an array in some cases. However, these things could also be described in a test and it would not need any documentation in the code, but the tests serve as documentation.

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1109*
