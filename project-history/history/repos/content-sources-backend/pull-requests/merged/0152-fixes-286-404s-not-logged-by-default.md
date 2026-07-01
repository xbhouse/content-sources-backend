---
type: pull_request
number: 152
title: "Fixes 286: 404s not logged by default"
state: merged
author: avisiedo
created: 2022-11-30T10:08:29Z
updated: 2023-01-16T11:31:27Z
closed: 2022-12-08T15:01:16Z
merged: 2022-12-08T15:01:16Z
base_branch: main
head_branch: hmscontent-286
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/152
---

# Pull Request #152: Fixes 286: 404s not logged by default

**Author**: @avisiedo
**Created**: November 30, 2022 at 10:08 AM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `hmscontent-286`

## Description

- Add default log level value to info.
- Add LOGGING_LEVEL parameter to the deployment level (we could increase the level in deployment time for ephemeral environment if needed).
- Fix '/ping' and '/ping/' endpoint and update the tests.

TODO:

- [x] Pending to check last time in ephemeral environment (I have to move some commits to the branch I cherry pick that help me on this, which works with bonfire 4.11.0).

---

## Discussion

### Comment by @jlsherrill on November 30, 2022 at 10:09 AM UTC

https://issues.redhat.com/browse/HMSCONTENT-286

### Comment by @jlsherrill on November 30, 2022 at 01:39 PM UTC

I confirmed in ephemeral:
```
{"level":"error","error":"code=404, message=Not Found","remote_ip":"10.128.8.9","host":"env-ephemeral-yejy47-gcn8a3ez.apps.c-rh-c-eph.8p0c.p1.openshiftapps.com","method":"GET","uri":"/api/content-sources/v1/zombies/","user_agent":"curl/7.82.0","status":404,"referer":"","latency":0.013939,"latency_human":"13.939µs","bytes_in":"0","bytes_out":"49","time":"2022-11-30T13:38:36Z","time":"2022-11-30T13:38:36Z","caller":"/go/pkg/mod/github.com/ziflex/lecho/v3@v3.1.0/middleware.go:124"}
```

With bonfire v 4.11.1

### Comment by @avisiedo on December 02, 2022 at 07:50 AM UTC

There is another situation that I will address in a different ticket; the new situation is that I cannot set the log level to filter messages; I mean, if I set the log level to "error" and the message to print out is "info" level, I would expect the message is not printed out, however, I am getting that message printed out.

Debugging I have seen that at some point, the log level for the logger is "debug", so everything is going to be printed out. I have tried several changes, but nothing gave me the expected result. By the way, the current root cause for this pr is addressed, I will open a new card to address that other situation, as this pr is letting me to deploy into the ephemeral environment.


### Comment by @jlsherrill on December 02, 2022 at 04:23 PM UTC

linting error, otherwise looks good

### Comment by @jlsherrill on December 05, 2022 at 05:16 PM UTC

feel free to merge!

---

## Reviews

### Review by @jlsherrill - Commented on November 30, 2022 at 01:24 PM UTC

### Review by @jlsherrill - Commented on November 30, 2022 at 01:40 PM UTC

### Review by @avisiedo - Commented on December 01, 2022 at 06:00 PM UTC

### Review by @avisiedo - Commented on December 01, 2022 at 06:12 PM UTC

### Review by @jlsherrill - Approved on December 05, 2022 at 05:16 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/152*
