---
type: pull_request
number: 88
title: "Fixes 115: Create authentication middleware"
state: merged
author: avisiedo
created: 2022-08-25T14:02:55Z
updated: 2022-09-07T19:57:51Z
closed: 2022-09-07T19:57:51Z
merged: 2022-09-07T19:57:51Z
base_branch: main
head_branch: hmscontent-115-wrapper-use-context-x-rh-identity
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/88
---

# Pull Request #88: Fixes 115: Create authentication middleware

**Author**: @avisiedo
**Created**: August 25, 2022 at 02:02 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `hmscontent-115-wrapper-use-context-x-rh-identity`

## Description

- Add usage of go-platform-middleware.
- Add wrapper to reuse the middleware leetting skiping route configurations.
- Add tests for it.
- Update some current tests.


---

## Discussion

### Comment by @jlsherrill on August 25, 2022 at 04:00 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-115

### Comment by @mshriver on September 06, 2022 at 07:51 PM UTC

/retest


### Comment by @jlsherrill on September 07, 2022 at 07:57 PM UTC

going ahead and merging this as i gotta change the routes

---

## Reviews

### Review by @jlsherrill - Commented on August 26, 2022 at 05:26 PM UTC

### Review by @jlsherrill - Commented on August 26, 2022 at 05:55 PM UTC

### Review by @jlsherrill - Commented on August 26, 2022 at 05:59 PM UTC

### Review by @jlsherrill - Commented on August 26, 2022 at 06:17 PM UTC

### Review by @avisiedo - Commented on August 29, 2022 at 01:09 PM UTC

### Review by @jlsherrill - Commented on August 30, 2022 at 01:11 PM UTC

### Review by @jlsherrill - Commented on August 30, 2022 at 01:44 PM UTC

### Review by @avisiedo - Commented on August 30, 2022 at 03:09 PM UTC

### Review by @avisiedo - Commented on August 31, 2022 at 09:09 AM UTC

### Review by @jlsherrill - Approved on September 07, 2022 at 06:25 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/88*
