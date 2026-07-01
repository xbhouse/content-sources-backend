---
type: pull_request
number: 850
title: "Fixes 4852: close pgpool after integration tests"
state: merged
author: jlsherrill
created: 2024-10-16T14:51:10Z
updated: 2024-10-21T08:07:22Z
closed: 2024-10-21T08:07:22Z
merged: 2024-10-21T08:07:22Z
base_branch: main
head_branch: pg_cancel
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/850
---

# Pull Request #850: Fixes 4852: close pgpool after integration tests

**Author**: @jlsherrill
**Created**: October 16, 2024 at 02:51 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `pg_cancel`

## Description

## Summary
run the integration tests
```
make test-integration
```

while its running check db connections:
```
make db-cli-connect


select * from
(select count(*) used from pg_stat_activity) q1,
(select setting::int res_for_super from pg_settings where name=$$superuser_reserved_connections$$) q2,
(select setting::int max_conn from pg_settings where name=$$max_connections$$) q3;
```

Prior to this is would grow as the tests ran to approach the 100 default limit we have through our docker compose

## Testing steps



---

## Discussion

### Comment by @jlsherrill on October 16, 2024 at 03:00 PM UTC

https://issues.redhat.com/browse/HMS-TODO

### Comment by @jlsherrill on October 16, 2024 at 04:00 PM UTC

https://issues.redhat.com/browse/HMS-4852

---

## Reviews

### Review by @jlsherrill - Commented on October 16, 2024 at 03:00 PM UTC

### Review by @rverdile - Commented on October 18, 2024 at 01:32 PM UTC

### Review by @rverdile - Commented on October 18, 2024 at 01:34 PM UTC

### Review by @jlsherrill - Commented on October 18, 2024 at 02:03 PM UTC

### Review by @jlsherrill - Commented on October 18, 2024 at 02:56 PM UTC

### Review by @jlsherrill - Commented on October 18, 2024 at 06:26 PM UTC

### Review by @rverdile - Commented on October 18, 2024 at 07:54 PM UTC

### Review by @jlsherrill - Commented on October 18, 2024 at 07:58 PM UTC

### Review by @rverdile - Approved on October 18, 2024 at 08:10 PM UTC

looks good!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/850*
