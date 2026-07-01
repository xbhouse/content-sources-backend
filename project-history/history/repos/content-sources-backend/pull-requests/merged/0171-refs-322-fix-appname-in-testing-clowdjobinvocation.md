---
type: pull_request
number: 171
title: "Refs 322: Fix appName in testing clowdjobinvocation param"
state: merged
author: mshriver
created: 2023-01-09T22:34:46Z
updated: 2023-01-09T22:43:43Z
closed: 2023-01-09T22:43:43Z
merged: 2023-01-09T22:43:43Z
base_branch: main
head_branch: fix-app-name
labels: []
url: https://github.com/content-services/content-sources-backend/pull/171
---

# Pull Request #171: Refs 322: Fix appName in testing clowdjobinvocation param

**Author**: @mshriver
**Created**: January 09, 2023 at 10:34 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `fix-app-name`

## Description

```
message: Some Jobs are still incomplete
reason: ClowdApp.cloud.redhat.com "content-sources" not found
status: 'False'
type: JobInvocationComplete
```

Resolving the above failure when this CJI tries to start

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/171*
