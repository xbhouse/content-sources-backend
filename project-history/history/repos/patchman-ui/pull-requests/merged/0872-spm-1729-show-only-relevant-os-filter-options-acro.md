---
type: pull_request
number: 872
title: "SPM-1729: Show only relevant OS filter options across tables"
state: merged
author: gkarat
created: 2022-09-15T13:30:35Z
updated: 2022-09-26T08:06:25Z
closed: 2022-09-26T07:50:16Z
merged: 2022-09-26T07:50:16Z
base_branch: master
head_branch: spm-1729-dynamic-os-version-values
labels: ["enhancement", "released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/872
---

# Pull Request #872: SPM-1729: Show only relevant OS filter options across tables

**Author**: @gkarat
**Created**: September 15, 2022 at 01:30 PM UTC
**Status**: Merged
**Labels**: `enhancement`, `released`
**Base**: `master` ← **Head**: `spm-1729-dynamic-os-version-values`

## Description

Implements https://issues.redhat.com/browse/SPM-1729.

This changes the behavior of the OS filter across tables. The filter now only shows the options that are relevant and available for the user's account. The options are requested from the Inventory API.

## How to test

1. Check the following pages and components to contain the same values for the OS filter:
1.1. https://stage.foo.redhat.com:1337/beta/insights/patch/systems
1.2. https://stage.foo.redhat.com:1337/beta/insights/patch/advisories/${advisory_id}
1.3. https://stage.foo.redhat.com:1337/beta/insights/patch/templates: create patch template wizard component.
2. Check whether a call to the Inventory API's /operating_system is made in all pages.
3. Check chips/filter selection behavior and URL parameters update.

## Screenshots

![Screenshot 2022-09-15 at 15-27-15 Patch template - Patch Red Hat Insights](https://user-images.githubusercontent.com/31385370/190415978-caffe7d1-73f1-492e-94a9-82a16026195f.png)
![Screenshot 2022-09-15 at 15-26-53 FEDORA-EPEL-2022-929440e4dc - Advisories - Patch Red Hat Insights](https://user-images.githubusercontent.com/31385370/190415985-1bdea00b-26d1-42cb-96d1-ae55b247f237.png)
![Screenshot 2022-09-15 at 15-26-19 Systems - Patch Red Hat Insights](https://user-images.githubusercontent.com/31385370/190415990-1f4243de-0510-409f-82b4-d561073ff422.png)


---

## Discussion

### Comment by @jiridostal on September 26, 2022 at 08:06 AM UTC

:tada: This PR is included in version 1.51.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.51.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.51.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @mkholjuraev - Changes Requested on September 19, 2022 at 10:15 AM UTC

Overall code looks good. But, I would love to see a few duplicated code to be removed by making created OS filter hook more dynamic.  (I am checking how much effort is needed to use the default `operatingSystem` filter from [the Inventory table component](https://github.com/RedHatInsights/insights-inventory-frontend/blob/master/src/components/InventoryTable/EntityTableToolbar.js#L133))

### Review by @gkarat - Commented on September 20, 2022 at 03:04 PM UTC

### Review by @mkholjuraev - Approved on September 20, 2022 at 03:25 PM UTC

Looks perfect. @gkarat thank you! Did not run it though.  

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/872*
