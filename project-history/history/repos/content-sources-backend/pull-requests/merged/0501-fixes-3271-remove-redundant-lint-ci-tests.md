---
type: pull_request
number: 501
title: "Fixes 3271: remove redundant lint CI tests"
state: merged
author: rverdile
created: 2023-12-11T22:01:27Z
updated: 2023-12-12T16:02:28Z
closed: 2023-12-12T16:02:25Z
merged: 2023-12-12T16:02:25Z
base_branch: main
head_branch: rm-lint
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/501
---

# Pull Request #501: Fixes 3271: remove redundant lint CI tests

**Author**: @rverdile
**Created**: December 11, 2023 at 10:01 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `rm-lint`

## Description

## Summary
- govet and gofmt are already tested by golangci-lint
- there's also no need to explicitly enable tests in the args for the golangci-lint action because we have a config file for it

## Testing steps
tests pass
## Checklist

- [x] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @rverdile on December 11, 2023 at 10:07 PM UTC

Review similar yummy PR as well: https://github.com/content-services/yummy/pull/19

### Comment by @jlsherrill on December 11, 2023 at 10:30 PM UTC

https://issues.redhat.com/browse/HMS-3271

---

## Reviews

### Review by @jlsherrill - Approved on December 12, 2023 at 02:46 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/501*
