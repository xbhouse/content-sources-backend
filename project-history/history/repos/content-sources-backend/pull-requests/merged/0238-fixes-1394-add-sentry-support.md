---
type: pull_request
number: 238
title: "Fixes 1394: Add sentry support"
state: merged
author: jlsherrill
created: 2023-03-27T15:04:15Z
updated: 2023-04-04T11:56:26Z
closed: 2023-04-04T00:49:25Z
merged: 2023-04-04T00:49:25Z
base_branch: main
head_branch: 1394
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/238
---

# Pull Request #238: Fixes 1394: Add sentry support

**Author**: @jlsherrill
**Created**: March 27, 2023 at 03:04 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `1394`

## Description

## Summary
This adds sentry/glitchtip support for alerts on errors, panics, and fatal messages

## Testing steps
I tried setting up glitchtip locally, and fought my way through email setup issues, and ended up testing with glitchtip directly.  So you can login to glitchtip grab the DSN for hcc-content-sources-stage and run the server locally with  SENTRY_DSN="FROM_GLITCHTIP"  make run

Then trigger an error.  Everyone should get a notification, so only do this sparingly :)


---

## Discussion

### Comment by @jlsherrill on March 27, 2023 at 03:30 PM UTC

https://issues.redhat.com/browse/HMS-1394

---

## Reviews

### Review by @rverdile - Approved on April 03, 2023 at 03:25 PM UTC

tested a few different errors, can see them in the glitchtip interface. lgtm!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/238*
