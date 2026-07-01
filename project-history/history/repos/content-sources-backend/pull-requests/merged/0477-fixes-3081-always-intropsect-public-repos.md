---
type: pull_request
number: 477
title: "Fixes 3081: always intropsect public repos"
state: merged
author: jlsherrill
created: 2023-11-20T17:26:44Z
updated: 2023-11-20T17:59:44Z
closed: 2023-11-20T17:59:36Z
merged: 2023-11-20T17:59:36Z
base_branch: main
head_branch: 3081
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/477
---

# Pull Request #477: Fixes 3081: always intropsect public repos

**Author**: @jlsherrill
**Created**: November 20, 2023 at 05:26 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `3081`

## Description

## Summary
even if failed counts has been reached

As public repos should never fail  in the real world, some action may need to be taken, and we can retry.

We add RHEL repos before they are actually on the cdn, which means it can fail introspection for some time before its officially there.  Once it is, we want to introspect it properly.

## Testing steps

I wrote a test to test this, i think test passing is sufficient. 

## Checklist

- [x] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on November 20, 2023 at 05:30 PM UTC

https://issues.redhat.com/browse/HMS-3081

### Comment by @jlsherrill on November 20, 2023 at 05:59 PM UTC

since iqe tests passed, and the only test failure is some ci issue, i'm merging

---

## Reviews

### Review by @Andrewgdewar - Approved on November 20, 2023 at 05:31 PM UTC

LGTM!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/477*
