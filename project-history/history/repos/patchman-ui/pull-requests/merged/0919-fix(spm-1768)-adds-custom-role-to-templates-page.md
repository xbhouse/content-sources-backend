---
type: pull_request
number: 919
title: "fix(SPM-1768): Adds custom role to templates page"
state: merged
author: adonispuente
created: 2022-11-23T16:34:56Z
updated: 2022-11-29T21:11:37Z
closed: 2022-11-29T20:54:24Z
merged: 2022-11-29T20:54:24Z
base_branch: master
head_branch: SPM-1752-2
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/919
---

# Pull Request #919: fix(SPM-1768): Adds custom role to templates page

**Author**: @adonispuente
**Created**: November 23, 2022 at 04:34 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `SPM-1752-2`

## Description

# Description

Associated Jira ticket: # (SPM-[1768](https://issues.redhat.com/browse/SPM-1768))

This adds a new role to the allow list of roles that can create/edit patch templates. Admin has access to everything, and I had set it so that only admin can perform template actions. The custom role should have access too it as well. This change allows both admin and a user with patch:template:write permissions to perform template actions

# How to test the PR

Navigate the Templates page with an account that has patch:template:write access and read permissions, and verify that you can create and edit a template
Do the same for an account with admin permissions
Verify that if a user with read permissions goes to the page they dont have access to the button and cant edit the templates.

If you need access to accounts with these permission's, please let me know
# Before the change

The custom role didnt work

# After the change
The custom role works

# Checklist:

- [x] The commit message has the Jira ticket linked
- [x] PR has a short description
- [ ] Screenshots before and after the change are added
- [ ] Tests for the changes have been added
- [ ] README.md is updated if necessary
- [ ] Needs additional dependent work


---

## Discussion

### Comment by @codecov-commenter on November 23, 2022 at 04:43 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/919?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **69.75**% // Head: **69.75**% // No change to project coverage :thumbsup:
> Coverage data is based on head [(`dd6c99c`)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/919?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`b84932b`)](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/b84932b03075d19ec5ee603e38e0a9cb0767841d?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch has no changes to coverable lines.

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master     #919   +/-   ##
=======================================
  Coverage   69.75%   69.75%           
=======================================
  Files         111      111           
  Lines        2605     2605           
  Branches      677      677           
=======================================
  Hits         1817     1817           
  Misses        712      712           
  Partials       76       76           
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/919?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/SmartComponents/PatchSet/PatchSet.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/919/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldC9QYXRjaFNldC5qcw==) | `0.00% <ø> (ø)` | |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/919?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on November 29, 2022 at 03:44 PM UTC

@adonispuente users with admin rights are not checked by RBAC, rather it is checked by the platform/engine. Am I right? I want to understand how users with admin rights are checked.

### Comment by @adonispuente on November 29, 2022 at 07:43 PM UTC

@mkholjuraev The answer is both RBAC and the engine check for it. Its like having 2 layers of security. Its essential to have it on the backend (at-least thats what I think if you're only going to pick one), but on the front-end (taking the templates page for example) if a user with invalid access was somehow able to send a POST request and send a template, it would be denied by the engine. Having it checked on the frontend just allows you to do things like this where a user can see the page without having the ability to make any templates. It really just helps add more options to give a user to experience the app. 

### Comment by @mkholjuraev on November 29, 2022 at 08:28 PM UTC

And now the PR looks good with `'patch:*:*'`. With this, we are checking admin users. Approving

### Comment by @jiridostal on November 29, 2022 at 09:11 PM UTC

:tada: This PR is included in version 1.57.2 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.57.2)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.57.2)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @mkholjuraev - Approved on November 29, 2022 at 08:29 PM UTC

LGTM! @adonispuente  thanks

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/919*
