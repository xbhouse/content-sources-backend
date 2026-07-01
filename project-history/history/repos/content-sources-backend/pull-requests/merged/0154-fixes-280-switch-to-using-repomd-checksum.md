---
type: pull_request
number: 154
title: "Fixes 280: switch to using repomd checksum"
state: merged
author: jlsherrill
created: 2022-12-12T18:30:15Z
updated: 2023-01-03T16:30:36Z
closed: 2023-01-03T16:29:00Z
merged: 2023-01-03T16:29:00Z
base_branch: main
head_branch: 280
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/154
---

# Pull Request #154: Fixes 280: switch to using repomd checksum

**Author**: @jlsherrill
**Created**: December 12, 2022 at 06:30 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `280`

## Description

to determine if a repo was updated.

---

## Discussion

### Comment by @jlsherrill on December 12, 2022 at 06:30 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-280

### Comment by @jlsherrill on December 12, 2022 at 10:05 PM UTC

To test this, you can use this repository url: https://jlsherrill.fedorapeople.org/fake-repos/revision/current/

You should see one package.  Ask me to 'switch' it and add a 2nd package.  Then introspect it again, or wait ~24 hours.  you should now see 2 packages

### Comment by @mshriver on January 03, 2023 at 03:56 PM UTC

/retest


---

## Reviews

### Review by @Andrewgdewar - Approved on December 13, 2022 at 07:05 PM UTC

Tested for functionality. Works as intended. 
Code looks good but will defer to @rverdile's opinion in that regard.
Ack

### Review by @rverdile - Commented on December 13, 2022 at 07:09 PM UTC

### Review by @jlsherrill - Commented on December 13, 2022 at 10:23 PM UTC

### Review by @rverdile - Commented on December 14, 2022 at 03:07 PM UTC

### Review by @mshriver - Approved on January 03, 2023 at 04:11 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/154*
