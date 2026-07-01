---
type: pull_request
number: 1037
title: "HMS-5395: add all cert expirations to grafana"
state: merged
author: rverdile
created: 2025-03-18T19:21:31Z
updated: 2025-03-19T09:51:36Z
closed: 2025-03-19T09:51:36Z
merged: 2025-03-19T09:51:36Z
base_branch: main
head_branch: cert-alert-grafana
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/1037
---

# Pull Request #1037: HMS-5395: add all cert expirations to grafana

**Author**: @rverdile
**Created**: March 18, 2025 at 07:21 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `cert-alert-grafana`

## Description

## Summary
Replaces the gauge of the rh cdn expiry days with a group of guages that show expiry days for all the certs. Any new certs added to this metric will automatically show up within this group

Looks like this:
![image](https://github.com/user-attachments/assets/b46cc2e9-6e6f-4c17-995a-beb0a650bb44)


## Testing steps
You can look at the board after this merges and I can open a new PR with changes if needed


---

## Discussion

### Comment by @jlsherrill on March 18, 2025 at 07:30 PM UTC

https://issues.redhat.com/browse/HMS-5395

---

## Reviews

### Review by @dominikvagner - Approved on March 19, 2025 at 09:50 AM UTC

looks great! ✨  let's see this in Grafana 🚀😄 
approved! ✅  

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1037*
