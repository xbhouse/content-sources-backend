---
type: pull_request
number: 1122
title: "Build: Adjust job name for transform-pulp-logs-fix"
state: merged
author: jlsherrill
created: 2025-06-04T12:40:02Z
updated: 2025-06-04T12:57:54Z
closed: 2025-06-04T12:57:27Z
merged: 2025-06-04T12:57:27Z
base_branch: main
head_branch: jobfix
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1122
---

# Pull Request #1122: Build: Adjust job name for transform-pulp-logs-fix

**Author**: @jlsherrill
**Created**: June 04, 2025 at 12:40 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `jobfix`

## Description

## Summary

The transform-pulp-logs-fix-2025-05 job didn't run in stage and I'm not sure why, a couple theories:

* the job name in deployment.yaml and jobs-stage.yaml were the same, maybe this isn't supported?
* The job is very similar to the other job (and cron jobs get created with a similar name), so make the start different

## Testing steps

none


---

## Discussion

### Comment by @jlsherrill on June 04, 2025 at 12:57 PM UTC

oh yes, in this case its fixing it for may.... which isn't really the intention of those dates :)   will keep that in mind and adjust if i have to do a 2nd round

---

## Reviews

### Review by @dominikvagner - Approved on June 04, 2025 at 12:52 PM UTC

looks good, hopefully that fixes that problem 🙏🏼 ack ✅ 

_note: if you're changing the dates on the jobs, you could also bump the month, not that it matters, but it's not May anymore 😝_ 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1122*
