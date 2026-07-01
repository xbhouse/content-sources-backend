---
type: pull_request
number: 1714
title: "Update module github.com/prometheus/procfs to v0.17.0"
state: merged
author: red-hat-konflux
created: 2025-07-06T07:10:35Z
updated: 2026-04-05T14:12:19Z
closed: 2025-07-06T19:03:48Z
merged: 2025-07-06T19:03:48Z
base_branch: security-compliance
head_branch: konflux/mintmaker/security-compliance/github.com-prometheus-procfs-0.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1714
---

# Pull Request #1714: Update module github.com/prometheus/procfs to v0.17.0

**Author**: @red-hat-konflux
**Created**: July 06, 2025 at 07:10 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `security-compliance` ← **Head**: `konflux/mintmaker/security-compliance/github.com-prometheus-procfs-0.x`

## Description

This PR contains the following updates:

| Package | Type | Update | Change |
|---|---|---|---|
| [github.com/prometheus/procfs](https://redirect.github.com/prometheus/procfs) | indirect | minor | `v0.16.1` -> `v0.17.0` |

---

> [!WARNING]
> Some dependencies could not be looked up. Check the warning logs for more information.

---

### Release Notes

<details>
<summary>prometheus/procfs (github.com/prometheus/procfs)</summary>

### [`v0.17.0`](https://redirect.github.com/prometheus/procfs/releases/tag/v0.17.0)

[Compare Source](https://redirect.github.com/prometheus/procfs/compare/v0.16.1...v0.17.0)

#### What's Changed

-   Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [https://github.com/prometheus/procfs/pull/718](https://redirect.github.com/prometheus/procfs/pull/718)
-   Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [https://github.com/prometheus/procfs/pull/721](https://redirect.github.com/prometheus/procfs/pull/721)
-   btrfs: correct allocation ratios for raid1c\[34] by [@&#8203;SimSaladin](https://redirect.github.com/SimSaladin) in [https://github.com/prometheus/procfs/pull/722](https://redirect.github.com/prometheus/procfs/pull/722)
-   build(deps): bump golang.org/x/sync from 0.13.0 to 0.14.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [https://github.com/prometheus/procfs/pull/724](https://redirect.github.com/prometheus/procfs/pull/724)
-   build(deps): bump golang.org/x/sys from 0.32.0 to 0.33.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [https://github.com/prometheus/procfs/pull/723](https://redirect.github.com/prometheus/procfs/pull/723)
-   Supports collection of process shared memory by [@&#8203;SilenceAdele](https://redirect.github.com/SilenceAdele) in [https://github.com/prometheus/procfs/pull/719](https://redirect.github.com/prometheus/procfs/pull/719)
-   build(deps): bump golang.org/x/sync from 0.14.0 to 0.15.0 by [@&#8203;dependabot](https://redirect.github.com/dependabot) in [https://github.com/prometheus/procfs/pull/732](https://redirect.github.com/prometheus/procfs/pull/732)
-   nvme: Add ControllerID output by [@&#8203;ShashwatHiregoudar](https://redirect.github.com/ShashwatHiregoudar) in [https://github.com/prometheus/procfs/pull/731](https://redirect.github.com/prometheus/procfs/pull/731)
-   Synchronize common files from prometheus/prometheus by [@&#8203;prombot](https://redirect.github.com/prombot) in [https://github.com/prometheus/procfs/pull/727](https://redirect.github.com/prometheus/procfs/pull/727)
-   sysfs: Add support to collect link status for  PCIe devices by [@&#8203;naoki9911](https://redirect.github.com/naoki9911) in [https://github.com/prometheus/procfs/pull/728](https://redirect.github.com/prometheus/procfs/pull/728)
-   nfs/parse.go: fix ClientV4Stats' GetDeviceInfo/LayoutGet -- values were swapped by [@&#8203;johnleslie](https://redirect.github.com/johnleslie) in [https://github.com/prometheus/procfs/pull/726](https://redirect.github.com/prometheus/procfs/pull/726)
-   Fix linting issue by [@&#8203;SuperQ](https://redirect.github.com/SuperQ) in [https://github.com/prometheus/procfs/pull/733](https://redirect.github.com/prometheus/procfs/pull/733)
-   feat(mdstat): recognize reshape status by [@&#8203;tamcore](https://redirect.github.com/tamcore) in [https://github.com/prometheus/procfs/pull/679](https://redirect.github.com/prometheus/procfs/pull/679)
-   Nvidia/Mellanox expose ROCE ECN information on sysfs on the path by [@&#8203;dasturiasArista](https://redirect.github.com/dasturiasArista) in [https://github.com/prometheus/procfs/pull/695](https://redirect.github.com/prometheus/procfs/pull/695)
-   added zswap, zswapped, secpagetables, filehugepages, hugetlb and unaccepted to meminfo by [@&#8203;navidys](https://redirect.github.com/navidys) in [https://github.com/prometheus/procfs/pull/655](https://redirect.github.com/prometheus/procfs/pull/655)
-   Parse StartCode, EndCode, and StartStack in `Proc.Stat()` by [@&#8203;pgimalac](https://redirect.github.com/pgimalac) in [https://github.com/prometheus/procfs/pull/659](https://redirect.github.com/prometheus/procfs/pull/659)
-   Add node_guid to infiniband class by [@&#8203;di3go-sona](https://redirect.github.com/di3go-sona) in [https://github.com/prometheus/procfs/pull/670](https://redirect.github.com/prometheus/procfs/pull/670)
-   Fix linting issues by [@&#8203;SuperQ](https://redirect.github.com/SuperQ) in [https://github.com/prometheus/procfs/pull/734](https://redirect.github.com/prometheus/procfs/pull/734)

#### New Contributors

-   [@&#8203;SimSaladin](https://redirect.github.com/SimSaladin) made their first contribution in [https://github.com/prometheus/procfs/pull/722](https://redirect.github.com/prometheus/procfs/pull/722)
-   [@&#8203;SilenceAdele](https://redirect.github.com/SilenceAdele) made their first contribution in [https://github.com/prometheus/procfs/pull/719](https://redirect.github.com/prometheus/procfs/pull/719)
-   [@&#8203;ShashwatHiregoudar](https://redirect.github.com/ShashwatHiregoudar) made their first contribution in [https://github.com/prometheus/procfs/pull/731](https://redirect.github.com/prometheus/procfs/pull/731)
-   [@&#8203;naoki9911](https://redirect.github.com/naoki9911) made their first contribution in [https://github.com/prometheus/procfs/pull/728](https://redirect.github.com/prometheus/procfs/pull/728)
-   [@&#8203;johnleslie](https://redirect.github.com/johnleslie) made their first contribution in [https://github.com/prometheus/procfs/pull/726](https://redirect.github.com/prometheus/procfs/pull/726)
-   [@&#8203;tamcore](https://redirect.github.com/tamcore) made their first contribution in [https://github.com/prometheus/procfs/pull/679](https://redirect.github.com/prometheus/procfs/pull/679)
-   [@&#8203;navidys](https://redirect.github.com/navidys) made their first contribution in [https://github.com/prometheus/procfs/pull/655](https://redirect.github.com/prometheus/procfs/pull/655)
-   [@&#8203;pgimalac](https://redirect.github.com/pgimalac) made their first contribution in [https://github.com/prometheus/procfs/pull/659](https://redirect.github.com/prometheus/procfs/pull/659)
-   [@&#8203;di3go-sona](https://redirect.github.com/di3go-sona) made their first contribution in [https://github.com/prometheus/procfs/pull/670](https://redirect.github.com/prometheus/procfs/pull/670)

**Full Changelog**: https://github.com/prometheus/procfs/compare/v0.16.1...v0.17.0

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
<!--renovate-debug:eyJjcmVhdGVkSW5WZXIiOiIzOS4yNjQuMC1ycG0iLCJ1cGRhdGVkSW5WZXIiOiIzOS4yNjQuMC1ycG0iLCJ0YXJnZXRCcmFuY2giOiJzZWN1cml0eS1jb21wbGlhbmNlIiwibGFiZWxzIjpbXX0=-->


---

## Discussion

### Comment by @jira-linking on July 06, 2025 at 07:10 AM UTC

Commits missing Jira IDs:
d531912ce120f34797a056539b6a5fe87b931210


### Comment by @sourcery-ai on July 06, 2025 at 07:10 AM UTC

<!-- Generated by sourcery-ai[bot]: start review_guide -->

## Reviewer's Guide

This PR upgrades the indirect Prometheus procfs dependency from v0.16.1 to v0.17.0 by updating the version constraint in go.mod and regenerating module checksums in go.sum.

### File-Level Changes

| Change | Details | Files |
| ------ | ------- | ----- |
| Bump procfs dependency to v0.17.0 | <ul><li>Incremented procfs version in go.mod</li><li>Refreshed checksum entries in go.sum</li></ul> | `go.mod`<br/>`go.sum` |

---

<details>
<summary>Tips and commands</summary>

#### Interacting with Sourcery

- **Trigger a new review:** Comment `@sourcery-ai review` on the pull request.
- **Continue discussions:** Reply directly to Sourcery's review comments.
- **Generate a GitHub issue from a review comment:** Ask Sourcery to create an
  issue from a review comment by replying to it. You can also reply to a
  review comment with `@sourcery-ai issue` to create an issue from it.
- **Generate a pull request title:** Write `@sourcery-ai` anywhere in the pull
  request title to generate a title at any time. You can also comment
  `@sourcery-ai title` on the pull request to (re-)generate the title at any time.
- **Generate a pull request summary:** Write `@sourcery-ai summary` anywhere in
  the pull request body to generate a PR summary at any time exactly where you
  want it. You can also comment `@sourcery-ai summary` on the pull request to
  (re-)generate the summary at any time.
- **Generate reviewer's guide:** Comment `@sourcery-ai guide` on the pull
  request to (re-)generate the reviewer's guide at any time.
- **Resolve all Sourcery comments:** Comment `@sourcery-ai resolve` on the
  pull request to resolve all Sourcery comments. Useful if you've already
  addressed all the comments and don't want to see them anymore.
- **Dismiss all Sourcery reviews:** Comment `@sourcery-ai dismiss` on the pull
  request to dismiss all existing Sourcery reviews. Especially useful if you
  want to start fresh with a new review - don't forget to comment
  `@sourcery-ai review` to trigger a new review!

#### Customizing Your Experience

Access your [dashboard](https://app.sourcery.ai) to:
- Enable or disable review features such as the Sourcery-generated pull request
  summary, the reviewer's guide, and others.
- Change the review language.
- Add, remove or edit custom review instructions.
- Adjust other review settings.

#### Getting Help

- [Contact our support team](mailto:support@sourcery.ai) for questions or feedback.
- Visit our [documentation](https://docs.sourcery.ai) for detailed guides and information.
- Keep in touch with the Sourcery team by following us on [X/Twitter](https://x.com/SourceryAI), [LinkedIn](https://www.linkedin.com/company/sourcery-ai/) or [GitHub](https://github.com/sourcery-ai).

</details>

<!-- Generated by sourcery-ai[bot]: end review_guide -->

### Comment by @codecov-commenter on July 06, 2025 at 07:15 AM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1714?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:white_check_mark: All modified and coverable lines are covered by tests.
:white_check_mark: Project coverage is 57.25%. Comparing base ([`4ed8a04`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/4ed8a04cf04ccf15c1640d67bc2f59dc7ed5c7df?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`d531912`](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/d531912ce120f34797a056539b6a5fe87b931210?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).
:warning: Report is 549 commits behind head on security-compliance.

<details><summary>Additional details and impacted files</summary>



```diff
@@                 Coverage Diff                  @@
##           security-compliance    #1714   +/-   ##
====================================================
  Coverage                57.25%   57.25%           
====================================================
  Files                      138      138           
  Lines                    10776    10776           
====================================================
  Hits                      6170     6170           
  Misses                    4046     4046           
  Partials                   560      560           
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1714/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1714/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `57.25% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1714?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1714*
