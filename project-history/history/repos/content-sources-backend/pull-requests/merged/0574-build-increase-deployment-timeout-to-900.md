---
type: pull_request
number: 574
title: "Build: increase deployment timeout to 900"
state: merged
author: swadeley
created: 2024-02-13T10:19:03Z
updated: 2025-09-08T07:30:42Z
closed: 2024-02-13T17:40:56Z
merged: 2024-02-13T17:40:56Z
base_branch: main
head_branch: swadeley/increase_deployment_timeout
labels: []
url: https://github.com/content-services/content-sources-backend/pull/574
---

# Pull Request #574: Build: increase deployment timeout to 900

**Author**: @swadeley
**Created**: February 13, 2024 at 10:19 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `swadeley/increase_deployment_timeout`

## Description

## Summary

In smoke tests I see `DEPLOY_TIMEOUT:-"600` coming from appsre:

`$ bonfire deploy ${APP_NAME} ${DEPLOY_ARGS} --source=appsre --ref-env ${IQE_REF_ENV} --namespace ${NAMESPACE} --timeout ${DEPLOY_TIMEOUT:-"600"} ${FRONTEND_ARG}`

but deploy time seem too short:
`2024-02-12 20:48:10 [   ERROR] [          thread-358] [clowdjobinvocation/create-contentsources-user] timed out waiting for resource to be ready, details:   clowdjobinvocation/create-contentsources-user not ready, status conditions:`

Frontend already has `export DEPLOY_TIMEOUT="900"  # 15min` so I will try that for backend.

## Testing steps
pr checks pass



---

## Reviews

### Review by @jlsherrill - Approved on February 13, 2024 at 03:46 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/574*
