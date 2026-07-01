---
type: pull_request
number: 613
title: "Fixes 3813: correct repo label for 9 appstream"
state: merged
author: jlsherrill
created: 2024-03-21T11:53:25Z
updated: 2024-03-21T16:00:23Z
closed: 2024-03-21T15:34:35Z
merged: 2024-03-21T15:34:35Z
base_branch: main
head_branch: 3813
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/613
---

# Pull Request #613: Fixes 3813: correct repo label for 9 appstream

**Author**: @jlsherrill
**Created**: March 21, 2024 at 11:53 AM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `3813`

## Description

## Summary
fixes a repo label
## Testing steps

on main:  ```make repos-import```
```
make db-cli-connect
 select label from repository_configurations where name = 'Red Hat Enterprise Linux 9 for x86_64 - AppStream (RPMs)';
```

switch to the branch
```make repos-import```

```
make db-cli-connect
 select label from repository_configurations where name = 'Red Hat Enterprise Linux 9 for x86_64 - AppStream (RPMs)';
```
You should see the updated label

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on March 21, 2024 at 12:00 PM UTC

https://issues.redhat.com/browse/HMS-3813

---

## Reviews

### Review by @rverdile - Approved on March 21, 2024 at 02:20 PM UTC

works!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/613*
