---
type: pull_request
number: 141
title: "Fixes 261: add limits when returning rpm searches"
state: merged
author: avisiedo
created: 2022-11-04T09:44:54Z
updated: 2023-01-16T11:30:57Z
closed: 2022-12-05T17:08:37Z
merged: 2022-12-05T17:08:37Z
base_branch: main
head_branch: hmscontent-261-limit-pkg-search
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/141
---

# Pull Request #141: Fixes 261: add limits when returning rpm searches

**Author**: @avisiedo
**Created**: November 04, 2022 at 09:44 AM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `hmscontent-261-limit-pkg-search`

## Description

- This update the request structure to add limit field.
- By default it set the value of 100.
- The maximum threshold is 500.

TODO
- [x] Add unit test for the handler.

Signed-off-by: Alejandro Visiedo <avisiedo@redhat.com>

---

## Discussion

### Comment by @jlsherrill on November 04, 2022 at 10:00 AM UTC

https://issues.redhat.com/browse/HMSCONTENT-261

### Comment by @avisiedo on November 15, 2022 at 04:13 PM UTC

Test scenarios

**Scenario 1: Default limit**

Given a repository config with more than 100 packages associated to the user organization
When we invoke the search endpoint '/api/content-sources/rpms/names'
   and search field is empty (to retrieve all the packages)
Then the returned list does not have more than 100 items

**Scenario 2: Set the limit**

Given a repository config with more than 50 packages associated to the user organization
When we invoke the search endpoint '/api/content-sources/rpms/names'
   and search field is empty (to retrieve all the packages)
   and the limit field value is 10
Then the returned list does not have more than 10 items

**Scenario 3: Restrict to maximum limit**

Given a repository config with more than 500 packages associated to the user organization
When we invoke the search endpoint '/api/content-sources/rpms/names'
   and search field is empty (to retrieve all the packages)
   and the limit field value is 501
Then the returned list does not have more than 500 items


### Comment by @swadeley on November 30, 2022 at 09:16 AM UTC

@avisiedo Please rebase

### Comment by @swadeley on November 30, 2022 at 10:58 AM UTC

/retest

### Comment by @avisiedo on December 02, 2022 at 10:29 AM UTC

thanks mates for the review! @rverdile I have added some comments; If I am overlooking something please let me know, happy to update the pr if needed.

I will create a new card to cover the unit test for the second handler `listRepositoriesRpm` (or if there some card related with it, it could be an opportunity to add it), so the coverage is increased.

Sorry for not providing the unit tests for the handlers when I created this code, at that moment I didn't know very well how to add them; now I am more capable to add unit tests for the http handlers.

### Comment by @rverdile on December 02, 2022 at 07:48 PM UTC

@avisiedo addressed all my concerns :)

---

## Reviews

### Review by @rverdile - Commented on December 01, 2022 at 06:39 PM UTC

### Review by @rverdile - Commented on December 01, 2022 at 07:00 PM UTC

### Review by @rverdile - Commented on December 01, 2022 at 07:05 PM UTC

### Review by @Andrewgdewar - Approved on December 01, 2022 at 08:17 PM UTC

Functionally works really well, code wise, looks like @rverdile has a few comments, so you may want to address those first.
This gets an ack from though!

### Review by @avisiedo - Commented on December 02, 2022 at 08:25 AM UTC

### Review by @avisiedo - Commented on December 02, 2022 at 09:57 AM UTC

### Review by @avisiedo - Commented on December 02, 2022 at 09:57 AM UTC

### Review by @rverdile - Commented on December 02, 2022 at 03:12 PM UTC

### Review by @rverdile - Commented on December 02, 2022 at 03:26 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/141*
