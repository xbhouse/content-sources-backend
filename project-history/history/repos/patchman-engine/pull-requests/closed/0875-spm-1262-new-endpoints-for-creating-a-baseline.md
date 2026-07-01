---
type: pull_request
number: 875
title: "SPM-1262: new endpoints for creating a baseline"
state: closed
author: mkholjuraev
created: 2021-12-13T22:13:24Z
updated: 2022-01-13T12:35:39Z
closed: 2022-01-13T12:35:38Z
base_branch: master
head_branch: create-baseline
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/875
---

# Pull Request #875: SPM-1262: new endpoints for creating a baseline

**Author**: @mkholjuraev
**Created**: December 13, 2021 at 10:13 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `create-baseline`

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


 PR for creating a baseline. It includes changes in rbac.go to add write a permission to PUT requests.
 There is an issue with tests from terminal. However, I do not see any errors in VS code testing tool. Please let me know if you know what the problem is. Thanks.


---

## Discussion

### Comment by @josef-hak on January 13, 2022 at 12:35 PM UTC

Included into #884, closing.

---

## Reviews

### Review by @josef-hak - Changes Requested on December 14, 2021 at 02:28 PM UTC

failing test:
~~~
test         | 2021/12/13 22:19:37 [Recovery] 2021/12/13 - 22:19:37 panic recovered:
test         | runtime error: invalid memory address or nil pointer dereference
test         | /usr/lib/golang/src/runtime/panic.go:212 (0x4a2a9a)
test         | 	panicmem: panic(memoryError)
test         | /usr/lib/golang/src/runtime/signal_unix.go:734 (0x4bc872)
test         | 	sigpanic: panicmem()
test         | /go/src/app/manager/controllers/baseline_create.go:69 (0xfb6912)
test         | 	CreateBaselineHandler: LogAndRespError(c, query.Error, "Database error")
test         | /go/pkg/mod/github.com/gin-gonic/gin@v1.7.2/context.go:165 (0xefd849)
test         | 	(*Context).Next: c.handlers[c.index](c)
test         | /go/src/app/manager/middlewares/authentication.go:93 (0xefd82c)
test         | 	MockAuthenticator.func1: c.Next()
test         | /go/pkg/mod/github.com/gin-gonic/gin@v1.7.2/context.go:165 (0xefd91e)
test         | 	(*Context).Next: c.handlers[c.index](c)
test         | /go/src/app/manager/middlewares/logger.go:17 (0xefd8d7)
test         | 	RequestResponseLogger.func1: c.Next()
test         | /go/pkg/mod/github.com/gin-gonic/gin@v1.7.2/context.go:165 (0xcc5939)
test         | 	(*Context).Next: c.handlers[c.index](c)
test         | /go/pkg/mod/github.com/gin-gonic/gin@v1.7.2/recovery.go:99 (0xcc5920)
test         | 	CustomRecoveryWithWriter.func1: c.Next()
test         | /go/pkg/mod/github.com/gin-gonic/gin@v1.7.2/context.go:165 (0xcc4a13)
test         | 	(*Context).Next: c.handlers[c.index](c)
test         | /go/pkg/mod/github.com/gin-gonic/gin@v1.7.2/logger.go:241 (0xcc49d2)
test         | 	LoggerWithConfig.func1: c.Next()
test         | /go/pkg/mod/github.com/gin-gonic/gin@v1.7.2/context.go:165 (0xcbaec9)
test         | 	(*Context).Next: c.handlers[c.index](c)
test         | /go/pkg/mod/github.com/gin-gonic/gin@v1.7.2/gin.go:489 (0xcbaeaf)
test         | 	(*Engine).handleHTTPRequest: c.Next()
test         | /go/pkg/mod/github.com/gin-gonic/gin@v1.7.2/gin.go:445 (0xcba99b)
test         | 	(*Engine).ServeHTTP: engine.handleHTTPRequest(c)
test         | /go/src/app/manager/controllers/baseline_create_test.go:31 (0xfd7292)
test         | 	TestCreateBaseline: core.InitRouterWithParams(CreateBaselineHandler, 1, "PUT", "/").ServeHTTP(w, req)
test         | /usr/lib/golang/src/testing/testing.go:1193 (0x5988ae)
test         | 	tRunner: fn(t)
test         | /usr/lib/golang/src/runtime/asm_amd64.s:1371 (0x4deea0)
test         | 	goexit: BYTE	$0x90	// NOP
test         | 
test         | [GIN] 2021/12/13 - 22:19:37 | 500 |   11.321934ms |                 | PUT      "/"
test         |     baseline_create_test.go:33: 
test         |         	Error Trace:	baseline_create_test.go:33
test         |         	Error:      	Not equal: 
test         |         	            	expected: 200
test         |         	            	actual  : 500
test         |         	Test:       	TestCreateBaseline
~~~

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/875*
