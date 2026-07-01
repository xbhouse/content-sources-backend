---
type: pull_request
number: 1398
title: "Build: fix jobs"
state: merged
author: xbhouse
created: 2026-02-25T19:44:08Z
updated: 2026-02-25T20:04:49Z
closed: 2026-02-25T20:04:49Z
merged: 2026-02-25T20:04:49Z
base_branch: main
head_branch: fix-jobs
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1398
---

# Pull Request #1398: Build: fix jobs

**Author**: @xbhouse
**Created**: February 25, 2026 at 07:44 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `fix-jobs`

## Description

## Summary

- Shortens job name for `migrate-templates-to-shared-epel`
  - Job labels (like `content-sources-backend-migrate-templates-to-shared-epel-fichspi`) need to be no more than 63 characters :exploding_head: so this job was failing to create
- Removes an old job definition from jobs-stage.yaml

## Testing steps



---

## Reviews

### Review by @rverdile - Approved on February 25, 2026 at 08:03 PM UTC

ack!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1398*
