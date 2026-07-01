---
type: pull_request
number: 1220
title: "HMS-9442: Only include lifecycle dates for matching rhel release"
state: merged
author: jlsherrill
created: 2025-09-29T12:54:18Z
updated: 2025-10-07T13:40:17Z
closed: 2025-10-07T13:40:13Z
merged: 2025-10-07T13:40:13Z
base_branch: main
head_branch: modularity_fix
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/1220
---

# Pull Request #1220: HMS-9442: Only include lifecycle dates for matching rhel release

**Author**: @jlsherrill
**Created**: September 29, 2025 at 12:54 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `modularity_fix`

## Description

## Summary

This adds filtering on the roadmap/lifecycle api responses to only include 'package' type app-streams if the osMajor matches at least one red hat repo passed into the search

## Testing steps

on a fresh db run:

```
OPTIONS_REPOSITORY_IMPORT_FILTER=rhel8 go run ./cmd/external-repos/main.go import
go run cmd/external-repos/main.go process-repos
```

and then curl:
```
curl -X POST --location "http://localhost:8000/api/content-sources/v1.0/rpms/names" \
    -H "x-rh-identity: eyJpZGVudGl0eSI6eyJvcmdfaWQiOiI5IiwgInR5cGUiOiJVc2VyIiwidXNlciI6eyJ1c2VybmFtZSI6ImZvbyJ9LCJhY2NvdW50X251bWJlciI6ImZvbyIsImludGVybmFsIjp7Im9yZ19pZCI6IjkifX19Cg==" \
    -H "Content-Type: application/json" \
    -d '{
          "urls": ["https://cdn.redhat.com/content/dist/rhel8/8/x86_64/appstream/os/"],
          "search": "ansible-core",
          "include_package_sources": true
        }'
```

In this case, you should only get a single value:

```
      {
        "type": "package",
        "name": "ansible-core",
        "stream": "2.12",
        "start_date": "2022-05-10",
        "end_date": "2023-11-10"
      }
```

but on main you get 3.

---

## Discussion

### Comment by @jlsherrill on September 29, 2025 at 01:00 PM UTC

https://issues.redhat.com/browse/HMS-9442

---

## Reviews

### Review by @xbhouse - Approved on October 02, 2025 at 07:50 PM UTC

lgtm!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1220*
