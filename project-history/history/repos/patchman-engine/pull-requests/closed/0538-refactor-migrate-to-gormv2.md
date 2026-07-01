---
type: pull_request
number: 538
title: "refactor: migrate to GORMv2"
state: closed
author: AlesKas
created: 2021-02-16T10:21:03Z
updated: 2021-07-12T11:03:01Z
closed: 2021-04-29T15:08:31Z
base_branch: master
head_branch: SPM-568
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/538
---

# Pull Request #538: refactor: migrate to GORMv2

**Author**: @AlesKas
**Created**: February 16, 2021 at 10:21 AM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `SPM-568`

## Description

## Secure Coding Practices Checklist GitHub Link
- https://github.com/RedHatInsights/secure-coding-checklist

## Secure Coding Checklist
- [ ] Input Validation
- [ ] Output Encoding
- [ ] Authentication and Password Management
- [ ] Session Management
- [ ] Access Control
- [ ] Cryptographic Practices
- [ ] Error Handling and Logging
- [ ] Data Protection
- [ ] Communication Security
- [ ] System Configuration
- [ ] Database Security
- [ ] File Management
- [ ] Memory Management
- [ ] General Coding Practices


---

## Discussion

### Comment by @codecov-io on March 02, 2021 at 09:11 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/538?src=pr&el=h1) Report
> Merging [#538](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/538?src=pr&el=desc) (4a3195f) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/9d3f83745aede0e72f577dbd0d8eb191783238ab?el=desc) (9d3f837) will **increase** coverage by `0.00%`.
> The diff coverage is `75.42%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/538/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/538?src=pr&el=tree)

```diff
@@           Coverage Diff           @@
##           master     #538   +/-   ##
=======================================
  Coverage   60.36%   60.37%           
=======================================
  Files          67       67           
  Lines        2942     2983   +41     
=======================================
+ Hits         1776     1801   +25     
- Misses        912      926   +14     
- Partials      254      256    +2     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/538?src=pr&el=tree) | Coverage Δ | |
|---|---|---|
| [base/database/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/538/diff?src=pr&el=tree#diff-YmFzZS9kYXRhYmFzZS90ZXN0aW5nLmdv) | `0.00% <0.00%> (ø)` | |
| [base/database/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/538/diff?src=pr&el=tree#diff-YmFzZS9kYXRhYmFzZS91dGlscy5nbw==) | `0.00% <ø> (ø)` | |
| [manager/controllers/advisories\_export.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/538/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzX2V4cG9ydC5nbw==) | `82.14% <ø> (ø)` | |
| [manager/controllers/filter.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/538/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9maWx0ZXIuZ28=) | `68.18% <ø> (ø)` | |
| [manager/controllers/package\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/538/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX3N5c3RlbXMuZ28=) | `72.72% <ø> (ø)` | |
| [manager/controllers/status.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/538/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zdGF0dXMuZ28=) | `0.00% <0.00%> (ø)` | |
| [manager/controllers/systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/538/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zLmdv) | `89.28% <ø> (ø)` | |
| [manager/controllers/system\_packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/538/diff?src=pr&el=tree#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fcGFja2FnZXMuZ28=) | `66.66% <33.33%> (ø)` | |
| [base/core/probes.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/538/diff?src=pr&el=tree#diff-YmFzZS9jb3JlL3Byb2Jlcy5nbw==) | `58.33% <40.00%> (-16.67%)` | :arrow_down: |
| [base/database/database.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/538/diff?src=pr&el=tree#diff-YmFzZS9kYXRhYmFzZS9kYXRhYmFzZS5nbw==) | `41.17% <40.00%> (-51.14%)` | :arrow_down: |
| ... and [22 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/538/diff?src=pr&el=tree-more) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/538?src=pr&el=continue).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/538?src=pr&el=footer). Last update [9d3f837...4a3195f](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/538?src=pr&el=lastupdated). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments).


### Comment by @codecov-commenter on April 29, 2021 at 01:37 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/538?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#538](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/538?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (c27e648) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/11444200aae9f796c39a1cc9798f8921f8ef5f86?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (1144420) will **increase** coverage by `0.09%`.
> The diff coverage is `72.82%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/538/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/538?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #538      +/-   ##
==========================================
+ Coverage   58.59%   58.68%   +0.09%     
==========================================
  Files          72       72              
  Lines        3251     3297      +46     
==========================================
+ Hits         1905     1935      +30     
- Misses       1068     1081      +13     
- Partials      278      281       +3     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `58.68% <72.82%> (+0.09%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/538?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/database/testing.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/538/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS90ZXN0aW5nLmdv) | `0.00% <0.00%> (ø)` | |
| [base/database/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/538/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS91dGlscy5nbw==) | `0.00% <ø> (ø)` | |
| [manager/controllers/advisories\_export.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/538/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzX2V4cG9ydC5nbw==) | `82.14% <ø> (ø)` | |
| [manager/controllers/filter.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/538/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9maWx0ZXIuZ28=) | `68.18% <ø> (ø)` | |
| [manager/controllers/package\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/538/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX3N5c3RlbXMuZ28=) | `72.72% <ø> (ø)` | |
| [manager/controllers/status.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/538/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zdGF0dXMuZ28=) | `0.00% <0.00%> (ø)` | |
| [manager/controllers/systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/538/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1zLmdv) | `89.28% <ø> (ø)` | |
| [manager/controllers/system\_packages.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/538/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fcGFja2FnZXMuZ28=) | `66.66% <33.33%> (ø)` | |
| [base/core/probes.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/538/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9jb3JlL3Byb2Jlcy5nbw==) | `58.33% <40.00%> (-16.67%)` | :arrow_down: |
| [base/database/database.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/538/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS9kYXRhYmFzZS5nbw==) | `41.17% <40.00%> (-51.14%)` | :arrow_down: |
| ... and [25 more](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/538/diff?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/538?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/538?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [1144420...c27e648](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/538?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @josef-hak on April 29, 2021 at 03:08 PM UTC

We will finish it in #653 

---

## Reviews

### Review by @josef-hak - Commented on February 16, 2021 at 10:32 AM UTC

Firstly, great thanks for doing this big update and make it working and tests passing. Looks good in general. However I would like to know more about some changes so I wrote them down. I wonder whether all changes are needed.

Definitely I would like @semtexzv would check it as he did many gorm-related things in our current implementation.

Also @MichaelMraka 's ideas will be useful as he is our database guru :wink: .

### Review by @AlesKas - Commented on February 16, 2021 at 11:29 AM UTC

### Review by @AlesKas - Commented on February 16, 2021 at 11:34 AM UTC

### Review by @AlesKas - Commented on February 16, 2021 at 11:40 AM UTC

### Review by @AlesKas - Commented on February 16, 2021 at 12:01 PM UTC

### Review by @AlesKas - Commented on February 16, 2021 at 12:50 PM UTC

### Review by @AlesKas - Commented on February 16, 2021 at 01:27 PM UTC

### Review by @AlesKas - Commented on February 16, 2021 at 01:31 PM UTC

### Review by @AlesKas - Commented on February 16, 2021 at 01:37 PM UTC

### Review by @AlesKas - Commented on February 16, 2021 at 01:48 PM UTC

### Review by @AlesKas - Commented on February 16, 2021 at 01:52 PM UTC

### Review by @MichaelMraka - Approved on February 23, 2021 at 05:53 AM UTC

### Review by @semtexzv - Commented on February 23, 2021 at 09:45 AM UTC

Great work, few comments:

### Review by @AlesKas - Commented on February 23, 2021 at 10:04 AM UTC

### Review by @AlesKas - Commented on February 23, 2021 at 10:18 AM UTC

### Review by @semtexzv - Commented on March 02, 2021 at 11:04 AM UTC

### Review by @AlesKas - Commented on March 02, 2021 at 11:14 AM UTC

### Review by @semtexzv - Approved on April 27, 2021 at 12:14 PM UTC

Since we're done with stuff for summit, I think it's the perfect time to merge this change. Due to the size of this PR, there are many possible bugs, but our test suite passes, and we'll have enough time to resolve them. 

So, I'd say, let's do it ! :+1: 

### Review by @josef-hak - Changes Requested on April 29, 2021 at 10:17 AM UTC

some unprocessed conflicts, tests failed
~~~
 go: errors parsing go.mod:
/go/src/app/go.mod:39: malformed module path "<<<<<<<": invalid char '<'
~~~

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/538*
