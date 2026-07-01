---
type: pull_request
number: 428
title: "Refs 2776: capture pool stats"
state: closed
author: jlsherrill
created: 2023-10-12T18:25:18Z
updated: 2024-01-19T13:55:36Z
closed: 2024-01-19T13:55:36Z
base_branch: main
head_branch: 2776_pool
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/428
---

# Pull Request #428: Refs 2776: capture pool stats

**Author**: @jlsherrill
**Created**: October 12, 2023 at 06:25 PM UTC
**Status**: Closed
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `2776_pool`

## Description

## Summary

## Testing steps


---

## Discussion

### Comment by @jlsherrill on October 12, 2023 at 06:30 PM UTC

https://issues.redhat.com/browse/HMS-2776

### Comment by @jlsherrill on October 12, 2023 at 06:30 PM UTC

⚠️ This task currently requires qe-approval, but this PR does not fully resolve the task.  Please contact QE to determine appropriate testing required.

### Comment by @jlsherrill on October 17, 2023 at 05:50 PM UTC

I'm thinking i could also stick the logging in the pgx_pool_wrapper instead of here

### Comment by @Andrewgdewar on November 07, 2023 at 06:41 PM UTC

/retest

---

## Reviews

### Review by @rverdile - Commented on October 17, 2023 at 05:45 PM UTC

Since you said you were trying to think of a different way to do this, what if you added some fields to the logger? Kinda like here: https://github.com/content-services/content-sources-backend/blob/main/pkg/tasks/introspect.go#L50-L56

Maybe it could be done conditionally based on some config.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/428*
