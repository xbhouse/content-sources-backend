---
type: pull_request
number: 1067
title: "Include only installable advisories when remediating all"
state: merged
author: leSamo
created: 2023-05-30T14:06:01Z
updated: 2023-06-05T21:41:53Z
closed: 2023-05-30T16:01:35Z
merged: 2023-05-30T16:01:35Z
base_branch: master
head_branch: fix-remediate-applicable
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1067
---

# Pull Request #1067: Include only installable advisories when remediating all

**Author**: @leSamo
**Created**: May 30, 2023 at 02:06 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `fix-remediate-applicable`

## Description

https://issues.redhat.com/browse/SPM-2069

1. Go to system list page
2. Find a system with a template, that has some applicable (non-installable) advisories
3. In row actions menu select "Apply all applicable advisories"
4. Observe how many issues are selected -- previously it selected applicable and installable advisories, now it should only select installable 

---

## Discussion

### Comment by @mkholjuraev on June 05, 2023 at 09:41 PM UTC

:tada: This PR is included in version 1.63.3 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.63.3)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.63.3)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @mkholjuraev - Approved on May 30, 2023 at 02:09 PM UTC

LGTM! Happy to see this

### Review by @psegedy - Commented on May 30, 2023 at 04:25 PM UTC

### Review by @leSamo - Commented on May 30, 2023 at 04:31 PM UTC

### Review by @psegedy - Commented on May 30, 2023 at 04:52 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1067*
