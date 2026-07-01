---
type: pull_request
number: 15
title: "Performance measurement skeleton"
state: merged
author: semtexzv
created: 2019-11-18T14:51:34Z
updated: 2019-11-20T09:14:54Z
closed: 2019-11-20T09:14:54Z
merged: 2019-11-20T09:14:54Z
base_branch: master
head_branch: master
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/15
---

# Pull Request #15: Performance measurement skeleton

**Author**: @semtexzv
**Created**: November 18, 2019 at 02:51 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `master`

## Description

*No description provided*

---

## Reviews

### Review by @josef-hak - Commented on November 18, 2019 at 03:15 PM UTC

I found one test in go broken, probably from previous PR:
- list_test.go:42
~~~
Expected :{"id":1,"request":"r","checksum":"454349e422f05297191ead13e21d3db520e5abef52055e4964b82fb213f593a1"}
Actual   :{"id":1,"request":"r","checksum":"454349e422f05297191ead13e21d3db520e5abef52055e4964b82fb213f593a1","updated":"0001-01-01T00:00:00Z"}
~~~

### Review by @josef-hak - Changes Requested on November 18, 2019 at 03:23 PM UTC

### Review by @semtexzv - Commented on November 20, 2019 at 09:14 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/15*
