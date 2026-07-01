---
type: pull_request
number: 1224
title: "Fixes 5033: Expose old patch templates UI via feature flag"
state: merged
author: Andrewgdewar
created: 2024-11-21T22:01:44Z
updated: 2024-12-11T18:38:11Z
closed: 2024-12-11T18:38:11Z
merged: 2024-12-11T18:38:11Z
base_branch: master
head_branch: HMS-5033
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1224
---

# Pull Request #1224: Fixes 5033: Expose old patch templates UI via feature flag

**Author**: @Andrewgdewar
**Created**: November 21, 2024 at 10:01 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `HMS-5033`

## Description

# Description

Associated Jira ticket: [HMS-5033](https://issues.redhat.com/browse/HMS-5033)

Goal: 

[This](https://issues.redhat.com/browse/RHINENG-7807)  caused all of the patch template UI to be removed, and all systems list references to either be removed or to start using the new content template data
All of this now should been controlled with a feature flag  🫤 

Acceptance Criteria:

Old UI for patch and patch templates exists in stable
Patch uses the new content templates in preview, old patch templates are not available
Stable patch UI can be flipped from the old  templates to new templates in the future easily


# How to test the PR

Run with proxy against prod, test changes by toggling preview on and off on various screens.
Original list of changes can be found [here](https://issues.redhat.com/browse/RHINENG-7807).

Alternatively, one can use the feature flag for [stage](https://insights-stage.unleash.devshift.net/projects/default/features/patchman-ui.template-update.enabled) or [prod](https://insights.unleash.devshift.net/projects/default/features/patchman-ui.template-update.enabled) 
To test functionality directly.

# Before the change
All environments would show new UI implementation.

# After the change
Prod (non-preview) will now show the old ui, until the unleashed feature flag "patchman-ui.template-update.enabled" is enabled for that environment.

# Checklist:

- [x] The commit message has the Jira ticket linked
- [x] PR has a short description
- [ ] Screenshots before and after the change are added


---

## Discussion

### Comment by @app-sre-bot on November 21, 2024 at 10:04 PM UTC

Can one of the admins verify this patch?

### Comment by @codecov-commenter on November 21, 2024 at 10:10 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1224?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: Patch coverage is `25.33333%` with `56 lines` in your changes missing coverage. Please review.
> Project coverage is 63.60%. Comparing base [(`096d743`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/096d743b11678edb7b0022830f7e13231bde34dd?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`7decbca`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/7decbcaae4c664a2eac97c082fac1148215de055?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1224?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [src/SmartComponents/Systems/SystemsListAssets.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1224?src=pr&el=tree&filepath=src%2FSmartComponents%2FSystems%2FSystemsListAssets.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXNMaXN0QXNzZXRzLmpz) | 24.13% | [22 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1224?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [...rc/SmartComponents/SystemDetail/InventoryDetail.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1224?src=pr&el=tree&filepath=src%2FSmartComponents%2FSystemDetail%2FInventoryDetail.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1EZXRhaWwvSW52ZW50b3J5RGV0YWlsLmpz) | 0.00% | [13 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1224?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [src/Routes.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1224?src=pr&el=tree&filepath=src%2FRoutes.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1JvdXRlcy5qcw==) | 0.00% | [8 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1224?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [...rtComponents/PatchSetWizard/steps/ReviewSystems.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1224?src=pr&el=tree&filepath=src%2FSmartComponents%2FPatchSetWizard%2Fsteps%2FReviewSystems.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9zdGVwcy9SZXZpZXdTeXN0ZW1zLmpz) | 0.00% | [4 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1224?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [src/SmartComponents/Systems/SystemsTable.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1224?src=pr&el=tree&filepath=src%2FSmartComponents%2FSystems%2FSystemsTable.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXNUYWJsZS5qcw==) | 55.55% | [4 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1224?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [...Components/AdvisorySystems/AdvisorySystemsTable.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1224?src=pr&el=tree&filepath=src%2FSmartComponents%2FAdvisorySystems%2FAdvisorySystemsTable.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yeVN5c3RlbXMvQWR2aXNvcnlTeXN0ZW1zVGFibGUuanM=) | 33.33% | [2 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1224?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [...c/SmartComponents/PatchSetDetail/PatchSetDetail.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1224?src=pr&el=tree&filepath=src%2FSmartComponents%2FPatchSetDetail%2FPatchSetDetail.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldERldGFpbC9QYXRjaFNldERldGFpbC5qcw==) | 0.00% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1224?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [src/SmartComponents/PatchSetWizard/WizardAssets.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1224?src=pr&el=tree&filepath=src%2FSmartComponents%2FPatchSetWizard%2FWizardAssets.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9XaXphcmRBc3NldHMuanM=) | 50.00% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1224?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [src/Utilities/DataMappers.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1224?src=pr&el=tree&filepath=src%2FUtilities%2FDataMappers.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9EYXRhTWFwcGVycy5qcw==) | 0.00% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1224?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1224      +/-   ##
==========================================
- Coverage   63.98%   63.60%   -0.39%     
==========================================
  Files         124      124              
  Lines        3210     3267      +57     
  Branches      821      860      +39     
==========================================
+ Hits         2054     2078      +24     
- Misses       1156     1189      +33     
```

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1224?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @Andrewgdewar on November 27, 2024 at 04:16 PM UTC

> When I create a new template with some systems and try deleting one, I get this error ![image](https://private-user-images.githubusercontent.com/20592948/390331346-32b057cf-86c9-40bd-87d7-4cd6cc65692e.png?jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3MzI3MjQ0NjksIm5iZiI6MTczMjcyNDE2OSwicGF0aCI6Ii8yMDU5Mjk0OC8zOTAzMzEzNDYtMzJiMDU3Y2YtODZjOS00MGJkLTg3ZDctNGNkNmNjNjU2OTJlLnBuZz9YLUFtei1BbGdvcml0aG09QVdTNC1ITUFDLVNIQTI1NiZYLUFtei1DcmVkZW50aWFsPUFLSUFWQ09EWUxTQTUzUFFLNFpBJTJGMjAyNDExMjclMkZ1cy1lYXN0LTElMkZzMyUyRmF3czRfcmVxdWVzdCZYLUFtei1EYXRlPTIwMjQxMTI3VDE2MTYwOVomWC1BbXotRXhwaXJlcz0zMDAmWC1BbXotU2lnbmF0dXJlPWMyMDEwOWFlM2UwMjUzYTU3ZWJlZmZjN2Y2NzY4ODM0YjZiOTAyZWZhY2RkZWMzOGMzNzNhNTNiMmJjMDRmNjcmWC1BbXotU2lnbmVkSGVhZGVycz1ob3N0In0.azPzncNhlHbh1Oh21DHdfd8IF4h_nLF3BhA0bWTkSyU)

Will look at this. 

Thanks for the review! 🍻 

### Comment by @Andrewgdewar on November 29, 2024 at 08:13 PM UTC

> When I create a new template with some systems and try deleting one, I get this error ![image](https://private-user-images.githubusercontent.com/20592948/390331346-32b057cf-86c9-40bd-87d7-4cd6cc65692e.png?jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3MzI5MTE0NzIsIm5iZiI6MTczMjkxMTE3MiwicGF0aCI6Ii8yMDU5Mjk0OC8zOTAzMzEzNDYtMzJiMDU3Y2YtODZjOS00MGJkLTg3ZDctNGNkNmNjNjU2OTJlLnBuZz9YLUFtei1BbGdvcml0aG09QVdTNC1ITUFDLVNIQTI1NiZYLUFtei1DcmVkZW50aWFsPUFLSUFWQ09EWUxTQTUzUFFLNFpBJTJGMjAyNDExMjklMkZ1cy1lYXN0LTElMkZzMyUyRmF3czRfcmVxdWVzdCZYLUFtei1EYXRlPTIwMjQxMTI5VDIwMTI1MlomWC1BbXotRXhwaXJlcz0zMDAmWC1BbXotU2lnbmF0dXJlPTIxMjVkZDhiNWVhMjZlNWExMTg5ZGZiOGI0ZTU3NTM1YWZjYzY2ZWJiZmZmY2MzYTQzY2Y5ZGRkZGQyNzZmMWUmWC1BbXotU2lnbmVkSGVhZGVycz1ob3N0In0.0k9AoHUYusNdW8JcdBVUJ457SW0IJMUNWZpm73waEe0)

in which environment is that?
I wasn't able to recreate the issue.

### Comment by @adonispuente on December 11, 2024 at 04:31 PM UTC

/ok-to-test

### Comment by @adonispuente on December 11, 2024 at 04:33 PM UTC

/ok-test-ok

### Comment by @adonispuente on December 11, 2024 at 04:35 PM UTC

/retest

### Comment by @adonispuente on December 11, 2024 at 05:36 PM UTC

/retest

### Comment by @Andrewgdewar on December 11, 2024 at 06:23 PM UTC

I'm unsure if the above Pr test failures are inherit to the changes in my PR (putting certain component functionality behind feature flags). 

@niyazRedhat can you confirm these failures would be expected? and comment here.
I'd like to get this merged.

---

## Reviews

### Review by @AsToNlele - Commented on November 26, 2024 at 04:03 PM UTC

When I create a new template with some systems and try deleting one, I get this error
![image](https://github.com/user-attachments/assets/32b057cf-86c9-40bd-87d7-4cd6cc65692e)


### Review by @Andrewgdewar - Commented on November 27, 2024 at 04:13 PM UTC

### Review by @Andrewgdewar - Commented on November 27, 2024 at 04:14 PM UTC

### Review by @johnsonm325 - Approved on December 10, 2024 at 10:00 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1224*
