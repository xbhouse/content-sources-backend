---
type: pull_request
number: 519
title: "Fixes 3171: add gpg key path for red hat repos"
state: merged
author: rverdile
created: 2024-01-04T19:15:48Z
updated: 2024-01-10T18:30:34Z
closed: 2024-01-10T18:06:00Z
merged: 2024-01-10T18:06:00Z
base_branch: main
head_branch: rh-gpg
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/519
---

# Pull Request #519: Fixes 3171: add gpg key path for red hat repos

**Author**: @rverdile
**Created**: January 04, 2024 at 07:15 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `rh-gpg`

## Description

## Summary
The repository config file wasn't providing the correct gpg key path for red hat repos. All the red hat repos have the same gpg key path, so this hard codes that value. 

## Testing steps
1. Import red hat repos and create a snapshot
2. Fetch the config.repo file from the `/repositories/:uuid/snapshots/:snap_uuid/config.repo/` endpoint
3. The config.repo file should have the red hat gpg key path

## Checklist

- [x] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on January 08, 2024 at 05:23 PM UTC

https://issues.redhat.com/browse/HMS-3171

### Comment by @swadeley on January 09, 2024 at 08:19 PM UTC

/retest

---

## Reviews

### Review by @jlsherrill - Commented on January 05, 2024 at 02:45 PM UTC

### Review by @jlsherrill - Commented on January 05, 2024 at 02:46 PM UTC

### Review by @rverdile - Commented on January 05, 2024 at 05:00 PM UTC

### Review by @rverdile - Commented on January 08, 2024 at 06:08 PM UTC

### Review by @jlsherrill - Approved on January 08, 2024 at 06:48 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/519*
