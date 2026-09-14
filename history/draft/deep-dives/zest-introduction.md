# zest: generated Pulp Go bindings

Repo: [content-services/zest](https://github.com/content-services/zest)  
Started: 2023-03-09 by @Andrewgdewar (`Init`)  
Period: pulp-integration

## Problem

content-sources needed to talk to Pulp’s HTTP API (remotes, repositories, sync, publications, distributions). Hand-writing that client would have lagged Pulp’s OpenAPI and broken on every Pulp upgrade. The team needed a Go client that could be regenerated whenever Pulp’s spec changed.

## Approach

@Andrewgdewar created zest as a separate repository whose job is to generate Go bindings from Pulp’s OpenAPI and publish them as a module the backend imports. Most zest history is automated `Update pulp bindings to <sha>` commits — those are excluded from this analysis. The interesting work is the generation pipeline itself: which Pulp image to generate from, how to version the module, and how to paper over Pulp spec bugs (enum fixes, operation-id collisions, `TaskResponse.Error` type).

@jlsherrill and @rverdile later patched the generator (retry on 404s, UBI image, domain flags, operation-id renames). @TenSt took over binding updates in 2026.

## Design decisions

- Separate repo, not a package inside the backend. Binding churn (hundreds of commits) would have drowned product history. The backend pins a zest version and moves on.
- Generate from Pulp, do not fork the API. When Pulp’s spec was wrong, zest patched the generated types rather than wrapping Pulp by hand (see @rverdile’s `TaskResponse.Error` → string change, 2023-07).
- Versioning experiments. Mid-2023 the team tried `v$year` module versions and Pulp content checksums as versions, then settled back on a conventional module path. That thrash is visible in zest’s March–June 2023 log.

## Impact

Every snapshot, domain, and later Lightwell query that hits Pulp goes through zest. When Pulp domains landed (backend#331), the client had to expose a domain-less “global” interface and a domain-scoped one. Lightwell Maven/Python queries still use the same generated types.

## Quote

There is no single design PR; the closest artifact is the March 2023 commit storm while @Andrewgdewar got GitHub Actions to tag and release (`Finally!!!`). The product decision is implicit: Pulp’s API is the source of truth; zest is a generated adapter.
