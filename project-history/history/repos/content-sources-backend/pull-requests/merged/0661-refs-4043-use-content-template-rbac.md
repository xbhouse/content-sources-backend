---
type: pull_request
number: 661
title: "Refs 4043: use content-template rbac"
state: merged
author: jlsherrill
created: 2024-05-06T18:36:44Z
updated: 2024-05-14T12:44:55Z
closed: 2024-05-14T12:44:52Z
merged: 2024-05-14T12:44:52Z
base_branch: main
head_branch: 4043
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/661
---

# Pull Request #661: Refs 4043: use content-template rbac

**Author**: @jlsherrill
**Created**: May 06, 2024 at 06:36 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `4043`

## Description

## Summary

Adds content template permission to content template apis

## Testing steps
Probably needs to be tested in stage after merge

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on May 06, 2024 at 07:00 PM UTC

https://issues.redhat.com/browse/HMS-4043

### Comment by @xbhouse on May 13, 2024 at 01:55 PM UTC

this looks good, just needs some updated tests (unless i missed them) 

### Comment by @jlsherrill on May 13, 2024 at 06:11 PM UTC

added a small test, but we don't really test rbac for each handler, we more test the infrastructure around it.  Probably better suited for an api integration suite or iqe tests

---

## Reviews

### Review by @xbhouse - Commented on May 08, 2024 at 02:08 PM UTC

### Review by @xbhouse - Commented on May 08, 2024 at 02:12 PM UTC

### Review by @jlsherrill - Commented on May 08, 2024 at 05:13 PM UTC

### Review by @jlsherrill - Commented on May 08, 2024 at 05:13 PM UTC

### Review by @xbhouse - Approved on May 13, 2024 at 06:15 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/661*
