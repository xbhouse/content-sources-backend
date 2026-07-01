---
type: pull_request
number: 283
title: "Fixes 602: Revert previous commit (Glitchtip)"
state: merged
author: Andrewgdewar
created: 2023-05-23T20:36:10Z
updated: 2023-05-23T20:42:26Z
closed: 2023-05-23T20:42:26Z
merged: 2023-05-23T20:42:26Z
base_branch: main
head_branch: HMS-602-1
labels: []
url: https://github.com/content-services/content-sources-backend/pull/283
---

# Pull Request #283: Fixes 602: Revert previous commit (Glitchtip)

**Author**: @Andrewgdewar
**Created**: May 23, 2023 at 08:36 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `HMS-602-1`

## Description

## Summary

Previous commit caused an error to do with an unfound dummy kafka ingress. 
This reverts that and prevents the error, it will likely not fix the failing notifications.

## Testing steps


---

*Archived from: https://github.com/content-services/content-sources-backend/pull/283*
