---
type: pull_request
number: 1811
title: "Update module github.com/getkin/kin-openapi to v0.133.0"
state: merged
author: red-hat-konflux
created: 2025-08-31T12:29:43Z
updated: 2025-10-08T16:16:10Z
closed: 2025-08-31T16:48:20Z
merged: 2025-08-31T16:48:20Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-getkin-kin-openapi-0.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1811
---

# Pull Request #1811: Update module github.com/getkin/kin-openapi to v0.133.0

**Author**: @red-hat-konflux
**Created**: August 31, 2025 at 12:29 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-getkin-kin-openapi-0.x`

## Description

This PR contains the following updates:

| Package | Change | Age | Confidence |
|---|---|---|---|
| [github.com/getkin/kin-openapi](https://redirect.github.com/getkin/kin-openapi) | `v0.132.0` -> `v0.133.0` | [![age](https://developer.mend.io/api/mc/badges/age/go/github.com%2fgetkin%2fkin-openapi/v0.133.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) | [![confidence](https://developer.mend.io/api/mc/badges/confidence/go/github.com%2fgetkin%2fkin-openapi/v0.132.0/v0.133.0?slim=true)](https://docs.renovatebot.com/merge-confidence/) |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>getkin/kin-openapi (github.com/getkin/kin-openapi)</summary>

### [`v0.133.0`](https://redirect.github.com/getkin/kin-openapi/releases/tag/v0.133.0)

[Compare Source](https://redirect.github.com/getkin/kin-openapi/compare/v0.132.0...v0.133.0)

#### What's Changed

- openapi3: resolve Snyk security warning with path traversal by [@&#8203;seborama](https://redirect.github.com/seborama) in [https://github.com/getkin/kin-openapi/pull/1066](https://redirect.github.com/getkin/kin-openapi/pull/1066)
- openapi3: replace bigfloat with decimal128 to fix rounding errors during validation by [@&#8203;Revolyssup](https://redirect.github.com/Revolyssup) in [https://github.com/getkin/kin-openapi/pull/1068](https://redirect.github.com/getkin/kin-openapi/pull/1068)
- openapi2conv: Preserve externalDocs on operations during conversion by [@&#8203;hwustrack](https://redirect.github.com/hwustrack) in [https://github.com/getkin/kin-openapi/pull/1070](https://redirect.github.com/getkin/kin-openapi/pull/1070)
- openapi3: fix ineffectual caching of compiled regexps by [@&#8203;philpearl](https://redirect.github.com/philpearl) in [https://github.com/getkin/kin-openapi/pull/1076](https://redirect.github.com/getkin/kin-openapi/pull/1076)
- openapi3: use Ptr instead of BoolPtr,Float64Ptr,Int64Ptr,Uint64Ptr by [@&#8203;alexandear](https://redirect.github.com/alexandear) in [https://github.com/getkin/kin-openapi/pull/1033](https://redirect.github.com/getkin/kin-openapi/pull/1033)
- openapi3: resolve refs in parameter examples by [@&#8203;reuvenharrison](https://redirect.github.com/reuvenharrison) in [https://github.com/getkin/kin-openapi/pull/1086](https://redirect.github.com/getkin/kin-openapi/pull/1086)
- openapifilter: Add support for RFC 7396 application/merge-patch+json by [@&#8203;byted](https://redirect.github.com/byted) in [https://github.com/getkin/kin-openapi/pull/1084](https://redirect.github.com/getkin/kin-openapi/pull/1084)
- openapi3filter: use FileBodyDecoder if the format is specified as binary by [@&#8203;dbarrosop](https://redirect.github.com/dbarrosop) in [https://github.com/getkin/kin-openapi/pull/1088](https://redirect.github.com/getkin/kin-openapi/pull/1088)
- openapi3: preserve all validation errors for allOf by [@&#8203;alexbakker](https://redirect.github.com/alexbakker) in [https://github.com/getkin/kin-openapi/pull/1087](https://redirect.github.com/getkin/kin-openapi/pull/1087)
- openapi3filter: support primitive parsing for individual text like parts in multipart/form-data by [@&#8203;nmeheus](https://redirect.github.com/nmeheus) in [https://github.com/getkin/kin-openapi/pull/1090](https://redirect.github.com/getkin/kin-openapi/pull/1090)
- Some coding style fixes and cleaning up by [@&#8203;fenollp](https://redirect.github.com/fenollp) in [https://github.com/getkin/kin-openapi/pull/1093](https://redirect.github.com/getkin/kin-openapi/pull/1093)
- openapi2conv: preserve x-fields when converting from v2 to v3 by [@&#8203;saltbo](https://redirect.github.com/saltbo) in [https://github.com/getkin/kin-openapi/pull/1092](https://redirect.github.com/getkin/kin-openapi/pull/1092)

#### New Contributors

- [@&#8203;seborama](https://redirect.github.com/seborama) made their first contribution in [https://github.com/getkin/kin-openapi/pull/1066](https://redirect.github.com/getkin/kin-openapi/pull/1066)
- [@&#8203;Revolyssup](https://redirect.github.com/Revolyssup) made their first contribution in [https://github.com/getkin/kin-openapi/pull/1068](https://redirect.github.com/getkin/kin-openapi/pull/1068)
- [@&#8203;hwustrack](https://redirect.github.com/hwustrack) made their first contribution in [https://github.com/getkin/kin-openapi/pull/1070](https://redirect.github.com/getkin/kin-openapi/pull/1070)
- [@&#8203;philpearl](https://redirect.github.com/philpearl) made their first contribution in [https://github.com/getkin/kin-openapi/pull/1076](https://redirect.github.com/getkin/kin-openapi/pull/1076)
- [@&#8203;byted](https://redirect.github.com/byted) made their first contribution in [https://github.com/getkin/kin-openapi/pull/1084](https://redirect.github.com/getkin/kin-openapi/pull/1084)
- [@&#8203;dbarrosop](https://redirect.github.com/dbarrosop) made their first contribution in [https://github.com/getkin/kin-openapi/pull/1088](https://redirect.github.com/getkin/kin-openapi/pull/1088)
- [@&#8203;nmeheus](https://redirect.github.com/nmeheus) made their first contribution in [https://github.com/getkin/kin-openapi/pull/1090](https://redirect.github.com/getkin/kin-openapi/pull/1090)
- [@&#8203;saltbo](https://redirect.github.com/saltbo) made their first contribution in [https://github.com/getkin/kin-openapi/pull/1092](https://redirect.github.com/getkin/kin-openapi/pull/1092)

**Full Changelog**: https://github.com/getkin/kin-openapi/compare/v0.132.0...v0.133.0

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
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0MS43LjAtcnBtIiwidXBkYXRlZEluVmVyIjoiNDEuNy4wLXJwbSIsInRhcmdldEJyYW5jaCI6Im1hc3RlciIsImxhYmVscyI6W119-->


---

## Discussion

### Comment by @jira-linking on August 31, 2025 at 12:29 PM UTC

Commits missing Jira IDs:
532635a491a2f37cb6a67e4600726e2f1b69c2ac


### Comment by @codecov-commenter on August 31, 2025 at 12:34 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1811?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 54.85%. Comparing base ([`2cc4469`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/2cc4469ac891ed9a08dc3306cb02415cd0b9694a?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`532635a`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/532635a491a2f37cb6a67e4600726e2f1b69c2ac?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #1811   +/-   ##
=======================================
  Coverage   54.85%   54.85%           
=======================================
  Files         140      140           
  Lines       10875    10875           
=======================================
  Hits         5966     5966           
  Misses       4373     4373           
  Partials      536      536           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1811/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1811/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `54.85% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1811?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1811*
