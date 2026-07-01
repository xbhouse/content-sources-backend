---
type: pull_request
number: 249
title: "Move integration tests from ClowdJobInvocation to Job"
state: merged
author: tpapaioa
created: 2023-04-11T13:00:41Z
updated: 2023-04-11T13:43:03Z
closed: 2023-04-11T13:42:58Z
merged: 2023-04-11T13:42:58Z
base_branch: main
head_branch: use_job_based_post_deploy_integration_test
labels: []
url: https://github.com/content-services/content-sources-backend/pull/249
---

# Pull Request #249: Move integration tests from ClowdJobInvocation to Job

**Author**: @tpapaioa
**Created**: April 11, 2023 at 01:00 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `use_job_based_post_deploy_integration_test`

## Description

## Summary
In order to be able to run UI tests against stage, we've moved the integration testing deployments to a cluster that has VPN access. That cluster doesn't support clowder, so this PR changes the corresponding template to use `Job` instead of `ClowdJobInvocation`.

## Testing steps
None

---

## Reviews

### Review by @jlsherrill - Approved on April 11, 2023 at 01:03 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/249*
