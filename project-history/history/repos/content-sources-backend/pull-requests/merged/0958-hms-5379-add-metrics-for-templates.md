---
type: pull_request
number: 958
title: "HMS-5379: add metrics for templates"
state: merged
author: dominikvagner
created: 2025-01-29T14:29:25Z
updated: 2025-02-05T08:11:11Z
closed: 2025-02-05T08:11:10Z
merged: 2025-02-05T08:11:10Z
base_branch: main
head_branch: 5379
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/958
---

# Pull Request #958: HMS-5379: add metrics for templates

**Author**: @dominikvagner
**Created**: January 29, 2025 at 02:29 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `5379`

## Description

## Summary
This PR adds new Prometheus metrics that are gonna be used in our grafana dashboard.
```python
# HELP content_sources_templates_age_average Average age (days between the set template date and now) of templates.
content_sources_templates_age_average 
# HELP content_sources_templates_count Total number of templates.
content_sources_templates_count 
# HELP content_sources_templates_updated_in_last_24_hour_count Number of templates that have been updated in the last 24 hours.
content_sources_templates_updated_in_last_24_hour_count 
# HELP content_sources_templates_use_date_count Number of templates which have a set date.
content_sources_templates_use_date_count 
# HELP content_sources_templates_use_latest_count Number of templates which are set to use latest.
content_sources_templates_use_latest_count 
# HELP content_sources_task_pending_time_average_by_type Average pending time of the tasks by their type.
content_sources_task_pending_time_average_by_type{task_type="add-uploads-repository"}
content_sources_task_pending_time_average_by_type{task_type="delete-repository-snapshots"}
content_sources_task_pending_time_average_by_type{task_type="delete-snapshots"}
content_sources_task_pending_time_average_by_type{task_type="delete-templates"}
content_sources_task_pending_time_average_by_type{task_type="introspect"}
content_sources_task_pending_time_average_by_type{task_type="snapshot"}
content_sources_task_pending_time_average_by_type{task_type="update-latest-snapshot"}
content_sources_task_pending_time_average_by_type{task_type="update-repository"}
content_sources_task_pending_time_average_by_type{task_type="update-template-content"}
```
## Testing steps
1. Verify that the newly added metrics are showing up on our  `:9000/metrics` endpoint.
2. Create some templates which are using latest (at least 1) and have set dates (at least two, with different dates).
3. Verify that the `*_count` metrics have increased/changed accordingly.
4. Verify that the `content_sources_templates_age_average` is the average age of templates you created.
5. Run `make repos-import` and nightly jobs.
6. Wait for the pending time metrics to show up in the output of the metrics output.

\
_(tip: if you have the `bat` tool installed, you can pipe the output from that endpoint into `bat` and when you specify the language as Python, it will be a lot more readable, ex.: `http :9000/metrics | bat --language=python -P -p`)_ 😅 

---

## Discussion

### Comment by @jlsherrill on January 29, 2025 at 02:30 PM UTC

https://issues.redhat.com/browse/HMS-5379

---

## Reviews

### Review by @rverdile - Commented on January 31, 2025 at 08:17 PM UTC

### Review by @rverdile - Commented on January 31, 2025 at 08:22 PM UTC

### Review by @dominikvagner - Commented on February 04, 2025 at 11:52 AM UTC

### Review by @dominikvagner - Commented on February 04, 2025 at 11:53 AM UTC

### Review by @rverdile - Commented on February 04, 2025 at 04:41 PM UTC

### Review by @rverdile - Commented on February 04, 2025 at 04:41 PM UTC

### Review by @rverdile - Approved on February 04, 2025 at 09:29 PM UTC

nice job!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/958*
