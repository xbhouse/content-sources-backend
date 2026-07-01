---
type: pull_request
number: 757
title: "Fixes 4519: remove epel7 repo"
state: merged
author: swadeley
created: 2024-08-01T14:47:16Z
updated: 2025-09-08T07:30:52Z
closed: 2024-08-01T17:44:05Z
merged: 2024-08-01T17:44:05Z
base_branch: main
head_branch: swadeley/remove_epel7_repo
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/757
---

# Pull Request #757: Fixes 4519: remove epel7 repo

**Author**: @swadeley
**Created**: August 01, 2024 at 02:47 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `swadeley/remove_epel7_repo`

## Description

## Summary


Hi

Tests are failing because Epel 7 repo is EOL and has been disabled.

http://dl.fedoraproject.org/pub/epel/7/x86_64/

## Testing steps

pr_check test run passes

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on August 01, 2024 at 03:00 PM UTC

https://issues.redhat.com/browse/HMS-4519

### Comment by @swadeley on August 01, 2024 at 05:44 PM UTC

Thank you

---

## Reviews

### Review by @jlsherrill - Approved on August 01, 2024 at 04:55 PM UTC

ACK, tested locally

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/757*
