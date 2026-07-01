---
type: pull_request
number: 695
title: "run test kafka instance directly from git"
state: merged
author: MichaelMraka
created: 2021-06-04T13:07:22Z
updated: 2021-09-14T11:16:42Z
closed: 2021-09-14T11:16:42Z
merged: 2021-09-14T11:16:42Z
base_branch: master
head_branch: pr2
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/695
---

# Pull Request #695: run test kafka instance directly from git

**Author**: @MichaelMraka
**Created**: June 04, 2021 at 01:07 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pr2`

## Description

requires #699 


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

### Comment by @codecov-commenter on June 08, 2021 at 11:28 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/695?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#695](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/695?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (4ddb368) into [master](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/2e072717ea9abe20119e20d63c4750d2337c8e6b?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (2e07271) will **increase** coverage by `0.06%`.
> The diff coverage is `n/a`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/695/graphs/tree.svg?width=650&height=150&src=pr&token=I7QOHXwVxB&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/695?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@            Coverage Diff             @@
##           master     #695      +/-   ##
==========================================
+ Coverage   56.12%   56.19%   +0.06%     
==========================================
  Files          80       81       +1     
  Lines        3695     3726      +31     
==========================================
+ Hits         2074     2094      +20     
- Misses       1308     1315       +7     
- Partials      313      317       +4     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `56.19% <ø> (+0.06%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/695?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [manager/controllers/package\_systems.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/695/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX3N5c3RlbXMuZ28=) | `72.72% <0.00%> (ø)` | |
| [manager/controllers/package\_versions.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/695/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9wYWNrYWdlX3ZlcnNpb25zLmdv) | `64.51% <0.00%> (ø)` | |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/695?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/695?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [2e07271...4ddb368](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/695?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @MichaelMraka on June 21, 2021 at 01:12 PM UTC

SAN cert issue has been fixed by different PR.

---

## Reviews

### Review by @josef-hak - Changes Requested on June 04, 2021 at 02:07 PM UTC

Test failed:
~~~
test         | === RUN   TestRoundTrip
test         |     mqueue_test.go:43: 
test         |         	Error Trace:	mqueue_test.go:43
test         |         	Error:      	Received unexpected error:
test         |         	            	x509: certificate relies on legacy Common Name field, use SANs or temporarily enable Common Name matching with GODEBUG=x509ignoreCN=0
test         |         	Test:       	TestRoundTrip
test         | time="2021-06-04T13:55:56Z" level=error msg="Unable to establish connection to consumer group coordinator for group patchman: x509: certificate relies on legacy Common Name field, use SANs or temporarily enable Common Name matching with GODEBUG=x509ignoreCN=0" type=kafka
test         | time="2021-06-04T13:55:56Z" level=error msg="x509: certificate relies on legacy Common Name field, use SANs or temporarily enable Common Name matching with GODEBUG=x509ignoreCN=0" type=kafka
test         | time="2021-06-04T13:56:01Z" level=error msg="Unable to establish connection to consumer group coordinator for group patchman: x509: certificate relies on legacy Common Name field, use SANs or temporarily enable Common Name matching with GODEBUG=x509ignoreCN=0" type=kafka
test         | time="2021-06-04T13:56:01Z" level=error msg="x509: certificate relies on legacy Common Name field, use SANs or temporarily enable Common Name matching with GODEBUG=x509ignoreCN=0" type=kafka
test         |     testing.go:54: 
test         |         	Error Trace:	testing.go:54
test         |         	            				mqueue_test.go:44
test         |         	Error:      	Not equal: 
test         |         	            	expected: "99c0ffee-0000-0000-0000-0000000050de"
test         |         	            	actual  : ""
test         |         	            	
test         |         	            	Diff:
test         |         	            	--- Expected
test         |         	            	+++ Actual
test         |         	            	@@ -1 +1 @@
test         |         	            	-99c0ffee-0000-0000-0000-0000000050de
test         |         	            	+
test         |         	Test:       	TestRoundTrip
test         | --- FAIL: TestRoundTrip (8.11s)
~~~

### Review by @semtexzv - Approved on September 10, 2021 at 11:15 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/695*
