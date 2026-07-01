---
type: pull_request
number: 1173
title: "HMS-8568: add kessel client integration"
state: merged
author: rverdile
created: 2025-08-12T21:56:37Z
updated: 2025-09-08T14:25:13Z
closed: 2025-09-08T14:25:13Z
merged: 2025-09-08T14:25:13Z
base_branch: main
head_branch: kessel
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/1173
---

# Pull Request #1173: HMS-8568: add kessel client integration

**Author**: @rverdile
**Created**: August 12, 2025 at 09:56 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `kessel`

## Description

## Summary
Adds support for Kessel. Can be enabled for an individual org, account, or user using our feature flags.

Also has an MR to setting the environment variables: https://gitlab.cee.redhat.com/service/app-interface/-/merge_requests/155278

## Testing steps
This can theoretically be tested in ephemeral if you want to try it: https://project-kessel.github.io/docs/building-with-kessel/how-to/migrate-from-rbac-v1-to-v2/#validate-the-changes-in-the-ephemeral-environment

This can also be tested using the mock server (but the mock server won't test auth)

1. enable rbac, add the kessel_server and rbac_url to your config.yaml
Configure the mock permissions in your config, something like
```
mocks:
  rbac:
    # Update with your account number for admin
    user_read_write: ["snapUser", "uses-rbac"]
    # Update with your account number for viewer
    user_read: ["snapUser"]
    user_no_permissions: ["noperms"]
  kessel:
    # User permission categories (applied to all organizations)
    user_read_write: ["snapUser", "uses-kessel"]
    user_read: ["snapUser"]
    user_no_permissions: ["noperms"]
```
2. Enable kessel for the kessel user
```
  kessel:
    enabled: true
    users: ["uses-kessel"]
```
3. Test that requests for the kessel user use kessel, and others use rbac


---

## Discussion

### Comment by @jlsherrill on August 12, 2025 at 10:00 PM UTC

https://issues.redhat.com/browse/HMS-8568

### Comment by @jlsherrill on August 13, 2025 at 06:42 PM UTC

Overall this looks great!  No major comments.  I do wonder if we should cache GetRootWorkspaceID as i don't see that changing at all, but i know they'd prefer we not cache in general so i'm okay holding off and see how slow things are.

### Comment by @rverdile on August 28, 2025 at 01:47 PM UTC

/retest

### Comment by @rverdile on August 28, 2025 at 04:04 PM UTC

Ephemeral is working today so I am going to try and test again there

edit: couldn't fully test it in ephemeral, but could at least test deploying the schema and making the api calls

### Comment by @jlsherrill on August 29, 2025 at 05:34 PM UTC

The only other comment i think is to add a new section to the readme near here:

https://github.com/content-services/content-sources-backend?tab=readme-ov-file#start--stop-mock-for-rbac

for handling the kessel mock setup

### Comment by @rverdile on September 02, 2025 at 03:49 PM UTC

this is ready for another review

### Comment by @rverdile on September 04, 2025 at 01:32 PM UTC

/retest

### Comment by @rverdile on September 05, 2025 at 06:33 PM UTC

Ignoring snyk error because it is a vulnerability only found in the tests of our dependencies. Our app is not actually using this yaml.v3 package version

---

## Reviews

### Review by @jlsherrill - Commented on August 28, 2025 at 09:24 PM UTC

### Review by @rverdile - Commented on August 29, 2025 at 01:17 PM UTC

### Review by @jlsherrill - Commented on August 29, 2025 at 01:42 PM UTC

### Review by @jlsherrill - Commented on August 29, 2025 at 02:01 PM UTC

### Review by @rverdile - Commented on August 29, 2025 at 03:03 PM UTC

### Review by @jlsherrill - Commented on August 29, 2025 at 03:08 PM UTC

### Review by @jlsherrill - Commented on August 29, 2025 at 03:09 PM UTC

### Review by @rverdile - Commented on August 29, 2025 at 03:23 PM UTC

### Review by @rverdile - Commented on August 29, 2025 at 03:23 PM UTC

### Review by @rverdile - Commented on September 02, 2025 at 03:25 PM UTC

### Review by @jlsherrill - Approved on September 04, 2025 at 12:59 AM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1173*
