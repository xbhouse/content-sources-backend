---
type: pull_request
number: 310
title: "Refs 1947: parametrize CJI fields for IQE"
state: merged
author: mshriver
created: 2023-06-15T13:06:39Z
updated: 2023-06-15T13:56:18Z
closed: 2023-06-15T13:56:18Z
merged: 2023-06-15T13:56:18Z
base_branch: main
head_branch: parametrize-CJI
labels: []
url: https://github.com/content-services/content-sources-backend/pull/310
---

# Pull Request #310: Refs 1947: parametrize CJI fields for IQE

**Author**: @mshriver
**Created**: June 15, 2023 at 01:06 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `parametrize-CJI`

## Description

## Summary
Allow for running the CJI with parameters set via SAAS


## Testing steps
`bonfire deploy-iqe-cji <args> --template-file <path-to-file-in-branch>`

Resulting CJI had parameters values set correctly via defaults. Resulting pod had correct env config and pytest args.

---

## Discussion

### Comment by @jlsherrill on June 15, 2023 at 01:30 PM UTC

https://issues.redhat.com/browse/HMS-1947

---

## Reviews

### Review by @jlsherrill - Approved on June 15, 2023 at 01:21 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/310*
