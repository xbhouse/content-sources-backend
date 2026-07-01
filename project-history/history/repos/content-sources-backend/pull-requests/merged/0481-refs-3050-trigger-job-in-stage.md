---
type: pull_request
number: 481
title: "Refs 3050: trigger job in stage"
state: merged
author: jlsherrill
created: 2023-11-21T19:54:18Z
updated: 2023-11-21T20:26:14Z
closed: 2023-11-21T20:26:14Z
merged: 2023-11-21T20:26:13Z
base_branch: main
head_branch: 3050_3
labels: []
url: https://github.com/content-services/content-sources-backend/pull/481
---

# Pull Request #481: Refs 3050: trigger job in stage

**Author**: @jlsherrill
**Created**: November 21, 2023 at 07:54 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `3050_3`

## Description

## Summary

I believe this would cause it to be triggered in stage

## Testing steps

## Checklist

- [x] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on November 21, 2023 at 08:00 PM UTC

https://issues.redhat.com/browse/HMS-3050

### Comment by @rverdile on November 21, 2023 at 08:21 PM UTC

is this file we would append to every time we wanted to trigger a job?

### Comment by @jlsherrill on November 21, 2023 at 08:23 PM UTC

yes, i believe so.  i modeled this after https://github.com/RedHatInsights/insights-host-inventory/blob/470f2264ed7be2f5b028f880a7aae2790ca9e660/docs/running_special_job.md?plain=1#L29

---

## Reviews

### Review by @Andrewgdewar - Approved on November 21, 2023 at 08:11 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/481*
