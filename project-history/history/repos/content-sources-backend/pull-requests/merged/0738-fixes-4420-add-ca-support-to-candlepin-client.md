---
type: pull_request
number: 738
title: "Fixes 4420: add ca support to candlepin client"
state: merged
author: jlsherrill
created: 2024-07-10T15:11:10Z
updated: 2024-07-10T17:26:50Z
closed: 2024-07-10T17:26:45Z
merged: 2024-07-10T17:26:45Z
base_branch: main
head_branch: 4420
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/738
---

# Pull Request #738: Fixes 4420: add ca support to candlepin client

**Author**: @jlsherrill
**Created**: July 10, 2024 at 03:11 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `4420`

## Description

## Summary

support a custom ca cert in teh candlepin connection

## Testing steps

I would just do a code review on this.  I was able to see this work using custom certs, but it was a very manual process of loading files from the filesystem into the config (since the config takes the file contents, not file names).   We should see it fix the current issue in stage quickly

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on July 10, 2024 at 03:30 PM UTC

https://issues.redhat.com/browse/HMS-4420

### Comment by @jlsherrill on July 10, 2024 at 05:03 PM UTC

/retest

---

## Reviews

### Review by @Andrewgdewar - Approved on July 10, 2024 at 03:21 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/738*
