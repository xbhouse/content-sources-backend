---
type: pull_request
number: 1214
title: "HMS-5988: prevent creation of custom EPEL repo"
state: merged
author: xbhouse
created: 2025-09-23T20:47:55Z
updated: 2025-10-07T14:10:17Z
closed: 2025-10-07T14:10:17Z
merged: 2025-10-07T14:10:17Z
base_branch: main
head_branch: 5988
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/1214
---

# Pull Request #1214: HMS-5988: prevent creation of custom EPEL repo

**Author**: @xbhouse
**Created**: September 23, 2025 at 08:47 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `5988`

## Description

## Summary

- Prevents creating custom EPEL repos
- If the `allow_custom_epel_creation` feature flag is true for some orgs, custom EPEL repos can be created by those

## Testing steps

1. Set `features.community_repos.enabled: true` and `features.allow_custom_epel_creation.enabled: false` in your config
2. Try to create an EPEL repo and a repo with `origin: community`. Both should fail with a 400
3. Do the same with the bulk create API. This should also fail with a 400
4. Add an org ID to the `features.allow_custom_epel_creation.organizations` field and flip `features.allow_custom_epel_creation.enabled`  to true
5. Request to create these repos with that org ID. The repos should be created (or fail with `Repository with this URL already belongs to organization` if you have one already)

---

## Discussion

### Comment by @jlsherrill on September 23, 2025 at 09:00 PM UTC

https://issues.redhat.com/browse/HMS-5988

### Comment by @swadeley on September 24, 2025 at 07:57 AM UTC

Hi,
 `FAILED tests/test_repository_api_only.py::test_bulk_import` because of timeout waiting for repo vaild status.
`>           repo_2.wait_for_status(["Valid"], delay=20, timeout=1200)`

So not related to this PR

### Comment by @swadeley on September 24, 2025 at 07:58 AM UTC

/retest

### Comment by @xbhouse on September 25, 2025 at 12:24 PM UTC

/retest

### Comment by @xbhouse on September 29, 2025 at 05:40 PM UTC

/retest

---

## Reviews

### Review by @rverdile - Commented on September 24, 2025 at 12:33 PM UTC

I think we'd also want to be able to allow creation for certain orgs so that we can test or debug different migration scenarios in stage

### Review by @rverdile - Commented on October 03, 2025 at 03:37 PM UTC

### Review by @rverdile - Commented on October 03, 2025 at 04:27 PM UTC

### Review by @xbhouse - Commented on October 03, 2025 at 04:38 PM UTC

### Review by @xbhouse - Commented on October 03, 2025 at 08:48 PM UTC

### Review by @rverdile - Commented on October 06, 2025 at 05:37 PM UTC

### Review by @rverdile - Approved on October 06, 2025 at 06:13 PM UTC

looks good! nice job

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1214*
