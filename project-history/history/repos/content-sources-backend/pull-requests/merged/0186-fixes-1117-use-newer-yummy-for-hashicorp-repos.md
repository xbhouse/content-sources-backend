---
type: pull_request
number: 186
title: "Fixes 1117: use newer yummy for hashicorp repos"
state: merged
author: jlsherrill
created: 2023-01-26T15:37:45Z
updated: 2023-01-26T21:23:28Z
closed: 2023-01-26T21:23:28Z
merged: 2023-01-26T21:23:28Z
base_branch: main
head_branch: 1117
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/186
---

# Pull Request #186: Fixes 1117: use newer yummy for hashicorp repos

**Author**: @jlsherrill
**Created**: January 26, 2023 at 03:37 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `1117`

## Description

## Summary
Allows introspection of hashicorp repos
## Testing steps

Create a repo with url of 'https://rpm.releases.hashicorp.com/RHEL/9/x86_64/stable/'
Wait for it to be introspected.

Requires https://github.com/content-services/yummy/pull/10 to be merged and tagged.

---

## Discussion

### Comment by @jlsherrill on January 26, 2023 at 04:00 PM UTC

https://issues.redhat.com/browse/HMS-1117

### Comment by @swadeley on January 26, 2023 at 08:30 PM UTC

/retest

### Comment by @swadeley on January 26, 2023 at 09:23 PM UTC

Working now, thank you.

---

## Reviews

### Review by @rverdile - Approved on January 26, 2023 at 09:22 PM UTC

lgtm

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/186*
