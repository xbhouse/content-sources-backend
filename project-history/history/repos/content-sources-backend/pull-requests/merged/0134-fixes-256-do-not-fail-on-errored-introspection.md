---
type: pull_request
number: 134
title: "Fixes 256: do not fail on errored introspection"
state: merged
author: jlsherrill
created: 2022-11-01T19:25:49Z
updated: 2022-11-01T20:04:43Z
closed: 2022-11-01T20:04:34Z
merged: 2022-11-01T20:04:34Z
base_branch: main
head_branch: 256
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/134
---

# Pull Request #134: Fixes 256: do not fail on errored introspection

**Author**: @jlsherrill
**Created**: November 01, 2022 at 07:25 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `256`

## Description

*No description provided*

---

## Discussion

### Comment by @jlsherrill on November 01, 2022 at 07:27 PM UTC

To test, create some badly formed repository:
```
curl -X POST http://localhost:8000/api/content-sources/v1/repositories/bulk_create/  -d '[{"name":"foobar2", "url":"https://jlsherrill.fedorapeople.org/badurl/"}]'  -H "`./scripts/header.sh 1 1`"   -H "Content-Type: application/json" 
```

Run introspect all:
```
go run cmd/external-repos/main.go introspect-all
```

notice the better failure and zero exit code:
```
{"level":"error","error":"Error introspecting http://foobar.com/foobar: Status error: 403","time":"2022-11-01T15:18:51-04:00","message":"Introspection Error"}
```

```
echo $?
0
```

### Comment by @jlsherrill on November 01, 2022 at 07:30 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-256

---

## Reviews

### Review by @Andrewgdewar - Approved on November 01, 2022 at 08:03 PM UTC

LGTM! ⚡ 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/134*
