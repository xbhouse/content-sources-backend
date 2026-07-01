---
type: pull_request
number: 122
title: "Fixes 180: add name and url repo filters"
state: merged
author: jlsherrill
created: 2022-10-13T20:17:10Z
updated: 2022-10-24T13:00:37Z
closed: 2022-10-24T12:42:15Z
merged: 2022-10-24T12:42:14Z
base_branch: main
head_branch: 180
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/122
---

# Pull Request #122: Fixes 180: add name and url repo filters

**Author**: @jlsherrill
**Created**: October 13, 2022 at 08:17 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `180`

## Description

*No description provided*

---

## Discussion

### Comment by @jlsherrill on October 13, 2022 at 08:19 PM UTC

To test:

Create a repository via UI or api such as:
```
curl -X POST http://localhost:8000/api/content-sources/v1/repositories/bulk_create/  -d '[{"name":"fooz", "url":"https://jlsherrill.fedorapeople.org/fake-repos/needed-errata/"}]'  -H "`./scripts/header.sh 1 1`"   -H "Content-Type: application/json" 
```
and then:

```
curl http://localhost:8000/api/content-sources/v1/repositories/?name=fooz  -H "`./scripts/header.sh 1 1`"   -H "Content-Type: application/json"
```
or
```
curl http://localhost:8000/api/content-sources/v1/repositories/?url=https://jlsherrill.fedorapeople.org/fake-repos/needed-errata/  -H "`./scripts/header.sh 1 1`"   -H "Content-Type: application/json"
```


### Comment by @jlsherrill on October 13, 2022 at 08:30 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-180

### Comment by @jlsherrill on October 13, 2022 at 08:41 PM UTC

updated!  just has to wait on https://github.com/content-services/content-sources-backend/pull/121 to go in

---

## Reviews

### Review by @Andrewgdewar - Commented on October 13, 2022 at 08:29 PM UTC

### Review by @Andrewgdewar - Commented on October 13, 2022 at 08:30 PM UTC

### Review by @Andrewgdewar - Commented on October 13, 2022 at 08:30 PM UTC

### Review by @Andrewgdewar - Approved on October 13, 2022 at 08:39 PM UTC

This is working well.

Optional: Update some of the comment text if you prefer the offered alternatives.

### Review by @swadeley - Commented on October 14, 2022 at 07:40 AM UTC

### Review by @swadeley - Commented on October 14, 2022 at 07:41 AM UTC

### Review by @swadeley - Commented on October 14, 2022 at 08:05 AM UTC

### Review by @swadeley - Commented on October 14, 2022 at 08:07 AM UTC

typos and suggestions

### Review by @jlsherrill - Commented on October 14, 2022 at 06:52 PM UTC

### Review by @jlsherrill - Commented on October 14, 2022 at 06:52 PM UTC

### Review by @jlsherrill - Commented on October 14, 2022 at 06:52 PM UTC

### Review by @jlsherrill - Commented on October 14, 2022 at 06:59 PM UTC

### Review by @swadeley - Commented on October 18, 2022 at 10:30 AM UTC

### Review by @swadeley - Commented on October 24, 2022 at 11:06 AM UTC

### Review by @jlsherrill - Commented on October 24, 2022 at 12:37 PM UTC

### Review by @swadeley - Commented on October 24, 2022 at 12:41 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/122*
