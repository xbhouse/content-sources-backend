---
type: pull_request
number: 1272
title: "HMS-8642: save successful pulp log parse, and retry if failed"
state: merged
author: jlsherrill
created: 2025-10-31T20:48:36Z
updated: 2025-11-30T23:20:21Z
closed: 2025-11-30T23:20:18Z
merged: 2025-11-30T23:20:18Z
base_branch: main
head_branch: pulp_logs_retry
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1272
---

# Pull Request #1272: HMS-8642: save successful pulp log parse, and retry if failed

**Author**: @jlsherrill
**Created**: October 31, 2025 at 08:48 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `pulp_logs_retry`

## Description

## Summary

This adds the ability for the transform_pulp_logs to both report failure, but also retry from the previously successful time.
It saves the previous successful time via a new table 'memos'.  This can be reused for other things in the future

## Testing steps

Its not easy to test, we may have to merge this and verify in stage.

You can confirm that the metric is being reported via:  

```
$ curl localhost:9000/metrics | grep content_sources_pulp_transform_logs_days_since_success
```

You can add a memo:

```
INSERT INTO MEMOS (uuid, key, memo)
VALUES (
    'a1b2c3d4-e5f6-7890-abcd-ef1234567890'::uuid,
    'last_successful_pulp_log_date',
    '{"date": "2024-01-15"}'::jsonb
);
```

and then wait ~30-60 seconds and re-check:



```
$ curl localhost:9000/metrics | grep content_sources_pulp_transform_logs_days_since_success
content_sources_pulp_transform_logs_days_since_success 665

```


---

## Discussion

### Comment by @xbhouse on October 31, 2025 at 09:00 PM UTC

https://issues.redhat.com/browse/HMS-8642

### Comment by @jlsherrill on November 10, 2025 at 09:28 PM UTC

/retest

### Comment by @jlsherrill on November 21, 2025 at 07:12 PM UTC

updated!

---

## Reviews

### Review by @jlsherrill - Commented on October 31, 2025 at 08:48 PM UTC

### Review by @xbhouse - Commented on November 12, 2025 at 03:35 PM UTC

### Review by @xbhouse - Commented on November 17, 2025 at 05:03 PM UTC

### Review by @xbhouse - Commented on November 17, 2025 at 05:16 PM UTC

### Review by @xbhouse - Approved on November 24, 2025 at 03:41 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1272*
