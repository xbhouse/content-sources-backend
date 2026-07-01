---
type: pull_request
number: 1342
title: "Test: update playwright deps, remove custom epel/popular repo test"
state: merged
author: xbhouse
created: 2025-12-16T15:20:02Z
updated: 2025-12-17T19:50:47Z
closed: 2025-12-17T19:50:47Z
merged: 2025-12-17T19:50:46Z
base_branch: main
head_branch: update-pw-setup
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1342
---

# Pull Request #1342: Test: update playwright deps, remove custom epel/popular repo test

**Author**: @xbhouse
**Created**: December 16, 2025 at 03:20 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `update-pw-setup`

## Description

## Summary

- Removes caching for playwright browsers and updates the package install and playwright browser install commands to match the frontend CI
- Updates the yarn install commands to specify separate cache directories for the frontend and backend installs
  - It seems that, with multiple yarn installs running in parallel, we were running into a race condition when both processes tried to modify the global yarn cache. This resulted in intermittent errors like these: `ENOTEMPTY: directory not empty...'` and/or `Extracting tar content of undefined failed, the file appears to be corrupt: "ENOENT: no such file or directory...'"`
- Removes the popular repo test that creates and deletes a custom EPEL repository as this isn't allowed anymore with the introduction of community epel repos, and we have tests for the community epel repos

## Testing steps

Tests pass, specifically the `Setup backend, frontend, and playwright tests` step in the playwright-tests job

---

## Reviews

### Review by @TenSt - Approved on December 17, 2025 at 09:59 AM UTC

thanks 👍 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1342*
