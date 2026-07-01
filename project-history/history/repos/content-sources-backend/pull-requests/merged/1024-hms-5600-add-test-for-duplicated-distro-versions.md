---
type: pull_request
number: 1024
title: "HMS-5600: Add test for duplicated distro versions"
state: merged
author: dominikvagner
created: 2025-03-12T15:43:26Z
updated: 2025-04-09T08:47:55Z
closed: 2025-04-09T08:47:55Z
merged: 2025-04-09T08:47:55Z
base_branch: main
head_branch: 5600
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/1024
---

# Pull Request #1024: HMS-5600: Add test for duplicated distro versions

**Author**: @dominikvagner
**Created**: March 12, 2025 at 03:43 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `5600`

## Description

## Summary
This adds a new PW test (migrates it from IQE) which tests that repository creation will ignore duplicate distribution versions.

## Testing steps
- All tests pass as expected.
- The new migrated test verifies the same functionality as the one in IQE.

---

## Discussion

### Comment by @jlsherrill on March 19, 2025 at 10:00 AM UTC

https://issues.redhat.com/browse/HMS-5600

### Comment by @rverdile on March 19, 2025 at 08:01 PM UTC

looks good but needs rebase!

---

## Reviews

### Review by @rverdile - Approved on March 20, 2025 at 05:42 PM UTC

lgtm!

### Review by @swadeley - Commented on March 31, 2025 at 01:51 PM UTC

### Review by @dominikvagner - Commented on April 08, 2025 at 09:51 AM UTC

### Review by @swadeley - Commented on April 09, 2025 at 08:46 AM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1024*
