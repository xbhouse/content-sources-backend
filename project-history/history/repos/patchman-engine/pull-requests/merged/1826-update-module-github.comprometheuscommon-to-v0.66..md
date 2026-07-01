---
type: pull_request
number: 1826
title: "Update module github.com/prometheus/common to v0.66.1"
state: merged
author: red-hat-konflux
created: 2025-09-07T08:15:28Z
updated: 2025-09-07T08:31:38Z
closed: 2025-09-07T08:31:38Z
merged: 2025-09-07T08:31:38Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-prometheus-common-0.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1826
---

# Pull Request #1826: Update module github.com/prometheus/common to v0.66.1

**Author**: @red-hat-konflux
**Created**: September 07, 2025 at 08:15 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-prometheus-common-0.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/prometheus/common](https://redirect.github.com/prometheus/common) | `v0.65.0` -> `v0.66.1` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fprometheus%2fcommon/v0.66.1?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fprometheus%2fcommon/v0.65.0/v0.66.1?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

### Release Notes

<details>
<summary>prometheus/common (github.com/prometheus/common)</summary>

### [`v0.66.1`](https://redirect.github.com/prometheus/common/blob/HEAD/CHANGELOG.md#v0661--2025-09-05)

[Compare Source](https://redirect.github.com/prometheus/common/compare/v0.66.0...v0.66.1)

This release has no functional changes, it just drops the dependencies `github.com/grafana/regexp` and `go.uber.org/atomic` and replaces `gopkg.in/yaml.v2` with `go.yaml.in/yaml/v2` (a drop-in replacement).

##### What's Changed

- Revert "Use github.com/grafana/regexp instead of regexp" by [@&#8203;aknuds1](https://redirect.github.com/aknuds1) in [#&#8203;835](https://redirect.github.com/prometheus/common/pull/835)
- Move to supported version of yaml parser by [@&#8203;dims](https://redirect.github.com/dims) in [#&#8203;834](https://redirect.github.com/prometheus/common/pull/834)
- Revert "Use go.uber.org/atomic instead of sync/atomic ([#&#8203;825](https://redirect.github.com/prometheus/common/issues/825))" by [@&#8203;aknuds1](https://redirect.github.com/aknuds1) in [#&#8203;838](https://redirect.github.com/prometheus/common/pull/838)

**Full Changelog**: <https://github.com/prometheus/common/compare/v1.20.99...v0.66.1>

### [`v0.66.0`](https://redirect.github.com/prometheus/common/blob/HEAD/CHANGELOG.md#v0660--2025-09-02)

[Compare Source](https://redirect.github.com/prometheus/common/compare/v0.65.0...v0.66.0)

##### ⚠️ Breaking Changes ⚠️

- A default-constructed TextParser will be invalid. It must have a valid `scheme` set, so users should use the NewTextParser function to create a valid TextParser. Otherwise parsing will panic with "Invalid name validation scheme requested: unset".

##### What's Changed

- model: add constants for type and unit labels. by [@&#8203;bwplotka](https://redirect.github.com/bwplotka) in [#&#8203;801](https://redirect.github.com/prometheus/common/pull/801)

- model.ValidationScheme: Support encoding as YAML by [@&#8203;aknuds1](https://redirect.github.com/aknuds1) in [#&#8203;799](https://redirect.github.com/prometheus/common/pull/799)

- fix(promslog): always print time.Duration values as go duration strings by [@&#8203;tjhop](https://redirect.github.com/tjhop) in [#&#8203;798](https://redirect.github.com/prometheus/common/pull/798)

- Add `ValidationScheme` methods `IsValidMetricName` and `IsValidLabelName` by [@&#8203;aknuds1](https://redirect.github.com/aknuds1) in [#&#8203;806](https://redirect.github.com/prometheus/common/pull/806)

- Fix delimited proto not escaped correctly by [@&#8203;thampiotr](https://redirect.github.com/thampiotr) in [#&#8203;809](https://redirect.github.com/prometheus/common/pull/809)

- Decoder: Remove use of global name validation and add validation by [@&#8203;ywwg](https://redirect.github.com/ywwg) in [#&#8203;808](https://redirect.github.com/prometheus/common/pull/808)

- ValidationScheme implements pflag.Value and json.Marshaler/Unmarshaler interfaces by [@&#8203;juliusmh](https://redirect.github.com/juliusmh) in [#&#8203;807](https://redirect.github.com/prometheus/common/pull/807)

- expfmt: Add NewTextParser function by [@&#8203;aknuds1](https://redirect.github.com/aknuds1) in [#&#8203;816](https://redirect.github.com/prometheus/common/pull/816)

- Enable the godot linter by [@&#8203;aknuds1](https://redirect.github.com/aknuds1) in [#&#8203;821](https://redirect.github.com/prometheus/common/pull/821)

- Enable usestdlibvars linter by [@&#8203;aknuds1](https://redirect.github.com/aknuds1) in [#&#8203;820](https://redirect.github.com/prometheus/common/pull/820)

- Enable unconvert linter by [@&#8203;aknuds1](https://redirect.github.com/aknuds1) in [#&#8203;819](https://redirect.github.com/prometheus/common/pull/819)

- Enable the fatcontext linter by [@&#8203;aknuds1](https://redirect.github.com/aknuds1) in [#&#8203;822](https://redirect.github.com/prometheus/common/pull/822)

- Enable gocritic linter by [@&#8203;aknuds1](https://redirect.github.com/aknuds1) in [#&#8203;818](https://redirect.github.com/prometheus/common/pull/818)

- Use go.uber.org/atomic instead of sync/atomic by [@&#8203;aknuds1](https://redirect.github.com/aknuds1) in [#&#8203;825](https://redirect.github.com/prometheus/common/pull/825)

- Enable revive rule unused-parameter by [@&#8203;aknuds1](https://redirect.github.com/aknuds1) in [#&#8203;824](https://redirect.github.com/prometheus/common/pull/824)

- Enable revive rules by [@&#8203;aknuds1](https://redirect.github.com/aknuds1) in [#&#8203;823](https://redirect.github.com/prometheus/common/pull/823)

- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;802](https://redirect.github.com/prometheus/common/pull/802)

- Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [#&#8203;803](https://redirect.github.com/prometheus/common/pull/803)

- Sync .golangci.yml with prometheus/prometheus by [@&#8203;aknuds1](https://redirect.github.com/aknuds1) in [#&#8203;817](https://redirect.github.com/prometheus/common/pull/817)

- ci: update upload-actions by [@&#8203;ywwg](https://redirect.github.com/ywwg) in [#&#8203;814](https://redirect.github.com/prometheus/common/pull/814)

- docs: fix typo in expfmt.Negotiate by [@&#8203;wmcram](https://redirect.github.com/wmcram) in [#&#8203;813](https://redirect.github.com/prometheus/common/pull/813)

- build(deps): bump golang.org/x/net from 0.40.0 to 0.41.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;800](https://redirect.github.com/prometheus/common/pull/800)

- build(deps): bump golang.org/x/net from 0.41.0 to 0.42.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;810](https://redirect.github.com/prometheus/common/pull/810)

- build(deps): bump github.com/stretchr/testify from 1.10.0 to 1.11.1 in /assets by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;826](https://redirect.github.com/prometheus/common/pull/826)

- build(deps): bump google.golang.org/protobuf from 1.36.6 to 1.36.8 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;830](https://redirect.github.com/prometheus/common/pull/830)

- build(deps): bump golang.org/x/net from 0.42.0 to 0.43.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;829](https://redirect.github.com/prometheus/common/pull/829)

- build(deps): bump github.com/stretchr/testify from 1.10.0 to 1.11.1 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;827](https://redirect.github.com/prometheus/common/pull/827)

##### New Contributors

- [@&#8203;aknuds1](https://redirect.github.com/aknuds1) made their first contribution in [#&#8203;799](https://redirect.github.com/prometheus/common/pull/799)
- [@&#8203;thampiotr](https://redirect.github.com/thampiotr) made their first contribution in [#&#8203;809](https://redirect.github.com/prometheus/common/pull/809)
- [@&#8203;wmcram](https://redirect.github.com/wmcram) made their first contribution in [#&#8203;813](https://redirect.github.com/prometheus/common/pull/813)
- [@&#8203;juliusmh](https://redirect.github.com/juliusmh) made their first contribution in [#&#8203;807](https://redirect.github.com/prometheus/common/pull/807)

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

### Comment by @jira-linking on September 07, 2025 at 08:15 AM UTC

Commits missing Jira IDs:
f0e104cd9c46cc71ba882ee5ced02eccfbefcdb1


### Comment by @codecov-commenter on September 07, 2025 at 08:20 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1826?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 54.85%. Comparing base ([`500dc4c`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/500dc4c7de3b85648173e55bbaae188a6f19e728?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`f0e104c`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/f0e104cd9c46cc71ba882ee5ced02eccfbefcdb1?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1826   +/-   ##
=======================================
  Coverage   54.85%   54.85%           
=======================================
  Files         140      140           
  Lines       10876    10876           
=======================================
  Hits         5966     5966           
  Misses       4374     4374           
  Partials      536      536           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1826/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1826/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `54.85% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1826?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1826*
