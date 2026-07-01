---
type: pull_request
number: 440
title: "Fixes 2890: don't query pulp if not enabled"
state: merged
author: rverdile
created: 2023-10-24T13:54:39Z
updated: 2023-10-25T17:45:27Z
closed: 2023-10-25T08:39:49Z
merged: 2023-10-25T08:39:49Z
base_branch: main
head_branch: 2890
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/440
---

# Pull Request #440: Fixes 2890: don't query pulp if not enabled

**Author**: @rverdile
**Created**: October 24, 2023 at 01:54 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `2890`

## Description

## Summary
Seeing error where if pulp isn't enabled, certain API endpoints are broken because there are dao functions that query pulp. This checks if snapshots are enabled in all of those places before querying pulp.

## Testing steps
Turn off pulp and see if the errors are still present when listing/fetching repositories.

---

## Discussion

### Comment by @jlsherrill on October 24, 2023 at 02:00 PM UTC

https://issues.redhat.com/browse/HMS-2890

### Comment by @jlsherrill on October 24, 2023 at 02:06 PM UTC

code looks good, i will test

### Comment by @jlsherrill on October 24, 2023 at 02:53 PM UTC

You should be able to test this in ephemeral with:

--set-parameter content-sources-backend/FEATURES_SNAPSHOTS_ENABLED=false   --set-parameter content-sources-backend/CLIENTS_PULP_SERVER=""

### Comment by @swadeley on October 25, 2023 at 08:39 AM UTC

Hi, as per Ryan's suggestion I had to use the bonfire conf to get `CLIENTS_PULP_SERVER=""` setting to work.

All API tests pass.

Thank you

---

## Reviews

### Review by @jlsherrill - Approved on October 24, 2023 at 02:42 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/440*
