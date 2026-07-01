---
type: pull_request
number: 1136
title: "chore: Update npm run start scripts and documentation"
state: merged
author: gkarat
created: 2023-10-26T16:41:06Z
updated: 2023-11-09T11:01:04Z
closed: 2023-10-31T11:40:27Z
merged: 2023-10-31T11:40:27Z
base_branch: master
head_branch: update-scripts
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1136
---

# Pull Request #1136: chore: Update npm run start scripts and documentation

**Author**: @gkarat
**Created**: October 26, 2023 at 04:41 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `update-scripts`

## Description

no jira.

This is a quick PR that aims to make npm run start* scripts consistent across Insights frontend applications.

---

## Discussion

### Comment by @gkarat on October 26, 2023 at 04:41 PM UTC

Relates to https://github.com/RedHatInsights/vulnerability-ui/pull/2013

### Comment by @gkarat on October 26, 2023 at 04:43 PM UTC

@mkholjuraev, could you please help me checking if dev.webpack.config.js contains all the necessary parameters? For now, I put only those that are common across all apps. I am not sure about some extra ones that Patch added like useCloud or useFileHash. Since you are more experienced with Patch, you might have more insights if these parameters are used at all and if we need to keep them further

### Comment by @codecov-commenter on October 26, 2023 at 05:09 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1136?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:

[see 2 files with indirect coverage changes](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1136/indirect-changes?src=pr&el=tree-more&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)


:loudspeaker: Thoughts on this report? [Let us know!](https://about.codecov.io/pull-request-comment-report/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on November 09, 2023 at 11:00 AM UTC

:tada: This PR is included in version 1.65.0 :tada:

The release is available on [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.65.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @mkholjuraev - Commented on October 27, 2023 at 08:18 AM UTC

The app runs pretty well. According to team's decision, we can keep or remove following params:

1. useFileHash: this is needed to turn on caching in dev env and platform advices to do so. Ref: [consuming-module](http://front-end-docs-insights.apps.ocp4.prod.psi.redhat.com/fed/consuming-module) 
2. debug: as this is dev env, do we want to have debug mode enabled?
3. proxyVerbose: this is intended for the output of what the proxy is doing when we run the app. I am not sure if we need it (Attached image)

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1136*
