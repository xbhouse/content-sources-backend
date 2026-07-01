---
type: pull_request
number: 362
title: "Fixes 2356: rename s3 bucket"
state: closed
author: jlsherrill
created: 2023-08-15T18:53:52Z
updated: 2023-08-16T12:23:19Z
closed: 2023-08-16T12:23:16Z
base_branch: main
head_branch: 2356
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/362
---

# Pull Request #362: Fixes 2356: rename s3 bucket

**Author**: @jlsherrill
**Created**: August 15, 2023 at 06:53 PM UTC
**Status**: Closed
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `2356`

## Description

## Summary
We need to rename the bucket resource we consume, due to previously using an unversioned bucket.

## Testing steps
tests pass

---

## Discussion

### Comment by @jlsherrill on August 15, 2023 at 06:56 PM UTC

Note, this requires https://gitlab.cee.redhat.com/service/app-interface/-/merge_requests/77491 to go in first

### Comment by @jlsherrill on August 15, 2023 at 07:00 PM UTC

https://issues.redhat.com/browse/HMS-2356

### Comment by @mshriver on August 16, 2023 at 11:55 AM UTC

@jlsherrill with https://gitlab.cee.redhat.com/service/app-interface/-/merge_requests/77513, is this still needed?

### Comment by @jlsherrill on August 16, 2023 at 12:23 PM UTC

nope!  not anymore

---

## Reviews

### Review by @mshriver - Approved on August 15, 2023 at 06:59 PM UTC

Can see this getting parametrized in the future and defined within SRE

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/362*
