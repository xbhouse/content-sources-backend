---
type: pull_request
number: 213
title: "Fixes 1249: Add sonarQube scanning"
state: merged
author: jlsherrill
created: 2023-02-15T14:14:26Z
updated: 2023-02-15T15:13:58Z
closed: 2023-02-15T15:09:49Z
merged: 2023-02-15T15:09:49Z
base_branch: main
head_branch: 1249
labels: []
url: https://github.com/content-services/content-sources-backend/pull/213
---

# Pull Request #213: Fixes 1249: Add sonarQube scanning

**Author**: @jlsherrill
**Created**: February 15, 2023 at 02:14 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `1249`

## Description

## Summary

## Testing steps


---

## Discussion

### Comment by @jlsherrill on February 15, 2023 at 02:30 PM UTC

https://issues.redhat.com/browse/HMS-1249

### Comment by @mshriver on February 15, 2023 at 02:30 PM UTC

/retest

### Comment by @mshriver on February 15, 2023 at 02:40 PM UTC

@jlsherrill the scan script is causing the ci.int.devshift failure:
```09:35:56 14:35:56.229 ERROR: Error during SonarScanner execution
09:35:56 java.lang.IllegalStateException: Failed to upload report: You're not authorized to run analysis. Please contact the project administrator.```

### Comment by @ezr-ondrej on February 15, 2023 at 03:01 PM UTC

/retest

---

## Reviews

### Review by @mshriver - Commented on February 15, 2023 at 02:41 PM UTC

Intent wise and script wise this looks fine, but the scan isn't actually passing

### Review by @mshriver - Approved on February 15, 2023 at 02:41 PM UTC

Intent wise and script wise this looks fine, but the scan isn't actually passing

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/213*
