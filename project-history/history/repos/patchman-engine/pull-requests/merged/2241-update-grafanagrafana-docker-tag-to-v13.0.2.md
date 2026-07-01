---
type: pull_request
number: 2241
title: "Update grafana/grafana Docker tag to v13.0.2"
state: merged
author: red-hat-konflux
created: 2026-06-22T05:24:00Z
updated: 2026-06-22T05:47:03Z
closed: 2026-06-22T05:24:13Z
merged: 2026-06-22T05:24:13Z
base_branch: master
head_branch: konflux/mintmaker/master/grafana-monorepo
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/2241
---

# Pull Request #2241: Update grafana/grafana Docker tag to v13.0.2

**Author**: @red-hat-konflux
**Created**: June 22, 2026 at 05:24 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/grafana-monorepo`

## Description

This PR contains the following updates:

| Package | Type | Update | Change |
|---|---|---|---|
| [grafana/grafana](https://redirect.github.com/grafana/grafana) | final | patch | `13.0.1` → `13.0.2` |

---

### Release Notes

<details>
<summary>grafana/grafana (grafana/grafana)</summary>

### [`v13.0.2`](https://redirect.github.com/grafana/grafana/releases/tag/v13.0.2): 13.0.2

[Compare Source](https://redirect.github.com/grafana/grafana/compare/v13.0.1...v13.0.2)

[Download page](https://grafana.com/grafana/download/13.0.2)
[What's new highlights](https://grafana.com/docs/grafana/latest/whatsnew/)

##### Features and enhancements

- **Dashboards:** Show k8s format in provisioned save [#&#8203;123045](https://redirect.github.com/grafana/grafana/pull/123045), [@&#8203;stephaniehingtgen](https://redirect.github.com/stephaniehingtgen)
- **Docker:** Bump Alpine-based images to 3.23.4 [#&#8203;122938](https://redirect.github.com/grafana/grafana/pull/122938), [@&#8203;Proximyst](https://redirect.github.com/Proximyst)
- **Go:** Update version to 1.26.3 [#&#8203;124454](https://redirect.github.com/grafana/grafana/pull/124454), [@&#8203;macabu](https://redirect.github.com/macabu)
- **Homepage:** Support v2 dashboards if defined by a file [#&#8203;123029](https://redirect.github.com/grafana/grafana/pull/123029), [@&#8203;stephaniehingtgen](https://redirect.github.com/stephaniehingtgen)
- **LibraryPanels:** Return 403 instead of 500 for insufficient permissions [#&#8203;123467](https://redirect.github.com/grafana/grafana/pull/123467), [@&#8203;MissingRoberto](https://redirect.github.com/MissingRoberto)
- **Provisioning:** Don't mark folders pending due to \_folder.json metadata [#&#8203;124139](https://redirect.github.com/grafana/grafana/pull/124139), [@&#8203;MissingRoberto](https://redirect.github.com/MissingRoberto)
- **Provisioning:** Honor ruleset bypass for write workflow validation [#&#8203;124128](https://redirect.github.com/grafana/grafana/pull/124128), [@&#8203;MissingRoberto](https://redirect.github.com/MissingRoberto)
- **Provisioning:** Negotiate receive-pack capabilities for git pushes [#&#8203;124130](https://redirect.github.com/grafana/grafana/pull/124130), [@&#8203;MissingRoberto](https://redirect.github.com/MissingRoberto)
- **Provisioning:** Per-verb fallback for the files subresource [#&#8203;123900](https://redirect.github.com/grafana/grafana/pull/123900), [@&#8203;MissingRoberto](https://redirect.github.com/MissingRoberto)
- **Provisioning:** Scope repository uniqueness by (URL, branch, path) [#&#8203;124121](https://redirect.github.com/grafana/grafana/pull/124121), [@&#8203;ferruvich](https://redirect.github.com/ferruvich)
- **Provisioning:** Surface folder uid-too-long and other validation 4xx as sync warnings [#&#8203;123888](https://redirect.github.com/grafana/grafana/pull/123888), [@&#8203;MissingRoberto](https://redirect.github.com/MissingRoberto)
- **Provisioning:** add public\_root\_url instance setting for external URLs [#&#8203;124258](https://redirect.github.com/grafana/grafana/pull/124258), [@&#8203;MissingRoberto](https://redirect.github.com/MissingRoberto)

##### Bug fixes

- **DashboardDS:** Fix Mixed panels not updating on time-range change with stale upstreams [#&#8203;124894](https://redirect.github.com/grafana/grafana/pull/124894), [@&#8203;ivanortegaalba](https://redirect.github.com/ivanortegaalba)
- **Jaeger:** Fix log event timestamp unit conversion in trace view [#&#8203;123707](https://redirect.github.com/grafana/grafana/pull/123707), [@&#8203;ktw4071](https://redirect.github.com/ktw4071)
- **K8s Dashboards:** Fix folder permission check to use dashboards:create [#&#8203;124942](https://redirect.github.com/grafana/grafana/pull/124942), [@&#8203;mihai-turdean](https://redirect.github.com/mihai-turdean)
- **PostgreSQL:** Allow sql\_engine to return results for EXPLAIN queries [#&#8203;123246](https://redirect.github.com/grafana/grafana/pull/123246), [@&#8203;sdague](https://redirect.github.com/sdague)
- **Provisioning:** Bump nanogit to v0.17.0 to fix pushes with repositories using git modules [#&#8203;124140](https://redirect.github.com/grafana/grafana/pull/124140), [@&#8203;MissingRoberto](https://redirect.github.com/MissingRoberto)
- **RBAC:** Quick fix for global datasource permissions (Enterprise)
- **Security**: CVE-2026-9029
- **Security**: CVE-2026-33382
- **Security**: CVE-2026-42127
- **Security**: CVE-2026-42129
- **Security**: CVE-2026-10601
- **Security**: CVE-2026-8609
- **Security**: CVE-2026-8595

</details>

---

### Configuration

📅 **Schedule**: Branch creation - Between 03:00 AM and 10:59 AM, only on Monday ( * 3-10 * * 1 ) in timezone Europe/Prague, Automerge - At any time (no schedule defined).

🚦 **Automerge**: Enabled.

♻ **Rebasing**: Whenever PR is behind base branch, or you tick the rebase/retry checkbox.

🔕 **Ignore**: Close this PR and you won't be reminded about this update again.

---

 - [ ] <!-- rebase-check -->If you want to rebase/retry this PR, check this box

---

To execute skipped test pipelines write comment `/ok-to-test`.

---
### Documentation

Find out how to configure dependency updates in [MintMaker documentation](https://konflux-ci.dev/docs/mintmaker/user/) or see all available configuration options in [Renovate documentation](https://docs.renovatebot.com/configuration-options/).
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiI0Mi45OS4wLXJwbSIsInVwZGF0ZWRJblZlciI6IjQyLjk5LjAtcnBtIiwidGFyZ2V0QnJhbmNoIjoibWFzdGVyIiwibGFiZWxzIjpbXX0=-->


---

## Discussion

### Comment by @codecov-commenter on June 22, 2026 at 05:29 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2241?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 59.06%. Comparing base ([`95cf64c`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/95cf64ca476e3464850f0deb1dd7f41e0f0ee2ce?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`834dc81`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/834dc811c9be6cd758c3370752db41c6e24e5603?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 4 commits behind head on master.

<details><summary>Additional details and impacted files</summary>



```diff
@@           Coverage Diff           @@
##           master    #2241   +/-   ##
=======================================
  Coverage   59.06%   59.06%           
=======================================
  Files         138      138           
  Lines        8848     8848           
=======================================
  Hits         5226     5226           
  Misses       3076     3076           
  Partials      546      546           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2241/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2241/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `59.06% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Harness](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/2241?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/2241*
