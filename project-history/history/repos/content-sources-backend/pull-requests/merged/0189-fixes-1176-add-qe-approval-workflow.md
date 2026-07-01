---
type: pull_request
number: 189
title: "Fixes 1176: add qe approval workflow"
state: merged
author: jlsherrill
created: 2023-02-01T22:52:36Z
updated: 2023-02-01T23:07:08Z
closed: 2023-02-01T23:07:08Z
merged: 2023-02-01T23:07:08Z
base_branch: main
head_branch: workflow
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/189
---

# Pull Request #189: Fixes 1176: add qe approval workflow

**Author**: @jlsherrill
**Created**: February 01, 2023 at 10:52 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `workflow`

## Description

## Summary

This adds a couple of actions:
1) a pr is red if it has 'qe-testing-needed' label but does not have a qe-approved label
2) if someone from our quality engineering team approves a pr, it gets a 'qe-approved' label.  This uses a github action i forked and modified here:  https://github.com/content-services/label-when-approved-action

## Testing steps


---

## Discussion

### Comment by @jlsherrill on February 01, 2023 at 11:00 PM UTC

https://issues.redhat.com/browse/HMS-1176

---

## Reviews

### Review by @mshriver - Approved on February 01, 2023 at 11:07 PM UTC

Thanks for adding this check, great visual indicator as the labels themselves don't stand out enough to highlight this situation. 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/189*
