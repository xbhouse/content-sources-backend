---
type: pull_request
number: 1228
title: "HMS-9203: Add detected OS version to snapshots"
state: merged
author: rverdile
created: 2025-10-07T17:23:40Z
updated: 2025-10-22T13:35:06Z
closed: 2025-10-22T13:35:03Z
merged: 2025-10-22T13:35:03Z
base_branch: main
head_branch: detected-release
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/1228
---

# Pull Request #1228: HMS-9203: Add detected OS version to snapshots

**Author**: @rverdile
**Created**: October 07, 2025 at 05:23 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `detected-release`

## Description

## Summary
For BaseOS Red Hat repositories, we can detect and return an expected OS version for the snapshot by checking the version of the redhat-release package

## Testing steps
1. Without this PR, import and snapshot RHEL repos
2. Switch to this PR
3. If you list the RHEL repos, you should see a new field under "last_snapshot" called "detected_os_version"
4. This field should be blank for all repos, including BaseOS
5. Run the job this PR adds: `go run cmd/jobs/main.go set-detected-os-versions`
6. Check the snapshots for the BaseOS repos, they should now have that field set
7. `make compose-clean compose-up`
8. Import and snapshot RHEL repos
9. The BaseOS repo snapshots should have a "detected_os_version"



---

## Discussion

### Comment by @swadeley on October 08, 2025 at 10:35 AM UTC

Hi, 'detected_os_version' is not supported by the old IQE tests. The time has come to shut them down.

### Comment by @TenSt on October 08, 2025 at 12:26 PM UTC

Just my 2 cents as a new guy:
- it would be great to add unit tests as there are some new functions and some old ones gets changed
- there is no check if a `detected_os_version` field is already populated - I assume this is a recurring job that runs on some schedule, right? If no, ignore this. If yes, adding a check at the start might save time and performance

### Comment by @rverdile on October 08, 2025 at 02:26 PM UTC

> Just my 2 cents as a new guy:
> 
> * it would be great to add unit tests as there are some new functions and some old ones gets changed
> * there is no check if a `detected_os_version` field is already populated - I assume this is a recurring job that runs on some schedule, right? If no, ignore this. If yes, adding a check at the start might save time and performance

- I agree. I forgot to add a unit test :)
- The job just runs once to update the existing snapshots

### Comment by @jlsherrill on October 08, 2025 at 03:32 PM UTC

Code and testing looks good, one small request about the test

### Comment by @jlsherrill on October 08, 2025 at 05:36 PM UTC

/retest

### Comment by @rverdile on October 08, 2025 at 07:30 PM UTC

@swadeley would the IQE bindings need to be regenerated in the meantime?

### Comment by @xbhouse on October 09, 2025 at 07:30 PM UTC

https://issues.redhat.com/browse/HMS-9203

### Comment by @swadeley on October 13, 2025 at 08:52 AM UTC

/retest

### Comment by @rverdile on October 13, 2025 at 01:18 PM UTC

/retest

### Comment by @rverdile on October 14, 2025 at 01:06 PM UTC

/retest

### Comment by @swadeley on October 14, 2025 at 05:13 PM UTC

> IQE

I don't see any failure, so , seems not. But lets check nightly tests.
EDIT: Konflux `run-iqe-cji` did fail

### Comment by @rverdile on October 21, 2025 at 05:32 PM UTC

/retest

### Comment by @rverdile on October 21, 2025 at 07:09 PM UTC

/retest

### Comment by @rverdile on October 21, 2025 at 07:32 PM UTC

/retest

### Comment by @rverdile on October 21, 2025 at 08:15 PM UTC

skipped snyk check because failure is unrelated

---

## Reviews

### Review by @jlsherrill - Commented on October 08, 2025 at 01:50 PM UTC

### Review by @rverdile - Commented on October 08, 2025 at 02:31 PM UTC

### Review by @jlsherrill - Commented on October 08, 2025 at 03:32 PM UTC

### Review by @rverdile - Commented on October 08, 2025 at 05:32 PM UTC

### Review by @jlsherrill - Approved on October 08, 2025 at 05:36 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1228*
