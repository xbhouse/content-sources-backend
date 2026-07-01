---
type: pull_request
number: 647
title: "Fixes 3793: add make command for mockery"
state: merged
author: xbhouse
created: 2024-04-24T20:04:01Z
updated: 2024-04-30T13:29:49Z
closed: 2024-04-30T13:29:49Z
merged: 2024-04-30T13:29:49Z
base_branch: main
head_branch: add-make-mockery
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/647
---

# Pull Request #647: Fixes 3793: add make command for mockery

**Author**: @xbhouse
**Created**: April 24, 2024 at 08:04 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `add-make-mockery`

## Description

## Summary

Adds make command to install mockery and regenerate mocks

## Testing steps

- Run `make mock`. This should install the latest version of mockery if it isn't already installed in the release directory and regenerate the mocks we have

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on April 24, 2024 at 08:30 PM UTC

https://issues.redhat.com/browse/HMS-3793

### Comment by @swadeley on April 26, 2024 at 06:38 PM UTC

/retest

---

## Reviews

### Review by @jlsherrill - Commented on April 26, 2024 at 03:52 PM UTC

### Review by @xbhouse - Commented on April 26, 2024 at 06:07 PM UTC

### Review by @jlsherrill - Commented on April 29, 2024 at 12:55 PM UTC

### Review by @xbhouse - Commented on April 29, 2024 at 02:51 PM UTC

### Review by @jlsherrill - Commented on April 29, 2024 at 02:54 PM UTC

### Review by @jlsherrill - Commented on April 29, 2024 at 02:55 PM UTC

### Review by @jlsherrill - Approved on April 29, 2024 at 08:27 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/647*
