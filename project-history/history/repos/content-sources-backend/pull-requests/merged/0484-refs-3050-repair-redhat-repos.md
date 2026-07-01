---
type: pull_request
number: 484
title: "Refs 3050: repair redhat repos"
state: merged
author: jlsherrill
created: 2023-11-22T17:32:52Z
updated: 2023-11-22T18:22:41Z
closed: 2023-11-22T18:22:41Z
merged: 2023-11-22T18:22:41Z
base_branch: main
head_branch: repair_redhat
labels: []
url: https://github.com/content-services/content-sources-backend/pull/484
---

# Pull Request #484: Refs 3050: repair redhat repos

**Author**: @jlsherrill
**Created**: November 22, 2023 at 05:32 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `repair_redhat`

## Description

and delete newest version if does not match latest snapshot

## Summary

Currently in stage we have a couple issues, stemming from a repository snapshotting, and the publish failed due to a missing metadata file (that is downloaded, not generated).  

So for each red hat org's repository, we grab the latest_version_href from pulp and trigger a repair.  If this latest version doesn't match the repo's latest snapshot, we delete it. 

## Testing steps

You might edit the redhat_repo.json file and delete all except for the `Red Hat Ansible Engine 2 for RHEL 8 x86_64 (RPMs)` repo.  
Run 
```make repos-import```
```go run cmd/external_repos/main.go nightly-jobs```

This will sync the red hat repo.   Then run:

```go run cmd/repair_latest/main.go --force```

## Checklist

- [x] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on November 22, 2023 at 05:33 PM UTC

https://issues.redhat.com/browse/HMS-3050

---

## Reviews

### Review by @rverdile - Approved on November 22, 2023 at 06:07 PM UTC

lgtm

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/484*
