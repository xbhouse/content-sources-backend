---
type: pull_request
number: 20
title: "Add RHEL dockerfiles"
state: merged
author: semtexzv
created: 2019-11-27T09:41:31Z
updated: 2021-03-16T09:28:44Z
closed: 2019-11-29T10:07:21Z
merged: 2019-11-29T10:07:21Z
base_branch: master
head_branch: rhel-dockerfiles
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/20
---

# Pull Request #20: Add RHEL dockerfiles

**Author**: @semtexzv
**Created**: November 27, 2019 at 09:41 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `rhel-dockerfiles`

## Description

- Added Dockerfiles for running the app on RHEL 7 & 8.
- Added script to check Dockerfiles consistency.
- We're waiting for CentOS images to catch up in order to migrate database to CentOS & Rhel 8

---

## Discussion

### Comment by @josef-hak on November 29, 2019 at 09:20 AM UTC

Please remove ".hel8/7" suffixes in production dockerfiles as discussed. 

### Comment by @semtexzv on November 29, 2019 at 09:54 AM UTC

I'd rather keep it as is. If someone will try to run the project, and they will try to use rhel based dockerfiles, they will encounter errors because of missing subscriptions. This variant contains most  information

---

## Reviews

### Review by @josef-hak - Changes Requested on November 27, 2019 at 01:55 PM UTC

### Review by @josef-hak - Commented on November 27, 2019 at 03:00 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/20*
