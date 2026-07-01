---
type: pull_request
number: 721
title: "fix(InventoryDetail): Changed Alert to inline"
state: merged
author: AsToNlele
created: 2022-02-23T13:50:55Z
updated: 2022-03-08T12:06:06Z
closed: 2022-03-08T11:18:58Z
merged: 2022-03-08T11:18:58Z
base_branch: master
head_branch: ADVISOR-1964
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/721
---

# Pull Request #721: fix(InventoryDetail): Changed Alert to inline

**Author**: @AsToNlele
**Created**: February 23, 2022 at 01:50 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `ADVISOR-1964`

## Description

This ticket is in a wrong project: ADVISOR-1964

New look
![image](https://user-images.githubusercontent.com/20592948/155332419-1aa5c8b8-4dc4-4204-9655-5056d8d2710f.png)


---

## Discussion

### Comment by @leSamo on March 03, 2022 at 03:39 PM UTC

I believe there should also be a md (16px) spacer above the alert, now it looks really crammed

### Comment by @AsToNlele on March 04, 2022 at 02:02 PM UTC

You are right, now it looks much better
![image](https://user-images.githubusercontent.com/20592948/156776818-02edd047-9116-47fa-bb27-421f221bfbb6.png)

### Comment by @jiridostal on March 08, 2022 at 12:05 PM UTC

:tada: This PR is included in version 1.40.1 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.40.1)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.40.1)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @MichalSajdik - Approved on February 24, 2022 at 07:49 AM UTC

### Review by @leSamo - Approved on March 07, 2022 at 10:45 AM UTC

Perfect, thank you for the additional fix and for providing pictures in your PR!

One tiny nitpick: instead of `variant="info"` you could use `variant={AlertVariant.info}` to make it more clear, but this is just a heads up into the future :grinning: 

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/721*
