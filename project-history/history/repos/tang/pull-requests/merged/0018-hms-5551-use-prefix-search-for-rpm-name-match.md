---
type: pull_request
number: 18
title: "HMS-5551: use prefix search for rpm name match"
state: merged
author: jlsherrill
created: 2025-02-27T21:52:06Z
updated: 2025-03-10T14:06:11Z
closed: 2025-03-10T14:06:05Z
merged: 2025-03-10T14:06:05Z
base_branch: main
head_branch: 5551
labels: []
url: https://github.com/content-services/tang/pull/18
---

# Pull Request #18: HMS-5551: use prefix search for rpm name match

**Author**: @jlsherrill
**Created**: February 27, 2025 at 09:52 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `5551`

## Description

*No description provided*

---

## Discussion

### Comment by @jlsherrill on March 05, 2025 at 09:11 PM UTC

doh! i think i backed out one too many changes!

### Comment by @jlsherrill on March 05, 2025 at 09:45 PM UTC

thx, updated!

---

## Reviews

### Review by @rverdile - Commented on March 04, 2025 at 08:38 PM UTC

snapshot RPM search is still returning a result for me if I search for a suffix. I think it's because there's a leading wildcard here: https://github.com/content-services/tang/blob/8feff3369981d0518aa6bec4be39c5404a010f20/pkg/tangy/rpm.go#L115

### Review by @rverdile - Approved on March 06, 2025 at 03:24 PM UTC

looks good!

---

*Archived from: https://github.com/content-services/tang/pull/18*
