---
type: pull_request
number: 2093
title: "Update module github.com/project-kessel/kessel-sdk-go to v1.5.0"
state: merged
author: red-hat-konflux
created: 2026-03-09T05:25:23Z
updated: 2026-03-09T21:28:42Z
closed: 2026-03-09T18:18:29Z
merged: 2026-03-09T18:18:29Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-project-kessel-kessel-sdk-go-1.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2093
---

# Pull Request #2093: Update module github.com/project-kessel/kessel-sdk-go to v1.5.0

**Author**: @red-hat-konflux
**Created**: March 09, 2026 at 05:25 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-project-kessel-kessel-sdk-go-1.x`

## Description

This PR contains the following updates:

| Package | Change | [Age](https://docs.renovatebot.com/merge-confidence/) | [Confidence](https://docs.renovatebot.com/merge-confidence/) |
|---|---|---|---|
| [github.com/project-kessel/kessel-sdk-go](https://redirect.github.com/project-kessel/kessel-sdk-go) | `v1.4.0` -> `v1.5.0` | ![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fproject-kessel%2fkessel-sdk-go/v1.5.0?slim=true) | ![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fproject-kessel%2fkessel-sdk-go/v1.4.0/v1.5.0?slim=true) |

---

### Release Notes

<details>
<summary>project-kessel/kessel-sdk-go (github.com/project-kessel/kessel-sdk-go)</summary>

### [`v1.5.0`](https://redirect.github.com/project-kessel/kessel-sdk-go/releases/tag/v1.5.0)

[Compare Source](https://redirect.github.com/project-kessel/kessel-sdk-go/compare/v1.4.0...v1.5.0)

#### What's Changed

- RHCLOUD-43506 - Add CI worklows by [@&#8203;lennysgarage](https://redirect.github.com/lennysgarage) in [#&#8203;24](https://redirect.github.com/project-kessel/kessel-sdk-go/pull/24)
- Update generated protobuf files by [@&#8203;github-actions](https://redirect.github.com/github-actions)\[bot] in [#&#8203;25](https://redirect.github.com/project-kessel/kessel-sdk-go/pull/25)
- Bump actions/checkout from 5 to 6 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;23](https://redirect.github.com/project-kessel/kessel-sdk-go/pull/23)
- Bump peter-evans/create-pull-request from 7 to 8 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;27](https://redirect.github.com/project-kessel/kessel-sdk-go/pull/27)
- Bump buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go from 1.36.10-20250912141014-52f32327d4b0.1 to 1.36.10-20251209175733-2a1774d88802.1 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;28](https://redirect.github.com/project-kessel/kessel-sdk-go/pull/28)
- Bump buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go from 1.36.10-20251209175733-2a1774d88802.1 to 1.36.11-20251209175733-2a1774d88802.1 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;30](https://redirect.github.com/project-kessel/kessel-sdk-go/pull/30)
- Update generated protobuf files by [@&#8203;github-actions](https://redirect.github.com/github-actions)\[bot] in [#&#8203;33](https://redirect.github.com/project-kessel/kessel-sdk-go/pull/33)

**Full Changelog**: <https://github.com/project-kessel/kessel-sdk-go/compare/v1.4.0...v1.5.0>

</details>

---

### Configuration

📅 **Schedule**: Branch creation - Between 03:00 AM and 10:59 AM, only on Monday ( * 3-10 * * 1 ) in timezone Europe/Prague, Automerge - At any time (no schedule defined).

🚦 **Automerge**: Enabled.

♻ **Rebasing**: Whenever PR becomes conflicted, or you tick the rebase/retry checkbox.

🔕 **Ignore**: Close this PR and you won't be reminded about this update again.

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

### Comment by @github-actions on March 09, 2026 at 05:25 AM UTC

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

### Comment by @codecov-commenter on March 09, 2026 at 05:31 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2093?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.40%. Comparing base ([`b1792ab`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/b1792abb3861101c26f9b602ec2c244ca15a4437?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`515e47e`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/515e47e14736cd2456f704387f9f79fa439ce1b9?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 1 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #2093      +/-   ##
==========================================
+ Coverage   59.37%   59.40%   +0.02%     
==========================================
  Files         134      134              
  Lines        8707     8707              
==========================================
+ Hits         5170     5172       +2     
+ Misses       2991     2989       -2     
  Partials      546      546              
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2093/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2093/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.40% <ø> (+0.02%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2093?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

## Reviews

### Review by @MichaelMraka - Approved on March 09, 2026 at 09:05 AM UTC

### Review by @TenSt - Approved on March 09, 2026 at 10:10 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2093*
