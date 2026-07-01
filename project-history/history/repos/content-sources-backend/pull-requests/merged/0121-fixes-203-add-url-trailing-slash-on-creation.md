---
type: pull_request
number: 121
title: "Fixes 203: Add URL trailing slash on creation"
state: merged
author: rverdile
created: 2022-10-07T20:49:42Z
updated: 2022-10-24T11:30:44Z
closed: 2022-10-24T11:04:54Z
merged: 2022-10-24T11:04:54Z
base_branch: main
head_branch: url-slash
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/121
---

# Pull Request #121: Fixes 203: Add URL trailing slash on creation

**Author**: @rverdile
**Created**: October 07, 2022 at 08:49 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `url-slash`

## Description

Adds URL trailing slash on creation, if missing, and removes extra trailing slashes (e.g. http://example/path/////)

Also:
* Trims leading/trailing whitespace from URL before creation
* Reject requests where URLs contain other whitespace

---

## Discussion

### Comment by @jlsherrill on October 07, 2022 at 09:00 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-203

### Comment by @jlsherrill on October 12, 2022 at 04:02 PM UTC

one small testing comment  otherwise this all looks good!

### Comment by @jlsherrill on October 13, 2022 at 02:06 PM UTC

we should also add a trailing slash for the validation api

### Comment by @rverdile on October 19, 2022 at 02:14 PM UTC

updated. everything should be addressed now. most of the renames were resolved just by rebasing

---

## Reviews

### Review by @jlsherrill - Commented on October 12, 2022 at 03:41 PM UTC

### Review by @rverdile - Commented on October 13, 2022 at 03:09 PM UTC

### Review by @jlsherrill - Commented on October 13, 2022 at 07:25 PM UTC

### Review by @jlsherrill - Commented on October 13, 2022 at 09:02 PM UTC

### Review by @jlsherrill - Commented on October 13, 2022 at 09:02 PM UTC

### Review by @rverdile - Commented on October 13, 2022 at 09:53 PM UTC

### Review by @rverdile - Commented on October 13, 2022 at 09:53 PM UTC

### Review by @rverdile - Commented on October 14, 2022 at 04:18 PM UTC

### Review by @jlsherrill - Commented on October 14, 2022 at 06:59 PM UTC

### Review by @jlsherrill - Commented on October 14, 2022 at 08:36 PM UTC

### Review by @jlsherrill - Commented on October 14, 2022 at 08:36 PM UTC

### Review by @Andrewgdewar - Commented on October 17, 2022 at 06:19 PM UTC

### Review by @rverdile - Commented on October 17, 2022 at 06:29 PM UTC

### Review by @Andrewgdewar - Commented on October 17, 2022 at 06:35 PM UTC

### Review by @Andrewgdewar - Approved on October 17, 2022 at 06:36 PM UTC

This looks good to me!

### Review by @Andrewgdewar - Commented on October 17, 2022 at 06:38 PM UTC

### Review by @Andrewgdewar - Commented on October 17, 2022 at 06:39 PM UTC

### Review by @swadeley - Changes Requested on October 18, 2022 at 09:54 AM UTC

Please review the use of domain names with unknown owners.

### Review by @rverdile - Commented on October 18, 2022 at 06:39 PM UTC

### Review by @rverdile - Commented on October 18, 2022 at 06:42 PM UTC

### Review by @rverdile - Commented on October 18, 2022 at 06:44 PM UTC

### Review by @jlsherrill - Commented on October 19, 2022 at 12:30 PM UTC

### Review by @swadeley - Approved on October 24, 2022 at 10:41 AM UTC

lgtm

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/121*
