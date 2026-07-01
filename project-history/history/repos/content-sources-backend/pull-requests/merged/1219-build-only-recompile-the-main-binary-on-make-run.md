---
type: pull_request
number: 1219
title: "Build: Only recompile the main binary on make run"
state: merged
author: jlsherrill
created: 2025-09-29T12:40:04Z
updated: 2025-09-29T17:41:00Z
closed: 2025-09-29T17:40:57Z
merged: 2025-09-29T17:40:57Z
base_branch: main
head_branch: build-small
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1219
---

# Pull Request #1219: Build: Only recompile the main binary on make run

**Author**: @jlsherrill
**Created**: September 29, 2025 at 12:40 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `build-small`

## Description

## Summary

After updating some code and re-running 'make run', the makefile would recompile every binary, which wasn't really necessary.  This limits it to only rebuilding content-sources.

## Testing steps

Modify some dao file or model
run 'make run'

look for:
```
go build  -o "/home/jlsherri/git/content-sources-backend/release/content-sources" "cmd/content-sources/main.go"
```

previously you'd see about ~5 of these.


---

## Discussion

### Comment by @jlsherrill on September 29, 2025 at 01:31 PM UTC

/retest

---

## Reviews

### Review by @xbhouse - Approved on September 29, 2025 at 05:38 PM UTC

lgtm!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1219*
