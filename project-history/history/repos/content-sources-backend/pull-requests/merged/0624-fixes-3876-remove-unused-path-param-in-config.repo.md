---
type: pull_request
number: 624
title: "Fixes 3876: remove unused path param in config.repo spec"
state: merged
author: croissanne
created: 2024-04-02T12:45:05Z
updated: 2024-04-02T19:30:22Z
closed: 2024-04-02T19:17:11Z
merged: 2024-04-02T19:17:11Z
base_branch: main
head_branch: path-param-fix
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/624
---

# Pull Request #624: Fixes 3876: remove unused path param in config.repo spec

**Author**: @croissanne
**Created**: April 02, 2024 at 12:45 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `path-param-fix`

## Description

The path only has a single parameter, being the identifier of the snapshot.

---

Error I encountered:
```
+ /home/sanne/go/bin/go1.20.14 generate -mod=mod ./...
error generating code: error creating operation definitions: path '/snapshots/{snapshot_uuid}/config.repo' has 1 positional parameters, but spec has 2 declared
exit status 1
```
with
```
//go:generate go run -mod=mod github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen --config client.cfg.yaml content-sources.v1.json
```
---

## Summary

## Testing steps

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @app-sre-bot on April 02, 2024 at 12:46 PM UTC

Can one of the admins verify this patch?

### Comment by @jlsherrill on April 02, 2024 at 01:00 PM UTC

https://issues.redhat.com/browse/HMS-3876

### Comment by @jlsherrill on April 02, 2024 at 02:42 PM UTC

[test]

---

## Reviews

### Review by @jlsherrill - Approved on April 02, 2024 at 02:01 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/624*
