---
type: pull_request
number: 906
title: "(SPM-1752): RBAC implementation | Templates Page permissions"
state: merged
author: adonispuente
created: 2022-11-04T00:12:13Z
updated: 2022-11-17T15:50:00Z
closed: 2022-11-17T15:32:34Z
merged: 2022-11-17T15:32:34Z
base_branch: master
head_branch: SPM-1752
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/906
---

# Pull Request #906: (SPM-1752): RBAC implementation | Templates Page permissions

**Author**: @adonispuente
**Created**: November 04, 2022 at 12:12 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `SPM-1752`

## Description

In accordance with https://issues.redhat.com/browse/SPM-1768
This adds RBAC permissions to patch. If navigating to Templates page with only patch viewer permissions, a user cannot select the 'Create a Template' button and a tooltip hovers over the button. The same goes for editing a template.

---

## Discussion

### Comment by @adonispuente on November 10, 2022 at 09:43 PM UTC

@bastilian @mkholjuraev Regarding the breadcrumbs, it seems passing in paths.path is broken. 
It will resolve by either having a hard string, are by slapping in something like paths[0].path
You've both had different suggestions for both so I'd like to come to an agreement.

I actually prefer having strings in there, or I can make constants and put them there as well if thats something you guys would prefer. Reverting back to strings also resolves the snapshot 'undefined' issue in the 'to' field. Im not sure how exactly its possible to grab path dynamically.
With paths.path, I realize that its the wrong path entirely. The first breadcrumb is supposed to take you back to the original page.  So packagedetail 1st breadcrumb should take you to /package, advisoryDetail takes you the /advisory and so on. Even if paths.paths is grabbed properly, thats not where we want the to field to be 

### Comment by @mkholjuraev on November 10, 2022 at 09:46 PM UTC

@adonispuente I also opt for the string rather than paths[0].path. 

### Comment by @jiridostal on November 17, 2022 at 03:49 PM UTC

:tada: This PR is included in version 1.56.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.56.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.56.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @bastilian - Commented on November 04, 2022 at 08:54 AM UTC

### Review by @mkholjuraev - Commented on November 04, 2022 at 11:31 AM UTC

### Review by @adonispuente - Commented on November 04, 2022 at 01:13 PM UTC

### Review by @bastilian - Commented on November 07, 2022 at 11:28 AM UTC

### Review by @bastilian - Commented on November 09, 2022 at 08:31 AM UTC

### Review by @bastilian - Commented on November 09, 2022 at 08:32 AM UTC

### Review by @bastilian - Commented on November 09, 2022 at 08:34 AM UTC

### Review by @bastilian - Commented on November 09, 2022 at 08:34 AM UTC

### Review by @bastilian - Commented on November 09, 2022 at 08:35 AM UTC

### Review by @bastilian - Commented on November 09, 2022 at 08:37 AM UTC

### Review by @bastilian - Commented on November 09, 2022 at 08:37 AM UTC

### Review by @adonispuente - Commented on November 09, 2022 at 08:35 PM UTC

### Review by @adonispuente - Commented on November 09, 2022 at 08:39 PM UTC

### Review by @adonispuente - Commented on November 09, 2022 at 08:42 PM UTC

### Review by @adonispuente - Commented on November 09, 2022 at 09:03 PM UTC

### Review by @adonispuente - Commented on November 09, 2022 at 09:19 PM UTC

### Review by @mkholjuraev - Commented on November 10, 2022 at 01:04 PM UTC

### Review by @mkholjuraev - Approved on November 14, 2022 at 09:51 AM UTC

LGTM! tested locally using account `psegedy_patch` which has only read permission. For accounts with full permission, everything seems to be as described.

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/906*
