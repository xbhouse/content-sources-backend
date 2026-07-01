---
type: pull_request
number: 1543
title: "HMS-10544: Add lookup for orphaned distributions"
state: open
author: swadeley
created: 2026-06-18T16:50:23Z
updated: 2026-06-25T16:06:25Z
base_branch: main
head_branch: swadeley/HMS-10544
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1543
---

# Pull Request #1543: HMS-10544: Add lookup for orphaned distributions

**Author**: @swadeley
**Created**: June 18, 2026 at 04:50 PM UTC
**Status**: Open
**Labels**: None
**Base**: `main` ← **Head**: `swadeley/HMS-10544`

## Description



## Summary
Add lookup for orphaned distributions
To work with delete-snapshots task
## Testing steps

start new clean backend without any repos, i.e. `make compose-clean`

Automated
```
go test ./pkg/tasks/ -run TestDeleteRepositoryVersionCleansUntrackedBlockingDistributions -v
go test ./test/integration/ -run TestDeleteSnapshotCleansUntrackedBlockingDistribution -v (with local compose stack)
```

Manual (local)

Setup

make compose-up and make run
`Org header: HDR="$(./scripts/header.sh $ORG_ID snapUser)"`

Repo with snapshotting enabled using https://fixtures.pulpproject.org/rpm-unsigned/
Simulate untracked Pulp distribution

Enable a temporary failure before GetRpmRepositoryVersion in snapshot_helper.go 

I added this to `pkg/tasks/snapshot_helper.go`:
```
	 if os.Getenv("SIMULATE_GET_RPM_REPOSITORY_VERSION_TIMEOUT") == "1" {
	 	return fmt.Errorf("simulated timeout fetching repository version: %w", context.DeadlineExceeded)
	 }
```
and then reran the backend with `SIMULATE_GET_RPM_REPOSITORY_VERSION_TIMEOUT=1` 

Create the repo and let the first snapshot task fail after creating Pulp resources but before writing the snapshot row.
Disable the simulation (temporary failure code) and restart the backend.
Change the repo URL (e.g. to rpm-with-sha-512) and trigger a second snapshot so a valid snapshot exists in the DB.

Exercise delete-snapshots cleanup

Confirm two snapshots exist:
`curl -s ".../repositories/$REPO_CONFIG_UUID/snapshots/" -H "$HDR" | jq '.meta.count'`
On the older snapshot, set a bogus distribution_href so delete skips the real Pulp distribution:
```
UPDATE snapshots
SET distribution_href = 'bogus-untracked-distribution-href'
WHERE uuid = '<older_snapshot_uuid>';
Delete the older snapshot via API (no trailing slash after $REPO_CONFIG_UUID):
curl -s -X DELETE \
  ".../repositories/$REPO_CONFIG_UUID/snapshots/<older_snapshot_uuid>" \
  -H "$HDR"
```
Expected results

DELETE returns 204
delete-snapshots task status is completed with no error
Snapshot count drops from 2 to 1
Worker logs include: Deleting untracked distribution blocking repository version deletion
Pulp distribution at the snapshot’s distribution_path and its repository version are removed despite the bogus DB distribution_href

Look for a line like this in the terminal where backend is running:
`1:19PM INF Deleting untracked distribution blocking repository version deletion base_path=`


---

## Discussion

### Comment by @xbhouse on June 18, 2026 at 05:00 PM UTC

https://issues.redhat.com/browse/HMS-10544

### Comment by @rverdile on June 24, 2026 at 07:00 PM UTC

A few comments, looks good overall :)

---

## Reviews

### Review by @rverdile - Commented on June 24, 2026 at 06:41 PM UTC

### Review by @swadeley - Commented on June 25, 2026 at 09:25 AM UTC

### Review by @swadeley - Commented on June 25, 2026 at 04:03 PM UTC

### Review by @swadeley - Commented on June 25, 2026 at 04:03 PM UTC

### Review by @swadeley - Commented on June 25, 2026 at 04:04 PM UTC

### Review by @swadeley - Commented on June 25, 2026 at 04:06 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1543*
