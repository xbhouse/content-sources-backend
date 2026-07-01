---
type: pull_request
number: 789
title: "Fixes 4681: Update to go 1.23"
state: merged
author: jlsherrill
created: 2024-08-27T13:59:56Z
updated: 2024-09-10T19:43:06Z
closed: 2024-09-10T19:41:41Z
merged: 2024-09-10T19:41:41Z
base_branch: main
head_branch: go123
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/789
---

# Pull Request #789: Fixes 4681: Update to go 1.23

**Author**: @jlsherrill
**Created**: August 27, 2024 at 01:59 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `go123`

## Description

## Summary
you can upgrade with
```
go get go@1.23.0
```
## Testing steps
tests pass


---

## Discussion

### Comment by @swadeley on August 29, 2024 at 11:27 AM UTC

/retest

### Comment by @swadeley on September 03, 2024 at 10:31 AM UTC

Hi @jlsherrill 

You mentioned the bug in openapi that changes the order of these values:
```
"sha256",
 "file"
```
and I notice the only change to the api docs in this PR is that change you mentioned.

Does that mean I can just ignore this, you will you revert that change, or do I still need to regenerate API docs?

### Comment by @jlsherrill on September 03, 2024 at 03:25 PM UTC

The ordering of  those items be non-deterministic, and possibly change anytime 'make openapi' is run until https://github.com/getkin/kin-openapi is merged.   I think we can just ignore this until that happens

### Comment by @jlsherrill on September 03, 2024 at 03:25 PM UTC

but i guess i could change it back to the correct order :) 

### Comment by @jlsherrill on September 03, 2024 at 03:27 PM UTC

actually this seems like the more common order that is generated, so lets leave it like this until that new dependency is released

### Comment by @jlsherrill on September 06, 2024 at 08:21 PM UTC

rebased and removed openapi change

### Comment by @swadeley on September 09, 2024 at 09:04 AM UTC

Hi

pr check test run says:
The ENV_FOR_DYNACONF is ephemeral

yet this test ran (it is tagged for stage only):
```
    @pytest.mark.stage
    def test_search_non_added_popular_repo
```

I'll trigger a retest, see if it happens again (might have to specifically omit it).

### Comment by @swadeley on September 09, 2024 at 09:04 AM UTC

/retest

### Comment by @jlsherrill on September 09, 2024 at 01:30 PM UTC

https://issues.redhat.com/browse/HMS-4681

### Comment by @jlsherrill on September 09, 2024 at 05:15 PM UTC

/restest

### Comment by @jlsherrill on September 09, 2024 at 05:28 PM UTC

/retest

### Comment by @swadeley on September 10, 2024 at 07:55 AM UTC

/retest

### Comment by @jlsherrill on September 10, 2024 at 01:22 PM UTC

/retest

### Comment by @jlsherrill on September 10, 2024 at 05:13 PM UTC

/retest

---

## Reviews

### Review by @rverdile - Approved on September 10, 2024 at 02:18 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/789*
