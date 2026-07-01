---
type: pull_request
number: 28
title: "CONTENT-68 - support multiple versions for repo config"
state: merged
author: jlsherrill
created: 2022-06-10T15:26:15Z
updated: 2022-06-13T14:33:12Z
closed: 2022-06-13T14:26:55Z
merged: 2022-06-13T14:26:55Z
base_branch: main
head_branch: versions
labels: []
url: https://github.com/content-services/content-sources-backend/pull/28
---

# Pull Request #28: CONTENT-68 - support multiple versions for repo config

**Author**: @jlsherrill
**Created**: June 10, 2022 at 03:26 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `versions`

## Description

*No description provided*

---

## Discussion

### Comment by @jlsherrill on June 10, 2022 at 03:26 PM UTC

simple command to create a repo with multiple versions:
```
http -v localhost:8000/api/content_sources/v1.0/repositories/ name=foo url=http://foo.com/api/content_sources/v1.0/repositories/   distribution_versions:='["8", "9"]'   "`scripts//header.sh 1 2`" 
```

---

## Reviews

### Review by @Andrewgdewar - Approved on June 10, 2022 at 03:28 PM UTC

LGTM, testing with UI now.

### Review by @jlsherrill - Commented on June 10, 2022 at 03:30 PM UTC

### Review by @jlsherrill - Commented on June 10, 2022 at 03:30 PM UTC

### Review by @rverdile - Commented on June 10, 2022 at 03:48 PM UTC

### Review by @jlsherrill - Commented on June 10, 2022 at 03:59 PM UTC

### Review by @rverdile - Commented on June 10, 2022 at 04:12 PM UTC

### Review by @Andrewgdewar - Approved on June 13, 2022 at 01:49 PM UTC

Still looks good, was able to get this up and running locally.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/28*
