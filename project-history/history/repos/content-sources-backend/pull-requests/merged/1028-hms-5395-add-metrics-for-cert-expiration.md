---
type: pull_request
number: 1028
title: "HMS-5395: add metrics for cert expiration"
state: merged
author: rverdile
created: 2025-03-14T21:20:00Z
updated: 2025-03-18T18:31:02Z
closed: 2025-03-18T18:30:59Z
merged: 2025-03-18T18:30:59Z
base_branch: main
head_branch: cert-alert
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/1028
---

# Pull Request #1028: HMS-5395: add metrics for cert expiration

**Author**: @rverdile
**Created**: March 14, 2025 at 09:20 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `cert-alert`

## Description

## Summary
- Adds gauge vector with "certificate_label" type to filter by 
- This replaces the existing red hat cdn metric, so eventually that can be removed. But not until the existing alert is removed.

## Testing steps
1. Configure all of the certs: cdn, candlepin, and feature service
2. `make run`
3. verify that `http:/localhost:9000/metrics` shows the CertificationExpiryDays metric, with each cert's days separated by the label



---

## Discussion

### Comment by @jlsherrill on March 14, 2025 at 09:30 PM UTC

https://issues.redhat.com/browse/HMS-5395

---

## Reviews

### Review by @dominikvagner - Approved on March 18, 2025 at 02:11 PM UTC

works as expected, will be nice to have this! 😌 
nice improvement to the certificates file by splitting it into smaller functions ✨ 

approved! good job 👏🏼
 _just one small question about the `TODO` comments in this commit 💭_ 


### Review by @rverdile - Commented on March 18, 2025 at 03:06 PM UTC

### Review by @dominikvagner - Commented on March 18, 2025 at 03:43 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1028*
