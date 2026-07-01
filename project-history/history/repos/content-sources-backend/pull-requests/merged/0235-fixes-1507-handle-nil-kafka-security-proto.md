---
type: pull_request
number: 235
title: "Fixes 1507: handle nil kafka security proto"
state: merged
author: jlsherrill
created: 2023-03-24T14:34:12Z
updated: 2023-03-24T14:50:07Z
closed: 2023-03-24T14:48:48Z
merged: 2023-03-24T14:48:47Z
base_branch: main
head_branch: 1507
labels: []
url: https://github.com/content-services/content-sources-backend/pull/235
---

# Pull Request #235: Fixes 1507: handle nil kafka security proto

**Author**: @jlsherrill
**Created**: March 24, 2023 at 02:34 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `1507`

## Description

## Summary
Fixes this error:
```
{"level":"warn","time":"2023-03-24T08:00:51Z","message":"config.yaml
 file not loaded: Config File \"config.yaml\" Not Found in 
\"[/configs]\""}{
"level":"debug","platform.content-sources.introspect":"platform-mq-stage.platform.content-sources.introspect","time":"2023-03-24T08:00:51Z","message":"internalToReal"}
{"level":"debug","platform-mq-stage.platform.content-sources.introspect":"platform.content-sources.introspect","time":"2023-03-24T08:00:51Z","message":"realToInternal"}
panic: runtime error: invalid memory address or nil pointer dereference[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x9244d0]
goroutine 1 [running]:github.com/content-services/content-sources-backend/pkg/config.addEventConfigDefaults(0xc0004a6340?)	/go/src/app/pkg/config/event.go:50 +0x790github.com/content-services/content-sources-backend/pkg/config.setDefaults(0x10ea1a0?)	
/go/src/app/pkg/config/config.go:144 +0x305github.com/content-services/content-sources-backend/pkg/config.Load()	
/go/src/app/pkg/config/config.go:154 +0x147main.main()	/go/src/app/cmd/external-repos/main.go:24 +0x5b
```
## Testing steps
IDK that there is an easy way to test this, as it requires clowder to specify this information, and we may just have to see it work in stage

---

## Reviews

### Review by @rverdile - Approved on March 24, 2023 at 02:40 PM UTC

ack

### Review by @swadeley - Commented on March 24, 2023 at 02:48 PM UTC

lgtm

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/235*
