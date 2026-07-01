---
type: pull_request
number: 657
title: "Build: add make repos-import-rhel9"
state: merged
author: jlsherrill
created: 2024-05-01T17:46:30Z
updated: 2024-05-07T15:58:05Z
closed: 2024-05-07T15:58:00Z
merged: 2024-05-07T15:58:00Z
base_branch: main
head_branch: repos_import_rhel9
labels: []
url: https://github.com/content-services/content-sources-backend/pull/657
---

# Pull Request #657: Build: add make repos-import-rhel9

**Author**: @jlsherrill
**Created**: May 01, 2024 at 05:46 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `repos_import_rhel9`

## Description

## Summary

Support importing only a subset of repos for testing rhel9 (for example), and a handy make target.

## Testing steps

on a fresh db:  
```
make repos-import-rhel9
```

then check the red hat repos list and see only these two red hat repos

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @swadeley on May 03, 2024 at 12:37 PM UTC

Hi

I deployed with bonfire and:
 `--set-parameter content-sources-backend/OPTIONS_REPOSITORY_IMPORT_FILTER=rhel9`

and under the Red Hat tab I can see the baseos and appstream rhel9 repos are present and no others.

ACK from me

Thank you

### Comment by @swadeley on May 03, 2024 at 05:31 PM UTC

@rverdile merge on ACK please

---

## Reviews

### Review by @rverdile - Commented on May 01, 2024 at 07:08 PM UTC

### Review by @jlsherrill - Commented on May 03, 2024 at 05:21 PM UTC

### Review by @rverdile - Approved on May 07, 2024 at 01:45 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/657*
