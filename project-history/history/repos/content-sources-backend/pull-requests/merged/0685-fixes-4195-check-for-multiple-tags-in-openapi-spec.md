---
type: pull_request
number: 685
title: "Fixes 4195: check for multiple tags in openapi spec"
state: merged
author: xbhouse
created: 2024-06-03T16:40:23Z
updated: 2024-06-04T18:34:54Z
closed: 2024-06-04T18:34:54Z
merged: 2024-06-04T18:34:54Z
base_branch: main
head_branch: 4195
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/685
---

# Pull Request #685: Fixes 4195: check for multiple tags in openapi spec

**Author**: @xbhouse
**Created**: June 03, 2024 at 04:40 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `4195`

## Description

## Summary

Adds a linter to check for multiple tags in the openapi spec when running `make openapi`. Each tag is associated with the resulting data. For example, for listing repository package groups, the tag would be `packagegroups`.

## Testing steps

- If there are any cases of multiple tags in the openapi spec, `make openapi` should fail and tell you which endpoint has more than one tag
- `make openapi` should not fail if each endpoint has only one tag
- For QE, some endpoints will be moved around (for example, if an endpoint was originally found under the `repositories_api` and it lists repository rpms, it would now only be found under the `rpms_api`) 

## Checklist

- [ ] Tested with snapshotting feature disabled and pulp server URL not configured if appropriate


---

## Discussion

### Comment by @xbhouse on June 03, 2024 at 04:57 PM UTC

i can't see what's causing the snyk failure here, how can i see/fix that?

### Comment by @jlsherrill on June 03, 2024 at 05:00 PM UTC

https://issues.redhat.com/browse/HMS-4195

---

## Reviews

### Review by @jlsherrill - Approved on June 04, 2024 at 05:48 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/685*
