---
type: pull_request
number: 395
title: "Fixed invalid offset handling"
state: merged
author: josef-hak
created: 2020-10-19T16:15:07Z
updated: 2020-10-29T11:40:14Z
closed: 2020-10-20T07:58:55Z
merged: 2020-10-20T07:58:55Z
base_branch: master
head_branch: fix-offset-resp
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/395
---

# Pull Request #395: Fixed invalid offset handling

**Author**: @josef-hak
**Created**: October 19, 2020 at 04:15 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `fix-offset-resp`

## Description

- don't return 500 (server error) but 400 (bad request) when too big `offset` param set for manager's listing endpoints

---

## Reviews

### Review by @semtexzv - Approved on October 19, 2020 at 07:02 PM UTC

Thank you for fixing this. I thing for now this'll work, but long term, we need to do something about the convention to do error handling at lower layers,  because we seem to overwrite them in lot of places, causing bugs like these.

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/395*
