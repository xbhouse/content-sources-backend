---
type: pull_request
number: 658
title: "Refs 3565: add pulp connectivity metric"
state: merged
author: jlsherrill
created: 2024-05-02T13:17:47Z
updated: 2024-05-08T17:59:49Z
closed: 2024-05-08T17:59:37Z
merged: 2024-05-08T17:59:37Z
base_branch: main
head_branch: 3565
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/658
---

# Pull Request #658: Refs 3565: add pulp connectivity metric

**Author**: @jlsherrill
**Created**: May 02, 2024 at 01:17 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `3565`

## Description

## Summary

adds a metric to represent the status of the pulp connection.  

## Testing steps

curl localhost:9000/metrics 

look for:
```
# HELP content_sources_PulpConnectivity Status of pulp connection
# TYPE content_sources_PulpConnectivity gauge
content_sources_PulpConnectivity 1
```

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on May 02, 2024 at 01:18 PM UTC

I'll rebase this after https://github.com/content-services/content-sources-backend/pull/645 goes in

### Comment by @jlsherrill on May 02, 2024 at 01:30 PM UTC

https://issues.redhat.com/browse/HMS-3565

### Comment by @jlsherrill on May 02, 2024 at 01:30 PM UTC

⚠️ This task currently requires qe-approval, but this PR does not fully resolve the task.  Please contact QE to determine appropriate testing required.

### Comment by @swadeley on May 08, 2024 at 10:37 AM UTC

@jlsherrill rebase please

---

## Reviews

### Review by @xbhouse - Approved on May 07, 2024 at 01:44 PM UTC

lgtm!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/658*
