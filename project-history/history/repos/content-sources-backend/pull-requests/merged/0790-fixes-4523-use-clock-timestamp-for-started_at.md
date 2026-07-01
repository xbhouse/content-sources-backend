---
type: pull_request
number: 790
title: "Fixes 4523: use clock timestamp for started_at"
state: merged
author: jlsherrill
created: 2024-08-28T13:02:11Z
updated: 2024-08-28T18:30:25Z
closed: 2024-08-28T18:21:39Z
merged: 2024-08-28T18:21:39Z
base_branch: main
head_branch: 4523
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/790
---

# Pull Request #790: Fixes 4523: use clock timestamp for started_at

**Author**: @jlsherrill
**Created**: August 28, 2024 at 01:02 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `4523`

## Description

## Summary

statement_timestamp() calculates the timestamp of when the statement starts.  I suspect that a new task is being added while this query is executing in rare cases (but don't actually have much evidence for this).  By using clock_timestamp(), we could (in theory) ensure that the started_at is always set after the queued_at is set.

We'll probably have to commit this and monitor errors for a couple weeks to see if its still occuring.

## Testing steps

tests pass

---

## Discussion

### Comment by @jlsherrill on August 28, 2024 at 01:30 PM UTC

https://issues.redhat.com/browse/HMS-4523

### Comment by @rverdile on August 28, 2024 at 03:40 PM UTC

/retest

### Comment by @jlsherrill on August 28, 2024 at 04:47 PM UTC

[test]

### Comment by @jlsherrill on August 28, 2024 at 06:22 PM UTC

oh this didn't actually need or warrent qe beyond passing tests, will update card

---

## Reviews

### Review by @rverdile - Commented on August 28, 2024 at 02:27 PM UTC

this sounds like it could be the cause to me. could you change the other uses of statement_timestamp() to clock_timestamp()? clock_timestamp() seems like the better function to use

### Review by @rverdile - Approved on August 28, 2024 at 05:39 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/790*
