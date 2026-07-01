---
type: pull_request
number: 96
title: "Fixes 162: Let search packages by uuid"
state: merged
author: avisiedo
created: 2022-09-01T08:02:31Z
updated: 2022-09-09T06:00:27Z
closed: 2022-09-09T05:52:46Z
merged: 2022-09-09T05:52:46Z
base_branch: main
head_branch: hmscontent-162
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/96
---

# Pull Request #96: Fixes 162: Let search packages by uuid

**Author**: @avisiedo
**Created**: September 01, 2022 at 08:02 AM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `hmscontent-162`

## Description

This change will let to search rpm packages by specifying a list of repository_configurations uuids; it is combined as a logic or with the list of urls.

For /api/content_sources/v1/rpms/names route it will let to:

```json
{
    "urls": ["https://MzjLlbBGbpZTNGIIVnoN.com/IuQyV"],
    "uuids": ["98f88705-ba61-496f-9ed9-12268df2f790"],
    "search": "a"
}
```

or

```json
{
    "urls": ["https://MzjLlbBGbpZTNGIIVnoN.com/IuQyV"],
    "search": "a"
}
```

or

```json
{
    "uuids": ["98f88705-ba61-496f-9ed9-12268df2f790"],
    "search": "a"
}
```

a list of urls or uuids should be present, or both, but not none of them.

List of changes:

- Update request structure.
- Update dao object to consider the list of uuids.
- Update tests.
- Update api documentation.

---

## Discussion

### Comment by @jlsherrill on September 01, 2022 at 08:30 AM UTC

https://issues.redhat.com/browse/HMSCONTENT-162

---

## Reviews

### Review by @swadeley - Commented on September 06, 2022 at 12:15 PM UTC

Typos

### Review by @swadeley - Commented on September 06, 2022 at 12:27 PM UTC

### Review by @rverdile - Commented on September 06, 2022 at 02:47 PM UTC

### Review by @rverdile - Commented on September 06, 2022 at 02:48 PM UTC

### Review by @rverdile - Commented on September 06, 2022 at 05:18 PM UTC

### Review by @rverdile - Commented on September 06, 2022 at 05:18 PM UTC

### Review by @rverdile - Commented on September 07, 2022 at 06:32 PM UTC

### Review by @avisiedo - Commented on September 08, 2022 at 12:17 PM UTC

### Review by @avisiedo - Commented on September 08, 2022 at 07:07 PM UTC

### Review by @rverdile - Approved on September 08, 2022 at 09:25 PM UTC

looks good! nice! 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/96*
