---
type: pull_request
number: 1306
title: "HMS-9771: fix(repos): use default on error from feature service"
state: merged
author: TenSt
created: 2025-11-19T09:39:38Z
updated: 2025-11-21T10:24:40Z
closed: 2025-11-21T10:24:40Z
merged: 2025-11-21T10:24:40Z
base_branch: main
head_branch: stepan/hotfix-use-default-features-on-err-from-get-entitled-features
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1306
---

# Pull Request #1306: HMS-9771: fix(repos): use default on error from feature service

**Author**: @TenSt
**Created**: November 19, 2025 at 09:39 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `stepan/hotfix-use-default-features-on-err-from-get-entitled-features`

## Description

## Summary
This PR adds default option ("RHEL-OS-x86_64") on error from feature service.

## Testing steps



---

## Discussion

### Comment by @xbhouse on November 19, 2025 at 10:30 AM UTC

https://issues.redhat.com/browse/HMS-9771

### Comment by @marusak on November 19, 2025 at 11:51 AM UTC

Thanks! I was talking with Dominik if we have test for features. I would expect this test to still fail but others to succeed. Or do we have only integration test for it? :thinking: (since all tests passed in this PR)

### Comment by @TenSt on November 19, 2025 at 12:06 PM UTC

> Thanks! I was talking with Dominik if we have test for features. I would expect this test to still fail but others to succeed. Or do we have only integration test for it? 🤔 (since all tests passed in this PR)

We don't have tests for features client, but we have tests for other endpoints that uses it. I've added one more to cover the error case as it was missing and updated existing to use the default value instead of an empty one for all the mock calls.

As for the integration tests, I don't think that we can actually test this there as this is a call made by our API to the external service. So when we run integration tests, we can't mock that. We can mock the response from our API to the UI, but it will be the same as a regular success test case - returning repos filtered by just one feature name.

### Comment by @TenSt on November 19, 2025 at 12:09 PM UTC

I can add feature client tests, but I don't see a lot of value in them. The client is very simple and how our other code works with it is what brings value. 

### Comment by @dominikvagner on November 19, 2025 at 01:08 PM UTC

just for information completeness:
if we make this change, then we won't be "alerted" by our app going down when the feature service isn't working (😂), as there will be the base RHEL feature returned, but we do have an integration [test](https://github.com/content-services/content-sources-frontend/blob/d6ec283a8455424f46e3e9668805813983d38ee8/_playwright-tests/Integration/LayeredRepoAccess.spec.ts#L70) that will fail if the correct features aren't returned 👍🏼
so this change should be fine to make 🫡 

---

## Reviews

### Review by @rverdile - Approved on November 20, 2025 at 08:42 PM UTC

Tested locally by pointing feature service to stage and using non-stage orgID in request. looks good!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1306*
