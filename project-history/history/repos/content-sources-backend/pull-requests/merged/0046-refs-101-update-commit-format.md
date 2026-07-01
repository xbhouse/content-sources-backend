---
type: pull_request
number: 46
title: "Refs 101: update commit format"
state: merged
author: jlsherrill
created: 2022-06-29T15:52:30Z
updated: 2022-07-19T20:20:36Z
closed: 2022-07-05T18:01:42Z
merged: 2022-07-05T18:01:42Z
base_branch: main
head_branch: 101
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/46
---

# Pull Request #46: Refs 101: update commit format

**Author**: @jlsherrill
**Created**: June 29, 2022 at 03:52 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `101`

## Description

Changes the commit format such that:
* A jira card is now required
* You can use Refs or Fixes to indicate if a PR
  is related to or fully fixes a card
* Adds a github action to validate PR Title

---

## Discussion

### Comment by @jlsherrill on June 29, 2022 at 05:29 PM UTC

updated to drop the '#' symbol, as that can cause confusion since github interprets '#num' as a link to a pr

### Comment by @jlsherrill on June 29, 2022 at 05:50 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-101

### Comment by @jlsherrill on June 30, 2022 at 08:24 PM UTC

This task currently requires qe-approval, but this PR does not fully resolve the task.  Please contact QE to determine appropriate testing required.

### Comment by @jlsherrill on June 30, 2022 at 08:28 PM UTC

⚠️ This task currently requires qe-approval, but this PR does not fully resolve the task.  Please contact QE to determine appropriate testing required.

### Comment by @jlsherrill on July 19, 2022 at 08:20 PM UTC

https://issues.redhat.com/browse/14709226

---

## Reviews

### Review by @swadeley - Approved on June 29, 2022 at 04:17 PM UTC

ACK

### Review by @rverdile - Approved on June 29, 2022 at 04:40 PM UTC

no objections from me

### Review by @mshriver - Approved on June 29, 2022 at 07:52 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/46*
