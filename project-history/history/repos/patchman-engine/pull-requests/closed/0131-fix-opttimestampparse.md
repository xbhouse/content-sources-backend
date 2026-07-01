---
type: pull_request
number: 131
title: "Fix optTimestampParse"
state: closed
author: semtexzv
created: 2020-02-17T09:21:44Z
updated: 2021-03-16T09:27:29Z
closed: 2020-02-17T14:55:40Z
base_branch: master
head_branch: fix-timestamp-parse
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/131
---

# Pull Request #131: Fix optTimestampParse

**Author**: @semtexzv
**Created**: February 17, 2020 at 09:21 AM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `fix-timestamp-parse`

## Description

Contained an error, which resulted in always returning null, and because of thad not saving culling info into the database.

---

## Discussion

### Comment by @josef-hak on February 17, 2020 at 09:27 AM UTC

@semtexzv I've fixed it in #126 

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/131*
