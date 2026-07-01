---
type: pull_request
number: 873
title: "Build: Update Konflux references"
state: merged
author: red-hat-konflux
created: 2024-11-05T20:07:23Z
updated: 2024-11-05T21:03:20Z
closed: 2024-11-05T21:03:16Z
merged: 2024-11-05T21:03:16Z
base_branch: main
head_branch: konflux/references/main
labels: []
url: https://github.com/content-services/content-sources-backend/pull/873
---

# Pull Request #873: Build: Update Konflux references

**Author**: @red-hat-konflux
**Created**: November 05, 2024 at 08:07 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `konflux/references/main`

## Description

This PR contains the following updates:

| Package | Change | Notes |
|---|---|---|
| quay.io/konflux-ci/tekton-catalog/task-apply-tags | `f485e25` -> `87fd7fc` |  |
| quay.io/konflux-ci/tekton-catalog/task-build-image-index | `327d745` -> `715fa1f` |  |
| quay.io/konflux-ci/tekton-catalog/task-buildah | `67f0290` -> `fedcfe0` |  |
| quay.io/konflux-ci/tekton-catalog/task-clair-scan | `28fee4b` -> `90e371f` |  |
| quay.io/konflux-ci/tekton-catalog/task-clamav-scan | `a94b652` -> `1981b5a` |  |
| quay.io/konflux-ci/tekton-catalog/task-deprecated-image-check | `b4f9599` -> `443ffa8` |  |
| quay.io/konflux-ci/tekton-catalog/task-init | `092c113` -> `f239f38` |  |
| quay.io/konflux-ci/tekton-catalog/task-prefetch-dependencies | `fe7234e` -> `f53fe54` |  |
| quay.io/konflux-ci/tekton-catalog/task-push-dockerfile | `674e70f` -> `a216178` |  |
| quay.io/konflux-ci/tekton-catalog/task-rpms-signature-scan | `7aa4d3c` -> `0c9667f` |  |
| quay.io/konflux-ci/tekton-catalog/task-sast-snyk-check | `0.2` -> `0.3` | :warning:[migration](https://togithub.com/redhat-appstudio/build-definitions/blob/main/task/sast-snyk-check/0.3/MIGRATION.md):warning: |

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiIzOC41NS4yLXJwbSIsInVwZGF0ZWRJblZlciI6IjM4LjU1LjItcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFpbiIsImxhYmVscyI6W119-->


---

## Discussion

### Comment by @app-sre-bot on November 05, 2024 at 08:09 PM UTC

Can one of the admins verify this patch?

### Comment by @ezr-ondrej on November 05, 2024 at 08:42 PM UTC

/retest

### Comment by @ezr-ondrej on November 05, 2024 at 08:42 PM UTC

/ok-to-test

---

## Reviews

### Review by @jlsherrill - Approved on November 05, 2024 at 09:03 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/873*
