---
type: pull_request
number: 1426
title: "Test: Add resolveAuthStorageRoot helper"
state: merged
author: swadeley
created: 2026-03-25T18:29:28Z
updated: 2026-03-30T13:23:40Z
closed: 2026-03-30T13:23:40Z
merged: 2026-03-30T13:23:40Z
base_branch: main
head_branch: swadeley/fix_path_to_auth_storage
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1426
---

# Pull Request #1426: Test: Add resolveAuthStorageRoot helper

**Author**: @swadeley
**Created**: March 25, 2026 at 06:29 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `swadeley/fix_path_to_auth_storage`

## Description

## Summary

Add resolveAuthStorageRoot helper
Resolves path to auth files and sets PLAYWRIGHT_AUTH_DIR.

## Testing steps

Integration tests pass and no more "Failed to read token from file ADMIN_TOKEN.json" errors.

#testwith https://github.com/content-services/content-sources-frontend/pull/928

---

## Discussion

### Comment by @dominikvagner on March 26, 2026 at 03:44 PM UTC

I am still seeing: `Failed to read token from file ADMIN_TOKEN.json` in the Playwright logs on this PR 🥲🤔

don't see it in the frontend test runs, but here it still seems to persist 💭 

### Comment by @swadeley on March 26, 2026 at 07:25 PM UTC

> I am still seeing: `Failed to read token from file ADMIN_TOKEN.json` in the Playwright logs on this PR 🥲🤔
> 
> don't see it in the frontend test runs, but here it still seems to persist 💭

Hi @dominikvagner , I think the issue is that tests that don't use "page"  use IDENTITY_HEADER for API commands. So doing browser  JWT refresh is not going to work and is not required for pure API tests.

Fix coming soon.

Thank you

---

## Reviews

### Review by @dominikvagner - Approved on March 30, 2026 at 12:41 PM UTC

still seeing the `failed to load ...` messages in the frontend tests run on this PR, but that's caused by the incorrect frontend "branch" used 😅
_the backend workflow doesn't check for testwith in it's PR description, but looks for a PR in our frontend repo that has a testwith of the backend PR number 💭_ 

didn't see it in the frontend PR when it had the testwith link in it, so I guess it's good 😄 
ack! ✅ 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1426*
