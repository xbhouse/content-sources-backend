---
type: pull_request
number: 1151
title: "Build: pin the downloaded mockery version to v2"
state: merged
author: dominikvagner
created: 2025-07-23T15:00:59Z
updated: 2025-07-23T16:33:23Z
closed: 2025-07-23T16:33:23Z
merged: 2025-07-23T16:33:23Z
base_branch: main
head_branch: pin-mockery-version
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1151
---

# Pull Request #1151: Build: pin the downloaded mockery version to v2

**Author**: @dominikvagner
**Created**: July 23, 2025 at 03:00 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `pin-mockery-version`

## Description

## Summary
This pins the version of the downloaded mockery binary to `v2.*` as our config isn't v3 compatible. This should be reverted/removed when we or if we update the config.

Task for updating our config was created here [[HMS-8889](https://issues.redhat.com/browse/HMS-8889)].

## Testing steps
- When you remove your currently downloaded version of the mockery tool (`release/mockery`), the `make mock` command should download the latest v2 version and the mocks should be generated successfully.


---

## Reviews

### Review by @xbhouse - Approved on July 23, 2025 at 03:18 PM UTC

lgtm! thanks!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1151*
