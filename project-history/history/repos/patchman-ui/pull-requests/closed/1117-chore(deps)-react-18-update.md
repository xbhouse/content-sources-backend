---
type: pull_request
number: 1117
title: "chore(deps): react 18 update"
state: closed
author: Fewwy
created: 2023-09-01T14:25:21Z
updated: 2024-03-05T11:25:30Z
closed: 2024-03-05T11:25:30Z
base_branch: master
head_branch: react18update
labels: ["dependencies"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1117
---

# Pull Request #1117: chore(deps): react 18 update

**Author**: @Fewwy
**Created**: September 01, 2023 at 02:25 PM UTC
**Status**: Closed
**Labels**: `dependencies`
**Base**: `master` ← **Head**: `react18update`

## Description

# Description
https://issues.redhat.com/browse/RHIF-133
PR for updating react 18 and testing what is broken by updating react version

Replaced enzyme snapshots with RTL snapshots. In some cases, had to use the react-shallow-renderer.
1) AdvisoryHeader.test.js - updated 
2) AdvisoryType.test.js - updated 
3) Header.test.js - 1 test should be refactored or dropped
4) InfoBox.test.js - updated 
5) EmptyStates.test.js - updated 
6) Error.test.js - updated 
7) Label.test.js - updated 
8) NoRegisteredSystems.test.js - updated 
9) SystemUpToDate.test.js - updated 
10) TableFooter.test.js - updated 
11) TableView.test.js - replaced with react-shallow-renderer
12) AdvisoryDetail.test.js - 1 test should be refactored or dropped
13) CveModal.test.js - Updated
14) AdvisorySystem.test.js - Updated
15) UnassignSystemModal.test.js - Updated
16) PackageDetail.test.js - Updated
17) PackageSystems.test.js - snapshot looks into props - I think it has to be refactored.
18) PatchSet.test.js - updated 
19) SelectExistingSets.test.js snapshot looks into selectOptions - I think it has to be refactored.
20) ReviewPatchSet.test.js - Updated 
21) SystemPackages.test.js - Updated
22) System.test.js - Updated
23) DataMappers.test.js - snapshot looks into packages - I think it has to be refactored.
24) Helpers.test.js - Updated
25) Packages.test.js - updated
26) SystemDetail.test.js - updated

---

## Reviews

### Review by @gkarat - Commented on September 08, 2023 at 11:47 AM UTC

### Review by @Fewwy - Commented on September 11, 2023 at 01:00 PM UTC

### Review by @Fewwy - Commented on September 11, 2023 at 01:00 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1117*
