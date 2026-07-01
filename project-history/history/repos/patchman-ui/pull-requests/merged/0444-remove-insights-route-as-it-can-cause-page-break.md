---
type: pull_request
number: 444
title: "Remove insights route as it can cause page break"
state: merged
author: karelhala
created: 2021-02-18T15:14:19Z
updated: 2021-02-19T10:25:24Z
closed: 2021-02-19T07:49:11Z
merged: 2021-02-19T07:49:11Z
base_branch: master
head_branch: remove-insights-route
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/444
---

# Pull Request #444: Remove insights route as it can cause page break

**Author**: @karelhala
**Created**: February 18, 2021 at 03:14 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `remove-insights-route`

## Description

### Insights route should not be used

Since we are in SPA now InsightsRoute is redundant and only the `Route` from `react-router-dom` can be used.

---

## Discussion

### Comment by @jiridostal on February 19, 2021 at 10:25 AM UTC

:tada: This PR is included in version 1.11.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.11.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.11.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @Hyperkid123 - Commented on February 18, 2021 at 04:08 PM UTC

### Review by @karelhala - Commented on February 18, 2021 at 05:40 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/444*
