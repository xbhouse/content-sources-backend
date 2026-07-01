---
type: pull_request
number: 1177
title: "Bump golang.org/x/net from 0.0.0-20221014081412-f15817d10f9b to 0.7.0"
state: merged
author: dependabot
created: 2023-02-25T02:29:17Z
updated: 2023-02-27T09:43:21Z
closed: 2023-02-27T09:43:11Z
merged: 2023-02-27T09:43:10Z
base_branch: master
head_branch: dependabot/go_modules/golang.org/x/net-0.7.0
labels: ["dependencies"]
url: https://github.com/RedHatInsights/patchman-engine/pull/1177
---

# Pull Request #1177: Bump golang.org/x/net from 0.0.0-20221014081412-f15817d10f9b to 0.7.0

**Author**: @dependabot
**Created**: February 25, 2023 at 02:29 AM UTC
**Status**: Merged
**Labels**: `dependencies`
**Base**: `master` ← **Head**: `dependabot/go_modules/golang.org/x/net-0.7.0`

## Description

Bumps [golang.org/x/net](https://github.com/golang/net) from 0.0.0-20221014081412-f15817d10f9b to 0.7.0.
<details>
<summary>Commits</summary>
<ul>
<li>See full diff in <a href="https://github.com/golang/net/commits/v0.7.0">compare view</a></li>
</ul>
</details>
<br />


[![Dependabot compatibility score](https://dependabot-badges.githubapp.com/badges/compatibility_score?dependency-name=golang.org/x/net&package-manager=go_modules&previous-version=0.0.0-20221014081412-f15817d10f9b&new-version=0.7.0)](https://docs.github.com/en/github/managing-security-vulnerabilities/about-dependabot-security-updates#about-compatibility-scores)

Dependabot will resolve any conflicts with this PR as long as you don't alter it yourself. You can also trigger a rebase manually by commenting `@dependabot rebase`.

[//]: # (dependabot-automerge-start)
[//]: # (dependabot-automerge-end)

---

<details>
<summary>Dependabot commands and options</summary>
<br />

You can trigger Dependabot actions by commenting on this PR:
- `@dependabot rebase` will rebase this PR
- `@dependabot recreate` will recreate this PR, overwriting any edits that have been made to it
- `@dependabot merge` will merge this PR after your CI passes on it
- `@dependabot squash and merge` will squash and merge this PR after your CI passes on it
- `@dependabot cancel merge` will cancel a previously requested merge and block automerging
- `@dependabot reopen` will reopen this PR if it is closed
- `@dependabot close` will close this PR and stop Dependabot recreating it. You can achieve the same result by closing it manually
- `@dependabot ignore this major version` will close this PR and stop Dependabot creating any more for this major version (unless you reopen the PR or upgrade to it yourself)
- `@dependabot ignore this minor version` will close this PR and stop Dependabot creating any more for this minor version (unless you reopen the PR or upgrade to it yourself)
- `@dependabot ignore this dependency` will close this PR and stop Dependabot creating any more for this dependency (unless you reopen the PR or upgrade to it yourself)
- `@dependabot use these labels` will set the current labels as the default for future PRs for this repo and language
- `@dependabot use these reviewers` will set the current reviewers as the default for future PRs for this repo and language
- `@dependabot use these assignees` will set the current assignees as the default for future PRs for this repo and language
- `@dependabot use this milestone` will set the current milestone as the default for future PRs for this repo and language

You can disable automated security fix PRs for this repo from the [Security Alerts page](https://github.com/RedHatInsights/patchman-engine/network/alerts).

</details>

---

## Discussion

### Comment by @jira-linking on February 25, 2023 at 02:29 AM UTC

Commits missing Jira IDs:
3a0be6b30a30017fbea5822950db728778e5a2de


### Comment by @codecov-commenter on February 25, 2023 at 02:36 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1177?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **63.21**% // Head: **63.26**% // Increases project coverage by **`+0.05%`** :tada:
> Coverage data is based on head [(`3a0be6b`)](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1177?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`4999506`)](https://codecov.io/gh/RedHatInsights/patchman-engine/commit/499950637e9720a44342710aa18ca5843aa44b6d?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch coverage: 85.18% of modified lines in pull request are covered.

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1177      +/-   ##
==========================================
+ Coverage   63.21%   63.26%   +0.05%     
==========================================
  Files         102      102              
  Lines        5878     5897      +19     
==========================================
+ Hits         3716     3731      +15     
- Misses       1696     1699       +3     
- Partials      466      467       +1     
```

| Flag | Coverage Δ | |
|---|---|---|
| unittests | `63.26% <85.18%> (+0.05%)` | :arrow_up: |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1177?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [manager/middlewares/authentication.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1177?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9taWRkbGV3YXJlcy9hdXRoZW50aWNhdGlvbi5nbw==) | `20.28% <0.00%> (-0.30%)` | :arrow_down: |
| [manager/controllers/utils.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1177?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy91dGlscy5nbw==) | `87.53% <66.66%> (-0.17%)` | :arrow_down: |
| [manager/controllers/baseline\_create.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1177?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9jcmVhdGUuZ28=) | `76.71% <75.00%> (-0.57%)` | :arrow_down: |
| [manager/controllers/advisories.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1177?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9hZHZpc29yaWVzLmdv) | `90.36% <100.00%> (-0.12%)` | :arrow_down: |
| [manager/controllers/baseline\_detail.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1177?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV9kZXRhaWwuZ28=) | `89.36% <100.00%> (+0.98%)` | :arrow_up: |
| [manager/controllers/baseline\_update.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1177?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZV91cGRhdGUuZ28=) | `79.79% <100.00%> (+0.20%)` | :arrow_up: |
| [manager/controllers/baselines.go](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1177?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-bWFuYWdlci9jb250cm9sbGVycy9iYXNlbGluZXMuZ28=) | `92.45% <100.00%> (+0.61%)` | :arrow_up: |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-engine/pull/1177?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/1177*
