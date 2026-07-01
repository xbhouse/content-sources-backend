---
type: pull_request
number: 84
title: "Refs 121: introspect one public repo on startup"
state: merged
author: jlsherrill
created: 2022-08-23T17:44:24Z
updated: 2022-08-26T15:12:46Z
closed: 2022-08-26T15:12:46Z
merged: 2022-08-26T15:12:46Z
base_branch: main
head_branch: 121
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/84
---

# Pull Request #84: Refs 121: introspect one public repo on startup

**Author**: @jlsherrill
**Created**: August 23, 2022 at 05:44 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `121`

## Description

This introspects a single repositroy at startup such that
QE can better test integration points

This repo would get introspected once a day as it is

---

## Discussion

### Comment by @jlsherrill on August 23, 2022 at 06:00 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-121

### Comment by @jlsherrill on August 23, 2022 at 06:00 PM UTC

⚠️ This task currently requires qe-approval, but this PR does not fully resolve the task.  Please contact QE to determine appropriate testing required.

### Comment by @rverdile on August 24, 2022 at 07:37 PM UTC

should this have gotten the "Ready For QE" flag automatically? Maybe it didn't work because it's a "Refs"?

### Comment by @jlsherrill on August 25, 2022 at 01:08 PM UTC

yeah, its because its a Refs PR and no fixes Pr, which isn't something handled very well.

Multiple GH prs per issue makes things complicated ;)

---

## Reviews

### Review by @rverdile - Approved on August 23, 2022 at 08:21 PM UTC

tested and I can see that the repository was introspected 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/84*
