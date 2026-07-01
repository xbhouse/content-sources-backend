---
type: pull_request
number: 216
title: "Fixes 1251: better error msgs"
state: merged
author: jlsherrill
created: 2023-02-16T14:16:30Z
updated: 2023-02-17T10:00:22Z
closed: 2023-02-17T09:41:20Z
merged: 2023-02-17T09:41:20Z
base_branch: main
head_branch: 1251
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/216
---

# Pull Request #216: Fixes 1251: better error msgs

**Author**: @jlsherrill
**Created**: February 16, 2023 at 02:16 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `1251`

## Description

by pulling in a newer yummy


## Testing steps
Force add a bad repository via the api
see the error message returned in the error tooltip
It should now include what file errored (404)


---

## Discussion

### Comment by @jlsherrill on February 16, 2023 at 02:30 PM UTC

https://issues.redhat.com/browse/HMS-1251

### Comment by @swadeley on February 17, 2023 at 09:41 AM UTC

Hello

I added bad repo
Tooltip now says:
Cannot fetch https://fixtures.pulpproject.org/404.html/repodata/repomd.xml: 404

Thank you

---

## Reviews

### Review by @rverdile - Approved on February 16, 2023 at 02:25 PM UTC

ack

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/216*
