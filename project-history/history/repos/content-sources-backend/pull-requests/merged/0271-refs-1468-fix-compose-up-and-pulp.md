---
type: pull_request
number: 271
title: "Refs 1468: fix compose up and pulp"
state: merged
author: jlsherrill
created: 2023-05-15T16:55:16Z
updated: 2023-05-18T15:03:41Z
closed: 2023-05-18T15:02:53Z
merged: 2023-05-18T15:02:53Z
base_branch: main
head_branch: 1468_2
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/271
---

# Pull Request #271: Refs 1468: fix compose up and pulp

**Author**: @jlsherrill
**Created**: May 15, 2023 at 04:55 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `1468_2`

## Description

## Summary
the previous change enabled pulp during compose commands by default improperly, this fixes it


## Testing steps
```make compose-clean compose-up```

check if pulp is enabled: `podman ps`

---

## Discussion

### Comment by @jlsherrill on May 15, 2023 at 05:00 PM UTC

https://issues.redhat.com/browse/HMS-1468

### Comment by @mshriver on May 16, 2023 at 04:29 PM UTC

/retest

### Comment by @swadeley on May 16, 2023 at 05:16 PM UTC

heh, openapi validate test failed `Can't load config class with name 'python'`

### Comment by @swadeley on May 16, 2023 at 05:16 PM UTC

/retest

---

## Reviews

### Review by @swadeley - Approved on May 16, 2023 at 05:21 PM UTC

lgtm

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/271*
