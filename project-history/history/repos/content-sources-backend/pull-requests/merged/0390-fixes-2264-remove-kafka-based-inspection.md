---
type: pull_request
number: 390
title: "Fixes 2264: remove kafka based inspection"
state: merged
author: jlsherrill
created: 2023-09-14T18:01:22Z
updated: 2023-09-21T16:15:58Z
closed: 2023-09-21T16:15:55Z
merged: 2023-09-21T16:15:55Z
base_branch: main
head_branch: 2264
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/390
---

# Pull Request #390: Fixes 2264: remove kafka based inspection

**Author**: @jlsherrill
**Created**: September 14, 2023 at 06:01 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `2264`

## Description

## Summary
This removes the kafka based introspection and workers.  Some kafka code is still present to service notifications.

## Testing steps

To see that notification aren't broken run:
```
make kafka-topic-consume KAFKA_TOPIC=platform.notifications.ingress
```

then create some repository.  You should see a couple messages come through


---

## Discussion

### Comment by @jlsherrill on September 14, 2023 at 06:30 PM UTC

https://issues.redhat.com/browse/HMS-2264

---

## Reviews

### Review by @Andrewgdewar - Commented on September 19, 2023 at 06:29 PM UTC

✅  

- Everything seems to be working fine. 
- Notifications appear to be sent without issue.
- All tests are passing


### Review by @rverdile - Approved on September 20, 2023 at 09:08 PM UTC

tested and ack from me too

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/390*
