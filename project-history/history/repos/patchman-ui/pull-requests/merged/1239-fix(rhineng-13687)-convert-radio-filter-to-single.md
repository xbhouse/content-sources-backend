---
type: pull_request
number: 1239
title: "fix(RHINENG-13687): convert radio filter to single select"
state: merged
author: johnsonm325
created: 2025-01-10T17:08:22Z
updated: 2025-02-05T17:12:01Z
closed: 2025-02-05T17:12:01Z
merged: 2025-02-05T17:12:01Z
base_branch: master
head_branch: rhineng-13687
labels: ["enhancement"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1239
---

# Pull Request #1239: fix(RHINENG-13687): convert radio filter to single select

**Author**: @johnsonm325
**Created**: January 10, 2025 at 05:08 PM UTC
**Status**: Merged
**Labels**: `enhancement`
**Base**: `master` ← **Head**: `rhineng-13687`

## Description

This PR leverages the new conditional filter type singleSelect to convert radio dropdown filters to single select dropdown filters. These filters can be found on Advisories page -> Publish date filter and Systems page -> Updatable filter.

---

## Discussion

### Comment by @codecov-commenter on January 23, 2025 at 04:58 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1239?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: Patch coverage is `50.00000%` with `1 line` in your changes missing coverage. Please review.
> Project coverage is 63.59%. Comparing base [(`1169d86`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/1169d86f5b5a1f1d41739218f0f46b7cac362191?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`ae495b5`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/ae495b537b4fe1ff573097a02402737e998212ea?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

| [Files with missing lines](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1239?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [src/Utilities/Helpers.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1239?src=pr&el=tree&filepath=src%2FUtilities%2FHelpers.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | 50.00% | [1 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1239?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1239      +/-   ##
==========================================
- Coverage   63.60%   63.59%   -0.01%     
==========================================
  Files         124      124              
  Lines        3267     3269       +2     
  Branches      860      860              
==========================================
+ Hits         2078     2079       +1     
- Misses       1189     1190       +1     
```

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1239?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @LightOfHeaven1994 on January 27, 2025 at 12:11 PM UTC

@johnsonm325 we get ```Error: Invalid conditional filter component type! Expected one of text,checkbox,radio,custom,group, got singleSelect``` on Systems page for some reason. But it works on Advisory page

### Comment by @johnsonm325 on January 30, 2025 at 06:48 PM UTC

@LightOfHeaven1994 The Systems table filter should work now. The PR to fix the Inventory Table has been merged: https://github.com/RedHatInsights/insights-inventory-frontend/pull/2337.

And the release to production was done today.

### Comment by @johnsonm325 on January 30, 2025 at 06:57 PM UTC

/retest

### Comment by @johnsonm325 on January 30, 2025 at 07:08 PM UTC

/retest

### Comment by @johnsonm325 on January 30, 2025 at 07:19 PM UTC

/retest

### Comment by @AsToNlele on February 04, 2025 at 04:30 PM UTC

To fix the build, rename the `deploy/frontend.yml` to `frontend.yaml`
You can verify by running `npx fec build`
Hopefully this doesn't break anything

### Comment by @johnsonm325 on February 04, 2025 at 07:09 PM UTC

/retest

### Comment by @johnsonm325 on February 04, 2025 at 07:21 PM UTC

/retest

### Comment by @johnsonm325 on February 05, 2025 at 04:29 PM UTC

/retest

---

## Reviews

### Review by @gkarat - Commented on January 24, 2025 at 12:59 PM UTC

### Review by @AsToNlele - Approved on February 04, 2025 at 04:28 PM UTC

LGTM, well done! 😄 

### Review by @Fewwy - Approved on February 04, 2025 at 08:13 PM UTC

LGTM

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1239*
