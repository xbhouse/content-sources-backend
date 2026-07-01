---
type: pull_request
number: 144
title: "Fixes 266: use gorm config for batch creation"
state: merged
author: jlsherrill
created: 2022-11-09T22:03:55Z
updated: 2022-11-21T18:31:38Z
closed: 2022-11-21T18:25:43Z
merged: 2022-11-21T18:25:43Z
base_branch: main
head_branch: 266
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/144
---

# Pull Request #144: Fixes 266: use gorm config for batch creation

**Author**: @jlsherrill
**Created**: November 09, 2022 at 10:03 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `266`

## Description

Gorm has a built in way to batch creations, but we were not using it.  By using it, we can handle all creations, and introspect larger repos that were failing before

---

## Discussion

### Comment by @jlsherrill on November 09, 2022 at 10:30 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-266

### Comment by @jlsherrill on November 10, 2022 at 01:38 PM UTC

To test:
Create a repo with this url:  https://jlsherrill.fedorapeople.org/fake-repos/rhel7/
wait a minute or two, verify that the repo is introspected and shows packages.

### Comment by @jlsherrill on November 14, 2022 at 03:00 PM UTC

rebased!

---

## Reviews

### Review by @rverdile - Commented on November 15, 2022 at 03:33 PM UTC

### Review by @rverdile - Commented on November 15, 2022 at 03:34 PM UTC

### Review by @rverdile - Approved on November 15, 2022 at 05:06 PM UTC

lgtm!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/144*
