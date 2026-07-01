---
type: pull_request
number: 1058
title: "Test: re-add caching to buildspec"
state: closed
author: xbhouse
created: 2025-04-04T14:25:01Z
updated: 2025-04-28T16:57:34Z
closed: 2025-04-28T16:57:34Z
base_branch: main
head_branch: update-buildspec
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1058
---

# Pull Request #1058: Test: re-add caching to buildspec

**Author**: @xbhouse
**Created**: April 04, 2025 at 02:25 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `main` ← **Head**: `update-buildspec`

## Description

## Summary

The build-playwright-test job has been taking a bit longer than 30 minutes (about 40 minutes) so this adds caching back to the buildspec to see if this speeds it up a bit

## Testing steps

build-playwright-test finishes, codebuild doesn't time out



---

## Discussion

### Comment by @marusak on April 08, 2025 at 09:18 AM UTC

What was the time difference with and without this change?

### Comment by @xbhouse on April 08, 2025 at 12:53 PM UTC

> What was the time difference with and without this change?

@marusak this doesn't seem to have helped :/ both with and without this change the job took about 39 minutes 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1058*
