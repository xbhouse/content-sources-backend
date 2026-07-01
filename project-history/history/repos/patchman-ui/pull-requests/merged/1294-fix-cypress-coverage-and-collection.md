---
type: pull_request
number: 1294
title: "fix Cypress coverage and collection"
state: merged
author: opacut
created: 2025-02-24T13:20:43Z
updated: 2025-02-26T06:55:58Z
closed: 2025-02-26T06:55:58Z
merged: 2025-02-26T06:55:58Z
base_branch: master
head_branch: opacut/fix_cypress_collection
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/1294
---

# Pull Request #1294: fix Cypress coverage and collection

**Author**: @opacut
**Created**: February 24, 2025 at 01:20 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `opacut/fix_cypress_collection`

## Description

Fix code coverage scripts and collection so that we have a cypress report generated for codecov.

Review steps:
- npm i
- npm run travis:verify
- check that the coverage-cypress/lcov.info is not empty
- check that the coverage-jest/lcov.info is not empty
- check that .nyc_output/out.json is not empty

---

## Discussion

### Comment by @codecov-commenter on February 24, 2025 at 01:28 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1294?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
All modified and coverable lines are covered by tests :white_check_mark:
> Project coverage is 62.04%. Comparing base [(`c2beca0`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/c2beca04371525d9dda28a182cb95dc9258cae96?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`9d0c860`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/9d0c860553722a7b230296ccdcf3a512da219547?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1294      +/-   ##
==========================================
- Coverage   63.65%   62.04%   -1.61%     
==========================================
  Files         124      126       +2     
  Lines        3282     3367      +85     
  Branches      868      868              
==========================================
  Hits         2089     2089              
- Misses       1193     1278      +85     
```

| [Flag](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1294/flags?src=pr&el=flags&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [combined](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1294/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `62.04% <ø> (?)` | |
| [cypress](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1294/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `100.00% <ø> (?)` | |
| [jest](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1294/flags?src=pr&el=flag&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | `62.04% <ø> (?)` | |

Flags with carried forward coverage won't be shown. [Click here](https://docs.codecov.io/docs/carryforward-flags?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#carryforward-flags-in-the-pull-request-comment) to find out more.

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1294?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


---

## Reviews

### Review by @LightOfHeaven1994 - Approved on February 25, 2025 at 09:50 AM UTC

LGTM! Great work!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1294*
