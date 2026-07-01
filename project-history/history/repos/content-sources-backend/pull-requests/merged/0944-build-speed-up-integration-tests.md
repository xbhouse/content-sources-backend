---
type: pull_request
number: 944
title: "Build: speed up integration tests"
state: merged
author: jlsherrill
created: 2025-01-17T18:21:30Z
updated: 2025-01-31T16:47:41Z
closed: 2025-01-31T15:56:12Z
merged: 2025-01-31T15:56:12Z
base_branch: main
head_branch: faster
labels: []
url: https://github.com/content-services/content-sources-backend/pull/944
---

# Pull Request #944: Build: speed up integration tests

**Author**: @jlsherrill
**Created**: January 17, 2025 at 06:21 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `faster`

## Description

## Summary

* By moving the wait for pulp to after candlepin, we hopefully don't have to wait as long
* unsure if we realy need these 'sleep 30' statements, lets try to get rid of them

## Testing steps



---

## Discussion

### Comment by @jlsherrill on January 21, 2025 at 06:54 PM UTC

/retest

### Comment by @jlsherrill on January 24, 2025 at 08:57 PM UTC

one more time to see if we get a 3rd pass in a row

/retest

---

## Reviews

### Review by @xbhouse - Approved on January 29, 2025 at 04:40 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/944*
