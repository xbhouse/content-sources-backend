---
type: pull_request
number: 237
title: "Fixes 1507: use deprecated security proto config"
state: merged
author: jlsherrill
created: 2023-03-24T16:25:07Z
updated: 2023-03-24T16:43:29Z
closed: 2023-03-24T16:43:29Z
merged: 2023-03-24T16:43:29Z
base_branch: main
head_branch: 1507_1
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/237
---

# Pull Request #237: Fixes 1507: use deprecated security proto config

**Author**: @jlsherrill
**Created**: March 24, 2023 at 04:25 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `1507_1`

## Description

## Summary
https://github.com/content-services/content-sources-backend/commit/3f26de3b5798caf97e2c20e7d4c7bf48d7562cf6 

moved us away from using a now deprecated clowder setting, but it turns out the non-deprecated setting isn't actually populated.  This now checks and tries to use either if they are populated.

## Testing steps
Deploy to stage.

---

## Discussion

### Comment by @jlsherrill on March 24, 2023 at 04:30 PM UTC

https://issues.redhat.com/browse/HMS-1507

---

## Reviews

### Review by @rverdile - Approved on March 24, 2023 at 04:37 PM UTC

ack

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/237*
