---
type: pull_request
number: 749
title: "Fixes 4423: handle 404 on delete template"
state: merged
author: xbhouse
created: 2024-07-17T15:14:29Z
updated: 2024-07-19T08:00:20Z
closed: 2024-07-19T07:34:26Z
merged: 2024-07-19T07:34:26Z
base_branch: main
head_branch: 4423
labels: ["qe-testing-needed"]
url: https://github.com/content-services/content-sources-backend/pull/749
---

# Pull Request #749: Fixes 4423: handle 404 on delete template

**Author**: @xbhouse
**Created**: July 17, 2024 at 03:14 PM UTC
**Status**: Merged
**Labels**: `qe-testing-needed`
**Base**: `main` ← **Head**: `4423`

## Description

## Summary

Ignores 404s when trying to delete an environment in candlepin

## Testing steps

- Start the app with candlepin disabled (set the candlepin.server to blank in config.yaml) and create a template
- Restart the app with candlepin enabled and delete the template. You shouldn't see any errors in the logs
- Without this PR you would see `couldn't delete environment: 404 Not Found`
- This [PR](https://github.com/content-services/content-sources-backend/pull/745) resolved the issue we were seeing on fetching environments when removing a repo from a template

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @jlsherrill on July 17, 2024 at 03:30 PM UTC

https://issues.redhat.com/browse/HMS-4423

---

## Reviews

### Review by @rverdile - Approved on July 18, 2024 at 09:01 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/749*
