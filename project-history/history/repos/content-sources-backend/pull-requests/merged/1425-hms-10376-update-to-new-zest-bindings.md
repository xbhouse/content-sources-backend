---
type: pull_request
number: 1425
title: "HMS-10376: update to new zest bindings"
state: merged
author: jlsherrill
created: 2026-03-24T22:04:55Z
updated: 2026-03-25T13:45:54Z
closed: 2026-03-25T13:45:43Z
merged: 2026-03-25T13:45:43Z
base_branch: main
head_branch: zest_update
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1425
---

# Pull Request #1425: HMS-10376: update to new zest bindings

**Author**: @jlsherrill
**Created**: March 24, 2026 at 10:04 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `zest_update`

## Description

## Summary

Because we had not switched to 2026, we were a few months behind on our zest builds. This updates to the new bindings with some changes:

* for remote & distribution update apis, pulp is now properly returning different types depending on if an update is needed or not.  If its not, it returns a 200 with the object that was not actually updated.  If it does, it returns a wrapped task href.  The go generator cannot handle this, so for those two apis, lets just look at the response code and decode it to a task href if its a 202.
* updated 2025 to 2026

## Testing steps

* tests pass



---

## Discussion

### Comment by @xbhouse on March 24, 2026 at 10:30 PM UTC

https://issues.redhat.com/browse/HMS-10376

---

## Reviews

### Review by @dominikvagner - Commented on March 25, 2026 at 01:12 PM UTC

### Review by @jlsherrill - Commented on March 25, 2026 at 01:13 PM UTC

### Review by @jlsherrill - Commented on March 25, 2026 at 01:13 PM UTC

### Review by @dominikvagner - Approved on March 25, 2026 at 01:26 PM UTC

looks good 👍🏼 thanks! 
ack! ✅ 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1425*
