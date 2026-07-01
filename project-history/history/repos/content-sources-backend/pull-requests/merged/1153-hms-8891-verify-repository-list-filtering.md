---
type: pull_request
number: 1153
title: "HMS-8891: Verify repository list filtering"
state: merged
author: mayurilahane
created: 2025-07-24T04:08:56Z
updated: 2025-07-28T19:52:05Z
closed: 2025-07-28T19:52:05Z
merged: 2025-07-28T19:52:05Z
base_branch: main
head_branch: mlahane/HMS-8891
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/1153
---

# Pull Request #1153: HMS-8891: Verify repository list filtering

**Author**: @mayurilahane
**Created**: July 24, 2025 at 04:08 AM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `mlahane/HMS-8891`

## Description

## Summary
migrating https://gitlab.cee.redhat.com/insights-qe/iqe-content-sources-plugin/-/blob/master/iqe_content_sources/tests/test_repository_api_only.py?ref_type=heads#L1540




---

## Discussion

### Comment by @jlsherrill on July 24, 2025 at 04:30 AM UTC

https://issues.redhat.com/browse/HMS-8891

### Comment by @mayurilahane on July 24, 2025 at 04:52 PM UTC

/retest

### Comment by @xbhouse on July 28, 2025 at 06:44 PM UTC

since this test is checking for `any` filters specifically, it might be good to adjust the test name to make that more clear? maybe `Verify filtering by 'any' in repository list filter`?

### Comment by @xbhouse on July 28, 2025 at 06:45 PM UTC

just one comment, otherwise this looks good to me! nice job! :smile: 

---

## Reviews

### Review by @xbhouse - Approved on July 28, 2025 at 07:46 PM UTC

lgtm!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1153*
