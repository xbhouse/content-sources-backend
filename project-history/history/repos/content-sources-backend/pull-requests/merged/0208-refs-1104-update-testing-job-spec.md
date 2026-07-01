---
type: pull_request
number: 208
title: "Refs 1104: Update testing job spec"
state: merged
author: mshriver
created: 2023-02-14T13:14:59Z
updated: 2023-02-14T13:40:10Z
closed: 2023-02-14T13:40:09Z
merged: 2023-02-14T13:40:09Z
base_branch: main
head_branch: update-testing-job-spec
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/208
---

# Pull Request #208: Refs 1104: Update testing job spec

**Author**: @mshriver
**Created**: February 14, 2023 at 01:14 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `update-testing-job-spec`

## Description

ImagePullSecret is part of the PodSpec not ContainerSpec, and RestartPolicy is required by qontract-reconcile. 

There will be an additional change for the shm volume for the selenium container spec.

---

## Discussion

### Comment by @jlsherrill on February 14, 2023 at 01:30 PM UTC

https://issues.redhat.com/browse/HMS-1104

### Comment by @jlsherrill on February 14, 2023 at 01:30 PM UTC

⚠️ This task currently requires qe-approval, but this PR does not fully resolve the task.  Please contact QE to determine appropriate testing required.

### Comment by @mshriver on February 14, 2023 at 01:40 PM UTC

With @jlsherrill input I'm going to ignore the `build-test / Test` failure in the checks.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/208*
