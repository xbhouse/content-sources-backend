---
type: pull_request
number: 1043
title: "HMS-5633: test automated snapshot cleanup with DB intervention"
state: merged
author: dominikvagner
created: 2025-03-25T09:05:56Z
updated: 2025-03-26T15:13:17Z
closed: 2025-03-26T15:13:17Z
merged: 2025-03-26T15:13:17Z
base_branch: main
head_branch: 5633_2
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/1043
---

# Pull Request #1043: HMS-5633: test automated snapshot cleanup with DB intervention

**Author**: @dominikvagner
**Created**: March 25, 2025 at 09:05 AM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `5633_2`

## Description

## Summary
This PR adds a new test, that directly interacts with the DB to replicate a scenario needed for verifying the functionality of automatic snapshot cleanup (which is a part of the `process-repos`). 
To do this a new DB helper fixture was added to our PW tests that enables DB intervention in tests and also a new standalone `snapshot-cleanup` command was added to overcome queue bog downs in CI or local env while testing.

## Testing steps
1. Tests pass and the new snapshot cleanup test verifies the snapshot cleanup functionality in a similar way to [this PR](https://github.com/content-services/content-sources-backend/pull/988)s testing steps.

---

## Discussion

### Comment by @jlsherrill on March 25, 2025 at 09:30 AM UTC

https://issues.redhat.com/browse/HMS-5633

---

## Reviews

### Review by @marusak - Approved on March 26, 2025 at 01:05 PM UTC

Awesome, thanks!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1043*
