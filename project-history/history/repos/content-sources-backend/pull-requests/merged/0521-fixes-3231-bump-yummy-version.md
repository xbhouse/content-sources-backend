---
type: pull_request
number: 521
title: "Fixes 3231: Bump yummy version"
state: merged
author: xbhouse
created: 2024-01-05T19:35:23Z
updated: 2024-02-13T15:16:00Z
closed: 2024-01-19T09:02:40Z
merged: 2024-01-19T09:02:40Z
base_branch: main
head_branch: update-yummy
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/521
---

# Pull Request #521: Fixes 3231: Bump yummy version

**Author**: @xbhouse
**Created**: January 05, 2024 at 07:35 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `update-yummy`

## Description

## Summary

- bumped yummy version to latest

## Testing steps

- tests pass
- repos with compressed comps (https://fixtures.pulpproject.org/rpm-pkglists-updateinfo/, https://fixtures.pulpproject.org/rpm-unsigned/) can be introspected without error

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on January 08, 2024 at 01:26 PM UTC

/retest

### Comment by @jlsherrill on January 08, 2024 at 05:23 PM UTC

https://issues.redhat.com/browse/HMS-3231

### Comment by @jlsherrill on January 08, 2024 at 05:23 PM UTC

⚠️ This task currently requires qe-approval, but this PR does not fully resolve the task.  Please contact QE to determine appropriate testing required.

### Comment by @swadeley on January 09, 2024 at 09:01 AM UTC

/retest

### Comment by @jlsherrill on January 15, 2024 at 07:16 PM UTC

/retest

### Comment by @jlsherrill on January 18, 2024 at 03:08 PM UTC

[test]

### Comment by @jlsherrill on January 18, 2024 at 03:08 PM UTC

/retest

### Comment by @swadeley on January 19, 2024 at 09:02 AM UTC

Hi

Introspection works with the two repos in comment 0

---

## Reviews

### Review by @jlsherrill - Approved on January 09, 2024 at 05:32 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/521*
