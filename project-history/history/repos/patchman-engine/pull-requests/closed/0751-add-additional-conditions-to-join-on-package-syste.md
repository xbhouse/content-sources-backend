---
type: pull_request
number: 751
title: "Add additional conditions to join on package systems endpoint"
state: closed
author: semtexzv
created: 2021-07-30T12:33:32Z
updated: 2021-08-02T09:24:43Z
closed: 2021-08-02T09:24:43Z
base_branch: master
head_branch: join-optim
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/751
---

# Pull Request #751: Add additional conditions to join on package systems endpoint

**Author**: @semtexzv
**Created**: July 30, 2021 at 12:33 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `join-optim`

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

### Comment by @codecov-commenter on August 02, 2021 at 07:23 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/751?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#751](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/751?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (7029d53) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/ca7796ee14754408a938f55c2a3e4a3d3933545c?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (ca7796e) will **not change** coverage.
> The diff coverage is `0.00%`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/751/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/751?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@           Coverage Diff           @@
##           master     #751   +/-   ##
=======================================
  Coverage   57.03%   57.03%           
=======================================
  Files          81       81           
  Lines        3833     3833           
=======================================
  Hits         2186     2186           
  Misses       1329     1329           
  Partials      318      318           
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `57.03% <0.00%> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/751?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [base/database/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/751/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS9kYXRhYmFzZS91dGlscy5nbw==) | `0.00% <0.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/751?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/751?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [ca7796e...7029d53](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/751?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @josef-hak on August 02, 2021 at 09:24 AM UTC

dupl. of #752 

---

## Reviews

### Review by @josef-hak - Changes Requested on July 30, 2021 at 01:36 PM UTC

Good, but obviously there is something wrong:
~~~
test         | time="2021-07-30T12:39:31Z" level=info msg="config for aws CloudWatch not loaded"
test         | time="2021-07-30T12:39:31Z" level=warning msg="using mocking account id" account_id=3
test         | time="2021-07-30T12:39:31Z" level=error msg="database error" err="ERROR: column pn.name_id does not exist (SQLSTATE 42703)"
test         | time="2021-07-30T12:39:31Z" level=error msg=request account=3 content_encoding= durationMs=3 method=GET param_package_name=kernel remote_addr= status_code=500 url="/kernel/systems?sort=id" user_agent=
test         | [GIN] 2021/07/30 - 12:39:31 | 500 |    3.793754ms |                 | GET      "/kernel/systems?sort=id"
test         |     package_systems_export_test.go:23: 
test         |         	Error Trace:	package_systems_export_test.go:23
test         |         	Error:      	Not equal: 
test         |         	            	expected: 200
test         |         	            	actual  : 500
test         |         	Test:       	TestPackageSystemsExportHandlerJSON
test         |     test_utils.go:11: 
test         |         	Error Trace:	test_utils.go:11
test         |         	            				package_systems_export_test.go:25
test         |         	Error:      	Expected nil, but got: &json.UnmarshalTypeError{Value:"object", Type:(*reflect.rtype)(0x146eb60), Offset:1, Struct:"", Field:""}
test         |         	Test:       	TestPackageSystemsExportHandlerJSON
test         |         	Messages:   	{"error":"database error"}
test         |     package_systems_export_test.go:26: 
test         |         	Error Trace:	package_systems_export_test.go:26
test         |         	Error:      	Not equal: 
test         |         	            	expected: 2
test         |         	            	actual  : 0
test         |         	Test:       	TestPackageSystemsExportHandlerJSON
platform     | time="2021-07-30T12:39:36Z" level=info msg=request account=0 content_encoding=gzip durationMs=0 method=POST remote_addr="172.18.0.6:57730" status_code=200 url=/api/v3/pkgtree user_agent=OpenAPI-Generator/1.0.0/go
platform     | time="2021-07-30T12:39:36Z" level=info msg=request account=0 content_encoding=gzip durationMs=0 method=POST remote_addr="172.18.0.6:57730" status_code=200 url=/api/v3/pkgtree user_agent=OpenAPI-Generator/1.0.0/go
test         | --- FAIL: TestPackageSystemsExportHandlerJSON (0.01s)
test         | panic: runtime error: index out of range [0] with length 0 [recovered]
test         | 	panic: runtime error: index out of range [0] with length 0
test         | 
test         | goroutine 162 [running]:
test         | testing.tRunner.func1.1(0x15a5780, 0xc001a5d4c0)
test         | 	/usr/lib/golang/src/testing/testing.go:1072 +0x30d
~~~

### Review by @josef-hak - Changes Requested on July 30, 2021 at 01:47 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/751*
