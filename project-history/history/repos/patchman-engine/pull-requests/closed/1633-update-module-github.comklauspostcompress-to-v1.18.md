---
type: pull_request
number: 1633
title: "Update module github.com/klauspost/compress to v1.18.0"
state: closed
author: red-hat-konflux
created: 2025-04-06T13:14:07Z
updated: 2025-04-23T12:21:14Z
closed: 2025-04-23T09:03:42Z
base_branch: master
head_branch: konflux/mintmaker/master/github.com-klauspost-compress-1.x
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/1633
---

# Pull Request #1633: Update module github.com/klauspost/compress to v1.18.0

**Author**: @red-hat-konflux
**Created**: April 06, 2025 at 01:14 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `konflux/mintmaker/master/github.com-klauspost-compress-1.x`

## Description

This PR contains the following updates:

| Package | Type | Update | Change |
|---|---|---|---|
| [github.com/klauspost/compress](https://redirect.github.com/klauspost/compress) | indirect | minor | `v1.17.11` -> `v1.18.0` |

---

### Release Notes

<details>
<summary>klauspost/compress (github.com/klauspost/compress)</summary>

### [`v1.18.0`](https://redirect.github.com/klauspost/compress/releases/tag/v1.18.0)

[Compare Source](https://redirect.github.com/klauspost/compress/compare/v1.17.11...v1.18.0)

#### What's Changed

-   Deprecate Go 1.21 and add 1.24 by [@&#8203;klauspost](https://redirect.github.com/klauspost) in [https://github.com/klauspost/compress/pull/1055](https://redirect.github.com/klauspost/compress/pull/1055)
-   Add unsafe little endian loaders by [@&#8203;klauspost](https://redirect.github.com/klauspost) in [https://github.com/klauspost/compress/pull/1036](https://redirect.github.com/klauspost/compress/pull/1036)
-   fix: check `r.err != nil` but return a nil value error `err` by [@&#8203;alingse](https://redirect.github.com/alingse) in [https://github.com/klauspost/compress/pull/1028](https://redirect.github.com/klauspost/compress/pull/1028)
-   refactor: use built-in `min` function by [@&#8203;Juneezee](https://redirect.github.com/Juneezee) in [https://github.com/klauspost/compress/pull/1038](https://redirect.github.com/klauspost/compress/pull/1038)
-   zstd: use `slices.Max` for max value in slice by [@&#8203;Juneezee](https://redirect.github.com/Juneezee) in [https://github.com/klauspost/compress/pull/1041](https://redirect.github.com/klauspost/compress/pull/1041)
-   flate: Simplify L4-6 loading by [@&#8203;klauspost](https://redirect.github.com/klauspost) in [https://github.com/klauspost/compress/pull/1043](https://redirect.github.com/klauspost/compress/pull/1043)
-   flate: Simplify matchlen (remove asm) by [@&#8203;klauspost](https://redirect.github.com/klauspost) in [https://github.com/klauspost/compress/pull/1045](https://redirect.github.com/klauspost/compress/pull/1045)
-   s2: Add block decode fuzzer by [@&#8203;klauspost](https://redirect.github.com/klauspost) in [https://github.com/klauspost/compress/pull/1044](https://redirect.github.com/klauspost/compress/pull/1044)
-   s2: Improve small block compression speed w/o asm by [@&#8203;klauspost](https://redirect.github.com/klauspost) in [https://github.com/klauspost/compress/pull/1048](https://redirect.github.com/klauspost/compress/pull/1048)
-   flate: Fix matchlen L5+L6 by [@&#8203;klauspost](https://redirect.github.com/klauspost) in [https://github.com/klauspost/compress/pull/1049](https://redirect.github.com/klauspost/compress/pull/1049)
-   flate: Cleanup & reduce casts by [@&#8203;klauspost](https://redirect.github.com/klauspost) in [https://github.com/klauspost/compress/pull/1050](https://redirect.github.com/klauspost/compress/pull/1050)

#### New Contributors

-   [@&#8203;tcpdumppy](https://redirect.github.com/tcpdumppy) made their first contribution in [https://github.com/klauspost/compress/pull/1021](https://redirect.github.com/klauspost/compress/pull/1021)
-   [@&#8203;sam9291](https://redirect.github.com/sam9291) made their first contribution in [https://github.com/klauspost/compress/pull/1022](https://redirect.github.com/klauspost/compress/pull/1022)
-   [@&#8203;dezza](https://redirect.github.com/dezza) made their first contribution in [https://github.com/klauspost/compress/pull/1023](https://redirect.github.com/klauspost/compress/pull/1023)
-   [@&#8203;alingse](https://redirect.github.com/alingse) made their first contribution in [https://github.com/klauspost/compress/pull/1028](https://redirect.github.com/klauspost/compress/pull/1028)
-   [@&#8203;hyunsooda](https://redirect.github.com/hyunsooda) made their first contribution in [https://github.com/klauspost/compress/pull/1031](https://redirect.github.com/klauspost/compress/pull/1031)
-   [@&#8203;Juneezee](https://redirect.github.com/Juneezee) made their first contribution in [https://github.com/klauspost/compress/pull/1038](https://redirect.github.com/klauspost/compress/pull/1038)
-   [@&#8203;Bbulatov](https://redirect.github.com/Bbulatov) made their first contribution in [https://github.com/klauspost/compress/pull/1052](https://redirect.github.com/klauspost/compress/pull/1052)

**Full Changelog**: https://github.com/klauspost/compress/compare/v1.17.11...v1.18.0

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

### Comment by @jira-linking on April 06, 2025 at 01:14 PM UTC

Commits missing Jira IDs:
9bc69db731b3878e3572a2baf6248da5a1e66c24


### Comment by @codecov-commenter on April 06, 2025 at 01:19 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1633?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 58.16%. Comparing base [(`dde5824`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/dde582420b242db97254d841074ded8e0ba5d6eb?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`9bc69db`)](https://app.codecov.io/gh/RedHatInsights/patchman-engine/commit/9bc69db731b3878e3572a2baf6248da5a1e66c24?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1633   +/-   ##
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

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1633/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [unittests](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1633/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `58.16% <ø> (ø)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-engine/pull/1633?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
</details>

### Comment by @red-hat-konflux on April 16, 2025 at 09:33 AM UTC

### Edited/Blocked Notification

Renovate will not automatically rebase this PR, because it does not recognize the last commit author and assumes somebody else may have edited the PR.

You can manually request rebase by checking the rebase/retry box above.

 ⚠️ **Warning**: custom changes will be lost.

### Comment by @MichaelMraka on April 23, 2025 at 08:20 AM UTC

/retest

### Comment by @psegedy on April 23, 2025 at 09:03 AM UTC

closing in favor of https://github.com/RedHatInsights/patchman-engine/pull/1644

### Comment by @red-hat-konflux on April 23, 2025 at 12:21 PM UTC

### Renovate Ignore Notification

Because you closed this PR without merging, Renovate will ignore this update (`v1.18.0`). You will get a PR once a newer version is released. To ignore this dependency forever, add it to the `ignoreDeps` array of your Renovate config.

If you accidentally closed this PR, or if you changed your mind: rename this PR to get a fresh replacement PR.

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1633*
