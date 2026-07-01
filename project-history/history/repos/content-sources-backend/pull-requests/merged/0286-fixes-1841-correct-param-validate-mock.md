---
type: pull_request
number: 286
title: "Fixes 1841: correct param validate mock"
state: merged
author: jlsherrill
created: 2023-05-25T18:02:44Z
updated: 2023-05-26T19:28:50Z
closed: 2023-05-26T19:28:49Z
merged: 2023-05-26T19:28:49Z
base_branch: main
head_branch: 1841
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/286
---

# Pull Request #286: Fixes 1841: correct param validate mock

**Author**: @jlsherrill
**Created**: May 25, 2023 at 06:02 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `1841`

## Description

## Summary
The param validator mock was improperly defined, and thus improperly used.  When replaced with a proper mock, it caused a hang, this due to the wait group within the handler not properly marking Done() if an unexpected error occured.  By using defer, we can mark it Done() no matter what happens within the go routine.

## Testing steps
1.  checkout pr
2. make test
3. no hang!

---

## Discussion

### Comment by @jlsherrill on May 25, 2023 at 06:30 PM UTC

https://issues.redhat.com/browse/HMS-1841

---

## Reviews

### Review by @Andrewgdewar - Approved on May 26, 2023 at 06:51 PM UTC

ACK!

### Review by @mshriver - Approved on May 26, 2023 at 07:28 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/286*
