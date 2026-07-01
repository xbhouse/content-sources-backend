---
type: pull_request
number: 741
title: "fix(PatchSetWizard): separate request to update and create set"
state: merged
author: mkholjuraev
created: 2022-03-04T19:58:12Z
updated: 2024-04-03T09:21:39Z
closed: 2022-03-18T13:55:26Z
merged: 2022-03-18T13:55:26Z
base_branch: master
head_branch: fix-patch-set
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/741
---

# Pull Request #741: fix(PatchSetWizard): separate request to update and create set

**Author**: @mkholjuraev
**Created**: March 04, 2022 at 07:58 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `fix-patch-set`

## Description

We need to send  distinct requests for updating and creating a patch set according to the changed [API docs](https://console.stage.redhat.com/beta/docs/api/patch#operations-default-createBaseline). The API doc is missing description field on both endpoints:

1. For creating a patch set (**/api/patch/v1/baselines**): 
  `{
    "config": {
      "to_time": "2022-12-31T12:00:00-04:00"
    },
    "inventory_ids": [
      "string"
    ],
    "name": "string"
  }` 
2. For updating a patch set (**/api/patch/v1/baselines/{baseline_id}**):
```
  {
    "config": {
      "to_time": "2022-12-31T12:00:00-04:00"
    },
    "inventory_ids": {
      "additionalProp1": true,
      "additionalProp2": true,
      "additionalProp3": true
    },
    "name": "my-changed-baseline-name"
  }
```

---

## Discussion

### Comment by @mkholjuraev on March 18, 2022 at 01:55 PM UTC

I would like include this PR for my current patch set wizard integration work. Merging this as it will not be pushed to prod untill Q2 and can be tested in the next PRs

### Comment by @jiridostal on March 18, 2022 at 02:07 PM UTC

:tada: This PR is included in version 1.40.4 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.40.4)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.40.4)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/741*
