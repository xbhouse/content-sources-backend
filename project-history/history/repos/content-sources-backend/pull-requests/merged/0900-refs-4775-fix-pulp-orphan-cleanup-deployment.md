---
type: pull_request
number: 900
title: "Refs 4775: fix pulp-orphan-cleanup deployment"
state: merged
author: rverdile
created: 2024-11-20T18:36:16Z
updated: 2024-11-20T19:03:23Z
closed: 2024-11-20T19:03:19Z
merged: 2024-11-20T19:03:19Z
base_branch: main
head_branch: fix-orphan-cleanup
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/900
---

# Pull Request #900: Refs 4775: fix pulp-orphan-cleanup deployment

**Author**: @rverdile
**Created**: November 20, 2024 at 06:36 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `fix-orphan-cleanup`

## Description

## Summary
Orphan cleanup didn't run because the args were wrong. This fixes that. Changes the day to Thursday so I can check it again tomorrow.

## Testing steps
You could test this in ephemeral by changing the cron timer to every couple of minutes and checking the logs, but I already did :)

---

## Discussion

### Comment by @jlsherrill on November 20, 2024 at 07:00 PM UTC

https://issues.redhat.com/browse/HMS-4775

---

## Reviews

### Review by @xbhouse - Approved on November 20, 2024 at 06:56 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/900*
