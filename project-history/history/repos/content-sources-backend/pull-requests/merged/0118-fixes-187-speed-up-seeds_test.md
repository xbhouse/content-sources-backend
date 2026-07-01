---
type: pull_request
number: 118
title: "Fixes 187: speed up seeds_test"
state: merged
author: Andrewgdewar
created: 2022-09-30T17:39:52Z
updated: 2022-10-03T13:39:52Z
closed: 2022-10-03T13:39:52Z
merged: 2022-10-03T13:39:52Z
base_branch: main
head_branch: CONTENT-187
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/118
---

# Pull Request #118: Fixes 187: speed up seeds_test

**Author**: @Andrewgdewar
**Created**: September 30, 2022 at 05:39 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `CONTENT-187`

## Description

Currently seed test is the slowest test file, taking over a second.  

Some of this is due to 'batch size' for inserting, and a desire to test a number greater than that.  

This ticket makes the batch size configurable so the seeding can test with a loweramount

---

## Discussion

### Comment by @jlsherrill on September 30, 2022 at 06:00 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-187

---

## Reviews

### Review by @jlsherrill - Approved on October 03, 2022 at 11:56 AM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/118*
