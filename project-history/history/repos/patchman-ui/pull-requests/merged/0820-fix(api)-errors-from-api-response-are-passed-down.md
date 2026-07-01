---
type: pull_request
number: 820
title: "fix(api): errors from api response are passed down to components"
state: merged
author: mkholjuraev
created: 2022-06-15T09:42:18Z
updated: 2024-04-03T09:23:25Z
closed: 2022-06-16T12:24:13Z
merged: 2022-06-16T12:24:13Z
base_branch: master
head_branch: api-error-handling
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/820
---

# Pull Request #820: fix(api): errors from api response are passed down to components

**Author**: @mkholjuraev
**Created**: June 15, 2022 at 09:42 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `api-error-handling`

## Description

The PR is intended to passdown api error response to the smart components. An error response is used to detect error status and handle correspondingly inside Smart components. e.g ([Systems.js](https://github.com/RedHatInsights/patchman-ui/blob/master/src/SmartComponents/Systems/Systems.js#L212) receives error status code and passes that to [ErrorHandler](https://github.com/RedHatInsights/patchman-ui/blob/master/src/PresentationalComponents/Snippets/ErrorHandler.js) ). Furthermore, redundant user checking before an api call is removed.

It also refactors test coverage code for api.js:

- disabled eslint rules are removed
- first test case for JSON response is removed as it is not necessary anymore
- some dirty code cleanup.

---

## Discussion

### Comment by @mkholjuraev on June 16, 2022 at 12:21 PM UTC

@AsToNlele thank you for your quick response!

### Comment by @jiridostal on June 16, 2022 at 12:41 PM UTC

:tada: This PR is included in version 1.48.3 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.48.3)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.48.3)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @AsToNlele - Approved on June 16, 2022 at 12:16 PM UTC

Looks good!

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/820*
