---
type: pull_request
number: 295
title: "Refs 1770: add graph for cert expire"
state: merged
author: jlsherrill
created: 2023-05-30T15:57:50Z
updated: 2023-05-30T19:45:40Z
closed: 2023-05-30T19:45:39Z
merged: 2023-05-30T19:45:39Z
base_branch: main
head_branch: 1770_graph
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/295
---

# Pull Request #295: Refs 1770: add graph for cert expire

**Author**: @jlsherrill
**Created**: May 30, 2023 at 03:57 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `1770_graph`

## Description

## Summary

Adds a graph for cdn cert expiration, currently turns red at 45 days.:

![image](https://github.com/content-services/content-sources-backend/assets/395077/8a615524-739c-4500-abae-83847def87f0)




---

## Discussion

### Comment by @jlsherrill on May 30, 2023 at 04:00 PM UTC

https://issues.redhat.com/browse/HMS-1770

### Comment by @rverdile on May 30, 2023 at 04:45 PM UTC

I think that test failure is the same transient one we saw on an earlier PR?

### Comment by @jlsherrill on May 30, 2023 at 05:02 PM UTC

@rverdile yes, i opened https://issues.redhat.com/browse/HMS-1845  as i saw it a couple times last week too

---

## Reviews

### Review by @rverdile - Approved on May 30, 2023 at 04:44 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/295*
