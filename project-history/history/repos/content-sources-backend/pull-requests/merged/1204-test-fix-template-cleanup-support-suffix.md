---
type: pull_request
number: 1204
title: "Test: Fix template cleanup support suffix"
state: merged
author: swadeley
created: 2025-09-17T13:15:09Z
updated: 2025-09-17T15:00:01Z
closed: 2025-09-17T15:00:01Z
merged: 2025-09-17T15:00:01Z
base_branch: main
head_branch: swadeley/fix_template_delete
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1204
---

# Pull Request #1204: Test: Fix template cleanup support suffix

**Author**: @swadeley
**Created**: September 17, 2025 at 01:15 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `swadeley/fix_template_delete`

## Description

## Summary

Make template cleanup accept prefix name

## Testing steps

check that template tests cleanup afterwards


---

## Discussion

### Comment by @marusak on September 17, 2025 at 01:40 PM UTC

> imho a language design flaw of typescript 

Would not be the only one in JS :D But yeah, the fact that `in` iterates over keys, while `of` over values is mindboggling :D 

---

## Reviews

### Review by @dominikvagner - Approved on September 17, 2025 at 01:26 PM UTC

ahh, oops **my bad** 🤦🏼‍♂️ 
don't know why I used `in` here and the correct one in the other cleanup func 🫣

_the fact both of these exist, do something different and `for ... in ...` is used by several other languages is imho a language design flaw of typescript 🙄😅_

acked! thanks for fixing this! 🙇🏼‍♂️ ✅ 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1204*
