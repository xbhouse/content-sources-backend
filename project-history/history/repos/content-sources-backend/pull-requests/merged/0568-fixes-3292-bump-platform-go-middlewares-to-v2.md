---
type: pull_request
number: 568
title: "Fixes 3292: Bump platform-go-middlewares to v2"
state: merged
author: Andrewgdewar
created: 2024-02-08T22:07:11Z
updated: 2024-02-28T17:00:32Z
closed: 2024-02-28T16:51:56Z
merged: 2024-02-28T16:51:56Z
base_branch: main
head_branch: HMS-3292
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/568
---

# Pull Request #568: Fixes 3292: Bump platform-go-middlewares to v2

**Author**: @Andrewgdewar
**Created**: February 08, 2024 at 10:07 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `HMS-3292`

## Description

## Summary

## Testing steps

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on February 08, 2024 at 10:30 PM UTC

https://issues.redhat.com/browse/HMS-3292

### Comment by @jlsherrill on February 12, 2024 at 07:13 PM UTC

overall looks good. i couldn't figure out how (if possible) to get service accounts in ephemeral, so once this is merged, we can test in stage

### Comment by @Andrewgdewar on February 13, 2024 at 07:16 PM UTC

> was there a mock here that needed to be updated? I see a bunch of file changes on the mock files, but it looks like they're all just updating the version in the header comment.

I thought there was one update in one file, I'll double check though, and remove all changes that aren't necessary.

### Comment by @jlsherrill on February 14, 2024 at 09:06 PM UTC

Can you try pushing this again to see if the snyk checks clear up?

### Comment by @jlsherrill on February 21, 2024 at 08:59 PM UTC

needs a rebase, that may clear up the sync checks

### Comment by @swadeley on February 27, 2024 at 03:41 PM UTC

/retest

### Comment by @swadeley on February 28, 2024 at 04:51 PM UTC


Hi, I deployed the backend image, API tests pass locally.



---

## Reviews

### Review by @Andrewgdewar - Commented on February 09, 2024 at 08:41 PM UTC

### Review by @jlsherrill - Commented on February 12, 2024 at 06:20 PM UTC

### Review by @rverdile - Commented on February 12, 2024 at 08:38 PM UTC

was there a mock here that needed to be updated? I see a bunch of file changes on the mock files, but it looks like they're all just updating the version in the header comment.

### Review by @Andrewgdewar - Commented on February 13, 2024 at 07:15 PM UTC

### Review by @jlsherrill - Approved on February 27, 2024 at 04:32 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/568*
