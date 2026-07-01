---
type: pull_request
number: 838
title: "SPM-1185: Implement the order_query reflection attribute"
state: merged
author: semtexzv
created: 2021-10-04T09:02:03Z
updated: 2021-10-04T13:26:03Z
closed: 2021-10-04T13:26:03Z
merged: 2021-10-04T13:26:03Z
base_branch: master
head_branch: order-query
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/838
---

# Pull Request #838: SPM-1185: Implement the order_query reflection attribute

**Author**: @semtexzv
**Created**: October 04, 2021 at 09:02 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `order-query`

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

### Comment by @codecov-commenter on October 04, 2021 at 12:13 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/838?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#838](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/838?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (c0da3ab) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/ac3e50cc3d8c2753c19856c316fa35313055ef4f?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (ac3e50c) will **increase** coverage by `0.18%`.
> The diff coverage is `81.81%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/838/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/838?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #838      +/-   ##
==========================================
+ Coverage   58.18%   58.37%   +0.18%     
==========================================
  Files          81       81              
  Lines        3951     3957       +6     
==========================================
+ Hits         2299     2310      +11     
+ Misses       1328     1325       -3     
+ Partials      324      322       -2     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `58.37% <81.81%> (+0.18%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/838?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [manager/controllers/filter.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/838/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9maWx0ZXIuZ28=) | `68.18% <55.55%> (ø)` | |
| [base/database/query.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/838/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS9xdWVyeS5nbw==) | `82.27% <100.00%> (+1.45%)` | :arrow_up: |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/838/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `82.51% <100.00%> (+1.34%)` | :arrow_up: |
| [manager/controllers/advisories\_export.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/838/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzX2V4cG9ydC5nbw==) | `89.28% <0.00%> (+7.14%)` | :arrow_up: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/838?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/838?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [125567f...c0da3ab](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/838?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @josef-hak - Changes Requested on October 04, 2021 at 09:27 AM UTC

Some test error:
~~~
 test         | manager/controllers/filter_test.go:56:48: cannot use dummyParser (value of type func(v string) (interface{}, error)) as string value in struct literal (typecheck)
test         | 		attrMap := database.AttrMap{"test": {"test", dummyParser}}
test         | 		                                             ^
test         | manager/controllers/filter_test.go:76:54: cannot use dummyParser (value of type func(v string) (interface{}, error)) as string value in struct literal (typecheck)
test         | 		attrMap := database.AttrMap{"test": {"(NOT test)", dummyParser}}
~~~

And pls include SPM-1185, thx

### Review by @josef-hak - Commented on October 04, 2021 at 09:29 AM UTC

### Review by @semtexzv - Commented on October 04, 2021 at 11:57 AM UTC

### Review by @josef-hak - Changes Requested on October 04, 2021 at 12:04 PM UTC

Test fail:
~~~
test         | === RUN   TestOrderQuery
test         |     query_test.go:83: 
test         |         	Error Trace:	query_test.go:83
test         |         	Error:      	Not equal: 
test         |         	            	expected: ""
test         |         	            	actual  : "REVERSE(am.text_note)"
test         |         	            	
test         |         	            	Diff:
test         |         	            	--- Expected
test         |         	            	+++ Actual
test         |         	            	@@ -1 +1 @@
test         |         	            	-
test         |         	            	+REVERSE(am.text_note)
test         |         	Test:       	TestOrderQuery
test         | --- FAIL: TestOrderQuery (0.00s)
~~~

### Review by @josef-hak - Approved on October 04, 2021 at 01:25 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/838*
