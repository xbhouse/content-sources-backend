---
type: pull_request
number: 162
title: "Fixes 319: Order Arch list alphabetically"
state: merged
author: Andrewgdewar
created: 2023-01-03T19:37:53Z
updated: 2023-01-05T15:19:42Z
closed: 2023-01-05T15:16:46Z
merged: 2023-01-05T15:16:46Z
base_branch: main
head_branch: CONTENT-319
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/162
---

# Pull Request #162: Fixes 319: Order Arch list alphabetically

**Author**: @Andrewgdewar
**Created**: January 03, 2023 at 07:37 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `CONTENT-319`

## Description

*No description provided*

---

## Discussion

### Comment by @jlsherrill on January 03, 2023 at 08:00 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-319

### Comment by @swadeley on January 05, 2023 at 12:39 PM UTC

Hi

I have tested using image  pr-162-latest and it works. I was just wondering if some users might be confused by having `aarch64` before `Any` in the list of repos because all the other arches are on the other end of the sorted repo list. So you will see a column with lots of arches, but not `aarch64` , then all the `Any`, then the `aarch64` if you have some in your custom repos.

### Comment by @jlsherrill on January 05, 2023 at 01:26 PM UTC

@swadeley can you paste a screenshot of what you mean?  What i see looks fine:  
![image](https://user-images.githubusercontent.com/395077/210790530-00892c84-5b20-4a79-a4b0-dd9f17db474e.png)


### Comment by @swadeley on January 05, 2023 at 01:58 PM UTC

@justin this is what I mean, Any, which is not an arch, is separating aarch64 from the other arches:
![image](https://user-images.githubusercontent.com/11247237/210796798-d89f5a11-0bd3-4f03-91a0-eade1b8ff986.png)


### Comment by @swadeley on January 05, 2023 at 01:59 PM UTC

@jlsherrill Have I misunderstood the scope of this PR? is it just to sort the menu?

### Comment by @swadeley on January 05, 2023 at 03:16 PM UTC

> @jlsherrill Have I misunderstood the scope of this PR? is it just to sort the menu?

Yes, and I see the associated jira ticket name is more descriptive:  `   arch list on dropdown & creation screen is not sorted alphabetically`

### Comment by @swadeley on January 05, 2023 at 03:19 PM UTC

New Jira for the repo sort issue [Repos are not sorted by arch as expected](https://issues.redhat.com/browse/HMSCONTENT-327)

---

## Reviews

### Review by @Andrewgdewar - Commented on January 03, 2023 at 07:38 PM UTC

### Review by @Andrewgdewar - Commented on January 03, 2023 at 07:39 PM UTC

### Review by @jlsherrill - Approved on January 04, 2023 at 04:50 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/162*
