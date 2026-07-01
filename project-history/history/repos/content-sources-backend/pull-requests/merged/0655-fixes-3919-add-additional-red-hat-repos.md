---
type: pull_request
number: 655
title: "Fixes 3919: Add additional red hat repos"
state: merged
author: jlsherrill
created: 2024-04-29T21:25:30Z
updated: 2024-05-02T12:30:19Z
closed: 2024-05-02T12:00:57Z
merged: 2024-05-02T12:00:57Z
base_branch: main
head_branch: 3919
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/655
---

# Pull Request #655: Fixes 3919: Add additional red hat repos

**Author**: @jlsherrill
**Created**: April 29, 2024 at 09:25 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `3919`

## Description

## Summary

Adds some more rhel 8 and rhel 9 repository for snapahostting

## Testing steps

make repos-import
go run cmd/external-repos/main.go nightly-jobs

wait a while for it to sync all the repos.  Navigate in the UI to red hat repos and make sure they were snapshotted.

For QE:  i'd just check a deployment and see that the new rhel repos are there (if anything)

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on April 29, 2024 at 09:30 PM UTC

https://issues.redhat.com/browse/HMS-3919

### Comment by @jlsherrill on May 01, 2024 at 12:52 AM UTC

rhel9 supplementary doesn't seem to have much, let me take a closer look

### Comment by @swadeley on May 01, 2024 at 07:15 AM UTC

Hi, I see this in template modal:

![image](https://github.com/content-services/content-sources-backend/assets/11247237/1723b09d-6eec-492d-934d-97b4c1305767)


### Comment by @jlsherrill on May 01, 2024 at 05:26 PM UTC

adjusted to code ready repos

### Comment by @jlsherrill on May 01, 2024 at 05:26 PM UTC

/retest

### Comment by @swadeley on May 02, 2024 at 12:00 PM UTC

Hi

now with CodeReady repos

![image](https://github.com/content-services/content-sources-backend/assets/11247237/5d5c0775-50e5-43af-8a39-18ba6511e970)


---

## Reviews

### Review by @rverdile - Approved on May 01, 2024 at 05:59 PM UTC

worked

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/655*
