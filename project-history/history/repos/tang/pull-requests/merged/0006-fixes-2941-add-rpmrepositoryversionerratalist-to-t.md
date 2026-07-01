---
type: pull_request
number: 6
title: "Fixes 2941: Add RpmRepositoryVersionErrataList to tangy"
state: merged
author: Andrewgdewar
created: 2024-03-15T21:28:12Z
updated: 2024-04-17T18:55:12Z
closed: 2024-04-17T18:55:12Z
merged: 2024-04-17T18:55:12Z
base_branch: main
head_branch: HMS-2941
labels: []
url: https://github.com/content-services/tang/pull/6
---

# Pull Request #6: Fixes 2941: Add RpmRepositoryVersionErrataList to tangy

**Author**: @Andrewgdewar
**Created**: March 15, 2024 at 09:28 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `HMS-2941`

## Description

Adds a function to list errata in repository versions, supports pagination and filtering on errata id, type, and severity

Can be tested with this [backend PR](https://github.com/content-services/content-sources-backend/pull/608)

---

## Discussion

### Comment by @xbhouse on April 16, 2024 at 10:47 PM UTC

i haven't found a repo with packages that include severities in the errata, so i haven't been able to really test that filter yet here. i've verified it works with the backend and UI, but if you know of any repos i could use to test here let me know

### Comment by @swadeley on April 17, 2024 at 09:38 AM UTC

> i haven't found a repo with packages that include severities in the errata, so i haven't been able to really test that filter yet here. i've verified it works with the backend and UI, but if you know of any repos i could use to test here let me know

Hi, this repo has severity:
https://stephenw.fedorapeople.org/fakerepos/multiple_errata/

### Comment by @xbhouse on April 17, 2024 at 02:25 PM UTC

@swadeley do the package errata in that repo have an updatedDate? trying to view that repo's package errata has surfaced an error: `can't scan into dest[6]: cannot scan NULL into *string"`

these are the values expected and i'm just assuming updatedDate is missing since it's the 7th one:
`id, errataId, title, summary, description, issuedDate, updatedDate, type, severity, rebootSuggested`

if updatedDate can be null (which makes sense if it can), i'll need to add a change 

### Comment by @xbhouse on April 17, 2024 at 02:34 PM UTC

ok, changed updatedDate to *string and can now view the errata. thanks again

---

## Reviews

### Review by @Andrewgdewar - Commented on March 15, 2024 at 09:30 PM UTC

### Review by @rverdile - Commented on March 19, 2024 at 01:52 PM UTC

### Review by @rverdile - Commented on March 19, 2024 at 01:53 PM UTC

### Review by @Andrewgdewar - Commented on March 19, 2024 at 05:24 PM UTC

### Review by @jlsherrill - Commented on March 19, 2024 at 07:04 PM UTC

### Review by @rverdile - Commented on April 16, 2024 at 07:42 PM UTC

### Review by @rverdile - Commented on April 16, 2024 at 07:42 PM UTC

### Review by @rverdile - Approved on April 17, 2024 at 06:03 PM UTC

looks good! remember to do the release after you merge!

---

*Archived from: https://github.com/content-services/tang/pull/6*
