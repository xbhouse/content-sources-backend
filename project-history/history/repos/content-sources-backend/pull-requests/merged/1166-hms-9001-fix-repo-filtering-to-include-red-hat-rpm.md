---
type: pull_request
number: 1166
title: "HMS-9001: fix repo filtering to include red hat rpms in search results"
state: merged
author: xbhouse
created: 2025-08-05T14:57:39Z
updated: 2025-08-05T18:28:59Z
closed: 2025-08-05T18:28:59Z
merged: 2025-08-05T18:28:58Z
base_branch: main
head_branch: fix-query
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/1166
---

# Pull Request #1166: HMS-9001: fix repo filtering to include red hat rpms in search results

**Author**: @xbhouse
**Created**: August 05, 2025 at 02:57 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `fix-query`

## Description

## Summary

This [PR](https://github.com/content-services/content-sources-backend/pull/1160) surfaced a bug in the rpm search where if both popular repos and red hat repos were imported, rpms for red hat repos would not be found using the repo url

## Testing steps

1. On main, import and introspect/snapshot the small RH repo and a popular repo: 
```
OPTIONS_REPOSITORY_IMPORT_FILTER=small make repos-import
go run cmd/external-repos/main.go snapshot --url https://cdn.redhat.com/content/dist/rhel9/9/aarch64/codeready-builder/os/
OPTIONS_REPOSITORY_IMPORT_FILTER=epel10 make repos-import
go run cmd/external-repos/main.go introspect --url https://dl.fedoraproject.org/pub/epel/10/Everything/x86_64/ --force
```
2. Search for any rpms using the codeready repo url, you'll see this return an empty array:
```
POST /rpms/names
{
  "search": "",
  "urls": ["https://cdn.redhat.com/content/dist/rhel9/9/aarch64/codeready-builder/os/"]
}
```
3. Check out this PR and repeat step 2. You should see rpms in the response



---

## Discussion

### Comment by @jlsherrill on August 05, 2025 at 03:00 PM UTC

https://issues.redhat.com/browse/HMS-9001

---

## Reviews

### Review by @rverdile - Commented on August 05, 2025 at 05:02 PM UTC

### Review by @rverdile - Approved on August 05, 2025 at 05:02 PM UTC

good find!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1166*
