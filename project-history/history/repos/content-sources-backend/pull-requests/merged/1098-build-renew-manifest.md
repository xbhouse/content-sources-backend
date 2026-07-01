---
type: pull_request
number: 1098
title: "Build: renew manifest"
state: merged
author: xbhouse
created: 2025-04-29T17:41:14Z
updated: 2025-04-30T13:25:57Z
closed: 2025-04-30T13:25:56Z
merged: 2025-04-30T13:25:56Z
base_branch: main
head_branch: update-manifest
labels: ["no-qe-needed"]
url: https://github.com/content-services/content-sources-backend/pull/1098
---

# Pull Request #1098: Build: renew manifest

**Author**: @xbhouse
**Created**: April 29, 2025 at 05:41 PM UTC
**Status**: Merged
**Labels**: `no-qe-needed`
**Base**: `main` ← **Head**: `update-manifest`

## Description

## Summary

Our manifest expired on April 24th, this just replaces it with a new one

## Testing steps

1. Run make compose-clean and make compose-up on this PR, the new manifest should be imported
2. The subscription check endpoint (`/subscription_check/`) should return `"red_hat_enterprise_linux": true` and you should be able to create/edit templates with the UI

---

## Discussion

### Comment by @jlsherrill on April 29, 2025 at 07:39 PM UTC

/retest

---

## Reviews

### Review by @jlsherrill - Approved on April 29, 2025 at 05:54 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/1098*
