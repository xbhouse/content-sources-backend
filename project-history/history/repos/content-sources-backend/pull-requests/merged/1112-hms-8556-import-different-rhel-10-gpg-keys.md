---
type: pull_request
number: 1112
title: "HMS-8556: import different RHEL 10 gpg keys"
state: merged
author: jlsherrill
created: 2025-05-16T15:12:09Z
updated: 2025-05-16T16:18:02Z
closed: 2025-05-16T16:16:13Z
merged: 2025-05-16T16:16:13Z
base_branch: main
head_branch: 8556
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/1112
---

# Pull Request #1112: HMS-8556: import different RHEL 10 gpg keys

**Author**: @jlsherrill
**Created**: May 16, 2025 at 03:12 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `8556`

## Description

## Summary

While RHEL 10 uses the same GPG keys as rhel 9, the format is different for one of the keys and image builder is failing on a step

## Testing steps

on main:
# make repos-import

switch to this PR:
# make repos-import
# make db-cli-connect

you should see the new gpg keys with this query:
```
select gpg_key from repository_configurations where name ilike '%10%' limit 1;
```

and the old ones with this query:
```
select gpg_key from repository_configurations where name ilike '%9%' limit 1;
```


---

## Discussion

### Comment by @jlsherrill on May 16, 2025 at 03:30 PM UTC

https://issues.redhat.com/browse/HMS-8556

### Comment by @jlsherrill on May 16, 2025 at 04:16 PM UTC

merging, will verify in stage

---

## Reviews

### Review by @rverdile - Approved on May 16, 2025 at 03:16 PM UTC

works!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1112*
