---
type: pull_request
number: 737
title: "Refs 4414: set candlepin deploy params"
state: merged
author: jlsherrill
created: 2024-07-09T18:00:12Z
updated: 2024-07-09T18:41:32Z
closed: 2024-07-09T18:41:32Z
merged: 2024-07-09T18:41:32Z
base_branch: main
head_branch: 4414
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/737
---

# Pull Request #737: Refs 4414: set candlepin deploy params

**Author**: @jlsherrill
**Created**: July 09, 2024 at 06:00 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `4414`

## Description

## Summary

Set variables needed for candlepin communication.  Cert & key come from a secret.  This won't affect ephemeral at all

## Testing steps
tests are green

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on July 09, 2024 at 06:30 PM UTC

https://issues.redhat.com/browse/HMS-4414

### Comment by @jlsherrill on July 09, 2024 at 06:30 PM UTC

⚠️ This task currently requires qe-approval, but this PR does not fully resolve the task.  Please contact QE to determine appropriate testing required.

---

## Reviews

### Review by @xbhouse - Approved on July 09, 2024 at 06:20 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/737*
