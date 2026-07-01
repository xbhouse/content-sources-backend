---
type: pull_request
number: 1114
title: "HMS-8563: change to testing proxy for PW tests in actions"
state: merged
author: dominikvagner
created: 2025-05-19T12:46:19Z
updated: 2025-05-22T10:46:32Z
closed: 2025-05-22T10:46:32Z
merged: 2025-05-22T10:46:31Z
base_branch: main
head_branch: add-testing-proxy
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/1114
---

# Pull Request #1114: HMS-8563: change to testing proxy for PW tests in actions

**Author**: @dominikvagner
**Created**: May 19, 2025 at 12:46 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `add-testing-proxy`

## Description

## Summary
This PR changes the playwright action to use a new custom testing proxy, which should make tests faster and more stable.
Also increases the instance size and parallelizes the different long
running setups.

## Testing steps
Playwright test should pass without loading errors.



---

## Discussion

### Comment by @jlsherrill on May 19, 2025 at 01:00 PM UTC

https://issues.redhat.com/browse/HMS-8563

---

## Reviews

### Review by @marusak - Approved on May 20, 2025 at 11:37 AM UTC

Thank you!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1114*
