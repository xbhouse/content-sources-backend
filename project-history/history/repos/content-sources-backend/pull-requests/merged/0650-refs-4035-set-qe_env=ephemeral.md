---
type: pull_request
number: 650
title: "Refs 4035: set QE_ENV=\"ephemeral\""
state: merged
author: swadeley
created: 2024-04-26T18:07:31Z
updated: 2025-09-08T07:30:45Z
closed: 2024-04-26T18:28:29Z
merged: 2024-04-26T18:28:29Z
base_branch: main
head_branch: swadeley/try_emphemeral_venv
labels: []
url: https://github.com/content-services/content-sources-backend/pull/650
---

# Pull Request #650: Refs 4035: set QE_ENV="ephemeral"

**Author**: @swadeley
**Created**: April 26, 2024 at 06:07 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `swadeley/try_emphemeral_venv`

## Description

## Summary

Recent errors in pr_checks (which use clowder venv):
`HTTP response body: Bad Request: x-rh-identity header has an invalid or missing org_id`

possibly related to authentication changes, e.g. certs required, but clowder and ephemeral users are not configured for cert auth, they are user name and password only.

Lets try if ephemeral's use of caddy containers solves that problem.

## Testing steps

pr check runs OK

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @swadeley on April 26, 2024 at 06:13 PM UTC

https://issues.redhat.com/browse/HMS-4035

---

## Reviews

### Review by @jlsherrill - Approved on April 26, 2024 at 06:26 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/650*
