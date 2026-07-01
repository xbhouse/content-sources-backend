---
type: pull_request
number: 239
title: "Refs 1203: adjust message latency bucket sizes"
state: merged
author: jlsherrill
created: 2023-03-28T17:27:25Z
updated: 2023-03-28T19:02:32Z
closed: 2023-03-28T19:02:28Z
merged: 2023-03-28T19:02:28Z
base_branch: main
head_branch: 1203_5
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/239
---

# Pull Request #239: Refs 1203: adjust message latency bucket sizes

**Author**: @jlsherrill
**Created**: March 28, 2023 at 05:27 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `1203_5`

## Description

## Summary
The bucket sizes weren't conducive to writing alerts or really even defining the SLO, this increases them greatly.

## Testing steps
None needed?  check the metrics after creating a repo maybe?


---

## Discussion

### Comment by @jlsherrill on March 28, 2023 at 05:30 PM UTC

https://issues.redhat.com/browse/HMS-1203

---

## Reviews

### Review by @rverdile - Approved on March 28, 2023 at 05:56 PM UTC

tested and lgtm!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/239*
