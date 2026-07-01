---
type: pull_request
number: 388
title: "Add IQE_ENV=ephemeral"
state: merged
author: swadeley
created: 2023-09-13T09:06:55Z
updated: 2025-09-08T07:30:25Z
closed: 2023-09-13T12:14:51Z
merged: 2023-09-13T12:14:51Z
base_branch: main
head_branch: swadeley/add_export_env_ephemeral
labels: []
url: https://github.com/content-services/content-sources-backend/pull/388
---

# Pull Request #388: Add IQE_ENV=ephemeral

**Author**: @swadeley
**Created**: September 13, 2023 at 09:06 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `swadeley/add_export_env_ephemeral`

## Description

  and fix doc link

## Summary

One test[1] is failing[2] because it is not supported in ephemeral env, so I want to add this env variable.

## Testing steps

pr checks pass[3]

[1] https://gitlab.cee.redhat.com/insights-qe/iqe-content-sources-plugin/-/blob/master/iqe_content_sources/tests/test_repository_api_only.py#L613

[2] https://ci.int.devshift.net/job/content-sources-backend-pr-check-main/359/testReport/junit/tests/test_repository_api_only/test_introspection_of_persistent_user/

[3] https://ci.int.devshift.net/job/content-sources-backend-pr-check-main/


---

## Discussion

### Comment by @swadeley on September 13, 2023 at 09:25 AM UTC

Hi

watching the logs[1], I see  `--env ephemeral` 

The job that failed had `--env clowder_smoke`, not sure if this idea will work.

`10:24:09 Finished: SUCCESS`

build [2] is green !

[1] https://ci.int.devshift.net/job/content-sources-backend-pr-check-main/371/console
[2] https://ci.int.devshift.net/job/content-sources-backend-pr-check-main/371/

---

## Reviews

### Review by @jlsherrill - Approved on September 13, 2023 at 12:14 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/388*
