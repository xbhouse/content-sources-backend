---
type: pull_request
number: 1288
title: "HMS-9097: chore(db-logger): add err exception similar to api logger"
state: merged
author: TenSt
created: 2025-11-06T11:27:44Z
updated: 2025-11-11T14:28:50Z
closed: 2025-11-11T14:28:50Z
merged: 2025-11-11T14:28:50Z
base_branch: main
head_branch: stepan/HMS-9097-add-exceptions-to-dblogger-similar-to-api-logger
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1288
---

# Pull Request #1288: HMS-9097: chore(db-logger): add err exception similar to api logger

**Author**: @TenSt
**Created**: November 06, 2025 at 11:27 AM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `stepan/HMS-9097-add-exceptions-to-dblogger-similar-to-api-logger`

## Description

## Summary
This PR adds 2 more exceptions for the `dbLogger` to ignore errors about duplicate keys and invalid syntax. These types of errors are already ignored on the API level, so we're just syncing the behaviors.


---

## Discussion

### Comment by @TenSt on November 07, 2025 at 09:50 AM UTC

> this has a JIRA ticket, right? we should add that to the title so it associates to the PR

Yes, sorry, forgot to do that. Updated the PR's title

### Comment by @xbhouse on November 07, 2025 at 10:00 AM UTC

https://issues.redhat.com/browse/HMS-9097

### Comment by @TenSt on November 07, 2025 at 10:12 AM UTC

/retest

---

## Reviews

### Review by @rverdile - Commented on November 06, 2025 at 04:05 PM UTC

this has a JIRA ticket, right? we should add that to the title so it associates to the PR

### Review by @jlsherrill - Approved on November 11, 2025 at 02:27 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1288*
