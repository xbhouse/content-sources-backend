---
type: pull_request
number: 149
title: "Fixes 470: Add RBAC integration"
state: merged
author: avisiedo
created: 2022-11-24T22:15:41Z
updated: 2023-02-27T20:00:54Z
closed: 2023-02-27T11:11:12Z
merged: 2023-02-27T11:11:12Z
base_branch: main
head_branch: hmscontent-240-rbac
labels: ["qe-testing-needed", "qe-approved"]
url: https://github.com/content-services/content-sources-backend/pull/149
---

# Pull Request #149: Fixes 470: Add RBAC integration

**Author**: @avisiedo
**Created**: November 24, 2022 at 10:15 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `qe-approved`
**Base**: `main` ← **Head**: `hmscontent-240-rbac`

## Description

- Add Rbac middleware.
- Move middlewares to `pkg/middleware`
- Generate mock for the rbac client.
- Add initial unit test structure.
- Add configurations manage when enable middleware, base url and timeout.
- Refactor ConfigureEcho to the main package to avoid cycle depencies; not happy with this change, open to alternatives.
- add mock service to simulate clients response; this needs a re-thinking.

TODO

- [x] add use cases to cover all the paths for the rbac middleware and helper functions.
- [ ] refine the mock service client at least for rbac service.
- [x] compare with the results from ephemeral environment and update according to them.
- [x] Update deployment.yaml descriptor to consider CLIENTS_RBAC_BASE_URL


---

## Discussion

### Comment by @avisiedo on January 11, 2023 at 10:42 PM UTC

**test specs**

```raw
Scenario 1: check read actions for admin user

Given an admin user
When it list or get a repository
Then the operation is successful

----

Scenario 2: check write actions for admin user

Given an admin user
When it create, update or delete a repository
Then the operation is successful

----

Scenario 3: check read for normal user

Given a normal user (no write permissions)
When it list or retrieve a repository
Then the operation is authorized
----

Scenario 4: check write rejected for normal user

Given a normal user (no write permissions)
When it create, update or delete a repository
Then the operation is unauthorized

----

Scenario 5: check read/write rejected for user with no read nor write permissions

Given a user with no read or write permissions
When it create, update, delete, list or get a repository
Then the operation is unauthorized


```

### Comment by @swadeley on January 23, 2023 at 10:09 AM UTC

/retest

### Comment by @jlsherrill on January 23, 2023 at 09:27 PM UTC

https://issues.redhat.com/browse/HMS-470

### Comment by @swadeley on January 26, 2023 at 02:13 PM UTC

/retest

### Comment by @swadeley on January 31, 2023 at 03:55 PM UTC

/retest

### Comment by @swadeley on February 05, 2023 at 01:23 PM UTC

/retest

### Comment by @swadeley on February 08, 2023 at 09:58 AM UTC

/retest

### Comment by @swadeley on February 08, 2023 at 03:45 PM UTC

/retest

### Comment by @avisiedo on February 08, 2023 at 03:45 PM UTC

/retest

### Comment by @swadeley on February 13, 2023 at 11:49 AM UTC

/retest

### Comment by @swadeley on February 20, 2023 at 03:26 PM UTC

/retest

### Comment by @swadeley on February 21, 2023 at 07:20 PM UTC

/retest

### Comment by @swadeley on February 25, 2023 at 08:57 PM UTC

/retest

---

## Reviews

### Review by @avisiedo - Commented on November 24, 2022 at 10:16 PM UTC

### Review by @jlsherrill - Commented on November 30, 2022 at 03:14 PM UTC

### Review by @jlsherrill - Commented on November 30, 2022 at 03:15 PM UTC

### Review by @jlsherrill - Commented on November 30, 2022 at 03:16 PM UTC

### Review by @jlsherrill - Commented on November 30, 2022 at 03:16 PM UTC

### Review by @jlsherrill - Commented on November 30, 2022 at 03:24 PM UTC

### Review by @jlsherrill - Commented on November 30, 2022 at 03:35 PM UTC

### Review by @jlsherrill - Commented on November 30, 2022 at 03:39 PM UTC

### Review by @avisiedo - Commented on December 02, 2022 at 05:05 PM UTC

### Review by @avisiedo - Commented on December 02, 2022 at 06:07 PM UTC

### Review by @avisiedo - Commented on December 02, 2022 at 06:09 PM UTC

### Review by @avisiedo - Commented on January 10, 2023 at 12:33 PM UTC

### Review by @avisiedo - Commented on January 11, 2023 at 07:46 AM UTC

### Review by @avisiedo - Commented on January 11, 2023 at 08:29 AM UTC

### Review by @avisiedo - Commented on January 11, 2023 at 08:31 AM UTC

### Review by @jlsherrill - Commented on January 16, 2023 at 09:50 PM UTC

### Review by @jlsherrill - Commented on January 16, 2023 at 09:51 PM UTC

### Review by @jlsherrill - Commented on January 16, 2023 at 09:51 PM UTC

### Review by @jlsherrill - Commented on January 16, 2023 at 10:12 PM UTC

### Review by @jlsherrill - Commented on January 16, 2023 at 10:18 PM UTC

### Review by @avisiedo - Commented on January 17, 2023 at 10:19 PM UTC

### Review by @avisiedo - Commented on January 17, 2023 at 10:24 PM UTC

### Review by @avisiedo - Commented on January 18, 2023 at 12:35 AM UTC

### Review by @avisiedo - Commented on January 18, 2023 at 12:41 AM UTC

### Review by @rverdile - Commented on January 18, 2023 at 08:20 PM UTC

### Review by @jlsherrill - Commented on January 18, 2023 at 08:43 PM UTC

### Review by @jlsherrill - Commented on January 18, 2023 at 08:43 PM UTC

### Review by @jlsherrill - Commented on January 18, 2023 at 08:45 PM UTC

### Review by @jlsherrill - Commented on January 18, 2023 at 08:46 PM UTC

### Review by @rverdile - Commented on January 18, 2023 at 08:48 PM UTC

### Review by @rverdile - Commented on January 18, 2023 at 08:57 PM UTC

### Review by @jlsherrill - Commented on January 18, 2023 at 08:58 PM UTC

### Review by @jlsherrill - Commented on January 18, 2023 at 08:59 PM UTC

### Review by @jlsherrill - Commented on January 18, 2023 at 09:01 PM UTC

### Review by @rverdile - Commented on January 18, 2023 at 09:48 PM UTC

### Review by @avisiedo - Commented on January 19, 2023 at 07:57 AM UTC

### Review by @avisiedo - Commented on January 19, 2023 at 08:17 AM UTC

### Review by @avisiedo - Commented on January 19, 2023 at 08:25 AM UTC

### Review by @avisiedo - Commented on January 19, 2023 at 08:36 AM UTC

### Review by @jlsherrill - Approved on January 27, 2023 at 07:47 PM UTC

### Review by @mshriver - Commented on February 01, 2023 at 03:29 PM UTC

### Review by @avisiedo - Commented on February 07, 2023 at 02:47 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/149*
