---
type: pull_request
number: 1202
title: "Test: Add test CRUD honours OrgID"
state: merged
author: swadeley
created: 2025-09-16T13:22:32Z
updated: 2025-10-02T13:41:31Z
closed: 2025-10-02T13:41:31Z
merged: 2025-10-02T13:41:31Z
base_branch: main
head_branch: swadeley/check_CRUD_honors_orgid
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1202
---

# Pull Request #1202: Test: Add test CRUD honours OrgID

**Author**: @swadeley
**Created**: September 16, 2025 at 01:22 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `swadeley/check_CRUD_honors_orgid`

## Description

## Summary

[HMS-8886  migrate test_CRUD_honours_org_id](https://issues.redhat.com/browse/HMS-8886)

## Testing steps
tests pass



---

## Discussion

### Comment by @rverdile on September 17, 2025 at 06:02 PM UTC

It looks like there is a test failure related to this changes

### Comment by @swadeley on September 18, 2025 at 09:23 AM UTC

> It looks like there is a test failure related to this changes

Hi @rverdile 
I have added a separate function to get the org id rather than hard code it.  I have now moved that into the `_playwright-tests/test-utils/src/helpers/repositories.ts` file. Should that be a separate commit? To add a helper?

Thank you
EDIT: Was too complicated, removed that code.

### Comment by @xbhouse on October 01, 2025 at 03:32 PM UTC

/retest

---

## Reviews

### Review by @swadeley - Commented on September 18, 2025 at 12:29 PM UTC

### Review by @xbhouse - Commented on September 29, 2025 at 02:38 PM UTC

### Review by @xbhouse - Commented on September 29, 2025 at 05:39 PM UTC

just the one comment about logging, other than that this looks good!

### Review by @swadeley - Commented on October 01, 2025 at 08:28 AM UTC

### Review by @xbhouse - Commented on October 01, 2025 at 03:32 PM UTC

### Review by @xbhouse - Commented on October 01, 2025 at 03:33 PM UTC

### Review by @swadeley - Commented on October 02, 2025 at 08:35 AM UTC

### Review by @swadeley - Commented on October 02, 2025 at 09:09 AM UTC

### Review by @xbhouse - Approved on October 02, 2025 at 01:28 PM UTC

looks good! :) 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1202*
