---
type: pull_request
number: 986
title: "fix(SPM-1885): pass auth token to webworker"
state: merged
author: mkholjuraev
created: 2023-03-10T08:25:48Z
updated: 2024-04-03T09:22:36Z
closed: 2023-03-10T09:37:00Z
merged: 2023-03-10T09:37:00Z
base_branch: master
head_branch: fix-remediation
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/986
---

# Pull Request #986: fix(SPM-1885): pass auth token to webworker

**Author**: @mkholjuraev
**Created**: March 10, 2023 at 08:25 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `fix-remediation`

## Description

# Description

Associated Jira ticket: https://issues.redhat.com/browse/SPM-1885

Fixes failing remediation by passing auth token explicitly. After platform decomissioned auth cooki, we have to pass it to webworker.

# How to test the PR

Please make sure remediation works as expected, the remediation gets indeed created.

# Before the change


# After the change


# Dependent work link


# Checklist:

- [x] The commit message has the Jira ticket linked
- [x] PR has a short description
- [ ] Screenshots before and after the change are added
- [ ] Tests for the changes have been added
- [ ] README.md is updated if necessary
- [ ] Needs additional dependent work


---

## Discussion

### Comment by @mkholjuraev on March 10, 2023 at 09:53 AM UTC

:tada: This PR is included in version 1.62.2 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.62.2)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.62.2)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @gkarat - Approved on March 10, 2023 at 09:33 AM UTC

Worked for me on a couple of systems/remediations (ran under stage-beta env). LGTM!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/986*
