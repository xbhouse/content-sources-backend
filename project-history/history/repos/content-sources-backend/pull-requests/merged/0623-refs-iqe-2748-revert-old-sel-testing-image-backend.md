---
type: pull_request
number: 623
title: "Refs IQE-2748: Revert old sel testing image backend"
state: merged
author: akhil-jha
created: 2024-04-02T08:51:13Z
updated: 2024-04-02T20:08:17Z
closed: 2024-04-02T20:08:17Z
merged: 2024-04-02T20:08:17Z
base_branch: main
head_branch: revert_old_sel_image
labels: []
url: https://github.com/content-services/content-sources-backend/pull/623
---

# Pull Request #623: Refs IQE-2748: Revert old sel testing image backend

**Author**: @akhil-jha
**Created**: April 02, 2024 at 08:51 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `revert_old_sel_image`

## Description

New sel image is failing the test pipeline because it does not have a controlled shutdown command, making the pod stuck and resulting into timeout.

---

## Discussion

### Comment by @app-sre-bot on April 02, 2024 at 08:52 AM UTC

Can one of the admins verify this patch?

### Comment by @jlsherrill on April 02, 2024 at 09:00 AM UTC

https://issues.redhat.com/browse/HMS-IQE-2748

---

## Reviews

### Review by @mshriver - Approved on April 02, 2024 at 08:08 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/623*
