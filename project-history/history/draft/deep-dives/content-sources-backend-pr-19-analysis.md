---
repo: content-sources-backend
pr_number: 19
period: expansion
source: history/repos/content-sources-backend/pull-requests/merged/
---

# Deep Dive: content-sources-backend PR #19

## Problem / Motivation

With database models in place (PR #4), CONTENT-40 required the first customer-facing API: create a content source (repository configuration) with org-scoped isolation.

## Solution Approach

@rverdile established the **api/dao/handler** layering: API structs for request/response binding, dao layer for DB operations, handlers wire HTTP to dao. Create endpoint with OpenAPI comments, unique constraint on org ID + URL pair, identity via `x-rh-identity` header.

## Key Design Decisions

- **Layered architecture**: api/dao separation became convention for all endpoints.
- **Platform identity**: Uses platform-go-middlewares identity parsing — same as other Insights services.
- **Org-scoped uniqueness**: URL unique per org, not globally.

## Impact

Architectural template for 1,500+ subsequent PRs. Every endpoint follows the pattern introduced here.

## Review Notes

Merged Jun 2022 alongside frontend bootstrap — content-sources "go live" month.

---

[← Project History](../../PROJECT_HISTORY.md) · [Validation report](../validation_report.md)
