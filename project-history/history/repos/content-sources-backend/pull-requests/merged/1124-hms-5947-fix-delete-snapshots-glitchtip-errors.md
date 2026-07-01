---
type: pull_request
number: 1124
title: "HMS-5947: fix delete-snapshots glitchtip errors"
state: merged
author: xbhouse
created: 2025-06-11T16:02:33Z
updated: 2025-06-18T13:52:02Z
closed: 2025-06-18T13:52:02Z
merged: 2025-06-18T13:52:02Z
base_branch: main
head_branch: fix-snapshot-cleanup
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/1124
---

# Pull Request #1124: HMS-5947: fix delete-snapshots glitchtip errors

**Author**: @xbhouse
**Created**: June 11, 2025 at 04:02 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `fix-snapshot-cleanup`

## Description

## Summary

Template distributions were not being updated when cleaning up outdated snapshots for RH repos, resulting in failed deletion of repo versions. There might be other issues at play here, so I've left in the debug logs for now.

## Testing steps

I haven't found a way to test this without multiple RH snapshots, this hasn't been reproducible for me locally with custom or EPEL repos. If you have multiple RH snapshots, you can follow these steps:

1. Create a template with just the base RH repos for any version or arch. Set to a date earlier than the oldest snapshot so it uses the oldest
2. Update one of the RH repo's oldest snapshot so that it is set to be cleaned up. You can use this query:  `UPDATE snapshots SET created_at = (NOW() - CAST('365 days' AS INTERVAL)) WHERE uuid = '<snapshot_uuid>';`
3. Confirm the template now shows a warning that some repos have snapshots that are soon to be deleted.
4. Grab the pulp URL to the snapshot you'll be deleting so we can verify it doesn't exist later.
5. Run the snapshot cleanup task: `go run cmd/external-repos/main.go cleanup --type snapshot`
6. Follow the logs, you shouldn't see errors related to existing distributions and the task should succeed. 
7. In the template, the warning should disappear and you should see the next available snapshot of that RH repo on that template.
8. The RH snapshot should be fully deleted (not available when viewing snapshots in the UI or listing via the API, not stuck in a soft-deleted state in the DB, and the older snapshot URL should give you a 404).



---

## Discussion

### Comment by @jlsherrill on June 11, 2025 at 04:30 PM UTC

https://issues.redhat.com/browse/HMS-5947

---

## Reviews

### Review by @dominikvagner - Approved on June 18, 2025 at 06:58 AM UTC

looks good to me! thank for finding this issue and fixing it, good job 👏🏼😇
tested it locally, with simulated old RH snaps and everything seems to work as it should 🎉
approved ✅ 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1124*
