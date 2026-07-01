---
type: pull_request
number: 418
title: "Fixes 2588: Existing popular repositories not found error"
state: merged
author: Andrewgdewar
created: 2023-10-05T17:34:16Z
updated: 2023-10-09T08:00:36Z
closed: 2023-10-09T07:52:16Z
merged: 2023-10-09T07:52:16Z
base_branch: main
head_branch: HMS-2588
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/418
---

# Pull Request #418: Fixes 2588: Existing popular repositories not found error

**Author**: @Andrewgdewar
**Created**: October 05, 2023 at 05:34 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `HMS-2588`

## Description

## Summary

Fixes popular repositories not having a UUID added to their list due to a gorm default limit change. 
Previously, when limit was set to 0  (or not set, which defaults to 0) the limit was effectively infinite, the update changes that so that the limit is now actually 0. 
This made it so that on query of the repositoryConfig.List method (with the default Limit of 0), UUID updates for each item on the popular repositories list were not being applied. As "nothing" was found in the Data, despite a count of 1 being returned.

## Testing steps

- Add a popular repository
- Check that the UUID is now returned in the response
- This will also be evident in the UI by the "Remove" button being visible on the popular repositories view.

---

## Discussion

### Comment by @jlsherrill on October 05, 2023 at 06:00 PM UTC

https://issues.redhat.com/browse/HMS-2588

### Comment by @swadeley on October 09, 2023 at 07:52 AM UTC

Hi

Remove buttons work again. Thank you.

---

## Reviews

### Review by @rverdile - Commented on October 05, 2023 at 08:17 PM UTC

### Review by @jlsherrill - Approved on October 06, 2023 at 07:26 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/418*
