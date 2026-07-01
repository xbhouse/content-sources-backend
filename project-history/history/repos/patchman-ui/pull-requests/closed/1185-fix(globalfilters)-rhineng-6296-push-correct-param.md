---
type: pull_request
number: 1185
title: "fix(GlobalFilters): RHINENG-6296 push correct params to url"
state: closed
author: darkeriossss
created: 2024-06-17T12:25:33Z
updated: 2025-04-17T09:18:14Z
closed: 2025-04-17T09:18:14Z
base_branch: master
head_branch: global-filters
labels: ["bug"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1185
---

# Pull Request #1185: fix(GlobalFilters): RHINENG-6296 push correct params to url

**Author**: @darkeriossss
**Created**: June 17, 2024 at 12:25 PM UTC
**Status**: Closed
**Labels**: `bug`
**Base**: `master` ← **Head**: `global-filters`

## Description

# Description

Associated Jira ticket: # [RHINENG-6296](https://issues.redhat.com/browse/RHINENG-6296)

Fixed pushing global filters to URL


# How to test the PR

1. Open Patch systems page
2. Apply Global filters
3. Global Filters should be added to the URL

# Before the change

![image](https://github.com/RedHatInsights/patchman-ui/assets/62351699/f9eecdcb-78ac-45a4-8e93-adcd92530915)

# After the change

![image](https://github.com/RedHatInsights/patchman-ui/assets/62351699/869f4d7e-88fb-42ad-92c5-2fdd6c6b57b7)

# Checklist:

- [x] The commit message has the Jira ticket linked
- [x] PR has a short description
- [x] Screenshots before and after the change are added
- [ ] Tests for the changes have been added
- [ ] README.md is updated if necessary
- [ ] Needs additional dependent work


---

## Discussion

### Comment by @codecov-commenter on June 17, 2024 at 12:41 PM UTC

## [Codecov](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1185?dropdown=coverage&src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
Attention: Patch coverage is `75.00000%` with `2 lines` in your changes missing coverage. Please review.
> Project coverage is 64.04%. Comparing base [(`fdcced0`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/fdcced002244e1929c3b781a1d3eb3359b411f26?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) to head [(`d164f1f`)](https://app.codecov.io/gh/RedHatInsights/patchman-ui/commit/d164f1f4f27bc16f8b49f1c123ec1395dca5939f?dropdown=coverage&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).

| [Files](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1185?dropdown=coverage&src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Patch % | Lines |
|---|---|---|
| [src/SmartComponents/Systems/SystemsTable.js](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1185?src=pr&el=tree&filepath=src%2FSmartComponents%2FSystems%2FSystemsTable.js&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1NtYXJ0Q29tcG9uZW50cy9TeXN0ZW1zL1N5c3RlbXNUYWJsZS5qcw==) | 75.00% | [2 Missing :warning: ](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1185?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) |

<details><summary>Additional details and impacted files</summary>


```diff
@@            Coverage Diff             @@
##           master    #1185      +/-   ##
==========================================
+ Coverage   64.01%   64.04%   +0.02%     
==========================================
  Files         124      124              
  Lines        3235     3243       +8     
  Branches      831      833       +2     
==========================================
+ Hits         2071     2077       +6     
- Misses       1164     1166       +2     
```

</details>

[:umbrella: View full report in Codecov by Sentry](https://app.codecov.io/gh/RedHatInsights/patchman-ui/pull/1185?dropdown=coverage&src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).   
:loudspeaker: Have feedback on the report? [Share it here](https://about.codecov.io/codecov-pr-comment-feedback/?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @bastilian on June 17, 2024 at 03:24 PM UTC

Can you elaborate a bit more on what the issue was? 

### Comment by @darkeriossss on June 18, 2024 at 08:20 AM UTC

basically, `SystemsTable` did not use the hook for managing URL params which is used in other places in patch. using this hook fixed the issue

### Comment by @bastilian on June 18, 2024 at 10:16 AM UTC

@xmicha82 This doesn't feel like a fix, but a work around. We should also not need this `firstMount` state.

### Comment by @darkeriossss on June 19, 2024 at 08:57 AM UTC

@bastilian this change is consistent with other pages in patch, which also use the `historyPusher`. could you please explain why do you think it's not the right way to fix this issue?

### Comment by @bastilian on July 23, 2024 at 12:39 AM UTC

@xmicha82 I took a quick look at this. I used `master` without any changes, but this in the `SystemsTable`:

```js
    const location = useLocation();
    console.log('LOCATION CHANGE', location);
```

Here is what I saw:

https://github.com/user-attachments/assets/82305bc1-6e59-402f-a000-e138fa3400a1

If you pay attention to the console you will notice that, first of all, there are probably updates to the URL multiple times per render and re renders, and if you look that location after the global filter is set you'll see that at some point the `hash` is lost.

A proper fix for this issue will require us to take a closer look and see why, how and where changes to the URL happen when the `SystemsTable` loads and ensure that all of these places do not change the `hash` and preserve the global filter. 

Ideally we would also look into reducing the amount of updates to just one update to the URL per filter change, but for now just making sure that the global filter hash is not lost should be fix enough.


### Comment by @bastilian on July 23, 2024 at 12:59 AM UTC

A better place to start debugging might be the `SystemsMainContent` where we use `useSearchParams`. Adding this log output can give you an idea of how many times the URL is updated:

```js
    const setSearchParams = (...args) => {
        console.log('SetSearch', ...args);
        setSearchParamsOrig(...args);
    };
```

~The `setSerachParams` function provided by react router also allows navigate options as a second parameter, this might allow us to set the `hash` of the current location, that includes the global filter.~ No it doesn't help.

We might need to move those search param changes to use `useNavigate` instead.


### Comment by @bastilian on August 16, 2024 at 02:32 PM UTC

@xmicha82 Have you had time to revise this issue? If you want we can debug this together and see where it starts loosing the `hash` property. 

---

## Reviews

### Review by @bastilian - Commented on June 17, 2024 at 03:21 PM UTC

### Review by @darkeriossss - Commented on June 18, 2024 at 08:10 AM UTC

### Review by @bastilian - Commented on June 18, 2024 at 10:15 AM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1185*
