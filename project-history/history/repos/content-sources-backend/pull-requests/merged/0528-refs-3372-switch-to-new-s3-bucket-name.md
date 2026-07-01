---
type: pull_request
number: 528
title: "Refs 3372: switch to new s3 bucket name"
state: merged
author: jlsherrill
created: 2024-01-11T15:49:01Z
updated: 2024-01-15T15:13:30Z
closed: 2024-01-15T15:13:19Z
merged: 2024-01-15T15:13:19Z
base_branch: main
head_branch: 3372_s3_bucket
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/528
---

# Pull Request #528: Refs 3372: switch to new s3 bucket name

**Author**: @jlsherrill
**Created**: January 11, 2024 at 03:49 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `3372_s3_bucket`

## Description

## Summary

As part of switching to new central pulp server we need to switch to a new s3 bucket for domain creation.  All previous content will be deleted

## Testing steps

none

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on January 11, 2024 at 04:00 PM UTC

https://issues.redhat.com/browse/HMS-3372

### Comment by @jlsherrill on January 11, 2024 at 04:00 PM UTC

⚠️ This task currently requires qe-approval, but this PR does not fully resolve the task.  Please contact QE to determine appropriate testing required.

### Comment by @jlsherrill on January 12, 2024 at 09:14 PM UTC

pulp is failing with 

 MountVolume.SetUp failed for volume "pulp-settings" : secret "pulp-settings" not found
and secret "pulp-db-fields-encryption" not found

asking the pulp team

### Comment by @jlsherrill on January 14, 2024 at 11:14 PM UTC

[test]

### Comment by @jlsherrill on January 15, 2024 at 01:18 PM UTC

[test]

### Comment by @jlsherrill on January 15, 2024 at 01:36 PM UTC

[test]

### Comment by @jlsherrill on January 15, 2024 at 02:01 PM UTC

[test]

### Comment by @jlsherrill on January 15, 2024 at 02:36 PM UTC

[test]

### Comment by @jlsherrill on January 15, 2024 at 03:12 PM UTC

merging since at least the deployment works.

---

## Reviews

### Review by @swadeley - Approved on January 15, 2024 at 01:20 PM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/528*
