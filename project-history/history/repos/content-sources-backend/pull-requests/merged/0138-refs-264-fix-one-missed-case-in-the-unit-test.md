---
type: pull_request
number: 138
title: "Refs 264: Fix one missed case in the unit test"
state: merged
author: avisiedo
created: 2022-11-03T13:12:09Z
updated: 2022-11-04T13:39:31Z
closed: 2022-11-04T13:39:31Z
merged: 2022-11-04T13:39:31Z
base_branch: main
head_branch: hmscontent-264-second-fix
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/138
---

# Pull Request #138: Refs 264: Fix one missed case in the unit test

**Author**: @avisiedo
**Created**: November 03, 2022 at 01:12 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `hmscontent-264-second-fix`

## Description

- fix test "force error when message content fails validation", it was providing
  two errors which order is not deterministic; this change fix that providing only
  one failure, making the result deterministic.

Signed-off-by: Alejandro Visiedo <avisiedo@redhat.com>

---

## Discussion

### Comment by @jlsherrill on November 03, 2022 at 01:30 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-264

---

## Reviews

### Review by @jlsherrill - Approved on November 03, 2022 at 07:48 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/138*
