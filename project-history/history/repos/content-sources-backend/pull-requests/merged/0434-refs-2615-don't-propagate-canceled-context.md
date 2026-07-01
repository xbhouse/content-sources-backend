---
type: pull_request
number: 434
title: "Refs 2615: don't propagate canceled context"
state: merged
author: rverdile
created: 2023-10-19T13:49:11Z
updated: 2023-10-20T20:26:33Z
closed: 2023-10-20T20:16:31Z
merged: 2023-10-20T20:16:31Z
base_branch: main
head_branch: fix-broken-change
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/434
---

# Pull Request #434: Refs 2615: don't propagate canceled context

**Author**: @rverdile
**Created**: October 19, 2023 at 01:49 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `fix-broken-change`

## Description

## Summary
In PR #414, I mistakenly set context on a shared struct. This means that unless the context is manually reset by using `InitializePulpClient` before calling the `repoConfig.List` or `repoConfig.Fetch` dao methods, the request will fail with

`{"errors":[{"status":500,"title":"Error fetching repository","detail":"Get \"http://cs-pulp-api-svc.content-sources-stage.svc.cluster.local:24817/pulp/api/v3/status/\": context canceled"}]}`


This changes makes it so `context.TODO()` is used instead of the propagated context. This is only a quick and temporary fix to unblock things for now.

## Testing steps
1. Fetch a repository
2. Introspect the same repository
3. Without this change, the introspection will fail with the 500 error.


---

## Discussion

### Comment by @jlsherrill on October 19, 2023 at 02:00 PM UTC

https://issues.redhat.com/browse/HMS-2615

### Comment by @jlsherrill on October 19, 2023 at 02:01 PM UTC

⚠️ This task currently requires qe-approval, but this PR does not fully resolve the task.  Please contact QE to determine appropriate testing required.

### Comment by @swadeley on October 19, 2023 at 07:52 PM UTC

Hello


```
Type
    introspect
Status
    completed
Queued At
    19 Oct 2023 20:51 UTC+01:00
Started At
    19 Oct 2023 20:51 UTC+01:00
Finished At
    19 Oct 2023 20:51 UTC+01:00
Error
    None
```

lgtm

### Comment by @jlsherrill on October 19, 2023 at 08:44 PM UTC

/retest

### Comment by @swadeley on October 20, 2023 at 08:15 PM UTC

Works, UI and API crud, thank you

---

## Reviews

### Review by @jlsherrill - Commented on October 19, 2023 at 09:03 PM UTC

### Review by @rverdile - Commented on October 20, 2023 at 02:11 PM UTC

### Review by @swadeley - Approved on October 20, 2023 at 08:16 PM UTC

lgtm

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/434*
