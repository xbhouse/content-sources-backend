---
type: pull_request
number: 577
title: "Refs 3255: correct mistaken default config value"
state: merged
author: jlsherrill
created: 2024-02-13T19:14:46Z
updated: 2024-02-13T20:19:10Z
closed: 2024-02-13T20:19:07Z
merged: 2024-02-13T20:19:06Z
base_branch: main
head_branch: 3255_2
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/577
---

# Pull Request #577: Refs 3255: correct mistaken default config value

**Author**: @jlsherrill
**Created**: February 13, 2024 at 07:14 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `3255_2`

## Description

## Summary

Fixes a typo in the default value for this config settings.  If the default isn't set, the config reader doesn't actually use the value.

## Testing steps

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on February 13, 2024 at 07:30 PM UTC

https://issues.redhat.com/browse/HMS-3255

### Comment by @jlsherrill on February 13, 2024 at 07:52 PM UTC

merging as its unrelated to the rbac iqe test

---

## Reviews

### Review by @rverdile - Approved on February 13, 2024 at 07:18 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/577*
