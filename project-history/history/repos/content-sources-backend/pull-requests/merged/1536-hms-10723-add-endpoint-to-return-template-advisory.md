---
type: pull_request
number: 1536
title: "HMS-10723: add endpoint to return template advisory IDs"
state: merged
author: xbhouse
created: 2026-06-10T18:30:36Z
updated: 2026-06-18T19:17:08Z
closed: 2026-06-18T19:17:08Z
merged: 2026-06-18T19:17:08Z
base_branch: main
head_branch: 10723
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1536
---

# Pull Request #1536: HMS-10723: add endpoint to return template advisory IDs

**Author**: @xbhouse
**Created**: June 10, 2026 at 06:30 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `10723`

## Description

## Summary

- Adds an endpoint that returns a template's advisory IDs with no pagination
- Sends template update events whenever a template is updated

## Testing steps

1. Create a template with snapshots that have advisories.
2. Send a GET request to  `/templates/:uuid/advisories/ids`. This should return all the advisory IDs in the template.
3. Run `make kafka-topic-consume KAFKA_TOPIC=platform.content-sources.template`
4. Update the template. There should be 2 messages in the consumer output, one when the update request is sent and one during the `update-template-content` task.
5. Add a custom repo to a template that's using the latest snapshots. Triggering a new snapshot for that repo should enqueue the `update-latest-snapshot` task, and a message should be sent in the consumer output during that task.
6. Change the template to a date before the latest custom repo snapshot. Delete the custom repo snapshot used by the template, this should enqueue a `delete-snapshots` task and a message should be sent in the consumer output.
7. Delete the custom repository used by the template. This should enqueue a `delete-repository-snapshots` task and a message should be sent during this task too.



---

## Discussion

### Comment by @xbhouse on June 10, 2026 at 07:00 PM UTC

https://issues.redhat.com/browse/HMS-10723

### Comment by @xbhouse on June 16, 2026 at 01:50 PM UTC

> delete-snapshots and delete-repository-snapshots also modify templates, so those need to send an update event too

thank you! was about to ask which other 2 tasks needed to send an event :) will add that in

---

## Reviews

### Review by @TenSt - Commented on June 12, 2026 at 08:37 AM UTC

### Review by @TenSt - Commented on June 12, 2026 at 09:33 AM UTC

### Review by @TenSt - Dismissed on June 12, 2026 at 09:35 AM UTC

lgtm! 
There are a couple comments, but feel free to merge as is or I'll re-approve if you decide to do the changes.

### Review by @rverdile - Commented on June 12, 2026 at 11:10 AM UTC

### Review by @rverdile - Commented on June 12, 2026 at 01:06 PM UTC

### Review by @xbhouse - Commented on June 12, 2026 at 01:52 PM UTC

### Review by @xbhouse - Commented on June 12, 2026 at 02:15 PM UTC

### Review by @xbhouse - Commented on June 12, 2026 at 02:44 PM UTC

### Review by @rverdile - Commented on June 12, 2026 at 04:39 PM UTC

### Review by @rverdile - Commented on June 12, 2026 at 04:43 PM UTC

### Review by @TenSt - Commented on June 12, 2026 at 09:27 PM UTC

### Review by @TenSt - Commented on June 12, 2026 at 09:32 PM UTC

### Review by @xbhouse - Commented on June 12, 2026 at 11:02 PM UTC

### Review by @xbhouse - Commented on June 15, 2026 at 07:46 PM UTC

### Review by @TenSt - Dismissed on June 16, 2026 at 12:32 PM UTC

lgtm! great work 👍 

### Review by @rverdile - Commented on June 16, 2026 at 01:48 PM UTC

delete-snapshots and delete-repository-snapshots also modify templates, so those need to send an update event too

### Review by @TenSt - Dismissed on June 17, 2026 at 07:50 AM UTC

### Review by @TenSt - Approved on June 18, 2026 at 02:28 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1536*
