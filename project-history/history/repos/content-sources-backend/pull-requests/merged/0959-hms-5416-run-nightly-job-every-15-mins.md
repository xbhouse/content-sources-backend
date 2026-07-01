---
type: pull_request
number: 959
title: "HMS-5416: run nightly job every 15 mins"
state: merged
author: jlsherrill
created: 2025-01-29T14:55:57Z
updated: 2025-02-10T08:19:50Z
closed: 2025-02-10T08:19:50Z
merged: 2025-02-10T08:19:50Z
base_branch: main
head_branch: 5416
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/959
---

# Pull Request #959: HMS-5416: run nightly job every 15 mins

**Author**: @jlsherrill
**Created**: January 29, 2025 at 02:55 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `5416`

## Description

## Summary

runs our nightly job every 15 minutes, and renames it 'process-repos'

## Testing steps

```
make repos-import
```

```
go run cmd/external-repos/main.go process-repos 96
```

You should see all repos snapshotted at first. Anytime you re-run it, you should see 1 get snapshottted (roughly # of repos/96 rounded up)



---

## Discussion

### Comment by @jlsherrill on January 29, 2025 at 03:00 PM UTC

https://issues.redhat.com/browse/HMS-5416

---

## Reviews

### Review by @rverdile - Approved on February 04, 2025 at 06:05 PM UTC

lgtm!

### Review by @swadeley - Commented on February 04, 2025 at 06:40 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/959*
