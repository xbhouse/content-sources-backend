---
type: pull_request
number: 1367
title: "HMS-9470: add support for extended release repos"
state: merged
author: dominikvagner
created: 2026-01-20T12:49:06Z
updated: 2026-02-03T18:38:04Z
closed: 2026-02-03T18:38:04Z
merged: 2026-02-03T18:38:04Z
base_branch: main
head_branch: add-extended-support-repos
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1367
---

# Pull Request #1367: HMS-9470: add support for extended release repos

**Author**: @dominikvagner
**Created**: January 20, 2026 at 12:49 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `add-extended-support-repos`

## Description

## Summary
This adds support for extended release repositories (EUS, E4S). For that, two new attributes were added into repository configurations along with support for filtering by them. Also creates new snapshotted repos entries, with the currently supported versions of those features.

## Testing steps
Import the new repos and verify they are able to be snapshotted correctly.
Test out filtering with the `extended_release` and `extended_release_version` filters set.


---

## Discussion

### Comment by @xbhouse on January 20, 2026 at 01:00 PM UTC

https://issues.redhat.com/browse/HMS-9470

### Comment by @dominikvagner on January 22, 2026 at 09:15 AM UTC

> this looks good. but I think you were going to add more repository versions, right? I think that's the only thing missing

awesome! 😄 
I already did put the additional ones (all the currently supported) in here, or am I still missing some? 💭😅 

### Comment by @rverdile on January 22, 2026 at 03:07 PM UTC

> I already did put the additional ones (all the currently supported) in here, or am I still missing some?

oh I guess I just missed that you had already pushed them :)

### Comment by @rverdile on February 03, 2026 at 06:37 PM UTC

I'm going to merge this so I can rebase my PRs and make them ready for review

---

## Reviews

### Review by @rverdile - Commented on January 21, 2026 at 03:23 PM UTC

this looks good. but I think you were going to add more repository versions, right? I think that's the only thing missing

### Review by @rverdile - Commented on January 21, 2026 at 09:18 PM UTC

### Review by @dominikvagner - Commented on January 22, 2026 at 09:12 AM UTC

### Review by @rverdile - Commented on January 22, 2026 at 03:06 PM UTC

### Review by @dominikvagner - Commented on January 22, 2026 at 03:28 PM UTC

### Review by @rverdile - Dismissed on January 22, 2026 at 08:08 PM UTC

nice!

### Review by @rverdile - Commented on January 27, 2026 at 02:21 PM UTC

### Review by @rverdile - Commented on January 27, 2026 at 02:21 PM UTC

### Review by @dominikvagner - Commented on January 29, 2026 at 12:12 PM UTC

### Review by @rverdile - Commented on January 29, 2026 at 02:50 PM UTC

thanks :) two more small things

we also have to set a default [here](https://github.com/content-services/content-sources-backend/blob/main/pkg/config/config.go#L432) for the feature variable, and add it to the deployment.yaml

### Review by @rverdile - Approved on February 03, 2026 at 03:01 PM UTC

looks good!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1367*
