---
type: pull_request
number: 34
title: "increases default max xml size"
state: merged
author: xbhouse
created: 2025-06-17T15:00:08Z
updated: 2026-05-20T16:56:49Z
closed: 2025-06-18T13:58:56Z
merged: 2025-06-18T13:58:56Z
base_branch: master
head_branch: increase-max-xml-size
labels: []
url: https://github.com/content-services/yummy/pull/34
---

# Pull Request #34: increases default max xml size

**Author**: @xbhouse
**Created**: June 17, 2025 at 03:00 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `increase-max-xml-size`

## Description

Introspection of the RHEL9 arm baseos repo is failing with this error (`error introspecting https://cdn.redhat.com/content/dist/rhel9/9/aarch64/baseos/os/: XML syntax error on line 6371618: unexpected EOF`) because the uncompressed file is larger than the default we have set (about 3MB larger). This changes the default max xml size to be 1GB.

To test:

1. Replace yummy in the backend's go.mod with this PR (`replace github.com/content-services/yummy => /path/to/repo`)
2. Introspect the RHEL9 arm baseos repo: `./release/external-repos introspect --url https://cdn.redhat.com/content/dist/rhel9/9/aarch64/baseos/os/`. It should introspect successfully.   

---

## Reviews

### Review by @rverdile - Commented on June 18, 2025 at 01:40 PM UTC

### Review by @xbhouse - Commented on June 18, 2025 at 01:42 PM UTC

### Review by @rverdile - Approved on June 18, 2025 at 01:51 PM UTC

looks good!

---

*Archived from: https://github.com/content-services/yummy/pull/34*
