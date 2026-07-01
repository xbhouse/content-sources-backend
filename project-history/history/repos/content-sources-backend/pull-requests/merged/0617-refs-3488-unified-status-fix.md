---
type: pull_request
number: 617
title: "Refs 3488: unified status fix"
state: merged
author: xbhouse
created: 2024-03-27T17:53:39Z
updated: 2024-03-27T18:18:52Z
closed: 2024-03-27T18:18:52Z
merged: 2024-03-27T18:18:51Z
base_branch: main
head_branch: status-fix
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/617
---

# Pull Request #617: Refs 3488: unified status fix

**Author**: @xbhouse
**Created**: March 27, 2024 at 05:53 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `status-fix`

## Description

## Summary

fixes issues in stage from [this PR](https://github.com/content-services/content-sources-backend/pull/587) where repos with snapshotting disabled were not able to be listed or filtered 

## Testing steps

- add repo with snapshotting disabled, then edit it to turn it on
- listing repositories should return all repos
- listing repositories with a filter on the status should return all repos with that status

## Checklist

- [x] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on March 27, 2024 at 06:00 PM UTC

https://issues.redhat.com/browse/HMS-3488

---

## Reviews

### Review by @rverdile - Approved on March 27, 2024 at 06:11 PM UTC

tested and lgtm!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/617*
