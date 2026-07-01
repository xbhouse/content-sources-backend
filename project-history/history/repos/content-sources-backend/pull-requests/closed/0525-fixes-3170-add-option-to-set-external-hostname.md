---
type: pull_request
number: 525
title: "Fixes 3170: add option to set external hostname"
state: closed
author: jlsherrill
created: 2024-01-09T19:47:10Z
updated: 2024-01-23T15:24:05Z
closed: 2024-01-23T15:24:05Z
base_branch: main
head_branch: 3170
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/525
---

# Pull Request #525: Fixes 3170: add option to set external hostname

**Author**: @jlsherrill
**Created**: January 09, 2024 at 07:47 PM UTC
**Status**: Closed
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `3170`

## Description

## Summary

so we can have the repo.config file return the correct hostname for gpg keys

## Testing steps

1.  create a repo with a gpg key set, and  snapshotting enabled
2.  fetch your repos list to get uuids for the repo and latest snapshots:   GET  http://localhost:8000/api/content-sources/v1/repositories/
4.   fetch this url:  http://localhost:8000/api/content-sources/v1/repositories/60dd106b-d08e-436a-bf34-53180713d4bc/snapshots/8a68fc2d-c8eb-4046-8c56-df44603f8206/config.repo

replacing the UUIDs for the repo & snapshot UUID.  Notice the URL is likely https://localhost:8000.  now re-launch the server with:

`OPTIONS_EXTERNAL_HOST="https://foobar.com"  go run cmd/content-sources/main.go api`

Refetch the config.repo file, notice the host for the gpg key has changed to foobar.com


## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on January 09, 2024 at 08:00 PM UTC

https://issues.redhat.com/browse/HMS-3170

### Comment by @jlsherrill on January 10, 2024 at 02:20 PM UTC

Note, this PR is likely superseded by https://github.com/content-services/content-sources-backend/pull/526

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/525*
