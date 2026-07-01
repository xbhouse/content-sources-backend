---
type: pull_request
number: 305
title: "Refs 602: allow notifications to local kafka"
state: merged
author: jlsherrill
created: 2023-06-05T17:39:19Z
updated: 2023-06-05T18:11:11Z
closed: 2023-06-05T18:11:08Z
merged: 2023-06-05T18:11:08Z
base_branch: main
head_branch: 602
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/305
---

# Pull Request #305: Refs 602: allow notifications to local kafka

**Author**: @jlsherrill
**Created**: June 05, 2023 at 05:39 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `602`

## Description

## Summary
This allows local testing of notification message sending

## Testing steps

1) configure your config for kafka if not already:
```
kafka:
  auto:
    offset:
      reset: latest
    commit:
      interval:
        ms: 5000
  bootstrap:
    servers: localhost:9092
  group:
    id: content-sources
  message:
    send:
      max:
        retries: 15
  request:
    timeout:
      ms: 30000
    required:
      acks: -1
  retry:
    backoff:
      ms: 100
  timeout: 10000
  topics:
    - platform.content-sources.introspect
    - platform.notifications.ingress
```
2) run server `make run`
3) run listener ` KAFKA_TOPIC=platform.notifications.ingress make  kafka-topic-consume`
4) create a bulk repo and see the notification


---

## Discussion

### Comment by @jlsherrill on June 05, 2023 at 06:00 PM UTC

https://issues.redhat.com/browse/HMS-602

---

## Reviews

### Review by @Andrewgdewar - Approved on June 05, 2023 at 06:03 PM UTC

Works! 🎉 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/305*
