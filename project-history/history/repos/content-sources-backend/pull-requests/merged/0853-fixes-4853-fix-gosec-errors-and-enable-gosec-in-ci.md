---
type: pull_request
number: 853
title: "Fixes 4853: fix gosec errors and enable gosec in CI"
state: merged
author: dominikvagner
created: 2024-10-17T11:30:56Z
updated: 2024-10-22T06:32:05Z
closed: 2024-10-22T06:32:05Z
merged: 2024-10-22T06:32:05Z
base_branch: main
head_branch: 4853
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/853
---

# Pull Request #853: Fixes 4853: fix gosec errors and enable gosec in CI

**Author**: @dominikvagner
**Created**: October 17, 2024 at 11:30 AM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `4853`

## Description

## Summary
This PR addresses all problems currently reported by the `gosec` security checker and adds `gosec` to the list of checkers used in `golangci-lint`. Also fixes a problem with miss-matched Go versions for our pre-commit hooks.




---

## Discussion

### Comment by @jlsherrill on October 17, 2024 at 12:00 PM UTC

https://issues.redhat.com/browse/HMS-4853

---

## Reviews

### Review by @jlsherrill - Approved on October 21, 2024 at 02:35 PM UTC

ack nice work!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/853*
