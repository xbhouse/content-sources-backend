---
type: pull_request
number: 43
title: "fix: increase max XML size to 1.5GB"
state: merged
author: xbhouse
created: 2026-05-20T17:05:54Z
updated: 2026-05-20T17:43:58Z
closed: 2026-05-20T17:43:58Z
merged: 2026-05-20T17:43:58Z
base_branch: master
head_branch: increase-max-xml-size
labels: []
url: https://github.com/content-services/yummy/pull/43
---

# Pull Request #43: fix: increase max XML size to 1.5GB

**Author**: @xbhouse
**Created**: May 20, 2026 at 05:05 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `increase-max-xml-size`

## Description

Some repos (RHEL 9.8 EUS/E4S aarch64 baseos) exceed 1GB. This increases the max XML size yummy can parse.

Testing steps:

1. Use your local yummy directory as the backend yummy package in content-sources-backend/go.mod and run `go mod tidy`:

`replace [github.com/content-services/yummy](http://github.com/content-services/yummy) => /path/to/yummy`

2. Try to introspect a large repo (https://cdn.redhat.com/content/e4s/rhel9/9.8/aarch64/baseos/os/ or https://cdn.redhat.com/content/eus/rhel9/9.8/aarch64/baseos/os/). It shouldn't fail. 

3. Without this PR, you'll see an error like this: 

```
An introspection error occurred: XML syntax error on line 12654284: unexpected EOF
```

---

## Reviews

### Review by @rverdile - Approved on May 20, 2026 at 05:31 PM UTC

---

*Archived from: https://github.com/content-services/yummy/pull/43*
