---
type: pull_request
number: 498
title: "improve cyndi tag metrics query"
state: merged
author: MichaelMraka
created: 2021-01-26T12:45:30Z
updated: 2021-04-16T11:22:38Z
closed: 2021-01-26T14:25:15Z
merged: 2021-01-26T14:25:14Z
base_branch: master
head_branch: pr15
labels: []
url: https://github.com/RedHatInsights/patchman-engine/pull/498
---

# Pull Request #498: improve cyndi tag metrics query

**Author**: @MichaelMraka
**Created**: January 26, 2021 at 12:45 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `master` ← **Head**: `pr15`

## Description

it runs every 15s and does fullscan over inventory.hosts
this change reduces cost from
```
Aggregate  (cost=415838.35..415838.36 rows=1 width=8)
   ->  Unique  (cost=405794.02..408663.83 rows=573962 width=32)
         ->  Sort  (cost=405794.02..407228.92 rows=573962 width=32)
               Sort Key: (jsonb_array_elements(hosts_v1_1607594495056200400.tags))
               ->  Gather  (cost=1000.00..337158.84 rows=573962 width=32)
                     Workers Planned: 2
                     ->  ProjectSet  (cost=0.00..278762.64 rows=23915100 width=32)
                           ->  Parallel Seq Scan on hosts_v1_1607594495056200400  (cost=0.00..157393.51 rows=239151 width=319)

```
to
```
Aggregate  (cost=221569.25..221569.26 rows=1 width=8)
   ->  Gather  (cost=1000.00..219177.74 rows=191321 width=32)
         Workers Planned: 2
         ->  ProjectSet  (cost=0.00..199045.64 rows=7971700 width=32)
               ->  Parallel Seq Scan on hosts_v1_1607594495056200400  (cost=0.00..158589.26 rows=79717 width=319)
                     Filter: (jsonb_array_length(tags) > 0)
```

---

## Reviews

### Review by @josef-hak - Changes Requested on January 26, 2021 at 01:18 PM UTC

I am getting:
~~~
vmaas_sync/metrics_cyndi.go:62: line is 132 characters (lll)
~~~
... from tests

### Review by @josef-hak - Changes Requested on January 26, 2021 at 01:23 PM UTC

I am getting related test fail:
~~~
metrics_cyndi_test.go:85: 
test         |         	Error Trace:	metrics_cyndi_test.go:85
test         |         	Error:      	Not equal: 
test         |         	            	expected: 3
test         |         	            	actual  : 18
test         |         	Test:       	TestCyndiTags
~~~

### Review by @josef-hak - Approved on January 26, 2021 at 02:24 PM UTC

Great improvement. It's good you have aws console access :smile: 

---

*Archived from: https://github.com/RedHatInsights/patchman-engine/pull/498*
