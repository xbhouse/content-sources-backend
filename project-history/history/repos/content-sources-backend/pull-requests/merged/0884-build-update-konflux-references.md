---
type: pull_request
number: 884
title: "Build: Update Konflux references"
state: merged
author: red-hat-konflux
created: 2024-11-11T18:29:20Z
updated: 2024-11-18T19:23:29Z
closed: 2024-11-18T19:21:34Z
merged: 2024-11-18T19:21:34Z
base_branch: main
head_branch: konflux/references/main
labels: []
url: https://github.com/content-services/content-sources-backend/pull/884
---

# Pull Request #884: Build: Update Konflux references

**Author**: @red-hat-konflux
**Created**: November 11, 2024 at 06:29 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `konflux/references/main`

## Description

This PR contains the following updates:

| Package | Change |
|---|---|
| quay.io/konflux-ci/tekton-catalog/task-buildah | `27357fc` -> `e950993` |
| quay.io/konflux-ci/tekton-catalog/task-clair-scan | `90e371f` -> `eff773a` |
| quay.io/konflux-ci/tekton-catalog/task-clamav-scan | `1981b5a` -> `1465898` |
| quay.io/konflux-ci/tekton-catalog/task-deprecated-image-check | `443ffa8` -> `5a1a165` |
| quay.io/konflux-ci/tekton-catalog/task-ecosystem-cert-preflight-checks | `5131cce` -> `2e37ec3` |
| quay.io/konflux-ci/tekton-catalog/task-git-clone | `2cccdf8` -> `d091a9e` |
| quay.io/konflux-ci/tekton-catalog/task-init | `f239f38` -> `0523b51` |
| quay.io/konflux-ci/tekton-catalog/task-prefetch-dependencies | `f53fe54` -> `f59a214` |
| quay.io/konflux-ci/tekton-catalog/task-sast-snyk-check | `8aab0fd` -> `f165b1c` |
| quay.io/konflux-ci/tekton-catalog/task-summary | `d97c04a` -> `b0f049f` |

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiIzOC41NS4yLXJwbSIsInVwZGF0ZWRJblZlciI6IjM4LjU1LjItcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFpbiIsImxhYmVscyI6W119-->


---

## Discussion

### Comment by @app-sre-bot on November 11, 2024 at 06:31 PM UTC

Can one of the admins verify this patch?

---

## Reviews

### Review by @jlsherrill - Approved on November 18, 2024 at 07:21 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/884*
