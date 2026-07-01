---
type: pull_request
number: 151
title: "Fixes 267: Add a metrics endpoint"
state: merged
author: avisiedo
created: 2022-11-30T09:57:14Z
updated: 2023-01-13T20:30:55Z
closed: 2023-01-13T16:42:14Z
merged: 2023-01-13T16:42:14Z
base_branch: main
head_branch: hmscontent-267
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/151
---

# Pull Request #151: Fixes 267: Add a metrics endpoint

**Author**: @avisiedo
**Created**: November 30, 2022 at 09:57 AM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `hmscontent-267`

## Description

Currently:
- Provide and endpoint to pull the metrics (instrumentation service).
- Provide local infra to check the metrics from prometheus ui (make prometheus-{up,down,clean,logs,ui}).
- Provide local prometheus configuration to check it locally.
- Update dependencies.
- Fix a wrong behavior for containers in macos systems (healthcheck) when podman and docker are installed into the same system.
- Create a common place to set the metrics to export at `pkg/instrumentation/metrics.go`.

TODO

- [ ] Some metrics seems absolute values for the whole service; how could this impact when every pod export the same metric?
- [ ] If the response to the above point is that it can impact, then study and ask about decompose in some discrete metrics that could be used for a prometheus query that finally report the expected values.
- [x] About the introspection result to use it for calculate repositories not introspected into the last 24 hours, it currently happens in a different pod; now that they could be launched from the same binary, could make sense to use only one pod and start the kafka-consumer from the unique pod? Keeping api and kafka consumer running in different pods.
- [x] Declare in app-interface the metric endpoint for the service.
- [x] Unit tests.
- [ ] Top50 repositories is not implemented on this pr; I have not found a way to add that information with prometheus metric; we store counter, specific value that go up and down, histogram and summaries, but for the 50 repositories, not clear; cross my mind to store the position as the value, and the url repository as a label of that position; as soon we have 50 repositories or more, when a repository is out of the list, the last value scrapped is that remain for that label, which it does not represent the reality.
- [x] unit test for `collector.go` file.

---

## Discussion

### Comment by @jlsherrill on November 30, 2022 at 10:00 AM UTC

https://issues.redhat.com/browse/HMSCONTENT-267

### Comment by @jlsherrill on December 01, 2022 at 02:16 AM UTC

>  About the introspection result to use it for calculate repositories not introspected into the last 24 hours, it currently happens in a different pod; now that they could be launched from the same binary, could make sense to use only one pod and start the kafka-consumer from the unique pod?

I would assume this would be a db query, so it doesn't matter which pod is asked this question?  They would all return the same answer. (but maybe i'm missing something)

### Comment by @avisiedo on December 13, 2022 at 11:14 AM UTC

@jlsherrill @rverdile thanks for the comments! I will update the code based on the comments:
- I will add a go routine to update periodically the metrics from the database.
- I will try to avoid the leverage of DefaultRegisterer.
- I will use the path and port indicated by clowder.

### Comment by @avisiedo on January 03, 2023 at 12:21 PM UTC

I have updated this pr by adding the custom-metrics but the top50 repositories, and their unit tests, but for the collector.

### Comment by @avisiedo on January 04, 2023 at 01:38 PM UTC

/retest

### Comment by @rverdile on January 09, 2023 at 04:04 PM UTC

Tested again and looks like it's working. I think @jlsherrill wanted to review the sql queries and we also need his input on that one remaining question. Otherwise, I think this is just about done :).

### Comment by @avisiedo on January 09, 2023 at 05:34 PM UTC

/retest

### Comment by @avisiedo on January 11, 2023 at 11:57 AM UTC

Test requirements:

```raw
Scenario 1: metrics endpoint is up and listening

Given the service is up and running
    and the metrics endpoint configuration read from clowder
  When the metrics endpoint is reached out (http://content-sources-backend:9000/metrics)
   Then a get a response with the current metrics value and description.

----

Scenario 2: metrics are generated for total number of repositories

Given the service up and running
    and one repository created into the database
    and two repository configurations into the database pointing to the above repo
    and another one public repository introspected beyond 24 hours above in Invalid or Failed
    and another one public repository with Failed state before the 24 hours boundary
    and another one non public repository with state Invalid or Failed beyond 24 hours
  When the metrics endpoint is scrapped
   Then the number of repositories count is 4
     and the number of repository config count is 2
     and the number of public repositories not introspected in the last 24 hours is 1
     and the number of non publis repositories not introspected in the last 24 hours is 1
     and the number of failed repositories is 1 (we have 1 repo bayond 24 hours and 1 before)
     
** I was not able to think a way to get repos in the above states without create the entry directly into the database.

----

Scenario 3: api metrics

Given the service up and running
    and two repositories created with the api or ui
   and one of the above repositories deleted
  When the metrics endpoint is scrapped
   Then we get the metrics count of the api status according to the previous operations

Example of content printed out by `curl http://localhost:9000/metrics`:

# HELP content_sources_http_status_histogram Duration of HTTP requests
# TYPE content_sources_http_status_histogram histogram
content_sources_http_status_histogram_bucket{method="DELETE",path="/api/content-sources/v1/repositories/:uuid",status="2xx",le="0.005"} 0
content_sources_http_status_histogram_bucket{method="DELETE",path="/api/content-sources/v1/repositories/:uuid",status="2xx",le="0.01"} 0
content_sources_http_status_histogram_bucket{method="DELETE",path="/api/content-sources/v1/repositories/:uuid",status="2xx",le="0.025"} 1
content_sources_http_status_histogram_bucket{method="DELETE",path="/api/content-sources/v1/repositories/:uuid",status="2xx",le="0.05"} 1
content_sources_http_status_histogram_bucket{method="DELETE",path="/api/content-sources/v1/repositories/:uuid",status="2xx",le="0.1"} 1
content_sources_http_status_histogram_bucket{method="DELETE",path="/api/content-sources/v1/repositories/:uuid",status="2xx",le="0.25"} 1
content_sources_http_status_histogram_bucket{method="DELETE",path="/api/content-sources/v1/repositories/:uuid",status="2xx",le="0.5"} 1
content_sources_http_status_histogram_bucket{method="DELETE",path="/api/content-sources/v1/repositories/:uuid",status="2xx",le="1"} 1
content_sources_http_status_histogram_bucket{method="DELETE",path="/api/content-sources/v1/repositories/:uuid",status="2xx",le="2.5"} 1
content_sources_http_status_histogram_bucket{method="DELETE",path="/api/content-sources/v1/repositories/:uuid",status="2xx",le="5"} 1
content_sources_http_status_histogram_bucket{method="DELETE",path="/api/content-sources/v1/repositories/:uuid",status="2xx",le="10"} 1
content_sources_http_status_histogram_bucket{method="DELETE",path="/api/content-sources/v1/repositories/:uuid",status="2xx",le="+Inf"} 1
content_sources_http_status_histogram_sum{method="DELETE",path="/api/content-sources/v1/repositories/:uuid",status="2xx"} 0.011591979
content_sources_http_status_histogram_count{method="DELETE",path="/api/content-sources/v1/repositories/:uuid",status="2xx"} 1
content_sources_http_status_histogram_bucket{method="GET",path="/api/content-sources/v1.0/repositories/:uuid/rpms",status="2xx",le="0.005"} 0
content_sources_http_status_histogram_bucket{method="GET",path="/api/content-sources/v1.0/repositories/:uuid/rpms",status="2xx",le="0.01"} 0
content_sources_http_status_histogram_bucket{method="GET",path="/api/content-sources/v1.0/repositories/:uuid/rpms",status="2xx",le="0.025"} 0
content_sources_http_status_histogram_bucket{method="GET",path="/api/content-sources/v1.0/repositories/:uuid/rpms",status="2xx",le="0.05"} 1
content_sources_http_status_histogram_bucket{method="GET",path="/api/content-sources/v1.0/repositories/:uuid/rpms",status="2xx",le="0.1"} 1
content_sources_http_status_histogram_bucket{method="GET",path="/api/content-sources/v1.0/repositories/:uuid/rpms",status="2xx",le="0.25"} 1
content_sources_http_status_histogram_bucket{method="GET",path="/api/content-sources/v1.0/repositories/:uuid/rpms",status="2xx",le="0.5"} 1
content_sources_http_status_histogram_bucket{method="GET",path="/api/content-sources/v1.0/repositories/:uuid/rpms",status="2xx",le="1"} 1
content_sources_http_status_histogram_bucket{method="GET",path="/api/content-sources/v1.0/repositories/:uuid/rpms",status="2xx",le="2.5"} 1
content_sources_http_status_histogram_bucket{method="GET",path="/api/content-sources/v1.0/repositories/:uuid/rpms",status="2xx",le="5"} 1
content_sources_http_status_histogram_bucket{method="GET",path="/api/content-sources/v1.0/repositories/:uuid/rpms",status="2xx",le="10"} 1
content_sources_http_status_histogram_bucket{method="GET",path="/api/content-sources/v1.0/repositories/:uuid/rpms",status="2xx",le="+Inf"} 1
content_sources_http_status_histogram_sum{method="GET",path="/api/content-sources/v1.0/repositories/:uuid/rpms",status="2xx"} 0.03865349
content_sources_http_status_histogram_count{method="GET",path="/api/content-sources/v1.0/repositories/:uuid/rpms",status="2xx"} 1
content_sources_http_status_histogram_bucket{method="GET",path="/api/content-sources/v1/popular_repositories/",status="2xx",le="0.005"} 3
content_sources_http_status_histogram_bucket{method="GET",path="/api/content-sources/v1/popular_repositories/",status="2xx",le="0.01"} 4
content_sources_http_status_histogram_bucket{method="GET",path="/api/content-sources/v1/popular_repositories/",status="2xx",le="0.025"} 4
content_sources_http_status_histogram_bucket{method="GET",path="/api/content-sources/v1/popular_repositories/",status="2xx",le="0.05"} 4
content_sources_http_status_histogram_bucket{method="GET",path="/api/content-sources/v1/popular_repositories/",status="2xx",le="0.1"} 4
content_sources_http_status_histogram_bucket{method="GET",path="/api/content-sources/v1/popular_repositories/",status="2xx",le="0.25"} 4
content_sources_http_status_histogram_bucket{method="GET",path="/api/content-sources/v1/popular_repositories/",status="2xx",le="0.5"} 4
content_sources_http_status_histogram_bucket{method="GET",path="/api/content-sources/v1/popular_repositories/",status="2xx",le="1"} 4
content_sources_http_status_histogram_bucket{method="GET",path="/api/content-sources/v1/popular_repositories/",status="2xx",le="2.5"} 4
content_sources_http_status_histogram_bucket{method="GET",path="/api/content-sources/v1/popular_repositories/",status="2xx",le="5"} 4
content_sources_http_status_histogram_bucket{method="GET",path="/api/content-sources/v1/popular_repositories/",status="2xx",le="10"} 4
content_sources_http_status_histogram_bucket{method="GET",path="/api/content-sources/v1/popular_repositories/",status="2xx",le="+Inf"} 4
content_sources_http_status_histogram_sum{method="GET",path="/api/content-sources/v1/popular_repositories/",status="2xx"} 0.014736989999999998
content_sources_http_status_histogram_count{method="GET",path="/api/content-sources/v1/popular_repositories/",status="2xx"} 4
content_sources_http_status_histogram_bucket{method="GET",path="/api/content-sources/v1/repositories/",status="2xx",le="0.005"} 5
content_sources_http_status_histogram_bucket{method="GET",path="/api/content-sources/v1/repositories/",status="2xx",le="0.01"} 5
content_sources_http_status_histogram_bucket{method="GET",path="/api/content-sources/v1/repositories/",status="2xx",le="0.025"} 5
content_sources_http_status_histogram_bucket{method="GET",path="/api/content-sources/v1/repositories/",status="2xx",le="0.05"} 5
content_sources_http_status_histogram_bucket{method="GET",path="/api/content-sources/v1/repositories/",status="2xx",le="0.1"} 5
content_sources_http_status_histogram_bucket{method="GET",path="/api/content-sources/v1/repositories/",status="2xx",le="0.25"} 5
content_sources_http_status_histogram_bucket{method="GET",path="/api/content-sources/v1/repositories/",status="2xx",le="0.5"} 5
content_sources_http_status_histogram_bucket{method="GET",path="/api/content-sources/v1/repositories/",status="2xx",le="1"} 5
content_sources_http_status_histogram_bucket{method="GET",path="/api/content-sources/v1/repositories/",status="2xx",le="2.5"} 5
content_sources_http_status_histogram_bucket{method="GET",path="/api/content-sources/v1/repositories/",status="2xx",le="5"} 5
content_sources_http_status_histogram_bucket{method="GET",path="/api/content-sources/v1/repositories/",status="2xx",le="10"} 5
content_sources_http_status_histogram_bucket{method="GET",path="/api/content-sources/v1/repositories/",status="2xx",le="+Inf"} 5
content_sources_http_status_histogram_sum{method="GET",path="/api/content-sources/v1/repositories/",status="2xx"} 0.008500505
content_sources_http_status_histogram_count{method="GET",path="/api/content-sources/v1/repositories/",status="2xx"} 5
content_sources_http_status_histogram_bucket{method="GET",path="/api/content-sources/v1/repository_parameters/",status="2xx",le="0.005"} 1
content_sources_http_status_histogram_bucket{method="GET",path="/api/content-sources/v1/repository_parameters/",status="2xx",le="0.01"} 1
content_sources_http_status_histogram_bucket{method="GET",path="/api/content-sources/v1/repository_parameters/",status="2xx",le="0.025"} 1
content_sources_http_status_histogram_bucket{method="GET",path="/api/content-sources/v1/repository_parameters/",status="2xx",le="0.05"} 1
content_sources_http_status_histogram_bucket{method="GET",path="/api/content-sources/v1/repository_parameters/",status="2xx",le="0.1"} 1
content_sources_http_status_histogram_bucket{method="GET",path="/api/content-sources/v1/repository_parameters/",status="2xx",le="0.25"} 1
content_sources_http_status_histogram_bucket{method="GET",path="/api/content-sources/v1/repository_parameters/",status="2xx",le="0.5"} 1
content_sources_http_status_histogram_bucket{method="GET",path="/api/content-sources/v1/repository_parameters/",status="2xx",le="1"} 1
content_sources_http_status_histogram_bucket{method="GET",path="/api/content-sources/v1/repository_parameters/",status="2xx",le="2.5"} 1
content_sources_http_status_histogram_bucket{method="GET",path="/api/content-sources/v1/repository_parameters/",status="2xx",le="5"} 1
content_sources_http_status_histogram_bucket{method="GET",path="/api/content-sources/v1/repository_parameters/",status="2xx",le="10"} 1
content_sources_http_status_histogram_bucket{method="GET",path="/api/content-sources/v1/repository_parameters/",status="2xx",le="+Inf"} 1
content_sources_http_status_histogram_sum{method="GET",path="/api/content-sources/v1/repository_parameters/",status="2xx"} 0.000141457
content_sources_http_status_histogram_count{method="GET",path="/api/content-sources/v1/repository_parameters/",status="2xx"} 1
content_sources_http_status_histogram_bucket{method="PATCH",path="/api/content-sources/v1.0/repositories/:uuid",status="2xx",le="0.005"} 0
content_sources_http_status_histogram_bucket{method="PATCH",path="/api/content-sources/v1.0/repositories/:uuid",status="2xx",le="0.01"} 0
content_sources_http_status_histogram_bucket{method="PATCH",path="/api/content-sources/v1.0/repositories/:uuid",status="2xx",le="0.025"} 1
content_sources_http_status_histogram_bucket{method="PATCH",path="/api/content-sources/v1.0/repositories/:uuid",status="2xx",le="0.05"} 1
content_sources_http_status_histogram_bucket{method="PATCH",path="/api/content-sources/v1.0/repositories/:uuid",status="2xx",le="0.1"} 1
content_sources_http_status_histogram_bucket{method="PATCH",path="/api/content-sources/v1.0/repositories/:uuid",status="2xx",le="0.25"} 1
content_sources_http_status_histogram_bucket{method="PATCH",path="/api/content-sources/v1.0/repositories/:uuid",status="2xx",le="0.5"} 1
content_sources_http_status_histogram_bucket{method="PATCH",path="/api/content-sources/v1.0/repositories/:uuid",status="2xx",le="1"} 1
content_sources_http_status_histogram_bucket{method="PATCH",path="/api/content-sources/v1.0/repositories/:uuid",status="2xx",le="2.5"} 1
content_sources_http_status_histogram_bucket{method="PATCH",path="/api/content-sources/v1.0/repositories/:uuid",status="2xx",le="5"} 1
content_sources_http_status_histogram_bucket{method="PATCH",path="/api/content-sources/v1.0/repositories/:uuid",status="2xx",le="10"} 1
content_sources_http_status_histogram_bucket{method="PATCH",path="/api/content-sources/v1.0/repositories/:uuid",status="2xx",le="+Inf"} 1
content_sources_http_status_histogram_sum{method="PATCH",path="/api/content-sources/v1.0/repositories/:uuid",status="2xx"} 0.01137019
content_sources_http_status_histogram_count{method="PATCH",path="/api/content-sources/v1.0/repositories/:uuid",status="2xx"} 1
content_sources_http_status_histogram_bucket{method="POST",path="/api/content-sources/v1.0/repositories/bulk_create/",status="2xx",le="0.005"} 0
content_sources_http_status_histogram_bucket{method="POST",path="/api/content-sources/v1.0/repositories/bulk_create/",status="2xx",le="0.01"} 0
content_sources_http_status_histogram_bucket{method="POST",path="/api/content-sources/v1.0/repositories/bulk_create/",status="2xx",le="0.025"} 2
content_sources_http_status_histogram_bucket{method="POST",path="/api/content-sources/v1.0/repositories/bulk_create/",status="2xx",le="0.05"} 2
content_sources_http_status_histogram_bucket{method="POST",path="/api/content-sources/v1.0/repositories/bulk_create/",status="2xx",le="0.1"} 2
content_sources_http_status_histogram_bucket{method="POST",path="/api/content-sources/v1.0/repositories/bulk_create/",status="2xx",le="0.25"} 2
content_sources_http_status_histogram_bucket{method="POST",path="/api/content-sources/v1.0/repositories/bulk_create/",status="2xx",le="0.5"} 2
content_sources_http_status_histogram_bucket{method="POST",path="/api/content-sources/v1.0/repositories/bulk_create/",status="2xx",le="1"} 2
content_sources_http_status_histogram_bucket{method="POST",path="/api/content-sources/v1.0/repositories/bulk_create/",status="2xx",le="2.5"} 2
content_sources_http_status_histogram_bucket{method="POST",path="/api/content-sources/v1.0/repositories/bulk_create/",status="2xx",le="5"} 2
content_sources_http_status_histogram_bucket{method="POST",path="/api/content-sources/v1.0/repositories/bulk_create/",status="2xx",le="10"} 2
content_sources_http_status_histogram_bucket{method="POST",path="/api/content-sources/v1.0/repositories/bulk_create/",status="2xx",le="+Inf"} 2
content_sources_http_status_histogram_sum{method="POST",path="/api/content-sources/v1.0/repositories/bulk_create/",status="2xx"} 0.023565302
content_sources_http_status_histogram_count{method="POST",path="/api/content-sources/v1.0/repositories/bulk_create/",status="2xx"} 2
content_sources_http_status_histogram_bucket{method="POST",path="/api/content-sources/v1.0/repository_parameters/validate/",status="2xx",le="0.005"} 0
content_sources_http_status_histogram_bucket{method="POST",path="/api/content-sources/v1.0/repository_parameters/validate/",status="2xx",le="0.01"} 0
content_sources_http_status_histogram_bucket{method="POST",path="/api/content-sources/v1.0/repository_parameters/validate/",status="2xx",le="0.025"} 0
content_sources_http_status_histogram_bucket{method="POST",path="/api/content-sources/v1.0/repository_parameters/validate/",status="2xx",le="0.05"} 0
content_sources_http_status_histogram_bucket{method="POST",path="/api/content-sources/v1.0/repository_parameters/validate/",status="2xx",le="0.1"} 0
content_sources_http_status_histogram_bucket{method="POST",path="/api/content-sources/v1.0/repository_parameters/validate/",status="2xx",le="0.25"} 0
content_sources_http_status_histogram_bucket{method="POST",path="/api/content-sources/v1.0/repository_parameters/validate/",status="2xx",le="0.5"} 0
content_sources_http_status_histogram_bucket{method="POST",path="/api/content-sources/v1.0/repository_parameters/validate/",status="2xx",le="1"} 1
content_sources_http_status_histogram_bucket{method="POST",path="/api/content-sources/v1.0/repository_parameters/validate/",status="2xx",le="2.5"} 2
content_sources_http_status_histogram_bucket{method="POST",path="/api/content-sources/v1.0/repository_parameters/validate/",status="2xx",le="5"} 4
content_sources_http_status_histogram_bucket{method="POST",path="/api/content-sources/v1.0/repository_parameters/validate/",status="2xx",le="10"} 4
content_sources_http_status_histogram_bucket{method="POST",path="/api/content-sources/v1.0/repository_parameters/validate/",status="2xx",le="+Inf"} 4
content_sources_http_status_histogram_sum{method="POST",path="/api/content-sources/v1.0/repository_parameters/validate/",status="2xx"} 8.110344537
content_sources_http_status_histogram_count{method="POST",path="/api/content-sources/v1.0/repository_parameters/validate/",status="2xx"} 4
# HELP content_sources_non_public_repositories_not_introspected_last_24_hours_total Number of non public repositories not introspected in the last 24 hours
# TYPE content_sources_non_public_repositories_not_introspected_last_24_hours_total gauge
content_sources_non_public_repositories_not_introspected_last_24_hours_total 0
# HELP content_sources_public_repositories_not_introspected_last_24_hours_total Number of public repositories not introspected into the last 24 hours
# TYPE content_sources_public_repositories_not_introspected_last_24_hours_total gauge
content_sources_public_repositories_not_introspected_last_24_hours_total 0
# HELP content_sources_public_repositories_with_failed_introspection_total Number of repositories with failed introspection
# TYPE content_sources_public_repositories_with_failed_introspection_total gauge
content_sources_public_repositories_with_failed_introspection_total 0
# HELP content_sources_repositories_total Number of repositories
# TYPE content_sources_repositories_total gauge
content_sources_repositories_total 1
# HELP content_sources_repository_configs_total Number of repository configurations
# TYPE content_sources_repository_configs_total gauge
content_sources_repository_configs_total 1
# HELP go_build_info Build information about the main Go module.
# TYPE go_build_info gauge
go_build_info{checksum="",path="",version=""} 1
# HELP promhttp_metric_handler_errors_total Total number of internal errors encountered by the promhttp metric handler.
# TYPE promhttp_metric_handler_errors_total counter
promhttp_metric_handler_errors_total{cause="encoding"} 0
promhttp_metric_handler_errors_total{cause="gathering"} 0



```

from the above can be derived more specific scenarios as they are considered.

> by scrapping the metrics endpoint it is:
> for the workstation can be used: `curl http://localhost:9000/metrics`
> for the ephemeral environment can be used: `curl http://content-sources-backend-service:9000/metrics`
> not sure if there are some python helpers to scrap the response (I recall to see something for the golang client library).
> not sure what flexibility could allow that helper


### Comment by @swadeley on January 13, 2023 at 04:41 PM UTC

Hi

lgtm


```
(.iqe-env) [sjw@t480s iqe-content-sources-plugin]$ oc rsh -n $NAMESPACE $POD

sh-4.4$  curl http://content-sources-backend-service:9000/metrics
# HELP content_sources_non_public_repositories_not_introspected_last_24_hours_total Number of non public repositories not introspected in the last 24 hours
# TYPE content_sources_non_public_repositories_not_introspected_last_24_hours_total gauge
content_sources_non_public_repositories_not_introspected_last_24_hours_total 0
# HELP content_sources_public_repositories_not_introspected_last_24_hours_total Number of public repositories not introspected into the last 24 hours
# TYPE content_sources_public_repositories_not_introspected_last_24_hours_total gauge
content_sources_public_repositories_not_introspected_last_24_hours_total 0
# HELP content_sources_public_repositories_with_failed_introspection_total Number of repositories with failed introspection
# TYPE content_sources_public_repositories_with_failed_introspection_total gauge
content_sources_public_repositories_with_failed_introspection_total 0
# HELP content_sources_repositories_total Number of repositories
# TYPE content_sources_repositories_total gauge
content_sources_repositories_total 22
# HELP content_sources_repository_configs_total Number of repository configurations
# TYPE content_sources_repository_configs_total gauge
content_sources_repository_configs_total 0
# HELP go_build_info Build information about the main Go module.
# TYPE go_build_info gauge
go_build_info{checksum="",path="",version=""} 1
# HELP promhttp_metric_handler_errors_total Total number of internal errors encountered by the promhttp metric handler.
# TYPE promhttp_metric_handler_errors_total counter
promhttp_metric_handler_errors_total{cause="encoding"} 0
promhttp_metric_handler_errors_total{cause="gathering"} 0
```

---

## Reviews

### Review by @rverdile - Commented on December 12, 2022 at 02:56 PM UTC

### Review by @rverdile - Commented on December 12, 2022 at 03:00 PM UTC

### Review by @rverdile - Commented on December 12, 2022 at 03:24 PM UTC

### Review by @avisiedo - Commented on December 13, 2022 at 09:58 AM UTC

### Review by @avisiedo - Commented on December 13, 2022 at 10:06 AM UTC

### Review by @avisiedo - Commented on December 13, 2022 at 10:11 AM UTC

### Review by @avisiedo - Commented on December 13, 2022 at 11:09 AM UTC

### Review by @rverdile - Commented on December 13, 2022 at 02:31 PM UTC

### Review by @rverdile - Commented on December 13, 2022 at 02:44 PM UTC

### Review by @rverdile - Commented on December 13, 2022 at 03:53 PM UTC

### Review by @rverdile - Commented on December 14, 2022 at 07:26 PM UTC

### Review by @avisiedo - Commented on December 15, 2022 at 11:46 AM UTC

### Review by @rverdile - Commented on December 15, 2022 at 06:35 PM UTC

### Review by @avisiedo - Commented on December 16, 2022 at 01:13 PM UTC

### Review by @rverdile - Commented on December 19, 2022 at 09:31 PM UTC

### Review by @rverdile - Commented on December 19, 2022 at 09:32 PM UTC

### Review by @avisiedo - Commented on December 20, 2022 at 09:03 AM UTC

### Review by @avisiedo - Commented on December 20, 2022 at 12:30 PM UTC

### Review by @jlsherrill - Commented on December 20, 2022 at 09:08 PM UTC

### Review by @avisiedo - Commented on December 21, 2022 at 01:34 PM UTC

### Review by @rverdile - Commented on December 21, 2022 at 07:31 PM UTC

I know you're still working on this, so some of my comments might already be outdated by changes you've made (or plan to make) locally. Some of the comments are also on small things.

I'm just leaving a bunch since I'll be out tomorrow :).

### Review by @avisiedo - Commented on December 22, 2022 at 02:39 PM UTC

### Review by @avisiedo - Commented on December 22, 2022 at 02:58 PM UTC

### Review by @avisiedo - Commented on January 03, 2023 at 01:30 PM UTC

### Review by @avisiedo - Commented on January 03, 2023 at 08:53 PM UTC

### Review by @jlsherrill - Commented on January 03, 2023 at 09:12 PM UTC

### Review by @jlsherrill - Commented on January 03, 2023 at 09:19 PM UTC

### Review by @rverdile - Commented on January 04, 2023 at 03:15 PM UTC

### Review by @rverdile - Commented on January 04, 2023 at 03:35 PM UTC

### Review by @rverdile - Commented on January 04, 2023 at 04:23 PM UTC

### Review by @rverdile - Commented on January 04, 2023 at 04:25 PM UTC

### Review by @avisiedo - Commented on January 04, 2023 at 05:54 PM UTC

### Review by @avisiedo - Commented on January 04, 2023 at 06:21 PM UTC

### Review by @avisiedo - Commented on January 04, 2023 at 08:42 PM UTC

### Review by @avisiedo - Commented on January 04, 2023 at 08:56 PM UTC

### Review by @rverdile - Commented on January 05, 2023 at 02:17 PM UTC

### Review by @rverdile - Commented on January 05, 2023 at 02:55 PM UTC

### Review by @rverdile - Commented on January 05, 2023 at 03:34 PM UTC

### Review by @rverdile - Commented on January 05, 2023 at 03:55 PM UTC

### Review by @avisiedo - Commented on January 09, 2023 at 12:34 PM UTC

### Review by @rverdile - Commented on January 09, 2023 at 03:36 PM UTC

### Review by @avisiedo - Commented on January 10, 2023 at 01:28 PM UTC

### Review by @jlsherrill - Commented on January 10, 2023 at 06:18 PM UTC

### Review by @jlsherrill - Commented on January 10, 2023 at 06:21 PM UTC

### Review by @jlsherrill - Commented on January 10, 2023 at 06:22 PM UTC

### Review by @jlsherrill - Commented on January 10, 2023 at 06:22 PM UTC

### Review by @jlsherrill - Commented on January 10, 2023 at 06:23 PM UTC

### Review by @jlsherrill - Commented on January 10, 2023 at 06:24 PM UTC

### Review by @rverdile - Commented on January 10, 2023 at 07:52 PM UTC

### Review by @jlsherrill - Commented on January 10, 2023 at 08:05 PM UTC

### Review by @avisiedo - Commented on January 11, 2023 at 09:03 AM UTC

### Review by @avisiedo - Commented on January 11, 2023 at 09:41 AM UTC

### Review by @avisiedo - Commented on January 11, 2023 at 09:41 AM UTC

### Review by @rverdile - Approved on January 11, 2023 at 06:32 PM UTC

looks good!! 

### Review by @jlsherrill - Commented on January 11, 2023 at 06:42 PM UTC

### Review by @jlsherrill - Commented on January 11, 2023 at 06:43 PM UTC

### Review by @avisiedo - Commented on January 12, 2023 at 04:24 PM UTC

### Review by @jlsherrill - Approved on January 13, 2023 at 01:51 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/151*
