---
type: pull_request
number: 972
title: "HMS-5429: add popular repo test and api client"
state: merged
author: jlsherrill
created: 2025-02-10T21:04:17Z
updated: 2025-02-11T20:23:24Z
closed: 2025-02-11T20:15:15Z
merged: 2025-02-11T20:15:15Z
base_branch: main
head_branch: HMS-5429
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/972
---

# Pull Request #972: HMS-5429: add popular repo test and api client

**Author**: @jlsherrill
**Created**: February 10, 2025 at 09:04 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `HMS-5429`

## Description

## Summary

* change 'make openapi' to generate a new typescript client in for playwright tests
* changes the features api test to use the new client
* adds popular repo api test
* Adds a 'negative' template test with error check

## Testing steps

code review

---

## Discussion

### Comment by @jlsherrill on February 10, 2025 at 09:30 PM UTC

https://issues.redhat.com/browse/HMS-5429

### Comment by @mayurilahane on February 11, 2025 at 08:23 PM UTC

- Basically ran all API tests mentioned in this PR under API folder and all are passing 



---

## Reviews

### Review by @swadeley - Commented on February 11, 2025 at 04:15 PM UTC

### Review by @Andrewgdewar - Approved on February 11, 2025 at 04:25 PM UTC

Likely need a linter, as you have a number of unneeded imports and have formatting mistakes. 

But other than that this works and looks good.

### Review by @jlsherrill - Commented on February 11, 2025 at 06:52 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/972*
