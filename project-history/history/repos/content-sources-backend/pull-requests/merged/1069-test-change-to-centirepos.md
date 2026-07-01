---
type: pull_request
number: 1069
title: "Test: Change to centirepos"
state: merged
author: swadeley
created: 2025-04-09T14:24:50Z
updated: 2025-09-08T07:31:28Z
closed: 2025-04-30T08:42:16Z
merged: 2025-04-30T08:42:16Z
base_branch: main
head_branch: swadeley/change_to_centirepos
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1069
---

# Pull Request #1069: Test: Change to centirepos

**Author**: @swadeley
**Created**: April 09, 2025 at 02:24 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `swadeley/change_to_centirepos`

## Description



## Summary

I made a new batch of 100 repos without a number in the URL
The multirepos path with a rank number in URL causes problems when that number corresponds to a release version.

## Testing steps

tests pass

---

## Discussion

### Comment by @rverdile on April 10, 2025 at 01:59 PM UTC

looks good but there's a linting issue

### Comment by @rverdile on April 10, 2025 at 02:01 PM UTC

Doesn't break anything, but wanted to point out that some the repos are named with a leading 0 e.g. `repo03`. I'm not sure if you meant to include those.

### Comment by @rverdile on April 10, 2025 at 02:19 PM UTC

Oh and there's no `repo1`, so I think that would be an issue if the randomURL function returned `repo1`

### Comment by @swadeley on April 14, 2025 at 04:07 PM UTC

> Oh and there's no `repo1`, so I think that would be an issue if the randomURL function returned `repo1`

I think that code is OK, grok says:
randomNum: Exports a string representing numbers from 1 to 100, padded as "01" to "100".

### Comment by @swadeley on April 14, 2025 at 04:10 PM UTC

> Doesn't break anything, but wanted to point out that some the repos are named with a leading 0 e.g. `repo03`. I'm not sure if you meant to include those.

Hi, yes, I prefer to use leading zeros so that where the repo numbers are used in the names it helps with sorting sorting (looks better, easier to read, in pagination tests). I have now deleted repos 2 to 9 without leading zeros which were created by mistake.

Thank you

### Comment by @swadeley on April 14, 2025 at 07:43 PM UTC

/retest

### Comment by @rverdile on April 24, 2025 at 07:28 PM UTC

@swadeley once you update the URL here, should be good to go

### Comment by @swadeley on April 25, 2025 at 08:24 PM UTC

> @swadeley once you update the URL here, should be good to go

Hi @rverdile , OK, rebased and update to use `github.io` URL.

Thank you

---

## Reviews

### Review by @rverdile - Approved on April 29, 2025 at 01:14 PM UTC

lgtm!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1069*
