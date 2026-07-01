---
type: pull_request
number: 396
title: "Fixes 2669: Improving parameter descriptions"
state: merged
author: dubewarsagar
created: 2023-09-21T11:26:19Z
updated: 2023-09-27T17:40:02Z
closed: 2023-09-27T17:40:02Z
merged: 2023-09-27T17:40:02Z
base_branch: main
head_branch: enhancing-documentation-parameters-HCCDOC-1605
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/396
---

# Pull Request #396: Fixes 2669: Improving parameter descriptions

**Author**: @dubewarsagar
**Created**: September 21, 2023 at 11:26 AM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `enhancing-documentation-parameters-HCCDOC-1605`

## Description

## Summary
Scope of this PR is to improve parameter descriptions for the operations listed on: 
https://developers.redhat.com/api-catalog/api/repositories#content-operations

I found 2 scenarios where docs are published. One for upstream and another for downstream. Mention of Optional v/s Required is handled differently in each, hence I am not adding that information in each parameter description. 

## Testing steps


---

## Discussion

### Comment by @app-sre-bot on September 21, 2023 at 11:26 AM UTC

Can one of the admins verify this patch?

### Comment by @rverdile on September 21, 2023 at 01:52 PM UTC

Hi, thanks for your PR! You've done the right thing by making changes in the code comments to change the descriptions. You'll also need to run the `make openapi` command, which will regenerate the `openapi.json` file to update it with the changed descriptions.

### Comment by @rverdile on September 21, 2023 at 08:41 PM UTC

You'll also need to rebase again to remove the extra commits that you did not add. That was caused by another issue that has since been fixed. 

### Comment by @jlsherrill on September 26, 2023 at 09:00 PM UTC

https://issues.redhat.com/browse/HMS-2669

---

## Reviews

### Review by @swadeley - Commented on September 22, 2023 at 06:38 AM UTC

### Review by @swadeley - Commented on September 22, 2023 at 06:38 AM UTC

### Review by @swadeley - Commented on September 22, 2023 at 06:39 AM UTC

### Review by @swadeley - Commented on September 22, 2023 at 06:42 AM UTC

### Review by @swadeley - Commented on September 22, 2023 at 06:43 AM UTC

### Review by @swadeley - Commented on September 22, 2023 at 06:44 AM UTC

### Review by @swadeley - Commented on September 22, 2023 at 06:49 AM UTC

### Review by @swadeley - Commented on September 22, 2023 at 06:51 AM UTC

### Review by @swadeley - Commented on September 22, 2023 at 08:22 AM UTC

### Review by @swadeley - Commented on September 22, 2023 at 08:23 AM UTC

### Review by @swadeley - Commented on September 22, 2023 at 08:24 AM UTC

### Review by @swadeley - Commented on September 22, 2023 at 08:31 AM UTC

### Review by @jlsherrill - Commented on September 22, 2023 at 12:58 PM UTC

### Review by @swadeley - Approved on September 27, 2023 at 05:39 PM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/396*
