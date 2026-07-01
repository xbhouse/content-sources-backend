---
type: pull_request
number: 858
title: "feat(RHIF-83: introduce loading icon on remediation button in Advisor\u2026"
state: merged
author: mkholjuraev
created: 2022-08-01T15:54:45Z
updated: 2024-04-03T09:23:10Z
closed: 2022-08-02T07:54:59Z
merged: 2022-08-02T07:54:59Z
base_branch: master
head_branch: loader-remediation
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/858
---

# Pull Request #858: feat(RHIF-83: introduce loading icon on remediation button in Advisor…

**Author**: @mkholjuraev
**Created**: August 01, 2022 at 03:54 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `loader-remediation`

## Description

Partly resloves: https://issues.redhat.com/browse/RHIF-51.

It just introduces a loading icon while the remediation issues map is being built. There is a new ticket in the advisor to optimize the API requests to fetch the issues map. 

---

## Discussion

### Comment by @mkholjuraev on August 02, 2022 at 07:54 AM UTC

@adonispuente thank you for the review! YOU are awesome!

### Comment by @mkholjuraev on August 02, 2022 at 08:13 AM UTC

@adonispuente in a good world we should not see that modal at all on Advisories page. That modal is poping up because of API request failure. I have asked the backend to fix that endpoint.

### Comment by @jiridostal on August 24, 2022 at 09:47 PM UTC

:tada: This PR is included in version 1.49.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.49.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.49.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @adonispuente - Approved on August 01, 2022 at 09:29 PM UTC

Thank you, PR works as intended, nice organization. LGTM!

Just a side note, has it been considered to disable the remediate button if the selected items are remediable automatically by ansible? I know theres a popup that displays afterwards if all remediations are manual, but it takes a bit to get feedback from it. There could be a reason we dont do this, but I just wanted to through it out there and hear your thoughts!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/858*
