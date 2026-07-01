---
type: pull_request
number: 758
title: "Fixes 4413: create aggregate function in migration"
state: merged
author: jlsherrill
created: 2024-08-01T15:38:49Z
updated: 2024-08-01T20:00:32Z
closed: 2024-08-01T19:43:19Z
merged: 2024-08-01T19:43:19Z
base_branch: main
head_branch: 4413
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/758
---

# Pull Request #758: Fixes 4413: create aggregate function in migration

**Author**: @jlsherrill
**Created**: August 01, 2024 at 03:38 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `4413`

## Description

## Summary

Moves the aggregate function creation to a migration

## Testing steps

tests pass, this search seems well tested and should either work or not work.

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on August 01, 2024 at 04:00 PM UTC

https://issues.redhat.com/browse/HMS-4413

### Comment by @mayurilahane on August 01, 2024 at 06:29 PM UTC

/retest

---

## Reviews

### Review by @xbhouse - Approved on August 01, 2024 at 05:22 PM UTC

ack!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/758*
