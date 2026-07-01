---
type: pull_request
number: 277
title: "Implement system tagging"
state: merged
author: semtexzv
created: 2020-06-22T10:24:19Z
updated: 2021-03-16T09:28:32Z
closed: 2020-07-20T10:14:27Z
merged: 2020-07-20T10:14:27Z
base_branch: master
head_branch: tags
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/277
---

# Pull Request #277: Implement system tagging

**Author**: @semtexzv
**Created**: June 22, 2020 at 10:24 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `tags`

## Description

- Adds locally stored tags
- Allows implements tag filtering on /systems endpoints
- Does not perform deduplication of tags by moving them to separate table. Current implementation has simpler evaluation.

---

## Discussion

### Comment by @codecov-commenter on June 24, 2020 at 06:33 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/277?src=pr&el=h1) Report
> Merging [#277](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/277?src=pr&el=desc) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/acae0c620c69af0f54ebc8c28e97e9703cdccb29&el=desc) will **increase** coverage by `0.36%`.
> The diff coverage is `88.57%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/277/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/277?src=pr&el=tree)

```diff
@@            Coverage Diff             @@
##           master     #277      +/-   ##
==========================================
+ Coverage   62.20%   62.57%   +0.36%     
==========================================
  Files          51       51              
  Lines        2408     2442      +34     
==========================================
+ Hits         1498     1528      +30     
- Misses        722      724       +2     
- Partials      188      190       +2     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/277?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [listener/upload.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/277/diff?src=pr&el=tree#diff-bGlzdGVuZXIvdXBsb2FkLmdv) | `79.72% <80.95%> (+0.13%)` | :arrow_up: |
| [manager/controllers/advisory\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/277/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `77.08% <100.00%> (+0.99%)` | :arrow_up: |
| [manager/controllers/systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/277/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zLmdv) | `85.18% <100.00%> (+0.27%)` | :arrow_up: |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/277/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `78.57% <100.00%> (+1.84%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/277?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/277?src=pr&el=footer). Last update [acae0c6...fa0bc66](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/277?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


---

## Reviews

### Review by @josef-hak - Commented on June 30, 2020 at 11:19 AM UTC

Thanks for update, added some comments and questions.

### Review by @semtexzv - Commented on July 01, 2020 at 12:10 PM UTC

### Review by @semtexzv - Commented on July 01, 2020 at 12:10 PM UTC

### Review by @semtexzv - Commented on July 03, 2020 at 07:59 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/277*
