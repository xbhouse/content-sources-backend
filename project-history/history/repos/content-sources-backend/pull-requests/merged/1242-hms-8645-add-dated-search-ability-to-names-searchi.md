---
type: pull_request
number: 1242
title: "HMS-8645: add dated search ability to names searching APIs"
state: merged
author: dominikvagner
created: 2025-10-17T09:32:15Z
updated: 2025-10-24T11:16:54Z
closed: 2025-10-24T11:16:54Z
merged: 2025-10-24T11:16:54Z
base_branch: main
head_branch: fix-repeatable-build-package-search
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/1242
---

# Pull Request #1242: HMS-8645: add dated search ability to names searching APIs

**Author**: @dominikvagner
**Created**: October 17, 2025 at 09:32 AM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `fix-repeatable-build-package-search`

## Description

## Summary
This PR adds necessary changes for addressing the broken package search for repeatable build in Image Builder.

This adds a new 'date' argument to names searching APIs (packages, package groups, environments), which allows for searching the names at the specified date, without the client needing to provide a list of snapshot UUIDs.
The ability to do this will alleviate burden from the IB's package search step that is already way too complicated. It's much easier to solve this here than in the FE, where you have to adhere to React hook rules. 

Adds tests for the new feature and pulls some of the required logic into helpers/utils. Also adds a new doc section for testing work spanning this repo and the IB FE. Small adjustments to Playwright test for more stable local testing.

Issue Link: [[HMS-8645](https://issues.redhat.com/browse/HMS-8645)]

## Testing steps
1. Verify the name searching APIs work as expected (with and without a date specified, with and without package sources)
1.1 Test the RPM search API: `/rpms/names`
1.2 Test the Package Groups search API: `/package_groups/names`
1.3 Test the Environments search API: `/environments/names`
2. Spin up the Image Builder Frontend (from this [PR](https://github.com/osbuild/image-builder-frontend/pull/3731)) locally together with this repo (following the docs added in this PR, and verify that the package search and package groups search work as expected in the packages step in the IB's wizard.
2.1 Create a repo with 2 snapshots (good one to use are [these](https://content-services.github.io/fixtures/yum/comps-modules/) ones).
2.2 Verify the package search show the correct stuff depending on the option and date chosen in the "Repeatable Build" step.

---

## Discussion

### Comment by @xbhouse on October 17, 2025 at 10:00 AM UTC

https://issues.redhat.com/browse/HMS-8645

### Comment by @xbhouse on October 20, 2025 at 05:54 PM UTC

tested with your IB frontend PR, works well! :tada: thanks for adding in those instructions :) just a couple small comments

### Comment by @swadeley on October 23, 2025 at 08:01 AM UTC

/retest

---

## Reviews

### Review by @xbhouse - Commented on October 20, 2025 at 05:47 PM UTC

### Review by @xbhouse - Commented on October 20, 2025 at 05:50 PM UTC

### Review by @dominikvagner - Commented on October 22, 2025 at 08:24 AM UTC

### Review by @dominikvagner - Commented on October 22, 2025 at 08:25 AM UTC

### Review by @xbhouse - Commented on October 22, 2025 at 01:11 PM UTC

### Review by @xbhouse - Approved on October 22, 2025 at 01:13 PM UTC

lgtm! nice work :rocket: 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1242*
