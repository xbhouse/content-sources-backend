---
type: pull_request
number: 373
title: "Move syncing advisories into separate goroutine"
state: closed
author: semtexzv
created: 2020-10-12T08:16:56Z
updated: 2021-03-16T09:29:37Z
closed: 2020-10-12T15:21:51Z
base_branch: master
head_branch: vmaas-sync-async
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/373
---

# Pull Request #373: Move syncing advisories into separate goroutine

**Author**: @semtexzv
**Created**: October 12, 2020 at 08:16 AM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `vmaas-sync-async`

## Description

Should resolve vmaas_sync crashing issue

---

## Discussion

### Comment by @semtexzv on October 12, 2020 at 12:22 PM UTC

It seems that connection to vmaas times out in prod

### Comment by @josef-hak on October 12, 2020 at 12:25 PM UTC

So it does not solve the whole issue - advisories are not synced. It mill maybe just resolve issue of restarting, right?

### Comment by @josef-hak on October 12, 2020 at 03:04 PM UTC

@semtexzv ok, makes sense, so just ensure passing pipeline and then feel free to merge it.

### Comment by @semtexzv on October 12, 2020 at 03:22 PM UTC

Superseded by https://github.com/RedHatInsights/patchman-engine/pull/376

---

## Reviews

### Review by @josef-hak - Changes Requested on October 12, 2020 at 08:26 AM UTC

1) Why should this help?
2) Tests pipeline failed.

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/373*
