---
type: pull_request
number: 973
title: "Build: use full path for compose files"
state: merged
author: jlsherrill
created: 2025-02-11T16:17:16Z
updated: 2025-03-12T16:21:14Z
closed: 2025-03-12T16:21:14Z
merged: 2025-03-12T16:21:14Z
base_branch: main
head_branch: podman_fix
labels: []
url: https://github.com/content-services/content-sources-backend/pull/973
---

# Pull Request #973: Build: use full path for compose files

**Author**: @jlsherrill
**Created**: February 11, 2025 at 04:17 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `podman_fix`

## Description

## Summary

specifies full path to compose files to work around podman-compose bug

## Testing steps

make compose-clean compose-up 



---

## Reviews

### Review by @Andrewgdewar - Approved on March 12, 2025 at 03:00 PM UTC

Still works on mac! Ack from me!

### Review by @dominikvagner - Approved on March 12, 2025 at 03:03 PM UTC

awesome, works with updated `podman-compose` _(1.3.0-2)_ on fedora _(41)_ 🎉 
thanks! 🫶 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/973*
