---
type: pull_request
number: 307
title: "Fixes 1924: move notifications loading to startup"
state: merged
author: jlsherrill
created: 2023-06-08T13:28:18Z
updated: 2023-06-13T01:11:10Z
closed: 2023-06-13T01:09:31Z
merged: 2023-06-13T01:09:31Z
base_branch: main
head_branch: 1924
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/307
---

# Pull Request #307: Fixes 1924: move notifications loading to startup

**Author**: @jlsherrill
**Created**: June 08, 2023 at 01:28 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `1924`

## Description

## Summary
Notifications Client loading spawns a go routine which seemed to cause an issue with worker shutdown in tests.

## Testing steps
Tests pass, notifications still are sent.

---

## Discussion

### Comment by @jlsherrill on June 12, 2023 at 01:36 PM UTC

https://issues.redhat.com/browse/HMS-1924

---

## Reviews

### Review by @Andrewgdewar - Approved on June 09, 2023 at 03:07 PM UTC

Tested, working the same as prior to this change. 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/307*
