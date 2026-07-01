---
type: pull_request
number: 105
title: "Fixes 166: propery validate repomd exists"
state: merged
author: jlsherrill
created: 2022-09-12T19:14:40Z
updated: 2022-09-13T12:30:45Z
closed: 2022-09-13T12:26:10Z
merged: 2022-09-13T12:26:10Z
base_branch: main
head_branch: 166
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/105
---

# Pull Request #105: Fixes 166: propery validate repomd exists

**Author**: @jlsherrill
**Created**: September 12, 2022 at 07:14 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `166`

## Description

and not just that the passed in url returns 2XX

---

## Discussion

### Comment by @jlsherrill on September 12, 2022 at 07:30 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-166

### Comment by @jlsherrill on September 13, 2022 at 11:51 AM UTC

test instructions:

```
curl -X POST localhost:8000/api/content-sources/v1/repository_parameters/validate/    -H "`./scripts/header.sh 1`"  -H 'Content-Type: application/json'   -d '[{"url":"https://google.com"}]'
```
you'd expect this to return a 404 response, indicating the metadata  is not valid, but prior to this PR, it reports valid metadata and a 200

---

## Reviews

### Review by @Andrewgdewar - Approved on September 12, 2022 at 08:24 PM UTC

This works well. 🚀 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/105*
