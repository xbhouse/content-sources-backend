---
type: pull_request
number: 1438
title: "Standardize konflux-pipelines to main branch"
state: merged
author: ryemorris
created: 2025-11-18T16:45:09Z
updated: 2025-11-19T08:56:48Z
closed: 2025-11-19T08:56:48Z
merged: 2025-11-19T08:56:48Z
base_branch: security-compliance
head_branch: update-pipeline-to-main
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1438
---

# Pull Request #1438: Standardize konflux-pipelines to main branch

**Author**: @ryemorris
**Created**: November 18, 2025 at 04:45 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `security-compliance` ← **Head**: `update-pipeline-to-main`

## Description

## Summary
- Removed version pinning from konflux-pipelines URL in SC YAML files
- Updated to `main` branch reference

## Changes
Updated all `*-sc-*.yaml` files in `.tekton/` directory

## Reason
Standardizing pipeline references across the platform to use the main branch instead of version-specific URLs.

---

## Discussion

### Comment by @ryemorris on November 18, 2025 at 06:36 PM UTC

/retest

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1438*
