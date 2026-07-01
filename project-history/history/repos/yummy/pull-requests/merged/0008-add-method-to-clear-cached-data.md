---
type: pull_request
number: 8
title: "Add method to clear cached data"
state: merged
author: rverdile
created: 2022-12-06T19:14:20Z
updated: 2022-12-08T14:07:30Z
closed: 2022-12-08T14:07:25Z
merged: 2022-12-08T14:07:24Z
base_branch: master
head_branch: clear-method
labels: []
url: https://github.com/content-services/yummy/pull/8
---

# Pull Request #8: Add method to clear cached data

**Author**: @rverdile
**Created**: December 06, 2022 at 07:14 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `clear-method`

## Description

Adds `repo.Clear()`, which will set `repo.repomd`, `repo.packages`, and `repo.repomdSignature` to `nil`, so that the fetching methods will fetch new data.

---

## Discussion

### Comment by @jlsherrill on December 06, 2022 at 07:38 PM UTC

while you're at it, can you get rid of this line:  https://github.com/content-services/yummy/blob/master/pkg/yum/repository.go#L290

we may want proper logging support at some point, but right now i think its fine just to ignore this

### Comment by @jlsherrill on December 07, 2022 at 02:26 PM UTC

have a method to clear is fine, any reason not to when .Configure() is called?

---

## Reviews

### Review by @jlsherrill - Approved on December 07, 2022 at 09:18 PM UTC

---

*Archived from: https://github.com/content-services/yummy/pull/8*
