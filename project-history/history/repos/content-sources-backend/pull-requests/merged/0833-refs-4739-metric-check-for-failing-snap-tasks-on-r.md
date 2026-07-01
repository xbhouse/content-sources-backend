---
type: pull_request
number: 833
title: "Refs 4739: metric check for failing snap tasks on RH repos"
state: merged
author: dominikvagner
created: 2024-09-26T08:35:53Z
updated: 2024-10-14T06:29:21Z
closed: 2024-10-14T06:26:55Z
merged: 2024-10-14T06:26:55Z
base_branch: main
head_branch: 4739
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/833
---

# Pull Request #833: Refs 4739: metric check for failing snap tasks on RH repos

**Author**: @dominikvagner
**Created**: September 26, 2024 at 08:35 AM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `4739`

## Description

## Summary
This PR adds a new metric for checking if the snapshot tasks are failing on RH repos, specifically if there isn't a successful one in the last 36 hours and returns number of repos on which is this happening.

## Testing steps
1. Import RH repos.
2. Make the snapshot task fail for some repo(s) (i.e. status = failed) and see if the metric increases.




---

## Discussion

### Comment by @jlsherrill on September 26, 2024 at 09:00 AM UTC

https://issues.redhat.com/browse/HMS-4739

### Comment by @mshriver on October 14, 2024 at 06:29 AM UTC

@swadeley @jlsherrill @dominikvagner  I like this approach to production monitoring rather than trying to write an IQE test around this situation and identifying the failure mode from production test results. thanks!

---

## Reviews

### Review by @xbhouse - Approved on October 01, 2024 at 03:51 PM UTC

metrics look good to me! @rverdile can you take a look when you get a chance in case i missed anything?

### Review by @rverdile - Commented on October 02, 2024 at 06:44 PM UTC

Small comment, looks really good overall

### Review by @rverdile - Commented on October 02, 2024 at 07:35 PM UTC

### Review by @dominikvagner - Commented on October 11, 2024 at 03:45 PM UTC

### Review by @dominikvagner - Commented on October 11, 2024 at 03:47 PM UTC

### Review by @rverdile - Approved on October 11, 2024 at 04:34 PM UTC

looks good!!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/833*
