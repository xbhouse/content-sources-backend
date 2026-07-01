---
type: pull_request
number: 306
title: "Fixes 25: Remove selenium container from testing.yaml"
state: merged
author: tpapaioa
created: 2023-06-06T18:50:09Z
updated: 2023-06-07T14:37:56Z
closed: 2023-06-07T14:37:30Z
merged: 2023-06-07T14:37:30Z
base_branch: main
head_branch: remove_selenium_testing_yaml
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/306
---

# Pull Request #306: Fixes 25: Remove selenium container from testing.yaml

**Author**: @tpapaioa
**Created**: June 06, 2023 at 06:50 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `remove_selenium_testing_yaml`

## Description

## Summary
The test deployment in testing.yaml only runs api tests, so the `ui` section is unnecessary and creates an unused selenium container for each deployment. This PR removes the `ui` section so that the container isn't created at all.

## Testing steps
N/A

---

## Discussion

### Comment by @jlsherrill on June 06, 2023 at 07:00 PM UTC

https://issues.redhat.com/browse/HMS-25

---

## Reviews

### Review by @mshriver - Approved on June 06, 2023 at 07:14 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/306*
