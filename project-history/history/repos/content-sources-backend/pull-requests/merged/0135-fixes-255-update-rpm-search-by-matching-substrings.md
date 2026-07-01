---
type: pull_request
number: 135
title: "Fixes 255: update rpm search by matching substrings"
state: merged
author: avisiedo
created: 2022-11-02T06:41:24Z
updated: 2022-11-07T15:00:25Z
closed: 2022-11-07T14:46:10Z
merged: 2022-11-07T14:46:10Z
base_branch: main
head_branch: hmscontent-255
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/135
---

# Pull Request #135: Fixes 255: update rpm search by matching substrings

**Author**: @avisiedo
**Created**: November 02, 2022 at 06:41 AM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `hmscontent-255`

## Description

- It matches sub-strings instead of suffix for the package name.

Signed-off-by: Alejandro Visiedo <avisiedo@redhat.com>

---

## Discussion

### Comment by @jlsherrill on November 02, 2022 at 07:00 AM UTC

https://issues.redhat.com/browse/HMSCONTENT-255

### Comment by @avisiedo on November 02, 2022 at 12:01 PM UTC

Unit test error addressed at: https://github.com/content-services/content-sources-backend/pull/135

### Comment by @jlsherrill on November 02, 2022 at 04:19 PM UTC

you might wanna add testing steps for qe :)

---

## Reviews

### Review by @jlsherrill - Approved on November 02, 2022 at 04:13 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/135*
