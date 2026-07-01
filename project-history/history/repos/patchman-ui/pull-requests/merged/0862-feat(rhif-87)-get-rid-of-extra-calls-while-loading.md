---
type: pull_request
number: 862
title: "feat(RHIF-87): get rid of extra calls while loading remediation"
state: merged
author: mkholjuraev
created: 2022-08-16T11:21:33Z
updated: 2024-04-03T09:23:09Z
closed: 2022-08-24T21:31:17Z
merged: 2022-08-24T21:31:17Z
base_branch: master
head_branch: optimize-remediation
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/862
---

# Pull Request #862: feat(RHIF-87): get rid of extra calls while loading remediation

**Author**: @mkholjuraev
**Created**: August 16, 2022 at 11:21 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `optimize-remediation`

## Description

Resolves: https://issues.redhat.com/browse/RHIF-53.

There should not be any change in functionality, except for faster remediation loading behavior. Make sure that remediation loads on Advisories and Systems page as expected.

---

## Discussion

### Comment by @adonispuente on August 24, 2022 at 07:10 PM UTC

When reviewing this PR, I used time as a metric to see the performance boost. Im not sure if this is how you tested it, but if there is another metric or way you saw this boosted performance please let me know. 
For testing I grabbed 20 systems, and timed how long it took for the application to load the remediation modal after pressing the remediation button. I did this 5 times in both advisories and systems and then averaged out the numbers.

In master
The average time for Advisories was 24.15 seconds 
For systems it was 6.97 seconds

In this PR
Advisories averaged : 23.87
Systems averaged: 6.99

This PR does make Advisories a little faster, and makes the code cleaner. Before moving forward with review, I just want to be sure if there is another metric I should be using to test efficiency? 

### Comment by @mkholjuraev on August 24, 2022 at 08:05 PM UTC

@adonispuente thank you for the review. You are on the right track in testing the PR. But this PR does not optimize the loading remediation modal when only some number of advisories/systems are remediated. Rather, this optimizes the loading remediation modal when a user selects all of the advisories/systems and tries to load the remediation modal. This is achieved by getting rid of one of the extra time-consuming API requests to fetch entire advisories/systems before fetching remediation pairs. 

Thus, can you please test this scenario?

### Comment by @adonispuente on August 24, 2022 at 08:10 PM UTC

Ahhh ok, retesting now @mkholjuraev 

### Comment by @mkholjuraev on August 24, 2022 at 08:15 PM UTC

@adonispuente I should also mention that sometimes API is throwing errors because of too much time for selecting remediation pair. I have talked to the backend guys, we will work on more optimizations

### Comment by @adonispuente on August 24, 2022 at 09:29 PM UTC

Testing this again with select All (5 times and averaged out), With 236 advisories and 36 systems in both master and this branch the averages are as follows

This PR : 
Advisories : 28.8 seconds
Systems : 8.1 seconds
Master:
Advisories : 31.2 seconds 
Systems : 8.4 seconds

This PR does decrease the average time, LGTM!

### Comment by @mkholjuraev on August 24, 2022 at 09:31 PM UTC

@adonispuente this will be more improved with modifying the  backend SQL query hopefully.

### Comment by @jiridostal on August 24, 2022 at 09:47 PM UTC

:tada: This PR is included in version 1.49.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.49.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.49.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @adonispuente - Approved on August 24, 2022 at 09:29 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/862*
