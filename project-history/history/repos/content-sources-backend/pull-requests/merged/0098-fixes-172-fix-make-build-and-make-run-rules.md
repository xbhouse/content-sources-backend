---
type: pull_request
number: 98
title: "Fixes 172: fix make build and make run rules"
state: merged
author: avisiedo
created: 2022-09-06T12:25:58Z
updated: 2022-09-12T12:42:15Z
closed: 2022-09-12T12:42:15Z
merged: 2022-09-12T12:42:15Z
base_branch: main
head_branch: hmscontent-172-enhance-make-build-run
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/98
---

# Pull Request #98: Fixes 172: fix make build and make run rules

**Author**: @avisiedo
**Created**: September 06, 2022 at 12:25 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `hmscontent-172-enhance-make-build-run`

## Description

When a source code is updated or just touched, now `make build` and `make run` detect it has changed and the binaries are rerebuilt.

---

## Discussion

### Comment by @jlsherrill on September 06, 2022 at 01:00 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-172

---

## Reviews

### Review by @jlsherrill - Approved on September 12, 2022 at 12:31 PM UTC

ACK, now rebuilds on any changes!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/98*
