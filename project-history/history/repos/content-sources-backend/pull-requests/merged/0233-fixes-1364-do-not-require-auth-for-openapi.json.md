---
type: pull_request
number: 233
title: "Fixes 1364: do not require auth for openapi.json"
state: merged
author: jlsherrill
created: 2023-03-21T15:41:06Z
updated: 2023-03-21T16:30:17Z
closed: 2023-03-21T16:02:10Z
merged: 2023-03-21T16:02:10Z
base_branch: main
head_branch: 1364
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/233
---

# Pull Request #233: Fixes 1364: do not require auth for openapi.json

**Author**: @jlsherrill
**Created**: March 21, 2023 at 03:41 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `1364`

## Description

## Summary
Stop requiring auth for requests to openapi.json

## Testing steps
run server
run: ```curl localhost:8000/api/content-sources/v1/openapi.json```


---

## Discussion

### Comment by @jlsherrill on March 21, 2023 at 03:58 PM UTC

yeah, i think /ping is now only available at  /ping and not  /api/content-sources/v1/ping anymore

This is expected

### Comment by @jlsherrill on March 21, 2023 at 04:00 PM UTC

https://issues.redhat.com/browse/HMS-1364

---

## Reviews

### Review by @Andrewgdewar - Approved on March 21, 2023 at 03:49 PM UTC

Works well. One note, ping doesn't seem to be working for me locally but that may be something else.

### Review by @swadeley - Approved on March 21, 2023 at 04:02 PM UTC

LGTM

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/233*
