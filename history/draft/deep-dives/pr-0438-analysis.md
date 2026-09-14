# backend#438 — Fixes 2896: refactor pulp client in dao

Author: @rverdile  
Merged: 2023-11-16  
URL: https://github.com/content-services/content-sources-backend/pull/438  
Period: pulp-integration

## Problem

After #331, DAOs needed a domain-scoped Pulp client per request. Ryan’s first helper (in #414) set the domain client in a way that was easy to call with the wrong context or the wrong domain. Tests and handlers leaked `context.Background()` into domain-scoped calls.

## Approach

Two chained methods: construct the client once at dependency injection with `context.Background()`, then at request time `pulpClient.WithContext(ctx).WithDomain(domainName).GetContentPath()`. The underlying struct is reused; only context and domain change.

The first iteration commented out failing tests while the pattern was proven. @jlsherrill approved pending CI.

## Design decisions

- Immutable-style chain, not a setter on a singleton. Avoids one request’s domain sticking on the next.
- DAO owns Pulp, handlers do not. Snapshot, template, and later Lightwell import code all go through this client. That is the “shape Lightwell still uses” from the milestone list.

## Impact

A plumbing PR, but it is the last 2023 change that made Pulp-from-the-DAO safe. Without it, every new content type would reintroduce the #414 bug. No user-facing API change.

## Quote

> “The method I introduced [in #414] to set the domain scoped pulp client is bug-prone. … Now, we can initialize the pulp client with `context.Background()` during dependency injection. Then, replace the context later, during the request.” — @rverdile
