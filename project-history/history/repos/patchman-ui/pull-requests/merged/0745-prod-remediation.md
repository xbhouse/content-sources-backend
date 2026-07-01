---
type: pull_request
number: 745
title: "Prod remediation"
state: merged
author: mkholjuraev
created: 2022-03-10T08:55:38Z
updated: 2024-04-03T09:21:41Z
closed: 2022-03-10T10:21:24Z
merged: 2022-03-10T10:21:23Z
base_branch: prod-beta
head_branch: prod-remediation
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/745
---

# Pull Request #745: Prod remediation

**Author**: @mkholjuraev
**Created**: March 10, 2022 at 08:55 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `prod-beta` ← **Head**: `prod-remediation`

## Description

This is intended to remediate the issue on prod by adopting the latest technique of loading remediation.

The issue on prod was due to the depricated insights.experimental.loadRemediations insights.loadRemediations functions on Patch. Now the modal is loaded using AsyncComponent.

---

## Discussion

### Comment by @mkholjuraev on March 10, 2022 at 10:20 AM UTC

I would appreciate a quick review as it needs to land in prod ASAP. Some tests are disabled, I focused on resolving the bug, I am preparing a clean PR for later release.

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/745*
