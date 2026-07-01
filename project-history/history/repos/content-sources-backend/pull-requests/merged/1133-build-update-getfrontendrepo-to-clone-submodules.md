---
type: pull_request
number: 1133
title: "Build: update getFrontendRepo to clone submodules"
state: merged
author: dominikvagner
created: 2025-06-24T15:00:34Z
updated: 2025-06-25T13:09:39Z
closed: 2025-06-25T13:09:39Z
merged: 2025-06-25T13:09:39Z
base_branch: main
head_branch: fix-frontend-checkout
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1133
---

# Pull Request #1133: Build: update getFrontendRepo to clone submodules

**Author**: @dominikvagner
**Created**: June 24, 2025 at 03:00 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `fix-frontend-checkout`

## Description

## Summary
This updates the `getFrontendRepo` workflow helper script to clone the frontend repo with submodules. Needed for cloning test utils.
This also switches the proxy to use the CI specific routes config, which is needed for the proxy to work correctly inside DinD. And makes the Playwright action use the updated user secrets.

## Testing steps
Playwright tests pass in CI.


---

## Discussion

### Comment by @dominikvagner on June 25, 2025 at 09:36 AM UTC

/retest

### Comment by @jlsherrill on June 25, 2025 at 12:11 PM UTC

/retest

---

## Reviews

### Review by @jlsherrill - Approved on June 25, 2025 at 12:17 PM UTC

ACK, konflux failure is unrelated and being fixed by another team

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1133*
