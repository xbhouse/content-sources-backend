---
type: pull_request
number: 222
title: "Fixes 1358: set rbac timeout properly"
state: closed
author: jlsherrill
created: 2023-02-28T17:42:21Z
updated: 2023-02-28T21:43:28Z
closed: 2023-02-28T17:56:52Z
base_branch: main
head_branch: 1358
labels: []
url: https://github.com/content-services/content-sources-backend/pull/222
---

# Pull Request #222: Fixes 1358: set rbac timeout properly

**Author**: @jlsherrill
**Created**: February 28, 2023 at 05:42 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `main` ← **Head**: `1358`

## Description

## Summary

I'm about 75% sure this is the problem in stage, as the config value wasn't properly being converted to 
seconds in the main usage (but was as part of metrics).

## Testing steps
Tests pass
Works against mock rbac service

---

## Discussion

### Comment by @jlsherrill on February 28, 2023 at 05:43 PM UTC

actually double checking something


### Comment by @jlsherrill on February 28, 2023 at 05:57 PM UTC

this isn't actually needed, but i'm not sure what is

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/222*
