---
type: pull_request
number: 269
title: "Fixes 602: Send create notification"
state: merged
author: Andrewgdewar
created: 2023-05-11T14:07:19Z
updated: 2023-05-17T15:33:12Z
closed: 2023-05-17T15:33:11Z
merged: 2023-05-17T15:33:11Z
base_branch: main
head_branch: HMS-602-1
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/269
---

# Pull Request #269: Fixes 602: Send create notification

**Author**: @Andrewgdewar
**Created**: May 11, 2023 at 02:07 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `HMS-602-1`

## Description

## Summary

Initial test of the notifications service integration. 

This current PR only implements the "BulkCreate" API.

A follow up PR will implement the rest of the notifications required after this pattern is proven.

## Testing steps

To test locally:  

- Navigate to "/pkg/notifications/client.go"

- comment out lines 19 - 25, and the closing brace on line 67

- Replace the "kafkaServers" on line 32 with the following:
<img width="1025" alt="Screen Shot 2023-05-11 at 10 24 36 AM" src="https://github.com/content-services/content-sources-backend/assets/38083295/3f12f7d3-a0cf-49a5-94ff-3a57bfd3020a">

- Install/run [KafkaCat](https://docs.confluent.io/platform/current/clients/kafkacat-usage.html#consumer-mode) so that you can record the message when it is sent.
- Call bulkCreate (through the API or via UI)
- Notice that kafkaCat should output a json blob. 
- (Optional) For schema verification, you can copy the json output from KafkaCat and paste it [here](https://internal.console.stage.redhat.com/api/notifications/internal/#/utils/console-cloud-event-validator), to very it is a valid schema.



---

## Discussion

### Comment by @jlsherrill on May 11, 2023 at 02:30 PM UTC

https://issues.redhat.com/browse/HMS-602

### Comment by @swadeley on May 16, 2023 at 04:06 PM UTC

/retest

### Comment by @jlsherrill on May 16, 2023 at 06:03 PM UTC

This now needs a rebase

---

## Reviews

### Review by @jlsherrill - Commented on May 11, 2023 at 03:27 PM UTC

### Review by @jlsherrill - Commented on May 11, 2023 at 03:32 PM UTC

### Review by @Andrewgdewar - Commented on May 11, 2023 at 03:55 PM UTC

### Review by @Andrewgdewar - Commented on May 11, 2023 at 03:57 PM UTC

### Review by @jlsherrill - Commented on May 11, 2023 at 05:57 PM UTC

### Review by @jlsherrill - Approved on May 11, 2023 at 07:20 PM UTC

### Review by @rverdile - Commented on May 11, 2023 at 07:34 PM UTC

### Review by @rverdile - Commented on May 11, 2023 at 07:36 PM UTC

### Review by @Andrewgdewar - Commented on May 11, 2023 at 08:30 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/269*
