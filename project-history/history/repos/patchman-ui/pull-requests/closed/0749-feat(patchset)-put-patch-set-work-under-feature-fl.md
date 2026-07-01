---
type: pull_request
number: 749
title: "feat(PatchSet): put patch set work under feature flag"
state: closed
author: mkholjuraev
created: 2022-03-11T12:20:30Z
updated: 2022-03-23T12:50:10Z
closed: 2022-03-23T12:50:10Z
base_branch: master
head_branch: feature-flag
labels: []
url: https://github.com/RedHatInsights/patchman-ui/pull/749
---

# Pull Request #749: feat(PatchSet): put patch set work under feature flag

**Author**: @mkholjuraev
**Created**: March 11, 2022 at 12:20 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `master` ← **Head**: `feature-flag`

## Description

This PR is intended to put patch set work under a feature flag. This feature is inteded to be released into prod in Q2. 

**activateFeatureFlag** is used in App.js once to enable flags using URL search params and environment substring. The function calculates what features are enabled and stores them in localStorage. Later **checkEnabledFeature** can be used in any component down in the DOM to check a single feature. 

 All features are by default enabled in development and stage environments. So that we do not have to enter url search param everytime if we want a certain feature. I do not see any reason to hide a feature in these environments.

---

## Discussion

### Comment by @bastilian on March 14, 2022 at 01:09 PM UTC

@mkholjuraev How can the feature(s) be disabled?

---

## Reviews

### Review by @johnsonm325 - Changes Requested on March 14, 2022 at 07:34 PM UTC

@mkholjuraev Check out my suggestions and let me know if I'm misunderstanding how this all works.

### Review by @mkholjuraev - Commented on March 14, 2022 at 08:07 PM UTC

### Review by @mkholjuraev - Commented on March 14, 2022 at 08:11 PM UTC

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/749*
