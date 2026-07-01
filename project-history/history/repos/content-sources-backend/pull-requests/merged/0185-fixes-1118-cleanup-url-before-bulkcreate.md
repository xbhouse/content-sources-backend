---
type: pull_request
number: 185
title: "Fixes 1118: cleanup url before bulkcreate"
state: merged
author: rverdile
created: 2023-01-26T14:29:35Z
updated: 2023-01-27T01:05:23Z
closed: 2023-01-27T01:05:23Z
merged: 2023-01-27T01:05:23Z
base_branch: main
head_branch: url-claimed
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/185
---

# Pull Request #185: Fixes 1118: cleanup url before bulkcreate

**Author**: @rverdile
**Created**: January 26, 2023 at 02:29 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `url-claimed`

## Description

## Summary
During BulkCreate, the repository URL was not being cleaned up (trailing slashes add/removed) before searching for an existing repository. This would cause an error when creating a repository without a trailing slash, if the repository already existed (with a trailing slash) because `FirstOrCreate` would not find the repository.

## Testing steps
1. Create a repository https://jlsherrill.fedorapeople.org/fake-repos/needed-errata/
2. delete the repository
3. Create the repository again, but without a trailing slash: https://jlsherrill.fedorapeople.org/fake-repos/needed-errata

Without changes, this would say that it cannot create repository that already exists.


---

## Discussion

### Comment by @jlsherrill on January 26, 2023 at 02:30 PM UTC

https://issues.redhat.com/browse/HMS-1118

### Comment by @mshriver on January 26, 2023 at 11:30 PM UTC

/retest

### Comment by @mshriver on January 26, 2023 at 11:45 PM UTC

/retest


### Comment by @mshriver on January 26, 2023 at 11:46 PM UTC

I've got test coverage for this, and am retesting because there may just be a timeout issue. Addressing that timeout issue directly within the test, but may see green here earlier.

### Comment by @mshriver on January 27, 2023 at 12:57 AM UTC

/retest

### Comment by @mshriver on January 27, 2023 at 01:05 AM UTC

coverage included in E2E case for backend and frontend contexts.

---

## Reviews

### Review by @jlsherrill - Approved on January 26, 2023 at 05:58 PM UTC

### Review by @mshriver - Approved on January 27, 2023 at 01:04 AM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/185*
