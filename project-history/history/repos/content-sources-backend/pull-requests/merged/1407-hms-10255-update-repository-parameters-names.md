---
type: pull_request
number: 1407
title: "HMS-10255: update repository parameters names"
state: merged
author: Dugowitch
created: 2026-03-06T18:12:26Z
updated: 2026-03-09T12:55:52Z
closed: 2026-03-09T12:55:49Z
merged: 2026-03-09T12:55:49Z
base_branch: main
head_branch: hms10255
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1407
---

# Pull Request #1407: HMS-10255: update repository parameters names

**Author**: @Dugowitch
**Created**: March 06, 2026 at 06:12 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `hms10255`

## Description

## Summary
List of versions uses names  "RHEL 8", "RHEL 9", and "RHEL 10" instead of "el8, el9, el10". These are displayed in the UI.

## Testing steps
tests pass

---

## Discussion

### Comment by @xbhouse on March 06, 2026 at 06:30 PM UTC

https://issues.redhat.com/browse/HMS-10255

### Comment by @xbhouse on March 06, 2026 at 09:08 PM UTC

not sure if this will work, but you might try adding `#testwith https://github.com/content-services/content-sources-frontend/pull/893` to this PR description

### Comment by @Dugowitch on March 09, 2026 at 10:05 AM UTC

/retest

### Comment by @Dugowitch on March 09, 2026 at 11:06 AM UTC

It logged:
```
No PR title or description contains '#testwith https://github.com/content-services/content-sources-backend/pull/1407'. Cloning the main branch.
```

So it's just not picking it up

### Comment by @katarinazaprazna on March 09, 2026 at 12:29 PM UTC

Hey @xbhouse and @Dugowitch. Apologies for the delay! I missed the notification for this one

Regarding the changes, the frontend currently expects the following structure:
```
export type NameLabel = {
  name: string;
  label: string;
};
```

At the moment, we are receiving `{ name: "el9", label: "el9" }`. Once the backend changes are live, we're expecting the name to reflect the display value i.e., `{ name: "RHEL 9", label: "el9" }`

We're expecting both `version` and `extended_release_version` to change, so that the UI displays OS version names consistently. Since the frontend pulls these values from the repository_parameters API, I don’t anticipate any disruptions (aside from needing to update our tests)

---

## Reviews

### Review by @xbhouse - Commented on March 06, 2026 at 09:06 PM UTC

### Review by @Dugowitch - Commented on March 09, 2026 at 10:50 AM UTC

### Review by @katarinazaprazna - Commented on March 09, 2026 at 12:34 PM UTC

### Review by @xbhouse - Approved on March 09, 2026 at 12:53 PM UTC

thank you! this looks good to me :)  tested with Kate's frontend [PR](https://github.com/content-services/content-sources-frontend/pull/849) too

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1407*
