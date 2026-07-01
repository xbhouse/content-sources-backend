---
type: pull_request
number: 762
title: "Build: fix gh action docker-compose up"
state: merged
author: jlsherrill
created: 2024-08-05T12:09:01Z
updated: 2024-08-05T13:24:28Z
closed: 2024-08-05T13:24:24Z
merged: 2024-08-05T13:24:24Z
base_branch: main
head_branch: test_test
labels: []
url: https://github.com/content-services/content-sources-backend/pull/762
---

# Pull Request #762: Build: fix gh action docker-compose up

**Author**: @jlsherrill
**Created**: August 05, 2024 at 12:09 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `test_test`

## Description

## Summary

for some unknown reason specifying the full path to the docker compose no longer works.  Something must have changed in the way GH actions run.  Switching to the newer action for docker-compose, and specifying CWD seems to fix it.

## Testing steps
Tests pass!


---

## Discussion

### Comment by @jlsherrill on August 05, 2024 at 01:04 PM UTC

/retest

---

## Reviews

### Review by @xbhouse - Commented on August 05, 2024 at 12:45 PM UTC

### Review by @jlsherrill - Commented on August 05, 2024 at 12:48 PM UTC

### Review by @xbhouse - Approved on August 05, 2024 at 12:53 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/762*
