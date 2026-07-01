---
type: pull_request
number: 1469
title: "HMS-10325: allow minor and major versions in version filter"
state: merged
author: xbhouse
created: 2026-04-17T19:08:59Z
updated: 2026-04-21T14:42:01Z
closed: 2026-04-21T14:42:01Z
merged: 2026-04-21T14:42:01Z
base_branch: main
head_branch: version-filtering
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1469
---

# Pull Request #1469: HMS-10325: allow minor and major versions in version filter

**Author**: @xbhouse
**Created**: April 17, 2026 at 07:08 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `version-filtering`

## Description

## Summary

Updates version filtering logic for repos and templates to allow both minor and major versions

## Testing steps

1. Create a few non-EUS and EUS repos and templates 
2. Filtering repos by both major and minor version in the version filter should return the expected repos (e.g. `/repositories/?version=9,8.6` should return only RHEL 9 (non-EUS) and RHEL 8.6)
3. Filtering templates by both major and minor version in the version filter should return the expected templates (e.g. `/templates/?version=9,8.6` should return only RHEL 9 (non-EUS) and RHEL 8.6)


---

## Discussion

### Comment by @xbhouse on April 17, 2026 at 07:30 PM UTC

https://issues.redhat.com/browse/HMS-10325

---

## Reviews

### Review by @katarinazaprazna - Approved on April 20, 2026 at 10:45 AM UTC

LGTM. Thanks for implementing this! :) It provides a great foundation for the upcoming frontend work

### Review by @rverdile - Approved on April 21, 2026 at 02:31 PM UTC

lgtm!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1469*
