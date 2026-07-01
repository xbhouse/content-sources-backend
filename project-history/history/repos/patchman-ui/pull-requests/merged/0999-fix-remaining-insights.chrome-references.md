---
type: pull_request
number: 999
title: "fix: remaining insights.chrome references"
state: merged
author: johnsonm325
created: 2023-03-23T20:50:42Z
updated: 2023-03-28T08:57:34Z
closed: 2023-03-24T10:18:10Z
merged: 2023-03-24T10:18:10Z
base_branch: master
head_branch: fix-rest-insights-dot-chrome
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/999
---

# Pull Request #999: fix: remaining insights.chrome references

**Author**: @johnsonm325
**Created**: March 23, 2023 at 08:50 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `fix-rest-insights-dot-chrome`

## Description

# Description

Associated Jira ticket: # (issue)

Please include a summary of the change, what this fixes/creates/improves.


# How to test the PR

Please include steps to test your PR.

# Before the change


# After the change


# Dependent work link


# Checklist:

- [ ] The commit message has the Jira ticket linked
- [ ] PR has a short description
- [ ] Screenshots before and after the change are added
- [ ] Tests for the changes have been added
- [ ] README.md is updated if necessary
- [ ] Needs additional dependent work


---

## Discussion

### Comment by @mkholjuraev on March 24, 2023 at 09:53 AM UTC

One more commit is coming...

### Comment by @Fewwy on March 24, 2023 at 10:14 AM UTC

LGTM, tried it locally and haven't found any insights.chrome warnings.

### Comment by @mkholjuraev on March 28, 2023 at 08:56 AM UTC

:tada: This PR is included in version 1.62.4 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.62.4)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.62.4)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @gkarat - Commented on March 24, 2023 at 09:44 AM UTC

there is one place where we still keep insights.chrome call: https://github.com/johnsonm325/patchman-ui/blob/fix-rest-insights-dot-chrome/src/Utilities/constants.js#L35 (gonna fix it and commit to this PR)

### Review by @mkholjuraev - Approved on March 24, 2023 at 10:02 AM UTC

LGTM! tested locally

### Review by @Fewwy - Approved on March 24, 2023 at 10:14 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/999*
