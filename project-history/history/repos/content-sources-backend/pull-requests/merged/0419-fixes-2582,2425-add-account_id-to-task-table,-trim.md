---
type: pull_request
number: 419
title: "Fixes 2582,2425: add account_id to task table, trim error"
state: merged
author: jlsherrill
created: 2023-10-06T15:31:58Z
updated: 2023-10-16T13:01:04Z
closed: 2023-10-13T17:56:17Z
merged: 2023-10-13T17:56:17Z
base_branch: main
head_branch: 2582
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/419
---

# Pull Request #419: Fixes 2582,2425: add account_id to task table, trim error

**Author**: @jlsherrill
**Created**: October 06, 2023 at 03:31 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `2582`

## Description

## Summary

Adds account_id to tasks table.  This makes querying by account_id easier and faster.
Note that after this change, any existing tasks will have a null account_id, but since we are still in stage, that is okay.

Change error length to 4000 chars, and auto trim error if its longer than that

I did these in a single PR as they would normally conflict

## Testing steps

1. Deploy this pr and migrate the db
2. Snapshot some repository
3. Use the admin tasks interface
4. filter by account id

Verify nothing is broken, behavior should be the same

Unsure how to test the error length, i wasn't able to reproduce it but the unit test should prove that its solved



---

## Discussion

### Comment by @jlsherrill on October 06, 2023 at 05:00 PM UTC

https://issues.redhat.com/browse/HMS-2425

### Comment by @swadeley on October 09, 2023 at 02:44 PM UTC

Hi

I deployed this PR, ran automated tests and made manual check in UI. Works as described, I can filter on account ID, org ID, all works as expected.

API test result pending.

### Comment by @swadeley on October 10, 2023 at 09:54 AM UTC

Hi, introspection seems to be misbehaving. Repos with `example.com` never finish introspecting. I see three  "delete-repository-snapshots" tasks which have been running for nearly 24 hours.

### Comment by @jlsherrill on October 10, 2023 at 12:05 PM UTC

/retest

### Comment by @jlsherrill on October 10, 2023 at 02:30 PM UTC

https://issues.redhat.com/browse/HMS-2582

### Comment by @jlsherrill on October 10, 2023 at 06:36 PM UTC

i wasn't able to reproduce, are you able to?

### Comment by @swadeley on October 11, 2023 at 12:25 PM UTC

> i wasn't able to reproduce, are you able to?

Hi, I deployed new image and now introspection completes as expected.

Thank you.

### Comment by @swadeley on October 11, 2023 at 12:28 PM UTC

> > i wasn't able to reproduce, are you able to?
> 
> Hi, I deployed new image and now introspection completes as expected.
> 
> Thank you.

oops, spoke too soon. It completes OK for popular repo, but for small custom repo it changed to status Valid, but when I check Admin tab I see its still running. It is also trying to make a snapshot but I did not enable that feature.

### Comment by @swadeley on October 13, 2023 at 09:05 AM UTC

> Hi
> 
> I deployed this PR, ran automated tests and made manual check in UI. Works as described, I can filter on account ID, org ID, all works as expected.
> 
> API test result pending.

Tested again and ran all API tests.

Thank you.

### Comment by @swadeley on October 13, 2023 at 09:06 AM UTC

@rverdile Merge after your approval.

### Comment by @jlsherrill on October 13, 2023 at 12:33 PM UTC

/retest

---

## Reviews

### Review by @rverdile - Approved on October 13, 2023 at 05:12 PM UTC

tested and lgtm!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/419*
