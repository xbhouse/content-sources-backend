---
type: pull_request
number: 514
title: "Fix last migration"
state: merged
author: semtexzv
created: 2021-02-02T10:48:31Z
updated: 2021-03-16T09:30:29Z
closed: 2021-02-03T11:14:44Z
merged: 2021-02-03T11:14:44Z
base_branch: master
head_branch: fix-migration
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/514
---

# Pull Request #514: Fix last migration

**Author**: @semtexzv
**Created**: February 02, 2021 at 10:48 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `fix-migration`

## Description

*No description provided*

---

## Discussion

### Comment by @semtexzv on February 02, 2021 at 11:04 AM UTC

No, This was just a precaution, but after careful consideration, that would only lock up the calls to /packages, and in turn calls to other APIs


---

## Reviews

### Review by @josef-hak - Approved on February 02, 2021 at 10:56 AM UTC

As we are not locking the table we should stop evaluation not to write into the table during the migration, right?

### Review by @MichaelMraka - Approved on February 02, 2021 at 11:03 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/514*
