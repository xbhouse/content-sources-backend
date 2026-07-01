---
type: pull_request
number: 1056
title: "HMS-5798: move helpers and fixtures to shared dir"
state: merged
author: dominikvagner
created: 2025-04-02T11:26:54Z
updated: 2025-04-25T08:30:38Z
closed: 2025-04-25T08:30:38Z
merged: 2025-04-25T08:30:38Z
base_branch: main
head_branch: test-shared-utils
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/1056
---

# Pull Request #1056: HMS-5798: move helpers and fixtures to shared dir

**Author**: @dominikvagner
**Created**: April 02, 2025 at 11:26 AM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `test-shared-utils`

## Description

## Summary
This moves and in some cases slightly polishes helpers, fixtures and the API client to a new directory which is going to contain any shared test utilities between our projects (through git submodules and sparse checkout).
Along with this a new `package.json` file was added to the to be shared directory which should contain any dependencies required by these utilities. This shared utilities library should then be included in `package.json` files of repos consuming it and also to their TS configs for compilation.

## Testing steps
All tests should pass as they did before and helpers/fixtures work in the same way.


---

## Discussion

### Comment by @jlsherrill on April 02, 2025 at 11:30 AM UTC

https://issues.redhat.com/browse/HMS-5798

---

## Reviews

### Review by @xbhouse - Approved on April 24, 2025 at 08:42 PM UTC

this looks good! tests pass locally (except for snapshot cleanup but that also fails for me on main so not related to this 😃)

### Review by @marusak - Commented on April 25, 2025 at 07:23 AM UTC

### Review by @marusak - Commented on April 25, 2025 at 07:24 AM UTC

### Review by @marusak - Approved on April 25, 2025 at 07:24 AM UTC

Thank you! I have only two questions, but should be all good :) 

### Review by @dominikvagner - Commented on April 25, 2025 at 08:00 AM UTC

### Review by @dominikvagner - Commented on April 25, 2025 at 08:02 AM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1056*
