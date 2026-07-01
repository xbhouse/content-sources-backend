---
type: pull_request
number: 1394
title: "RHINENG-7099: fix evaluator-recalc nil pointer dereference"
state: merged
author: Dugowitch
created: 2024-03-07T15:55:27Z
updated: 2024-09-26T10:11:20Z
closed: 2024-03-08T10:50:21Z
merged: 2024-03-08T10:50:21Z
base_branch: master
head_branch: RHINENG-7099
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1394
---

# Pull Request #1394: RHINENG-7099: fix evaluator-recalc nil pointer dereference

**Author**: @Dugowitch
**Created**: March 07, 2024 at 03:55 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `RHINENG-7099`

## Description

## Why the error occurred 
The SIGSEGV originates in function `HTTPCallRetry` at line 40 (file *base/utils/http.go*):
```go
resp.Body.Close()
```
The `resp` comes from `httpCallFun` that is passed into the function as a parameter:
```go
func HTTPCallRetry(ctx context.Context, httpCallFun func() (outputDataPtr interface{},
	resp *http.Response, err error), exponentialRetry bool, maxRetries int, codesToRetry ...int) (
	outputDataPtr interface{}, err error) { ... }
```
```go
outDataPtr, resp, callErr := httpCallFun()
```
`HTTPCallRetry` function is called in `callVMaas` (file *evaluator/evaluate.go*) and `vmaasCallFunc` is passed in as `httpCallFun`.
```go
vmaasDataPtr, err := utils.HTTPCallRetry(base.Context, vmaasCallFunc, vmaasCallUseExpRetry,
	vmaasCallMaxRetries, http.StatusServiceUnavailable)
```
```go
vmaasCallFunc := func() (interface{}, *http.Response, error) { ... }
```
In `vmaasCallFunc`, `resp` comes from a `Request` method (file *base/api/client.go*). 
```go
resp, err := vmaasClient.Request(&ctx, http.MethodPost, vmaasUpdatesURL, request, &vmaasData)
```
```go
func (o *Client) Request(ctx *context.Context, method, url string,
	requestPtr interface{}, responseOutPtr interface{}) (*http.Response, error) { ... }
```
In the request method, when the request fails, the method returns *nil, err*. 
```go
httpReq, err := http.NewRequestWithContext(*ctx, method, url, body)
if err != nil {
	return nil, errors.Wrap(err, "Request failed")
}
```
Therefore `resp` in `vmaasCallFunc` as well as `resp` in `HTTPCallRetry` can be nil. A nil check is missing in `HTTPCallRetry`.

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

### Comment by @jira-linking on March 07, 2024 at 03:55 PM UTC

Referenced Jiras:
https://issues.redhat.com/browse/RHINENG-7099


### Comment by @codecov-commenter on March 07, 2024 at 04:00 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1394?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: Patch coverage is `0%` with `2 lines` in your changes are missing coverage. Please review.
> Project coverage is 63.67%. Comparing base [(`25f765e`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/25f765ec1438368c4e3d9d1d2d9094362619b9ac?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`eec5a8b`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1394?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Report is 134 commits behind head on master.

| [Files](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1394?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [base/utils/http.go](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1394?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-YmFzZS91dGlscy9odHRwLmdv) | 0.00% | [2 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1394?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1394      +/-   ##
==========================================
+ Coverage   61.82%   63.67%   +1.85%     
==========================================
  Files         108      107       -1     
  Lines        6814     6517     -297     
==========================================
- Hits         4213     4150      -63     
+ Misses       2062     1878     -184     
+ Partials      539      489      -50     
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1394/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1394/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `63.67% <0.00%> (+1.85%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.


</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1394?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @psegedy on March 08, 2024 at 10:50 AM UTC

Great job!
Looking at the code it seems like `nil` is only returned `Request` when we fail to make request not when we receive a http error code as I initially though. We might have incorrect data in the DB and see some other issues like pod restarts when this gets to stage, we will need to monitor it.

---

## Reviews

### Review by @psegedy - Approved on March 08, 2024 at 10:50 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1394*
