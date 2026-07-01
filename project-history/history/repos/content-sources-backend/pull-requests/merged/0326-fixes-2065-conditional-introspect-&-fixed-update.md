---
type: pull_request
number: 326
title: "Fixes 2065: conditional Introspect & fixed update "
state: merged
author: Andrewgdewar
created: 2023-06-30T19:46:42Z
updated: 2023-07-04T15:22:59Z
closed: 2023-07-04T15:22:59Z
merged: 2023-07-04T15:22:59Z
base_branch: main
head_branch: HMS-2065
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/326
---

# Pull Request #326: Fixes 2065: conditional Introspect & fixed update 

**Author**: @Andrewgdewar
**Created**: June 30, 2023 at 07:46 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `HMS-2065`

## Description

## Summary

This fixes the empty struct values for the update notification that were causing it to fail to be parsed by the schems. 

This adds conditional sending of the introspect notification so that it now will only send a notification when the introspection would result in an update.


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
4) create a bulk repo, edit it, delete it, and see the associated notifications. 
5) check the json for the "package_count": and "url": values (these were appearing as empty string "", when editing/updating a repo).






---

## Discussion

### Comment by @jlsherrill on June 30, 2023 at 08:00 PM UTC

https://issues.redhat.com/browse/HMS-2065

---

## Reviews

### Review by @rverdile - Commented on July 03, 2023 at 05:29 PM UTC

I noticed an issue with BulkCreate notifications while reviewing the BulkDelete PR.

Right now in BulkCreate, if there is an error and no repositories are created, a notification will be sent out because mappedValues here could be empty: https://github.com/content-services/content-sources-backend/blob/main/pkg/dao/repository_configs.go#L112-L116

Do you want to fix that here? You'd just check if it's empty or not before sending a notification.

### Review by @swadeley - Approved on July 04, 2023 at 03:22 PM UTC

lgtm

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/326*
