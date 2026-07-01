---
type: pull_request
number: 406
title: "Fixes 2668: use newer yummy"
state: merged
author: jlsherrill
created: 2023-09-26T16:56:51Z
updated: 2023-09-26T19:00:27Z
closed: 2023-09-26T18:38:10Z
merged: 2023-09-26T18:38:10Z
base_branch: main
head_branch: 2668
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/406
---

# Pull Request #406: Fixes 2668: use newer yummy

**Author**: @jlsherrill
**Created**: September 26, 2023 at 04:56 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `2668`

## Description

## Summary
to limit the size of an xml file we can process

## Testing steps

Try to introspect a repo: https://raw.githubusercontent.com/lpardoRH/custom-repo/zip_bomb2/
You will likely run out of memory in ephemeral (maybe locally).  With this fix it will simply return zero packages (as its not really an xml file).   Right now this does not trigger an error which is a result of underlying libraries.  Since this edge case is an attempted attack and not a legitimate case, this is okay.  A partial (or corrupted) XML file would error as expected

---

## Discussion

### Comment by @jlsherrill on September 26, 2023 at 05:00 PM UTC

https://issues.redhat.com/browse/HMS-2668

### Comment by @swadeley on September 26, 2023 at 06:38 PM UTC

lgtm

---

## Reviews

### Review by @rverdile - Approved on September 26, 2023 at 05:33 PM UTC

tested and working!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/406*
