---
type: pull_request
number: 167
title: "Refs 299: set default log stream"
state: merged
author: jlsherrill
created: 2023-01-06T17:43:47Z
updated: 2023-01-06T17:51:12Z
closed: 2023-01-06T17:51:08Z
merged: 2023-01-06T17:51:08Z
base_branch: main
head_branch: 299_2
labels: []
url: https://github.com/content-services/content-sources-backend/pull/167
---

# Pull Request #167: Refs 299: set default log stream

**Author**: @jlsherrill
**Created**: January 06, 2023 at 05:43 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `299_2`

## Description

to reproduce the error we were seeing in stage, run with: `CLOUDWATCH_STREAM=""`  on main branch.

With this change, if you run with `CLOUDWATCH_STREAM=""`  it will default to a 'default' stream.  Otherwise it will use the specified stream.  This could allow us to customize the stream with env variables in our clowdapp config

---

## Reviews

### Review by @rverdile - Approved on January 06, 2023 at 05:44 PM UTC

ack!

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/167*
