---
type: pull_request
number: 625
title: "Fixes 3902: Add wrapping struct to snapshot by date api"
state: merged
author: jlsherrill
created: 2024-04-05T14:37:07Z
updated: 2024-04-08T12:04:26Z
closed: 2024-04-08T12:04:26Z
merged: 2024-04-08T12:04:26Z
base_branch: main
head_branch: 3902
labels: []
url: https://github.com/content-services/content-sources-backend/pull/625
---

# Pull Request #625: Fixes 3902: Add wrapping struct to snapshot by date api

**Author**: @jlsherrill
**Created**: April 05, 2024 at 02:37 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `3902`

## Description

## Summary

Adds a wrapping struct to the api in case we need to add pagination/search/etc... in the future

## Testing steps
Probably easiest just to test in the UI: https://github.com/content-services/content-sources-frontend/pull/240
## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @ezr-ondrej on April 05, 2024 at 03:18 PM UTC

/retest

### Comment by @xbhouse on April 05, 2024 at 05:04 PM UTC

code and functionality looks good to me here. 

looking at the test error.. i was seeing an unauthorized response locally trying to pull the pulp images too earlier today, i just had to re-login to quay for some reason

### Comment by @xbhouse on April 05, 2024 at 05:14 PM UTC

rerunning failed tests to see if it was just a one-off

### Comment by @xbhouse on April 05, 2024 at 05:31 PM UTC

seems like this could be related to the rotation of the keys in the ephemeral cluster? teams are seeing issues in pr-checks, discussing [here](https://app.slack.com/client/E030G10V24F/C022YV4E0NA)

### Comment by @swadeley on April 05, 2024 at 07:32 PM UTC

/retest

### Comment by @ezr-ondrej on April 08, 2024 at 08:54 AM UTC

/retest

---

## Reviews

### Review by @xbhouse - Approved on April 05, 2024 at 06:14 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/625*
