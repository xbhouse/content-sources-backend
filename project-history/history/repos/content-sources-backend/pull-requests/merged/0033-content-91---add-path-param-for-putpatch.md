---
type: pull_request
number: 33
title: "CONTENT-91 - add path param for put/patch"
state: merged
author: jlsherrill
created: 2022-06-14T12:41:46Z
updated: 2022-07-20T08:13:40Z
closed: 2022-06-14T14:08:36Z
merged: 2022-06-14T14:08:36Z
base_branch: main
head_branch: openapi
labels: ["bug", "qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/33
---

# Pull Request #33: CONTENT-91 - add path param for put/patch

**Author**: @jlsherrill
**Created**: June 14, 2022 at 12:41 PM UTC
**Status**: Merged
**Labels**: `bug`, `qe-testing-needed`
**Base**: `main` ← **Head**: `openapi`

## Description

*No description provided*

---

## Discussion

### Comment by @jlsherrill on June 14, 2022 at 12:42 PM UTC

this doesn't actually fix the problem yet, it just verifies the validator will catch the problem

EDIT:  after seeing the generation action failing on 'main', i've added the fix

### Comment by @jlsherrill on July 19, 2022 at 08:20 PM UTC

https://issues.redhat.com/browse/14696935

### Comment by @jlsherrill on July 19, 2022 at 08:22 PM UTC

⚠️ This task currently requires qe-approval, but this PR does not fully resolve the task.  Please contact QE to determine appropriate testing required.

---

## Reviews

### Review by @mshriver - Commented on June 14, 2022 at 01:38 PM UTC

### Review by @mshriver - Approved on June 14, 2022 at 02:02 PM UTC

Pulled this into QE environment and was able to generate a client without failing validation.

### Review by @swadeley - Commented on July 20, 2022 at 08:12 AM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/33*
