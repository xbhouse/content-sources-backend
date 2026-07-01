---
type: pull_request
number: 110
title: "Fixes 185: Save revision properly"
state: merged
author: rverdile
created: 2022-09-20T18:14:41Z
updated: 2022-09-22T16:59:05Z
closed: 2022-09-22T16:59:00Z
merged: 2022-09-22T16:59:00Z
base_branch: main
head_branch: update-bug
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/110
---

# Pull Request #110: Fixes 185: Save revision properly

**Author**: @rverdile
**Created**: September 20, 2022 at 06:14 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `update-bug`

## Description

The `Update` function in the repository dao was not working correctly when passed zero values. It would update the model with the zero values.

It may be best to have the [internal Repository object](https://github.com/content-services/content-sources-backend/blob/main/pkg/dao/repositories.go#L10) split into two the way that the RepoConfig layer has RepositoryRequest and RepositoryResponse. This way we can distinguish between what should be updated and what shouldn't be with `nil`. I can make a change like that here if we want.

For now I fixed this in the least disruptive way possible because it's blocking #92 

To test:
Introspect and repository and the Revision field should be saved correctly. Without this change, it would be blank.

---

## Discussion

### Comment by @jlsherrill on September 20, 2022 at 06:30 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-185

### Comment by @jlsherrill on September 21, 2022 at 06:25 PM UTC

@rverdile mind filing a ticket to fix this in a more destructive way? 

### Comment by @jlsherrill on September 21, 2022 at 08:42 PM UTC

needs a rebase now, but looks good and works well

---

## Reviews

### Review by @jlsherrill - Approved on September 22, 2022 at 01:01 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/110*
