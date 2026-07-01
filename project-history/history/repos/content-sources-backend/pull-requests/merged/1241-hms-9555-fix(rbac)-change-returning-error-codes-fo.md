---
type: pull_request
number: 1241
title: "HMS-9555: fix(rbac): change returning error codes for checks"
state: merged
author: TenSt
created: 2025-10-16T14:09:35Z
updated: 2025-10-17T08:59:37Z
closed: 2025-10-17T08:59:37Z
merged: 2025-10-17T08:59:37Z
base_branch: main
head_branch: stepan/HMS-9555-change-returning-status-codes-for-perm-not-found-error
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/1241
---

# Pull Request #1241: HMS-9555: fix(rbac): change returning error codes for checks

**Author**: @TenSt
**Created**: October 16, 2025 at 02:09 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `stepan/HMS-9555-change-returning-status-codes-for-perm-not-found-error`

## Description

## Summary

This PR:
- changes returning codes for missing permissions and header
- changes log levels from "error" to "debug"
- small refactoring of log calls

GlitchTip link: https://glitchtip.devshift.net/insights/issues/3221286/events/69d117004c464159925b08eeddab1c63

---

## Discussion

### Comment by @xbhouse on October 16, 2025 at 02:30 PM UTC

https://issues.redhat.com/browse/HMS-9555

### Comment by @swadeley on October 17, 2025 at 08:44 AM UTC

Hi

the Playwright test that failed is not relevant to the changes in this PR and will be fixed by https://github.com/content-services/content-sources-frontend/pull/688

---

## Reviews

### Review by @swadeley - Approved on October 17, 2025 at 08:45 AM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1241*
