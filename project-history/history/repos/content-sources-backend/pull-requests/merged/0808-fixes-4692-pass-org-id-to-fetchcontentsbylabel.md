---
type: pull_request
number: 808
title: "Fixes 4692: pass org id to FetchContentsByLabel"
state: merged
author: jlsherrill
created: 2024-09-10T14:49:15Z
updated: 2024-09-10T17:13:11Z
closed: 2024-09-10T17:13:07Z
merged: 2024-09-10T17:13:06Z
base_branch: main
head_branch: 4692
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/808
---

# Pull Request #808: Fixes 4692: pass org id to FetchContentsByLabel

**Author**: @jlsherrill
**Created**: September 10, 2024 at 02:49 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `4692`

## Description

## Summary

FetchContentsByLabel takes an orgID not a templateID.  This wasn't detected in dev because the ID passed in isn't used, instead a global org ID is used.  

There isn't a test for this and right now, its not easily testable.   Proper testing in stage would have caught it, and i've opened this issue to try to come up with a better test 

## Testing steps





---

## Discussion

### Comment by @jlsherrill on September 10, 2024 at 03:00 PM UTC

https://issues.redhat.com/browse/HMS-4692

### Comment by @mayurilahane on September 10, 2024 at 03:11 PM UTC

/retest

### Comment by @jlsherrill on September 10, 2024 at 04:47 PM UTC

/retest

---

## Reviews

### Review by @Andrewgdewar - Approved on September 10, 2024 at 02:55 PM UTC

Ack!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/808*
