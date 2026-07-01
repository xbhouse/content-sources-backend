---
type: pull_request
number: 515
title: "Upgrade to GitHub-native Dependabot"
state: merged
author: dependabot-preview
created: 2021-04-29T15:26:38Z
updated: 2022-11-16T08:17:00Z
closed: 2021-05-04T12:25:52Z
merged: 2021-05-04T12:25:51Z
base_branch: master
head_branch: dependabot/add-v2-config-file
labels: ["dependencies", "released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/515
---

# Pull Request #515: Upgrade to GitHub-native Dependabot

**Author**: @dependabot-preview
**Created**: April 29, 2021 at 03:26 PM UTC
**Status**: Merged
**Labels**: `dependencies`, `released`
**Base**: `master` ← **Head**: `dependabot/add-v2-config-file`

## Description

_Dependabot Preview will be shut down on August 3rd, 2021. In order to keep getting Dependabot updates, please merge this PR and migrate to GitHub-native Dependabot before then._

Dependabot has been fully integrated into GitHub, so you no longer have to install and manage a separate app. This pull request updates your config file to the [new syntax][new_syntax]. When merged, we'll swap out `dependabot-preview` (me) for a new `dependabot` app, and you'll be all set!

With this change, you'll now use the [Dependabot page in GitHub][dependabot_page], rather than the [Dependabot dashboard][dashboard], to monitor your version updates, and you'll configure Dependabot through the new config file rather than a UI.







If you've got any questions or feedback for us, please let us know by creating an issue in the [dependabot/dependabot-core][issues] repository.

[Learn more about migrating to GitHub-native Dependabot][learn]

Please note that regular `@dependabot` commands do not work on this pull request.

[dashboard]: https://app.dependabot.com/
[dependabot_page]: https://github.com/RedHatInsights/patchman-ui/network/updates
[issues]: https://github.com/dependabot/dependabot-core/issues/new?assignees=%40dependabot%2Fpreview-migration-reviewers&labels=E%3A+preview-migration&template=migration-issue.md
[learn]: http://docs.github.com/code-security/supply-chain-security/upgrading-from-dependabotcom-to-github-native-dependabot
[new_syntax]: https://help.github.com/en/github/administering-a-repository/configuration-options-for-dependency-updates
[org_secrets_url]: https://github.com/organizations/RedHatInsights/settings/secrets/dependabot
[repo_secrets_url]: https://github.com/RedHatInsights/patchman-ui/settings/secrets/dependabot


---

## Discussion

### Comment by @codecov-commenter on April 29, 2021 at 03:36 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/515?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#515](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/515?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (55c5490) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/2980a5bfbb3be5fac0361fd4bc021162a4ffc631?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (2980a5b) will **not change** coverage.
> The diff coverage is `n/a`.

[![Impacted file tree graph](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/515/graphs/tree.svg?width=650&height=150&src=pr&token=28maKEO5Bi&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/515?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

```diff
@@           Coverage Diff           @@
##           master     #515   +/-   ##
=======================================
  Coverage   54.49%   54.49%           
=======================================
  Files          78       78           
  Lines        1558     1558           
  Branches      183      183           
=======================================
  Hits          849      849           
  Misses        647      647           
  Partials       62       62           
```



------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/515?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/515?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [2980a5b...55c5490](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/515?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @jiridostal on May 05, 2021 at 12:22 PM UTC

:tada: This PR is included in version 1.17.12 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.17.12)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.17.12)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/515*
