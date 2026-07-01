---
type: pull_request
number: 74
title: "Move evaluation into separate component"
state: merged
author: semtexzv
created: 2020-01-27T09:58:47Z
updated: 2020-01-30T09:56:15Z
closed: 2020-01-30T09:56:14Z
merged: 2020-01-30T09:56:14Z
base_branch: master
head_branch: evaluator
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/74
---

# Pull Request #74: Move evaluation into separate component

**Author**: @semtexzv
**Created**: January 27, 2020 at 09:58 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `evaluator`

## Description

*No description provided*

---

## Discussion

### Comment by @codecov-io on January 27, 2020 at 10:11 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/74?src=pr&el=h1) Report
> :exclamation: No coverage uploaded for pull request base (`master@d07f394`). [Click here to learn what that means](https://docs.codecov.io/docs/error-reference#section-missing-base-commit).
> The diff coverage is `51.42%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/74/graphs/tree.svg?width=650&token=I7QOHXwVxB&height=150&src=pr)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/74?src=pr&el=tree)

```diff
@@            Coverage Diff            @@
##             master      #74   +/-   ##
=========================================
  Coverage          ?   57.02%           
=========================================
  Files             ?       34           
  Lines             ?     1359           
  Branches          ?        0           
=========================================
  Hits              ?      775           
  Misses            ?      492           
  Partials          ?       92
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/74?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [base/mqueue/mqueue.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/74/diff?src=pr&el=tree#diff-YmFzZS9tcXVldWUvbXF1ZXVlLmdv) | `65.78% <ø> (ø)` | |
| [base/mqueue/event.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/74/diff?src=pr&el=tree#diff-YmFzZS9tcXVldWUvZXZlbnQuZ28=) | `33.33% <0%> (ø)` | |
| [base/database/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/74/diff?src=pr&el=tree#diff-YmFzZS9kYXRhYmFzZS90ZXN0aW5nLmdv) | `0% <0%> (ø)` | |
| [listener/listener.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/74/diff?src=pr&el=tree#diff-bGlzdGVuZXIvbGlzdGVuZXIuZ28=) | `50% <100%> (ø)` | |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/74/diff?src=pr&el=tree#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `72.18% <47.36%> (ø)` | |
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/74/diff?src=pr&el=tree#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `68.68% <87.5%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/74?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/74?src=pr&el=footer). Last update [d07f394...be3acd4](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/74?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @josef-hak on January 29, 2020 at 12:32 PM UTC

@semtexzv there are some conflicts here

---

## Reviews

### Review by @josef-hak - Changes Requested on January 29, 2020 at 03:38 PM UTC

conflicts

### Review by @josef-hak - Changes Requested on January 29, 2020 at 03:39 PM UTC

### Review by @semtexzv - Commented on January 30, 2020 at 08:33 AM UTC

### Review by @semtexzv - Commented on January 30, 2020 at 08:34 AM UTC

### Review by @josef-hak - Approved on January 30, 2020 at 09:56 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/74*
