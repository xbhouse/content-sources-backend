---
type: pull_request
number: 891
title: "improve paginated remediation pairs fetching"
state: merged
author: mkholjuraev
created: 2022-10-19T20:00:01Z
updated: 2022-10-26T08:00:38Z
closed: 2022-10-26T07:44:53Z
merged: 2022-10-26T07:44:53Z
base_branch: master
head_branch: master
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/891
---

# Pull Request #891: improve paginated remediation pairs fetching

**Author**: @mkholjuraev
**Created**: October 19, 2022 at 08:00 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `master`

## Description

This PR will improve paginated Remediation issues fetching by 

1. creating grouped API requests with time interval
2. when all items are selected, only /view/advisory/systems endpoint is used. This endpoint behaves the same when all entities are requested and is faster.
3. limit and offset are introduced into the API requests to speed up the SQL query.

Please verify that there is a correct number of API requests to remediate all selected entities and that pagination works as it should.

---

## Discussion

### Comment by @mkholjuraev on October 24, 2022 at 08:34 PM UTC

@adonispuente thank you for the review. I have already started refactoring and preparing remediation issues in the patch for a web worker. This means that all of this work will be done outside of the main thread. Surely, this will optimize the process, however, the remediation app still takes too much time and we can not do something with this from the Patch side.

### Comment by @mkholjuraev on October 26, 2022 at 07:44 AM UTC

@adonispuente thank you. I will create a ticket in remediation Jira board.

### Comment by @jiridostal on October 26, 2022 at 08:00 AM UTC

:tada: This PR is included in version 1.54.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.54.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.54.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @adonispuente - Commented on October 24, 2022 at 07:36 PM UTC

While testing it seems to work as intended. However I noticed that going through the actually remediation process is very slow. When I try to complete remediation, it errors out every time

### Review by @adonispuente - Approved on October 25, 2022 at 05:45 PM UTC

PR works as intended, a ticket needs to be created for the remediation side of things as discussed on call. LGTM!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/891*
