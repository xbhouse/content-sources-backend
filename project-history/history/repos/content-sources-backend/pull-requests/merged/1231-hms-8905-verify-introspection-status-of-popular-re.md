---
type: pull_request
number: 1231
title: "HMS-8905: Verify introspection status of popular repo after adding it in custom "
state: merged
author: mayurilahane
created: 2025-10-08T20:29:24Z
updated: 2025-10-15T15:14:56Z
closed: 2025-10-15T15:14:56Z
merged: 2025-10-15T15:14:56Z
base_branch: main
head_branch: mlahane/HMS-8905
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/1231
---

# Pull Request #1231: HMS-8905: Verify introspection status of popular repo after adding it in custom 

**Author**: @mayurilahane
**Created**: October 08, 2025 at 08:29 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `mlahane/HMS-8905`

## Description

## Summary

Goal was to migrate this test https://gitlab.cee.redhat.com/insights-qe/iqe-content-sources-plugin/-/blob/master/iqe_content_sources/tests/test_repository_api_only.py?ref_type=heads#L1019

but its is already implement in here https://github.com/content-services/content-sources-backend/blob/main/_playwright-tests/tests/API/Repositories.spec.ts#L253

just needed to varify if the introspection status is "valid" and this repo does not have failed introspection in past run



---

## Discussion

### Comment by @mayurilahane on October 08, 2025 at 08:31 PM UTC

Do not merge it until helper function PR is merged 
#1224 
https://github.com/content-services/content-sources-backend/pull/1224/files#diff-3012519d890c860c8cabc015676569b06036f131a0e8a3d02c15475da11aee2dR34


### Comment by @xbhouse on October 09, 2025 at 07:30 PM UTC

https://issues.redhat.com/browse/HMS-8905

---

## Reviews

### Review by @TenSt - Approved on October 15, 2025 at 08:58 AM UTC

lgtm

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1231*
