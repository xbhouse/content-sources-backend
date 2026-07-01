---
type: pull_request
number: 252
title: "Refs 1602: add floorist query for metrics gathering"
state: merged
author: jlsherrill
created: 2023-04-12T17:14:14Z
updated: 2023-04-17T17:46:59Z
closed: 2023-04-17T17:46:56Z
merged: 2023-04-17T17:46:56Z
base_branch: main
head_branch: 1602
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/252
---

# Pull Request #252: Refs 1602: add floorist query for metrics gathering

**Author**: @jlsherrill
**Created**: April 12, 2023 at 05:14 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `1602`

## Description

## Summary
This defines a query that gets run daily and collected into an S3 bucket for later analysis. 

## Testing steps
Manually make sure the query is correct


---

## Discussion

### Comment by @jlsherrill on April 12, 2023 at 05:30 PM UTC

https://issues.redhat.com/browse/HMS-1602

### Comment by @jlsherrill on April 12, 2023 at 05:30 PM UTC

⚠️ This task currently requires qe-approval, but this PR does not fully resolve the task.  Please contact QE to determine appropriate testing required.

---

## Reviews

### Review by @rverdile - Approved on April 14, 2023 at 05:00 PM UTC

Tested the query, looks like it's working as expected. I'm assuming the other stuff can't really be tested, but I'll add that I don't see any typos.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/252*
