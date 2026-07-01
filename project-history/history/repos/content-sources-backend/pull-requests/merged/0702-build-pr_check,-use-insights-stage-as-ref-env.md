---
type: pull_request
number: 702
title: "Build: pr_check, use insights-stage as ref-env"
state: merged
author: jrusz
created: 2024-06-12T12:08:33Z
updated: 2024-07-10T15:07:10Z
closed: 2024-06-17T18:13:28Z
merged: 2024-06-17T18:13:28Z
base_branch: main
head_branch: change-ref
labels: []
url: https://github.com/content-services/content-sources-backend/pull/702
---

# Pull Request #702: Build: pr_check, use insights-stage as ref-env

**Author**: @jrusz
**Created**: June 12, 2024 at 12:08 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `change-ref`

## Description

This will deploy the apps using whatever is in stage which is where PRs are deployed after being merged.

## Summary

## Testing steps

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @app-sre-bot on June 12, 2024 at 12:09 PM UTC

Can one of the admins verify this patch?

### Comment by @jlsherrill on June 12, 2024 at 03:09 PM UTC

[test]

### Comment by @jlsherrill on June 12, 2024 at 03:10 PM UTC

[test]

### Comment by @swadeley on June 17, 2024 at 07:49 AM UTC

/retest

---

## Reviews

### Review by @swadeley - Approved on June 17, 2024 at 06:13 PM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/702*
