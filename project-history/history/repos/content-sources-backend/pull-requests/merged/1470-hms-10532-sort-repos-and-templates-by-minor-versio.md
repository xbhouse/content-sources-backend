---
type: pull_request
number: 1470
title: "HMS-10532: sort repos and templates by minor version"
state: merged
author: katarinazaprazna
created: 2026-04-19T22:06:13Z
updated: 2026-04-21T15:21:29Z
closed: 2026-04-21T15:07:39Z
merged: 2026-04-21T15:07:39Z
base_branch: main
head_branch: support-sorting-by-minor-version
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1470
---

# Pull Request #1470: HMS-10532: sort repos and templates by minor version

**Author**: @katarinazaprazna
**Created**: April 19, 2026 at 10:06 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `support-sorting-by-minor-version`

## Description

## Summary

- Enhanced version sorting for repos and templates to handle minor and major versions. The order follows RHEL versioning: 8, 8.0, 8.1, ..., 9, 9.0, 9.1, ...
- Uses `extended_release_version` for minor version repos/templates, with this formula `major * 10000 + minor (or major * 10000 - 1 for major-only)` to compute an integer sort key

<img width="1303" height="1088" alt="Screenshot From 2026-04-20 00-19-48" src="https://github.com/user-attachments/assets/a97ea38f-090c-49aa-b988-155e5ea4d7f5" />


## Testing steps

- Existing tests pass
- Repositories sorted by version return "", "9.4", "9.4", "9.6" in ascending order
- Templates sorted by version return "8.6", "", "9.4" in ascending order

---

## Discussion

### Comment by @katarinazaprazna on April 19, 2026 at 10:16 PM UTC

Hi @rverdile and @xbhouse, I’m concerned about the performance of my sorting logic for repos, as the sort expression appears to run for every row on every request. This overhead could become a bottleneck in larger repo datasets. I’d like to get your thoughts on whether we should optimize this. Thanks!

### Comment by @xbhouse on April 19, 2026 at 10:30 PM UTC

https://issues.redhat.com/browse/HMS-10532

### Comment by @katarinazaprazna on April 20, 2026 at 10:22 AM UTC

/retest

### Comment by @xbhouse on April 20, 2026 at 03:07 PM UTC

@katarinazaprazna i took a look at this and ran `EXPLAIN ANALYZE` comparing the new sort vs the original.

for repositories, with the full set of RH and EPEL repos (65), the new sort actually looks faster :) 
- ~0.317ms (new)
- ~0.505ms (old)

i think since your new query uses `extended_release_version` to sort when it's present, the number of times we have to account for multiple versions in the version array is reduced :star_struck: 

for templates, i checked the current max number of templates an org has in prod, which is 37. with a set of 50 templates, the new sort is _slightly_ slower, which makes sense:
- ~0.202ms (new)
- ~0.135ms (old)

based on this i'd say we don't need to optimize right now, but we can revisit if these numbers grow or if we start seeing an impact on performance. @rverdile @jlsherrill - curious what you think?

### Comment by @rverdile on April 21, 2026 at 02:44 PM UTC

@xbhouse i agree! 

---

## Reviews

### Review by @xbhouse - Approved on April 21, 2026 at 03:02 PM UTC

looks great to me! thank you! 🫶 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1470*
