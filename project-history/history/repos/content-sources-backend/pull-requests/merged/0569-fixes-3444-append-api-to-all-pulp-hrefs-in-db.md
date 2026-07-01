---
type: pull_request
number: 569
title: "Fixes 3444: append /api/ to all pulp hrefs in db"
state: merged
author: xbhouse
created: 2024-02-09T17:18:49Z
updated: 2024-02-19T14:51:39Z
closed: 2024-02-19T14:51:38Z
merged: 2024-02-19T14:51:38Z
base_branch: main
head_branch: update-hrefs
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/569
---

# Pull Request #569: Fixes 3444: append /api/ to all pulp hrefs in db

**Author**: @xbhouse
**Created**: February 09, 2024 at 05:18 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `update-hrefs`

## Description

## Summary

Job runs on stage that updates pulp hrefs (version_href, publication_href, distribution_href) to prepend /api if not already. 

## Testing steps

- Run query on local snapshots table in db. Shouldn't update any rows if hrefs are already prepended with /api
- *Maybe* run query on stage db?  I haven't done this yet myself though
- Don't merge on a Friday

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on February 09, 2024 at 05:30 PM UTC

https://issues.redhat.com/browse/HMS-3444

### Comment by @xbhouse on February 09, 2024 at 05:45 PM UTC

i'm not sure what caused this test failure. and when running the tests locally they just hang, never gets to the rbac tests :/ 

EDIT: rbac test was fine locally, an rpm test was timing out after 10 min but passed after clearing out the rpms table
EDIT: reran failed jobs and the rbac unit test resolved itself

### Comment by @swadeley on February 09, 2024 at 08:49 PM UTC

/retest

### Comment by @swadeley on February 12, 2024 at 07:04 AM UTC

/retest

### Comment by @swadeley on February 12, 2024 at 04:40 PM UTC

/retest

---

## Reviews

### Review by @jlsherrill - Commented on February 14, 2024 at 09:57 PM UTC

### Review by @jlsherrill - Approved on February 14, 2024 at 09:57 PM UTC

### Review by @xbhouse - Commented on February 15, 2024 at 07:25 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/569*
