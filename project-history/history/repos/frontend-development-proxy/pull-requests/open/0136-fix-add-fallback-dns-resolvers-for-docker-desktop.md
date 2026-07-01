---
type: pull_request
number: 136
title: "fix: add fallback DNS resolvers for Docker Desktop"
state: open
author: platex-rehor-bot
created: 2026-05-07T16:53:13Z
updated: 2026-05-07T16:53:13Z
base_branch: main
head_branch: bot/RHCLOUD-46377
labels: []
url: https://github.com/RedHatInsights/frontend-development-proxy/pull/136
---

# Pull Request #136: fix: add fallback DNS resolvers for Docker Desktop

**Author**: @platex-rehor-bot
**Created**: May 07, 2026 at 04:53 PM UTC
**Status**: Open
**Labels**: None
**Base**: `main` ← **Head**: `bot/RHCLOUD-46377`

## Description

## Summary

Fixes RHCLOUD-46377 — `fec dev-proxy` broken for `prod.foo.redhat.com` due to Docker DNS resolution failures.

- Adds fallback DNS resolvers (8.8.8.8, 1.1.1.1) to `/etc/resolv.conf` at container startup
- Only modifies resolv.conf when writable and resolvers not already present
- Prevents intermittent DNS failures in Docker Desktop on macOS after sleep, network switch, or VPN toggle

## Test plan

- [ ] Run `fec dev-proxy` on Docker Desktop (macOS) after a VPN toggle or sleep/wake cycle
- [ ] Verify Caddy successfully proxies to `console.redhat.com`
- [ ] Confirm fallback resolvers are only added when `/etc/resolv.conf` is writable
- [ ] Verify no duplicate resolvers on repeated container starts

RHCLOUD-46377

---

*Archived from: https://github.com/RedHatInsights/frontend-development-proxy/pull/136*
