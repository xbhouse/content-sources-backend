---
type: pull_request
number: 1845
title: "Update module github.com/go-kratos/kratos/v2 to v2.9.1"
state: merged
author: red-hat-konflux
created: 2025-09-14T12:31:13Z
updated: 2025-10-08T16:16:13Z
closed: 2025-09-14T12:43:03Z
merged: 2025-09-14T12:43:03Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-go-kratos-kratos-v2-2.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1845
---

# Pull Request #1845: Update module github.com/go-kratos/kratos/v2 to v2.9.1

**Author**: @red-hat-konflux
**Created**: September 14, 2025 at 12:31 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-go-kratos-kratos-v2-2.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/go-kratos/kratos/v2](https://redirect.github.com/go-kratos/kratos) | `v2.8.4` -> `v2.9.1` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fgo-kratos%2fkratos%2fv2/v2.9.1?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fgo-kratos%2fkratos%2fv2/v2.8.4/v2.9.1?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

### Release Notes

<details>
<summary>go-kratos/kratos (github.com/go-kratos/kratos/v2)</summary>

### [`v2.9.1`](https://redirect.github.com/go-kratos/kratos/releases/tag/v2.9.1)

[Compare Source](https://redirect.github.com/go-kratos/kratos/compare/v2.9.0...v2.9.1)

##### Dependencies

- deps: retract v2.9.0 ([#&#8203;3723](https://redirect.github.com/go-kratos/kratos/issues/3723))

##### Bug Fixes

- fix(transport/http): resolve breaking change with ResponseTransporter interface ([#&#8203;3721](https://redirect.github.com/go-kratos/kratos/issues/3721))

##### Others

- Revert "perf(transport/http): optimize URL construction with url.URL for bett…" ([#&#8203;3722](https://redirect.github.com/go-kratos/kratos/issues/3722))

##### New Contributors

- [@&#8203;qingbozhang](https://redirect.github.com/qingbozhang) made their first contribution in [#&#8203;3721](https://redirect.github.com/go-kratos/kratos/pull/3721)

**Full Changelog**: <https://github.com/go-kratos/kratos/compare/v2.9.0...v2.9.1>

### [`v2.9.0`](https://redirect.github.com/go-kratos/kratos/releases/tag/v2.9.0)

[Compare Source](https://redirect.github.com/go-kratos/kratos/compare/v2.8.4...v2.9.0)

##### Dependencies

- deps(middleware): upgrade protovalidate to resolve compatibility issue ([#&#8203;3706](https://redirect.github.com/go-kratos/kratos/issues/3706))

##### New Features

- feat(protoc-gen-go-http): mark deprecated methods in generated code ([#&#8203;3711](https://redirect.github.com/go-kratos/kratos/issues/3711))
- feat: add MCP transport ([#&#8203;3646](https://redirect.github.com/go-kratos/kratos/issues/3646))

##### Bug Fixes

- fix(encoding): extract all form fields even if some unsupported  ([#&#8203;3694](https://redirect.github.com/go-kratos/kratos/issues/3694))
- fix meta variable ([#&#8203;3712](https://redirect.github.com/go-kratos/kratos/issues/3712))
- fix: honor shutdown timeout when provided context already canceled ([#&#8203;3695](https://redirect.github.com/go-kratos/kratos/issues/3695))
- fix stale entries before the each pick operation ([#&#8203;3690](https://redirect.github.com/go-kratos/kratos/issues/3690))
- fix: replace strconv.Itoa with strconv.FormatInt for int64 handling ([#&#8203;3667](https://redirect.github.com/go-kratos/kratos/issues/3667))
- fix: solve have syntax error when execute make command. ([#&#8203;3682](https://redirect.github.com/go-kratos/kratos/issues/3682))
- fix(internal/host): Extract optimizes the return of the minimum index IP ([#&#8203;3421](https://redirect.github.com/go-kratos/kratos/issues/3421))
- fix: Fixed zookeeper watcher exiting the loop under exceptional conditions, causing service discovery to terminate ([#&#8203;3517](https://redirect.github.com/go-kratos/kratos/issues/3517))

##### Chores

- chore: add deepwiki badge ([#&#8203;3701](https://redirect.github.com/go-kratos/kratos/issues/3701))

##### Others

- perf(transport/http): optimize URL construction with url.URL for better performance ([#&#8203;3678](https://redirect.github.com/go-kratos/kratos/issues/3678))
- perf(transport/http/binding): optimize EncodeURL performance for paths without placeholders ([#&#8203;3679](https://redirect.github.com/go-kratos/kratos/issues/3679))
- refactor(all): replace atomic functions with atomic types ([#&#8203;3699](https://redirect.github.com/go-kratos/kratos/issues/3699))
- contrib/register/nacos: refactor to use maps.Clone ([#&#8203;3703](https://redirect.github.com/go-kratos/kratos/issues/3703))
- \*: use net.JoinHostPort to improve IPv6 compatible ([#&#8203;3675](https://redirect.github.com/go-kratos/kratos/issues/3675))
- perf(config/env): use strings.LastIndexByte instead of strings.LastIndex ([#&#8203;3660](https://redirect.github.com/go-kratos/kratos/issues/3660))
- Add trendshift badge ([#&#8203;3655](https://redirect.github.com/go-kratos/kratos/issues/3655))
- perf(metadata): simplify Metadata.Add by avoiding redundant strings.ToLower call ([#&#8203;3671](https://redirect.github.com/go-kratos/kratos/issues/3671))
- Mcp options ([#&#8203;3652](https://redirect.github.com/go-kratos/kratos/issues/3652))
- Revert "refactor: replace repeated error reasons with constants ([#&#8203;3612](https://redirect.github.com/go-kratos/kratos/issues/3612))" ([#&#8203;3651](https://redirect.github.com/go-kratos/kratos/issues/3651))
- perf(config/env): use strings.Cut  to optimize env load method ([#&#8203;3645](https://redirect.github.com/go-kratos/kratos/issues/3645))
- add mcp middleware ([#&#8203;3649](https://redirect.github.com/go-kratos/kratos/issues/3649))
- refactor(contrib/registry/etcd): use Namespace Name and ID as the unique keys ([#&#8203;3594](https://redirect.github.com/go-kratos/kratos/issues/3594))
- go mod tidy ([#&#8203;3648](https://redirect.github.com/go-kratos/kratos/issues/3648))
- build(deps): bump golang.org/x/net from 0.23.0 to 0.33.0 ([#&#8203;3527](https://redirect.github.com/go-kratos/kratos/issues/3527))
- github: Remove submodule dependencies in Dependabot. ([#&#8203;3643](https://redirect.github.com/go-kratos/kratos/issues/3643))
- perf(encoding/form): optimize camelCase and snakeCase conversion ([#&#8203;3592](https://redirect.github.com/go-kratos/kratos/issues/3592))
- refactor: extract key for reuse and maintainability ([#&#8203;3602](https://redirect.github.com/go-kratos/kratos/issues/3602))

##### New Contributors

- [@&#8203;uucloud](https://redirect.github.com/uucloud) made their first contribution in [#&#8203;3561](https://redirect.github.com/go-kratos/kratos/pull/3561)
- [@&#8203;ch3nnn](https://redirect.github.com/ch3nnn) made their first contribution in [#&#8203;3624](https://redirect.github.com/go-kratos/kratos/pull/3624)
- [@&#8203;yuluo-yx](https://redirect.github.com/yuluo-yx) made their first contribution in [#&#8203;3625](https://redirect.github.com/go-kratos/kratos/pull/3625)
- [@&#8203;Piwriw](https://redirect.github.com/Piwriw) made their first contribution in [#&#8203;3642](https://redirect.github.com/go-kratos/kratos/pull/3642)
- [@&#8203;hengyumo](https://redirect.github.com/hengyumo) made their first contribution in [#&#8203;3517](https://redirect.github.com/go-kratos/kratos/pull/3517)
- [@&#8203;HoronLee](https://redirect.github.com/HoronLee) made their first contribution in [#&#8203;3706](https://redirect.github.com/go-kratos/kratos/pull/3706)
- [@&#8203;cuiweixie](https://redirect.github.com/cuiweixie) made their first contribution in [#&#8203;3703](https://redirect.github.com/go-kratos/kratos/pull/3703)

**Full Changelog**: <https://github.com/go-kratos/kratos/compare/v2.8.4...v2.9.0>

</details>

---

### Configuration

📅 **Schedule**: Branch creation - "after 5am on sunday" in timezone Europe/Prague, Automerge - At any time (no schedule defined).

🚦 **Automerge**: Enabled.

♻ **Rebasing**: Whenever PR is behind base branch, or you tick the rebase/retry checkbox.

🔕 **Ignore**: Close this PR and you won't be reminded about this update again.

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`.

This PR has been generated by [MintMaker](https://redirect.github.com/konflux-ci/mintmaker) (powered by [Renovate Bot](https://redirect.github.com/renovatebot/renovate)).
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0MS45MC4xLXJwbSIsInVwZGF0ZWRJblZlciI6IjQxLjkwLjEtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFzdGVyIiwibGFiZWxzIjpbXX0=-->


---

## Discussion

### Comment by @jira-linking on September 14, 2025 at 12:31 PM UTC

Commits missing Jira IDs:
24bcf6e8200ea0b7964bff6ec7755ba1fc996cd3


### Comment by @codecov-commenter on September 14, 2025 at 12:36 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1845?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 54.87%. Comparing base ([`15f589c`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/15f589cdd9fc3543cd0f164fd8dcbb07c5bded23?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`24bcf6e`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/24bcf6e8200ea0b7964bff6ec7755ba1fc996cd3?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1845   +/-   ##
=======================================
  Coverage   54.87%   54.87%           
=======================================
  Files         140      140           
  Lines       10878    10878           
=======================================
  Hits         5969     5969           
  Misses       4373     4373           
  Partials      536      536           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1845/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1845/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `54.87% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1845?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1845*
