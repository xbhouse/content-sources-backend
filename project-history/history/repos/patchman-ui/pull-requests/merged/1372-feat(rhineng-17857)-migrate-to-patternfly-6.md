---
type: pull_request
number: 1372
title: "feat(RHINENG-17857): migrate to PatternFly 6"
state: merged
author: xbhouse
created: 2025-09-04T12:46:26Z
updated: 2025-09-22T16:27:17Z
closed: 2025-09-22T16:27:17Z
merged: 2025-09-22T16:27:16Z
base_branch: master
head_branch: pf6-migration
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1372
---

# Pull Request #1372: feat(RHINENG-17857): migrate to PatternFly 6

**Author**: @xbhouse
**Created**: September 04, 2025 at 12:46 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pf6-migration`

## Description

# Description

Associated Jira ticket: [RHINENG-17857](https://issues.redhat.com/browse/RHINENG-17857)

Please include a summary of the change, what this fixes/creates/improves.

- Migrates Patch to Patternfly 6

# How to test the PR

- Verify functionality still works and nothing is broken or looks off (check stage or prod to compare)

**Things left to do:**

Since we'll have a few people working on this PR, let's split up as best we can which files to work on fixing so we have minimal conflicts :) 

- [x] Fix up any styles / colors in PresentationalComponents
- [x] Fix up any styles / colors in SmartComponents
- [x] Fix up any styles / colors that are left
- [x] Fix any broken functionality
- [x] Verify with UX that everything looks as expected

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

### Comment by @app-sre-bot on September 04, 2025 at 12:46 PM UTC

Can one of the admins verify this patch?

### Comment by @xbhouse on September 10, 2025 at 01:13 PM UTC

@adonispuente @bastilian @LightOfHeaven1994 i see yall have commited to this repo, would either of you all mind taking a look and/or approving the workflows if you have access? or if there is someone else i should ask please let me know :smile_cat: 

### Comment by @LightOfHeaven1994 on September 10, 2025 at 01:24 PM UTC

@xbhouse yes sure! Workflow approved. Some of us will take a look 👍 

### Comment by @codecov-commenter on September 10, 2025 at 03:04 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1372?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
:x: Patch coverage is `50.00000%` with `19 lines` in your changes missing coverage. Please review.
:white_check_mark: Project coverage is 62.46%. Comparing base ([`ef29dd3`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/ef29dd3635f3c354e1be2b9ae0c88651484ef274?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)) to head ([`6a4d487`](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/6a4d48779b4f614799683a5b29049b41099e7c87?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)).

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1372?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [...Components/PatchSetWizard/steps/RequestProgress.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1372?src=pr&el=tree&filepath=src%2FSmartComponents%2FPatchSetWizard%2Fsteps%2FRequestProgress.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9QYXRjaFNldFdpemFyZC9zdGVwcy9SZXF1ZXN0UHJvZ3Jlc3MuanM=) | 0.00% | [8 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1372?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [...c/PresentationalComponents/Snippets/EmptyStates.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1372?src=pr&el=tree&filepath=src%2FPresentationalComponents%2FSnippets%2FEmptyStates.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TbmlwcGV0cy9FbXB0eVN0YXRlcy5qcw==) | 28.57% | [5 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1372?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [...ntationalComponents/Header/HeaderBreadcrumbs.cy.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1372?src=pr&el=tree&filepath=src%2FPresentationalComponents%2FHeader%2FHeaderBreadcrumbs.cy.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9IZWFkZXIvSGVhZGVyQnJlYWRjcnVtYnMuY3kuanM=) | 0.00% | [3 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1372?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [.../PresentationalComponents/Filters/VersionFilter.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1372?src=pr&el=tree&filepath=src%2FPresentationalComponents%2FFilters%2FVersionFilter.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9GaWx0ZXJzL1ZlcnNpb25GaWx0ZXIuanM=) | 91.66% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1372?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [...tationalComponents/Snippets/DescriptionWithLink.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1372?src=pr&el=tree&filepath=src%2FPresentationalComponents%2FSnippets%2FDescriptionWithLink.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TbmlwcGV0cy9EZXNjcmlwdGlvbldpdGhMaW5rLmpz) | 0.00% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1372?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [...Components/StatusReports/AdvisoriesStatusReport.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1372?src=pr&el=tree&filepath=src%2FPresentationalComponents%2FStatusReports%2FAdvisoriesStatusReport.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1ByZXNlbnRhdGlvbmFsQ29tcG9uZW50cy9TdGF0dXNSZXBvcnRzL0Fkdmlzb3JpZXNTdGF0dXNSZXBvcnQuanM=) | 0.00% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1372?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>



```diff
@@            Coverage Diff             @@
##           master    #1372      +/-   ##
==========================================
- Coverage   62.57%   62.46%   -0.11%     
==========================================
  Files         126      126              
  Lines        3337     3312      -25     
  Branches      866      857       -9     
==========================================
- Hits         2088     2069      -19     
+ Misses       1128     1122       -6     
  Partials      121      121              
```
</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1372?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
<details><summary> :rocket: New features to boost your workflow: </summary>

- :snowflake: [Test Analytics](https://docs.codecov.com/docs/test-analytics): Detect flaky tests, report on failures, and find test suite problems.
- :package: [JS Bundle Analysis](https://docs.codecov.com/docs/javascript-bundle-analysis): Save yourself from yourself by tracking and limiting bundle sizes in JS merges.
</details>

### Comment by @LightOfHeaven1994 on September 10, 2025 at 03:23 PM UTC

/ok-to-test

### Comment by @xbhouse on September 12, 2025 at 04:14 PM UTC

@adonispuente thank you for reviewing! :heart: the UI i'm looking at when running the frontend locally looks different from your screenshots :thinking: this is what i see on Systems details:
<img width="1608" height="523" alt="Screenshot From 2025-09-12 12-05-30" src="https://github.com/user-attachments/assets/9dc6bf33-03c9-47f7-bc3f-229d29e6a3ad" />

i do see the color issue in the alert (should be more blue than purple?) but i don't see the grey background

i'll work on fixing this, changing over all the v5 classes, and making the icons consistent :+1: i'm curious though why you're seeing more of a pf5 UI

EDIT: hmm also not seeing any v5 classes in this branch after a global search for v5- :thinking: 

### Comment by @xbhouse on September 17, 2025 at 05:59 PM UTC

> Idk what the heck was goin on in the differences, I also now can find any v5- after global searching. Everything now looks great, this was the only noticable difference i could find. Overall though functionality looks good, and I wont block the PR on it. Its just resolving the merge conflict and this color difference
> <img alt="Screenshot 2025-09-17 at 11 13 52 AM" width="955" height="122" src="https://private-user-images.githubusercontent.com/60629070/490655987-e7ed1263-0d9a-4ddb-afcb-80c01b68472e.png?jwt=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3NTgxMzE1ODIsIm5iZiI6MTc1ODEzMTI4MiwicGF0aCI6Ii82MDYyOTA3MC80OTA2NTU5ODctZTdlZDEyNjMtMGQ5YS00ZGRiLWFmY2ItODBjMDFiNjg0NzJlLnBuZz9YLUFtei1BbGdvcml0aG09QVdTNC1ITUFDLVNIQTI1NiZYLUFtei1DcmVkZW50aWFsPUFLSUFWQ09EWUxTQTUzUFFLNFpBJTJGMjAyNTA5MTclMkZ1cy1lYXN0LTElMkZzMyUyRmF3czRfcmVxdWVzdCZYLUFtei1EYXRlPTIwMjUwOTE3VDE3NDgwMlomWC1BbXotRXhwaXJlcz0zMDAmWC1BbXotU2lnbmF0dXJlPTY4OWU4YTQ3MjA3MjZmMjEyNTY0NGI1NjhjNDlkNTJlOGE2MGZhZmJkNjE5MjM4YTg3YjQ5NDViYTJkZDBkMWQmWC1BbXotU2lnbmVkSGVhZGVycz1ob3N0In0.G8TyKZPhBdP6AT70YwnT_4CkNDy4Afod399tI1AuilE">

@adonispuente thank you!! just one more question, i'm a little confused as that looks like the repositories page. which color difference do you mean exactly? 

### Comment by @adonispuente on September 17, 2025 at 06:04 PM UTC

Sorry, I shouldve taken a larger picture. Its just the in progress cell on one of the columns. On stage right now its a lighter blue, but in the migrated pf6 styling its darker

### Comment by @xbhouse on September 17, 2025 at 06:15 PM UTC

@adonispuente oh weird :thinking:  hmm i'm not sure why the repositories page would be affected by the pf6 migration in patch. i don't see a difference in firefox, is this in chrome? 

### Comment by @xbhouse on September 22, 2025 at 01:52 PM UTC

@adonispuente if you think this is okay to merge, please go ahead! i don't have merge rights and this PR is blocking the FEO migration for patch and content-sources. i can open follow-up PRs afterward to fix up any remaining style weirdness etc but it would be good to get the bulk of the migration in asap :pray: :smiley: 

### Comment by @adonispuente on September 22, 2025 at 04:11 PM UTC

/retest

---

## Reviews

### Review by @adonispuente - Changes Requested on September 11, 2025 at 07:57 PM UTC

Visually, there are some things not consistent with Patternfly visuals. Most being fixed with the removal of  pf-v5-c-page__main-section class or related styling. We no longer aim for the grey background spacing/color. Everything is white. 
The following pages should be altered to have spacing but with white backgrounds,

Advisories page, Advisory details, Packages, Packages details, Systems (this page also needs some filterchip love), and system details pages. Below is the padding/grey im referring too 
<img width="738" height="548" alt="Screenshot 2025-09-11 at 2 51 03 PM" src="https://github.com/user-attachments/assets/7a7539ca-e1b1-4e91-a54d-48ed318de379" />

The aim is for things to look more like this: 
<img width="783" height="497" alt="Screenshot 2025-09-11 at 2 51 22 PM" src="https://github.com/user-attachments/assets/fbad35de-cbc2-4d24-a3e8-5ac1d27d6123" />
There is still padding from the nav bar, but all the grey border stuff is gone. Heres how filter chips look: 
<img width="606" height="204" alt="Screenshot 2025-09-11 at 2 52 08 PM" src="https://github.com/user-attachments/assets/372aff2e-68d5-47c0-9a92-d1290ccdd45b" />


Doing a global search still results in v5 classes imported in a few files like ADvisories/details for it/ a few others.  (Doing a global search for 'v5-')

This may also be an opportunity to fix the sizing differences between these icons in details pages. 
Overall functionality seems to be unaffected though on initial pass through 
<img width="228" height="280" alt="Screenshot 2025-09-11 at 2 54 29 PM" src="https://github.com/user-attachments/assets/bf8908d2-4fc2-4d64-ab32-75a49ca8eb63" />


### Review by @adonispuente - Approved on September 17, 2025 at 04:18 PM UTC

Idk what the heck was goin on in the differences, I also now can find any v5- after global searching. Everything now looks great, this was the only noticable difference i could find. Overall though functionality looks good, and I wont block the PR on it. Its just resolving the merge conflict and this color difference

<img width="955" height="122" alt="Screenshot 2025-09-17 at 11 13 52 AM" src="https://github.com/user-attachments/assets/e7ed1263-0d9a-4ddb-afcb-80c01b68472e" />


---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1372*
