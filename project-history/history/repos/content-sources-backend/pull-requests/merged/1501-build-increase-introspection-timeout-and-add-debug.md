---
type: pull_request
number: 1501
title: "Build: increase introspection timeout and add debug logs"
state: merged
author: xbhouse
created: 2026-05-21T21:03:01Z
updated: 2026-05-26T14:13:18Z
closed: 2026-05-26T14:13:18Z
merged: 2026-05-26T14:13:18Z
base_branch: main
head_branch: increase-introspection-timeout
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1501
---

# Pull Request #1501: Build: increase introspection timeout and add debug logs

**Author**: @xbhouse
**Created**: May 21, 2026 at 09:03 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `increase-introspection-timeout`

## Description

## Summary

- Doubles the client request timeout to 6 minutes for introspection of RH repos, keeps the response header timeout as is (3 minutes)
- Adds debug logs to each introspection step to make it easier to see where it times out

## Testing steps

Import and introspect a few larger RH repos:

```
make repos-import && go run cmd/external-repos/main.go introspect --url https://cdn.redhat.com/content/eus/rhel9/9.8/aarch64/baseos/os --url https://cdn.redhat.com/content/e4s/rhel9/9.8/aarch64/baseos/os --url https://cdn.redhat.com/content/dist/rhel9/9/aarch64/baseos/os --url https://cdn.redhat.com/content/dist/rhel9/9/aarch64/baseos/os --url https://cdn.redhat.com/content/e4s/rhel10/10.0/aarch64/appstream/os --url https://cdn.redhat.com/content/eus/rhel10/10.0/aarch64/appstream/os --url https://cdn.redhat.com/content/e4s/rhel10/10.0/aarch64/baseos/os --url https://cdn.redhat.com/content/e4s/rhel10/10.0/x86_64/appstream/os --url https://cdn.redhat.com/content/eus/rhel10/10.0/x86_64/appstream/os --url https://cdn.redhat.com/content/eus/rhel10/10.0/x86_64/baseos/os
```

This will likely introspect everything fine locally, we might only notice a difference in stage/prod. You could lower the timeouts locally to verify errors and messages are logged correctly


---

## Reviews

### Review by @dominikvagner - Approved on May 26, 2026 at 08:02 AM UTC

ack ✅ thanks 🚀 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1501*
