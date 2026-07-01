---
type: pull_request
number: 1468
title: "HMS-10346: Add support CSV list of repo features"
state: merged
author: swadeley
created: 2026-04-17T18:49:01Z
updated: 2026-05-19T16:07:29Z
closed: 2026-05-19T16:07:29Z
merged: 2026-05-19T16:07:29Z
base_branch: main
head_branch: swadeley/HMS-10346
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1468
---

# Pull Request #1468: HMS-10346: Add support CSV list of repo features

**Author**: @swadeley
**Created**: April 17, 2026 at 06:49 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `swadeley/HMS-10346`

## Description

## Summary
HMS-10346 support feature filter with overlapping product IDs
 Feature filter can now handle comma separated list.
## Testing steps

On a running backend:

Set both of RHEL-EEUS-x86_64,RHEL-E4S-x86_64 in configs/config.yaml

run `make repos-import`

Expect to see two rows with:

```
    ---------------------|           feature_name           |                                                 
------------------------ | RHEL-EEUS-x86_64,RHEL-E4S-x86_64 | 
```


Set one of RHEL-EEUS-x86_64,RHEL-E4S-x86_64 in  in configs/config.yaml

Expect to see the same two rows.



---

## Discussion

### Comment by @xbhouse on April 17, 2026 at 07:00 PM UTC

https://issues.redhat.com/browse/HMS-10346

### Comment by @swadeley on April 20, 2026 at 08:04 PM UTC

/retest

### Comment by @rverdile on May 18, 2026 at 08:41 PM UTC

one more comment, otherwise looks good

### Comment by @swadeley on May 19, 2026 at 04:07 PM UTC

> nice job! looks good

Thank you

---

## Reviews

### Review by @rverdile - Commented on April 28, 2026 at 02:58 PM UTC

A few comments, but looks good overall!

### Review by @swadeley - Commented on April 30, 2026 at 09:18 AM UTC

### Review by @swadeley - Commented on April 30, 2026 at 09:38 AM UTC

### Review by @swadeley - Commented on April 30, 2026 at 09:38 AM UTC

### Review by @swadeley - Commented on April 30, 2026 at 09:38 AM UTC

### Review by @rverdile - Commented on May 05, 2026 at 07:02 PM UTC

~~When multiple repos with the same feature name are processed concurrently (e.g. during a batch job or parallel snapshot tasks), two goroutines can both call `fetchFeatureGuard`, both get nil, and then both try to create the same guard — the second fails with 400~~

ignore this :)

### Review by @rverdile - Commented on May 05, 2026 at 08:57 PM UTC

Thanks for the updates, looking really good. Just a couple of comments. I was testing it and seems like it's working. Although we can't properly verify pulp until this is in stage.

### Review by @swadeley - Commented on May 06, 2026 at 02:44 PM UTC

### Review by @swadeley - Commented on May 06, 2026 at 02:44 PM UTC

### Review by @rverdile - Commented on May 13, 2026 at 06:30 PM UTC

### Review by @swadeley - Commented on May 14, 2026 at 06:16 PM UTC

### Review by @swadeley - Commented on May 14, 2026 at 06:23 PM UTC

### Review by @rverdile - Commented on May 18, 2026 at 08:35 PM UTC

### Review by @swadeley - Commented on May 19, 2026 at 09:47 AM UTC

### Review by @rverdile - Approved on May 19, 2026 at 02:22 PM UTC

nice job! looks good

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1468*
