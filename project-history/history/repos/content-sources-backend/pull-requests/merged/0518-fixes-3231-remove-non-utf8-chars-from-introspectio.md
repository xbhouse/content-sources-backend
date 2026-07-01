---
type: pull_request
number: 518
title: "Fixes 3231: remove non-UTF8 chars from introspection error"
state: merged
author: rverdile
created: 2024-01-03T20:55:41Z
updated: 2024-01-05T19:01:06Z
closed: 2024-01-05T19:01:02Z
merged: 2024-01-05T19:01:02Z
base_branch: main
head_branch: utf8
labels: []
url: https://github.com/content-services/content-sources-backend/pull/518
---

# Pull Request #518: Fixes 3231: remove non-UTF8 chars from introspection error

**Author**: @rverdile
**Created**: January 03, 2024 at 08:55 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `utf8`

## Description

## Summary
When introspecting, certain error messages may contain non-UTF8 characters. The database doesn't like non-UTF8 characters. This removes those characters before the error messages are inserted into the database.

## Testing steps

1. Assuming [this yummy PR](https://github.com/content-services/yummy/pull/21) is not integrated into the backend, introspect this repository: https://fixtures.pulpproject.org/rpm-pkglists-updateinfo/
2. Without this PR you will get the error: "invalid byte sequence for encoding \"UTF8\": 0xf4 0x62 0x20 0x28 (SQLSTATE 22021)"
3. With this PR, you should get no error

## Checklist

- [x] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @rverdile on January 03, 2024 at 09:48 PM UTC

/retest

### Comment by @xbhouse on January 04, 2024 at 08:50 PM UTC

tested functionality and no longer seeing that sql error, error message was inserted in the database without that character. lgtm! 

### Comment by @swadeley on January 05, 2024 at 09:44 AM UTC

@jlsherrill This looks like a "no qe required" (but no Jira link) and there are no API changes. So merge when ready I say.

---

## Reviews

### Review by @xbhouse - Approved on January 04, 2024 at 09:23 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/518*
