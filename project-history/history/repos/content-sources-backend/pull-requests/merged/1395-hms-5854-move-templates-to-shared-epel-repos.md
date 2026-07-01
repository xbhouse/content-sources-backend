---
type: pull_request
number: 1395
title: "HMS-5854: move templates to shared EPEL repos"
state: merged
author: xbhouse
created: 2026-02-19T21:49:03Z
updated: 2026-02-25T18:32:51Z
closed: 2026-02-25T18:32:51Z
merged: 2026-02-25T18:32:51Z
base_branch: main
head_branch: 5854
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1395
---

# Pull Request #1395: HMS-5854: move templates to shared EPEL repos

**Author**: @xbhouse
**Created**: February 19, 2026 at 09:49 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `5854`

## Description

## Summary

Adds a job to migrate templates with custom EPEL repos to shared EPEL repos

## Testing steps

1. Set `allow_custom_epel_creation.enabled` to `true` and create a custom EPEL repo via the API. Alternatively you can also set `community_repos.enabled` to `false` and add the custom EPEL repos from the popular repos tab in the UI.  

2. Import and snapshot the shared EPEL repos and the BaseOS and AppStream repos for x86_64:

```
OPTIONS_REPOSITORY_IMPORT_FILTER=epel10,epel9,epel8,base go run cmd/external-repos/main.go import && go run cmd/external-repos/main.go process-repos
```

3. Create a few latest and snapshot date templates with custom EPEL in different orgs, some with and without other custom repos

4. Assign one of the templates to a RHEL system following these [docs](https://github.com/content-services/content-sources-backend/blob/main/docs/register_client.md#registering-a-subscription-manager-client-to-your-local-environment). In /etc/yum.repos.d/redhat.repo, the baseurl should be pointing to the custom EPEL content: `baseurl = http://pulp.content:8081/pulp/content/cs-eb6e57cae0/templates/0375b08a-bcda-4928-ac46-98ee35d6c46e/1aba0dd8-9f36-4c82-bc11-dee0f6144b24`

5. Run the job: `go run cmd/jobs/main.go migrate-templates-to-shared-epel`

6. The update-template-content tasks should complete and all templates should be migrated to shared EPEL 

7. Run `subscription-manager refresh` on the RHEL system, the baseurl for EPEL should be pointing to the shared EPEL content: `baseurl = http://pulp.content:8081/pulp/content/cs-community/templates/0375b08a-bcda-4928-ac46-98ee35d6c46e/efa85b5c-258b-4f9d-9560-1eaf1a55e768`

---

## Discussion

### Comment by @xbhouse on February 19, 2026 at 10:00 PM UTC

https://issues.redhat.com/browse/HMS-5854

### Comment by @xbhouse on February 24, 2026 at 01:52 PM UTC

> One small comment. Other than that I think this looks good! Do you have templates with custom epel repos in stage that you can use to test that the job worked?

thank you! yep i do

### Comment by @xbhouse on February 25, 2026 at 06:32 PM UTC

merging! will keep an eye on it in stage

---

## Reviews

### Review by @rverdile - Commented on February 23, 2026 at 09:34 PM UTC

One small comment. Other than that I think this looks good! Do you have templates with custom epel repos in stage that you can use to test that the job worked?

### Review by @xbhouse - Commented on February 24, 2026 at 01:36 PM UTC

### Review by @rverdile - Approved on February 25, 2026 at 03:52 PM UTC

nice job!!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1395*
