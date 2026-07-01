---
type: pull_request
number: 438
title: "Fixes 2896: refactor pulp client in dao"
state: merged
author: rverdile
created: 2023-10-23T14:40:54Z
updated: 2023-11-18T08:02:54Z
closed: 2023-11-16T18:42:04Z
merged: 2023-11-16T18:42:04Z
base_branch: main
head_branch: pc-3
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/438
---

# Pull Request #438: Fixes 2896: refactor pulp client in dao

**Author**: @rverdile
**Created**: October 23, 2023 at 02:40 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `pc-3`

## Description

## Summary
The method I introduced [here](https://github.com/content-services/content-sources-backend/pull/414/files#diff-c90c06b1d5bb21209cf1d0b2bd26ae86ec3db770764e6653d5dc72580a10c16fR78-R89) to set the domain scoped pulp client is bug-prone.

This PR demonstrates a new way of setting the pulp client by adding two chained methods. Now, we can initialize the pulp client with `context.Background()` during dependency injection. Then, replace the context later, during the request, so that the pulp client is using the domain.

For example, the pulp client method `GetContentPath()` is called with `pulpClient.WithContext(ctx).WithDomain(domainName).GetContentPath()`

I've commented out most of the failing tests for now, as I was just trying out this approach.

---

## Discussion

### Comment by @jlsherrill on October 24, 2023 at 08:00 PM UTC

https://issues.redhat.com/browse/HMS-2896

### Comment by @jlsherrill on October 27, 2023 at 01:28 PM UTC

this needs a rebase now too

### Comment by @Andrewgdewar on November 07, 2023 at 06:40 PM UTC

/retest

### Comment by @rverdile on November 13, 2023 at 08:35 PM UTC

this is updated now

### Comment by @jlsherrill on November 15, 2023 at 03:50 PM UTC

/retest

### Comment by @jlsherrill on November 16, 2023 at 03:28 PM UTC

/retest

---

## Reviews

### Review by @rverdile - Commented on October 24, 2023 at 07:41 PM UTC

### Review by @jlsherrill - Commented on October 26, 2023 at 12:30 PM UTC

### Review by @jlsherrill - Commented on October 26, 2023 at 07:39 PM UTC

### Review by @rverdile - Commented on October 27, 2023 at 07:22 PM UTC

### Review by @jlsherrill - Commented on October 31, 2023 at 06:59 PM UTC

### Review by @rverdile - Commented on November 02, 2023 at 05:29 PM UTC

### Review by @jlsherrill - Approved on November 16, 2023 at 03:37 PM UTC

ACK pending tests passing

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/438*
