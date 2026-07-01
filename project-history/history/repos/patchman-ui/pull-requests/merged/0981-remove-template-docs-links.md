---
type: pull_request
number: 981
title: "Remove template docs links"
state: merged
author: leSamo
created: 2023-03-07T23:05:16Z
updated: 2023-05-08T18:16:58Z
closed: 2023-03-13T17:20:24Z
merged: 2023-03-13T17:20:24Z
base_branch: master
head_branch: remove-doc-links
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/981
---

# Pull Request #981: Remove template docs links

**Author**: @leSamo
**Created**: March 07, 2023 at 11:05 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `remove-doc-links`

## Description

As per PM's request (https://redhat-internal.slack.com/archives/CPS5FH07Q/p1678202549907769?thread_ts=1678119878.115079&cid=CPS5FH07Q) we are, for now, removing all documentation link related to templates.

These places include:
- Template list page empty state (displayed when user has 0 templates)
- Template list page title tooltip
- Template creation/editing wizard just below the wizard title

---

## Discussion

### Comment by @mkholjuraev on March 09, 2023 at 08:44 AM UTC

Maybe we can use a feature flag?

### Comment by @leSamo on March 09, 2023 at 08:48 AM UTC

> Maybe we can use a feature flag?

We theoretically could, but after the link has been updated we wouldn't need the feature flag anymore and would need to remove some code anyway. We would need to add feature flag hooks and I find that more intrusive than just commenting the code out and then back in.

### Comment by @mkholjuraev on March 28, 2023 at 08:57 AM UTC

:tada: This PR is included in version 1.62.4 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.62.4)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.62.4)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @bastilian - Approved on March 13, 2023 at 05:11 PM UTC

LGTM! Thank you @leSamo!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/981*
