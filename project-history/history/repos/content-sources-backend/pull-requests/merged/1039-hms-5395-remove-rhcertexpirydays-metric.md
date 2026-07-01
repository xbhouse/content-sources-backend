---
type: pull_request
number: 1039
title: "HMS-5395: remove RHCertExpiryDays metric"
state: merged
author: rverdile
created: 2025-03-20T18:47:24Z
updated: 2025-03-31T14:38:24Z
closed: 2025-03-31T14:38:21Z
merged: 2025-03-31T14:38:21Z
base_branch: main
head_branch: rm-rh-expiry
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/1039
---

# Pull Request #1039: HMS-5395: remove RHCertExpiryDays metric

**Author**: @rverdile
**Created**: March 20, 2025 at 06:47 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `rm-rh-expiry`

## Description

## Summary
Removes the RHCertExpiryDays metric because it is replaced by the CertificateExpiryDays metric

Alert has already been removed from stage

## Testing steps
1. `make run`
2. View metrics at `http://localhost:9000/metrics
3. This metric no longer shows up 


---

## Discussion

### Comment by @jlsherrill on March 20, 2025 at 07:00 PM UTC

https://issues.redhat.com/browse/HMS-5395

---

## Reviews

### Review by @copilot-pull-request-reviewer - Commented on March 28, 2025 at 08:33 PM UTC

## Pull Request Overview

This pull request removes the deprecated RHCertExpiryDays metric in favor of the CertificateExpiryDays metric.  
- Removed RHCertExpiryDays from the metrics constants and struct in metrics.go.  
- Eliminated the gauge creation for RHCertExpiryDays in metrics.go.  
- Removed the setting of RHCertExpiryDays and an associated log line from collector.go.

### Reviewed Changes

Copilot reviewed 2 out of 2 changed files in this pull request and generated no comments.

| File                                      | Description                                               |
| ----------------------------------------- | --------------------------------------------------------- |
| pkg/instrumentation/metrics.go            | Removal of the RHCertExpiryDays metric definition and gauge |
| pkg/instrumentation/custom/collector.go   | Removal of RHCertExpiryDays usage and redundant log line  |





### Review by @xbhouse - Approved on March 31, 2025 at 02:20 PM UTC

ack!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1039*
