---
type: pull_request
number: 1463
title: "HMS-10539: run remove-custom-epel as cron"
state: merged
author: rverdile
created: 2026-04-14T19:51:27Z
updated: 2026-04-15T14:21:28Z
closed: 2026-04-15T14:21:28Z
merged: 2026-04-15T14:21:28Z
base_branch: main
head_branch: cron-for-epel
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1463
---

# Pull Request #1463: HMS-10539: run remove-custom-epel as cron

**Author**: @rverdile
**Created**: April 14, 2026 at 07:51 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `cron-for-epel`

## Description

## Summary
Spreads removing the remaining custom epel repos to 8x a day, 50 at a time, so the tasking system does not get a huge backlog of deletion tasks.

## Testing steps
1. Create two custom epel repos
2. You can run `cmd/jobs/main.go remove-custom-epel-repos 2` just to make sure the batching works
3. This job has otherwise already run in stage/prod to success



---

## Discussion

### Comment by @xbhouse on April 14, 2026 at 08:00 PM UTC

https://issues.redhat.com/browse/HMS-10539

---

## Reviews

### Review by @TenSt - Approved on April 15, 2026 at 08:04 AM UTC

lgtm

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1463*
