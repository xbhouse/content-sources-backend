---
type: pull_request
number: 49
title: "Fixes 116: Add create response object and openAPI spec fixes"
state: merged
author: rverdile
created: 2022-07-06T19:24:15Z
updated: 2022-07-19T20:20:33Z
closed: 2022-07-08T14:04:26Z
merged: 2022-07-08T14:04:26Z
base_branch: main
head_branch: 116
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/49
---

# Pull Request #49: Fixes 116: Add create response object and openAPI spec fixes

**Author**: @rverdile
**Created**: July 06, 2022 at 07:24 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `116`

## Description

Adds:
1. Create returns api.RepositoryResponse object
2. With 201 status, create will return Location header with the url to the resource. I see many sources online saying this is part of the spec, so I thought we might want it.

Fixes:
1. Fetch spec success now says it returns api.RepositoryResponse object
2. Delete spec now says it returns 204 instead of 200
3. Spec examples for distributions_versions were being formatted strangely

old
```
"example": [ 
  "['7'",
  "'8']"
 ]
```

new (I think this is preferred?)
```
"example": [ 
  "7",
  "8"
 ]
```


---

## Discussion

### Comment by @jlsherrill on July 06, 2022 at 07:45 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-116

### Comment by @jlsherrill on July 07, 2022 at 05:47 PM UTC

reminder, hold off on merging until qe reviews :)

### Comment by @jlsherrill on July 19, 2022 at 08:20 PM UTC

https://issues.redhat.com/browse/14733471

---

## Reviews

### Review by @jlsherrill - Approved on July 07, 2022 at 05:47 PM UTC

### Review by @swadeley - Approved on July 08, 2022 at 02:04 PM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/49*
