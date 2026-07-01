---
type: pull_request
number: 549
title: "Fixes 3482: Do not run introspection tests"
state: merged
author: swadeley
created: 2024-01-26T13:50:50Z
updated: 2025-09-08T07:30:40Z
closed: 2024-01-26T14:47:42Z
merged: 2024-01-26T14:47:42Z
base_branch: main
head_branch: swadeley/dont_run_any_introspection_tests
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/549
---

# Pull Request #549: Fixes 3482: Do not run introspection tests

**Author**: @swadeley
**Created**: January 26, 2024 at 01:50 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `swadeley/dont_run_any_introspection_tests`

## Description

## Summary

We do not want to run any full introspection tests in pr checks.
There are simple checks within other CRUD tests.

  there are now two we need to skip
  so use just "introspection" to make it generic.


## Testing steps
Test pass

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @swadeley on January 26, 2024 at 01:56 PM UTC

https://issues.redhat.com/browse/HMS-3482

### Comment by @swadeley on January 26, 2024 at 02:47 PM UTC

Thank you

---

## Reviews

### Review by @jlsherrill - Approved on January 26, 2024 at 01:59 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/549*
