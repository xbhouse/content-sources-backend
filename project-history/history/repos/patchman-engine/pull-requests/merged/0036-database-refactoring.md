---
type: pull_request
number: 36
title: "Database refactoring"
state: merged
author: semtexzv
created: 2019-12-12T13:12:32Z
updated: 2021-03-16T09:26:43Z
closed: 2019-12-17T10:09:35Z
merged: 2019-12-17T10:09:35Z
base_branch: master
head_branch: db-refactor
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/36
---

# Pull Request #36: Database refactoring

**Author**: @semtexzv
**Created**: December 12, 2019 at 01:12 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `db-refactor`

## Description

- Rename all occurences of `errata` to `advisory`, to correspond with frontend and API nomenclature
- Change `advisory_type` names to correspond to the ones we receive from VMaaS
- Remove `s3_url` from hosts, since we don't download archives.

---

## Reviews

### Review by @beav - Commented on December 12, 2019 at 09:56 PM UTC

### Review by @semtexzv - Commented on December 13, 2019 at 08:55 AM UTC

### Review by @beav - Commented on December 13, 2019 at 01:36 PM UTC

### Review by @josef-hak - Changes Requested on December 17, 2019 at 09:20 AM UTC

There is failing test because test_data.sql script is not updated according to renaming.

### Review by @josef-hak - Approved on December 17, 2019 at 10:09 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/36*
