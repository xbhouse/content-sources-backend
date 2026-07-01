---
type: pull_request
number: 1085
title: "HMS-5923: add support for rhel10"
state: merged
author: xbhouse
created: 2025-04-18T19:10:06Z
updated: 2025-05-14T19:27:00Z
closed: 2025-05-14T19:27:00Z
merged: 2025-05-14T19:27:00Z
base_branch: main
head_branch: 5923
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/1085
---

# Pull Request #1085: HMS-5923: add support for rhel10

**Author**: @xbhouse
**Created**: April 18, 2025 at 07:10 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `5923`

## Description

## Summary

- Adds support for RHEL10 (including extensions channel) and adds EPEL10
- Adds a new manifest that includes these repos

## Testing steps

1. Start with a fresh environment to import the new manifest (`make compose-clean make compose-up`)
2. Run `make repos-import && go run cmd/external-repos/main.go process-repos`. All repos including RHEL10 should be introspected and snapshotted successfully
3. List the RH repos via API and you should see RHEL10 repos in the response
4. List the popular repositories and you should EPEL10 in the response
5. Sorting (`/repositories/?sort_by=distribution_versions:desc&origin=red_hat`) and filtering (`/repositories/?origin=red_hat&version=10`) these repos by version should still work as expected
6. You should be able to create a template with RHEL10 repos, assign it to a RHEL10 client (following these [docs](https://github.com/content-services/content-sources-backend/blob/main/docs/register_client.md#registering-a-subscription-manager-client-to-your-local-environment)), and install packages on the client from those repos

---

## Discussion

### Comment by @jlsherrill on April 18, 2025 at 07:30 PM UTC

https://issues.redhat.com/browse/HMS-5923

### Comment by @jlsherrill on May 14, 2025 at 11:41 AM UTC

/retest

### Comment by @jlsherrill on May 14, 2025 at 11:42 AM UTC

one last thing, can you add the appstream and baseos repos to external_repos.json?     

EDIT: actually,  nevermind, this is handled for us!

---

## Reviews

### Review by @rverdile - Commented on April 24, 2025 at 07:22 PM UTC

Tested repositories list with filtering, snapshots list, templates list, and the various content unit searches. Also made a rhel 10 beta system and was able to install rpms from the template. 

Looks good! will revisit when this is ready to merge

### Review by @jlsherrill - Commented on May 13, 2025 at 05:53 PM UTC

### Review by @xbhouse - Commented on May 13, 2025 at 06:36 PM UTC

### Review by @jlsherrill - Approved on May 14, 2025 at 06:01 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1085*
