---
type: pull_request
number: 770
title: "Fixes 4561: log/return domain on pulp task failures"
state: merged
author: jlsherrill
created: 2024-08-09T17:03:35Z
updated: 2024-08-12T20:00:36Z
closed: 2024-08-12T20:00:33Z
merged: 2024-08-12T20:00:33Z
base_branch: main
head_branch: 4561
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/770
---

# Pull Request #770: Fixes 4561: log/return domain on pulp task failures

**Author**: @jlsherrill
**Created**: August 09, 2024 at 05:03 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `4561`

## Description

## Summary

added the domain to logging/errors of failed pulp tasks

## Testing steps


create a repo with snapshotting pointing at a URL that is nonsense.  Monitor the logs:

```
error="domain ed88dfc6 had Pulp Task error 'traceback:   File...
```


---

## Discussion

### Comment by @jlsherrill on August 09, 2024 at 05:30 PM UTC

https://issues.redhat.com/browse/HMS-4561

### Comment by @mayurilahane on August 12, 2024 at 05:57 PM UTC

/retest

### Comment by @jlsherrill on August 12, 2024 at 07:57 PM UTC

oh, this doesn't need qe

---

## Reviews

### Review by @xbhouse - Approved on August 09, 2024 at 06:50 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/770*
