---
type: pull_request
number: 1009
title: "HMS-5530: add a helper for cleaning up test resources"
state: merged
author: dominikvagner
created: 2025-03-06T15:52:20Z
updated: 2025-03-24T17:14:48Z
closed: 2025-03-24T17:14:48Z
merged: 2025-03-24T17:14:48Z
base_branch: main
head_branch: 5530
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/1009
---

# Pull Request #1009: HMS-5530: add a helper for cleaning up test resources

**Author**: @dominikvagner
**Created**: March 06, 2025 at 03:52 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `5530`

## Description

## Summary
This adds new helpers and a fixture for easier cleaning up of test resources and better test idem-potency handling. And also "migrates" all existing tests that create test data to use this approach. 🧹😌 

## Testing steps
- All tests should pass repeatedly and even with existing data that would cause conflicts on resource creation. (Conflicts on repo name/URL or template name) 🔁 


---

## Discussion

### Comment by @jlsherrill on March 06, 2025 at 04:00 PM UTC

https://issues.redhat.com/browse/HMS-5530

---

## Reviews

### Review by @jlsherrill - Commented on March 12, 2025 at 03:09 PM UTC

### Review by @dominikvagner - Commented on March 12, 2025 at 03:38 PM UTC

### Review by @xbhouse - Commented on March 12, 2025 at 08:35 PM UTC

### Review by @dominikvagner - Commented on March 19, 2025 at 02:30 PM UTC

### Review by @xbhouse - Approved on March 20, 2025 at 01:58 PM UTC

looks great and tests are green 🤩 awesome work!!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1009*
