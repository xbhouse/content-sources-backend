---
type: pull_request
number: 1246
title: "HMS-4443: add User-Agent header to Candlepin client"
state: merged
author: katarinazaprazna
created: 2025-10-22T08:13:55Z
updated: 2025-10-22T12:38:38Z
closed: 2025-10-22T12:38:38Z
merged: 2025-10-22T12:38:38Z
base_branch: main
head_branch: add-user-agent-header-to-candlepin-client
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/1246
---

# Pull Request #1246: HMS-4443: add User-Agent header to Candlepin client

**Author**: @katarinazaprazna
**Created**: October 22, 2025 at 08:13 AM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `add-user-agent-header-to-candlepin-client`

## Description

## Summary
Added a user-friendly User-Agent value to requests sent to the Candlepin service, allowing easier identification of our service in Hosted Candlepin logs.

## Testing steps
Inspect the Candlepin logs using `podman exec -it cs_candlepin_1 cat /var/log/candlepin/candlepin.log`

<img width="2083" height="904" alt="Screenshot From 2025-10-22 09-51-09" src="https://github.com/user-attachments/assets/f66cad47-10f8-4cba-aa88-3c77b734562c" />


---

## Discussion

### Comment by @xbhouse on October 22, 2025 at 08:30 AM UTC

https://issues.redhat.com/browse/HMS-4443

### Comment by @swadeley on October 22, 2025 at 11:24 AM UTC

@katarinazaprazna "Code style issues found"

---

## Reviews

### Review by @swadeley - Approved on October 22, 2025 at 11:21 AM UTC

ACK

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1246*
