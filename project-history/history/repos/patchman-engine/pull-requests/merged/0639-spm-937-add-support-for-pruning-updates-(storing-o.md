---
type: pull_request
number: 639
title: "SPM-937: Add support for pruning updates (storing only latest)"
state: merged
author: semtexzv
created: 2021-04-23T13:03:21Z
updated: 2021-06-22T12:19:55Z
closed: 2021-06-22T12:19:55Z
merged: 2021-06-22T12:19:55Z
base_branch: master
head_branch: prune-updates
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/639
---

# Pull Request #639: SPM-937: Add support for pruning updates (storing only latest)

**Author**: @semtexzv
**Created**: April 23, 2021 at 01:03 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `prune-updates`

## Description

If we don't use intermediate updates anywhere, why store them ? 

---

## Discussion

### Comment by @codecov-commenter on April 23, 2021 at 01:12 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/639?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#639](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/639?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (517cf47) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/1b045996f1cf57fd873942cfefba40dd4509788f?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (1b04599) will **decrease** coverage by `0.07%`.
> The diff coverage is `30.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/639/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/639?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #639      +/-   ##
==========================================
- Coverage   58.53%   58.45%   -0.08%     
==========================================
  Files          72       72              
  Lines        3364     3375      +11     
==========================================
+ Hits         1969     1973       +4     
- Misses       1111     1118       +7     
  Partials      284      284              
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `58.45% <30.00%> (-0.08%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/639?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/database/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/639/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS90ZXN0aW5nLmdv) | `0.00% <0.00%> (ø)` | |
| [evaluator/evaluate.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/639/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlLmdv) | `66.33% <100.00%> (+0.17%)` | :arrow_up: |
| [evaluator/evaluate\_packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/639/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-ZXZhbHVhdG9yL2V2YWx1YXRlX3BhY2thZ2VzLmdv) | `68.91% <100.00%> (+0.42%)` | :arrow_up: |
| [manager/controllers/package\_detail.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/639/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX2RldGFpbC5nbw==) | `83.33% <0.00%> (+0.35%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/639?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/639?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [1b04599...517cf47](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/639?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @josef-hak on May 27, 2021 at 04:12 AM UTC

... also please assign a Jira ticket for that.

---

## Reviews

### Review by @josef-hak - Changes Requested on April 23, 2021 at 01:22 PM UTC

evaluator/evaluate_test.go:107: File is not `gofmt`-ed with `-s` (gofmt)

@semtexzv run `golangci-lint run` before push. It will show you issues like this.

### Review by @MichaelMraka - Changes Requested on April 26, 2021 at 08:35 AM UTC

### Review by @semtexzv - Commented on April 26, 2021 at 08:37 AM UTC

### Review by @josef-hak - Changes Requested on April 26, 2021 at 07:57 AM UTC

### Review by @semtexzv - Commented on April 26, 2021 at 09:20 AM UTC

### Review by @semtexzv - Commented on April 26, 2021 at 09:21 AM UTC

### Review by @semtexzv - Commented on April 26, 2021 at 09:21 AM UTC

### Review by @semtexzv - Commented on April 26, 2021 at 09:21 AM UTC

### Review by @MichaelMraka - Commented on April 26, 2021 at 09:56 AM UTC

### Review by @josef-hak - Changes Requested on April 26, 2021 at 11:28 AM UTC

### Review by @semtexzv - Commented on April 28, 2021 at 07:17 AM UTC

### Review by @josef-hak - Changes Requested on April 28, 2021 at 01:15 PM UTC

### Review by @semtexzv - Commented on April 29, 2021 at 09:22 AM UTC

### Review by @josef-hak - Changes Requested on May 03, 2021 at 02:22 PM UTC

test         | evaluator/evaluate.go:25: File is not `gofmt`-ed with `-s` (gofmt)

### Review by @josef-hak - Changes Requested on May 10, 2021 at 08:18 AM UTC

still failing test

### Review by @josef-hak - Commented on May 12, 2021 at 01:10 PM UTC

### Review by @josef-hak - Commented on May 27, 2021 at 04:09 AM UTC

@semtexzv There is some unanswered suggestion by @MichaelMraka. Could you process it please?

### Review by @MichaelMraka - Changes Requested on May 31, 2021 at 08:18 AM UTC

### Review by @josef-hak - Commented on June 18, 2021 at 09:20 AM UTC

We have some failing tests here, @semtexzv can you update it, please?

### Review by @josef-hak - Approved on June 22, 2021 at 12:19 PM UTC

Ok, thanks.

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/639*
