---
type: pull_request
number: 17
title: "CONTENT-54: Support delete of repositories"
state: merged
author: jlsherrill
created: 2022-05-26T12:44:04Z
updated: 2022-05-31T16:00:37Z
closed: 2022-05-31T16:00:37Z
merged: 2022-05-31T16:00:37Z
base_branch: main
head_branch: delete
labels: []
url: https://github.com/content-services/content-sources-backend/pull/17
---

# Pull Request #17: CONTENT-54: Support delete of repositories

**Author**: @jlsherrill
**Created**: May 26, 2022 at 12:44 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `delete`

## Description

*No description provided*

---

## Discussion

### Comment by @jlsherrill on May 26, 2022 at 01:22 PM UTC

For testing:
```
go run cmd/dbmigrate/main.go seed
go run cmd/content-sources/main.go
curl http://localhost/api/content_sources/v1.0/respoitories/ | python -m json.tool   #grab a uuid, notice the total count
curl -X DELETE http://localhost/api/content_sources/v1.0/respoitories/UUID
curl http://localhost/api/content_sources/v1.0/respoitories/ | python -m json.tool   #total count should be 1 less
```

---

## Reviews

### Review by @Andrewgdewar - Approved on May 31, 2022 at 03:49 PM UTC

Ack, Much Delete.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/17*
