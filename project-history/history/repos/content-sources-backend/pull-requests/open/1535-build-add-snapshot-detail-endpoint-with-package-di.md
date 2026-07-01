---
type: pull_request
number: 1535
title: "Build: Add snapshot detail endpoint with package diffs"
state: open
author: dominikvagner
created: 2026-06-10T14:53:57Z
updated: 2026-06-11T12:34:04Z
base_branch: main
head_branch: snapshot-diffs
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1535
---

# Pull Request #1535: Build: Add snapshot detail endpoint with package diffs

**Author**: @dominikvagner
**Created**: June 10, 2026 at 02:53 PM UTC
**Status**: Open
**Labels**: None
**Base**: `main` ← **Head**: `snapshot-diffs`

## Description

## Summary
- add `GET /repositories/{repo_uuid}/snapshots/{snapshot_uuid}` for repo-scoped snapshot detail
- include added and removed package diffs from Pulp using the stored snapshot version href, with `name`, `version`, `release`, and `arch`
- cover the new DAO and handler paths with tests and refresh generated mocks plus OpenAPI docs

## Test plan
- [x] `go test ./pkg/dao -run TestSnapshotsSuite`
- [x] `go test ./pkg/handler -run TestSnapshotSuite`


Made with [Cursor](https://cursor.com)

---

## Discussion

### Comment by @xbhouse on June 11, 2026 at 12:34 PM UTC

/retest

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1535*
