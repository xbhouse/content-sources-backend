# backend#331 — Fixes 1939: pulp domain integration

Author: @jlsherrill  
Merged: 2023-08-04  
URL: https://github.com/content-services/content-sources-backend/pull/331  
Period: pulp-integration

## Problem

Early snapshots lived in a shared Pulp namespace. Two Insights orgs snapshotting the same URL (or colliding on paths) had no isolation. Pulp 3 domains give each tenant a prefix on API and content paths. content-sources needed to create a domain per org and send every subsequent Pulp call through it.

## Approach

On repo create/update with snapshotting, a goroutine creates the Pulp domain if missing. The Pulp client split into global (no domain — domain CRUD itself) and domain-scoped interfaces over the same struct. Snapshot `repository_path` gained a domain prefix (`{domain}/{repo_uuid}/{snapshot_uuid}`). Local compose switched to the team’s own Pulp compose; ephemeral Pulp had to set `DOMAIN_ENABLED: "true"`.

@rverdile hit 500s on bulk-create for a new org (concurrent domain create). Justin treated “create at the same time” as an edge case to swallow. @swadeley confirmed two orgs produced two domain prefixes (`2d18751c` vs `12d903db`).

## Design decisions

- Domain name is not the org_id. Paths use a short Pulp domain identifier. Org mapping stays in Postgres.
- Create domain lazily on first snapshot, not at org signup. Orgs that never snapshot never touch Pulp.
- Global vs scoped client. Domain bootstrap cannot use a domain-scoped client; everything after must. #438 later made that chaining (`WithContext` / `WithDomain`) less error-prone.
- S3/object storage is part of the same change. Local minio and `custom_repo_objects` config landed here; ACL/hostname follow-ups were #355 and #395.

## Impact

Multi-tenant Pulp is a hard requirement for RH-hosted custom repos. Stage later needed repair jobs (#478–#492) because leftover pre-domain snapshots were wrong. Lightwell import (#1559) still creates org-scoped rows that sit on this domain model — and immediately tripped org-id middleware because imported Lightwell orgs do not match the caller.

## Quote

> “The domain is still created and no actual functionality is disrupted, but I'm consistently getting `Error creating domain error=\"500 Internal Server Error\"` when using bulk create with a new orgID.” — @rverdile, 2023-07-27
