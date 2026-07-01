---
type: pull_request
number: 170
title: "Fixes 326: Filter repositories by status"
state: merged
author: rverdile
created: 2023-01-09T15:09:49Z
updated: 2023-01-16T19:30:37Z
closed: 2023-01-16T19:01:23Z
merged: 2023-01-16T19:01:23Z
base_branch: main
head_branch: status-filter
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/170
---

# Pull Request #170: Fixes 326: Filter repositories by status

**Author**: @rverdile
**Created**: January 09, 2023 at 03:09 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `status-filter`

## Description

Adds filtering repository list by status e.g. `/api/content-sources/v1/repositories/?status=Pending,Valid`

To test, create repositories with different introspection statuses and try filtering.

---

## Discussion

### Comment by @jlsherrill on January 09, 2023 at 03:30 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-326

### Comment by @swadeley on January 16, 2023 at 04:08 PM UTC

/retest

### Comment by @swadeley on January 16, 2023 at 07:01 PM UTC

Hi, 

works as expected
e.g.:

In [11]: app.content_sources.rest_client.repositories_api.list_repositories(status='Pending,Valid')
2023-01-16 18:59:40.400 [    INFO] [iqe.base.rest_client] REST: METHOD=GET, request_id=gcnjIJMln1l22D60z7yyn5yW8pZ8wCod, params=[('status', 'Pending,Valid')]


---

## Reviews

### Review by @rverdile - Commented on January 09, 2023 at 03:10 PM UTC

### Review by @Andrewgdewar - Commented on January 12, 2023 at 03:16 PM UTC

### Review by @Andrewgdewar - Approved on January 12, 2023 at 05:38 PM UTC

Tested both via postman and front-end, working well. 


---

*Archived from: https://github.com/content-services/content-sources-backend/pull/170*
