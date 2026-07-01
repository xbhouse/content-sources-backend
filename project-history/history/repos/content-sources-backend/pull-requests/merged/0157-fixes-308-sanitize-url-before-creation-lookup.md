---
type: pull_request
number: 157
title: "Fixes 308: sanitize url before creation lookup"
state: merged
author: jlsherrill
created: 2022-12-13T19:03:16Z
updated: 2022-12-15T17:48:09Z
closed: 2022-12-14T22:02:05Z
merged: 2022-12-14T22:02:05Z
base_branch: main
head_branch: 308
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/157
---

# Pull Request #157: Fixes 308: sanitize url before creation lookup

**Author**: @jlsherrill
**Created**: December 13, 2022 at 07:03 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `308`

## Description

*No description provided*

---

## Discussion

### Comment by @jlsherrill on December 13, 2022 at 07:04 PM UTC


For testing, run this creation api:
```
 curl -X POST http://localhost:8000/api/content-sources/v1/repositories/  -d '{"name":"", "url":"something"}'  -H "`./scripts/header.sh 3 3`"   -H "Content-Type: application/json"
```
you should get:
```
{"errors":[{"status":400,"title":"Error creating repository","detail":"Name cannot be blank."}]}
```

and if you run it a 2nd time, you should get the same.

### Comment by @jlsherrill on December 13, 2022 at 08:20 PM UTC

good call it does!

### Comment by @jlsherrill on December 13, 2022 at 10:15 PM UTC

fixed Create & Update and added a test, both tests fail on master

### Comment by @jlsherrill on December 14, 2022 at 01:53 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-308

### Comment by @mshriver on December 14, 2022 at 10:01 PM UTC

Deployed the container and ran our test cases that cover this validation back-to-back with consistent results, thanks @jlsherrill !

---

## Reviews

### Review by @rverdile - Commented on December 13, 2022 at 07:05 PM UTC

Does Update need this too?

### Review by @rverdile - Approved on December 14, 2022 at 04:27 PM UTC

lgtm!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/157*
