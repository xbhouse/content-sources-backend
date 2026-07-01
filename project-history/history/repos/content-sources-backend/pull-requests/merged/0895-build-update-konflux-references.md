---
type: pull_request
number: 895
title: "Build: Update Konflux references"
state: merged
author: red-hat-konflux
created: 2024-11-18T20:38:43Z
updated: 2024-12-02T21:05:53Z
closed: 2024-12-02T21:05:50Z
merged: 2024-12-02T21:05:50Z
base_branch: main
head_branch: konflux/references/main
labels: []
url: https://github.com/content-services/content-sources-backend/pull/895
---

# Pull Request #895: Build: Update Konflux references

**Author**: @red-hat-konflux
**Created**: November 18, 2024 at 08:38 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `konflux/references/main`

## Description

This PR contains the following updates:

| Package | Change |
|---|---|
| quay.io/konflux-ci/tekton-catalog/task-build-image-index | `5da8c2f` -> `7b2c5ab` |
| quay.io/konflux-ci/tekton-catalog/task-buildah | `e950993` -> `7d3f090` |
| quay.io/konflux-ci/tekton-catalog/task-clair-scan | `eff773a` -> `0a54211` |
| quay.io/konflux-ci/tekton-catalog/task-clamav-scan | `1465898` -> `b4f450f` |
| quay.io/konflux-ci/tekton-catalog/task-ecosystem-cert-preflight-checks | `2e37ec3` -> `df8a25a` |
| quay.io/konflux-ci/tekton-catalog/task-prefetch-dependencies | `f59a214` -> `d3d8a8b` |
| quay.io/konflux-ci/tekton-catalog/task-push-dockerfile | `a216178` -> `48bb2ee` |
| quay.io/konflux-ci/tekton-catalog/task-rpms-signature-scan | `0c9667f` -> `8f3b23b` |
| quay.io/konflux-ci/tekton-catalog/task-sast-snyk-check | `f165b1c` -> `9fa8acb` |
| quay.io/konflux-ci/tekton-catalog/task-show-sbom | `52f8b96` -> `945a7c9` |
| quay.io/konflux-ci/tekton-catalog/task-source-build | `53a41b0` -> `957f765` |
| quay.io/konflux-ci/tekton-catalog/task-summary | `b0f049f` -> `870d9a0` |

---

### Configuration

📅 **Schedule**: Branch creation - "after 5am on saturday" (UTC), Automerge - At any time (no schedule defined).

🚦 **Automerge**: Disabled by config. Please merge this manually once you are satisfied.

♻ **Rebasing**: Whenever PR is behind base branch, or you tick the rebase/retry checkbox.

👻 **Immortal**: This PR will be recreated if closed unmerged. Get [config help](https://togithub.com/renovatebot/renovate/discussions) if that's undesired.

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiIzOC41NS4yLXJwbSIsInVwZGF0ZWRJblZlciI6IjM4LjU1LjItcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFpbiIsImxhYmVscyI6W119-->


---

## Discussion

### Comment by @jlsherrill on November 25, 2024 at 08:23 PM UTC

/retest

### Comment by @jlsherrill on December 02, 2024 at 05:51 PM UTC

/retest

---

## Reviews

### Review by @jlsherrill - Approved on December 02, 2024 at 09:05 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/895*
