---
type: pull_request
number: 1253
title: "HMS-9350: Add cdn certificate expiry action in playwright-actions.yml"
state: merged
author: mayurilahane
created: 2025-10-24T18:39:07Z
updated: 2025-10-24T20:26:12Z
closed: 2025-10-24T20:26:12Z
merged: 2025-10-24T20:26:12Z
base_branch: main
head_branch: mlahane/cdn_cert
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1253
---

# Pull Request #1253: HMS-9350: Add cdn certificate expiry action in playwright-actions.yml

**Author**: @mayurilahane
**Created**: October 24, 2025 at 06:39 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `mlahane/cdn_cert`

## Description

## Summary

1. Using openssl to check the certificate expiry
2. Fail if the certificate is expired
3. Log the number of days until the certificate expires

## Testing steps

To verify the certificate's status, have a look at the pipeline check "playwright test". It will show the remaining time until expiration or confirm if it has already expired. Refer to the screenshot below.



---

## Discussion

### Comment by @mayurilahane on October 24, 2025 at 06:45 PM UTC

<img width="500" height="300" alt="image" src="https://github.com/user-attachments/assets/b062da93-cc67-43e1-a08b-ef764a218d39" />


### Comment by @xbhouse on October 24, 2025 at 07:00 PM UTC

https://issues.redhat.com/browse/HMS-9350

### Comment by @mayurilahane on October 24, 2025 at 07:54 PM UTC

Frontend PR - https://github.com/content-services/content-sources-frontend/pull/699

---

## Reviews

### Review by @xbhouse - Approved on October 24, 2025 at 08:08 PM UTC

lgtm!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1253*
