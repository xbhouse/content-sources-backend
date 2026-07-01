---
type: pull_request
number: 854
title: "Fixes 4860: use timestamptz for date comparisons"
state: merged
author: jlsherrill
created: 2024-10-17T19:45:29Z
updated: 2024-10-18T14:00:22Z
closed: 2024-10-18T13:51:08Z
merged: 2024-10-18T13:51:08Z
base_branch: main
head_branch: date_fix
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/854
---

# Pull Request #854: Fixes 4860: use timestamptz for date comparisons

**Author**: @jlsherrill
**Created**: October 17, 2024 at 07:45 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `date_fix`

## Description

## Summary

Previously the change to this query improperly used the timestamp postgresql type which ignores timezones, by using timestamptz, the query now works properly.  We'll also just convert the date to UTC to be more consistent with using UTC dates.

## Testing steps

Fly to somewhere that doesn't use UTC.

Run the TestUseLatest integration test in test/integration/update_template_content_test.go


---

## Discussion

### Comment by @jlsherrill on October 17, 2024 at 08:00 PM UTC

https://issues.redhat.com/browse/HMS-4860

---

## Reviews

### Review by @rverdile - Approved on October 17, 2024 at 08:10 PM UTC

lgtm!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/854*
