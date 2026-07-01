---
type: pull_request
number: 225
title: "Refs 1203: Initial grafana update for new metrics"
state: merged
author: jlsherrill
created: 2023-03-06T19:59:53Z
updated: 2023-03-08T15:35:21Z
closed: 2023-03-07T14:42:26Z
merged: 2023-03-07T14:42:26Z
base_branch: main
head_branch: 1203_1
labels: []
url: https://github.com/content-services/content-sources-backend/pull/225
---

# Pull Request #225: Refs 1203: Initial grafana update for new metrics

**Author**: @jlsherrill
**Created**: March 06, 2023 at 07:59 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `1203_1`

## Description

This includes all metrics except for content_sources_kafka_message_result_total which does not seem to be working correctly


---

## Discussion

### Comment by @jlsherrill on March 06, 2023 at 08:00 PM UTC

https://issues.redhat.com/browse/HMS-1203

### Comment by @jlsherrill on March 06, 2023 at 08:00 PM UTC

![image](https://user-images.githubusercontent.com/395077/223217662-686ee438-9373-4db4-a025-c369505e2133.png)
![image](https://user-images.githubusercontent.com/395077/223217739-2601a730-9c11-4994-bd51-dc434e8667ca.png)


### Comment by @jlsherrill on March 06, 2023 at 08:12 PM UTC

i wanted to go ahead and get 'something' committed, as i have to redo the work if i get logged out.   
So more work tuning will be done, but i'd vote for just going ahead and merging this and tweaking later

### Comment by @swadeley on March 08, 2023 at 03:35 PM UTC

/retest

---

## Reviews

### Review by @rverdile - Approved on March 07, 2023 at 02:33 PM UTC

ack. looks good

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/225*
