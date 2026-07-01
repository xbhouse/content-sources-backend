---
type: pull_request
number: 966
title: "Bundle size"
state: merged
author: mkholjuraev
created: 2023-02-21T15:05:56Z
updated: 2024-04-03T09:22:36Z
closed: 2023-03-14T20:31:40Z
merged: 2023-03-14T20:31:40Z
base_branch: master
head_branch: bundle-size
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/966
---

# Pull Request #966: Bundle size

**Author**: @mkholjuraev
**Created**: February 21, 2023 at 03:05 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `bundle-size`

## Description

# Description

Associated Jira ticket: # (issue)

This PR disables loading `@patternfly/react-icons`,` @patternfly/react-table` packages, but forces to load the shared deps from the default scope in federated module. With this in place, dependant apps, will not load their local copy of the package, but use the shared one as both a fallback module and main instance in the shared scope. This PR, should be looked at together with this demo Patch PR.


# How to test the PR

Please include steps to test your PR.

# Before the change


# After the change


# Dependent work link


# Checklist:

- [ ] The commit message has the Jira ticket linked
- [ ] PR has a short description
- [ ] Screenshots before and after the change are added
- [ ] Tests for the changes have been added
- [ ] README.md is updated if necessary
- [ ] Needs additional dependent work


---

## Discussion

### Comment by @codecov-commenter on February 21, 2023 at 03:13 PM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/966?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Base: **72.10**% // Head: **72.10**% // No change to project coverage :thumbsup:
> Coverage data is based on head [(`db6a464`)](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/966?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) compared to base [(`b0ceadc`)](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/b0ceadca7885babba8d6ca3ffbd800a724f64610?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> Patch has no changes to coverable lines.

<details><summary>Additional details and impacted files</summary>


```diff
@@           Coverage Diff           @@
##           master     #966   +/-   ##
=======================================
  Coverage   72.10%   72.10%           
=======================================
  Files         110      110           
  Lines        2613     2613           
  Branches      659      659           
=======================================
  Hits         1884     1884           
  Misses        729      729           
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/966?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [...c/PresentationalComponents/Snippets/EmptyStates.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/966?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TbmlwcGV0cy9FbXB0eVN0YXRlcy5qcw==) | `71.42% <ø> (ø)` | |
| [src/PresentationalComponents/Snippets/Error.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/966?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TbmlwcGV0cy9FcnJvci5qcw==) | `100.00% <ø> (ø)` | |
| [.../PresentationalComponents/Snippets/ExternalLink.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/966?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TbmlwcGV0cy9FeHRlcm5hbExpbmsuanM=) | `100.00% <ø> (ø)` | |
| [...sentationalComponents/TableView/TableViewAssets.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/966?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9UYWJsZVZpZXcvVGFibGVWaWV3QXNzZXRzLmpz) | `100.00% <ø> (ø)` | |
| [src/SmartComponents/PatchSet/PatchSetAssets.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/966?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldC9QYXRjaFNldEFzc2V0cy5qcw==) | `54.54% <ø> (ø)` | |
| [...rtComponents/Remediation/AsyncRemediationButton.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/966?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9SZW1lZGlhdGlvbi9Bc3luY1JlbWVkaWF0aW9uQnV0dG9uLmpz) | `33.33% <ø> (ø)` | |
| [src/Utilities/Helpers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/966?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | `81.84% <ø> (ø)` | |
| [src/Utilities/Hooks.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/966?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9Ib29rcy5qcw==) | `74.34% <ø> (ø)` | |
| [src/Utilities/constants.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/966?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9jb25zdGFudHMuanM=) | `100.00% <ø> (ø)` | |

Help us with your feedback. Take ten seconds to tell us [how you rate us](https://about.codecov.io/nps?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Have a feature suggestion? [Share it here.](https://app.codecov.io/gh/feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)

</details>

[:umbrella: View full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/966?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Do you have feedback about the report comment? [Let us know in this issue](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @Fewwy on March 13, 2023 at 08:30 PM UTC

I couldn't open patch when I'm running this PR.
The instruction says nothing about this issue so it's a bug?
![image](https://user-images.githubusercontent.com/62722417/224825189-901ac2c5-e274-4c32-a217-b917544f6fec.png)

![image](https://user-images.githubusercontent.com/62722417/224825050-55a02f0e-b093-432f-a8ad-4b219008bbc2.png)


### Comment by @mkholjuraev on March 14, 2023 at 01:16 PM UTC

@Fewwy fixed!

### Comment by @Fewwy on March 14, 2023 at 07:09 PM UTC

Everything works as expected

### Comment by @mkholjuraev on March 28, 2023 at 08:57 AM UTC

:tada: This PR is included in version 1.62.4 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.62.4)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.62.4)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @leSamo - Commented on February 28, 2023 at 04:16 PM UTC

### Review by @Fewwy - Approved on March 14, 2023 at 07:09 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/966*
