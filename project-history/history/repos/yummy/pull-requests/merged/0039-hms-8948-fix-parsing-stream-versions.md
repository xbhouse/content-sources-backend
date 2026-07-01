---
type: pull_request
number: 39
title: "HMS-8948: fix parsing stream versions"
state: merged
author: rverdile
created: 2025-11-12T21:43:00Z
updated: 2025-11-26T14:01:24Z
closed: 2025-11-26T14:01:20Z
merged: 2025-11-26T14:01:20Z
base_branch: master
head_branch: decode-hook
labels: []
url: https://github.com/content-services/yummy/pull/39
---

# Pull Request #39: HMS-8948: fix parsing stream versions

**Author**: @rverdile
**Created**: November 12, 2025 at 09:43 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `decode-hook`

## Description

Stream versions that were floats with trailing zeros were not being parsed correctly. i.e. perl 5.30 -> perl 5.3. Adds custom unmarshal for streams so that the trailing zeros are preserved.

---

## Discussion

### Comment by @rverdile on November 13, 2025 at 03:32 PM UTC

@TenSt did you get a chance to test this through the backend?

---

## Reviews

### Review by @TenSt - Approved on November 12, 2025 at 10:14 PM UTC

lgtm

---

*Archived from: https://github.com/content-services/yummy/pull/39*
