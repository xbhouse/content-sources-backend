---
type: pull_request
number: 961
title: "HMS-5396: gather pulp logs and upload to s3"
state: merged
author: jlsherrill
created: 2025-02-03T18:50:56Z
updated: 2025-02-06T13:03:16Z
closed: 2025-02-06T13:03:10Z
merged: 2025-02-06T13:03:10Z
base_branch: main
head_branch: logs
labels: ["qe-testing-needed", "Ready for QE"]
url: https://github.com/content-services/content-sources-backend/pull/961
---

# Pull Request #961: HMS-5396: gather pulp logs and upload to s3

**Author**: @jlsherrill
**Created**: February 03, 2025 at 06:50 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`, `Ready for QE`
**Base**: `main` ← **Head**: `logs`

## Description

## Summary

Sets up a cronjob to run daily a job to pull pulp logs from cloudwatch, parse them, translate the domain to the orgid, assemble them into a gzippd csv and upload to s3

## Testing steps

Not much testing can be done, I'll have to monitor in stage and adjust.  



---

## Discussion

### Comment by @jlsherrill on February 03, 2025 at 07:00 PM UTC

https://issues.redhat.com/browse/HMS-5396

### Comment by @akhil-jha on February 05, 2025 at 07:59 AM UTC

/retest

---

## Reviews

### Review by @rverdile - Commented on February 05, 2025 at 07:22 PM UTC

### Review by @rverdile - Commented on February 05, 2025 at 07:22 PM UTC

one question, code looks good

### Review by @jlsherrill - Commented on February 05, 2025 at 07:59 PM UTC

### Review by @rverdile - Approved on February 05, 2025 at 08:14 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/961*
