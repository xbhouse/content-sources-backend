---
type: pull_request
number: 389
title: "Refs 1963: Bump mem request to 200Mi"
state: merged
author: swadeley
created: 2023-09-13T19:36:21Z
updated: 2025-09-08T07:30:25Z
closed: 2023-09-14T07:01:12Z
merged: 2023-09-14T07:01:12Z
base_branch: main
head_branch: swadeley/update_ram_request_size
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/389
---

# Pull Request #389: Refs 1963: Bump mem request to 200Mi

**Author**: @swadeley
**Created**: September 13, 2023 at 07:36 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `swadeley/update_ram_request_size`

## Description

## Summary

The kafka consumer init container is crashing while trying to introspect repos (On startup, all repos in the database are introspected.). This problem only occurs in ephemeral env.

As per @jrusz 's suggestion, this issue could be due to memory taking to long to be increased, so bumping the request size to see if that helps to get the memory increased before the pod crashes. 

## Testing steps

Check for content-sources-backend-kafka-consumer restarting multiple times.


---

## Discussion

### Comment by @jlsherrill on September 13, 2023 at 08:00 PM UTC

https://issues.redhat.com/browse/HMS-1963

### Comment by @jlsherrill on September 13, 2023 at 08:00 PM UTC

⚠️ This task currently requires qe-approval, but this PR does not fully resolve the task.  Please contact QE to determine appropriate testing required.

---

## Reviews

### Review by @jlsherrill - Approved on September 13, 2023 at 07:47 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/389*
