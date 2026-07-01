---
type: pull_request
number: 433
title: "Fixes 2738: support SCRAM-SHA-512 sasl mech for kafka"
state: merged
author: jlsherrill
created: 2023-10-17T20:11:22Z
updated: 2023-10-25T13:00:38Z
closed: 2023-10-25T12:55:43Z
merged: 2023-10-25T12:55:43Z
base_branch: main
head_branch: 2738
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/433
---

# Pull Request #433: Fixes 2738: support SCRAM-SHA-512 sasl mech for kafka

**Author**: @jlsherrill
**Created**: October 17, 2023 at 08:11 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `2738`

## Description

## Summary

based on  https://github.com/RedHatInsights/insights-results-aggregator/pull/1864/files

## Testing steps

We may need to just merge to stage to verify this works.

---

## Discussion

### Comment by @jlsherrill on October 17, 2023 at 08:30 PM UTC

https://issues.redhat.com/browse/HMS-2738

### Comment by @jlsherrill on October 18, 2023 at 12:37 PM UTC

/retest

### Comment by @jlsherrill on October 19, 2023 at 03:12 PM UTC

/retest

### Comment by @jlsherrill on October 19, 2023 at 08:44 PM UTC

/retest

### Comment by @jlsherrill on October 25, 2023 at 12:56 PM UTC

i'll monitor our stage deployment to see if the error goes away

---

## Reviews

### Review by @jlsherrill - Commented on October 17, 2023 at 08:11 PM UTC

### Review by @rverdile - Approved on October 24, 2023 at 07:30 PM UTC

I guess I can't test this, but it looks fine

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/433*
