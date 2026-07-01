---
type: pull_request
number: 755
title: "feat(PatchSet): put patch set work under feature flag"
state: merged
author: mkholjuraev
created: 2022-03-23T12:58:27Z
updated: 2022-04-12T13:14:04Z
closed: 2022-04-12T12:28:12Z
merged: 2022-04-12T12:28:12Z
base_branch: master
head_branch: feature-flag
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/755
---

# Pull Request #755: feat(PatchSet): put patch set work under feature flag

**Author**: @mkholjuraev
**Created**: March 23, 2022 at 12:58 PM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `feature-flag`

## Description

The PR enables using Unleash feature flags provided by the platform to manage path-sets feature. Current version of the Unleash is old, thus we are using only default enablement strategy. It only enables or disables the flag for now. There is a [ticket](https://issues.redhat.com/browse/APPSRE-3449) to upgrade its version. 

In order to enable/disable the flag you would need access to [Unleash instance](https://insights-stage.unleash.devshift.net/#/features/strategies/patch.patch_set) provided by app-sre. You will need to add these lines to your app-sre profile.

- $ref: /teams/insights/roles/insights-stage-unleash-admin.yml
- $ref: /teams/insights/roles/insights-unleash-admin.yml
- $ref: /teams/insights/roles/unleash-proxy.yml

For the sake of simplicity, you can test the PR by simply hard coding the flag value in the code. 

---

## Discussion

### Comment by @mkholjuraev on April 11, 2022 at 11:25 AM UTC

@leSamo thanks for the review. There was an ask to push patch-set work to prod-beta  ASAP from management thus, this PR is obsolete somehow and not important now. I am putting this in draft state, in case it will be needed.

### Comment by @mkholjuraev on April 12, 2022 at 08:36 AM UTC

@leSamo it looks like this PR is still. I have updated the PR. It would be great if you can have a look at it again. These are following cases when patch set work needs to be behind the flag and hidden. Whenver you enable/disable the flag, please do not forget to clean local storage and open Patch application in a new tab.

1. When you hit **Patch set** navigation, it should redirect you to landing page.
2. Systems row actions in Systems page should hide 'Assign to patch set' action.
3. Patch set column should get hidden on Systems page.
4. **Assign to patch sets** button on Systems table toolbar should get hidden.
5. No patch set work should be displayed whether the flag is enabled or disabled.


### Comment by @jiridostal on April 12, 2022 at 01:13 PM UTC

:tada: This PR is included in version 1.45.0 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.45.0)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.45.0)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @leSamo - Commented on April 11, 2022 at 10:51 AM UTC

Thank you for help with the permissions to the unleash plugin, I was able to change the flag.

I'm not sure what should be hidden by disabling the patch_set flag. I assume it is:
- if user navigates to Patch Set page, navigate him to the landing page (I assume that from the code) but that is not working for me, navigating to Patch Set page doesn't redirect me anywhere
- remove Apply to Patch Set and Remove from Patch Set buttons in tables, but even if flag is false there are these buttons in Advisory detail page:
![false](https://user-images.githubusercontent.com/8426204/162724667-38bf2a3a-06a8-4b67-a9bc-65850d70cb6c.png)

Regardless of the flag being on or off clicking on Apply to Patch Set button throws an error.
![patch-set](https://user-images.githubusercontent.com/8426204/162724619-7512c1a3-fe22-40ed-9351-29d0945ed967.gif)

Maybe I'm doing something wrong, but I made sure to check response from `/api/featureflags/` to be what I expect.

### Review by @leSamo - Approved on April 12, 2022 at 12:24 PM UTC

- When you hit Patch set navigation, it should redirect you to landing page.
    - disabled: OK
    - enabled: OK
- Systems row actions in Systems page should hide 'Assign to patch set' action.
    - disabled: OK
    - enabled: OK
- Patch set column should get hidden on Systems page.
    - disabled: OK 
    - enabled: OK
- Assign to patch sets button on Systems table toolbar should get hidden.
    - disabled: OK 
    - enabled: OK
- No patch set work should be displayed whether the flag is enabled or disabled.
    - advisory detail page has no patch set work regardless of the flag state

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/755*
