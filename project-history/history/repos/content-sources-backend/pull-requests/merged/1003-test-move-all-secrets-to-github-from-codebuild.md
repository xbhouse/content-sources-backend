---
type: pull_request
number: 1003
title: "Test: move all secrets to github from codebuild"
state: merged
author: xbhouse
created: 2025-03-03T19:17:03Z
updated: 2025-03-06T20:19:44Z
closed: 2025-03-06T20:19:44Z
merged: 2025-03-06T20:19:44Z
base_branch: main
head_branch: 5531-move-secrets
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1003
---

# Pull Request #1003: Test: move all secrets to github from codebuild

**Author**: @xbhouse
**Created**: March 03, 2025 at 07:17 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `5531-move-secrets`

## Description

## Summary

Switches secrets to be referenced from Github instead of Codebuild

## Testing steps

- PW tests pass (triggered manually on this branch and passed [here](https://github.com/content-services/content-sources-backend/actions/runs/13638994392))
- No secrets being pulled from Codebuild (these look like `$FOO` instead of `${{ secrets.FOO }}`)

---

## Reviews

### Review by @Andrewgdewar - Approved on March 06, 2025 at 08:19 PM UTC

Looks good! Thanks for this @xbhouse!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1003*
