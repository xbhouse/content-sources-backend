---
type: pull_request
number: 1252
title: "HMS-9279: wait for candlepin to be healthy before running init"
state: merged
author: xbhouse
created: 2025-10-24T16:09:57Z
updated: 2025-10-27T13:38:14Z
closed: 2025-10-27T13:38:14Z
merged: 2025-10-27T13:38:14Z
base_branch: main
head_branch: 9279
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1252
---

# Pull Request #1252: HMS-9279: wait for candlepin to be healthy before running init

**Author**: @xbhouse
**Created**: October 24, 2025 at 04:09 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `9279`

## Description

## Summary

Adds a healthcheck to candlepin and waits for the pod to be healthy before running init

## Testing steps

1. Run `make compose-clean`
2. When your laptop isn't charging or is in power-saver mode, run `make compose-up`
3. You shouldn't see candlepin time out


---

## Discussion

### Comment by @xbhouse on October 24, 2025 at 04:30 PM UTC

https://issues.redhat.com/browse/HMS-9279

---

## Reviews

### Review by @jlsherrill - Commented on October 24, 2025 at 05:09 PM UTC

### Review by @xbhouse - Commented on October 24, 2025 at 05:22 PM UTC

### Review by @jlsherrill - Commented on October 27, 2025 at 12:18 PM UTC

### Review by @jlsherrill - Approved on October 27, 2025 at 12:19 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1252*
