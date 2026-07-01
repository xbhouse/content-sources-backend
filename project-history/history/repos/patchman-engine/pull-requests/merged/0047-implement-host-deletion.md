---
type: pull_request
number: 47
title: "Implement host deletion"
state: merged
author: semtexzv
created: 2020-01-07T09:11:20Z
updated: 2021-03-16T09:26:46Z
closed: 2020-01-09T09:04:14Z
merged: 2020-01-09T09:04:14Z
base_branch: master
head_branch: delete-systems
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/47
---

# Pull Request #47: Implement host deletion

**Author**: @semtexzv
**Created**: January 07, 2020 at 09:11 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `delete-systems`

## Description

- Removed 'deleted_systems' table is not needed, since we are querying
inventory for systems, and can detect out-of-order deletions that way.

  This table was used as a guard against accidentally reordered upload,
and delete messages. If in a rare case, a delete message was processed
before an upload message for same host, it would mean that host would
be deleted and re-created. We can detect this when processing an upload
message. If the call to get 'system_profile' returns no data, host is deleted.

- Re-organized the handlers, to make event parsing code shared.

---

## Discussion

### Comment by @codecov-io on January 07, 2020 at 10:43 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/47?src=pr&el=h1) Report
> Merging [#47](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/47?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/37dc76a36738dce38662bfb85fc04c5a0f585b20?src=pr&el=desc) will **increase** coverage by `0.88%`.
> The diff coverage is `51.61%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/47/graphs/tree.svg?width=650&token=I7QOHXwVxB&height=150&src=pr)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/47?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master      #47      +/-   ##
==========================================
+ Coverage   56.68%   57.57%   +0.88%     
==========================================
  Files          26       27       +1     
  Lines         912      924      +12     
==========================================
+ Hits          517      532      +15     
+ Misses        337      334       -3     
  Partials       58       58
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/47?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [base/utils/gin.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/47/diff?src=pr&el=tree#diff-YmFzZS91dGlscy9naW4uZ28=) | `0% <ø> (ø)` | :arrow_up: |
| [base/utils/rpm.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/47/diff?src=pr&el=tree#diff-YmFzZS91dGlscy9ycG0uZ28=) | `85.71% <ø> (ø)` | :arrow_up: |
| [base/database/batch.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/47/diff?src=pr&el=tree#diff-YmFzZS9kYXRhYmFzZS9iYXRjaC5nbw==) | `68.51% <ø> (ø)` | :arrow_up: |
| [listener/delete.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/47/diff?src=pr&el=tree#diff-bGlzdGVuZXIvZGVsZXRlLmdv) | `44.44% <44.44%> (ø)` | |
| [listener/listener.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/47/diff?src=pr&el=tree#diff-bGlzdGVuZXIvbGlzdGVuZXIuZ28=) | `52.54% <45.45%> (+3.48%)` | :arrow_up: |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/47/diff?src=pr&el=tree#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `66.66% <63.63%> (+10.66%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/47?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/47?src=pr&el=footer). Last update [37dc76a...a6b7494](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/47?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Changes Requested on January 07, 2020 at 01:10 PM UTC

### Review by @semtexzv - Commented on January 08, 2020 at 09:28 AM UTC

### Review by @semtexzv - Commented on January 08, 2020 at 09:28 AM UTC

### Review by @semtexzv - Commented on January 08, 2020 at 09:32 AM UTC

### Review by @semtexzv - Commented on January 09, 2020 at 08:38 AM UTC

### Review by @josef-hak - Commented on January 09, 2020 at 09:02 AM UTC

### Review by @josef-hak - Approved on January 09, 2020 at 09:04 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/47*
