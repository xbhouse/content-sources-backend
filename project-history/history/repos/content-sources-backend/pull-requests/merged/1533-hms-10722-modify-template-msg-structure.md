---
type: pull_request
number: 1533
title: "HMS-10722: modify template msg structure"
state: merged
author: xbhouse
created: 2026-06-10T13:37:49Z
updated: 2026-06-10T17:31:44Z
closed: 2026-06-10T17:31:44Z
merged: 2026-06-10T17:31:44Z
base_branch: main
head_branch: 10722
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1533
---

# Pull Request #1533: HMS-10722: modify template msg structure

**Author**: @xbhouse
**Created**: June 10, 2026 at 01:37 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `10722`

## Description

## Summary

Updates the structure of template events to include added and removed advisories. These fields are kept empty for now.

## Testing steps

1. Create the template topic and start the Kafka consumer:

```
export KAFKA_TOPICS=platform.content-sources.template
make kafka-topics-create
make kafka-topic-consume KAFKA_TOPIC=platform.content-sources.template
```
2. Create, update, and delete a template. You should see the template data in the consumer output after each request. `added_advisories` and `removed_advisories` should be omitted

---

## Discussion

### Comment by @xbhouse on June 10, 2026 at 02:00 PM UTC

https://issues.redhat.com/browse/HMS-10722

### Comment by @xbhouse on June 10, 2026 at 04:06 PM UTC

/retest

---

## Reviews

### Review by @TenSt - Approved on June 10, 2026 at 02:58 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1533*
