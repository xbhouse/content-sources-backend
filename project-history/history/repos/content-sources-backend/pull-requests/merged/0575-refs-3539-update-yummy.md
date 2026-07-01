---
type: pull_request
number: 575
title: "Refs 3539: update yummy"
state: merged
author: xbhouse
created: 2024-02-13T15:21:19Z
updated: 2024-02-14T08:25:34Z
closed: 2024-02-14T08:25:34Z
merged: 2024-02-14T08:25:34Z
base_branch: main
head_branch: update-yummy
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/575
---

# Pull Request #575: Refs 3539: update yummy

**Author**: @xbhouse
**Created**: February 13, 2024 at 03:21 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `update-yummy`

## Description

## Summary

Updates yummy to latest version. Includes fix where repo comps with certain file extensions or of group_gz data type were not detected.

## Testing steps

These repos should be able to be introspected successfully: 
- http://yum.theforeman.org/pulpcore/3.16/el7/x86_64/
- http://yum.theforeman.org/pulpcore/3.16/el8/x86_64/

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on February 13, 2024 at 03:30 PM UTC

https://issues.redhat.com/browse/HMS-3539

### Comment by @jlsherrill on February 13, 2024 at 03:30 PM UTC

⚠️ This task currently requires qe-approval, but this PR does not fully resolve the task.  Please contact QE to determine appropriate testing required.

### Comment by @swadeley on February 13, 2024 at 06:03 PM UTC

Looks good

![image](https://github.com/content-services/content-sources-backend/assets/11247237/0820d373-f0f9-4b07-b204-ffd67e59aeab)


---

## Reviews

### Review by @rverdile - Approved on February 13, 2024 at 09:17 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/575*
