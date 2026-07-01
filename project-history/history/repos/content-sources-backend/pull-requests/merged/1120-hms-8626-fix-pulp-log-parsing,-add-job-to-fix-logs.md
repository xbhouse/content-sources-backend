---
type: pull_request
number: 1120
title: "HMS-8626: fix pulp log parsing, add job to fix logs"
state: merged
author: jlsherrill
created: 2025-06-03T02:06:55Z
updated: 2025-06-03T20:35:36Z
closed: 2025-06-03T20:35:13Z
merged: 2025-06-03T20:35:13Z
base_branch: main
head_branch: 8626
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/1120
---

# Pull Request #1120: HMS-8626: fix pulp log parsing, add job to fix logs

**Author**: @jlsherrill
**Created**: June 03, 2025 at 02:06 AM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `8626`

## Description

## Summary

* Adds the ability for arguments to be passed to jobs (most all jobs ignore this)
* Fixes the log parsing for pulp to handle the new format (example included in tests)
* Adds new arguments to the ./jobs transform-pulp-logs DATE COUNT
  * DATE 2025-01-01 is the date to start processing of logs
  * COUNT is the number of days to process
* Adds a stage and prod job to re-run these from May the 5th, through the end of may

## Testing steps



---

## Discussion

### Comment by @jlsherrill on June 03, 2025 at 02:30 AM UTC

https://issues.redhat.com/browse/HMS-8626

### Comment by @jlsherrill on June 03, 2025 at 08:35 PM UTC

no-qe-needed (fixed task), merging

---

## Reviews

### Review by @copilot-pull-request-reviewer - Commented on June 03, 2025 at 02:08 AM UTC

## Pull Request Overview

This PR fixes pulp log parsing by updating the log format handling and adds support for passing arguments to jobs so that logs can be processed for a specified date range. It also updates error handling for log processing and adjusts job function signatures and deployment configurations to support these changes.
- Updated test expectations in the pulp logs test to match the new log format.
- Modified the TransformPulpLogs job to accept arguments and improved error handling.
- Adjusted job function signatures and deployment configurations to pass arguments to jobs.

### Reviewed Changes

Copilot reviewed 10 out of 10 changed files in this pull request and generated 1 comment.

<details>
<summary>Show a summary per file</summary>

| File                                    | Description                                               |
| --------------------------------------- | --------------------------------------------------------- |
| pkg/jobs/transform_pulp_logs_test.go     | Updated expected log formats and test assertions          |
| pkg/jobs/transform_pulp_logs.go          | Updated function signature, added argument parsing, and improved error handling |
| pkg/jobs/set_domain_label.go              | Updated function signature to accept arguments            |
| pkg/jobs/retry_failed_tasks.go            | Updated function signature to accept arguments            |
| pkg/jobs/create_latest_distributions.go   | Updated function signature to accept arguments            |
| pkg/jobs/cleanup_missing_domains.go        | Updated function signature to accept arguments            |
| deployments/jobs-stage.yaml              | Added ClowdJobInvocation deployment for the new job         |
| deployments/jobs-prod.yaml               | Added ClowdJobInvocation deployment for the new job         |
| deployments/deployment.yaml                | Updated deployment to include the new job trigger          |
| cmd/jobs/main.go                         | Adjusted job function type and argument passing mechanism   |
</details>






### Review by @jlsherrill - Commented on June 03, 2025 at 02:20 AM UTC

### Review by @copilot-pull-request-reviewer - Commented on June 03, 2025 at 02:52 AM UTC

## Pull Request Overview

This PR fixes the pulp log parsing to handle a new log format and updates job function signatures to allow passing arguments. It also adds new job invocations for both stage and production deployments.

### Reviewed Changes

Copilot reviewed 10 out of 10 changed files in this pull request and generated 2 comments.

<details>
<summary>Show a summary per file</summary>

| File                                   | Description                                                      |
| -------------------------------------- | ---------------------------------------------------------------- |
| pkg/jobs/transform_pulp_logs_test.go   | Updated test log entries to match the new pulp log format.       |
| pkg/jobs/transform_pulp_logs.go        | Modified TransformPulpLogs to accept arguments and improved error handling. |
| pkg/jobs/set_domain_label.go            | Refactored function signature to accept arguments for consistency.|
| pkg/jobs/retry_failed_tasks.go          | Updated job function signature to accept arguments.            |
| pkg/jobs/create_latest_distributions.go | Updated job function signature for consistency.                  |
| pkg/jobs/cleanup_missing_domains.go    | Updated job function signature for consistency.                  |
| deployments/*.yaml                     | Added new job invocations with proper configuration.             |
| cmd/jobs/main.go                        | Updated job function type and argument passing in main.          |
</details>






### Review by @rverdile - Commented on June 03, 2025 at 03:19 PM UTC

### Review by @jlsherrill - Commented on June 03, 2025 at 03:35 PM UTC

### Review by @jlsherrill - Commented on June 03, 2025 at 03:37 PM UTC

### Review by @rverdile - Commented on June 03, 2025 at 03:41 PM UTC

### Review by @rverdile - Approved on June 03, 2025 at 03:42 PM UTC

code looks good to me

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1120*
