---
type: pull_request
number: 538
title: "Fixes 3423: add arm64 repos"
state: merged
author: jlsherrill
created: 2024-01-18T22:33:16Z
updated: 2024-01-19T11:00:30Z
closed: 2024-01-19T10:42:08Z
merged: 2024-01-19T10:42:08Z
base_branch: main
head_branch: 3423
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/538
---

# Pull Request #538: Fixes 3423: add arm64 repos

**Author**: @jlsherrill
**Created**: January 18, 2024 at 10:33 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `3423`

## Description

## Summary

adds the base arm64 repos 

## Testing steps

configure a cdn cert in the config if you have one, if you don't reach out to me to get one.
load in the repos:

`make  repos-import`
trigger the sync:

```go run cmd/external-repos/main.go nightly-jobs```

wait a while
make sure that all repos have snapshots.

## Checklist

- [x] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on January 18, 2024 at 11:00 PM UTC

https://issues.redhat.com/browse/HMS-3423

### Comment by @swadeley on January 19, 2024 at 10:42 AM UTC

Hi, looks good:


![image](https://github.com/content-services/content-sources-backend/assets/11247237/09aa6802-d90a-43e6-977c-9cf5d9e6fe31)


---

## Reviews

### Review by @Andrewgdewar - Approved on January 19, 2024 at 12:05 AM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/538*
