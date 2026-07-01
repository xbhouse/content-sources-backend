---
type: pull_request
number: 570
title: "Fixes 3491: bootstrap candlepin connection"
state: merged
author: jlsherrill
created: 2024-02-09T17:56:03Z
updated: 2024-02-28T15:08:34Z
closed: 2024-02-28T15:07:55Z
merged: 2024-02-28T15:07:55Z
base_branch: main
head_branch: cp_bootstrap
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/570
---

# Pull Request #570: Fixes 3491: bootstrap candlepin connection

**Author**: @jlsherrill
**Created**: February 09, 2024 at 05:56 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `cp_bootstrap`

## Description

## Summary

This adds an initial candlepin client bootstrap process and adds a couple items to our compose-up process:
* creates an org
* imports a manifest

The thought is that in stage/prod we will use the existing orgs that match the user's org
In dev, all actions will use this created 'devel org'
This will be controlled via the devel_org boolean config 

## Testing steps

Checkout this pr
update your config.yaml with the new example values.

```go get ./...```
```make compose-down compose-up```
``` go run cmd/candlepin/main.go list-contents```  # list the contents of the devel org



## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on February 09, 2024 at 06:00 PM UTC

https://issues.redhat.com/browse/HMS-3491

---

## Reviews

### Review by @rverdile - Commented on February 15, 2024 at 08:37 PM UTC

### Review by @rverdile - Commented on February 21, 2024 at 07:28 PM UTC

Looks good overall. I tested it and both commands work. Only thing I found was the one small comment I left.

### Review by @jlsherrill - Commented on February 21, 2024 at 08:39 PM UTC

### Review by @rverdile - Commented on February 22, 2024 at 04:06 PM UTC

### Review by @jlsherrill - Commented on February 22, 2024 at 05:01 PM UTC

### Review by @rverdile - Approved on February 22, 2024 at 06:13 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/570*
