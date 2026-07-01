---
type: pull_request
number: 1157
title: "Build: generate deployment.yaml from template"
state: closed
author: jlsherrill
created: 2025-07-25T20:06:35Z
updated: 2026-03-11T12:10:05Z
closed: 2026-03-11T12:10:05Z
base_branch: main
head_branch: auto_gen
labels: []
url: https://github.com/content-services/content-sources-backend/pull/1157
---

# Pull Request #1157: Build: generate deployment.yaml from template

**Author**: @jlsherrill
**Created**: July 25, 2025 at 08:06 PM UTC
**Status**: Closed
**Labels**: None
**Base**: `main` ← **Head**: `auto_gen`

## Description

## Summary

Creates some a method for generating the deployment.yaml file from a template and a yaml file with variables.  Generally all variables are used for all items, except for the transform-pulp-logs which needs very specific special secrets.

## Testing steps

Initially i'm looking for general feedback for the approach, not necessarily a full review at this time.

Largely created by cursor.


---

## Discussion

### Comment by @xbhouse on July 31, 2025 at 08:39 PM UTC

+1 to this idea!

i did a small test to generate the deployment with a new common env var and noticed that it wasn't added to the parameters section, that's something we'll probably want to add. maybe automate it if possible or a validation step to check that every env var has a matching parameter definition?

### Comment by @rverdile on December 05, 2025 at 03:52 PM UTC

more evidence for the usefulness of this: https://github.com/content-services/content-sources-backend/pull/1326

### Comment by @jlsherrill on March 11, 2026 at 12:09 PM UTC

closed in favor of https://github.com/content-services/content-sources-backend/pull/1406

---

## Reviews

### Review by @rverdile - Commented on July 28, 2025 at 08:41 PM UTC

I haven't tested it, but I like the idea. I say let's do it.

### Review by @jlsherrill - Commented on July 30, 2025 at 12:51 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1157*
