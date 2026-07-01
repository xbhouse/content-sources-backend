---
type: pull_request
number: 454
title: "Build: Fix title regex to prevent bad multi issues"
state: merged
author: jlsherrill
created: 2023-10-31T17:30:15Z
updated: 2023-11-02T12:25:00Z
closed: 2023-11-02T12:25:00Z
merged: 2023-11-02T12:25:00Z
base_branch: main
head_branch: regex_fix
labels: []
url: https://github.com/content-services/content-sources-backend/pull/454
---

# Pull Request #454: Build: Fix title regex to prevent bad multi issues

**Author**: @jlsherrill
**Created**: October 31, 2023 at 05:30 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `regex_fix`

## Description

## Summary

Tries to prevent "Fixes 123,Fixes 456: summary" titles when doing multiple issues, but keep "Fixes 123,456: summary"

## Testing steps

## Checklist

- [x] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on October 31, 2023 at 05:35 PM UTC


Error: Pull Request title "Fixes 123,Fixes 456: Fix title regex to prevent bad multi issues" failed to pass match regex - /^Build|^((Fixes |Refs )\d+(,\d+)*): .+/

### Comment by @jlsherrill on October 31, 2023 at 05:35 PM UTC

and works fine with "Fixes 123,456: summary"

---

## Reviews

### Review by @Andrewgdewar - Approved on October 31, 2023 at 08:48 PM UTC

Much reg/ex

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/454*
