---
type: pull_request
number: 339
title: "feat(OUIA): enable OUIA for Patch"
state: merged
author: mkholjuraev
created: 2020-11-24T12:24:35Z
updated: 2021-08-09T06:55:52Z
closed: 2020-11-25T13:22:11Z
merged: 2020-11-25T13:22:11Z
base_branch: master
head_branch: enable-ouia
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/339
---

# Pull Request #339: feat(OUIA): enable OUIA for Patch

**Author**: @mkholjuraev
**Created**: November 24, 2020 at 12:24 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `enable-ouia`

## Description

Partial work for: [SPM-580](https://issues.redhat.com/browse/SPM-580)

What components have OUIA enabled:
1.  Tables not including Inventory. e.g `data-ouia-component-id="advisories-table"`
2.  actionsConfig of PrimaryToolbar. e.g remediation button `data-ouia-component-id="toolbar-remediation-button"`  
3. bulkSelect toggle button of PrimaryToolbar.  `data-ouia-component-type="bulk-select-toggle-button"`
4. Header of all pages. e.g `data-ouia-component-type="advisory-details-page-header"`
5. Header breadcrumbs. e.g `data-ouia-component-type="advisory-details-breadcrumb" ` `data-ouia-component-id="breadcrumb-to-Advisories"` 
6. Header tabs. e.g `data-ouia-component-type="advisory-details-breadcrumb" ` `data-ouia-component-id="breadcrumb-to-Advisories"`

What is missing : 
1. Primary toolbar bulk select items.
2. Primary toolbar active filter config chips.
3. Primary toolbar filter config dropdown items.
4. OUIA for Inventory table
5. OUIA for pagination button 


---

## Discussion

### Comment by @jiridostal on December 01, 2020 at 01:03 PM UTC

:tada: This PR is included in version 1.0.3 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.0.3)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.0.3)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/339*
