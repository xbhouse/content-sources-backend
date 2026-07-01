---
type: pull_request
number: 1399
title: "HMS-5855: remove custom EPEL repositories"
state: merged
author: rverdile
created: 2026-02-26T19:54:09Z
updated: 2026-04-02T14:22:33Z
closed: 2026-04-02T14:22:22Z
merged: 2026-04-02T14:22:22Z
base_branch: main
head_branch: remove-custom-epel
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1399
---

# Pull Request #1399: HMS-5855: remove custom EPEL repositories

**Author**: @rverdile
**Created**: February 26, 2026 at 07:54 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `remove-custom-epel`

## Description

## Summary
Leaving this in draft until we're ready to run this job

Queries for all the custom repositories, soft deletes them, then enqueues the task to fully delete them and their snapshots. Very similar to the repository delete handler.

## Testing steps
1. Set `features.allow_custom_epel_creeation.enabled=true` in your config.yaml
2. Create some custom EPEL repositories under different orgs
3. Run the job: `go run cmd/jobs/main.go remove-custom-epel-repos`



---

## Discussion

### Comment by @xbhouse on February 26, 2026 at 08:00 PM UTC

https://issues.redhat.com/browse/HMS-5855

### Comment by @rverdile on March 05, 2026 at 09:57 PM UTC

Thanks! I realize this hasn't merged yet, so I'm gonna put this back in draft to avoid any accidents
https://github.com/osbuild/image-builder-crc/pull/1763

### Comment by @jlsherrill on March 12, 2026 at 08:02 PM UTC

we might also want to make sure  the new pulp auth is in place, so these deletes don't get rate limited

### Comment by @dominikvagner on March 23, 2026 at 08:55 AM UTC

the Image Builder crc [PR](https://github.com/osbuild/image-builder-crc/pull/1763) has been merged 🫡 _(don't know when it will get to prod)_
and our stage is using the packages Pulp endpoint now so maybe the stage job could be tested now 😅

### Comment by @dominikvagner on March 30, 2026 at 01:49 PM UTC

FYI: the prod migration of Pulp to packages is done 👍🏼 and the IB CRC change is in prod! 🫡 
so this is most likely unblocked 😅 

### Comment by @rverdile on March 31, 2026 at 01:44 PM UTC

/retest

### Comment by @rverdile on March 31, 2026 at 02:15 PM UTC

thanks all, rebased

### Comment by @rverdile on March 31, 2026 at 02:22 PM UTC

/retest

---

## Reviews

### Review by @xbhouse - Dismissed on March 05, 2026 at 09:48 PM UTC

nice job! looks good to me :)

### Review by @xbhouse - Approved on March 31, 2026 at 02:36 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1399*
