---
type: pull_request
number: 981
title: "Test: add cdn cert"
state: merged
author: xbhouse
created: 2025-02-14T18:06:26Z
updated: 2025-02-18T19:10:58Z
closed: 2025-02-18T19:10:58Z
merged: 2025-02-18T19:10:58Z
base_branch: main
head_branch: add-cdn-cert
labels: []
url: https://github.com/content-services/content-sources-backend/pull/981
---

# Pull Request #981: Test: add cdn cert

**Author**: @xbhouse
**Created**: February 14, 2025 at 06:06 PM UTC
**Status**: Merged
**Labels**: None
**Base**: `main` ← **Head**: `add-cdn-cert`

## Description

## Summary

- Adds the CDN cert for RH repos to the backend config

## Testing steps

- The same action steps exist on the frontend repo via this [PR](https://github.com/content-services/content-sources-frontend/pull/440) and were validated with the frontend RH repo tests
- You shouldn't see this warning anymore in the action step to trigger a snapshot for RH repos: `{"level":"warn","time":"2025-02-14T17:18:43Z","message":"No Red Hat CDN cert pair configured."}`

---

## Discussion

### Comment by @Andrewgdewar on February 18, 2025 at 07:10 PM UTC

Test failure unrelated and resolved in front-end test changes

---

## Reviews

### Review by @Andrewgdewar - Approved on February 18, 2025 at 07:10 PM UTC

---

*Archived from: https://github.com/content-services/content-sources-backend/pull/981*
