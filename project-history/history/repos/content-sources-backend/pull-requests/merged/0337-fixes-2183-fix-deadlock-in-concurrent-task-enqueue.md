---
type: pull_request
number: 337
title: "Fixes 2183: fix deadlock in concurrent task enqueue"
state: merged
author: jlsherrill
created: 2023-07-20T17:49:24Z
updated: 2023-07-21T08:00:20Z
closed: 2023-07-21T07:51:51Z
merged: 2023-07-21T07:51:51Z
base_branch: main
head_branch: 2183
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/337
---

# Pull Request #337: Fixes 2183: fix deadlock in concurrent task enqueue

**Author**: @jlsherrill
**Created**: July 20, 2023 at 05:49 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `2183`

## Description

## Summary

This resolves a couple of issues:
* Enqueue() was not properly releasing its acquired db pool connection, this resulted in deadlocks when more than ~20 create reqs came in at once
* Sets a limit on the postgres connections for gorm, as if you perform ~50 create reqs at the same time a postgres server with a default limit of connections will start to max out and start issuing errors
* Reworks pgqueue to use a wrapper around pgxpool.Pool, this allows for testing while giving pgqueue the ability to start/end transactions and acquire/release Pool connections


## Testing steps

Checkout this pr and 'make run'
create a new file `./cmd/test/main.go`   with the contents here: https://gist.github.com/jlsherrill/d07efa70bed152d8428c758b7332f95e

then run it:  `go run ./cmd/test/main.go`

You should only get back 201s, and all tasks should be processed without error (other than the 404)

---

## Discussion

### Comment by @jlsherrill on July 20, 2023 at 06:00 PM UTC

https://issues.redhat.com/browse/HMS-2183

### Comment by @rverdile on July 20, 2023 at 06:54 PM UTC

everything else looks good to me

### Comment by @jlsherrill on July 20, 2023 at 07:42 PM UTC

updated in a new commit

### Comment by @swadeley on July 21, 2023 at 07:51 AM UTC

Tested API OK.

---

## Reviews

### Review by @rverdile - Commented on July 20, 2023 at 06:30 PM UTC

### Review by @jlsherrill - Commented on July 20, 2023 at 06:59 PM UTC

### Review by @rverdile - Approved on July 20, 2023 at 08:02 PM UTC

nice!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/337*
