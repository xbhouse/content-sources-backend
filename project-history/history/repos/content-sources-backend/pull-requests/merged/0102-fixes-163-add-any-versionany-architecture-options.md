---
type: pull_request
number: 102
title: "Fixes 163: Add Any version/Any architecture options"
state: merged
author: Andrewgdewar
created: 2022-09-08T22:00:21Z
updated: 2022-09-21T20:30:28Z
closed: 2022-09-21T20:04:25Z
merged: 2022-09-21T20:04:25Z
base_branch: main
head_branch: CONTENT-163
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/102
---

# Pull Request #102: Fixes 163: Add Any version/Any architecture options

**Author**: @Andrewgdewar
**Created**: September 08, 2022 at 10:00 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `CONTENT-163`

## Description

*No description provided*

---

## Discussion

### Comment by @jlsherrill on September 08, 2022 at 10:30 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-163

### Comment by @Andrewgdewar on September 12, 2022 at 06:55 PM UTC

Made suggested changes, added a test. Woot for learning 😝 

---

## Reviews

### Review by @rverdile - Commented on September 09, 2022 at 07:07 PM UTC

Nice. One thing, I think you should move the validation to the model layer like here: https://github.com/content-services/content-sources-backend/blob/main/pkg/models/repository_configuration.go#L97-L102

### Review by @rverdile - Commented on September 13, 2022 at 02:12 PM UTC

I think it would be good to have a test that's in the models layer like here: https://github.com/content-services/content-sources-backend/blob/main/pkg/models/repository_configuration_test.go#L63-L82

If you have a test in the models layer, I'm not sure one would be needed in the dao layer, but I'm not necessarily opposed to having both.

### Review by @rverdile - Commented on September 13, 2022 at 02:29 PM UTC

### Review by @rverdile - Commented on September 13, 2022 at 02:53 PM UTC

### Review by @Andrewgdewar - Commented on September 13, 2022 at 02:57 PM UTC

### Review by @Andrewgdewar - Commented on September 13, 2022 at 02:58 PM UTC

### Review by @Andrewgdewar - Commented on September 13, 2022 at 02:59 PM UTC

### Review by @Andrewgdewar - Commented on September 13, 2022 at 03:29 PM UTC

### Review by @rverdile - Commented on September 13, 2022 at 03:29 PM UTC

### Review by @rverdile - Commented on September 13, 2022 at 03:30 PM UTC

### Review by @rverdile - Commented on September 13, 2022 at 03:32 PM UTC

### Review by @rverdile - Commented on September 13, 2022 at 07:34 PM UTC

If I create/update a repository with `distribution_versions:[]` in the body, it will make versions in the database an empty list.

### Review by @rverdile - Commented on September 14, 2022 at 01:58 PM UTC

### Review by @rverdile - Commented on September 14, 2022 at 02:06 PM UTC

### Review by @rverdile - Commented on September 14, 2022 at 02:13 PM UTC

### Review by @rverdile - Commented on September 14, 2022 at 06:52 PM UTC

### Review by @rverdile - Commented on September 14, 2022 at 06:55 PM UTC

### Review by @rverdile - Commented on September 14, 2022 at 07:22 PM UTC

### Review by @rverdile - Commented on September 14, 2022 at 07:24 PM UTC

### Review by @Andrewgdewar - Commented on September 14, 2022 at 09:45 PM UTC

### Review by @Andrewgdewar - Commented on September 14, 2022 at 09:45 PM UTC

### Review by @rverdile - Commented on September 15, 2022 at 02:32 PM UTC

### Review by @rverdile - Commented on September 19, 2022 at 03:21 PM UTC

### Review by @rverdile - Commented on September 19, 2022 at 03:25 PM UTC

### Review by @Andrewgdewar - Commented on September 19, 2022 at 09:14 PM UTC

### Review by @rverdile - Commented on September 19, 2022 at 09:42 PM UTC

### Review by @Andrewgdewar - Commented on September 20, 2022 at 06:28 PM UTC

### Review by @rverdile - Approved on September 20, 2022 at 08:08 PM UTC

nice! reminder not to merge, this still needs qe testing

### Review by @swadeley - Commented on September 21, 2022 at 06:35 PM UTC

### Review by @swadeley - Commented on September 21, 2022 at 08:04 PM UTC

lgtm

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/102*
