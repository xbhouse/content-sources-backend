---
type: pull_request
number: 1208
title: "HMS 4849: Remove/update features to support new content templates"
state: merged
author: Andrewgdewar
created: 2024-10-24T21:44:29Z
updated: 2024-11-12T17:57:55Z
closed: 2024-11-12T17:46:50Z
merged: 2024-11-12T17:46:50Z
base_branch: master
head_branch: HMS-4849
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1208
---

# Pull Request #1208: HMS 4849: Remove/update features to support new content templates

**Author**: @Andrewgdewar
**Created**: October 24, 2024 at 09:44 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `HMS-4849`

## Description

# Description

Associated Jira ticket: 
[RHINENG-7807](https://issues.redhat.com/browse/RHINENG-7807)
[HMS-4849](https://issues.redhat.com/browse/HMS-4849)

**NOTE: As many of these changes may be reverted/altered in the near-term, code has been commented out instead of removed; I understand this likely isn't the favored approach for all, but this makes things far easier to revert in pieces later.**

Remove/update features to support new content templates (All changes listed below):

- The template value on Content > Patch > Systems should pull from the 'template_name' instead of the 'baseline_name' attribute in the /patch/v3/systems endpoint
- The template name on Content > Patch > Systems should link to /insights/content/templates/template_id/details where the 'template_id' attribute is pulled from the /patch/v3/systems/inventory_id endpoint
- The template value on Content > Patch > Systems > inventory_id should pull from the 'template_name' instead of 'baseline_name' attribute in the /patch/v3/systems/inventory_id endpoint
- The template name on Content > Patch > Systems > inventory_id should link to /insights/content/templates/ {template_id}/details where the 'template_id' attribute is pulled from the /patch/v3/systems/inventory_id endpoint
- The template value on Content > Advisory > advisory_id system list should pull from the 'template_name' instead of 'baseline_name' attribute in the /patch/v3/systems/inventory_id endpoint
- The template name on Content > Advisory > advisory_id system list should link to /insights/content/templates/{template_id}
- /details where the 'template_id' attribute is pulled from the /patch/v3/systems/inventory_id endpoint
- The template value on Content > Package > package_name system list should pull from the 'template_name' instead of 'baseline_name' attribute in the /patch/v3/systems/inventory_id endpoint
- The template name on Content > Package > package_name system list should link to /insights/content/templates/ {template_id}/details where the 'template_id' attribute is pulled from the /patch/v3/systems/inventory_id endpoint
- Remove the bulk template assign and delete actions on the Content > Patch > Systems page (just remove the whole kebab menu)
- The "Assign to the template" and "Remove from a template" options in the table row kebab menu on the Content > Patch > Systems page and the Actions menu on the Content > Patch > Systems > UUID are removed (we might bring them back one day, but won't have them for now)

- **Remove the Content > Patch > Templates screens and all child screens from the UI**


# How to test the PR

# Before the change

The adjusted links should link to the patch>templates page
The items that are removed should all be present.

# After the change

The adjusted links should now link to the content>templates page
The items that have been removed should no longer be present.

# Dependent work link
[This needs to merged](https://github.com/RedHatInsights/patchman-ui/pull/1206)

# Checklist:

- [x] The commit message has the Jira ticket linked
- [x] PR has a short description
- [ ] Screenshots before and after the change are added
- [x] Tests related to the removal of features have been commented out



---

## Discussion

### Comment by @app-sre-bot on October 24, 2024 at 09:46 PM UTC

Can one of the admins verify this patch?

### Comment by @codecov-commenter on October 24, 2024 at 09:52 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1208?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: Patch coverage is `12.50000%` with `7 lines` in your changes missing coverage. Please review.
> Project coverage is 63.98%. Comparing base [(`e4e70f6`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/e4e70f6e4e3d43782ccf50fa3dbd04369b3c6f1b?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`9398221`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/939822140a407702bfd581985db52feb590c770b?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1208?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [src/SmartComponents/Systems/SystemsListAssets.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1208?src=pr&el=tree&filepath=src%2FSmartComponents%2FSystems%2FSystemsListAssets.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXNMaXN0QXNzZXRzLmpz) | 0.00% | [3 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1208?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [...rc/SmartComponents/SystemDetail/InventoryDetail.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1208?src=pr&el=tree&filepath=src%2FSmartComponents%2FSystemDetail%2FInventoryDetail.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1EZXRhaWwvSW52ZW50b3J5RGV0YWlsLmpz) | 0.00% | [2 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1208?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [...Components/AdvisorySystems/AdvisorySystemsTable.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1208?src=pr&el=tree&filepath=src%2FSmartComponents%2FAdvisorySystems%2FAdvisorySystemsTable.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9BZHZpc29yeVN5c3RlbXMvQWR2aXNvcnlTeXN0ZW1zVGFibGUuanM=) | 0.00% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1208?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |
| [src/Utilities/DataMappers.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1208?src=pr&el=tree&filepath=src%2FUtilities%2FDataMappers.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9EYXRhTWFwcGVycy5qcw==) | 0.00% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1208?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1208      +/-   ##
==========================================
- Coverage   64.01%   63.98%   -0.04%     
==========================================
  Files         124      124              
  Lines        3235     3210      -25     
  Branches      831      821      -10     
==========================================
- Hits         2071     2054      -17     
+ Misses       1164     1156       -8     
```

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1208?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @Andrewgdewar on October 25, 2024 at 07:38 PM UTC

**NOTE TO QE:** 
As this removes the "Templates" page and a number of key functionalities described above, you may have to comment-out/remove a number of tests!

### Comment by @Andrewgdewar on November 01, 2024 at 05:26 PM UTC

/ok-to-test

### Comment by @LightOfHeaven1994 on November 04, 2024 at 07:57 PM UTC

/retest

### Comment by @niyazRedhat on November 07, 2024 at 07:29 AM UTC

/retest

### Comment by @niyazRedhat on November 07, 2024 at 09:20 AM UTC

/retest

### Comment by @niyazRedhat on November 07, 2024 at 01:50 PM UTC

/retest

### Comment by @niyazRedhat on November 07, 2024 at 02:53 PM UTC

/retest

### Comment by @niyazRedhat on November 07, 2024 at 04:26 PM UTC

/retest

### Comment by @niyazRedhat on November 08, 2024 at 06:34 AM UTC

/ok-to-test



### Comment by @niyazRedhat on November 08, 2024 at 09:15 AM UTC

/retest

### Comment by @niyazRedhat on November 11, 2024 at 07:28 AM UTC

/ok-to-test

### Comment by @mkholjuraev on November 12, 2024 at 05:57 PM UTC

:tada: This PR is included in version 1.69.1 :tada:

The release is available on [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.69.1)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @AsToNlele - Approved on October 30, 2024 at 08:52 AM UTC

I think we can make an exception for the commented out code.
Looks good and changes are made as described.

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1208*
