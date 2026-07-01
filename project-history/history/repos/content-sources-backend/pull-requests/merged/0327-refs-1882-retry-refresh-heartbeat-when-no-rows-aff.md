---
type: pull_request
number: 327
title: "Refs 1882: retry refresh heartbeat when no rows affected"
state: merged
author: rverdile
created: 2023-07-03T14:16:58Z
updated: 2023-07-17T14:41:42Z
closed: 2023-07-17T14:41:38Z
merged: 2023-07-17T14:41:38Z
base_branch: main
head_branch: heartbeat-no-rows
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/327
---

# Pull Request #327: Refs 1882: retry refresh heartbeat when no rows affected

**Author**: @rverdile
**Created**: July 03, 2023 at 02:16 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `heartbeat-no-rows`

## Description

## Summary
There is still reoccurring error "No rows affected when refreshing heartbeat". This error seems to be caused by a race condition at the postgres level. The process refreshing the heartbeat and the process completing the task are two separate goroutines. If the process refreshing the heartbeat queries the task while it is finishing, it will get the old values (i.e. the status is running and not completed). But when it refreshes the heartbeat, the task status is "completed" so there is no longer a heartbeat to refresh.

These changes make it so if we've hit the error, we query the task status, and and try again to refresh the heartbeat only if the task is running.

## Testing steps
You can use this python script to help test by introspecting a repository repeatedly.
```
import requests

def sendReq():
  url = "<insert introspect endpoint here>"
  headers = { "x-rh-identity":"<insert identity here>" }
  x = requests.post(url, headers = headers)

for i in range(100):
  sendReq()
```
1. Make the heartbeat interval 100ms, which is unrealistic but will make the error more reproducible. 
2. Without the changes, this script should make the error, `"No rows affected when refreshing heartbeat"`, pop up at least a couple of times.
3. With the changes, you should no longer see the error. 
4. You can also test the tasking system by syncing, interrupting syncing, introspecting via command line, etc. Make sure no new errors were introduced.

---

## Discussion

### Comment by @jlsherrill on July 03, 2023 at 02:30 PM UTC

https://issues.redhat.com/browse/HMS-1882

---

## Reviews

### Review by @jlsherrill - Approved on July 14, 2023 at 05:48 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/327*
