---
type: pull_request
number: 1177
title: "fix(RHINENG-8948): update request content-type for vulnerability api calls"
state: merged
author: mkholjuraev
created: 2024-03-21T21:09:06Z
updated: 2024-04-03T13:00:34Z
closed: 2024-04-03T12:35:02Z
merged: 2024-04-03T12:35:02Z
base_branch: master
head_branch: rhineng-8948
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/1177
---

# Pull Request #1177: fix(RHINENG-8948): update request content-type for vulnerability api calls

**Author**: @mkholjuraev
**Created**: March 21, 2024 at 09:09 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `rhineng-8948`

## Description

Updates the request content-type from application/json to application/vnd.api+json. After the vulnerability api has upgraded their connexion package, application/json type is no longer supported

---

## Discussion

### Comment by @AsToNlele on April 03, 2024 at 12:04 PM UTC

Can't see why the PR build failed tho, I'm logged in ![image](https://github.com/RedHatInsights/patchman-ui/assets/20592948/51fa3f8d-ecc1-4390-bf5e-c77217c775d2)

Edit: I can see the new one


### Comment by @AsToNlele on April 03, 2024 at 12:05 PM UTC

/retest

### Comment by @mkholjuraev on April 03, 2024 at 01:00 PM UTC

:tada: This PR is included in version 1.67.3 :tada:

The release is available on [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.67.3)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @AsToNlele - Approved on April 03, 2024 at 12:04 PM UTC

LGTM, saw the discussion on slack

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/1177*
