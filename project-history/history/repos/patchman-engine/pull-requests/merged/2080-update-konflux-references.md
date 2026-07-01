---
type: pull_request
number: 2080
title: "Update Konflux references"
state: merged
author: red-hat-konflux
created: 2026-02-28T05:09:32Z
updated: 2026-02-28T05:16:20Z
closed: 2026-02-28T05:14:56Z
merged: 2026-02-28T05:14:56Z
base_branch: master
head_branch: konflux/references/master
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2080
---

# Pull Request #2080: Update Konflux references

**Author**: @red-hat-konflux
**Created**: February 28, 2026 at 05:09 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/references/master`

## Description

This PR contains the following updates:

| Package | Change | Notes |
|---|---|---|
| [quay.io/konflux-ci/tekton-catalog/task-build-image-index](https://redirect.github.com/konflux-ci/build-definitions/tree/95b9f85298d70df3f6ed8ea50db465b103fb994e/task/build-image-index/0.2) ([source](https://redirect.github.com/konflux-ci/build-definitions/tree/HEAD/task/build-image-index), [changelog](https://redirect.github.com/konflux-ci/build-definitions/blob/main/task/build-image-index/CHANGELOG.md)) | `30989fa` -> `ac4f8b5` |  |
| [quay.io/konflux-ci/tekton-catalog/task-buildah-oci-ta](https://redirect.github.com/konflux-ci/build-definitions/tree/dd5433145410ec8f2ee92d7803d09d7f223e6e0b/task/buildah-oci-ta/0.9) ([source](https://redirect.github.com/konflux-ci/build-definitions/tree/HEAD/task/buildah-oci-ta), [changelog](https://redirect.github.com/konflux-ci/build-definitions/blob/main/task/buildah-oci-ta/CHANGELOG.md)) | `0.8` -> `0.9` | :warning:[migration](https://redirect.github.com/konflux-ci/build-definitions/blob/main/task/buildah-oci-ta/0.9/MIGRATION.md):warning: |
| [quay.io/konflux-ci/tekton-catalog/task-clair-scan](https://redirect.github.com/konflux-ci/konflux-test-tasks/tree/0bf981e1e5ff7aaba40cf6ac3e69fec84605c49a/task/clair-scan/0.3) ([source](https://redirect.github.com/konflux-ci/konflux-test-tasks/tree/HEAD/task/clair-scan), [changelog](https://redirect.github.com/konflux-ci/konflux-test-tasks/blob/main/task/clair-scan/CHANGELOG.md)) | `b01d8e2` -> `1f19b5f` |  |
| [quay.io/konflux-ci/tekton-catalog/task-clamav-scan](https://redirect.github.com/konflux-ci/konflux-test-tasks/tree/0bf981e1e5ff7aaba40cf6ac3e69fec84605c49a/task/clamav-scan/0.3) ([source](https://redirect.github.com/konflux-ci/konflux-test-tasks/tree/HEAD/task/clamav-scan), [changelog](https://redirect.github.com/konflux-ci/konflux-test-tasks/blob/main/task/clamav-scan/CHANGELOG.md)) | `5b5b31e` -> `10054cd` |  |
| [quay.io/konflux-ci/tekton-catalog/task-coverity-availability-check](https://redirect.github.com/konflux-ci/konflux-sast-tasks/tree/55b0a7b67945bebc62c89060b824e46ba08fb52f/task/coverity-availability-check/0.2) ([source](https://redirect.github.com/konflux-ci/konflux-sast-tasks/tree/HEAD/task/coverity-availability-check), [changelog](https://redirect.github.com/konflux-ci/konflux-sast-tasks/blob/main/task/coverity-availability-check/CHANGELOG.md)) | `a24d8f3` -> `c9b9301` |  |
| [quay.io/konflux-ci/tekton-catalog/task-deprecated-image-check](https://redirect.github.com/konflux-ci/konflux-test-tasks/tree/0bf981e1e5ff7aaba40cf6ac3e69fec84605c49a/task/deprecated-image-check/0.5) ([source](https://redirect.github.com/konflux-ci/konflux-test-tasks/tree/HEAD/task/deprecated-image-check), [changelog](https://redirect.github.com/konflux-ci/konflux-test-tasks/blob/main/task/deprecated-image-check/CHANGELOG.md)) | `e3a55cc` -> `516ea66` |  |
| [quay.io/konflux-ci/tekton-catalog/task-ecosystem-cert-preflight-checks](https://redirect.github.com/konflux-ci/build-definitions/tree/d8df4b5f49ab0d7812c72ec3c1e271c5e5642d3e/task/ecosystem-cert-preflight-checks/0.2) ([source](https://redirect.github.com/konflux-ci/build-definitions/tree/HEAD/task/ecosystem-cert-preflight-checks), [changelog](https://redirect.github.com/konflux-ci/build-definitions/blob/main/task/ecosystem-cert-preflight-checks/CHANGELOG.md)) | `40bc4bc` -> `1c3bbac` |  |
| [quay.io/konflux-ci/tekton-catalog/task-init](https://redirect.github.com/konflux-ci/build-definitions/tree/04290fccfab17a95cccb7fc246f4cb3ec12ca00a/task/init/0.4) ([source](https://redirect.github.com/konflux-ci/build-definitions/tree/HEAD/task/init), [changelog](https://redirect.github.com/konflux-ci/build-definitions/blob/main/task/init/CHANGELOG.md)) | `0.3` -> `0.4` | :warning:[migration](https://redirect.github.com/konflux-ci/build-definitions/blob/main/task/init/0.4/MIGRATION.md):warning: |
| [quay.io/konflux-ci/tekton-catalog/task-prefetch-dependencies-oci-ta](https://redirect.github.com/konflux-ci/build-definitions/tree/e68de7a16f7c7242abcaa187eece03fe981e8ebc/task/prefetch-dependencies-oci-ta/0.2) ([source](https://redirect.github.com/konflux-ci/build-definitions/tree/HEAD/task/prefetch-dependencies-oci-ta), [changelog](https://redirect.github.com/konflux-ci/build-definitions/blob/main/task/prefetch-dependencies-oci-ta/CHANGELOG.md)) | `c07551e` -> `9a8b6e2` |  |
| [quay.io/konflux-ci/tekton-catalog/task-push-dockerfile-oci-ta](https://redirect.github.com/konflux-ci/build-definitions/tree/9ae3833bff4aa84ab5e9a27df152edd00a32dd7d/task/push-dockerfile-oci-ta/0.2) ([source](https://redirect.github.com/konflux-ci/build-definitions/tree/HEAD/task/push-dockerfile-oci-ta), [changelog](https://redirect.github.com/konflux-ci/build-definitions/blob/main/task/push-dockerfile-oci-ta/CHANGELOG.md)) | `0.1` -> `0.2` | :warning:[migration](https://redirect.github.com/konflux-ci/build-definitions/blob/main/task/push-dockerfile-oci-ta/0.2/MIGRATION.md):warning: |
| [quay.io/konflux-ci/tekton-catalog/task-rpms-signature-scan](https://redirect.github.com/konflux-ci/tekton-tools/tree/5cad89c6b0edbe5682b34553ca96b1081ece864b/tasks/rpms-signature-scan/0.2) ([source](https://redirect.github.com/konflux-ci/tekton-tools/tree/HEAD/task/rpms-signature-scan), [changelog](https://redirect.github.com/konflux-ci/tekton-tools/blob/main/task/rpms-signature-scan/CHANGELOG.md)) | `d762bda` -> `8252a4a` |  |
| [quay.io/konflux-ci/tekton-catalog/task-sast-shell-check-oci-ta](https://redirect.github.com/konflux-ci/konflux-sast-tasks/tree/55b0a7b67945bebc62c89060b824e46ba08fb52f/task/sast-shell-check-oci-ta/0.1,!task/sast-shell-check-oci-ta/0.1/recipe.yaml) ([source](https://redirect.github.com/konflux-ci/konflux-sast-tasks/tree/HEAD/task/sast-shell-check-oci-ta), [changelog](https://redirect.github.com/konflux-ci/konflux-sast-tasks/blob/main/task/sast-shell-check-oci-ta/CHANGELOG.md)) | `f475b4b` -> `a528f01` |  |
| [quay.io/konflux-ci/tekton-catalog/task-sast-snyk-check-oci-ta](https://redirect.github.com/konflux-ci/konflux-sast-tasks/tree/55b0a7b67945bebc62c89060b824e46ba08fb52f/task/sast-snyk-check-oci-ta/0.4,!task/sast-snyk-check-oci-ta/0.4/recipe.yaml) ([source](https://redirect.github.com/konflux-ci/konflux-sast-tasks/tree/HEAD/task/sast-snyk-check-oci-ta), [changelog](https://redirect.github.com/konflux-ci/konflux-sast-tasks/blob/main/task/sast-snyk-check-oci-ta/CHANGELOG.md)) | `0c2ab8c` -> `c4b79e0` |  |
| [quay.io/konflux-ci/tekton-catalog/task-sast-unicode-check-oci-ta](https://redirect.github.com/konflux-ci/konflux-sast-tasks/tree/55b0a7b67945bebc62c89060b824e46ba08fb52f/task/sast-unicode-check-oci-ta/0.4,!task/sast-unicode-check-oci-ta/0.4/recipe.yaml) ([source](https://redirect.github.com/konflux-ci/konflux-sast-tasks/tree/HEAD/task/sast-unicode-check-oci-ta), [changelog](https://redirect.github.com/konflux-ci/konflux-sast-tasks/blob/main/task/sast-unicode-check-oci-ta/CHANGELOG.md)) | `b38140b` -> `2bcccab` |  |
| [quay.io/konflux-ci/tekton-catalog/task-show-sbom](https://redirect.github.com/konflux-ci/build-definitions/tree/b6cb1d4475323de7cda1d3334911ed763fe118e4/task/show-sbom/0.1) ([source](https://redirect.github.com/konflux-ci/build-definitions/tree/HEAD/task/show-sbom), [changelog](https://redirect.github.com/konflux-ci/build-definitions/blob/main/task/show-sbom/CHANGELOG.md)) | `e2c1b4e` -> `e119aa8` |  |
| [quay.io/konflux-ci/tekton-catalog/task-source-build-oci-ta](https://redirect.github.com/konflux-ci/build-definitions/tree/95b9f85298d70df3f6ed8ea50db465b103fb994e/task/source-build-oci-ta/0.3) ([source](https://redirect.github.com/konflux-ci/build-definitions/tree/HEAD/task/source-build-oci-ta), [changelog](https://redirect.github.com/konflux-ci/build-definitions/blob/main/task/source-build-oci-ta/CHANGELOG.md)) | `c35ba21` -> `eb620d1` |  |

---

### Release Notes

<details>
<summary>konflux-ci/build-definitions (quay.io/konflux-ci/tekton-catalog/task-buildah-oci-ta)</summary>

### [`v0.9`](https://redirect.github.com/konflux-ci/build-definitions/blob/HEAD/task/buildah-oci-ta/CHANGELOG.md#09)

##### Removed

- BREAKING: Support for Dockerfile downloading in Konflux Build Pipeline.

</details>

<details>
<summary>konflux-ci/build-definitions (quay.io/konflux-ci/tekton-catalog/task-init)</summary>

### [`v0.4`](https://redirect.github.com/konflux-ci/build-definitions/blob/HEAD/task/init/CHANGELOG.md#04)

- Task started using konflux build cli instead of bash script.

</details>

<details>
<summary>konflux-ci/build-definitions (quay.io/konflux-ci/tekton-catalog/task-push-dockerfile-oci-ta)</summary>

### [`v0.2`](https://redirect.github.com/konflux-ci/build-definitions/blob/HEAD/task/push-dockerfile-oci-ta/CHANGELOG.md#02)

##### Removed

- BREAKING: Support for Dockerfile downloading in Konflux Build Pipeline.

</details>

---

### Configuration

📅 **Schedule**: Branch creation - Between 05:00 AM and 11:59 PM, only on Saturday ( * 5-23 * * 6 ) in timezone Europe/Prague, Automerge - At any time (no schedule defined).

🚦 **Automerge**: Enabled.

♻ **Rebasing**: Whenever PR is behind base branch, or you tick the rebase/retry checkbox.

👻 **Immortal**: This PR will be recreated if closed unmerged. Get [config help](https://redirect.github.com/renovatebot/renovate/discussions) if that's undesired.

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`.

---
### Documentation

Find out how to configure dependency updates in [MintMaker documentation](https://konflux-ci.dev/docs/mintmaker/user/) or see all available configuration options in [Renovate documentation](https://docs.renovatebot.com/configuration-options/).
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0Mi4yNi41LXJwbSIsInVwZGF0ZWRJblZlciI6IjQyLjI2LjUtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFzdGVyIiwibGFiZWxzIjpbXX0=-->


---

## Discussion

### Comment by @github-actions on February 28, 2026 at 05:09 AM UTC

<!-- sc-environment-impact-check -->
## SC Environment Impact Assessment

**Overall Impact:** ⚪ **NONE**

No SC Environment-specific impacts detected in this PR.

<details>
<summary>What was checked</summary>

This PR was automatically scanned for:
- Database migrations
- ClowdApp configuration changes
- Kessel integration changes
- AWS service integrations (S3, RDS, ElastiCache)
- Kafka topic changes
- Secrets management changes
- External dependencies
</details>

### Comment by @codecov-commenter on February 28, 2026 at 05:16 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2080?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.45%. Comparing base ([`1659116`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/1659116af8da6c0fc7eeebcfc869454337acafd3?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`d4c2816`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/d4c2816ccf6656dbd0f7f5e1a5ec873aa27923fb?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 2 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #2080   +/-   ##
=======================================
  Coverage   59.45%   59.45%           
=======================================
  Files         134      134           
  Lines        8699     8699           
=======================================
  Hits         5172     5172           
  Misses       2981     2981           
  Partials      546      546           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2080/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2080/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.45% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2080?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2080*
