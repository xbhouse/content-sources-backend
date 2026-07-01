---
type: pull_request
number: 462
title: "Fixes 3019: No need for pr check to run in venv ephemeral"
state: merged
author: swadeley
created: 2023-11-10T07:48:19Z
updated: 2025-09-08T07:30:29Z
closed: 2023-11-14T01:36:26Z
merged: 2023-11-14T01:36:26Z
base_branch: main
head_branch: swadeley/test_pr_test_run
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/462
---

# Pull Request #462: Fixes 3019: No need for pr check to run in venv ephemeral

**Author**: @swadeley
**Created**: November 10, 2023 at 07:48 AM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `swadeley/test_pr_test_run`

## Description

## Summary
The pr check does not need to run in ephemeral env, the main thing is not to run the introspection test as that works in stage_proxy and prod only

## Testing steps

Pr test run passes.








---

## Discussion

### Comment by @swadeley on November 10, 2023 at 03:49 PM UTC

/retest

### Comment by @swadeley on November 13, 2023 at 12:14 AM UTC

/retest

### Comment by @jlsherrill on November 13, 2023 at 03:30 AM UTC

https://issues.redhat.com/browse/HMS-3019

---

## Reviews

### Review by @swadeley - Commented on November 13, 2023 at 01:09 AM UTC

lgtm

### Review by @jlsherrill - Approved on November 13, 2023 at 02:48 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/462*
