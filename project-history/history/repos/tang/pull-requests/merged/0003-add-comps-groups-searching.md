---
type: pull_request
number: 3
title: "add comps groups searching"
state: merged
author: rverdile
created: 2024-01-15T20:47:01Z
updated: 2024-01-22T14:45:45Z
closed: 2024-01-22T14:45:42Z
merged: 2024-01-22T14:45:41Z
base_branch: main
head_branch: comps
labels: []
url: https://github.com/content-services/tang/pull/3
---

# Pull Request #3: add comps groups searching

**Author**: @rverdile
**Created**: January 15, 2024 at 08:47 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `comps`

## Description

Adds two new methods to the Tangy interface

- `RpmRepositoryVersionPackageGroupSearch(ctx context.Context, hrefs []string, search string, limit int) ([]RpmPackageGroupSearch, error)`
  - Returns list of package groups (with ID, Name, Description, and Package list) in repository version, sorted by name. Package lists for overlapping ID+Name will be combined and de-duplicated.
- `RpmRepositoryVersionEnvironmentSearch(ctx context.Context, hrefs []string, search string, limit int) ([]RpmEnvironmentSearch, error)`
  - Returns list of environments (with ID, Name, Description) in repository version, sorted by name, and de-duplicated by ID+Name.

---

## Discussion

### Comment by @xbhouse on January 17, 2024 at 05:13 PM UTC

this looks good to me and works as expected! maybe @jlsherrill should have a look over too in case i missed something? 

---

## Reviews

### Review by @jlsherrill - Approved on January 18, 2024 at 04:14 PM UTC

LGTM

---

*Archived from: https://github.com/content-services/tang/pull/3*
