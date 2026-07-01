---
type: pull_request
number: 1580
title: "fix(deps): update module github.com/getkin/kin-openapi to v0.129.0"
state: merged
author: red-hat-konflux
created: 2025-02-16T16:48:12Z
updated: 2025-02-24T12:30:33Z
closed: 2025-02-24T12:30:33Z
merged: 2025-02-24T12:30:33Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-getkin-kin-openapi-0.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1580
---

# Pull Request #1580: fix(deps): update module github.com/getkin/kin-openapi to v0.129.0

**Author**: @red-hat-konflux
**Created**: February 16, 2025 at 04:48 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-getkin-kin-openapi-0.x`

## Description

This PR contains the following updates:

| Package | Type | Update | Change |
|---|---|---|---|
| [github.com/getkin/kin-openapi](https://redirect.github.com/getkin/kin-openapi) | require | minor | `v0.128.0` -> `v0.129.0` |

---

### Release Notes

<details>
<summary>getkin/kin-openapi (github.com/getkin/kin-openapi)</summary>

### [`v0.129.0`](https://redirect.github.com/getkin/kin-openapi/releases/tag/v0.129.0)

[Compare Source](https://redirect.github.com/getkin/kin-openapi/compare/v0.128.0...v0.129.0)

#### What's Changed

-   README: add Fuego to dependents by [@&#8203;EwenQuim](https://redirect.github.com/EwenQuim) in [https://github.com/getkin/kin-openapi/pull/1017](https://redirect.github.com/getkin/kin-openapi/pull/1017)
-   openapi3: skip a test in CI to avoid 403s from some remote server by [@&#8203;fenollp](https://redirect.github.com/fenollp) in [https://github.com/getkin/kin-openapi/pull/1019](https://redirect.github.com/getkin/kin-openapi/pull/1019)
-   openapi3: introduce StringMap type to enable unmarshalling of maps with Origin by [@&#8203;reuvenharrison](https://redirect.github.com/reuvenharrison) in [https://github.com/getkin/kin-openapi/pull/1018](https://redirect.github.com/getkin/kin-openapi/pull/1018)
-   openapi3: reference originating locations in YAML specs - step 1 by [@&#8203;reuvenharrison](https://redirect.github.com/reuvenharrison) in [https://github.com/getkin/kin-openapi/pull/1007](https://redirect.github.com/getkin/kin-openapi/pull/1007)
-   openapi3: reference originating locations in YAML specs - step 2 by [@&#8203;reuvenharrison](https://redirect.github.com/reuvenharrison) in [https://github.com/getkin/kin-openapi/pull/1024](https://redirect.github.com/getkin/kin-openapi/pull/1024)
-   openapi3: process discriminator mapping values as refs by [@&#8203;jgresty](https://redirect.github.com/jgresty) in [https://github.com/getkin/kin-openapi/pull/1022](https://redirect.github.com/getkin/kin-openapi/pull/1022)
-   openapi3filter: register decoder for other JSON content types by [@&#8203;oliverli](https://redirect.github.com/oliverli) in [https://github.com/getkin/kin-openapi/pull/1026](https://redirect.github.com/getkin/kin-openapi/pull/1026)
-   Revert "openapi3: process discriminator mapping values as refs" by [@&#8203;fenollp](https://redirect.github.com/fenollp) in [https://github.com/getkin/kin-openapi/pull/1029](https://redirect.github.com/getkin/kin-openapi/pull/1029)
-   openapi3: fail to load spec because of schema names in mapping  by [@&#8203;reuvenharrison](https://redirect.github.com/reuvenharrison) in [https://github.com/getkin/kin-openapi/pull/1027](https://redirect.github.com/getkin/kin-openapi/pull/1027)
-   openapi2conv: convert schemaRef for additional props by [@&#8203;jayanth-tatina-groww](https://redirect.github.com/jayanth-tatina-groww) in [https://github.com/getkin/kin-openapi/pull/1030](https://redirect.github.com/getkin/kin-openapi/pull/1030)
-   openapi3: simplify by replacing math.Min with min by [@&#8203;alexandear](https://redirect.github.com/alexandear) in [https://github.com/getkin/kin-openapi/pull/1032](https://redirect.github.com/getkin/kin-openapi/pull/1032)
-   openapi3: fix deprecation comments by [@&#8203;alexandear](https://redirect.github.com/alexandear) in [https://github.com/getkin/kin-openapi/pull/1034](https://redirect.github.com/getkin/kin-openapi/pull/1034)
-   test: fix expected-actual parameters in require.Equal by [@&#8203;alexandear](https://redirect.github.com/alexandear) in [https://github.com/getkin/kin-openapi/pull/1035](https://redirect.github.com/getkin/kin-openapi/pull/1035)
-   use forked yaml modules without "replace" by [@&#8203;reuvenharrison](https://redirect.github.com/reuvenharrison) in [https://github.com/getkin/kin-openapi/pull/1038](https://redirect.github.com/getkin/kin-openapi/pull/1038)
-   openapi3: update date schema formats to not match months or days of '00' by [@&#8203;RulerOfTheQueendom](https://redirect.github.com/RulerOfTheQueendom) in [https://github.com/getkin/kin-openapi/pull/1042](https://redirect.github.com/getkin/kin-openapi/pull/1042)
-   openapi3,openapi3filter: replace interface{} with any by [@&#8203;alexandear](https://redirect.github.com/alexandear) in [https://github.com/getkin/kin-openapi/pull/1040](https://redirect.github.com/getkin/kin-openapi/pull/1040)
-   openapi3filter: Simplify ValidateRequest implementation by [@&#8203;alexandear](https://redirect.github.com/alexandear) in [https://github.com/getkin/kin-openapi/pull/1041](https://redirect.github.com/getkin/kin-openapi/pull/1041)
-   openapi3filter: validation of `x-www-form-urlencoded` with arbitrary nested allOf by [@&#8203;mikhalytch](https://redirect.github.com/mikhalytch) in [https://github.com/getkin/kin-openapi/pull/1046](https://redirect.github.com/getkin/kin-openapi/pull/1046)
-   openapi2conv: convert references in nested additionalProperties schemas by [@&#8203;travisnewhouse](https://redirect.github.com/travisnewhouse) in [https://github.com/getkin/kin-openapi/pull/1047](https://redirect.github.com/getkin/kin-openapi/pull/1047)

#### New Contributors

-   [@&#8203;EwenQuim](https://redirect.github.com/EwenQuim) made their first contribution in [https://github.com/getkin/kin-openapi/pull/1017](https://redirect.github.com/getkin/kin-openapi/pull/1017)
-   [@&#8203;jgresty](https://redirect.github.com/jgresty) made their first contribution in [https://github.com/getkin/kin-openapi/pull/1022](https://redirect.github.com/getkin/kin-openapi/pull/1022)
-   [@&#8203;oliverli](https://redirect.github.com/oliverli) made their first contribution in [https://github.com/getkin/kin-openapi/pull/1026](https://redirect.github.com/getkin/kin-openapi/pull/1026)
-   [@&#8203;jayanth-tatina-groww](https://redirect.github.com/jayanth-tatina-groww) made their first contribution in [https://github.com/getkin/kin-openapi/pull/1030](https://redirect.github.com/getkin/kin-openapi/pull/1030)
-   [@&#8203;RulerOfTheQueendom](https://redirect.github.com/RulerOfTheQueendom) made their first contribution in [https://github.com/getkin/kin-openapi/pull/1042](https://redirect.github.com/getkin/kin-openapi/pull/1042)
-   [@&#8203;mikhalytch](https://redirect.github.com/mikhalytch) made their first contribution in [https://github.com/getkin/kin-openapi/pull/1046](https://redirect.github.com/getkin/kin-openapi/pull/1046)
-   [@&#8203;travisnewhouse](https://redirect.github.com/travisnewhouse) made their first contribution in [https://github.com/getkin/kin-openapi/pull/1047](https://redirect.github.com/getkin/kin-openapi/pull/1047)

**Full Changelog**: https://github.com/getkin/kin-openapi/compare/v0.128.0...v0.129.0

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

### Comment by @jira-linking on February 16, 2025 at 04:48 PM UTC

Commits missing Jira IDs:
7cffff159d258009ae36e220330fb650f68d8f1f


### Comment by @codecov-commenter on February 16, 2025 at 04:52 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1580?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 58.00%. Comparing base [(`5d8dd09`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/5d8dd09e1d5d90445df8563b818ad50c95def18d?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`7cffff1`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/7cffff159d258009ae36e220330fb650f68d8f1f?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1580   +/-   ##
=======================================
  Coverage   58.00%   58.00%           
=======================================
  Files         143      143           
  Lines       11165    11165           
=======================================
  Hits         6476     6476           
  Misses       4107     4107           
  Partials      582      582           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1580/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1580/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `58.00% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1580?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1580*
