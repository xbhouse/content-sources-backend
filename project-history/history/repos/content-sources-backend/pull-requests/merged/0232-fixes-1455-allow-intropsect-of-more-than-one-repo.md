---
type: pull_request
number: 232
title: "Fixes 1455: allow intropsect of more than one repo"
state: merged
author: jlsherrill
created: 2023-03-17T16:18:42Z
updated: 2023-03-17T18:44:05Z
closed: 2023-03-17T18:44:05Z
merged: 2023-03-17T18:44:05Z
base_branch: main
head_branch: 1466
labels: []
url: https://github.com/content-services/content-sources-backend/pull/232
---

# Pull Request #232: Fixes 1455: allow intropsect of more than one repo

**Author**: @jlsherrill
**Created**: March 17, 2023 at 04:18 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `1466`

## Description

and add rhel 8.7 baseos to init containers

## Testing steps
* tests pass
* ` go run cmd/external-repos/main.go import`
*  `go run cmd/external-repos/main.go introspect  https://cdn.redhat.com/content/dist/rhel8/8.7/x86_64/baseos/os https://cdn.redhat.com/content/dist/rhel8/8.6/x86_64/baseos/os`
*  `go run cmd/external-repos/main.go introspect  https://cdn.redhat.com/content/dist/rhel8/8.7/x86_64/baseos/os https://cdn.redhat.com/content/dist/rhel8/8.6/x86_64/baseos/os --force`

---

## Discussion

### Comment by @jlsherrill on March 17, 2023 at 04:30 PM UTC

https://issues.redhat.com/browse/HMS-1455

### Comment by @Andrewgdewar on March 17, 2023 at 06:02 PM UTC

I tested this with some mixed results, will confer with you before acking.

---

## Reviews

### Review by @jlsherrill - Commented on March 17, 2023 at 04:19 PM UTC

### Review by @Andrewgdewar - Approved on March 17, 2023 at 06:24 PM UTC

Tested with normal repos, working well

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/232*
