---
type: pull_request
number: 1852
title: "Update module github.com/gin-gonic/gin to v1.11.0"
state: merged
author: red-hat-konflux
created: 2025-09-21T04:22:33Z
updated: 2025-10-08T16:16:12Z
closed: 2025-09-24T08:22:39Z
merged: 2025-09-24T08:22:39Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-gin-gonic-gin-1.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1852
---

# Pull Request #1852: Update module github.com/gin-gonic/gin to v1.11.0

**Author**: @red-hat-konflux
**Created**: September 21, 2025 at 04:22 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-gin-gonic-gin-1.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/gin-gonic/gin](https://redirect.github.com/gin-gonic/gin) | `v1.10.1` -> `v1.11.0` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fgin-gonic%2fgin/v1.11.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fgin-gonic%2fgin/v1.10.1/v1.11.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>gin-gonic/gin (github.com/gin-gonic/gin)</summary>

### [`v1.11.0`](https://redirect.github.com/gin-gonic/gin/blob/HEAD/CHANGELOG.md#Gin-v1110)

[Compare Source](https://redirect.github.com/gin-gonic/gin/compare/v1.10.1...v1.11.0)

##### Features

- feat(gin): Experimental support for HTTP/3 using quic-go/quic-go ([#&#8203;3210](https://redirect.github.com/gin-gonic/gin/pull/3210))
- feat(form): add array collection format in form binding ([#&#8203;3986](https://redirect.github.com/gin-gonic/gin/pull/3986)), add custom string slice for form tag unmarshal ([#&#8203;3970](https://redirect.github.com/gin-gonic/gin/pull/3970))
- feat(binding): add BindPlain ([#&#8203;3904](https://redirect.github.com/gin-gonic/gin/pull/3904))
- feat(fs): Export, test and document OnlyFilesFS ([#&#8203;3939](https://redirect.github.com/gin-gonic/gin/pull/3939))
- feat(binding): add support for unixMilli and unixMicro ([#&#8203;4190](https://redirect.github.com/gin-gonic/gin/pull/4190))
- feat(form): Support default values for collections in form binding ([#&#8203;4048](https://redirect.github.com/gin-gonic/gin/pull/4048))
- feat(context): GetXxx added support for more go native types ([#&#8203;3633](https://redirect.github.com/gin-gonic/gin/pull/3633))

##### Enhancements

- perf(context): optimize getMapFromFormData performance ([#&#8203;4339](https://redirect.github.com/gin-gonic/gin/pull/4339))
- refactor(tree): replace string(/) with "/" in node.insertChild ([#&#8203;4354](https://redirect.github.com/gin-gonic/gin/pull/4354))
- refactor(render): remove headers parameter from writeHeader ([#&#8203;4353](https://redirect.github.com/gin-gonic/gin/pull/4353))
- refactor(context): simplify "GetType()" functions ([#&#8203;4080](https://redirect.github.com/gin-gonic/gin/pull/4080))
- refactor(slice): simplify SliceValidationError Error method ([#&#8203;3910](https://redirect.github.com/gin-gonic/gin/pull/3910))
- refactor(context):Avoid using filepath.Dir twice in SaveUploadedFile ([#&#8203;4181](https://redirect.github.com/gin-gonic/gin/pull/4181))
- refactor(context): refactor context handling and improve test robustness ([#&#8203;4066](https://redirect.github.com/gin-gonic/gin/pull/4066))
- refactor(binding): use strings.Cut to replace strings.Index ([#&#8203;3522](https://redirect.github.com/gin-gonic/gin/pull/3522))
- refactor(context): add an optional permission parameter to SaveUploadedFile ([#&#8203;4068](https://redirect.github.com/gin-gonic/gin/pull/4068))
- refactor(context): verify URL is Non-nil in initQueryCache() ([#&#8203;3969](https://redirect.github.com/gin-gonic/gin/pull/3969))
- refactor(context): YAML judgment logic in Negotiate ([#&#8203;3966](https://redirect.github.com/gin-gonic/gin/pull/3966))
- tree: replace the self-defined 'min' to official one ([#&#8203;3975](https://redirect.github.com/gin-gonic/gin/pull/3975))
- context: Remove redundant filepath.Dir usage ([#&#8203;4181](https://redirect.github.com/gin-gonic/gin/pull/4181))

##### Bug Fixes

- fix: prevent middleware re-entry issue in HandleContext ([#&#8203;3987](https://redirect.github.com/gin-gonic/gin/pull/3987))
- fix(binding): prevent duplicate decoding and add validation in decodeToml ([#&#8203;4193](https://redirect.github.com/gin-gonic/gin/pull/4193))
- fix(gin): Do not panic when handling method not allowed on empty tree ([#&#8203;4003](https://redirect.github.com/gin-gonic/gin/pull/4003))
- fix(gin): data race warning for gin mode ([#&#8203;1580](https://redirect.github.com/gin-gonic/gin/pull/1580))
- fix(context): verify URL is Non-nil in initQueryCache() ([#&#8203;3969](https://redirect.github.com/gin-gonic/gin/pull/3969))
- fix(context): YAML judgment logic in Negotiate ([#&#8203;3966](https://redirect.github.com/gin-gonic/gin/pull/3966))
- fix(context): check handler is nil ([#&#8203;3413](https://redirect.github.com/gin-gonic/gin/pull/3413))
- fix(readme): fix broken link to English documentation ([#&#8203;4222](https://redirect.github.com/gin-gonic/gin/pull/4222))
- fix(tree): Keep panic infos consistent when wildcard type build faild ([#&#8203;4077](https://redirect.github.com/gin-gonic/gin/pull/4077))

##### Build process updates / CI

- ci: integrate Trivy vulnerability scanning into CI workflow ([#&#8203;4359](https://redirect.github.com/gin-gonic/gin/pull/4359))
- ci: support Go 1.25 in CI/CD ([#&#8203;4341](https://redirect.github.com/gin-gonic/gin/pull/4341))
- build(deps): upgrade github.com/bytedance/sonic from v1.13.2 to v1.14.0 ([#&#8203;4342](https://redirect.github.com/gin-gonic/gin/pull/4342))
- ci: add Go version 1.24 to GitHub Actions ([#&#8203;4154](https://redirect.github.com/gin-gonic/gin/pull/4154))
- build: update Gin minimum Go version to 1.21 ([#&#8203;3960](https://redirect.github.com/gin-gonic/gin/pull/3960))
- ci(lint): enable new linters (testifylint, usestdlibvars, perfsprint, etc.) ([#&#8203;4010](https://redirect.github.com/gin-gonic/gin/pull/4010), [#&#8203;4091](https://redirect.github.com/gin-gonic/gin/pull/4091), [#&#8203;4090](https://redirect.github.com/gin-gonic/gin/pull/4090))
- ci(lint): update workflows and improve test request consistency ([#&#8203;4126](https://redirect.github.com/gin-gonic/gin/pull/4126))

##### Dependency updates

- chore(deps): bump google.golang.org/protobuf from 1.36.6 to 1.36.9 ([#&#8203;4346](https://redirect.github.com/gin-gonic/gin/pull/4346), [#&#8203;4356](https://redirect.github.com/gin-gonic/gin/pull/4356))
- chore(deps): bump github.com/stretchr/testify from 1.10.0 to 1.11.1 ([#&#8203;4347](https://redirect.github.com/gin-gonic/gin/pull/4347))
- chore(deps): bump actions/setup-go from 5 to 6 ([#&#8203;4351](https://redirect.github.com/gin-gonic/gin/pull/4351))
- chore(deps): bump github.com/quic-go/quic-go from 0.53.0 to 0.54.0 ([#&#8203;4328](https://redirect.github.com/gin-gonic/gin/pull/4328))
- chore(deps): bump golang.org/x/net from 0.33.0 to 0.38.0 ([#&#8203;4178](https://redirect.github.com/gin-gonic/gin/pull/4178), [#&#8203;4221](https://redirect.github.com/gin-gonic/gin/pull/4221))
- chore(deps): bump github.com/go-playground/validator/v10 from 10.20.0 to 10.22.1 ([#&#8203;4052](https://redirect.github.com/gin-gonic/gin/pull/4052))

##### Documentation updates

- docs(changelog): update release notes for Gin v1.10.1 ([#&#8203;4360](https://redirect.github.com/gin-gonic/gin/pull/4360))
- docs: Fixing English grammar mistakes and awkward sentence structure in doc/doc.md ([#&#8203;4207](https://redirect.github.com/gin-gonic/gin/pull/4207))
- docs: update documentation and release notes for Gin v1.10.0 ([#&#8203;3953](https://redirect.github.com/gin-gonic/gin/pull/3953))
- docs: fix typo in Gin Quick Start ([#&#8203;3997](https://redirect.github.com/gin-gonic/gin/pull/3997))
- docs: fix comment and link issues ([#&#8203;4205](https://redirect.github.com/gin-gonic/gin/pull/4205), [#&#8203;3938](https://redirect.github.com/gin-gonic/gin/pull/3938))
- docs: fix route group example code ([#&#8203;4020](https://redirect.github.com/gin-gonic/gin/pull/4020))
- docs(readme): add Portuguese documentation ([#&#8203;4078](https://redirect.github.com/gin-gonic/gin/pull/4078))
- docs(context): fix some function names in comment ([#&#8203;4079](https://redirect.github.com/gin-gonic/gin/pull/4079))

***

</details>

---

### Configuration

📅 **Schedule**: Branch creation - "after 5am on sunday" in timezone Europe/Prague, Automerge - At any time (no schedule defined).

🚦 **Automerge**: Enabled.

♻ **Rebasing**: Whenever PR becomes conflicted, or you tick the rebase/retry checkbox.

🔕 **Ignore**: Close this PR and you won't be reminded about this update again.

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`.

This PR has been generated by [MintMaker](https://redirect.github.com/konflux-ci/mintmaker) (powered by [Renovate Bot](https://redirect.github.com/renovatebot/renovate)).
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0MS45MC4xLXJwbSIsInVwZGF0ZWRJblZlciI6IjQxLjkwLjEtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFzdGVyIiwibGFiZWxzIjpbXX0=-->


---

## Discussion

### Comment by @jira-linking on September 21, 2025 at 04:22 AM UTC

Commits missing Jira IDs:
177bb25a6db61192b44f56dd9d71dce171a64932


### Comment by @codecov-commenter on September 22, 2025 at 12:23 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1852?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 54.53%. Comparing base ([`03fcee1`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/03fcee1b614bc7397c067e30d856189cb919994c?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`177bb25`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/177bb25a6db61192b44f56dd9d71dce171a64932?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 1 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1852   +/-   ##
=======================================
  Coverage   54.53%   54.53%           
=======================================
  Files         138      138           
  Lines       10743    10743           
=======================================
  Hits         5859     5859           
  Misses       4346     4346           
  Partials      538      538           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1852/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1852/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `54.53% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1852?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1852*
