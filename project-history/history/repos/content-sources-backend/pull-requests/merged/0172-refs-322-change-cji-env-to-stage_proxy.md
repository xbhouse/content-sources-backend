---
type: pull_request
number: 172
title: "Refs 322: Change CJI env to stage_proxy"
state: merged
author: mshriver
created: 2023-01-10T15:54:51Z
updated: 2023-01-10T16:00:52Z
closed: 2023-01-10T16:00:52Z
merged: 2023-01-10T16:00:52Z
base_branch: main
head_branch: change-CJI-env
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/172
---

# Pull Request #172: Refs 322: Change CJI env to stage_proxy

**Author**: @mshriver
**Created**: January 10, 2023 at 03:54 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `change-CJI-env`

## Description

stage_post_deploy isn't modifying the HTTP defaults, so its pointing to production

---

## Discussion

### Comment by @jlsherrill on January 10, 2023 at 04:00 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-322

### Comment by @jlsherrill on January 10, 2023 at 04:00 PM UTC

⚠️ This task currently requires qe-approval, but this PR does not fully resolve the task.  Please contact QE to determine appropriate testing required.

---

## Reviews

### Review by @jlsherrill - Approved on January 10, 2023 at 03:59 PM UTC

lgtm

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/172*
