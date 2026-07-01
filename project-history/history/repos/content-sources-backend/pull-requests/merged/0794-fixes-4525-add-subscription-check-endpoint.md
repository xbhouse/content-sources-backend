---
type: pull_request
number: 794
title: "Fixes 4525: add subscription check endpoint"
state: merged
author: rverdile
created: 2024-08-30T14:59:52Z
updated: 2024-09-10T17:45:37Z
closed: 2024-09-10T17:45:34Z
merged: 2024-09-10T17:45:33Z
base_branch: main
head_branch: subscription-check
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/794
---

# Pull Request #794: Fixes 4525: add subscription check endpoint

**Author**: @rverdile
**Created**: August 30, 2024 at 02:59 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `subscription-check`

## Description

## Summary
Adds an endpoint to check if the user has the RHEL for x86_64 product. Will be used for gating content template creation in the UI.

## Testing steps
**For local testing**

1. Turn off "devel_org" in the config.yaml
2. `GET http://localhost:8000/api/content-sources/v1.0/subscription_check/` should return false
3. Turn on "devel_org"
4. `GET http://localhost:8000/api/content-sources/v1.0/subscription_check/` should return true
5. Running it again, you should see a faster response time from the caching



---

## Discussion

### Comment by @jlsherrill on August 30, 2024 at 03:00 PM UTC

https://issues.redhat.com/browse/HMS-4525

### Comment by @rverdile on September 10, 2024 at 05:06 PM UTC

/retest

---

## Reviews

### Review by @xbhouse - Approved on September 09, 2024 at 06:16 PM UTC

this looks good to me! just needs a rebase, that might fix the unit tests too

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/794*
