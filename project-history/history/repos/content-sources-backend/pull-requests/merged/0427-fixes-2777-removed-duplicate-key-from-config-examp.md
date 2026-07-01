---
type: pull_request
number: 427
title: "Fixes 2777: removed duplicate key from config example"
state: merged
author: xbhouse
created: 2023-10-12T16:37:16Z
updated: 2023-12-13T22:02:43Z
closed: 2023-10-12T17:15:07Z
merged: 2023-10-12T17:15:06Z
base_branch: main
head_branch: duplicate-pulp-key-fix
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/427
---

# Pull Request #427: Fixes 2777: removed duplicate key from config example

**Author**: @xbhouse
**Created**: October 12, 2023 at 04:37 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `duplicate-pulp-key-fix`

## Description

## Summary

Removed duplicate pulp key in config.yaml.example that was causing issues when running the app for the first time.

## Testing steps


---

## Discussion

### Comment by @app-sre-bot on October 12, 2023 at 04:38 PM UTC

Can one of the admins verify this patch?

### Comment by @jlsherrill on October 12, 2023 at 05:00 PM UTC

https://issues.redhat.com/browse/HMS-2777

### Comment by @rverdile on October 12, 2023 at 05:15 PM UTC

merging, QE wouldn't need to look at this

---

## Reviews

### Review by @rverdile - Approved on October 12, 2023 at 05:06 PM UTC

nice!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/427*
