---
type: pull_request
number: 1206
title: "chore(deps): Updating FEC and related packages"
state: merged
author: Andrewgdewar
created: 2024-10-17T20:37:21Z
updated: 2024-10-25T19:12:09Z
closed: 2024-10-25T18:38:09Z
merged: 2024-10-25T18:38:09Z
base_branch: master
head_branch: RHINENG-7807
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1206
---

# Pull Request #1206: chore(deps): Updating FEC and related packages

**Author**: @Andrewgdewar
**Created**: October 17, 2024 at 08:37 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `RHINENG-7807`

## Description

# Description

This: 
- updates a number of modules identified by snyk
- removes many unneeded babel artifacts
- updates test configurations
- updates the fec configuration file

Why: 
- I am running a macbook and thus use Docker, this application used an older fec-config module that had a bug preventing usage with docker.
- This fec-config bump necessitated a number of other updates, although some are not required, these changes help to simplify/modernify this application.


# How to test the PR

Please include steps to test your PR.

# After the change
- Sanity check, make sure all things still work as expected. 


# Checklist:

- [x] The commit message has the Jira ticket linked
- [x] PR has a short description
~~- [ ] Screenshots before and after the change are added~~
- [x] Pre-existing tests are still passing  ~~for the changes have been added~~
- [x] README.md is updated if necessary
~~- [ ] Needs additional dependent work~~


---

## Discussion

### Comment by @app-sre-bot on October 17, 2024 at 08:37 PM UTC

Can one of the admins verify this patch?

### Comment by @codecov-commenter on October 18, 2024 at 05:11 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1206?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 64.01%. Comparing base [(`d87c000`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/d87c000ff12dc6485d54946fa72f9990f9ed50dd?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`081e1ad`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/081e1addb348fadb9db8cbaf77023f49437ff53a?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master    #1206   +/-   ##
=======================================
  Coverage   64.01%   64.01%           
=======================================
  Files         124      124           
  Lines        3235     3235           
  Branches      831      831           
=======================================
  Hits         2071     2071           
  Misses       1164     1164           
```

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1206?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @skateman on October 22, 2024 at 07:31 PM UTC

/ok-to-test

### Comment by @LightOfHeaven1994 on October 22, 2024 at 09:33 PM UTC

/retest

### Comment by @AsToNlele on October 23, 2024 at 07:49 AM UTC

/retest

### Comment by @LightOfHeaven1994 on October 23, 2024 at 03:31 PM UTC

/retest

### Comment by @Andrewgdewar on October 24, 2024 at 09:23 PM UTC

@LightOfHeaven1994 
(Deployment does not have minimum availability.) << 
If this failure is unrelated to this PR, would we be able to merge this sooner rather than later? 
If so do close all the superseded/enclosed tickets from above.
I'd like to get this repo's PR list cleaned up if at all possible.

Ping me directly if there is anything I can do to speed this along.

### Comment by @Andrewgdewar on October 25, 2024 at 02:45 PM UTC

/retest

### Comment by @LightOfHeaven1994 on October 25, 2024 at 06:34 PM UTC

So the long story short - PR is fine itself but pr_check is failing on deploying `vmaas` service and patchman it set as dependency. Some other services are failing only because of `vmaas` pods can't get up. I run pr check on master and see the same issue. This one LGTM too

### Comment by @mkholjuraev on October 25, 2024 at 07:12 PM UTC

:tada: This PR is included in version 1.68.1 :tada:

The release is available on [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.68.1)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @bastilian - Commented on October 18, 2024 at 08:55 AM UTC

### Review by @bastilian - Commented on October 18, 2024 at 08:56 AM UTC

### Review by @Andrewgdewar - Commented on October 18, 2024 at 02:35 PM UTC

### Review by @Andrewgdewar - Commented on October 18, 2024 at 02:40 PM UTC

### Review by @adonispuente - Commented on October 18, 2024 at 04:58 PM UTC

The app itself seems to start up and work as it should, ignore the comment on the fec.config. Im having test fail locally when I run this PR 

### Review by @Andrewgdewar - Commented on October 18, 2024 at 05:07 PM UTC

### Review by @Andrewgdewar - Commented on October 18, 2024 at 05:10 PM UTC

### Review by @adonispuente - Commented on October 18, 2024 at 05:31 PM UTC

LGTM. Though because of the amount of packages changing id rather have at least 1 other person review the changes. 

Overall I havent found anything concerning. The title says youre updating FEC config FOR the ticket, but id rather it just be ‘updating FEC and related packages’ in the commit message/title. I dont see how the two items are related. This is just for cleaner commit history.

### Review by @Hyperkid123 - Commented on October 21, 2024 at 08:07 AM UTC

### Review by @AsToNlele - Commented on October 21, 2024 at 02:56 PM UTC

Left some questions

### Review by @Andrewgdewar - Commented on October 21, 2024 at 03:28 PM UTC

### Review by @Andrewgdewar - Commented on October 21, 2024 at 03:28 PM UTC

### Review by @Andrewgdewar - Commented on October 21, 2024 at 03:28 PM UTC

### Review by @gkarat - Commented on October 22, 2024 at 03:44 PM UTC

### Review by @gkarat - Commented on October 22, 2024 at 03:45 PM UTC

### Review by @gkarat - Approved on October 22, 2024 at 03:53 PM UTC

### Review by @LightOfHeaven1994 - Approved on October 25, 2024 at 06:37 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1206*
