---
type: pull_request
number: 15
title: "Feat: simplify logging format and add a valid container init"
state: merged
author: dominikvagner
created: 2025-08-21T12:08:30Z
updated: 2025-08-21T12:48:21Z
closed: 2025-08-21T12:48:21Z
merged: 2025-08-21T12:48:21Z
base_branch: main
head_branch: integration-improvements
labels: []
url: https://github.com/RedHatInsights/frontend-development-proxy/pull/15
---

# Pull Request #15: Feat: simplify logging format and add a valid container init

**Author**: @dominikvagner
**Created**: August 21, 2025 at 12:08 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `integration-improvements`

## Description

## Summary
- This changes the log format from the default debug one which is too
verbose to Apache common log format provided by the transform-encoder
plugin. This will come in handy when this container is used in a new fec
command with `proxyVerbose` on.

- This adds a new tiny but valid init for the proxy container to handle
signal forwarding.

---

*Archived from: https://github.com/RedHatInsights/frontend-development-proxy/pull/15*
