---
type: pull_request
number: 16
title: "Fixes 2442: add support for parsing groups and envs"
state: merged
author: xbhouse
created: 2023-10-26T14:16:42Z
updated: 2023-10-31T13:13:51Z
closed: 2023-10-31T13:13:51Z
merged: 2023-10-31T13:13:51Z
base_branch: master
head_branch: pgroups_and_envs
labels: []
url: https://github.com/content-services/yummy/pull/16
---

# Pull Request #16: Fixes 2442: add support for parsing groups and envs

**Author**: @xbhouse
**Created**: October 26, 2023 at 02:16 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pgroups_and_envs`

## Description

added:
- structs for Comps, PackageGroup, and Environment, and specified types for names and descriptions to avoid localized elements
- function to get the comps URL 
- methods to return Comps, PackageGroups, and Environments for a Repository
- parser for comps.xml to grab each group and environment element
- custom unmarshal methods for PackageGroup and Environment names and descriptions to ignore the elements with a language attribute (thank you @rverdile for this!!) 
- mock comps.xml and tests

testing steps:
- follow usage outlined in [readme](https://github.com/content-services/yummy/blob/master/README.md#usage)
- to get PackageGroups: `packageGroups, statusCode, err := repo.PackageGroups()`
- to get Environments: `environments, statusCode, err := repo.Environments()`


---

## Discussion

### Comment by @jlsherrill on October 26, 2023 at 08:03 PM UTC

it looks like this may result in the comps.xml file being downloaded and parsed multiple times if we call r.PackageGroups() and then r.Environments() .  You might consider caching the parsed xml of the comps.xml on the repository object like is done for repomd.xml:  https://github.com/content-services/yummy/blob/master/pkg/yum/repository.go#L140-L141

EDIT: oh i see you are caching the package groups and environment on the repository, but doing so individually.  You might think about making a Comps() method on repository that calls your ParseCompsXML() standalone function, and stores the results on the repository.  Then your "PackageGroups()" method could just call r.Comps()  and return the cached package groups.  (Yes i'm terrible at naming things)

Overall this looks great!

### Comment by @xbhouse on October 27, 2023 at 01:04 PM UTC

oo ok! i'll add that in

### Comment by @jlsherrill on October 27, 2023 at 08:30 PM UTC

Tiny nitpick, other things look fine.  @rverdile do you wanna take a look over?

---

## Reviews

### Review by @jlsherrill - Commented on October 27, 2023 at 08:25 PM UTC

### Review by @xbhouse - Commented on October 27, 2023 at 08:33 PM UTC

### Review by @rverdile - Commented on October 30, 2023 at 02:40 PM UTC

This looks good! One more thing to do would be update the README to include the new features.

### Review by @jlsherrill - Approved on October 31, 2023 at 01:13 PM UTC

Awesome work @xbhouse !

---

*Archived from: https://github.com/content-services/yummy/pull/16*
