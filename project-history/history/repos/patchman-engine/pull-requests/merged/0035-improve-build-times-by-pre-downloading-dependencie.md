---
type: pull_request
number: 35
title: "Improve build times by pre-downloading dependencies in module mode"
state: merged
author: semtexzv
created: 2019-12-12T08:23:59Z
updated: 2021-03-16T09:26:42Z
closed: 2019-12-13T13:44:29Z
merged: 2019-12-13T13:44:29Z
base_branch: master
head_branch: dockerfile-fixes
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/35
---

# Pull Request #35: Improve build times by pre-downloading dependencies in module mode

**Author**: @semtexzv
**Created**: December 12, 2019 at 08:23 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `dockerfile-fixes`

## Description

The `go vendor` created directory was not used by the build, forcing double downloading of dependencies.

---

## Reviews

### Review by @beav - Approved on December 12, 2019 at 09:51 PM UTC

change looks good and worked for me

### Review by @josef-hak - Approved on December 13, 2019 at 01:44 PM UTC

Cool, good improvement.

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/35*
