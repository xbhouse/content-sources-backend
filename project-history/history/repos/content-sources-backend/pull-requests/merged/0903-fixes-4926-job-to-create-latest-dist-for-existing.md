---
type: pull_request
number: 903
title: "Fixes 4926: job to create latest dist for existing repos"
state: merged
author: rverdile
created: 2024-11-22T21:37:04Z
updated: 2024-12-03T19:33:19Z
closed: 2024-12-03T19:33:12Z
merged: 2024-12-03T19:33:12Z
base_branch: main
head_branch: latest_config_job
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/903
---

# Pull Request #903: Fixes 4926: job to create latest dist for existing repos

**Author**: @rverdile
**Created**: November 22, 2024 at 09:37 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `latest_config_job`

## Description

## Summary
A job to create the latest snapshot distribution for existing repos, because existing ones would have been missed by this PR: https://github.com/content-services/content-sources-backend/pull/844#discussion_r1815501837

## Testing steps
1. Apply this patch `git apply mypatch.patch` to prevent the latest snapshot distribution being created in the snapshot task.
```
diff --git a/pkg/tasks/snapshot_helper.go b/pkg/tasks/snapshot_helper.go
index 1c0e812..75e9bc6 100644
--- a/pkg/tasks/snapshot_helper.go
+++ b/pkg/tasks/snapshot_helper.go
@@ -63,12 +63,12 @@ func (sh *SnapshotHelper) Run(versionHref string) error {
                return err
        }
 
-       latestPathIdent := helpers.GetLatestRepoDistPath(*sh.payload.GetSnapshotIdent())
-
-       _, _, err = helper.CreateOrUpdateDistribution(sh.orgId, sh.repo.UUID, latestPathIdent, publicationHref)
-       if err != nil {
-               return err
-       }
+       //latestPathIdent := helpers.GetLatestRepoDistPath(*sh.payload.GetSnapshotIdent())
+       //
+       //_, _, err = helper.CreateOrUpdateDistribution(sh.orgId, sh.repo.UUID, latestPathIdent, publicationHref)
+       //if err != nil {
+       //      return err
+       //}
        version, err := sh.pulpClient.GetRpmRepositoryVersion(sh.ctx, versionHref)
        if err != nil {
                return err
```
2. Create multiple repos for multiple orgs, some with and without snapshotting
3. Notice that the for the "latest_snapshot_url" returned by those repos, the URL gives 404.
4. Run the job: `go run cmd/create_latest_distributions/main.go`
5. Those URLs that were giving 404s should now serve the latest snapshot


---

## Discussion

### Comment by @jlsherrill on November 22, 2024 at 10:00 PM UTC

https://issues.redhat.com/browse/HMS-4926

---

## Reviews

### Review by @xbhouse - Commented on December 02, 2024 at 06:54 PM UTC

### Review by @xbhouse - Commented on December 02, 2024 at 07:06 PM UTC

### Review by @xbhouse - Commented on December 02, 2024 at 07:11 PM UTC

### Review by @rverdile - Commented on December 02, 2024 at 09:43 PM UTC

### Review by @rverdile - Commented on December 02, 2024 at 09:46 PM UTC

### Review by @rverdile - Commented on December 02, 2024 at 09:47 PM UTC

### Review by @xbhouse - Approved on December 03, 2024 at 04:03 PM UTC

looks good!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/903*
