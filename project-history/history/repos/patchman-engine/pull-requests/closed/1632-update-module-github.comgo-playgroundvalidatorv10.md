---
type: pull_request
number: 1632
title: "Update module github.com/go-playground/validator/v10 to v10.26.0"
state: closed
author: red-hat-konflux
created: 2025-04-06T09:25:46Z
updated: 2025-04-23T09:06:11Z
closed: 2025-04-23T09:06:05Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-go-playground-validator-v10-10.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1632
---

# Pull Request #1632: Update module github.com/go-playground/validator/v10 to v10.26.0

**Author**: @red-hat-konflux
**Created**: April 06, 2025 at 09:25 AM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-go-playground-validator-v10-10.x`

## Description

This PR contains the following updates:

| Package | Type | Update | Change |
|---|---|---|---|
| [github.com/go-playground/validator/v10](https://redirect.github.com/go-playground/validator) | indirect | minor | `v10.23.0` -> `v10.26.0` |

---

### Release Notes

<details>
<summary>go-playground/validator (github.com/go-playground/validator/v10)</summary>

### [`v10.26.0`](https://redirect.github.com/go-playground/validator/releases/tag/v10.26.0)

[Compare Source](https://redirect.github.com/go-playground/validator/compare/v10.25.0...v10.26.0)

#### What's Changed

-   Use correct pointer in errors.As(). Fix "panic: errors: \*target must be interface or implement error" in examples. by [@&#8203;antonsoroko](https://redirect.github.com/antonsoroko) in [https://github.com/go-playground/validator/pull/1378](https://redirect.github.com/go-playground/validator/pull/1378)
-   Create dependabot by [@&#8203;nodivbyzero](https://redirect.github.com/nodivbyzero) in [https://github.com/go-playground/validator/pull/1373](https://redirect.github.com/go-playground/validator/pull/1373)
-   Bump golangci/golangci-lint-action from 4 to 6 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [https://github.com/go-playground/validator/pull/1381](https://redirect.github.com/go-playground/validator/pull/1381)
-   Bump golang.org/x/text from 0.21.0 to 0.22.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [https://github.com/go-playground/validator/pull/1383](https://redirect.github.com/go-playground/validator/pull/1383)
-   Bump golang.org/x/crypto from 0.32.0 to 0.33.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [https://github.com/go-playground/validator/pull/1382](https://redirect.github.com/go-playground/validator/pull/1382)
-   feat(translations): improve Indonesian translations and add tests by [@&#8203;fathiraz](https://redirect.github.com/fathiraz) in [https://github.com/go-playground/validator/pull/1376](https://redirect.github.com/go-playground/validator/pull/1376)
-   Fix time.Duration translation error by [@&#8203;nodivbyzero](https://redirect.github.com/nodivbyzero) in [https://github.com/go-playground/validator/pull/1154](https://redirect.github.com/go-playground/validator/pull/1154)
-   Update Project Status button by [@&#8203;nodivbyzero](https://redirect.github.com/nodivbyzero) in [https://github.com/go-playground/validator/pull/1380](https://redirect.github.com/go-playground/validator/pull/1380)
-   Remove gitter.im link from README.md by [@&#8203;nodivbyzero](https://redirect.github.com/nodivbyzero) in [https://github.com/go-playground/validator/pull/1366](https://redirect.github.com/go-playground/validator/pull/1366)
-   Docs: fix `Base64RawURL` usage by [@&#8203;196Ikuchil](https://redirect.github.com/196Ikuchil) in [https://github.com/go-playground/validator/pull/1336](https://redirect.github.com/go-playground/validator/pull/1336)
-   Fix length check on dns_rfc1035\_label tag by [@&#8203;KimNorgaard](https://redirect.github.com/KimNorgaard) in [https://github.com/go-playground/validator/pull/1214](https://redirect.github.com/go-playground/validator/pull/1214)
-   Add Korean by [@&#8203;jkh0kr](https://redirect.github.com/jkh0kr) in [https://github.com/go-playground/validator/pull/1338](https://redirect.github.com/go-playground/validator/pull/1338)
-   add german translations by [@&#8203;max-weis](https://redirect.github.com/max-weis) in [https://github.com/go-playground/validator/pull/1322](https://redirect.github.com/go-playground/validator/pull/1322)
-   Update workflow to support the last three Go versions by [@&#8203;nodivbyzero](https://redirect.github.com/nodivbyzero) in [https://github.com/go-playground/validator/pull/1393](https://redirect.github.com/go-playground/validator/pull/1393)
-   Fix: Nil pointer dereference in Arabic translations by [@&#8203;BlackSud0](https://redirect.github.com/BlackSud0) in [https://github.com/go-playground/validator/pull/1391](https://redirect.github.com/go-playground/validator/pull/1391)
-   Translate to thai by [@&#8203;maetad](https://redirect.github.com/maetad) in [https://github.com/go-playground/validator/pull/1202](https://redirect.github.com/go-playground/validator/pull/1202)
-   Feat: add EIN validation by [@&#8203;henrriusdev](https://redirect.github.com/henrriusdev) in [https://github.com/go-playground/validator/pull/1384](https://redirect.github.com/go-playground/validator/pull/1384)
-   Fix reference to parameter name in docs by [@&#8203;yegvla](https://redirect.github.com/yegvla) in [https://github.com/go-playground/validator/pull/1400](https://redirect.github.com/go-playground/validator/pull/1400)
-   use mail.ParseAddress to cover missing email validations by [@&#8203;eladb2011](https://redirect.github.com/eladb2011) in [https://github.com/go-playground/validator/pull/1395](https://redirect.github.com/go-playground/validator/pull/1395)
-   Update linter to v2.0.2 by [@&#8203;nodivbyzero](https://redirect.github.com/nodivbyzero) in [https://github.com/go-playground/validator/pull/1405](https://redirect.github.com/go-playground/validator/pull/1405)

#### New Contributors

-   [@&#8203;antonsoroko](https://redirect.github.com/antonsoroko) made their first contribution in [https://github.com/go-playground/validator/pull/1378](https://redirect.github.com/go-playground/validator/pull/1378)
-   [@&#8203;dependabot](https://redirect.github.com/dependabot) made their first contribution in [https://github.com/go-playground/validator/pull/1381](https://redirect.github.com/go-playground/validator/pull/1381)
-   [@&#8203;fathiraz](https://redirect.github.com/fathiraz) made their first contribution in [https://github.com/go-playground/validator/pull/1376](https://redirect.github.com/go-playground/validator/pull/1376)
-   [@&#8203;196Ikuchil](https://redirect.github.com/196Ikuchil) made their first contribution in [https://github.com/go-playground/validator/pull/1336](https://redirect.github.com/go-playground/validator/pull/1336)
-   [@&#8203;KimNorgaard](https://redirect.github.com/KimNorgaard) made their first contribution in [https://github.com/go-playground/validator/pull/1214](https://redirect.github.com/go-playground/validator/pull/1214)
-   [@&#8203;jkh0kr](https://redirect.github.com/jkh0kr) made their first contribution in [https://github.com/go-playground/validator/pull/1338](https://redirect.github.com/go-playground/validator/pull/1338)
-   [@&#8203;max-weis](https://redirect.github.com/max-weis) made their first contribution in [https://github.com/go-playground/validator/pull/1322](https://redirect.github.com/go-playground/validator/pull/1322)
-   [@&#8203;BlackSud0](https://redirect.github.com/BlackSud0) made their first contribution in [https://github.com/go-playground/validator/pull/1391](https://redirect.github.com/go-playground/validator/pull/1391)
-   [@&#8203;maetad](https://redirect.github.com/maetad) made their first contribution in [https://github.com/go-playground/validator/pull/1202](https://redirect.github.com/go-playground/validator/pull/1202)
-   [@&#8203;henrriusdev](https://redirect.github.com/henrriusdev) made their first contribution in [https://github.com/go-playground/validator/pull/1384](https://redirect.github.com/go-playground/validator/pull/1384)
-   [@&#8203;yegvla](https://redirect.github.com/yegvla) made their first contribution in [https://github.com/go-playground/validator/pull/1400](https://redirect.github.com/go-playground/validator/pull/1400)
-   [@&#8203;eladb2011](https://redirect.github.com/eladb2011) made their first contribution in [https://github.com/go-playground/validator/pull/1395](https://redirect.github.com/go-playground/validator/pull/1395)

**Full Changelog**: https://github.com/go-playground/validator/compare/v10.25.0...v10.26.0

### [`v10.25.0`](https://redirect.github.com/go-playground/validator/releases/tag/v10.25.0): Release 10.25.0

[Compare Source](https://redirect.github.com/go-playground/validator/compare/v10.24.0...v10.25.0)

#### What's Changed

-   Fix postcode_iso3166\_alpha2\_field validation by [@&#8203;ddevcap](https://redirect.github.com/ddevcap) in [https://github.com/go-playground/validator/pull/1359](https://redirect.github.com/go-playground/validator/pull/1359)
-   Update README to replace the Travis CI badge with a GitHub Actions badge by [@&#8203;nodivbyzero](https://redirect.github.com/nodivbyzero) in [https://github.com/go-playground/validator/pull/1362](https://redirect.github.com/go-playground/validator/pull/1362)
-   chore: using errors.As instead of type assertion by [@&#8203;fatelei](https://redirect.github.com/fatelei) in [https://github.com/go-playground/validator/pull/1346](https://redirect.github.com/go-playground/validator/pull/1346)
-   Fix/remove issue template md by [@&#8203;ganeshdipdumbare](https://redirect.github.com/ganeshdipdumbare) in [https://github.com/go-playground/validator/pull/1375](https://redirect.github.com/go-playground/validator/pull/1375)
-   feat: Add support for omitting empty and zero values in validation (including nil pointer and empty content of pointer) by [@&#8203;zeewell](https://redirect.github.com/zeewell) in [https://github.com/go-playground/validator/pull/1289](https://redirect.github.com/go-playground/validator/pull/1289)

#### New Contributors

-   [@&#8203;ddevcap](https://redirect.github.com/ddevcap) made their first contribution in [https://github.com/go-playground/validator/pull/1359](https://redirect.github.com/go-playground/validator/pull/1359)
-   [@&#8203;fatelei](https://redirect.github.com/fatelei) made their first contribution in [https://github.com/go-playground/validator/pull/1346](https://redirect.github.com/go-playground/validator/pull/1346)
-   [@&#8203;zeewell](https://redirect.github.com/zeewell) made their first contribution in [https://github.com/go-playground/validator/pull/1289](https://redirect.github.com/go-playground/validator/pull/1289)

**Full Changelog**: https://github.com/go-playground/validator/compare/v10.24.0...v10.25.0

### [`v10.24.0`](https://redirect.github.com/go-playground/validator/releases/tag/v10.24.0): Release 10.24.0

[Compare Source](https://redirect.github.com/go-playground/validator/compare/v10.23.0...v10.24.0)

#### What's Changed

-   MSGV additions by [@&#8203;deankarn](https://redirect.github.com/deankarn) in [https://github.com/go-playground/validator/pull/1361](https://redirect.github.com/go-playground/validator/pull/1361)

The MSGV(Minimum Supported Go Version) has been bumped to v1.20 to address a security issues in the gaoling `net` package.

**Full Changelog**: https://github.com/go-playground/validator/compare/v10.23.0...v10.24.0

</details>

---

### Configuration

📅 **Schedule**: Branch creation - "after 5am on sunday" in timezone Europe/Prague, Automerge - At any time (no schedule defined).

🚦 **Automerge**: Disabled by config. Please merge this manually once you are satisfied.

♻ **Rebasing**: Whenever PR becomes conflicted, or you tick the rebase/retry checkbox.

🔕 **Ignore**: Close this PR and you won't be reminded about this update again.

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`.

This PR has been generated by [MintMaker](https://redirect.github.com/konflux-ci/mintmaker) (powered by [Renovate Bot](https://redirect.github.com/renovatebot/renovate)).
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiIzOS4xNTguMC1ycG0iLCJ1cGRhdGVkSW5WZXIiOiIzOS4xNTguMC1ycG0iLCJ0YXJnZXRCcmFuY2giOiJtYXN0ZXIiLCJsYWJlbHMiOltdfQ==-->


---

## Discussion

### Comment by @jira-linking on April 06, 2025 at 09:25 AM UTC

Commits missing Jira IDs:
d19ab0cdc025bf41fc759897189c87a7c07b1237
227f369249f395d4fb19f7d17d006ed7266c0173
089e27078540e34ebf3f1db579a1bcc40e35c10d
30cf9b14465721493a25907ac13c31dfb0b20af9


### Comment by @codecov-commenter on April 06, 2025 at 09:30 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1632?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 58.16%. Comparing base [(`dde5824`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/dde582420b242db97254d841074ded8e0ba5d6eb?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`30cf9b1`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/30cf9b14465721493a25907ac13c31dfb0b20af9?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Report is 2 commits behind head on master.

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1632   +/-   ##
=======================================
  Coverage   58.16%   58.16%           
=======================================
  Files         146      146           
  Lines       11333    11333           
=======================================
  Hits         6592     6592           
  Misses       4159     4159           
  Partials      582      582           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1632/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1632/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `58.16% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1632?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

### Comment by @red-hat-konflux on April 14, 2025 at 09:30 AM UTC

### Edited/Blocked Notification

Renovate will not automatically rebase this PR, because it does not recognize the last commit author and assumes somebody else may have edited the PR.

You can manually request rebase by checking the rebase/retry box above.

 ⚠️ **Warning**: custom changes will be lost.

### Comment by @MichaelMraka on April 23, 2025 at 08:20 AM UTC

/retest

### Comment by @psegedy on April 23, 2025 at 09:06 AM UTC

already updated in master branch

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1632*
