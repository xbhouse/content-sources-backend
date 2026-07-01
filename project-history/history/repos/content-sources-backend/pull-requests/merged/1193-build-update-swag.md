---
type: pull_request
number: 1193
title: "Build: update swag"
state: merged
author: rverdile
created: 2025-09-08T14:39:33Z
updated: 2025-09-08T16:44:26Z
closed: 2025-09-08T16:44:23Z
merged: 2025-09-08T16:44:23Z
base_branch: main
head_branch: no-u
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1193
---

# Pull Request #1193: Build: update swag

**Author**: @rverdile
**Created**: September 08, 2025 at 02:39 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `no-u`

## Description

## Summary
The "go get" command is unnecessary in this workflow because swag is installed by "make openapi". Also, the "-u" flag was causing the go.mod to be updated.

Also pins swag to be the latest version (it was previously 3 years old), and re-runs "make openapi" to get the updates


## Testing steps



---

## Reviews

### Review by @jlsherrill - Approved on September 08, 2025 at 04:23 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1193*
