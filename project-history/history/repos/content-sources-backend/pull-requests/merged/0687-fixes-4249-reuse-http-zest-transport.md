---
type: pull_request
number: 687
title: "Fixes 4249: Reuse HTTP zest transport"
state: merged
author: lzap
created: 2024-06-04T09:58:32Z
updated: 2024-06-14T11:22:09Z
closed: 2024-06-11T16:39:03Z
merged: 2024-06-11T16:39:03Z
base_branch: main
head_branch: share-transport
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/687
---

# Pull Request #687: Fixes 4249: Reuse HTTP zest transport

**Author**: @lzap
**Created**: June 04, 2024 at 09:58 AM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `share-transport`

## Description

Hello, I noticed this small detail - a new transport is created for each zest client, this effectively means that no HTTP connection pooling can be done. Here is an example of the [DefaultTransport](https://cs.opensource.google/go/go/+/refs/tags/go1.22.3:src/net/http/transport.go) from the standard library that is supposed to be reused - it is thread safe.

Easy help, this patch just extracts the variable into the package scope so it can be reused on every new client. In case you are wondering, all those timeouts are coming from the `DefaultTransport` from the standard library plus the `ResponseHeaderTimeout` that was already defined here.

_Edit: I was trying out github.dev, did not go well, force pushed the correct version._

---

## Discussion

### Comment by @app-sre-bot on June 04, 2024 at 10:00 AM UTC

Can one of the admins verify this patch?

### Comment by @jlsherrill on June 04, 2024 at 10:00 AM UTC

https://issues.redhat.com/browse/HMS-4249

### Comment by @jlsherrill on June 04, 2024 at 03:43 PM UTC

looks good!  i think you're just missing the 'net' import

### Comment by @lzap on June 07, 2024 at 07:29 AM UTC

Sure, rebased.

### Comment by @jlsherrill on June 11, 2024 at 02:00 PM UTC

[test]

### Comment by @jlsherrill on June 11, 2024 at 04:09 PM UTC

/retest

### Comment by @jlsherrill on June 11, 2024 at 04:39 PM UTC

thanks @lzap !

---

## Reviews

### Review by @jlsherrill - Approved on June 11, 2024 at 04:38 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/687*
