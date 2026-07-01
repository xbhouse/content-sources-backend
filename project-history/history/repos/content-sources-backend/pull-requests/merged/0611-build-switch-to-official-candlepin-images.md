---
type: pull_request
number: 611
title: "Build: switch to official candlepin images"
state: merged
author: jlsherrill
created: 2024-03-18T15:10:50Z
updated: 2024-03-18T16:51:10Z
closed: 2024-03-18T16:51:10Z
merged: 2024-03-18T16:51:10Z
base_branch: main
head_branch: cp_dev
labels: []
url: https://github.com/content-services/content-sources-backend/pull/611
---

# Pull Request #611: Build: switch to official candlepin images

**Author**: @jlsherrill
**Created**: March 18, 2024 at 03:10 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `cp_dev`

## Description

## Summary

switches to the candlepin:dev-latest image. 

Uses an init file to generate a 2nd db/user for the postgres container

## Testing steps

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Reviews

### Review by @Andrewgdewar - Approved on March 18, 2024 at 03:50 PM UTC

Tested and working!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/611*
