---
type: pull_request
number: 187
title: "Fixes 1138: use env variable for count grafana panels"
state: merged
author: rverdile
created: 2023-01-30T17:06:41Z
updated: 2023-01-30T19:30:13Z
closed: 2023-01-30T19:26:03Z
merged: 2023-01-30T19:26:03Z
base_branch: main
head_branch: grafana-panels
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/187
---

# Pull Request #187: Fixes 1138: use env variable for count grafana panels

**Author**: @rverdile
**Created**: January 30, 2023 at 05:06 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `grafana-panels`

## Description

## Summary
The grafana panels for the repository counts were using the wrong datasource. I changed the datasource to the environment variable that holds the value of the datasource selected for the whole dashboard (a dropdown in the outer UI).
## Testing steps
none

---

## Discussion

### Comment by @jlsherrill on January 30, 2023 at 05:30 PM UTC

https://issues.redhat.com/browse/HMS-1138

### Comment by @jlsherrill on January 30, 2023 at 07:25 PM UTC

going ahead and merging, no qe needed.

---

## Reviews

### Review by @jlsherrill - Approved on January 30, 2023 at 07:24 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/187*
