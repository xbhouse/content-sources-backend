---
type: pull_request
number: 1187
title: "HMS-5829: support for templates with shared EPEL"
state: merged
author: xbhouse
created: 2025-08-27T22:34:01Z
updated: 2025-09-12T12:27:05Z
closed: 2025-09-12T12:27:05Z
merged: 2025-09-12T12:27:05Z
base_branch: main
head_branch: 5829
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/1187
---

# Pull Request #1187: HMS-5829: support for templates with shared EPEL

**Author**: @xbhouse
**Created**: August 27, 2025 at 10:34 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `5829`

## Description

## Summary

Adds support for creating templates with the shared EPEL repos

## Testing steps

1. Import and snapshot either all repos (`make repos-import &&
go run cmd/external-repos/main.go process-repos`) or at least one RH repo and every EPEL repo: 
```
OPTIONS_REPOSITORY_IMPORT_FILTER=epel10 go run cmd/external-repos/main.go import && go run cmd/external-repos/main.go snapshot --url https://dl.fedoraproject.org/pub/epel/10/Everything/x86_64/

OPTIONS_REPOSITORY_IMPORT_FILTER=epel9 go run cmd/external-repos/main.go import && go run cmd/external-repos/main.go snapshot --url https://dl.fedoraproject.org/pub/epel/9/Everything/x86_64/

OPTIONS_REPOSITORY_IMPORT_FILTER=epel8 go run cmd/external-repos/main.go import && go run cmd/external-repos/main.go snapshot --url https://dl.fedoraproject.org/pub/epel/8/Everything/x86_64/
```
2. Set the `community_repos.enabled` flag in your config.yaml to false and add the popular repos in the UI.
3. Set the `community_repos.enabled` flag back to true.
4. Create a template to use the latest content using the UI with this [PR](https://github.com/content-services/content-sources-frontend/pull/628) with one of the shared EPEL repos (in the UI it will have a `Community` label and in the API it will have `origin: community`). The template should be created successfully.
5. Register a RHEL system with that template, following these [steps](https://github.com/content-services/content-sources-backend/blob/main/docs/register_client.md#registering-a-subscription-manager-client-to-your-local-environment).
7. You should see the shared EPEL content on the system with the correct domain (cs-community) in the URL. Rpms from the EPEL repo should be able to be installed.
8. Get another snapshot for one of the shared EPEL repos (waiting for repo changes might take a day or two, so you could force a snapshot by updating the repo URL manually in the db or in the community_repos.json). The template should update to the latest snapshot.
9. Create another template with a shared EPEL repo and set it to use a snapshot date (older than the oldest snapshot).
10. Update the `snapshot_retain_days_limit` in the config.yaml so it will include the oldest snapshot and run the snapshot cleanup (`go run cmd/external-repos/main.go cleanup --type snapshot`).
11. This task should complete successfully and the template with a snapshot date should be updated to the closest available snapshot.


---

## Discussion

### Comment by @jlsherrill on August 27, 2025 at 11:00 PM UTC

https://issues.redhat.com/browse/HMS-5829

### Comment by @swadeley on September 08, 2025 at 09:37 AM UTC

Hi, I see there was a fail in `test_bulk_import` on 3 Sept 2025, 23:03.
I'll try rerun.

### Comment by @swadeley on September 08, 2025 at 09:37 AM UTC

/retest

### Comment by @rverdile on September 10, 2025 at 01:07 PM UTC

not sure about the failed konflux, might try rebasing?

### Comment by @xbhouse on September 10, 2025 at 02:22 PM UTC

/retest

### Comment by @rverdile on September 10, 2025 at 06:24 PM UTC

/retest

---

## Reviews

### Review by @rverdile - Approved on September 10, 2025 at 01:07 PM UTC

no comments, this looks good to me! nice job

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1187*
