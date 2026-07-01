---
type: pull_request
number: 540
title: "Refs 3423: fix arm64 repo arch"
state: merged
author: jlsherrill
created: 2024-01-19T15:41:33Z
updated: 2024-01-19T17:49:49Z
closed: 2024-01-19T17:49:49Z
merged: 2024-01-19T17:49:49Z
base_branch: main
head_branch: 3423_2
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/540
---

# Pull Request #540: Refs 3423: fix arm64 repo arch

**Author**: @jlsherrill
**Created**: January 19, 2024 at 03:41 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `3423_2`

## Description

## Summary

This fixes an incorrect arch i had added.  

## Testing steps

Check out master, run `make repos-import`   notice the "Red Hat Ansible Engine 2 for RHEL 8 ARM 64 (RPMs)" repo has an arch of x86_64. 

Checkout this PR, run `make repos-import` again, notice its not correct

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on January 19, 2024 at 04:00 PM UTC

https://issues.redhat.com/browse/HMS-3423

### Comment by @jlsherrill on January 19, 2024 at 04:00 PM UTC

⚠️ This task currently requires qe-approval, but this PR does not fully resolve the task.  Please contact QE to determine appropriate testing required.

---

## Reviews

### Review by @Andrewgdewar - Approved on January 19, 2024 at 03:45 PM UTC

![no-yes ](https://github.com/content-services/content-sources-backend/assets/38083295/57b36d2e-02ff-4875-a017-6abcb4b2b6e7)


---

*Archived from: https://github.com/content-services/content-sources-backend/pull/540*
