---
type: pull_request
number: 482
title: "Refs 3081: ignore failed intro count for public repos"
state: merged
author: jlsherrill
created: 2023-11-21T20:47:00Z
updated: 2023-11-21T21:24:32Z
closed: 2023-11-21T21:21:57Z
merged: 2023-11-21T21:21:57Z
base_branch: main
head_branch: 3081_2
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/482
---

# Pull Request #482: Refs 3081: ignore failed intro count for public repos

**Author**: @jlsherrill
**Created**: November 21, 2023 at 08:47 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `3081_2`

## Description

## Summary

Found another place where its checking this

## Testing steps

## Checklist

- [x] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on November 21, 2023 at 08:49 PM UTC

I plan to simplify this a bit over on my pr:  https://github.com/content-services/content-sources-backend/pull/445

so the check isn't happening 3 different times

### Comment by @jlsherrill on November 21, 2023 at 09:00 PM UTC

https://issues.redhat.com/browse/HMS-3081

---

## Reviews

### Review by @Andrewgdewar - Approved on November 21, 2023 at 09:03 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/482*
