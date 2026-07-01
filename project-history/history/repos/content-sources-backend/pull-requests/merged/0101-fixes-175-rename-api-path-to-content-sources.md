---
type: pull_request
number: 101
title: "Fixes 175: rename api path to content-sources"
state: merged
author: jlsherrill
created: 2022-09-08T12:56:46Z
updated: 2022-09-08T16:30:32Z
closed: 2022-09-08T16:01:27Z
merged: 2022-09-08T16:01:26Z
base_branch: main
head_branch: 175
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/101
---

# Pull Request #101: Fixes 175: rename api path to content-sources

**Author**: @jlsherrill
**Created**: September 08, 2022 at 12:56 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `175`

## Description

because insights does not like content_sources
This also refactors a bit to only set the app name once in code

---

## Discussion

### Comment by @jlsherrill on September 08, 2022 at 01:00 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-175

### Comment by @mshriver on September 08, 2022 at 03:14 PM UTC

/retest


---

## Reviews

### Review by @mshriver - Approved on September 08, 2022 at 03:04 PM UTC

### Review by @rverdile - Approved on September 08, 2022 at 03:58 PM UTC

lgtm. endpoints still function as normal, including skipping middleware on /ping

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/101*
