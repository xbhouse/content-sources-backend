---
type: pull_request
number: 682
title: "SPM-925: enable os column sorting"
state: closed
author: mkholjuraev
created: 2021-05-24T09:29:19Z
updated: 2021-06-10T07:32:26Z
closed: 2021-06-10T07:32:26Z
base_branch: master
head_branch: fix-os-sort
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/682
---

# Pull Request #682: SPM-925: enable os column sorting

**Author**: @mkholjuraev
**Created**: May 24, 2021 at 09:29 AM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `fix-os-sort`

## Description

## Secure Coding Practices Checklist GitHub Link
- https://github.com/RedHatInsights/secure-coding-checklist

## Secure Coding Checklist
- [x] Input Validation
- [x] Output Encoding
- [x] Authentication and Password Management
- [x] Session Management
- [x] Access Control
- [x] Cryptographic Practices
- [x] Error Handling and Logging
- [x] Data Protection
- [x] Communication Security
- [x] System Configuration
- [x] Database Security
- [x] File Management
- [x] Memory Management
- [x] General Coding Practices


---

## Discussion

### Comment by @mkholjuraev on May 24, 2021 at 09:32 AM UTC

PR for [SPM-925](https://issues.redhat.com/browse/SPM-925). Some lines in utils.go is a bit tedious. From UI side we shall send os_name as sorting parameter.  But it is sorting as expected taking account os_minor and os_major version also. Ideas for improvement are welcome :)

### Comment by @codecov-commenter on May 24, 2021 at 09:36 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/682?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#682](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/682?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (dd881e3) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/92f43c2e6aca29d2cee9dca6488003402a435421?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (92f43c2) will **decrease** coverage by `0.02%`.
> The diff coverage is `75.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/682/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/682?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #682      +/-   ##
==========================================
- Coverage   58.56%   58.54%   -0.03%     
==========================================
  Files          72       72              
  Lines        3367     3389      +22     
==========================================
+ Hits         1972     1984      +12     
- Misses       1111     1119       +8     
- Partials      284      286       +2     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `58.54% <75.00%> (-0.03%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/682?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [manager/controllers/advisory\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/682/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zLmdv) | `75.55% <ø> (ø)` | |
| [manager/controllers/systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/682/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zLmdv) | `89.28% <ø> (ø)` | |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/682/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `78.28% <75.00%> (-0.40%)` | :arrow_down: |
| [base/mqueue/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/682/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvdGVzdGluZy5nbw==) | `75.00% <0.00%> (-25.00%)` | :arrow_down: |
| [base/mqueue/mqueue.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/682/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvbXF1ZXVlLmdv) | `51.28% <0.00%> (-16.52%)` | :arrow_down: |
| [listener/events.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/682/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bGlzdGVuZXIvZXZlbnRzLmdv) | `45.00% <0.00%> (ø)` | |
| [base/mqueue/event.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/682/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvZXZlbnQuZ28=) | `27.58% <0.00%> (ø)` | |
| [base/mqueue/message.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/682/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvbWVzc2FnZS5nbw==) | `0.00% <0.00%> (ø)` | |
| [base/mqueue/encryption.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/682/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvZW5jcnlwdGlvbi5nbw==) | | |
| [base/mqueue/mqueue\_impl\_gokafka.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/682/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvbXF1ZXVlX2ltcGxfZ29rYWZrYS5nbw==) | `78.84% <0.00%> (ø)` | |
| ... and [1 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/682/diff?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/682?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/682?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [92f43c2...dd881e3](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/682?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @josef-hak on May 27, 2021 at 03:40 AM UTC

@mkholjuraev it's not rebased with master, right? Could you do it, pls? It should fix the 4th testing pipeline.

### Comment by @semtexzv on June 01, 2021 at 11:30 AM UTC

@Josca No particular reason. 

### Comment by @josef-hak on June 02, 2021 at 12:44 PM UTC

/retest

---

## Reviews

### Review by @MichaelMraka - Approved on May 24, 2021 at 11:44 AM UTC

nice

### Review by @josef-hak - Changes Requested on May 24, 2021 at 11:56 AM UTC

Good, could you please add test for new sorting option to both endpoints?

### Review by @josef-hak - Changes Requested on May 24, 2021 at 12:21 PM UTC

Nice, just please look at my one note. And please. Could you please use SPM-XYZ to mark both PR title and commit message? Look at e.g. #684.

### Review by @semtexzv - Commented on May 24, 2021 at 12:33 PM UTC

Try splitting up the long lines. Everything else looks ok.

### Review by @josef-hak - Approved on May 25, 2021 at 07:01 AM UTC

Just small improvement suggested. But can be merged imo.

### Review by @josef-hak - Changes Requested on May 25, 2021 at 07:56 PM UTC

### Review by @josef-hak - Changes Requested on May 25, 2021 at 08:22 PM UTC

### Review by @josef-hak - Changes Requested on May 27, 2021 at 03:04 AM UTC

### Review by @mkholjuraev - Commented on May 30, 2021 at 10:59 PM UTC

### Review by @josef-hak - Changes Requested on May 31, 2021 at 04:51 AM UTC

### Review by @mkholjuraev - Commented on May 31, 2021 at 08:43 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/682*
