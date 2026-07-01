---
type: pull_request
number: 927
title: "HMS-5062: fix del-templ task fail if repo is deleted first"
state: merged
author: dominikvagner
created: 2024-12-19T09:52:00Z
updated: 2025-01-13T00:57:59Z
closed: 2025-01-13T00:57:59Z
merged: 2025-01-13T00:57:59Z
base_branch: main
head_branch: 5062
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/927
---

# Pull Request #927: HMS-5062: fix del-templ task fail if repo is deleted first

**Author**: @dominikvagner
**Created**: December 19, 2024 at 09:52 AM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `5062`

## Description

## Summary
This PR fixes a bug that happens when a repository exists within a template, and both are deleted, but the repository config is deleted before the template is deleted, then the `delete-templates` task gives a "repository not found" error and if this happens the template-repository distribution doesn't get deleted and remains an orphan.

The ticket assumed that the `delete-repository-snapshots` task didn't try deleting the template-repository distribution, but it actually already was trying to do so, but that step just needed to be moved earlier in the task and also do that for soft deleted templates. The `delete-template` task just needed to be changed to ignore not found repository configs.

## Testing steps
1. Create a custom repository and add it to a new template.
2. Note down the template-repository distribution href (look it up by uuids in `SELECT * FROM templates_repository_configurations;`) and verify it exists (`http -a admin:password :8080/${DISTRIBUTION_HREF}`).
3. Change `retry_wait_upper_bound` in `config.yaml` to 3m.
4. Edit the `Run()` method of `delete_templates.go` to force a failure by returning an error (add `return errors.new("test error")` at the beginning).
5. Delete the template.
6. Delete the repository configuration.
7. Once the `delete-repository-snapshot` task finishes, stop the server.
8. Fix the `Run()` from step 3, so it runs correctly, then restart the server.
9. The `delete-templates` tasks will retry and it should be completed successfully.
10. Verify the template-repository distribution was deleted (`http -a admin:password :8080/${DISTRIBUTION_HREF}` -> `404`).


---

## Discussion

### Comment by @jlsherrill on December 20, 2024 at 04:38 PM UTC

https://issues.redhat.com/browse/HMS-5062

### Comment by @swadeley on January 08, 2025 at 05:03 AM UTC

/retest

### Comment by @jlsherrill on January 10, 2025 at 02:20 PM UTC

code looks good, testing now

---

## Reviews

### Review by @jlsherrill - Approved on January 10, 2025 at 02:58 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/927*
