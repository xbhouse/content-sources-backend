---
type: pull_request
number: 566
title: "Fixes 3420: add service notify for templates"
state: merged
author: rverdile
created: 2024-02-06T19:36:26Z
updated: 2024-02-09T09:31:53Z
closed: 2024-02-09T09:31:53Z
merged: 2024-02-09T09:31:53Z
base_branch: main
head_branch: template-notify
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/566
---

# Pull Request #566: Fixes 3420: add service notify for templates

**Author**: @rverdile
**Created**: February 06, 2024 at 07:36 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `template-notify`

## Description

## Summary
Sends an event on a new kafka topic when a template is created, updated, or deleted. 

This has nothing to do with sending an event (notification) to the notifications service or getting emails. I've used the term "notification" in places because it is still technically a notification (we're notifying the patch service), I wasn't sure what else to call it, and there's an opportunity here to share code between the templates and notifications. I'm open to other names!

## Testing steps
1. In a terminal, `make kafka-topic-consume KAFKA_TOPIC=platform.content-sources.template`
2. Start the API server
3. Create a template. Check the terminal where you've just run the kafka consumer, you should see a notification.
4. Update a template. Check for a notification.
5. Delete a template. Check for a notification.
## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on February 06, 2024 at 08:00 PM UTC

https://issues.redhat.com/browse/HMS-3420

### Comment by @swadeley on February 08, 2024 at 12:09 PM UTC

/retest

### Comment by @jlsherrill on February 08, 2024 at 07:41 PM UTC

You'll need to add the new topic name here:  https://github.com/content-services/content-sources-backend/blob/main/deployments/deployment.yaml#L21-L27

Otherwise this looks good and behave as expected !

### Comment by @jlsherrill on February 08, 2024 at 08:51 PM UTC

/retest

### Comment by @swadeley on February 09, 2024 at 09:31 AM UTC

> ACK, but we may need to reach out to Qe about the test failure. I'll rekick it for good measure

Hi, the one fail in the test run is related to the read-only user in the rbac test. There have been some changes to user creation code. So, not related to this PR.

---

## Reviews

### Review by @rverdile - Commented on February 06, 2024 at 07:40 PM UTC

### Review by @jlsherrill - Commented on February 07, 2024 at 06:57 PM UTC

### Review by @jlsherrill - Commented on February 07, 2024 at 07:11 PM UTC

### Review by @jlsherrill - Commented on February 07, 2024 at 07:12 PM UTC

### Review by @jlsherrill - Commented on February 07, 2024 at 07:13 PM UTC

### Review by @rverdile - Commented on February 07, 2024 at 09:12 PM UTC

### Review by @rverdile - Commented on February 07, 2024 at 09:12 PM UTC

### Review by @rverdile - Commented on February 07, 2024 at 09:13 PM UTC

### Review by @rverdile - Commented on February 07, 2024 at 09:14 PM UTC

### Review by @jlsherrill - Commented on February 08, 2024 at 08:51 PM UTC

ACK, but we may need to reach out to Qe about the test failure.  I'll rekick it for good measure

### Review by @jlsherrill - Approved on February 08, 2024 at 08:51 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/566*
