---
type: pull_request
number: 323
title: "Refs 1947: Update IQE CJI names and parameter defaults"
state: merged
author: tpapaioa
created: 2023-06-29T15:51:16Z
updated: 2023-07-18T16:51:22Z
closed: 2023-07-03T16:38:33Z
merged: 2023-07-03T16:38:33Z
base_branch: main
head_branch: update_cji_template
labels: []
url: https://github.com/content-services/content-sources-backend/pull/323
---

# Pull Request #323: Refs 1947: Update IQE CJI names and parameter defaults

**Author**: @tpapaioa
**Created**: June 29, 2023 at 03:51 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `update_cji_template`

## Description

## Summary
Now that testing.yaml is used for both stage and prod post-deploy tests, I've made a couple updates:
- Remove references to "stage" in the template and CJI names
- Remove the default IQE test marker and filter values, and marked the parameters as required, so that the values are passed in from the saas files in app-interface.

## Testing steps


---

## Discussion

### Comment by @jlsherrill on June 29, 2023 at 04:00 PM UTC

https://issues.redhat.com/browse/HMS-1947

---

## Reviews

### Review by @jlsherrill - Approved on June 30, 2023 at 06:06 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/323*
