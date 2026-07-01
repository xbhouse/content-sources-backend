---
type: pull_request
number: 324
title: "Fixes 2058: Empty url on introspect notification"
state: merged
author: Andrewgdewar
created: 2023-06-29T20:08:16Z
updated: 2023-06-29T20:35:10Z
closed: 2023-06-29T20:35:09Z
merged: 2023-06-29T20:35:09Z
base_branch: main
head_branch: HMS-2058
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/324
---

# Pull Request #324: Fixes 2058: Empty url on introspect notification

**Author**: @Andrewgdewar
**Created**: June 29, 2023 at 08:08 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `HMS-2058`

## Description

## Summary
Fixed a bug that would return an empty URL for the introspect notification, causing the schema to blow up. 
This was a simple DB call order mistake.


---

## Discussion

### Comment by @jlsherrill on June 29, 2023 at 08:30 PM UTC

https://issues.redhat.com/browse/HMS-2058

---

## Reviews

### Review by @rverdile - Approved on June 29, 2023 at 08:32 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/324*
