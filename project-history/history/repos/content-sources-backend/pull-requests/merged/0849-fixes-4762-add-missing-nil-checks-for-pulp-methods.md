---
type: pull_request
number: 849
title: "Fixes 4762: Add missing nil checks for pulp methods"
state: merged
author: rverdile
created: 2024-10-15T15:19:19Z
updated: 2024-10-18T20:12:15Z
closed: 2024-10-18T20:12:12Z
merged: 2024-10-18T20:12:12Z
base_branch: main
head_branch: stop-panic
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/849
---

# Pull Request #849: Fixes 4762: Add missing nil checks for pulp methods

**Author**: @rverdile
**Created**: October 15, 2024 at 03:19 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `stop-panic`

## Description

## Summary
I could not reproduce the original issue, but this adds checks to places I believe could cause a panic under the right conditions.

## Testing steps



---

## Discussion

### Comment by @jlsherrill on October 15, 2024 at 03:30 PM UTC

https://issues.redhat.com/browse/HMS-4762

---

## Reviews

### Review by @dominikvagner - Commented on October 17, 2024 at 12:38 PM UTC

one small question, otherwise looks good! 👏🏼 

### Review by @rverdile - Commented on October 17, 2024 at 01:46 PM UTC

### Review by @rverdile - Commented on October 17, 2024 at 01:47 PM UTC

### Review by @rverdile - Commented on October 17, 2024 at 07:04 PM UTC

### Review by @dominikvagner - Approved on October 18, 2024 at 07:26 AM UTC

everything looks good now 🎉 
approved ✅ nice work!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/849*
