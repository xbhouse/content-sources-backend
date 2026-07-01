---
type: pull_request
number: 311
title: "Fixes 1948: Add introspection/edit/delete notifications"
state: merged
author: Andrewgdewar
created: 2023-06-19T19:39:20Z
updated: 2023-06-28T16:22:45Z
closed: 2023-06-28T16:22:44Z
merged: 2023-06-28T16:22:44Z
base_branch: main
head_branch: HMS-1948
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/311
---

# Pull Request #311: Fixes 1948: Add introspection/edit/delete notifications

**Author**: @Andrewgdewar
**Created**: June 19, 2023 at 07:39 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `HMS-1948`

## Description

## Summary



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
4) create a bulk repo, edit it, delete it, and see the associated notifications


---

## Discussion

### Comment by @jlsherrill on June 19, 2023 at 08:00 PM UTC

https://issues.redhat.com/browse/HMS-1948

### Comment by @jlsherrill on June 23, 2023 at 01:50 PM UTC

I think we need a notification for non-bulk create?

### Comment by @jlsherrill on June 28, 2023 at 03:22 PM UTC

needs a rebase, but ack from a code perspective.  I trust @rverdile's testing

---

## Reviews

### Review by @jlsherrill - Commented on June 20, 2023 at 05:18 PM UTC

### Review by @Andrewgdewar - Commented on June 20, 2023 at 05:19 PM UTC

### Review by @rverdile - Commented on June 21, 2023 at 08:15 PM UTC

### Review by @Andrewgdewar - Commented on June 22, 2023 at 09:52 PM UTC

### Review by @rverdile - Commented on June 23, 2023 at 01:31 PM UTC

### Review by @jlsherrill - Commented on June 23, 2023 at 01:50 PM UTC

### Review by @Andrewgdewar - Commented on June 23, 2023 at 02:22 PM UTC

### Review by @rverdile - Commented on June 23, 2023 at 03:03 PM UTC

### Review by @Andrewgdewar - Commented on June 23, 2023 at 04:26 PM UTC

### Review by @Andrewgdewar - Commented on June 23, 2023 at 04:29 PM UTC

### Review by @Andrewgdewar - Commented on June 23, 2023 at 04:57 PM UTC

### Review by @rverdile - Commented on June 26, 2023 at 06:12 PM UTC

### Review by @rverdile - Commented on June 26, 2023 at 06:18 PM UTC

### Review by @rverdile - Commented on June 27, 2023 at 07:43 PM UTC

Not sure if @jlsherrill wants to take another look, but I tested all the endpoints notifications were added to and it's working.

### Review by @jlsherrill - Approved on June 28, 2023 at 03:22 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/311*
