---
type: pull_request
number: 917
title: "Fixes 5144: nightly job failing"
state: merged
author: xbhouse
created: 2024-12-10T16:44:05Z
updated: 2024-12-13T15:24:57Z
closed: 2024-12-13T15:24:57Z
merged: 2024-12-13T15:24:57Z
base_branch: main
head_branch: 5144
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/917
---

# Pull Request #917: Fixes 5144: nightly job failing

**Author**: @xbhouse
**Created**: December 10, 2024 at 04:44 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `5144`

## Description

## Summary

This fixes the nightly job failure by checking if the context was canceled before closing the pool. Here I added back a separate context for the listener, there might be a better way to do this without that though 🤔 

## Testing steps

1. Start the server and run `make repos-import-rhel9 && go run cmd/external-repos/main.go nightly-jobs`
2. Should not panic and RH repos should be snapshotted



---

## Discussion

### Comment by @jlsherrill on December 10, 2024 at 05:00 PM UTC

https://issues.redhat.com/browse/HMS-5144

### Comment by @jlsherrill on December 12, 2024 at 02:15 PM UTC

This solved most issues, but i'm still seeing:

```
9:14AM ERR Error unlistening for tasks in dequeue error="conn closed"
```

sometimes, maybe like only 1/4th of the time

### Comment by @jlsherrill on December 12, 2024 at 02:44 PM UTC

Also worth noting, on main "go run cmd/external-repos/main.go snapshot https://cdn.redhat.com/content/dist/layered/rhel8/x86_64/ansible/2/os/" seems to hang, but this resolves it!

### Comment by @rverdile on December 12, 2024 at 05:43 PM UTC

I took a look at this 
> 9:14AM ERR Error unlistening for tasks in dequeue error="conn closed"

I don't really have a good solution. If closing the pool is causing too much weirdness, then I'm not sure we _need_ to close the pool outside of the tests. Similarly, I don't know how important it really is to unlisten if the connection is already closed?


### Comment by @jlsherrill on December 13, 2024 at 01:51 PM UTC

@rverdile and i chatted about this, and think a waitgroup may help here?

### Comment by @jlsherrill on December 13, 2024 at 01:52 PM UTC

we might also go ahead and merge this so fix the stage jobs and then we can deal with this issue separately ?

### Comment by @xbhouse on December 13, 2024 at 02:21 PM UTC

> we might also go ahead and merge this so fix the stage jobs and then we can deal with this issue separately ?

ok, yea i think merging to fix the nightly-jobs is a good idea. could you open a card for the other issue?

### Comment by @jlsherrill on December 13, 2024 at 02:32 PM UTC

will do!  

---

## Reviews

### Review by @jlsherrill - Approved on December 13, 2024 at 02:32 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/917*
