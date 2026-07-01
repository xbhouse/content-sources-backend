---
type: pull_request
number: 133
title: "Fixes 235: add /ping for kafka consumer"
state: merged
author: jlsherrill
created: 2022-10-31T20:34:07Z
updated: 2022-11-14T14:26:22Z
closed: 2022-11-14T14:26:18Z
merged: 2022-11-14T14:26:18Z
base_branch: main
head_branch: 235
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/133
---

# Pull Request #133: Fixes 235: add /ping for kafka consumer

**Author**: @jlsherrill
**Created**: October 31, 2022 at 08:34 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `235`

## Description

and use a single binary for api and consumer

---

## Discussion

### Comment by @jlsherrill on October 31, 2022 at 09:00 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-235

### Comment by @jlsherrill on November 03, 2022 at 03:01 PM UTC

To test:

1.  You can now use 'make run' for both api & kafka
2. alternatively, you can run:   `go run cmd/content_sources/main.go api consumer`
3. i'd test those out individually as well:     `go run cmd/content_sources/main.go  api`  &    `go run cmd/content_sources/main.go  consumer`   (although you'll probably get an error trying to run them on the same machine, as they will both try to listen for /ping on the same port)
4.  verify that `curl localhost:8000/ping`  works  for all 3 versions of the command

### Comment by @jlsherrill on November 14, 2022 at 02:13 PM UTC

@avisiedo   Updated with suggestions! thank you!

### Comment by @avisiedo on November 14, 2022 at 02:22 PM UTC

lgtm

---

## Reviews

### Review by @avisiedo - Changes Requested on November 10, 2022 at 07:49 PM UTC

:+1: for the changes on this pr; only a few changes requested.

thinks I have learned:
- I have learned and dig more about mechanisms to synchronize in golang and how to use the context to broadcast the termination.
- I have learned how to debug this stuffs from vscode in a productive way; the below is not in the pr, but I found very helpful when debugging the changes; I suggest to add the below to `.vscode/launch.json` for the service:

```raw
    "console": "integratedTerminal"
```

This open a terminal to launch the debug process, and it allows to invoke the signal handler when ctl+c is pressed into that terminal, sending a sigint to the process. This is important because from the integrated debugger panel in vscode only a sigkill is possible to be sent to the process (when we click stop, press the key-binding to stop debugger, or click the menu item to stop the debugger).

Furthermore; number of replicas for the kafka consumer verified in ephemeral environment:

```raw
$ oc get deployments
NAME                                     READY   UP-TO-DATE   AVAILABLE   AGE
content-sources-backend-kafka-consumer   3/3     3            3           4m30s
content-sources-backend-service          3/3     3            3           4m30s
```

And ping to kafka consumer verified locally and from the ephemeral environment.

### Review by @avisiedo - Approved on November 14, 2022 at 02:21 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/133*
