---
type: pull_request
number: 928
title: "HMS-5202: don't always rerun failed snaps"
state: merged
author: jlsherrill
created: 2024-12-20T14:50:48Z
updated: 2024-12-20T19:19:32Z
closed: 2024-12-20T19:19:28Z
merged: 2024-12-20T19:19:28Z
base_branch: main
head_branch: 5202
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/928
---

# Pull Request #928: HMS-5202: don't always rerun failed snaps

**Author**: @jlsherrill
**Created**: December 20, 2024 at 02:50 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `5202`

## Description

## Summary

After some investigation it was discovered that currently all failed snapshots are being re-run every hour.  This results in no 'extra' repos being synced.   This change results in failed snapshots only being re-run every day.  

This also fixes a couple of things, such as sorting by finished_at (instead of queued_at), and re-snapshotting every 25 hours instead of 24 (need to subtract one from the query)

## Testing steps

1.   Create 3 repos with snapshotting, two that are real and one that isn't:

https://fixtures.pulpproject.org/rpm-unsigned/
https://jlsherrill.fedorapeople.org/fake-repos/needed-errata/
https://redhat.com/notarepo

2.  on main, run `go run cmd/external-repos/main.go nightly-jobs 24`

for consistentcy do not re-run it until snapshots are done
everytime you run it, you'll see one snapshot get run and its always the same one, the one that failed

3.  re run `go run cmd/external-repos/main.go nightly-jobs 24`  on this PR, 
now everytime you run it, you'll get a different  one of the repos which should alternate between all 3:

i added some logging to make this easier to see: ```enqueued snapshot for repository config 147df40e-e1be-40a5-bcb2-726b6e63818c```


---

## Discussion

### Comment by @jlsherrill on December 20, 2024 at 03:37 PM UTC

[test]

### Comment by @jlsherrill on December 20, 2024 at 04:38 PM UTC

https://issues.redhat.com/browse/HMS-5202

### Comment by @jlsherrill on December 20, 2024 at 07:19 PM UTC

This actually doesn't require qe

---

## Reviews

### Review by @rverdile - Approved on December 20, 2024 at 07:10 PM UTC

looks good!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/928*
