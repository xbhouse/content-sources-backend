---
type: pull_request
number: 165
title: "Refs 322: Add clowdjobinvocation template for stage testing"
state: merged
author: mshriver
created: 2023-01-06T16:38:15Z
updated: 2023-01-09T15:34:55Z
closed: 2023-01-09T15:34:55Z
merged: 2023-01-09T15:34:54Z
base_branch: main
head_branch: testing-saas-322
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/165
---

# Pull Request #165: Refs 322: Add clowdjobinvocation template for stage testing

**Author**: @mshriver
**Created**: January 06, 2023 at 04:38 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `testing-saas-322`

## Description

This will be deployed by app-interface after stage deployment, and should execute the iqe-content-sources container

---

## Discussion

### Comment by @jlsherrill on January 06, 2023 at 05:00 PM UTC

https://issues.redhat.com/browse/HMSCONTENT-322

### Comment by @jlsherrill on January 06, 2023 at 05:00 PM UTC

⚠️ This task currently requires qe-approval, but this PR does not fully resolve the task.  Please contact QE to determine appropriate testing required.

### Comment by @mshriver on January 09, 2023 at 03:34 PM UTC

File may need some tweaks, but app-interface is now failing because this isn't in the repo. Isn't referenced by any deployments, and thus adding the file should not change clowder deployment at all. 

---

## Reviews

### Review by @jlsherrill - Approved on January 06, 2023 at 04:47 PM UTC

LGTM

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/165*
