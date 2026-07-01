---
type: pull_request
number: 1483
title: "HMS-10605: replace zerolog-sentry with sentry-go/zerolog"
state: merged
author: rverdile
created: 2026-04-28T17:40:01Z
updated: 2026-04-29T15:00:56Z
closed: 2026-04-29T15:00:55Z
merged: 2026-04-29T15:00:55Z
base_branch: main
head_branch: replace-zerolog-sentry
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1483
---

# Pull Request #1483: HMS-10605: replace zerolog-sentry with sentry-go/zerolog

**Author**: @rverdile
**Created**: April 28, 2026 at 05:40 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `replace-zerolog-sentry`

## Description

## Summary
The archdx/zerolog-sentry package is no longer maintained and causes linter errors with newer sentry-go versions

## Testing steps
You could test this by pointing your local env to our stage glitchtip dsn and seeing if an error message gets sent, but I already did that: https://glitchtip.devshift.net/insights/issues/4410906/events/87a2e211614844b0941a3aed203c083b


---

## Discussion

### Comment by @xbhouse on April 28, 2026 at 06:00 PM UTC

https://issues.redhat.com/browse/HMS-10605

---

## Reviews

### Review by @xbhouse - Approved on April 28, 2026 at 10:11 PM UTC

thank you! 

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1483*
