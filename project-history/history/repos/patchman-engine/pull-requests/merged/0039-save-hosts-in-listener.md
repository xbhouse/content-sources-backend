---
type: pull_request
number: 39
title: "Save hosts in listener"
state: merged
author: semtexzv
created: 2019-12-17T16:34:41Z
updated: 2021-03-16T09:26:45Z
closed: 2020-01-03T14:33:03Z
merged: 2020-01-03T14:33:03Z
base_branch: master
head_branch: listener-db-save
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/39
---

# Pull Request #39: Save hosts in listener

**Author**: @semtexzv
**Created**: December 17, 2019 at 04:34 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `listener-db-save`

## Description

Moves code which processes uploads to separate file.  Implements process to get system_profile from inventory, create VMaaS request, And store both RHAccount and SystemProfile into the database.

Tests for parts which do not require external APIs are implemented.

---

## Reviews

### Review by @josef-hak - Changes Requested on January 03, 2020 at 11:10 AM UTC

Thanks for contribution. I've added some proposals to updates.

### Review by @semtexzv - Commented on January 03, 2020 at 12:03 PM UTC

### Review by @semtexzv - Commented on January 03, 2020 at 12:44 PM UTC

### Review by @semtexzv - Commented on January 03, 2020 at 12:45 PM UTC

### Review by @semtexzv - Commented on January 03, 2020 at 12:45 PM UTC

### Review by @semtexzv - Commented on January 03, 2020 at 12:48 PM UTC

### Review by @josef-hak - Approved on January 03, 2020 at 02:32 PM UTC

ok, thanks

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/39*
