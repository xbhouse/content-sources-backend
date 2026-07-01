---
type: pull_request
number: 644
title: "Refs 3897: fix sentry name"
state: merged
author: rverdile
created: 2024-04-23T15:40:05Z
updated: 2024-04-23T16:45:28Z
closed: 2024-04-23T16:45:21Z
merged: 2024-04-23T16:45:21Z
base_branch: main
head_branch: glitchtip-bug
labels: []
url: https://github.com/content-services/content-sources-backend/pull/644
---

# Pull Request #644: Refs 3897: fix sentry name

**Author**: @rverdile
**Created**: April 23, 2024 at 03:40 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `glitchtip-bug`

## Description

## Summary

I'm pretty sure this is the reason we're not getting alerts for service errors. I noticed in the logs that sentry was being configured for the task worker, but not for the service. The name for the task worker is "content-sources-glitchtip". Both pods were working fine in ephemeral if I used an env variable.

I included an error log to verify if this works. If it does, I'll remove the error log in a followup.

## Testing steps
Just merge and see I think

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Reviews

### Review by @jlsherrill - Approved on April 23, 2024 at 03:48 PM UTC

makes perfect sense, ack!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/644*
