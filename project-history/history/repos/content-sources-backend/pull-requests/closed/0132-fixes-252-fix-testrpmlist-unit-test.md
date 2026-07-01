---
type: pull_request
number: 132
title: "Fixes 252: Fix TestRpmList unit test"
state: closed
author: avisiedo
created: 2022-10-31T12:36:20Z
updated: 2022-11-01T12:48:29Z
closed: 2022-11-01T04:48:04Z
base_branch: main
head_branch: hmscontent-252
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/132
---

# Pull Request #132: Fixes 252: Fix TestRpmList unit test

**Author**: @avisiedo
**Created**: October 31, 2022 at 12:36 PM UTC
**Status**: Closed
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `hmscontent-252`

## Description

Update limit to `-1` into TestRpmList

Signed-off-by: Alejandro Visiedo <avisiedo@redhat.com>

---

## Discussion

### Comment by @jlsherrill on October 31, 2022 at 01:00 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-252

### Comment by @avisiedo on November 01, 2022 at 04:48 AM UTC

closing as I was not able to replay again from the `main` branch.

### Comment by @jlsherrill on November 01, 2022 at 11:55 AM UTC

i suspect that maybe the problem actually has to do with go 1.18 in the pr https://github.com/content-services/content-sources-backend/pull/125

have we seen it elsewhere?

### Comment by @jlsherrill on November 01, 2022 at 12:48 PM UTC

I was actually able to reproduce this locally randomly :)  
I fixed this here:  https://github.com/content-services/content-sources-backend/pull/125/files#diff-73ba5cff26e1cda8426ce77b3206b3b368b202a1ce40582dc8a37f96f23f217b
i think its related to the upgraded of go and deps, its treating limit 0 as actual limit 0, i'm not 100% why it wasn't doing that before

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/132*
