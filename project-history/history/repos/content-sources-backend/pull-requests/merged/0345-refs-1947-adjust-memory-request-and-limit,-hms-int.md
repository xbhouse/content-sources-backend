---
type: pull_request
number: 345
title: "Refs 1947: Adjust memory request and limit, hms-integration"
state: merged
author: mshriver
created: 2023-08-02T17:50:23Z
updated: 2023-08-03T06:15:41Z
closed: 2023-08-03T06:15:40Z
merged: 2023-08-03T06:15:40Z
base_branch: main
head_branch: update-pod-request
labels: ["bug", "no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/345
---

# Pull Request #345: Refs 1947: Adjust memory request and limit, hms-integration

**Author**: @mshriver
**Created**: August 02, 2023 at 05:50 PM UTC
**Status**: Merged
**Labels**: `bug`, `no-qe-needed`
**Base**: `main` ← **Head**: `update-pod-request`

## Description

The selenium container needs a higher memory limit, but can request in smaller chunks.

The CPU chunks for the iqe container can be smaller, and its memory limit reduced




---

## Discussion

### Comment by @jlsherrill on August 02, 2023 at 06:00 PM UTC

https://issues.redhat.com/browse/HMS-1947

---

## Reviews

### Review by @jlsherrill - Approved on August 02, 2023 at 06:08 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/345*
