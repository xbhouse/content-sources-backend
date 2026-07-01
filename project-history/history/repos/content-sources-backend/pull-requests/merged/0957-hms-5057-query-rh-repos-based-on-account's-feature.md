---
type: pull_request
number: 957
title: "HMS-5057: query RH repos based on account's features"
state: merged
author: xbhouse
created: 2025-01-28T21:18:21Z
updated: 2025-02-04T13:54:02Z
closed: 2025-02-04T13:54:02Z
merged: 2025-02-04T13:54:02Z
base_branch: main
head_branch: 5057
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/957
---

# Pull Request #957: HMS-5057: query RH repos based on account's features

**Author**: @xbhouse
**Created**: January 28, 2025 at 09:18 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `5057`

## Description

## Summary

- Adds functionality to query the /featureStatus endpoint from the feature service API and cache the org ID and list of features that account has access to
  - Note: account ID for this endpoint == org ID
- This endpoint should only be called in stage. In prod, accounts should only have access to the RHEL repos. In ephemeral, accounts should have access to all layered repos

## Testing steps

1. `makes repos-import` the small and layered repos (you can add a selector on the openshift and HA repos to just import those) or import them all
2. Set `redis.expiration.subscription_check` to `1m` (and maybe we should use a different expiration field for features?) 
3. List the RH repos via the API and the UI using an account that has access to all features (yours should have access, otherwise you can use this org: 17684632). You should see all the RH repos that were imported, including the layered repos
4. List the RH repos via the API and the UI using an account without access (this org does not have the necessary subscriptions: 18080467). You should see no layered repos
5. Fetch a layered RH repo with both accounts. The account with access should be able to fetch the repo, account without access should not
6. Verify you can fetch a non-layered RH repo with both accounts

---

## Discussion

### Comment by @jlsherrill on January 28, 2025 at 09:30 PM UTC

https://issues.redhat.com/browse/HMS-5057

### Comment by @xbhouse on January 30, 2025 at 07:03 PM UTC

@dominikvagner moving this back to draft, a few more changes are needed to control this behavior in different environments

### Comment by @swadeley on January 31, 2025 at 09:37 AM UTC

Hi, IQE API docs have been updated.

### Comment by @xbhouse on January 31, 2025 at 02:38 PM UTC

/retest


---

## Reviews

### Review by @dominikvagner - Approved on January 30, 2025 at 12:21 PM UTC

looks great, works as desired, nice job! 👏🏼🎉 
approved! 🚀 


### Review by @dominikvagner - Approved on January 31, 2025 at 04:36 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/957*
