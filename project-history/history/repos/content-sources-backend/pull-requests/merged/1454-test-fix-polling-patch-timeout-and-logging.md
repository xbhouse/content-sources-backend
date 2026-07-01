---
type: pull_request
number: 1454
title: "Test: Fix polling patch timeout and logging"
state: merged
author: swadeley
created: 2026-04-08T16:25:49Z
updated: 2026-04-14T12:00:08Z
closed: 2026-04-14T12:00:08Z
merged: 2026-04-14T12:00:08Z
base_branch: main
head_branch: swadeley/fix_getTemplateSystemsCount
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1454
---

# Pull Request #1454: Test: Fix polling patch timeout and logging

**Author**: @swadeley
**Created**: April 08, 2026 at 04:25 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `swadeley/fix_getTemplateSystemsCount`

## Description



## Summary
Fix polling patch timeout and logging
Fix INVENTORY_PATCH_POLL_TIMEOUT_MS and make logs less verbose.

Add patchTemplateSystems helpers
## Testing steps

_playwright-tests/Integration/AssociatedTemplateCRUD.spec.ts passes

Blocks: [HMS-10435: Use patch API to remove system in AssociatedTemplateCRUD](https://github.com/content-services/content-sources-frontend/pull/961)

---

## Discussion

### Comment by @swadeley on April 10, 2026 at 07:49 PM UTC

/retest-playwright

---

## Reviews

### Review by @TenSt - Approved on April 14, 2026 at 11:30 AM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1454*
