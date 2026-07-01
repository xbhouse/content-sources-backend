---
type: pull_request
number: 745
title: "Fixes 4431: detect env fetch 404 and use env type"
state: merged
author: jlsherrill
created: 2024-07-15T13:58:48Z
updated: 2024-07-15T15:12:59Z
closed: 2024-07-15T15:12:55Z
merged: 2024-07-15T15:12:55Z
base_branch: main
head_branch: 4431
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/745
---

# Pull Request #745: Fixes 4431: detect env fetch 404 and use env type

**Author**: @jlsherrill
**Created**: July 15, 2024 at 01:58 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `4431`

## Description

## Summary

Within the stage the 404 error is not the same as in dev,  `couldn't fetch environment: 404 Not Found:`  instead of `couldn't fetch environment: 404:`   So instead we should just check the response code.  

Also we should be setting an environment type, which can be an arbitrary string, set here to 'content-template'.

## Testing steps

Create a content template as normal, verify no errors

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on July 15, 2024 at 02:00 PM UTC

https://issues.redhat.com/browse/HMS-4431

---

## Reviews

### Review by @xbhouse - Approved on July 15, 2024 at 02:54 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/745*
