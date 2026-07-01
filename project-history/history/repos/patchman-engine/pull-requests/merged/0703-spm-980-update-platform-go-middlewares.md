---
type: pull_request
number: 703
title: "SPM-980: Update platform-go-middlewares"
state: merged
author: semtexzv
created: 2021-06-16T10:35:58Z
updated: 2021-06-18T09:17:53Z
closed: 2021-06-18T09:17:53Z
merged: 2021-06-18T09:17:53Z
base_branch: master
head_branch: update-middeleware-libs
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/703
---

# Pull Request #703: SPM-980: Update platform-go-middlewares

**Author**: @semtexzv
**Created**: June 16, 2021 at 10:35 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `update-middeleware-libs`

## Description

SteveHNH has fixed an issue in this library that caused some missing logs

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

### Comment by @codecov-commenter on June 16, 2021 at 10:43 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/703?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#703](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/703?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (f05326b) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/8541db51ffda52a7a9b1f2005cdc53bfbcc9854b?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (8541db5) will **decrease** coverage by `0.09%`.
> The diff coverage is `n/a`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/703/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/703?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #703      +/-   ##
==========================================
- Coverage   56.55%   56.45%   -0.10%     
==========================================
  Files          75       77       +2     
  Lines        3517     3592      +75     
==========================================
+ Hits         1989     2028      +39     
- Misses       1245     1267      +22     
- Partials      283      297      +14     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `56.45% <ø> (-0.10%)` | :arrow_down: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/703?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/database/setup.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/703/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS9zZXR1cC5nbw==) | `70.83% <0.00%> (-1.90%)` | :arrow_down: |
| [base/mqueue/event.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/703/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9tcXVldWUvZXZlbnQuZ28=) | `27.58% <0.00%> (ø)` | |
| [manager/controllers/advisory\_systems\_export.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/703/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yeV9zeXN0ZW1zX2V4cG9ydC5nbw==) | `50.00% <0.00%> (ø)` | |
| [manager/controllers/system\_advisories\_export.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/703/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9zeXN0ZW1fYWR2aXNvcmllc19leHBvcnQuZ28=) | `54.28% <0.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/703?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/703?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [8541db5...f05326b](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/703?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @josef-hak - Changes Requested on June 17, 2021 at 10:48 AM UTC

Command `go mod tidy` updates go.sum file for me:
~~~
 github.com/alecthomas/units v0.0.0-20190717042225-c3de453c63f4/go.mod h1:ybxpYRFXyAe+OPACYpWeL0wqObRcbAqCMya13uyzqw0=
 github.com/apache/arrow/go/arrow v0.0.0-20200601151325-b2287a20f230/go.mod h1:QNYViu/X0HXDHw7m3KXzWSVXIbfUvJqBFe6Gj8/pYA0=
 github.com/aws/aws-sdk-go v1.17.7/go.mod h1:KmX6BPdI08NWTb3/sm4ZGu5ShLoqVDhKgpiN924inxo=
-github.com/aws/aws-sdk-go v1.30.9 h1:DntpBUKkchINPDbhEzDRin1eEn1TG9TZFlzWPf0i8to=
-github.com/aws/aws-sdk-go v1.30.9/go.mod h1:5zCpMtNQVjRREroY7sYe8lOMRSxkhG6MZveU8YkpAk0=
+github.com/aws/aws-sdk-go v1.38.51 h1:aKQmbVbwOCuQSd8+fm/MR3bq0QOsu9Q7S+/QEND36oQ=
 github.com/aws/aws-sdk-go v1.38.51/go.mod h1:hcU610XS61/+aQV88ixoOzUoG7v3b31pl2zKMmprdro=
 github.com/beorn7/perks v0.0.0-20180321164747-3a771d992973/go.mod h1:Dwedo/Wpr24TaqPxmxbtue+5NUziq4I4S80YR8gNf3Q=
 github.com/beorn7/perks v1.0.0/go.mod h1:KWe93zE9D1o94FZ5RNwFwVgaQK1VOXiVxmqh+CedLV8=
@@ -312,9 +311,9 @@ github.com/jinzhu/inflection v1.0.0/go.mod h1:h+uFLlag+Qp1Va5pdKtLDYj+kHp5pxUVkr
 github.com/jinzhu/now v1.1.1 h1:g39TucaRWyV3dwDO++eEc6qf8TVIQ/Da48WmqjZ3i7E=
 github.com/jinzhu/now v1.1.1/go.mod h1:d3SSVoowX0Lcu0IBviAWJpolVfI5UJVZZ7cO71lE/z8=
 github.com/jmespath/go-jmespath v0.0.0-20180206201540-c2b33e8439af/go.mod h1:Nht3zPeWKUH0NzdCt2Blrr5ys8VGpn0CEB0cQHVjt7k=
-github.com/jmespath/go-jmespath v0.3.0 h1:OS12ieG61fsCg5+qLJ+SsW9NicxNkg3b25OyT2yCeUc=
-github.com/jmespath/go-jmespath v0.3.0/go.mod h1:9QtRXoHjLGCJ5IBSaohpXITPlowMeeYCZ7fLUTSywik=
+github.com/jmespath/go-jmespath v0.4.0 h1:BEgLn5cpjn8UN1mAw4NjwDrS35OdebyEtFe+9YPoQUg=
 github.com/jmespath/go-jmespath v0.4.0/go.mod h1:T8mJZnbsbmF+m6zOOFylbeCJqk5+pHWvzYPziyZiYoo=
+github.com/jmespath/go-jmespath/internal/testify v1.5.1 h1:shLQSRRSCCPj3f2gpwzGwWFoC7ycTf1rcQZHOlsJ6N8=
 github.com/jmespath/go-jmespath/internal/testify v1.5.1/go.mod h1:L3OGu8Wl2/fWfCI6z80xFu9LTZmf1ZRjMHUOPmWr69U=
...
~~~

### Review by @josef-hak - Approved on June 18, 2021 at 09:17 AM UTC

Good, thanks.

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/703*
