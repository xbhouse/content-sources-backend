# backend#570 — Fixes 3491: bootstrap Candlepin connection (context)

Author: @jlsherrill  
Merged: 2024-02-28  
URL: https://github.com/content-services/content-sources-backend/pull/570  
Period: content-sources-expansion  
Weight: one-pager

## Problem

A template that only exists in content-sources cannot attach to a RHEL system. Candlepin environments are how Insights/RHSM consume a content set. The backend needed a Candlepin client and a safe way to develop against it.

## Approach

Bootstrap a Candlepin client. Compose creates a devel org and imports a manifest. Stage/prod use the user’s real org; local uses `devel_org`. `go run cmd/candlepin/main.go list-contents` verifies the devel org. @rverdile tested both commands.

Content-for-templates (#639), env path overrides (#670), and GPG/modular hotfixes (#682) are the functional follow-ons.

## Why it matters here

Candlepin is an *entitlement* side effect of templates, which are a *use* of snapshots. Lightwell does not use this path. Mention it so Pulp → template → environment is one sentence, then stop.

## Quote

> “in stage/prod we will use the existing orgs that match the user's org. In dev, all actions will use this created 'devel org'.” — @jlsherrill
