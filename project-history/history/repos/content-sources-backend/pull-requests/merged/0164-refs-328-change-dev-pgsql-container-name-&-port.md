---
type: pull_request
number: 164
title: "Refs 328: change dev pgsql container name & port"
state: merged
author: jlsherrill
created: 2023-01-05T18:42:40Z
updated: 2023-01-23T21:27:20Z
closed: 2023-01-13T19:59:54Z
merged: 2023-01-13T19:59:54Z
base_branch: main
head_branch: 328
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/164
---

# Pull Request #164: Refs 328: change dev pgsql container name & port

**Author**: @jlsherrill
**Created**: January 05, 2023 at 06:42 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `328`

## Description

this will allow us to use the existing docker compose for spinning up a pulp server

---

## Discussion

### Comment by @jlsherrill on January 05, 2023 at 06:43 PM UTC

NOTE, this requires a change in all our configs: 
From:
```
database:
  host: localhost
  port: 5432
```
to
```
database:
  host: localhost
  port: 5433
```

### Comment by @jlsherrill on January 05, 2023 at 07:00 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-328

### Comment by @jlsherrill on January 23, 2023 at 09:27 PM UTC

https://issues.redhat.com/browse/HMS-540

---

## Reviews

### Review by @rverdile - Commented on January 06, 2023 at 03:30 PM UTC

### Review by @rverdile - Commented on January 06, 2023 at 03:30 PM UTC

### Review by @jlsherrill - Commented on January 06, 2023 at 10:23 PM UTC

### Review by @rverdile - Commented on January 09, 2023 at 03:22 PM UTC

### Review by @rverdile - Approved on January 09, 2023 at 05:57 PM UTC

lgtm

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/164*
