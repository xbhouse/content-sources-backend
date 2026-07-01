---
type: pull_request
number: 204
title: "Refs 1104: Add params to CJI for iqe"
state: merged
author: mshriver
created: 2023-02-10T18:39:06Z
updated: 2023-02-10T19:02:59Z
closed: 2023-02-10T19:02:59Z
merged: 2023-02-10T19:02:59Z
base_branch: main
head_branch: add-iqe-job-params
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/204
---

# Pull Request #204: Refs 1104: Add params to CJI for iqe

**Author**: @mshriver
**Created**: February 10, 2023 at 06:39 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `add-iqe-job-params`

## Description

Include ui param to add selenium container
explicitly specify plugin

I'm not entirely sure this will correctly include the container in the CJI pod for iqe, but this param will be needed regardless of current working order. The CJI spec at least recognizes it right now.

https://github.com/RedHatInsights/clowder/blob/f43aee65bbbe9c1303c23de7661613f2bdf66110/apis/cloud.redhat.com/v1alpha1/clowdjobinvocation_types.go#L59




---

## Discussion

### Comment by @mshriver on February 10, 2023 at 06:43 PM UTC

/retest

### Comment by @jlsherrill on February 10, 2023 at 07:00 PM UTC

https://issues.redhat.com/browse/HMS-1104

### Comment by @jlsherrill on February 10, 2023 at 07:00 PM UTC

⚠️ This task currently requires qe-approval, but this PR does not fully resolve the task.  Please contact QE to determine appropriate testing required.

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/204*
