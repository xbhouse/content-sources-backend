# tang: Pulp content search service

Repo: [content-services/tang](https://github.com/content-services/tang)  
Founded: 2023-11-30 by @rverdile  
First real PR: tang#1, merged 2024-01-03 — “Add rpm name search and integration testing”  
Period: pulp-integration → expansion

## Problem

Listing RPMs (and later errata, module streams, Maven, Python) against Pulp/Postgres from the API process is heavy and couples content-sources-backend to Pulp’s query shape. The backend already had zest for *mutating* Pulp. It needed a read-side service that could search a repository version by name given Pulp hrefs.

## Approach

@rverdile created tang with an exported `Tangy` interface. tang#1 added `RpmRepositoryVersionPackageSearch()`, zerolog, a local Pulp compose, and CI integration tests. @jlsherrill asked for the mock to live in tang, not in every consumer — Ryan added it.

The backend calls tang for snapshot content lists. Errata followed as tang#6 + backend#608 (Apr 2024). Module streams as tang#14 (Dec 2024). Maven/Python as tang#25/#26 (Jul 2026) for Lightwell.

## Design decisions

- Library-shaped Go module, not an HTTP microservice. The backend imports tang and points it at Pulp’s database. That is why backend#629 had to pass the Clowder RDS CA for “tangy db.”
- Interface + mock in tang. Keeps content-sources-backend tests from inventing a second fake.
- Content-type-agnostic list/detail later. #1560 / tang#25 deliberately made package APIs not RPM-specific (`group` may be blank for Python). That reuse is why Lightwell did not need a fourth repo.

## Impact

tang is the current read path for anything inside a Pulp repository version. Snapshots without tang are just a distribution URL. With tang they are searchable packages, errata, and (in 2026) Maven GAVs. Lightwell package pages are tang queries with a different content type today. That Lightwell use is scheduled to move onto Pulp plugin catalog APIs ([pulp_maven#451](https://github.com/pulp/pulp_maven/pull/451), [pulp_python#1359](https://github.com/pulp/pulp_python/pull/1359)); those PRs are open.

## Quote

> “Also, would it make sense to include a Mock of the interface here? Instead of in our application.” — @jlsherrill on tang#1
