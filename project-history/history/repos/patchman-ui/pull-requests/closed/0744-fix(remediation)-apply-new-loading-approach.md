---
type: pull_request
number: 744
title: "fix(Remediation): apply new loading approach"
state: closed
author: mkholjuraev
created: 2022-03-09T22:39:02Z
updated: 2022-03-10T09:40:02Z
closed: 2022-03-10T09:40:02Z
base_branch: master
head_branch: fix-remediation
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/744
---

# Pull Request #744: fix(Remediation): apply new loading approach

**Author**: @mkholjuraev
**Created**: March 09, 2022 at 10:39 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `fix-remediation`

## Description

This is intended to remediate the issue on prod by adopting the latest technique of loading remediation. 

The issue on prod was due to the depricated `insights.experimental.loadRemediations
insights.loadRemediations` functions on Patch. Now the modal is loaded using AsyncComponent.

TODO: I need to find a way to open the modal from systems table row actions. 

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/744*
