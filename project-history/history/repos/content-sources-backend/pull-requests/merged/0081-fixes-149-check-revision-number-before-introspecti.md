---
type: pull_request
number: 81
title: "Fixes 149: Check revision number before introspection"
state: merged
author: rverdile
created: 2022-08-16T21:07:50Z
updated: 2022-08-24T14:07:10Z
closed: 2022-08-24T14:07:10Z
merged: 2022-08-24T14:07:10Z
base_branch: main
head_branch: revision
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/81
---

# Pull Request #81: Fixes 149: Check revision number before introspection

**Author**: @rverdile
**Created**: August 16, 2022 at 09:07 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `revision`

## Description

Adds
- A check before introspection to see if the repository revision number has changed. If not, don't introspect because it is assumed the repository hasn't changed since the last introspection. 
  - If it has not changed, it won't error, just "successfully update 0 packages". 
- A "revision" column to the "repositories" table that saves the previous revision number.

I tested this by creating a yum repo locally and then serving it with apache. Then I could change the revision number to whatever I wanted.

Depends on this yummy PR: https://github.com/content-services/yummy/pull/6

~~To use this yummy PR, check it out locally and then add `replace github.com/content-services/yummy => path/to/local/yummy`, to `go.mod`, above `require`~~

Still needs:

- [ ] better tests I think

---

## Discussion

### Comment by @jlsherrill on August 16, 2022 at 09:30 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-149

### Comment by @jlsherrill on August 22, 2022 at 08:33 PM UTC

Code looks good, worked great, will just wait on better tests

---

## Reviews

### Review by @rverdile - Commented on August 16, 2022 at 09:09 PM UTC

### Review by @jlsherrill - Commented on August 16, 2022 at 09:20 PM UTC

### Review by @rverdile - Commented on August 17, 2022 at 02:48 PM UTC

### Review by @jlsherrill - Commented on August 22, 2022 at 07:23 PM UTC

### Review by @rverdile - Commented on August 23, 2022 at 06:01 PM UTC

Added to the current introspect to check that introspecting with the given revision number results in 0 package updates.

I don't know how comprehensive this test is, but it feels a little tricky to test more since the introspect tests mock the dao layer.

### Review by @jlsherrill - Approved on August 23, 2022 at 08:52 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/81*
