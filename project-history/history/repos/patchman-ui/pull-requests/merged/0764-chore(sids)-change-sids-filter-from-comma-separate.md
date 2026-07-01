---
type: pull_request
number: 764
title: "chore(SIDs): change SIDs filter from comma separated values to single\u2026"
state: merged
author: mkholjuraev
created: 2022-04-07T10:12:57Z
updated: 2022-05-17T08:50:46Z
closed: 2022-04-13T20:22:38Z
merged: 2022-04-13T20:22:38Z
base_branch: master
head_branch: fix-sid
labels: ["released"]
url: https://github.com/RedHatInsights/patchman-ui/pull/764
---

# Pull Request #764: chore(SIDs): change SIDs filter from comma separated values to single…

**Author**: @mkholjuraev
**Created**: April 07, 2022 at 10:12 AM UTC
**Status**: Merged
**Labels**: `released`
**Base**: `master` ← **Head**: `fix-sid`

## Description

Using the shared [generateFilter](https://github.com/RedHatInsights/patchman-ui/commit/efa5df4cbaf8caac2805f417f526b2d53c36617d#diff-4966f4e1c143ce9542deb4f826c300c57035c4fc796c615b9dd65aad9c3ce09bR250) caused SIDs filter to change to comma separated multivalue filter (`filter[system_profile][sid_filter][in]=a,v`). The engine team asked to avoid this pattern due to perfomance issues. This PR brings back original single value filters pattern: `filter[system_profile][sap_sids][in]=AV1&filter[system_profile][sap_sids][in]=AV2&limit=1&offset=0`

---

## Discussion

### Comment by @codecov-commenter on April 07, 2022 at 10:26 AM UTC

# [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/764?src=pr&el=h1&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) Report
> Merging [#764](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/764?src=pr&el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (6038efb) into [master](https://codecov.io/gh/RedHatInsights/patchman-ui/commit/fb7229cb31b4fe507c0a7fb524d3117085f25f45?el=desc&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) (fb7229c) will **decrease** coverage by `0.02%`.
> The diff coverage is `40.00%`.

> :exclamation: Current head 6038efb differs from pull request most recent head 47b487a. Consider uploading reports for the commit 47b487a to get more accurate results

```diff
@@            Coverage Diff             @@
##           master     #764      +/-   ##
==========================================
- Coverage   71.15%   71.12%   -0.03%     
==========================================
  Files         101      101              
  Lines        2371     2372       +1     
  Branches      608      609       +1     
==========================================
  Hits         1687     1687              
  Misses        622      622              
- Partials       62       63       +1     
```


| [Impacted Files](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/764?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights) | Coverage Δ | |
|---|---|---|
| [src/Utilities/Helpers.js](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/764/diff?src=pr&el=tree&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights#diff-c3JjL1V0aWxpdGllcy9IZWxwZXJzLmpz) | `76.71% <40.00%> (-0.27%)` | :arrow_down: |

------

[Continue to review full report at Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/764?src=pr&el=continue&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).
> **Legend** - [Click here to learn more](https://docs.codecov.io/docs/codecov-delta?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights)
> `Δ = absolute <relative> (impact)`, `ø = not affected`, `? = missing data`
> Powered by [Codecov](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/764?src=pr&el=footer&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Last update [fb7229c...47b487a](https://codecov.io/gh/RedHatInsights/patchman-ui/pull/764?src=pr&el=lastupdated&utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights). Read the [comment docs](https://docs.codecov.io/docs/pull-request-comments?utm_medium=referral&utm_source=github&utm_content=comment&utm_campaign=pr+comments&utm_term=RedHatInsights).


### Comment by @mkholjuraev on April 11, 2022 at 10:13 AM UTC

@johnsonm325 thanks for the review. I am working on a fix.

### Comment by @mkholjuraev on April 11, 2022 at 11:53 AM UTC

@johnsonm325 I have found out why this is happening. It is because patch workloads filter syntax is different from Global filter syntax in the URL. Patch URL syntax is` filter[system_profile][sap_sids][in]=AV1` while Global filter syntax is `#SIDs=AC1`. This is causing Global filter no to trigger when you enter `filter[system_profile][sap_sids][in]=AV1` manually. If I add `&SIDs=AC1#SIDs=AC1&tags=` in the URL Global filter gets triggered. 

The fix can be adding this part of url from global filter into the url not from Patch SIDs filter. But this will break SIDs filter in Patch app... I am ready to jump to a short call, if you want.

### Comment by @mkholjuraev on April 11, 2022 at 12:50 PM UTC

Hopefully, the issue is fixed. I have added global filter url syntax (`#sids=sid1,sid2`) together with Patch URL syntax. Now Global filter is getting triggered when user copies the URL and opens new page with copied URL link or refreshes the current page with the manually entered SIDs filter.

### Comment by @johnsonm325 on April 12, 2022 at 01:50 PM UTC

@mkholjuraev There's a new bug. If you select SAP workload in global filters, you get a 'Something went wrong' error. Here's the error in the console.
![image](https://user-images.githubusercontent.com/15373136/162977391-77e29b70-7467-4945-85f3-569965d4d05a.png)


### Comment by @mkholjuraev on April 12, 2022 at 01:57 PM UTC

Ups, I will fix just a sec

### Comment by @johnsonm325 on April 12, 2022 at 03:33 PM UTC

@mkholjuraev if you add `&filter[system_profile][sap_system]=true` in the url and load it, it does not add the SAP workload filter. I think you'll need to add `&workloads=`

### Comment by @mkholjuraev on April 12, 2022 at 04:11 PM UTC

@johnsonm325 I think that this needs to be discussed with backend to change the global filtering syntax from their side. I would suggest that I will create a new bug ticket in the Jira board. 

When you enter `&workloads=`, it triggers the imported global filter, but does nothing from the API side. When you enter `&filter[system_profile][sap_system]=true` in the URL and load, it does nothing on imported global filter and also on the API side.

Let me know your feedback.

### Comment by @johnsonm325 on April 12, 2022 at 05:36 PM UTC

These workloads confuse me. I need an education in how they work. If I hit the following url in drift, `https://console.stage.redhat.com/insights/drift#workloads=SAP&SIDs=&tags=`, the SAP workload is set and when I load systems, I only get SAP systems.

### Comment by @mkholjuraev on April 12, 2022 at 09:23 PM UTC

Lets have a short call tomorrow to discuss our findings if you do not mind.

### Comment by @mkholjuraev on April 13, 2022 at 03:02 PM UTC

@johnsonm325 as we have discussed, I am merging this PR if you do not mind. I have created a new ticket for the global filter issue: https://issues.redhat.com/browse/SPM-1446

### Comment by @mkholjuraev on April 13, 2022 at 08:22 PM UTC

Thanks for the review!

### Comment by @jiridostal on April 14, 2022 at 01:45 PM UTC

:tada: This PR is included in version 1.45.1 :tada:

The release is available on:
- [GitHub release](https://github.com/RedHatInsights/patchman-ui/releases/tag/v1.45.1)
- [npm package (@latest dist-tag)](https://www.npmjs.com/package/@redhat-cloud-services/frontend-components-inventory-patchman/v/1.45.1)

Your **[semantic-release](https://github.com/semantic-release/semantic-release)** bot :package::rocket:

---

## Reviews

### Review by @johnsonm325 - Changes Requested on April 08, 2022 at 10:57 PM UTC

I'm only seeing one problem, and I'm not sure what the expected action is. The url is correctly set on global filter update, but if you set the global filters in the url manually, it does not update the filters in the UI.

For example, if I have no filters set in the UI, and then I paste the following in to the url, `https://stage.foo.redhat.com:1337/insights/patch/advisories?page=1&page_size=20&offset=0&sort=-applicable_systems&filter[system_profile][sap_sids][in]=AV1#SIDs=&tags=` and press enter, the global filters will not be set in the UI.

Drift actually does this correctly, though I don't think drift follows the correct url pattern. The problem is, I can't remember how it was done. I'm either going to look more tonight or take another look on Monday. Here's where it _might_ be happening: https://github.com/RedHatInsights/drift-frontend/blob/master/src/App.js#L86-L92

Let me know if you think we should fix this in another PR and I'll approve this one.

### Review by @johnsonm325 - Approved on April 13, 2022 at 04:02 PM UTC

Yes, this is good.

---

*Archived from: https://github.com/RedHatInsights/patchman-ui/pull/764*
