---
type: pull_request
number: 479
title: "Fixes 3066: introspect aarch64 repositories"
state: merged
author: croissanne
created: 2023-11-21T11:35:23Z
updated: 2023-11-21T15:35:43Z
closed: 2023-11-21T15:35:01Z
merged: 2023-11-21T15:35:01Z
base_branch: main
head_branch: introspect
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/479
---

# Pull Request #479: Fixes 3066: introspect aarch64 repositories

**Author**: @croissanne
**Created**: November 21, 2023 at 11:35 AM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `introspect`

## Description

## Summary
Google aarch64 repositories don't need to be introspected yet, as we don't build aarch64 images for google cloud.



## Testing steps


## Checklist

- [x] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @app-sre-bot on November 21, 2023 at 11:36 AM UTC

Can one of the admins verify this patch?

### Comment by @jlsherrill on November 21, 2023 at 11:36 AM UTC

https://issues.redhat.com/browse/HMS-3066

### Comment by @jlsherrill on November 21, 2023 at 01:43 PM UTC

oktotest

### Comment by @jlsherrill on November 21, 2023 at 01:50 PM UTC

/test

### Comment by @jlsherrill on November 21, 2023 at 01:51 PM UTC

[test]  


### Comment by @jlsherrill on November 21, 2023 at 03:34 PM UTC

failing test is unrelated and fixed in master.

---

## Reviews

### Review by @jlsherrill - Approved on November 21, 2023 at 01:43 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/479*
