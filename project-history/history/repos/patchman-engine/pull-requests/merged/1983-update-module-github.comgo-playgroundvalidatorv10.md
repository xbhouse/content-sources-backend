---
type: pull_request
number: 1983
title: "Update module github.com/go-playground/validator/v10 to v10.30.1"
state: merged
author: red-hat-konflux
created: 2025-12-15T08:43:18Z
updated: 2025-12-28T20:50:58Z
closed: 2025-12-28T20:50:54Z
merged: 2025-12-28T20:50:54Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-go-playground-validator-v10-10.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1983
---

# Pull Request #1983: Update module github.com/go-playground/validator/v10 to v10.30.1

**Author**: @red-hat-konflux
**Created**: December 15, 2025 at 08:43 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-go-playground-validator-v10-10.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/go-playground/validator/v10](https://redirect.github.com/go-playground/validator) | `v10.28.0` -> `v10.30.1` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fgo-playground%2fvalidator%2fv10/v10.30.1?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fgo-playground%2fvalidator%2fv10/v10.28.0/v10.30.1?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>go-playground/validator (github.com/go-playground/validator/v10)</summary>

### [`v10.30.1`](https://redirect.github.com/go-playground/validator/releases/tag/v10.30.1): Release 10.30.1

[Compare Source](https://redirect.github.com/go-playground/validator/compare/v10.30.0...v10.30.1)

#### What's Changed

- Feat: uds\_exists validator by [@&#8203;barash-asenov](https://redirect.github.com/barash-asenov) in [#&#8203;1482](https://redirect.github.com/go-playground/validator/pull/1482)
- fix: Revert min limit of e164 regex by [@&#8203;zemzale](https://redirect.github.com/zemzale) in [#&#8203;1516](https://redirect.github.com/go-playground/validator/pull/1516)
- Fix 1513 update ISO 3166-2 codes by [@&#8203;xyz27900](https://redirect.github.com/xyz27900) in [#&#8203;1514](https://redirect.github.com/go-playground/validator/pull/1514)

#### New Contributors

- [@&#8203;barash-asenov](https://redirect.github.com/barash-asenov) made their first contribution in [#&#8203;1482](https://redirect.github.com/go-playground/validator/pull/1482)
- [@&#8203;xyz27900](https://redirect.github.com/xyz27900) made their first contribution in [#&#8203;1514](https://redirect.github.com/go-playground/validator/pull/1514)

**Full Changelog**: <https://github.com/go-playground/validator/compare/v10.30.0...v10.30.1>

### [`v10.30.0`](https://redirect.github.com/go-playground/validator/releases/tag/v10.30.0): Release 10.30.0

[Compare Source](https://redirect.github.com/go-playground/validator/compare/v10.29.0...v10.30.0)

#### What's Changed

- Bump golang.org/x/crypto from 0.45.0 to 0.46.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;1504](https://redirect.github.com/go-playground/validator/pull/1504)
- Bump github.com/gabriel-vasile/mimetype from 1.4.11 to 1.4.12 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;1505](https://redirect.github.com/go-playground/validator/pull/1505)
- docs: document omitzero by [@&#8203;minoritea](https://redirect.github.com/minoritea) in [#&#8203;1509](https://redirect.github.com/go-playground/validator/pull/1509)
- fix: add missing translations for alpha validators by [@&#8203;shindonghwi](https://redirect.github.com/shindonghwi) in [#&#8203;1510](https://redirect.github.com/go-playground/validator/pull/1510)
- fix: resolve panic when using aliases with OR operator by [@&#8203;shindonghwi](https://redirect.github.com/shindonghwi) in [#&#8203;1507](https://redirect.github.com/go-playground/validator/pull/1507)
- fix: resolve panic when using cross-field validators with ValidateMap by [@&#8203;shindonghwi](https://redirect.github.com/shindonghwi) in [#&#8203;1508](https://redirect.github.com/go-playground/validator/pull/1508)

#### New Contributors

- [@&#8203;minoritea](https://redirect.github.com/minoritea) made their first contribution in [#&#8203;1509](https://redirect.github.com/go-playground/validator/pull/1509)
- [@&#8203;shindonghwi](https://redirect.github.com/shindonghwi) made their first contribution in [#&#8203;1510](https://redirect.github.com/go-playground/validator/pull/1510)

**Full Changelog**: <https://github.com/go-playground/validator/compare/v10.29.0...v10.30.0>

### [`v10.29.0`](https://redirect.github.com/go-playground/validator/releases/tag/v10.29.0)

[Compare Source](https://redirect.github.com/go-playground/validator/compare/v10.28.0...v10.29.0)

#### What's Changed

- fix: minor spelling fix in docs by [@&#8203;Perfect5th](https://redirect.github.com/Perfect5th) in [#&#8203;1472](https://redirect.github.com/go-playground/validator/pull/1472)
- Bump golang.org/x/text from 0.29.0 to 0.30.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;1473](https://redirect.github.com/go-playground/validator/pull/1473)
- Bump golang.org/x/crypto from 0.42.0 to 0.43.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;1474](https://redirect.github.com/go-playground/validator/pull/1474)
- Fix integer overflows in test when run on 32bit systems by [@&#8203;gibmat](https://redirect.github.com/gibmat) in [#&#8203;1479](https://redirect.github.com/go-playground/validator/pull/1479)
- fix: exclude modernize linter by [@&#8203;nodivbyzero](https://redirect.github.com/nodivbyzero) in [#&#8203;1487](https://redirect.github.com/go-playground/validator/pull/1487)
- Bump golangci/golangci-lint-action from 8 to 9 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;1490](https://redirect.github.com/go-playground/validator/pull/1490)
- Bump github.com/gabriel-vasile/mimetype from 1.4.10 to 1.4.11 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;1485](https://redirect.github.com/go-playground/validator/pull/1485)
- Support for ISO 9362:2022 BIC (SWIFT) codes by [@&#8203;fira42073](https://redirect.github.com/fira42073) in [#&#8203;1478](https://redirect.github.com/go-playground/validator/pull/1478)
- Bump golang.org/x/crypto from 0.43.0 to 0.44.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;1492](https://redirect.github.com/go-playground/validator/pull/1492)
- Fix: validation now rejects phone codes starting with +0 by [@&#8203;nodivbyzero](https://redirect.github.com/nodivbyzero) in [#&#8203;1476](https://redirect.github.com/go-playground/validator/pull/1476)
- Bump golang.org/x/crypto from 0.44.0 to 0.45.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;1495](https://redirect.github.com/go-playground/validator/pull/1495)
- Bump actions/checkout from 5 to 6 by [@&#8203;dependabot](https://redirect.github.com/dependabot)\[bot] in [#&#8203;1497](https://redirect.github.com/go-playground/validator/pull/1497)
- fix/1500:Update Sierra Leone currency code from SLL to SLE by [@&#8203;princekm096](https://redirect.github.com/princekm096) in [#&#8203;1501](https://redirect.github.com/go-playground/validator/pull/1501)
- Fix/1481 skip invalid type validations by [@&#8203;KaranLathiya](https://redirect.github.com/KaranLathiya) in [#&#8203;1498](https://redirect.github.com/go-playground/validator/pull/1498)
- Fix 1502 update ccy codes by [@&#8203;princekm096](https://redirect.github.com/princekm096) in [#&#8203;1503](https://redirect.github.com/go-playground/validator/pull/1503)
- Added alphanumspace string validator by [@&#8203;haribabuk113](https://redirect.github.com/haribabuk113) in [#&#8203;1484](https://redirect.github.com/go-playground/validator/pull/1484)
- `excluded_unless` bug fix by [@&#8203;chargraves85](https://redirect.github.com/chargraves85) in [#&#8203;1307](https://redirect.github.com/go-playground/validator/pull/1307)

#### New Contributors

- [@&#8203;Perfect5th](https://redirect.github.com/Perfect5th) made their first contribution in [#&#8203;1472](https://redirect.github.com/go-playground/validator/pull/1472)
- [@&#8203;gibmat](https://redirect.github.com/gibmat) made their first contribution in [#&#8203;1479](https://redirect.github.com/go-playground/validator/pull/1479)
- [@&#8203;fira42073](https://redirect.github.com/fira42073) made their first contribution in [#&#8203;1478](https://redirect.github.com/go-playground/validator/pull/1478)
- [@&#8203;princekm096](https://redirect.github.com/princekm096) made their first contribution in [#&#8203;1501](https://redirect.github.com/go-playground/validator/pull/1501)
- [@&#8203;KaranLathiya](https://redirect.github.com/KaranLathiya) made their first contribution in [#&#8203;1498](https://redirect.github.com/go-playground/validator/pull/1498)
- [@&#8203;haribabuk113](https://redirect.github.com/haribabuk113) made their first contribution in [#&#8203;1484](https://redirect.github.com/go-playground/validator/pull/1484)
- [@&#8203;chargraves85](https://redirect.github.com/chargraves85) made their first contribution in [#&#8203;1307](https://redirect.github.com/go-playground/validator/pull/1307)

**Full Changelog**: <https://github.com/go-playground/validator/compare/v10.28.0...v10.29.0>

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
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0MS45MC4xLXJwbSIsInVwZGF0ZWRJblZlciI6IjQxLjkwLjEtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFzdGVyIiwibGFiZWxzIjpbXX0=-->


---

## Discussion

### Comment by @jira-linking on December 15, 2025 at 08:43 AM UTC

Commits missing Jira IDs:
f85155cb9721708fb34eedf94d73f38f5d710fd4


### Comment by @codecov-commenter on December 22, 2025 at 04:48 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1983?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.01%. Comparing base ([`deb9cf6`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/deb9cf64184a89ca320feb0d18475b1426ebdd8b?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`f85155c`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/f85155cb9721708fb34eedf94d73f38f5d710fd4?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 3 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1983   +/-   ##
=======================================
  Coverage   59.01%   59.01%           
=======================================
  Files         131      131           
  Lines        8493     8493           
=======================================
  Hits         5012     5012           
  Misses       2947     2947           
  Partials      534      534           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1983/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1983/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.01% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1983?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1983*
