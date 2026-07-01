---
type: pull_request
number: 1163
title: "HMS-8897: Test to fetch and verify repository configuration file"
state: merged
author: mayurilahane
created: 2025-07-31T05:42:23Z
updated: 2025-08-06T12:36:49Z
closed: 2025-08-06T12:36:48Z
merged: 2025-08-06T12:36:48Z
base_branch: main
head_branch: mlahane/HMS-8897
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/1163
---

# Pull Request #1163: HMS-8897: Test to fetch and verify repository configuration file

**Author**: @mayurilahane
**Created**: July 31, 2025 at 05:42 AM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `mlahane/HMS-8897`

## Description

## Summary
migrate - https://gitlab.cee.redhat.com/insights-qe/iqe-content-sources-plugin/-/blob/master/iqe_content_sources/tests/test_repository_api_only.py?ref_type=heads#L1387




---

## Discussion

### Comment by @jlsherrill on July 31, 2025 at 06:00 AM UTC

https://issues.redhat.com/browse/HMS-8897

### Comment by @jlsherrill on August 01, 2025 at 04:36 PM UTC

overall this looks good, i would change the title though to mention 'fetch config.repo', as this isn't really fetching the repodmd.xml

Was a problem in the original test though :) 

### Comment by @swadeley on August 04, 2025 at 07:28 AM UTC

/retest

---

## Reviews

### Review by @swadeley - Approved on August 06, 2025 at 07:28 AM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1163*
