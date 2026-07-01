---
type: pull_request
number: 613
title: "fixed dereference bug"
state: merged
author: MichaelMraka
created: 2021-04-14T11:31:30Z
updated: 2021-04-14T13:28:00Z
closed: 2021-04-14T13:28:00Z
merged: 2021-04-14T13:27:59Z
base_branch: master
head_branch: pr2
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/613
---

# Pull Request #613: fixed dereference bug

**Author**: @MichaelMraka
**Created**: April 14, 2021 at 11:31 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pr2`

## Description

{"@timestamp":"2021-04-14T11:20:00.791Z","count":500,"level":"debug","message":"Downloaded advisories"}
2021/04/14 11:20:00 http: panic serving [::1]:39904: runtime error: invalid memory address or nil pointer dereference
goroutine 4063 [running]:
net/http.(*conn).serve.func1(0xc0004c4780)
	/usr/lib/golang/src/net/http/server.go:1808 +0x13e
panic(0xf7f600, 0x240baa0)
	/usr/lib/golang/src/runtime/panic.go:975 +0x3e3
app/vmaas_sync.parseAdvisories(0xc0013580f0, 0x5, 0x0, 0xc0006be060, 0x15, 0xc000126078)
	/go/src/app/vmaas_sync/advisory_sync.go:78 +0xc23
app/vmaas_sync.storeAdvisories(0xc0013580f0, 0x5, 0xc0004713c0, 0x1)
	/go/src/app/vmaas_sync/advisory_sync.go:122 +0x43
app/vmaas_sync.syncAdvisories(0x0, 0x0)
	/go/src/app/vmaas_sync/advisory_sync.go:179 +0xb41
app/vmaas_sync.syncData(0xc0003ca070, 0x4)
	/go/src/app/vmaas_sync/vmaas_sync.go:127 +0x26
app/vmaas_sync.sync(0xc000690000)
	/go/src/app/vmaas_sync/admin_api.go:33 +0x9b
github.com/gin-gonic/gin.(*Context).Next(0xc000690000)
	/go/pkg/mod/github.com/gin-gonic/gin@v1.6.3/context.go:161 +0x3b
github.com/gin-gonic/gin.(*Engine).handleHTTPRequest(0xc0002fa000, 0xc000690000)
	/go/pkg/mod/github.com/gin-gonic/gin@v1.6.3/gin.go:409 +0x666
github.com/gin-gonic/gin.(*Engine).ServeHTTP(0xc0002fa000, 0x13a6b00, 0xc0007b3880, 0xc000442800)
	/go/pkg/mod/github.com/gin-gonic/gin@v1.6.3/gin.go:367 +0x14d
net/http.serverHandler.ServeHTTP(0xc0007b22a0, 0x13a6b00, 0xc0007b3880, 0xc000442800)
	/usr/lib/golang/src/net/http/server.go:2848 +0xa3
net/http.(*conn).serve(0xc0004c4780, 0x13abe00, 0xc000484a00)
	/usr/lib/golang/src/net/http/server.go:1936 +0x86c
created by net/http.(*Server).Serve
	/usr/lib/golang/src/net/http/server.go:2974 +0x361

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

### Comment by @semtexzv on April 14, 2021 at 12:11 PM UTC

Test failed, You need to manually check for nil value.

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/613*
