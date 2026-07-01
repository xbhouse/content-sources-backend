---
type: pull_request
number: 705
title: "Refs 3600: wrap template delete calls"
state: merged
author: jlsherrill
created: 2024-06-12T16:46:54Z
updated: 2024-06-12T18:31:12Z
closed: 2024-06-12T18:31:07Z
merged: 2024-06-12T18:31:07Z
base_branch: main
head_branch: 3600_2
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/705
---

# Pull Request #705: Refs 3600: wrap template delete calls

**Author**: @jlsherrill
**Created**: June 12, 2024 at 04:46 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `3600_2`

## Description

## Summary

Don't call out to candlepin when deleting a template if its not configured

## Testing steps

Locally disable the candlepin server by commenting out the candlepin server url
create a template and then try to delete it

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on June 12, 2024 at 05:00 PM UTC

https://issues.redhat.com/browse/HMS-3600

---

## Reviews

### Review by @rverdile - Approved on June 12, 2024 at 05:50 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/705*
