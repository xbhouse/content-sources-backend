---
type: pull_request
number: 139
title: "Fixes 262: ignore metalinks in external repos download"
state: merged
author: jlsherrill
created: 2022-11-03T13:25:18Z
updated: 2022-11-07T20:30:02Z
closed: 2022-11-07T20:29:58Z
merged: 2022-11-07T20:29:57Z
base_branch: main
head_branch: 262
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/139
---

# Pull Request #139: Fixes 262: ignore metalinks in external repos download

**Author**: @jlsherrill
**Created**: November 03, 2022 at 01:25 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `262`

## Description

image builder started supporting fedora but is using metalinks which we do not support.  This ignores those and also fixes a bug & deprecated method use.

---

## Discussion

### Comment by @jlsherrill on November 03, 2022 at 01:25 PM UTC

so in the end, there was no new repositories added, this just fixes the download/save process

### Comment by @jlsherrill on November 03, 2022 at 01:30 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-262

---

## Reviews

### Review by @rverdile - Approved on November 07, 2022 at 07:47 PM UTC

two small comments otherwise tested and lgtm!

### Review by @jlsherrill - Commented on November 07, 2022 at 07:57 PM UTC

### Review by @jlsherrill - Commented on November 07, 2022 at 08:09 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/139*
