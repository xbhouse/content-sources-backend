---
type: pull_request
number: 66
title: "Fixes 122: Add validation to Repository update"
state: merged
author: rverdile
created: 2022-08-01T16:29:47Z
updated: 2022-08-02T14:48:16Z
closed: 2022-08-02T14:48:16Z
merged: 2022-08-02T14:48:16Z
base_branch: main
head_branch: beforeupdate
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/66
---

# Pull Request #66: Fixes 122: Add validation to Repository update

**Author**: @rverdile
**Created**: August 01, 2022 at 04:29 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `beforeupdate`

## Description

Adds validation to updating repositories.

This could also be done using a `BeforeSave` hook that would run before both Update and Create, but keeping the hooks separate seems more future-proof to me.

---

## Discussion

### Comment by @jlsherrill on August 01, 2022 at 04:30 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-122

---

## Reviews

### Review by @swadeley - Approved on August 02, 2022 at 02:48 PM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/66*
