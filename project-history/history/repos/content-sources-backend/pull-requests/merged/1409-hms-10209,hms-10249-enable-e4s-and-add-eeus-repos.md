---
type: pull_request
number: 1409
title: "HMS-10209,HMS-10249:  enable e4s and add eeus repos"
state: merged
author: dominikvagner
created: 2026-03-11T13:17:41Z
updated: 2026-03-23T13:59:31Z
closed: 2026-03-23T13:59:31Z
merged: 2026-03-23T13:59:31Z
base_branch: main
head_branch: add-e4s-and-eeus-repos
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1409
---

# Pull Request #1409: HMS-10209,HMS-10249:  enable e4s and add eeus repos

**Author**: @dominikvagner
**Created**: March 11, 2026 at 01:17 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `add-e4s-and-eeus-repos`

## Description

## Summary
This fixes the config for E4S repos and removes their temporary
disablement and adds configs for EEUS (aarch64 & x86_64).

Also adds EUS aarch64 repos.

## Testing Steps
- Verify the repos are returned in the list repos request, if your user has the permissions for them.
- Verify the repos can be snapshotted correctly. _(the `AppStream` issued should be resolved now)_

---

## Discussion

### Comment by @xbhouse on March 11, 2026 at 01:30 PM UTC

https://issues.redhat.com/browse/HMS-10209

### Comment by @xbhouse on March 11, 2026 at 01:30 PM UTC

https://issues.redhat.com/browse/HMS-10249

---

## Reviews

### Review by @TenSt - Approved on March 20, 2026 at 01:34 PM UTC

lgtm

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1409*
